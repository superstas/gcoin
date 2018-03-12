package keys

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewKeys(t *testing.T) {
	k, err := New()
	require.Nil(t, err)
	assert.NotNil(t, k.PrivateKey)
	assert.NotNil(t, k.PublicKey)
	assert.NotNil(t, k.AddressPKH)
}

func TestKeysJSON(t *testing.T) {
	k, err := New()
	require.Nil(t, err)
	data, err := json.Marshal(k)
	require.Nil(t, err)
	assert.NotNil(t, data)

	var k2 Keys
	err = json.Unmarshal(data, &k2)
	require.Nil(t, err)
	assert.Equal(t, k, &k2)
}
