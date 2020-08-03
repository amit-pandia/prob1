package main

import (
	"fmt"
)

type testWrapper struct {
	TestCount int
	testData  []test
}

type roseDetails struct {
	roseCount int
	idx       int
}

type test struct {
	shopCount    int
	rosesPerShop []roseDetails
	result       int
	ignoredShops map[int]int
}

func main() {

	// get inputs
	tw := testWrapper{}
	_, err := fmt.Scanf("%d", &tw.TestCount)
	if err != nil {
		return
	}

	for i := 1; i <= tw.TestCount; i++ {
		t := test{}

		// get no of test cases
		_, err := fmt.Scanf("%d", &t.shopCount)
		if err != nil {
			return
		}

		// read test inputs
		t.rosesPerShop = make([]roseDetails, t.shopCount)
		for i := 0; i < t.shopCount; i++ {
			t.rosesPerShop[i].idx = i
			fmt.Scan(&t.rosesPerShop[i].roseCount)
			t.ignoredShops = map[int]int{}
		}

		tw.testData = append(tw.testData, t)
	}

	// Bubble sort: sort in decending order of rose count
	for _, t := range tw.testData {
		for i := 0; i < len(t.rosesPerShop); i++ {
			for j := i + 1; j < len(t.rosesPerShop); j++ {
				if t.rosesPerShop[i].roseCount < t.rosesPerShop[j].roseCount {
					tmp := t.rosesPerShop[j]
					t.rosesPerShop[j] = t.rosesPerShop[i]
					t.rosesPerShop[i] = tmp
				}
			}
		}
	}

	// calculate result
	for tIdx, t := range tw.testData {
		for _, r := range t.rosesPerShop {
			_, ok := t.ignoredShops[r.idx]
			if !ok {
				t.result += r.roseCount
				if r.idx < len(t.rosesPerShop)-1 {
					t.ignoredShops[r.idx+1] = r.idx + 1
				}
				if r.idx > 0 {
					t.ignoredShops[r.idx-1] = r.idx - 1
				}
			}
		}
		tw.testData[tIdx] = t
	}

	fmt.Println("Output:")
	for _, t := range tw.testData {
		fmt.Println(t.result)
	}
	return
}
