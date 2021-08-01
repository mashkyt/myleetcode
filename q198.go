package leetcode

func rob(nums []int) int {

    n := len(nums)
    // var tmpMax []int

    if n == 1 {return nums[0]}
    if n == 2 {return mymax(nums[0], nums[1])}
    if n == 3 {return mymax(nums[1], nums[2] + nums[0])}

    nums[1] = mymax(nums[0], nums[1])
    nums[2] = mymax(nums[0] + nums[2], nums[1])
    

    for i := 3; i < n; i++ {
        t1 := nums[i] + nums[i-2]
        t2 := nums[i-1]
        nums[i] = mymax(t1, t2)
    }

    tmpMax := nums[0]
    for i := 0; i < n; i++ {
        if tmpMax < nums[i] {
            tmpMax = nums[i]
        }
    }

    return tmpMax
}

func mymin(a int, b int) int {
    if a < b {return a}
    return b
}

func mymax(a int, b int) int {
    if a > b {return a}
    return b
}