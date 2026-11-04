# Task
You are given an array of integers `arr` and a positive integer `k` such that `1 <= k <= len(arr)`.
Write a function `solve` that returns an array with the maximum element of each `k` long sub-array (contiguous sub sequence) of `arr`.

# Example
For the array `[2, -1, 3, 2, 5, 7, 6, 8, 11]` and `k=4` the answer is `[3, 5, 7, 7, 8, 11]`.

# Explanation
The sub-arrays are:
1. `[2, -1, 3, 2]`
2. `[-1, 3, 2, 5]`
3. `[3, 2, 5, 7]`
4. `[2, 5, 7, 6]`
5. `[5, 7, 6, 8]`
6. `[7, 6, 8, 11]`

The maximum element in each is:
1. `3`
2. `5`
3. `7`
4. `7`
5. `8`
6. `11`