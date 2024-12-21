package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nevernorbo/kruskal-mst/kruskal"
)

// Here so that the terminal window doesn't close automatically when the program is finished
func showCloseDialogue(reader *bufio.Reader) {
	fmt.Println("\nType 'exit' to close the program")
	for {
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(input)) == "exit" {
			break
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of vertices: ")

	verticesStr, _ := reader.ReadString('\n')
	vertices, err := strconv.Atoi(strings.TrimSpace(verticesStr))
	if err != nil || vertices <= 0 {
		fmt.Println("Invalid number of vertices")
		showCloseDialogue(reader)
		return
	}

	g := kruskal.NewGraph(vertices)

	fmt.Println("\nEnter edges in format 'source destination weight'")
	fmt.Println("Example: 0 1 5")
	fmt.Println("Type 'done' to finish entering edges")

	// Take user input until "done" is input
	for {
		fmt.Print("Edge: ")
		edgeStr, _ := reader.ReadString('\n')
		input := strings.TrimSpace(edgeStr)

		if strings.ToLower(input) == "done" {
			break
		}

		edgeData := strings.Fields(input)
		if len(edgeData) != 3 {
			fmt.Println("Invalid input format. Please try again.")
			continue
		}

		src, err1 := strconv.Atoi(edgeData[0])
		dest, err2 := strconv.Atoi(edgeData[1])
		weight, err3 := strconv.Atoi(edgeData[2])

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Invalid input numbers. Please try again.")
			continue
		}

		if src >= vertices || dest >= vertices || src < 0 || dest < 0 {
			fmt.Printf("Vertex numbers must be between 0 and %d\n", vertices-1)
			continue
		}

		if weight < 0 {
			fmt.Println("Weight must be non-negative")
			continue
		}

		g.AddEdge(src, dest, weight)
		fmt.Printf("Added edge: %d -- %d with weight %d\n", src, dest, weight)
	}

	if g.IsEmpty() {
		fmt.Println("No edges were added to the graph")
		showCloseDialogue(reader)
		return
	}

	// Find and display MST
	mst := g.KruskalMST()
	kruskal.DisplayKruskal(mst)

	showCloseDialogue(reader)
}
