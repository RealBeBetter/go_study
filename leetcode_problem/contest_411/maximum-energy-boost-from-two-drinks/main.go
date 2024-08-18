package main

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	n := len(energyDrinkA)
	dp := make([][2]int64, n+1)

	// 初始化状态
	dp[0][0] = 0
	dp[0][1] = 0

	// 动态规划填充状态
	for i := 1; i <= n; i++ {
		// 继续饮用 A 类型饮料
		dp[i][0] = dp[i-1][0] + int64(energyDrinkA[i-1])
		// 切换到 A 类型饮料
		dp[i][0] = max(dp[i][0], dp[i-1][1]-int64(energyDrinkB[i-1]))

		// 继续饮用 B 类型饮料
		dp[i][1] = dp[i-1][1] + int64(energyDrinkB[i-1])
		// 切换到 B 类型饮料
		dp[i][1] = max(dp[i][1], dp[i-1][0]-int64(energyDrinkA[i-1]))
	}

	// 返回最终状态的最大值
	return max(dp[n][0], dp[n][1])
}
