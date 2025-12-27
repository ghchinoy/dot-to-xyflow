package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
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
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	dot := string(content)
	
	// Pre-cleanup: remove comments and rankdir
	dot = regexp.MustCompile(`(?m)//.*`).ReplaceAllString(dot, "")
	dot = regexp.MustCompile(`/\[\s\S]*?\*/`).ReplaceAllString(dot, "")
	dot = regexp.MustCompile(`rankdir=.*;`).ReplaceAllString(dot, "")

	nodes := make(map[string]Node)
	var nodeIds []string
	var edges []Edge

	// 1. First Pass: Get all Edges (Src -> Tgt [label="..."])
	edgePat := regexp.MustCompile(`(\w+)\s*->\s*(\w+)(?:\s*\[[^\]]*label\s*=\s*"([^"]+)"[^\]]*\])?`)
	eMatches := edgePat.FindAllStringSubmatch(dot, -1)
	for i, m := range eMatches {
		src, tgt, label := m[1], m[2], m[3]
		
		if src == "node" || src == "graph" || src == "edge" { continue }

		if _, ok := nodes[src]; !ok {
			nodes[src] = Node{ID: src, Data: NodeData{Label: src}, Type: "default"}
			nodeIds = append(nodeIds, src)
		}
		if _, ok := nodes[tgt]; !ok {
			nodes[tgt] = Node{ID: tgt, Data: NodeData{Label: tgt}, Type: "default"}
			nodeIds = append(nodeIds, tgt)
		}

		edges = append(edges, Edge{
			ID: fmt.Sprintf("e%d", i),
			Source: src, Target: tgt,
			MarkerEnd: MarkerEnd{Type: "arrowclosed"},
			Label: label,
		})
	}

	// 2. Second Pass: Explicit Node definitions
	nodePat := regexp.MustCompile(`(?m)^\s*(\w+)\s*\[([^\]]*label\s*=\s*"([^"]+)"[^\]]*)\]`)
	nMatches := nodePat.FindAllStringSubmatch(dot, -1)
	for _, m := range nMatches {
		id, attr, label := m[1], m[2], strings.ReplaceAll(m[3], "\\n", "\n")
		if id == "node" || id == "graph" || id == "edge" { continue }

		nodeType := "default"
		if strings.Contains(attr, "shape=component") { nodeType = "output" }
		if strings.Contains(strings.ToLower(id), "input") { nodeType = "input" }

		n, exists := nodes[id]
		if !exists {
			nodeIds = append(nodeIds, id)
			n = Node{ID: id}
		}
		n.Data.Label = label
		n.Type = nodeType
		nodes[id] = n
	}

	var finalNodes []Node
	for i, id := range nodeIds {
		node := nodes[id]
		node.Position = Position{X: 250, Y: float64(i * 120)}
		finalNodes = append(finalNodes, node)
	}

	out, _ := json.MarshalIndent(XYFlowData{Nodes: finalNodes, Edges: edges}, "", "  ")
	fmt.Println(string(out))
}