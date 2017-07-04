package main 

import (
	"fmt"
	"os"
	"sort"
	"bufio"
	"strings"
	"bytes"
	"index/suffixarray"
)


type ByLength []string


func (s ByLength) Len() int {
	return len(s)
}


func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}


// GetSubstrings returns a list of all unique
// substrings in a string
func GetSubstrings(strand string) []string {
	var subs []string
	found := make(map[string]int) 
	for i := 0; i < len(strand); i++ {
		for j := i; j <= len(strand); j++ {
			sub := strand[i:j]
			_, ok := found[sub]
			if !ok  && sub != "" {
				subs = append(subs, sub)
			}
			found[sub]++
		}
	}
	return subs
}


// MapSubstrings returns a map of substring, count 
// key value pairs of the number of times a substring
// appears across a group of strings
func MapSubstrings(strands []string) map[string]int {
	common := make(map[string]int)
	for _, strand  := range strands {
		subs := GetSubstrings(string(strand))
		for _, sub := range subs {
			common[string(sub)]++
		}
	}
	return common
}


// FindLCS finds the longest common substring 
// using a brute force algorithm
func FindLCS(strands []string) string {
	var best string
	var length int
	n := len(strands)
	subs := MapSubstrings(strands)
	for sub, count := range subs {
		if count == n && len(sub) > length {
			best = sub
			length = len(sub)
		}
	}
	return best
}

// FindLCSFast finds the longest common substring 
// in a group of strings using a suffix array data structure
func FindLCSFast(strands []string) string {
	data := []byte("\x00" + strings.Join(strands, "\x00") + "\x00")
	sa := suffixarray.New(data)
	// get all unique substrings for one strands,
	// since we expect there to be a common substring
	// amongst the strands
	subs := GetSubstrings(strands[0])
	// reverse sort, since we want the longest substring
	sort.Sort(ByLength(subs))
	var best string
	for _, sub := range subs {
		indices := sa.Lookup([]byte(sub), -1)
		if len(indices) >= len(strands) {
			best = string(sub)
			break
		}
	}
	return best
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var strands []string
	var strand bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// check if fasta has been intialized
			s := strand.String()
			if s != "" {
				strands = append(strands, s)
			}
			strand.Reset()
		} else {
			strand.WriteString(strings.Trim(scanner.Text(), "\n"))
		}
	}
	strands = append(strands, strand.String())

	// Slower Algorithm 
	fmt.Println(FindLCS(strands))

	// Faster Algorithm
	fmt.Println(FindLCSFast(strands))
}