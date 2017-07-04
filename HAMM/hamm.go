package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


// HammingDistance returns the number of times 
// characters at the same index position in two
// equal-length strings are different
func HammingDistance(s1 string, s2 string) int {
	dist := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			dist++
		}
	}
	return dist
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	s2, _ := reader.ReadString('\n')
	s1 = strings.Trim(s1, "\n")
	s2 = strings.Trim(s2, "\n")
	fmt.Println(HammingDistance(s1, s2))
}