package main

import "fmt"

func main() {
	num := 250

	indexArr := []int{7, 6, 5, 4, 3, 2, 1, 0}

	res := ""

	for num > 0 {
		for i := 0; i < len(indexArr); i++ {
			if factorial(2, indexArr[i]) > num {
				res += "0"
			} else {
				fmt.Println(num,"-",factorial(2,indexArr[i]))
				num = num - factorial(2, indexArr[i])
				res += "1"
			}
		}
		fmt.Println(num)
	}

	fmt.Println(res)

}

func factorial(num, pow int) int {
	if pow == 0 {
		return 1
	}
	startNum := num
	for i := 1; i < pow; i++ {
		startNum = startNum * num
	}
	return startNum
}