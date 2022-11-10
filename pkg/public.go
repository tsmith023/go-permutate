package permutations

func NewHeapsAlgorithmFlat[T any]() *HeapsAlgorithmFlat[T] {
	return &HeapsAlgorithmFlat[T]{}
}

func NewHeapsAlgorithmRecursive[T any]() *HeapsAlgorithmRecursive[T] {
	return &HeapsAlgorithmRecursive[T]{}
}

func NewReservoirSampling[T any](howManyPerms int, reservoirSize int) *ReservoirSampling[T] {
	return &ReservoirSampling[T]{
		howManyPerms:  howManyPerms,
		reservoirSize: reservoirSize,
	}
}

func NewFisherYatesShuffle[T any](howManyPerms int) *FisherYatesShuffle[T] {
	return &FisherYatesShuffle[T]{
		howManyPerms: howManyPerms,
	}
}

func (haf *HeapsAlgorithmFlat[T]) Permutate(ch chan *[]T, data []T) {
	haf.permutateFlatly(ch, data)
}

func (har *HeapsAlgorithmRecursive[T]) Permutate(ch chan *[]T, data []T) {
	har.permutateRecursively(ch, data, len(data))
}

func (rs *ReservoirSampling[T]) Permutate(ch chan *[]T, data []T) {
	rs.permutate(ch, data)
}

func (fys *FisherYatesShuffle[T]) Permutate(ch chan *[]T, data []T) {
	fys.permutate(ch, data)
}

func GeneratePermutations[T any, P Permutator[T]](data []T, p P) chan *[]T {
	ch := make(chan *[]T)
	if len(data) < 2 {
		// No possible permutations for [] nor [T]
		ch <- &data
		return ch
	}
	go func(ch chan *[]T) {
		defer close(ch)
		p.Permutate(ch, data)
	}(ch)
	return ch
}
