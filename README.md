# Go range behavior with arrays/slices

This repository demonstrates a common pitfall in Go when using the `range` keyword with arrays or slices.  The `range` function iterates over indices rather than values.  This can be unexpected and lead to errors if you're not careful.

## Bug Description

The `bug.go` file shows a simple example of how the `range` function iterates over indices, resulting in unexpected output when attempting to print array elements.

## Solution

The `bugSolution.go` demonstrates a correct way to iterate through slices or arrays using range, utilizing both the index and the value returned by range for accessing elements correctly.