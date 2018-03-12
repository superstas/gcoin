// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package amount

const (
	// SatoshiPerBitcent is the number of satoshi in one gophercoin cent.
	SatoshiPerBitcent = 1e6

	// GoshPerGcoin is the number of gosh in one gophercoin (1 GOC).
	GoshPerGcoin = 1e8

	// MaxGosh is the maximum transaction amount allowed in gosh.
	MaxGosh = 21e6 * GoshPerGcoin
)
