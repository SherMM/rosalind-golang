package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


type Vertex struct {
	val int
	degree int
	adj []int 
}

type Graph struct {
	nodes []Vertex
	n int
	m int 
}


func (g Graph) Init(n, m int) Graph {
	graph := Graph{n: n, m: m}
	for i := 0; i < n; i++ {
		vert := Vertex{}
		graph.nodes = append(graph.nodes, vert)
	}
	return graph
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := 0
	var n, m int
	var g Graph
	for scanner.Scan() {
		if c < 1 {
			t := strings.Split(scanner.Text(), " ")
			n, _ = strconv.Atoi(t[0])
			m, _ = strconv.Atoi(t[1])
			g = g.Init(n, m)
			fmt.Println(g)
			c++
		} else {
			t := strings.Split(scanner.Text(), " ")
			u, _ := strconv.Atoi(t[0])
			v, _ := strconv.Atoi(t[1])
			fmt.Println(u, v)
		}
	}
}