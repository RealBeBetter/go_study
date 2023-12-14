/**
 * @author wei.song
 * @since 2023/10/6 14:21
 */
package main

func main() {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(array)
	println(area)
}

func maxArea(height []int) int {
	length := len(height)
	maxVolume := 0

	left := 0
	right := length - 1

	for left < right {
		currentVolume := minInt(height[left], height[right]) * (right - left)
		if currentVolume > maxVolume {
			maxVolume = currentVolume
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxVolume
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
