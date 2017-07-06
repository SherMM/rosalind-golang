package main 

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)


func BinarySearch(a []int, val int) int {
	l, h := 0, len(a)-1
	i := -1
	for l < h {
		m := (h + l) / 2
		if a[m] == val {
			i = m + 1
			break
		} else if val < a[m] {
			h = m
		} else {
			l = m + 1
		}
	}
	return i
}


func StrArrayToIntArray(s []string) []int {
	var nums []int 
	for i := 0; i < len(s); i++ {
		val, _ := strconv.Atoi(s[i])
		nums = append(nums, val)
	}
	return nums
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	c := 0
	var nums [][]int
	for scanner.Scan() {
		if c >= 2 {
			t := strings.Split(scanner.Text(), " ")
			nums = append(nums, StrArrayToIntArray(t))
		} else {
			c++
		}
	}

	arr1, arr2 := nums[0], nums[1]

	for _, val := range arr2 {
		fmt.Print(BinarySearch(arr1, val))
		fmt.Print(" ")
	}
	fmt.Println()
}