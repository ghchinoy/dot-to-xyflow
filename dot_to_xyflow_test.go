package main

import (
	"context"
	"io"
	"testing"

	"github.com/goccy/go-graphviz"
)

func TestTranslate(t *testing.T) {
	dot := `
digraph G {
	A -> B [label="edge label"];
	B [shape=component, label="Node B"];
}
`
	ctx := context.Background()
	gv, err := graphviz.New(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer gv.Close()

	graph, err := graphviz.ParseBytes([]byte(dot))
	if err != nil {
		t.Fatal(err)
	}
	defer graph.Close()

	gv.SetLayout(graphviz.DOT)
	if err := gv.Render(ctx, graph, graphviz.XDOT, io.Discard); err != nil {
		t.Fatal(err)
	}

	// Verify nodes
	n, err := graph.FirstNode()
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for n != nil {
		count++
		name, _ := n.Name()
		if name != "A" && name != "B" {
			t.Errorf("unexpected node name: %s", name)
		}
		if name == "B" {
			if n.GetStr("shape") != "component" {
				t.Errorf("expected shape component for B, got %s", n.GetStr("shape"))
			}
		}
		n, _ = graph.NextNode(n)
	}
	if count != 2 {
		t.Errorf("expected 2 nodes, got %d", count)
	}

	// Verify edges
	eCount := 0
	n, _ = graph.FirstNode()
	for n != nil {
		e, _ := graph.FirstOut(n)
		for e != nil {
			eCount++
			if e.GetStr("label") != "edge label" {
				t.Errorf("expected edge label 'edge label', got '%s'", e.GetStr("label"))
			}
			e, _ = graph.NextOut(e)
		}
		n, _ = graph.NextNode(n)
	}
	if eCount != 1 {
		t.Errorf("expected 1 edge, got %d", eCount)
	}
}
