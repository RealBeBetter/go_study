package sort

type MergeSort struct {
}

func (sort *MergeSort) Sort(array []int) {
	mergeSortRange(array, 0, len(array)-1)
}

func mergeSortRange(array []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSortRange(array, start, mid)
	mergeSortRange(array, mid+1, end)
	merge(array, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	index := 0
	// 左半部分[0:mid]，右半部分[mid+1:end]
	for ; i <= mid && j <= end; index++ {
		if arr[i] <= arr[j] {
			tmpArr[index] = arr[i]
			i++
		} else {
			tmpArr[index] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[index] = arr[i]
		index++
	}
	for ; j <= end; j++ {
		tmpArr[index] = arr[j]
		index++
	}
	copy(arr[start:end+1], tmpArr)
}
