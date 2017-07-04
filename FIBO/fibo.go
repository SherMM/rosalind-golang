package main 


import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


// Fibonacci calculates the nth fibonacci number
func Fibonacci(n int) int {
	x, y := 0, 1
	for n > 0 {
		x, y = y, x+y
		n--
	}
	return x
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")
	n, _ := strconv.Atoi(s)
	fmt.Println(Fibonacci(n))
}