package main 


import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// FindMotifLocations returns a list of all index (1-based) positions
// in a dna strand where a slice starts that matches a substring
func FindMotifLocations(dna string, sub string) []int {
	var locs []int
	s := len(sub)
	for i:= 0; i < len(dna) - s; i++ {
		curr := dna[i: i+s]
		if strings.EqualFold(curr, sub) {
			locs = append(locs, i+1)
		}
	}
	return locs
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	dna, _ := reader.ReadString('\n')
	sub, _ := reader.ReadString('\n')
	dna = strings.Trim(dna, "\n")
	sub = strings.Trim(sub, "\n")
	locs := FindMotifLocations(dna, sub)
	for _, loc := range locs {
		fmt.Print(loc)
		fmt.Print(" ")
	}
	fmt.Println()
}