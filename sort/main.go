package main

import "fmt"

var array = []int{21,33,456,4,1,344,534,122,456,35,67,84,22}

func main() {
	sort()
	fmt.Println(array)
}

func sort() {
	N := len(array)
	for i := 0; i < N-1; i++ {
		currMin := i
		for k := i + 1; k < N; k++ {
			if array[k] < array[currMin] {
				currMin = k
				fmt.Printf("currMin[%d] = %d \n",k,array[k])
			}
		}
		// swap
		fmt.Println(currMin," ",i)
		if currMin != i {
			temp := array[i]
			array[i] = array[currMin]
			array[currMin] = temp
			fmt.Println(array)
		}
	}
}
