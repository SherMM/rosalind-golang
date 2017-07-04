package main 

import (
	"fmt"
	"os"
	"bytes"
	"bufio"
	"strings"
)

// ReverseComplement returns the reverse complement
// of a strand of DNA which is swapping 'A' with 'T'
// and 'C' with 'G', and vice-versa for both, then 
// reversing the entire strand
func ReverseComplement(dna string) string {
	comps := map[string]string {
		"A": "T",
		"C": "G",
		"G": "C",
		"T": "A",
	}
	var revc bytes.Buffer
	for i := len(dna)-1; i >= 0; i-- {
		revc.WriteString(comps[string(dna[i])])
	} 
	return revc.String()
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	dna := strings.Trim(s, "\n")
	revc := ReverseComplement(dna)
	fmt.Println(revc)
}