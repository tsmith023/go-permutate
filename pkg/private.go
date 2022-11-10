package permutations

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

func swap[T any](i int, j int, data []T) {
	tmp := data[j]
	data[j] = data[i]
	data[i] = tmp
}

func cloneSlice[T any](data []T) []T {
	new := make([]T, len(data))
	copy(new, data)
	return new
}

func (haf *HeapsAlgorithmFlat[T]) permutateFlatly(ch chan *[]T, data []T) {
	// https://en.wikipedia.org/wiki/Heap%27s_algorithm
	// https://sedgewick.io/

	n := len(data)
	// c is an encoding of the stack state. c[k] encodes the for-loop counter for when generate(k - 1, A) is called
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = 0
	}

	i := 0
	for i < n {
		if i == 0 {
			out := cloneSlice(data)
			ch <- &out
		}
		if c[i] < i {
			if i%2 == 0 {
				swap(0, i, data)
			} else {
				swap(c[i], i, data)
			}
			// Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
			c[i] += 1
			// Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
			i = 0
		} else {
			// Calling generate(i+1, A) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
			c[i] = 0
			i += 1
		}
	}
}

func (har *HeapsAlgorithmRecursive[T]) permutateRecursively(ch chan *[]T, data []T, whereInRecursion int) {
	if whereInRecursion == 1 {
		out := cloneSlice(data)
		ch <- &out
	} else {
		nextStep := whereInRecursion - 1
		har.permutateRecursively(ch, data, nextStep)
		for i := 0; i < nextStep; i++ {
			if whereInRecursion%2 == 0 {
				swap(i, nextStep, data)
			} else {
				swap(0, nextStep, data)
			}
			har.permutateRecursively(ch, data, nextStep)
		}
	}
}

func (rs *ReservoirSampling[T]) permutate(ch chan *[]T, data []T) {
	if rs.reservoirSize > len(data) {
		panic("The size of the reservoir cannot be greater than the size of the sample space")
	}
	var wg sync.WaitGroup
	for perm := 0; perm < rs.howManyPerms; perm++ {
		wg.Add(1)
		go func(ch chan *[]T, data []T) {
			defer wg.Done()
			ch <- rs.sample(data)
		}(ch, cloneSlice(data))
	}
	wg.Wait()
}

func (rs *ReservoirSampling[T]) sample(sample []T) *[]T {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	out := make([]T, rs.reservoirSize)
	for i := 0; i < rs.reservoirSize; i++ {
		out[i] = sample[i]
	}

	weight := math.Exp(math.Log(r.Float64()) / float64(rs.reservoirSize))

	i := 0
	for i < len(sample) {
		i = i + int(math.Floor(math.Log(r.Float64())/math.Log(1-weight))+1)
		if i < len(sample) {
			out[r.Intn(rs.reservoirSize)] = sample[i]
			weight = weight * math.Exp(math.Log(r.Float64())/float64(rs.reservoirSize))
		}
	}

	return &out
}

func (fys *FisherYatesShuffle[T]) sample(sample []T) *[]T {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	n := len(sample)
	for i := 0; i < n-1; i++ {
		j := r.Intn(n)
		swap(i, j, sample)
	}
	return &sample
}

func (fys *FisherYatesShuffle[T]) permutate(ch chan *[]T, data []T) {
	var wg sync.WaitGroup
	for i := 0; i < fys.howManyPerms; i++ {
		wg.Add(1)
		go func(ch chan *[]T, sample []T) {
			defer wg.Done()
			ch <- fys.sample(sample)
		}(ch, cloneSlice(data))
	}
	wg.Wait()
}
