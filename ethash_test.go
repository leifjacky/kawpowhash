// Copyright 2015 The go-ethereum Authors
// Copyright 2015 Lefteris Karapetsas <lefteris@refu.co>
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package kawpowhash

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// glog.SetV(6)
	// glog.SetToStderr(true)
}

type testBlock struct {
	difficulty  *big.Int
	hashNoNonce common.Hash
	nonce       uint64
	mixDigest   common.Hash
	number      uint64
	hash        common.Hash
}

func (b *testBlock) Difficulty() *big.Int     { return b.difficulty }
func (b *testBlock) HashNoNonce() common.Hash { return b.hashNoNonce }
func (b *testBlock) Nonce() uint64            { return b.nonce }
func (b *testBlock) MixDigest() common.Hash   { return b.mixDigest }
func (b *testBlock) NumberU64() uint64        { return b.number }

var validBlocks = []*testBlock{
	{
		number:      1727606,
		hashNoNonce: common.HexToHash("3e973c133bf9e79233037e9b3d57f4d7a8920bff3b8d771168909b8ebf79d102"),
		difficulty:  big.NewInt(1532671),
		nonce:       0x7693a100059d3d14,
		mixDigest:   common.HexToHash("f112e8acfe8276ae546b021f3f9b102a160b214a8bd0d3f886994090d2d6f6c2"),
		hash:        common.HexToHash("000000c5695bc9e6bdb3130a163e8ca7b1407d0befbe1fd3c89f785e08bf2e44"),
	},
}

func TestEthashVerifyValid(t *testing.T) {
	eth := New()
	for i, block := range validBlocks {
		if !eth.Verify(block) {
			t.Errorf("block %d (%x) did not validate.", i, block.hashNoNonce[:6])
		}
	}
}

func TestCompute(t *testing.T) {
	eth := New()
	for _, block := range validBlocks {
		_, mixDigest, hash := eth.Compute(block.hashNoNonce, block.nonce, block.number)
		if mixDigest != block.mixDigest {
			t.Errorf("invalid mix digest %064x", mixDigest)
		}
		if hash != block.hash {
			t.Errorf("invalid hash %064x", hash)
		}
	}
}
