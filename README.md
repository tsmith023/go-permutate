# Go Permutate

This repository hosts the beginnings of a Go package for computing the permutations of generic slices as generator functions leveraging goroutines under-the-hood.

<!---
The package exposes a simple API containing a generator function and types that define the different available algorithms.

To use it, simply choose your applicable algorithm and pass it to `GeneratePermutations`. Then loop over the generator with `for perm := range GeneratePermutations` to retrieve all the possible permutations.
--->

So far implemented are [Heap's Algorithm](https://en.wikipedia.org/wiki/Heap's_algorithm), in both flat and recursive forms, for finding all possible permutations and [Fisher-Yates-Shuffling](https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle) & [Reservoir Sampling](https://en.wikipedia.org/wiki/Reservoir_sampling) for finding permutations randomly until a predefined number have been obtained.
