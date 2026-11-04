# Task
You are given array of distinct integers `arr`. We treat this array as circular.
In this circular array every element has a next element.
The next element of `arr[i]` is `arr[(i + 1)%len(arr)]`.
We define the following function:
```go
func Next(arr []int, i int) []int {
    n := len(arr)
    next := make([]int, n-1)
    for j := 1; j < n; j++ {
        next[j] = arr[(i + j)%n]
    }
    return next
}
```
Write a function `func solve(arr []int) (res []int)` such that:
1. `len(res) == len(arr)`
2. for each `i` in `{0, 1, ..., len(arr)}` is true that:
- `res[i]` is the first element of `Next(arr, i)` that is larger than `arr[i]`, if such exists.
- `res[i] == -1`, otherwise.

Do this in the most efficient way.

# Example
For the array `[2, 3, 1, 5, 6, 4]` the result should be `[3, 5, 5, 6, -1, 5]`.

# Explanation
1. `Next(arr, 0) == []int{3, 1, 5, 6, 4}` and the first element larger than `2` is `3`.
2. `Next(arr, 1) == []int{1, 5, 6, 4, 2}` and the first element larger than `3` is `5`.
3. `Next(arr, 2) == []int{5, 6, 4, 2, 3}` and the first element larger than `1` is `5`.
4. `Next(arr, 3) == []int{6, 4, 2, 3, 1}` and the first element larger than `5` is `6`.
5. `Next(arr, 4) == []int{4, 2, 3, 1, 5}` and there is no element larger than `6` hence the element at index 4 should be `-1`.
6. `Next(arr, 5) == []int{2, 3, 1, 5, 6}` and the first element larger than `4` is `5`.