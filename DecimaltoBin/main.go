package main

import "fmt"

func main() {
	// num := 250

	// bitsNo := 8
	// startIndex := -1

	// indexArray := []int{}

	// for range bitsNo{
	// 	startIndex++
	// 	indexArray = append(indexArray, startIndex)
	// }

	// for range indexArray{
	// 	for i := range indexArray{
	// 		if i+1 < len(indexArray) && indexArray[i] < indexArray[i+1]{
	// 			indexArray[i], indexArray[i+1] = indexArray[i+1], indexArray[i]
	// 		}
	// 	}
	// }

	// fmt.Println(indexArray)

	// indexArr := []int{7, 6, 5, 4, 3, 2, 1, 0}

	// res := ""

	// for num > 0 {
	// 	for i := 0; i < len(indexArr); i++ {
	// 		if factorial(2, indexArr[i]) > num {
	// 			res += "0"
	// 		} else {
	// 			fmt.Println(num,"-",factorial(2,indexArr[i]))
	// 			num = num - factorial(2, indexArr[i])
	// 			res += "1"
	// 		}
	// 	}
	// 	fmt.Println(num)
	// }

	// fmt.Println(res)
	fmt.Println(DecimalToBin(255,8))
	fmt.Println(stringToInt("456"))

}


func DecimalToBin(num, bitsNumb int) string {
	res := ""
	startIndex := -1

	indexArray := []int{}

	for range bitsNumb{
		startIndex++
		indexArray = append(indexArray, startIndex)
	}

	for range indexArray{
		for i := range indexArray{
			if i+1 < len(indexArray) && indexArray[i] < indexArray[i+1]{
				indexArray[i], indexArray[i+1] = indexArray[i+1], indexArray[i]
			}
		}
	}

	

	for num > 0 {
		for i := 0; i < len(indexArray); i++ {
			if factorial(2, indexArray[i]) > num {
				res += "0"
			} else {
				// fmt.Println(num,"-",factorial(2,indexArray[i]))
				num = num - factorial(2, indexArray[i])
				res += "1"
			}
		}
		// fmt.Println(num)
	}
	return res
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


func stringToInt(s string) int {
	res := 0
	for _,ch := range s{
		startRune := '0'
		iteration := int(ch - startRune)
		count := 0
		for range iteration{
			count++
		}
		res = res * 10 + count
	}

	return res
}