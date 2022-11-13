# Go Permutate

This repository hosts the beginnings of a Go package for computing the [permutations](https://en.wikipedia.org/wiki/Permutation) of generic slices as generator functions leveraging goroutines under-the-hood.

<!---
The package exposes a simple API containing a generator function and types that define the different available algorithms.

To use it, simply choose your applicable algorithm and pass it to `GeneratePermutations`. Then loop over the generator with `for perm := range GeneratePermutations` to retrieve all the possible permutations.
--->

So far implemented are [Heap's Algorithm](https://en.wikipedia.org/wiki/Heap's_algorithm), in both flat and recursive forms, for finding all possible permutations and [Fisher-Yates-Shuffling](https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle) & [Reservoir Sampling](https://en.wikipedia.org/wiki/Reservoir_sampling) for finding permutations randomly until a predefined number have been obtained.

The benefit of the latter type of generation over the former is if the permutated slice has a length greater than eleven such that finding all its possible permutations is unfeasible, since Heap's Algorithm has order ${\cal O}(n!)$ so that finding (>12)! permutations takes many minutes, if not hours. The latter technique can find a representative sample of the underlying permutations of a slice when the length of the slice is much greater than 12.
