package bloomfilter_test

import (
	"crypto/rand"
	"testing"

	"github.com/hugoluchessi/bloomfilter"
)

func TestCreateBloomFilter(t *testing.T) {
	var validN uint64 = 1000
	var validP float64 = 0.1

	t.Run("Invalid P", func(t *testing.T) {
		_, err := bloomfilter.CreateBloomFilter(validN, 0)

		if err == nil {
			t.Error("Must return error")
		}

		_, err = bloomfilter.CreateBloomFilter(validN, -0.1)

		if err == nil {
			t.Error("Must return error")
		}
	})

	t.Run("Invalid N", func(t *testing.T) {
		_, err := bloomfilter.CreateBloomFilter(0, validP)

		if err == nil {
			t.Error("Must return error")
		}
	})

	t.Run("Valid Parameters", func(t *testing.T) {
		_, err := bloomfilter.CreateBloomFilter(validN, validP)

		if err != nil {
			t.Error("Must not return error")
		}
	})
}

func TestAddAndContains(t *testing.T) {
	var validN uint64 = 1000
	var validP float64 = 0.1

	v1 := []byte{1, 2, 3, 4, 5}
	v2 := []byte{11, 22, 33, 44, 55}

	bf, _ := bloomfilter.CreateBloomFilter(validN, validP)

	t.Run("Test adding a value", func(t *testing.T) {
		err := bf.Add(v1)

		if err != nil {
			t.Error("Must not return error")
		}

		isSet, err := bf.Contains(v1)

		if err != nil {
			t.Error("Must not return error")
		}

		if !isSet {
			t.Error("Value must have been set")
		}

		isSet, err = bf.Contains(v2)

		if err != nil {
			t.Error("Must not return error")
		}

		if isSet {
			t.Error("Value must not have been set")
		}
	})
}

func BenchmarkAddN1000K01Bytes1(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000K01Bytes1000(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000K01Bytes1000000(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000000)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000K00001Bytes1(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000K00001Bytes1000(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000K00001Bytes1000000(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000000)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000000K01Bytes1(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000000K01Bytes1000(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000000K01Bytes1000000(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000000)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000000K00001Bytes1(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000000K00001Bytes1000(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkAddN1000000K00001Bytes1000000(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000000)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Add(bytes)
	}
}

func BenchmarkContainsN1000K01Bytes1(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000K01Bytes1000(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000K01Bytes1000000(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000000)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000K00001Bytes1(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000K00001Bytes1000(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000K00001Bytes1000000(b *testing.B) {
	var n uint64 = 1000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000000)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000000K01Bytes1(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000000K01Bytes1000(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000000K01Bytes1000000(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.1

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000000)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000000K00001Bytes1(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000000K00001Bytes1000(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func BenchmarkContainsN1000000K00001Bytes1000000(b *testing.B) {
	var n uint64 = 1000000
	var p float64 = 0.0001

	bf, _ := bloomfilter.CreateBloomFilter(n, p)
	bytes := createBytes(1000000)
	bf.Add(bytes)

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bf.Contains(bytes)
	}
}

func createBytes(s uint32) []byte {
	bytes := make([]byte, s)
	rand.Read(bytes)
	return bytes
}
