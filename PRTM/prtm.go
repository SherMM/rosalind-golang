package main 


import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


// MapWeights returns a map of amino acid strings to 
// their corresponding mass
func MapWeights(filename string) map[string]float64 {
	weights := make(map[string]float64)
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		p, w := line[0], line[1]
		wght, _ := strconv.ParseFloat(w, 64)
		weights[p] = wght
	}
	return weights
}


// CalculateProteinMass returns the total mass of all amino acid
// chars in a protein strand
func CalculateProteinMass(strand string, w map[string]float64) float64 {
	total := 0.0
	for _, acid := range strand {
		total += w[string(acid)]
	}
	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	strand := strings.Trim(s, "\n")

	filename := "masses.txt"
	weights := MapWeights(filename)
	fmt.Println(CalculateProteinMass(strand, weights))
}