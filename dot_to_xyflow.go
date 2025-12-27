package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/goccy/go-graphviz"
)

type NodeData struct {
	Label string `json:"label"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Node struct {
	ID       string   `json:"id"`
	Data     NodeData `json:"data"`
	Position Position `json:"position"`
	Type     string   `json:"type"`
}

type MarkerEnd struct {
	Type string `json:"type"`
}

type Edge struct {
	ID        string    `json:"id"`
	Source    string    `json:"source"`
	Target    string    `json:"target"`
	MarkerEnd MarkerEnd `json:"markerEnd"`
	Label     string    `json:"label,omitempty"`
}

type XYFlowData struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run dot_to_xyflow.go <filename.dot>")
		return
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	ctx := context.Background()
	gv, err := graphviz.New(ctx)
	if err != nil {
		log.Fatalf("Error creating graphviz context: %v", err)
	}
	defer gv.Close()

	graph, err := graphviz.ParseBytes(content)
	if err != nil {
		log.Fatalf("Error parsing DOT: %v", err)
	}
	defer graph.Close()

	// Set layout engine (default to dot)
	gv.SetLayout(graphviz.DOT)

	// Render to XDOT to trigger layout and populate attributes
	if err := gv.Render(ctx, graph, graphviz.XDOT, io.Discard); err != nil {
		log.Fatalf("Error performing layout: %v", err)
	}

	// Get graph bounding box to flip Y axis
	bb := graph.GetStr("bb")
	var llx, lly, urx, ury float64
	fmt.Sscanf(bb, "%f,%f,%f,%f", &llx, &lly, &urx, &ury)
	graphHeight := ury

	nodes := []Node{}
	edges := []Edge{}

	// Iterate nodes
	n, err := graph.FirstNode()
	for n != nil && err == nil {
		id, _ := n.Name()
		label := n.GetStr("label")
		if label == "" {
			label = id
		}
		label = strings.ReplaceAll(label, "\n", "\n")

		posStr := n.GetStr("pos")
		var x, y float64
		if posStr != "" {
			fmt.Sscanf(posStr, "%f,%f", &x, &y)
		}

		// Graphviz pos is center. XYFlow is top-left.
		// width/height are in inches, convert to points (72 dpi)
		width, _ := strconv.ParseFloat(n.GetStr("width"), 64)
		height, _ := strconv.ParseFloat(n.GetStr("height"), 64)
		wPts := width * 72
		hPts := height * 72
		
		// Adjust x from center to left
		x = x - (wPts / 2)
		// Flip Y: Graphviz is bottom-up, XYFlow is top-down
		// y is center, so top-left Y in XYFlow is graphHeight - (y + hPts/2)
		y = graphHeight - (y + hPts/2)

		nodeType := "default"
		shape := n.GetStr("shape")
		if shape == "component" || shape == "doublecircle" {
			nodeType = "output"
		} else if shape == "note" || shape == "invhouse" || strings.Contains(strings.ToLower(id), "input") {
			nodeType = "input"
		}

		nodes = append(nodes, Node{
			ID: id,
			Data: NodeData{Label: label},
			Position: Position{X: x, Y: y},
			Type: nodeType,
		})
		n, err = graph.NextNode(n)
	}

	// Iterate edges
	edgeCount := 0
	n, err = graph.FirstNode()
	for n != nil && err == nil {
		e, err := graph.FirstOut(n)
		for e != nil && err == nil {
			tail, _ := e.Tail()
			head, _ := e.Head()
			src, _ := tail.Name()
			tgt, _ := head.Name()
			
			label := e.GetStr("label")

			edges = append(edges, Edge{
				ID: fmt.Sprintf("e%d", edgeCount),
				Source: src,
				Target: tgt,
				MarkerEnd: MarkerEnd{Type: "arrowclosed"},
				Label: label,
			})
			edgeCount++
			e, err = graph.NextOut(e)
		}
		n, err = graph.NextNode(n)
	}

	out, _ := json.MarshalIndent(XYFlowData{Nodes: nodes, Edges: edges}, "", "  ")
	fmt.Println(string(out))
}
