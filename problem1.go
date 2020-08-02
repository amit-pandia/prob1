package main

import (
	"fmt"
)

type testWrapper struct {
	TestCount int
	testData  []test
}

type test struct {
	shopCount    int
	rosesPerShop []int
	result       int
}

func main() {

	tw := testWrapper{}
	// read no. of tests
	// fmt.Println("Enter Test cases:")
	_, err := fmt.Scanf("%d", &tw.TestCount)
	if err != nil {
		return
	}

	for i := 1; i <= tw.TestCount; i++ {
		t := test{}

		// read no. of shops
		// fmt.Println("Enter no of shops:")
		_, err := fmt.Scanf("%d", &t.shopCount)
		if err != nil {
			return
		}

		length := t.shopCount
		//fmt.Println("Enter the inputs: ")
		t.rosesPerShop = make([]int, length)
		for i := 0; i < length; i++ {
			fmt.Scan(&t.rosesPerShop[i])
		}

		// calculate result
		oddSum := 0
		evenSum := 0

		for i := 0; i < t.shopCount; i++ {
			if (i % 2) == 0 {
				evenSum += t.rosesPerShop[i] // sum of roses provided by even numbered shops
			} else {
				oddSum += t.rosesPerShop[i] // sum of roses provided by odd numbered shops
			}
		}

		if oddSum > evenSum {
			t.result = oddSum
		} else if evenSum > oddSum {
			t.result = evenSum
		} else {
			t.result = evenSum
		}

		tw.testData = append(tw.testData, t)
	}

	fmt.Println("Output:")
	for _, t := range tw.testData {
		fmt.Println(t.result)
	}

	return
}
