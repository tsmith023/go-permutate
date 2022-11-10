package permutations

import (
	"testing"
	"time"
)

func factorial(N int) int {
	if N != 0 {
		return N * factorial(N-1)
	} else {
		return 1
	}
}

func TestHeapsAlgorithmFlat(t *testing.T) {
	alg := NewHeapsAlgorithmFlat[int]()
	data := []int{2, 1, 3, 4}
	t.Log("Beginning heaps algorithm flat testing")
	start := time.Now()
	counts := 0
	for perm := range GeneratePermutations(data, alg) {
		if perm == nil {
			t.Fatalf(`A permutation was returned as nil %v`, data)
		} else {
			// fmt.Println(perm)
			counts += 1
		}
	}
	if counts != 24 {
		t.Fatalf(`Number of permutations returned does not match %v! = %v: %v`, len(data), factorial(len(data)), counts)
	}
	end := time.Now()
	t.Logf(`Heaps Algorithm (Flat) took %v seconds to compute %v permutations of %v (length %v)`, end.Sub(start), factorial(len(data)), data, len(data))
}

func TestReservoirSampling(t *testing.T) {
	howManyPerms := 6
	reservoirSize := 3
	alg := NewReservoirSampling[int](howManyPerms, reservoirSize)
	data := []int{1, 2, 3, 4}
	t.Log("Beginning reservoir sampling testing")
	start := time.Now()
	counts := 0
	for perm := range GeneratePermutations(data, alg) {
		if perm == nil {
			t.Fatalf(`A permutation was returned as nil %v`, data)
		} else {
			if len(*perm) != reservoirSize {
				t.Fatalf(`A permutation was returned with the wrong length: %v`, perm)
			}
			counts += 1
		}
	}
	if counts != howManyPerms {
		t.Fatalf(`Number of permutations returned does not match %v`, howManyPerms)
	}
	end := time.Now()
	t.Logf(`Reservoir Sampling took %v seconds to compute %v permutations of %v (length %v)`, end.Sub(start), howManyPerms, data, len(data))
}

func TestFisherYatesShuffle(t *testing.T) {
	howManyPerms := 6
	alg := NewFisherYatesShuffle[int](howManyPerms)
	data := []int{1, 2, 3, 4}
	t.Log("Beginning Fisher-Yates-Shuffle testing")
	start := time.Now()
	counts := 0
	for perm := range GeneratePermutations(data, alg) {
		if perm == nil {
			t.Fatalf(`A permutation was returned as nil %v`, data)
		} else {
			counts += 1
		}
	}
	if counts != howManyPerms {
		t.Fatalf(`Number of permutations returned does not match %v`, howManyPerms)
	}
	end := time.Now()
	t.Logf(`Fisher Yates Shuffle took %v seconds to compute %v permutations of %v (length %v)`, end.Sub(start), howManyPerms, data, len(data))
}

func BenchmarkHeapsAlgorithmFlat(b *testing.B) {
	alg := NewHeapsAlgorithmFlat[int]()
	for i := 2; i < b.N; i++ {
		start := time.Now()
		minIter := i
		data := make([]int, minIter)
		for j := 0; j < i; j++ {
			data[j] = j + 1
		}
		counts := 0
		for perm := range GeneratePermutations(data, alg) {
			if perm == nil {
				b.Fatalf(`A permutation was returned as nil %v`, data)
			}
			counts += 1
		}
		if counts != factorial(len(data)) {
			b.Fatalf(`Number of permutations returned does not match %v! = %v: %v`, len(data), factorial(len(data)), counts)
		}
		end := time.Now()
		b.Logf(`Heaps Algorithm (Flat) took %v seconds to compute %v permutations of %v (length %v)`, end.Sub(start), factorial(len(data)), data, len(data))
	}
}

func BenchmarkFisherYatesShuffle(b *testing.B) {
	for i := 2; i < b.N; i++ {
		alg := NewFisherYatesShuffle[int](factorial(i))
		start := time.Now()
		minIter := i
		data := make([]int, minIter)
		for j := 0; j < i; j++ {
			data[j] = j + 1
		}
		counts := 0
		for perm := range GeneratePermutations(data, alg) {
			if perm == nil {
				b.Fatalf(`A permutation was returned as nil %v`, data)
			}
			counts += 1
		}
		if counts != factorial(i) {
			b.Fatalf(`Number of permutations returned does not match %v! = %v: %v`, len(data), factorial(i), counts)
		}
		end := time.Now()
		b.Logf(`Fisher Yates Shuffle took %v seconds to compute %v permutations of %v (length %v)`, end.Sub(start), factorial(i), data, len(data))
	}
}

func BenchmarkReservoirSampling(b *testing.B) {
	for i := 2; i < b.N; i++ {
		alg := NewReservoirSampling[int](factorial(i), i)
		start := time.Now()
		minIter := i
		data := make([]int, minIter)
		for j := 0; j < i; j++ {
			data[j] = j + 1
		}
		counts := 0
		for perm := range GeneratePermutations(data, alg) {
			if perm == nil {
				b.Fatalf(`A permutation was returned as nil %v`, data)
			}
			counts += 1
		}
		if counts != factorial(i) {
			b.Fatalf(`Number of permutations returned does not match %v! = %v: %v`, len(data), factorial(i), counts)
		}
		end := time.Now()
		b.Logf(`Reservoir Sampling took %v seconds to compute %v permutations of %v (length %v)`, end.Sub(start), factorial(i), data, len(data))
	}
}
