package sort

type QuickSort struct {
}

func (s *QuickSort) Sort(array []int) {
	if len(array) < 2 {
		return
	}

	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	middle := partition(array, left, right)
	quickSort(array, left, middle-1)
	quickSort(array, middle+1, right)
}

func partition(array []int, left, right int) int {
	// pivot 为基数，即待排序数组的第一个数
	pivot := array[left]
	i, j := left, right

	for i < j {
		// 从右往左找到第一个小于基数的数
		for array[j] >= pivot && i < j {
			j--
		}

		// 从左往右找到第一个大于基数的数
		for array[i] <= pivot && i < j {
			i++
		}

		array[i], array[j] = array[j], array[i]
	}

	// 使得划分好的数分布在 pivot 两侧，左侧小一点的数放在首位，基数放在中间位置
	array[left], array[i] = array[i], array[left]
	return i
}
