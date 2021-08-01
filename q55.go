package leetcode

func canJump(nums []int) bool {
    result := false
    tmpMax := nums[0]
    lastMax := 0
    newMax := tmpMax
    n := len(nums)

    for tmpMax < n - 1 {
        // get the range
        // in the range
        flag := false
        for i := lastMax; i <= tmpMax && i < n; i++ {
            if newMax < nums[i] + i {
                flag = true
                //lastMax = newMax // NOT here!
                newMax = nums[i] + i
            }
        }
        lastMax = tmpMax // !!! BUT here!
        tmpMax = newMax
        if flag == false {break}
    }

    if newMax >= n - 1 {result = true}

    return result
}

/*
func myJudge(nums []int, maxIndex int) bool {
    if maxIndex <= 0 {return true}
    flag := false
    for i := 0; i < maxIndex; i++ {
        if nums[i] >= maxIndex {
            flag = flag || myJudge(nums, i)
        }
        if flag == true {break}
    }
    return flag
}
*/


