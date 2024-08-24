package main

import (
	"fmt"
)

// partition выполняет разделение массива и возвращает индекс опорного элемента.
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Выбираем последний элемент как опорный
	i := low           // Индекс для меньшего элемента

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// quickSelect возвращает k-й наименьший элемент в массиве.
func quickSelect(arr []int, low, high, k int) int {
	if low == high {
		return arr[low]
	}

	// Получаем индекс опорного элемента после разделения
	pivotIndex := partition(arr, low, high)

	// Сравниваем индекс опорного элемента с k
	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelect(arr, low, pivotIndex-1, k)
	} else {
		return quickSelect(arr, pivotIndex+1, high, k)
	}
}

func main() {
	// Пример массива
	arr := []int{12, 3, 5, 7, 4, 19, 26}
	k := 2 // Ищем 3-й наименьший элемент (k=3, значит он будет на позиции k-1=2)

	result := quickSelect(arr, 0, len(arr)-1, k-1)
	fmt.Printf("%d-й наименьший элемент: %d\n", k, result)
	fmt.Println(arr)
}
