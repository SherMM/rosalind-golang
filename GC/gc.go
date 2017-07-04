package main 

import (
	"fmt"
	"os"
	"bufio"
	"bytes"
	"strings"
)


// CalculateGCContent returns the percentage that
// G and C nucleotides appear in a dna strand
func CalculateGCContent(dna string) float64 {
	total := 0.0
	g, c := 0.0, 0.0
	for _, letter := range dna {
		nuc := string(letter)
		if nuc == "G" {
			g++
		} else if nuc == "C" {
			c++
		} 
		total++
	}
	return (g + c) / total
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fastas := make(map[string]float64)
	var dna bytes.Buffer
	var fasta string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// check if fasta has been intialized
			if fasta != "" {
				gc := CalculateGCContent(dna.String())
				fastas[fasta] = gc
				dna.Reset()
			} 
			fasta = strings.Trim(line, ">\n")
		} else {
			dna.WriteString(strings.Trim(scanner.Text(), "\n"))
		}
	}

	// find fasta with max gc
	var best_gc float64
	var best_fasta string
	for k, v := range fastas {
		if v > best_gc {
			best_gc = v
			best_fasta = k
		}
	}
	fmt.Println(best_fasta)
	fmt.Println(best_gc*100)
}