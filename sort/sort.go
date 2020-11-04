package sort

// 插入排序
func InsertionSort(data []int) {
	count := len(data)
	if count <= 1 {
		return
	}

	for i := 1; i < count; i++ {
		value := data[i]
		j := i - 1
		// 插入插入的位置
		for ; j >= 0; j-- {
			if data[j] > value {
				data[j+1] = data[j] // 数据移动
			} else {
				break
			}
		}

		data[j+1] = value
	}

	return
}

// 选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
func SelectionSort(data []int) {
	count := len(data)
	if count <= 1 {
		return
	}
	for i := 0; i < count; i++ {
		min := i
		for j := i + 1; j < count; j++ {
			if data[min] > data[j] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}

	return
}

// 快速排序
func QuickSort(data []int) {

}

// 归并排序
func MergeSort(data []int) {
	if len(data) <= 1 {
		return
	}

}

func merge(left []int, right []int) []int {
	if len(left) <= 0 {
		return right
	}
	if len(right) <= 0 {
		return left
	}

	leftIndex := 0
	rightIndex := 0
	result := make([]int, 0, len(left)+len(right))
	for i := 0; i < len(result); i++ {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else if left[leftIndex] > right[rightIndex] {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	if leftIndex < len(left) {
		result = append(result, left[leftIndex:]...)
	}

	if rightIndex < len(right) {
		result = append(result, right[rightIndex:]...)
	}
	return result
}
