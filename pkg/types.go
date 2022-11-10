package permutations

type Permutator[T any] interface {
	Permutate(ch chan *[]T, data []T)
}

type HeapsAlgorithmFlat[T any] struct{}

type HeapsAlgorithmRecursive[T any] struct{}

type ReservoirSampling[T any] struct {
	howManyPerms  int
	reservoirSize int
}

type FisherYatesShuffle[T any] struct {
	howManyPerms int
}
