package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


// CountNucleotides returns the nucleotide counts of 
// 'A' adenine, 'C' cytosine, 'G' guanine, and 'T' thymine in
// a strand of DNA
func CountNucleotides(dna string) (map[string]int) {
	counts := make(map[string]int)
	for _, letter := range strand {
		counts[string(letter)]++
	}
	return counts
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	dna := strings.Trim(s, "\n")
	counts := CountNucleotides(dna)
	nucs := [4]string{"A", "C", "G", "T"}
	for i := 0; i < len(nucs); i++ {
		fmt.Print(counts[nucs[i]])
		fmt.Print(" ")
	}
	fmt.Println("")
}