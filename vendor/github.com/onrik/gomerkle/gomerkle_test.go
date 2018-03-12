package gomerkle

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"testing"

	"github.com/stretchr/testify/require"
)

// SimpleHash: does nothing
var SimpleHashData []byte

type SimpleHash struct{}

func NewSimpleHash() hash.Hash {
	return SimpleHash{}
}

func (h SimpleHash) Write(p []byte) (int, error) {
	size := h.Size()
	datalen := (len(p) / size) * size
	if len(p) == 0 || len(p)%size != 0 {
		datalen += size
	}
	data := make([]byte, datalen)
	copy(data, p)

	block := make([]byte, size)
	copy(block, data[:size])
	for i := 1; i < len(data)/size; i++ {
		_block := data[i*size : (i+1)*size]
		for j, c := range _block {
			block[j] += c
		}
	}

	SimpleHashData = append(SimpleHashData, block...)
	return size, nil
}
func (h SimpleHash) Sum(p []byte) []byte {
	p = append(p[:], SimpleHashData[:]...)
	return p
}
func (h SimpleHash) Reset() {
	SimpleHashData = nil
}
func (h SimpleHash) Size() int {
	return 32
}
func (h SimpleHash) BlockSize() int {
	return 32
}

type NotHash struct{}

func NewNotHash() hash.Hash {
	return NotHash{}
}
func (h NotHash) Write(p []byte) (int, error) {
	return 32, nil
}
func (h NotHash) Sum(p []byte) []byte {
	return p
}
func (h NotHash) Reset() {
}
func (h NotHash) Size() int {
	return 32
}
func (h NotHash) BlockSize() int {
	return 32
}

// FailingHash: always returns error on Write
type FailingHash struct {
	attempts int
}

func NewFailingHashAt(n int) *FailingHash {
	return &FailingHash{
		attempts: n,
	}
}

func NewFailingHash() *FailingHash {
	return NewFailingHashAt(0)
}

func (h *FailingHash) Write(p []byte) (int, error) {
	h.attempts--
	if h.attempts <= 0 {
		return 0, errors.New("Failed to write hash")
	}

	return 0, nil
}
func (h *FailingHash) Sum(p []byte) []byte {
	return p
}
func (h *FailingHash) Reset() {
}
func (h *FailingHash) Size() int {
	return 0
}
func (h *FailingHash) BlockSize() int {
	return 0
}

func failNotEqual(t *testing.T, label string, input interface{},
	expect interface{}, result interface{}) {
	t.Errorf("%s(%v) != %v (%v, instead)", label, input, expect, result)
}

/* Utils */

func TestCalculateTreeHeight(t *testing.T) {
	inputs := [][]uint64{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 3},
		{5, 4},
		{6, 4},
		{7, 4},
		{8, 4},
		{9, 5},
		{15, 5},
		{16, 5},
		{17, 6},
		{31, 6},
		{32, 6},
		{63, 7},
		{64, 7},
		{65, 8},
	}
	for _, i := range inputs {
		r := calculateTreeHeight(i[0])
		require.Equal(t, i[1], r)

		r, _ = calculateHeightAndNodeCount(i[0])
		require.Equal(t, i[1], r)
	}
}

func TestLog2(t *testing.T) {
	inputs := [][]uint64{
		{0, 0},
		{1, 0},
		{2, 1},
		{4, 2},
		{8, 3},
		{16, 4},
		{32, 5},
		{64, 6},
	}
	for _, i := range inputs {
		r := log2(i[0])
		require.Equal(t, i[1], r)
	}
}

func TestCalculateNodeCount(t *testing.T) {
	inputs := [][]uint64{
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 7},
		{9, 20},
		{10, 21},
		{11, 23},
		{12, 24},
		{13, 27},
		{21, 44},
		{22, 45},
	}
	for _, i := range inputs {
		height := calculateTreeHeight(i[0])
		r := calculateNodeCount(height, i[0])
		require.Equal(t, i[1], r)

		_, r = calculateHeightAndNodeCount(i[0])
		require.Equal(t, i[1], r)
	}
}

func TestNextPowerOfTwo(t *testing.T) {
	inputs := [][]uint64{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 4},
		{5, 8},
		{8, 8},
		{14, 16},
		{16, 16},
		{65535, 65536},
		{65536, 65536},
		{65537, 131072},
	}
	for _, i := range inputs {
		r := nextPowerOfTwo(i[0])
		require.Equal(t, i[1], r)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	type powerOfTwoResult struct {
		input  uint64
		output bool
	}
	inputs := []powerOfTwoResult{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{16, true},
		{65534, false},
		{65535, false},
		{65536, true},
		{65537, false},
		{2032131433, false},
	}
	for _, i := range inputs {
		r := isPowerOfTwo(i.input)
		require.Equal(t, i.output, r)
	}
}

/* Tree */

func containsNode(nodes []Node, node *Node) bool {
	/* Returns trueif a *Node is in a []Node */
	for i := 0; i < len(nodes); i++ {
		if node == &nodes[i] {
			return true
		}
	}
	return false
}

func createDummyTreeData(count, size int, useRand bool) [][]byte {
	/* Creates an array of bytes with nonsense in them */
	data := make([][]byte, count)
	for i := 0; i < count; i++ {
		garbage := make([]byte, size)
		if useRand {
			read := 0
			for read < size {
				n, _ := rand.Read(garbage[read:])
				read += n
			}
		} else {
			for i := 0; i < size; i++ {
				garbage[i] = byte((i + 1) % 0xFF)
			}
		}
		data[i] = garbage
	}
	return data
}

func verifyGeneratedTree(t *testing.T, tree *Tree) {
	/* Given a generated tree, confirm its state is correct */

	// Nodes should have been created
	require.NotNil(t, tree.Nodes)
	require.Equal(t, len(tree.Nodes), cap(tree.Nodes),
		"tree.Nodes len should equal its cap")

	for i := tree.Height() - 1; i > 0; i-- {
		// All the other nodes should have children, and their children
		// should be in the deeper level
		deeper := tree.GetNodesAtHeight(i + 1)
		row := tree.GetNodesAtHeight(i)
		for j, n := range row {
			require.NotNil(t, n.Left, "Left child should never be nil")
			require.Equal(t, n.Left, &deeper[j*2])
			if j == len(row)-1 && len(deeper)%2 == 1 {
				// Last node in this level should have nil right child
				// if its unbalanced
				require.Nil(t, n.Right)
				// Its hash should be the same as the left node hash
				require.Equal(t, n.Left.Hash, n.Hash,
					"Left child hash should equal node hash when right child is nil")
			} else {
				require.NotNil(t, n.Right)
				require.Equal(t, n.Right, &deeper[j*2+1])
				require.NotEqual(t, bytes.Equal(n.Right.Hash, n.Hash), true,
					"Right child hash should not equal node hash")
				require.NotEqual(t, bytes.Equal(n.Left.Hash, n.Hash), true,
					"Left child hash should not equal node hash")
			}
		}

		// Each row should have prev/2 + prev%2 nodes
		prev := len(deeper)
		require.Equal(t, prev/2+prev%2, len(row))
	}

	rootRow := tree.GetNodesAtHeight(1)
	// The root row should exist
	require.NotNil(t, rootRow)

	// The root row should be of length 1
	require.Equal(t, 1, len(rootRow))

	// the Root() should be the only item in the top row
	require.Equal(t, rootRow[0].Hash, tree.Root())
}

func verifyInitialState(t *testing.T, tree *Tree) {
	require.Nil(t, tree.Nodes)
	require.Nil(t, tree.Levels)
}

func TestNewNode(t *testing.T) {
	h := NewSimpleHash()
	block := createDummyTreeData(1, h.Size(), true)[0]
	n, err := NewNode(h, block)
	require.Nil(t, err)
	require.True(t, bytes.Equal(n.Hash, block))

	// Any nil argument should return blank node, no error
	n, err = NewNode(nil, nil)
	require.Nil(t, err)
	require.Nil(t, n.Hash)

	n, err = NewNode(h, nil)
	require.Nil(t, err)
	require.Nil(t, n.Hash)

	// Hashed data
	n, err = NewNode(nil, block)
	require.Nil(t, err)
	require.Equal(t, block, n.Hash)

	// Check hash error handling
	h = NewFailingHash()
	n, err = NewNode(h, block)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "Failed to write hash")
}

func TestNewTree(t *testing.T) {
	tree := NewTree(NewSimpleHash())
	verifyInitialState(t, &tree)
}

func TestTreeUngenerated(t *testing.T) {
	tree := NewTree(NewSimpleHash())
	// If data is nil, it should handle that:
	err := tree.Generate()
	require.NotNil(t, err)
	require.Equal(t, "Empty tree", err.Error())
	require.Nil(t, tree.Root())
	require.Equal(t, uint64(0), tree.Height())
	require.Nil(t, tree.Nodes)
	require.Nil(t, tree.GetLeaf(0))
}

func TestTreeGenerate(t *testing.T) {
	tree := NewTree(NewSimpleHash())
	// Setup some dummy data
	blockCount := 13
	blockSize := 16
	data := createDummyTreeData(blockCount, blockSize, true)

	// Generate the tree
	tree.AddData(data...)
	err := tree.Generate()
	require.Nil(t, err)
	verifyGeneratedTree(t, &tree)

	for i := range data {
		hash := tree.hash(data[i])
		require.Equal(t, hash, tree.GetLeaf(i))
	}

	// Generating with no blocks should return error
	tree = NewTree(NewSimpleHash())
	tree.AddData(make([][]byte, 0, 1)...)
	err = tree.Generate()
	require.NotNil(t, err)
	require.Equal(t, "Empty tree", err.Error())
}

func TestGenerateFailedHash(t *testing.T) {
	tree := NewTree(NewFailingHash())
	data := createDummyTreeData(16, 16, true)
	// Fail hash during the leaf generation
	tree.AddData(data...)
	err := tree.Generate()
	require.NotNil(t, err)
	require.Equal(t, "Failed to write hash", err.Error())

	// Fail hash during internal node generation
	data = createDummyTreeData(16, 16, true)
	tree = NewTree(NewFailingHashAt(20))
	tree.AddData(data...)
	err = tree.Generate()
	require.NotNil(t, err)
	require.Equal(t, "Failed to write hash", err.Error())
}

func TestGenerateWithOneNode(t *testing.T) {
	value := []byte("value")
	tree := NewTree(md5.New())
	tree.AddData(value)

	err := tree.Generate()
	require.Nil(t, err)

	proof := tree.GetProof(0)
	require.Equal(t, 0, len(proof))
	require.Equal(t, tree.hash(value), tree.Root())
}

func TestGetNodesAtHeight(t *testing.T) {
	// ungenerate tree should return nil
	tree := NewTree(NewSimpleHash())
	require.Nil(t, tree.GetNodesAtHeight(1))

	count := 15
	size := 16
	data := createDummyTreeData(count, size, true)
	tree.AddData(data...)
	tree.Generate()
	verifyGeneratedTree(t, &tree)

	// invalid height should return nil
	require.Nil(t, tree.GetNodesAtHeight(0))
	require.Nil(t, tree.GetNodesAtHeight(tree.Height()+1))

	// check valid height = 1
	nodes := tree.GetNodesAtHeight(tree.Height())
	require.Equal(t, len(nodes), count)
	expect := tree.Nodes[:count]
	for i := 0; i < len(nodes); i++ {
		require.Equal(t, &expect[i], &nodes[i])
	}
}

// Returns the root hash for an array of hashes
func simpleMerkle(data [][]byte) []byte {
	h := sha256.New()
	// Build the leaves
	h0 := make([][]byte, len(data))
	for i, b := range data {
		h.Reset()
		h.Write(b)
		h0[i] = h.Sum(nil)
	}

	h1 := make([][]byte, (len(h0)+len(h0)%2)/2)
	for {
		for i := 0; i < len(h0); i += 2 {
			var sum []byte
			if len(h0)%2 == 1 && i == len(h0)-1 {
				sum = h0[i]
			} else {
				c := append(h0[i], h0[i+1]...)
				h.Reset()
				h.Write(c)
				sum = h.Sum(nil)
			}
			h1[i/2] = sum
		}
		if len(h1) == 1 {
			break
		}
		h0 = h1
		h1 = make([][]byte, (len(h0)+len(h0)%2)/2)
	}
	return h1[0]
}

func TestRootHashValue(t *testing.T) {
	// Check the root hash made by Tree against a simpler implementation
	// that finds only the root hash

	tree := NewTree(sha256.New())
	// Setup some dummy data
	blockCount := 16
	blockSize := 16
	data := createDummyTreeData(blockCount, blockSize, true)

	// Generate the tree
	tree.AddData(data...)
	err := tree.Generate()
	require.Nil(t, err)
	verifyGeneratedTree(t, &tree)

	// Calculate the root hash with the simpler method
	merk := simpleMerkle(data)

	require.True(t, bytes.Equal(tree.Root(), merk))
}

func TestGetProof(t *testing.T) {
	data := make([][]byte, 20)
	for i := 0; i < len(data); i++ {
		data[i] = []byte(fmt.Sprintf("value%d", i))
	}

	tree := NewTree(md5.New())
	tree.AddData(data...)
	err := tree.Generate()

	require.Nil(t, err)

	proof := tree.GetProof(5)
	expected := Proof{
		map[string][]byte{
			"left": {182, 0, 252, 107, 110, 161, 149, 93, 17, 72, 97, 196, 41, 52, 198, 89},
		},
		map[string][]byte{
			"right": {254, 179, 55, 70, 100, 189, 9, 46, 252, 26, 168, 125, 179, 98, 109, 86},
		},
		map[string][]byte{
			"left": {98, 40, 179, 120, 254, 245, 143, 248, 208, 179, 3, 215, 114, 144, 115, 87},
		},
		map[string][]byte{
			"right": {86, 22, 191, 37, 72, 68, 16, 116, 208, 179, 61, 82, 214, 0, 3, 226},
		},
		map[string][]byte{
			"right": {37, 153, 197, 153, 191, 14, 1, 134, 75, 185, 94, 199, 7, 212, 201, 178},
		},
	}
	require.Equal(t, expected, proof)
}

func TestVerifyProof(t *testing.T) {
	data := make([][]byte, 10)
	for i := 0; i < len(data); i++ {
		data[i] = []byte(fmt.Sprintf("zzz%d", i))
	}

	tree := NewTree(md5.New())
	tree.AddData(data...)
	err := tree.Generate()
	require.Nil(t, err)

	root := tree.Root()
	for i := 0; i < len(data); i++ {
		proof := tree.GetProof(i)
		leaf := tree.GetLeaf(i)
		require.True(t, tree.VerifyProof(proof, root, leaf))

	}

	// Test verify with empty tree
	tree = NewTree(md5.New())
	proof := Proof{
		map[string][]byte{
			"left": {182, 0, 252, 107, 110, 161, 149, 93, 17, 72, 97, 196, 41, 52, 198, 89},
		},
		map[string][]byte{
			"right": {254, 179, 55, 70, 100, 189, 9, 46, 252, 26, 168, 125, 179, 98, 109, 86},
		},
		map[string][]byte{
			"left": {98, 40, 179, 120, 254, 245, 143, 248, 208, 179, 3, 215, 114, 144, 115, 87},
		},
		map[string][]byte{
			"right": {86, 22, 191, 37, 72, 68, 16, 116, 208, 179, 61, 82, 214, 0, 3, 226},
		},
		map[string][]byte{
			"right": {37, 153, 197, 153, 191, 14, 1, 134, 75, 185, 94, 199, 7, 212, 201, 178},
		},
	}
	root = []byte{198, 171, 151, 198, 99, 102, 171, 160, 114, 231, 230, 66, 133, 203, 93, 244}
	leaf := []byte{0x8a, 0xfb, 0x7a, 0xa8, 0xa1, 0xa0, 0xdd, 0x7b, 0xd7, 0x7c, 0xa8, 0x85, 0x7b, 0xcc, 0x42, 0x2d}
	require.True(t, tree.VerifyProof(proof, root, leaf))

	// Test invalid proof
	proof = Proof{
		map[string][]byte{
			"left": {182, 0, 252, 107, 110, 161, 149, 93, 17, 72, 97, 196, 41, 52, 198, 89},
		},
		map[string][]byte{
			"right": {254, 179, 55, 70, 100, 189, 9, 46, 252, 26, 168, 125, 179, 98, 109, 86},
		},
		map[string][]byte{
			"left": {98, 40, 179, 120, 254, 245, 143, 248, 208, 179, 3, 215, 114, 144, 115, 87},
		},
		map[string][]byte{
			"right": {86, 22, 191, 37, 72, 68, 16, 116, 208, 179, 61, 82, 214, 0, 3, 226},
		},
		map[string][]byte{
			"left": {37, 153, 197, 153, 191, 14, 1, 134, 75, 185, 94, 199, 7, 212, 201, 178},
		},
	}
	require.False(t, tree.VerifyProof(proof, root, leaf))

	proof = Proof{
		map[string][]byte{
			"left": {182, 0, 252, 107, 110, 161, 149, 93, 17, 72, 97, 196, 41, 52, 198, 89},
		},
		map[string][]byte{},
		map[string][]byte{
			"right": {86, 22, 191, 37, 72, 68, 16, 116, 208, 179, 61, 82, 214, 0, 3, 226},
		},
	}
	require.False(t, tree.VerifyProof(proof, root, leaf))
}
