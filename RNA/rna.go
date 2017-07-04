package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// TranscribeToRNA converts the 'T' thymine nucleotides
// in a strand of DNA to 'U' ribose to produce an RNA strand
func TranscribeToRNA(dna string) string {
	rna := strings.Replace(dna, "T", "U", -1)
	return rna
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	dna := strings.Trim(s, "\n")
	rna := TranscribeToRNA(dna)
	fmt.Println(rna)
}