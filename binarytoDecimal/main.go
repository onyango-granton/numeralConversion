package main

import "fmt"

func main() {
	num := 11000000

	res := ""
	revRes := ""

	for num != 0 {
		startRune := '0'
		modulus := num % 10
		for i := 0; i < modulus; i++ {
			startRune++
		}
		res += string(startRune)

		num = num / 10
	}

	for i:=len(res) - 1 ; i >=0; i--{
		revRes += string(res[i])
	}

	fmt.Println(revRes)

	indexpower := -1
	indexArray := []int{}

	for i:=len(revRes) - 1; i >= 0; i--{
		indexpower++
		indexArray = append(indexArray, indexpower)
	}

	for i:=0 ; i < len(indexArray); i++{
		for j:=0; j<len(indexArray);j++{
			if j+1 < len(indexArray) && indexArray[j] < indexArray[j+1]{
				indexArray[j], indexArray[j+1] = indexArray[j+1], indexArray[j]
			}
		}
	}

	fmt.Println(indexArray)

	sumTotal := 0

	for i:=0; i < len(revRes); i++{
		if revRes[i] == '1'{
			fmt.Println(2,indexArray[i])
			sumTotal += factorial(2,indexArray[i]) * 1
		} else {
			sumTotal += factorial(2,indexArray[i]) * 0
		}
	}

	fmt.Println(factorial(2,0))
	fmt.Println(sumTotal)
}


func factorial(num, pow int) int {
	if pow == 0{
		return 1
	}
	startNum := num
	for i:=1; i < pow; i++{
		startNum = startNum * num	
	}
	return startNum
}
