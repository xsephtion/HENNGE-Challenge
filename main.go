package main

import (
	"fmt"
	"sort"
)

var numbers []int
var testnum []int
var testsol int
var sol []int

// The challenge here is not to use for loop

func main() {
	var num int
	fmt.Scanf("%d", &num)
	arrsol := testcount(num)
	sort.SliceStable(arrsol, func(i, j int) bool {
		return true
	})
	printsol(arrsol, num)
}

// Printing reversed output
func printsol(arr []int, l int) {
	if l == 0 {
		return
	}
	fmt.Println(arr[l-1])
	printsol(arr, l-1)
}

//getting the testcount i.e.: 2
func testcount(num int) []int {
	if num <= 0 || num >= 100 {
		return sol
	} else {
		var testc int
		fmt.Scanf("%d", &testc)
		testsol = tests(testc)
		testnum = []int{}
		sol = append(sol, testsol)
		testcount(num - 1)
	}
	return sol
}

//getting the how many numbers and the value of numbers squared
func tests(num int) int {
	if num <= 0 || num >= 100 {
		return addall(testnum, len(testnum))
	} else {
		var testval int
		fmt.Scanf("%d", &testval)
		if testval > 0 {
			testnum = append(testnum, testval*testval)
			tests(num - 1)
		} else {
			tests(num - 1)
		}
	}
	return addall(testnum, len(testnum))
}

//adding all the squared numbers inside array
func addall(arr []int, num int) int {
	if num == 0 {
		return 0
	}
	return (addall(arr, num-1) + arr[num-1])
}
