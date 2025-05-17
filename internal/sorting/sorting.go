package sorting

func SelectionSort(data *[]int) (compareCount int, shiftCount int) {
	length := len(*data)

	for i := 0; i < length-1; i++ {
		minIndex := i

		for j := i + 1; j < length; j++ {
			if (*data)[j] < (*data)[minIndex] {
				minIndex = j
			}
			compareCount++
		}

		if minIndex != i {
			(*data)[i], (*data)[minIndex] = (*data)[minIndex], (*data)[i]
			shiftCount++
		}
	}

	return
}

func InsertionSort(data *[]int) (compareCount int, shiftCount int) {
	length := len(*data)

	for i := 1; i < length; i++ {
		current := (*data)[i]
		j := i - 1

		for j >= 0 {
			compareCount++
			if (*data)[j] > current {
				(*data)[j+1] = (*data)[j]
				shiftCount++
				j--
			} else {
				break
			}
		}

		(*data)[j+1] = current
		if j+1 != i {
			shiftCount++
		}
	}

	return
}

func BubbleSort(data *[]int) (compareCount int, shiftCount int) {
	length := len((*data))

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if (*data)[j] > (*data)[j+1] {
				(*data)[j], (*data)[j+1] = (*data)[j+1], (*data)[j]
				shiftCount++
			}
			compareCount++
		}
	}

	return
}
