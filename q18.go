package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
    var result [][]int
    //var tmpAns []int
    n := len(nums)
    if n < 4 {return result}

    p0 := 0
    p1 := 1
    p2 := 2
    p3 := 3

    sort.Ints(nums)

    for p0 = 0; p0 < n - 3; p0 = moveRight(nums, p0, n) { // NOT p0++
        for p1 = p0 + 1; p1 < n - 2; p1 = moveRight(nums, p1, n) { // NOT p1++
            tmpSum1 := nums[p0] + nums[p1]
            p2 = p1 + 1
            p3 = n - 1
            for p2 < p3 && p2 < n && p3 > p1 {
                tmpSum2 := nums[p2] + nums[p3]
                if tmpSum1 + tmpSum2 == target {
                    tmpAns := []int{nums[p0], nums[p1], nums[p2], nums[p3]}
                    result = append(result, tmpAns)
                    p2 = moveRight(nums, p2, n)
                    p3 = moveLeft(nums, p3, p1)
                } else if tmpSum1 + tmpSum2 < target {
                    p2 = moveRight(nums, p2, n)
                } else {
                    p3 = moveLeft(nums, p3, p1)
                }
            }
        }
    }
    return result
}

func moveRight(nums []int, index int, n int) int {
    if index >= n {return n}
    for index < n - 1 {
        if nums[index] != nums[index + 1] {return index + 1}
        index++
    }
    return n
}

func moveLeft(nums []int, index int, l int) int {
    if index < l {return l - 1}
    for index > l {
        if nums[index] != nums[index - 1] {return index - 1}
        index--
    }
    return l - 1
}