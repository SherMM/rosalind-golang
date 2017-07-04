package main 


import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"bytes"
)


// MapCodons reads in a file of codon->acid mappings
// and returns a map of these mappings
func MapCodons(filename string) map[string]string {
	codons := make(map[string]string)
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		codon, acid := line[0], line[1]
		codons[codon] = acid
	}
	return codons
}


// TranslateRNAtoProtein returns a string that translates 
// the codons (3-char slices) to their corresponding acids
// and stops when reaching a codon indicating a stopping position
func TranslateRNAtoProtein(rna string, codons map[string]string) string {
	var protein bytes.Buffer
	i := 0
	for {
		codon := rna[i: i+3]
		acid := codons[codon]
		if acid == "Stop" {
			break
		}
		protein.WriteString(acid)
		i += 3
	}
	return protein.String()
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	rna := strings.Trim(s, "\n")

	filename := "codons.txt"
	codons := MapCodons(filename)
	protein := TranslateRNAtoProtein(rna, codons)
	fmt.Println(protein)
}