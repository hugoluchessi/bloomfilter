# bloomfilter
Simple bloom filter data structure implementation

Bloom filter is a memory efficient data structure to persist elemets within a set. Specific for looking if an item **IS NOT**  in a set.

It uses [OneOfOne/xxhash](https://github.com/OneOfOne/xxhash) hash algorithm for greater performance and randomness.

## Good and bad:
A Bloom filter ensures an element is NOT in a set, but it may generates false positives, it also does not suport removing an item, you can just add new items.

## Usage

```go
import "github.com/hugoluchessi/bloomfilter"

var (
    maxExpectedElements = 1000000
    falsePositiveProbability = 0.001
)

bf := bloomfilter.CreateBloomFilter(maxExpectedElements, falsePositiveProbability)

value := []byte{1, 2, 3, 4, 5}

err := bf.Add(value)

isContained, err := bf.Contains(value)

if !isContained {
    // Do work
}

```

## Benchmark

### Add method

|N|P|Bytes length|Operations|Time per operation|
|---|---|---|---|---|
|1000|0.1|1|1000000|1105 ns/op|
|1000|0.1|1000|500000|3479 ns/op|
|1000|0.1|1000000|2000|796466 ns/op|
|1000|0.0001|1|500000|3633 ns/op|
|1000|0.0001|1000|200000|12265 ns/op|
|1000|0.0001|1000000|500|2724019 ns/op|
|1000000|0.1|1|1000000|1079 ns/op|
|1000000|0.1|1000|500000|3346 ns/op|
|1000000|0.1|1000000|2000|697185 ns/op|
|1000000|0.0001|1|500000|2610 ns/op|
|1000000|0.0001|1000|200000|10007 ns/op|
|1000000|0.0001|1000000|1000|2277241 ns/op|

### Contains Method

|N|P|Bytes length|Operations|Time per operation|
|---|---|---|---|---|
|1000|0.1|1|1000000|1105 ns/op|
|1000|0.1|1000|500000|3583 ns/op|
|1000|0.1|1000000|2000|807841 ns/op|
|1000|0.0001|1|500000|3827 ns/op|
|1000|0.0001|1000|200000|12274 ns/op|
|1000|0.0001|1000000|500|2786809 ns/op|
|1000000|0.1|1|1000000|1086 ns/op|
|1000000|0.1|1000|500000|3431 ns/op|
|1000000|0.1|1000000|3000|704018 ns/op|
|1000000|0.0001|1|500000|2651 ns/op|
|1000000|0.0001|1000|200000|9911 ns/op|
|1000000|0.0001|1000000|1000|2257919 ns/op|