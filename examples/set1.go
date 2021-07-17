package main

import (
	"fmt"
	"os"

	"github.com/guiyomh/gobuchheim/pkg/buchheim"
	"github.com/guiyomh/gobuchheim/pkg/treechart"
)

func main() {

	root := buchheim.NewDependencyNode("r", "R")
	nodes := buchheim.NodeList{
		root,
		buchheim.NewDependencyNode("a", "A"),
		buchheim.NewDependencyNode("b", "B"),
		buchheim.NewDependencyNode("c", "C"),
		buchheim.NewDependencyNode("d", "D"),
	}

	links := buchheim.LinkList{
		buchheim.NewLink("r", "a"),
		buchheim.NewLink("r", "b"),
		buchheim.NewLink("a", "c"),
		buchheim.NewLink("b", "d"),
	}

	g := buchheim.NewGraphService(root, nodes, links)
	g.Layout()

	for _, node := range nodes {
		fmt.Println(node)
	}

	file, err := os.Create("set1.svg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// var treeNodes []treechart.Node = nodes

	graph := treechart.NewTreeChart(
		treechart.ConvertBuchheimNodeList(nodes),
		treechart.ConvertBuchheimLinkList(links),
		file,
	)
	graph.Draw()

}
