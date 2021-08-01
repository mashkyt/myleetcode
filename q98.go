package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
    return myIsValid(root, math.MinInt64, math.MaxInt64)
}

func myIsValid(root *TreeNode, bot int, top int) bool {
    if root == nil {return true}

    if !(root.Val > bot && root.Val < top) {
        return false
    }
    /*
    flag := true
    if root.Left != nil {
        //if !(root.Left.Val < root.Val) {return false}
        flag = flag && myIsValid(root.Left, bot, mymin(top, root.Val))
    }
    if flag == false {return false}
    if root.Right != nil {
        //if !(root.Right.Val > root.Val) {return false}
        flag = flag && myIsValid(root.Right, mymax(bot, root.Val), top)
    }
    */
    return myIsValid(root.Left, bot, mymin(top, root.Val)) && myIsValid(root.Right, mymax(bot, root.Val), top)
}

func mymin(a int, b int) int {
    if a < b {return a}
    return b
}

func mymax(a int, b int) int {
    if a > b {return a}
    return b
}