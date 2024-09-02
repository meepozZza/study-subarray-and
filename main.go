package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter array via comma: ")
	arrayString := readStringInput()
	array := strings.Split(arrayString, ",")
	intArray := make([]int, len(array))

	for i, v := range array {
		v, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		intArray[i] = v
	}

	fmt.Print("Enter k: ")
	k := readIntInput()

	finalState := countSubarrays(intArray, k)

	fmt.Print("Result: ", finalState, "\n")
}

func readStringInput() (val string) {
	_, err := fmt.Scanf("%s", &val)

	if err != nil {
		panic(err)
	}

	return
}

func readIntInput() (val int) {
	_, err := fmt.Scanf("%d", &val)

	if err != nil {
		panic(err)
	}

	return
}

func countSubarrays(nums []int, k int) int64 {
	cntByVal := make(map[int]int)
	total := int64(0)
	for _, n := range nums {
		newCntByVal := make(map[int]int)
		if n&k == k {
			newCntByVal[n] = 1
			for v, cnt := range cntByVal {
				newVal := v & n
				if newVal&k == k {
					newCntByVal[newVal] += cnt
				}
			}
			total += int64(newCntByVal[k])
		}
		cntByVal = newCntByVal
	}
	return total
}
