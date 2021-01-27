package main

import "fmt"

func binSearch(sl []int, num int) int {
	for {
		if len(sl) == 1 {
			if num == sl[0] {
				return num
			} else {
				return 0
			}
		}

		if len(sl) == 2 {
			switch num {
			case sl[0]:
				return num
			case sl[1]:
				return num
			default:
				return 0
			}
		}

		mid := (len(sl) - 1) / 2
		midEl := sl[mid]

		switch {
		case num == midEl:
			return num
		case num < midEl:
			sl = sl[:mid]
		case num > midEl:
			sl = sl[(mid + 1):]
		default:
			return 0
		}
	}
}

func main() {
	var num, res int
	slice := []int{5, 8, 13, 22, 34, 48, 66, 81, 89, 90, 99, 105, 110, 121, 134, 145}

	fmt.Println("Введите целое положительное число")
	for fmt.Scan(&num); num <= 0; fmt.Scan(&num) {
		fmt.Println("нужно ввести целое положительное число")
	}

	if res = binSearch(slice, num); res == 0 {
		fmt.Println("указанное число не найдено в массиве")
	} else {
		fmt.Println("число", res, "найдено в массиве")
	}
}
