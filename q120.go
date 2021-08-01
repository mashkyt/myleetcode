package leetcode

func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    if n <= 0{return 0}

    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            if i > 0 {
                if j == 0 {
                    triangle[i][j] += triangle[i-1][j]
                }
                if j > 0 && j < i {
                    triangle[i][j] += mymin(triangle[i-1][j-1], triangle[i-1][j])
                }
                if j == i {
                    triangle[i][j] += triangle[i-1][j-1]
                }
            }
        }
    }

    tmpMin := triangle[n-1][0]
    for j := 0; j < n; j++ {
        if tmpMin > triangle[n-1][j] {tmpMin = triangle[n-1][j]}
    }

    return tmpMin
}

func mymin(a int, b int) int {
    if a < b {return a}
    return b
}