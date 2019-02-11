/*
Package bloomfilter was created for learning purposes
https://github.com/hugoluchessi/bloomfilter

Copyright © 2019 Hugo Luchessi

MIT license

All theory was based on the following source:
https://en.wikipedia.org/wiki/Bloom_filter

this bloom filter assumes:
- n (uint64): as the size of the bloom filter, the max number of elements contained in it
- k (uint64): number of hashing functions
- p (float64): probability of false positives
- m (uint64): size of the bit array in which elements will be placed

*/

package bloomfilter

import (
	"bytes"
	"encoding/binary"
	"hash"
	"math"
	"math/rand"

	"github.com/OneOfOne/xxhash"
	"github.com/hugoluchessi/bitarray"
	"github.com/hugoluchessi/goerrors"
)

const (
	invalidNErrorMessage = "Paramater [n] should be greater than zero."
	invalidPErrorMessage = "Paramater [p] should be greater than zero."
)

/*
BloomFilter is the main structure, hold the bit array with elements and controls number of
hashing funcions and max size.

It is thread unsafe on purpose, it may be used as value, and application
layer must control the distribution of its reference so it can stay in stack
instead of


k: number of hashing algorithms to be used
n: max number of elements
m: number of bits to use to persist elements
a: bit array to contain elements
hashes: hashes which content will be hashed before persisting
*/
type BloomFilter struct {
	k uint64
	n uint64
	m uint64
	a bitarray.BitArray
	h []hash.Hash
}

/*
CreateBloomFilter creates a filter based on:
n: max number of elements expected
p: probability of false positives for the given n
*/
func CreateBloomFilter(n uint64, p float64) (BloomFilter, error) {
	if n == 0 {
		return BloomFilter{}, NewInvalidArgumentError(invalidNErrorMessage, "n", n)
	}

	if p <= 0 {
		return BloomFilter{}, NewInvalidArgumentError(invalidPErrorMessage, "p", p)
	}

	m := optimalM(n, p)
	k := optimalK(n, m)
	a, err := bitarray.NewBitArray(m)

	if err != nil {
		return BloomFilter{}, goerrors.WrapError(err)
	}

	h := make([]hash.Hash, k)

	var i uint64

	for i = 0; i < k; i++ {
		h[i] = xxhash.NewS64(rand.Uint64())
	}

	return BloomFilter{k, n, m, a, h}, nil
}

func (bf BloomFilter) Add(v []byte) error {
	indexes := bf.bitIndexesToCheck(v)

	for _, i := range indexes {
		err := bf.a.TurnOn(i)

		if err != nil {
			return goerrors.WrapError(err)
		}
	}

	return nil
}

func (bf BloomFilter) Contains(v []byte) (bool, error) {
	indexes := bf.bitIndexesToCheck(v)

	for _, i := range indexes {
		isSet, err := bf.a.IndexValue(i)

		if err != nil {
			return false, goerrors.WrapError(err)
		}

		if !isSet {
			return false, nil
		}
	}

	return true, nil
}

func (bf BloomFilter) bitIndexesToCheck(v []byte) []uint64 {
	indexes := make([]uint64, bf.k)

	for i, hash := range bf.h {
		var hashValue uint64
		hashBinary := hash.Sum(v)
		buf := bytes.NewBuffer(hashBinary)
		binary.Read(buf, binary.LittleEndian, &hashValue)

		bitToBeSet := hashValue % bf.m
		indexes[i] = bitToBeSet
	}

	return indexes
}

/*
The optimal value for k varies regading the size of the bit array and the max number
elements expected fot the filter, maxed by the number of selected algorithms

Formulae explanation https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
*/
func optimalK(n, m uint64) uint64 {
	floatK := math.Ceil(float64(m) * math.Ln2 / float64(n))
	return uint64(floatK)
}

/*
The optimal value for m (size of the bit array) varies regading the expected number of
hashed keys (number of elements in the bloom filter) and the probability of collision,
both parameters for constructing the bloom filter

Formulae explanation https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
*/
func optimalM(n uint64, p float64) uint64 {
	floatM := math.Ceil(-float64(n) * math.Log(p) / (math.Ln2 * math.Ln2))
	return uint64(floatM)
}
