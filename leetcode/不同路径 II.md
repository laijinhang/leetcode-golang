[不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)
```golang
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    dp := make([][]int, len(obstacleGrid))
    for i := 0;i < len(obstacleGrid);i++ {
        dp[i] = make([]int, len(obstacleGrid[0]))
    }
    for i := 0;i < len(obstacleGrid);i++ {
        if obstacleGrid[i][0] == 1 {
            dp[i][0] = 0
            break
        }
        dp[i][0] = 1
    }
    for i := 0;i < len(obstacleGrid[0]);i++ {
        if obstacleGrid[0][i] == 1 {
            dp[0][i] = 0
            break
        }
        dp[0][i] = 1
    }

    for i := 1;i < len(obstacleGrid);i++ {
        for j := 1;j < len(obstacleGrid[0]);j++ {
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
                continue
            }
            dp[i][j] = dp[i][j-1] + dp[i-1][j]
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
}
```
