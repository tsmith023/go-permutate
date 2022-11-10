# Go Permutate

This repository hosts the beginnings of a Go package for computing the permutations of generic slices as generator functions leveraging goroutines under-the-hood.

<!---
The package exposes a simple API containing a generator function and types that define the different available algorithms.

To use it, simply choose your applicable algorithm and pass it to `GeneratePermutations`. Then loop over the generator with `for perm := range GeneratePermutations` to retrieve all the possible permutations.
--->