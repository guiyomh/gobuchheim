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
		buchheim.NewDependencyNode("a", "a"),
		buchheim.NewDependencyNode("b", "b"),
		buchheim.NewDependencyNode("c", "c"),
		buchheim.NewDependencyNode("d", "d"),
		buchheim.NewDependencyNode("e", "e"),
		buchheim.NewDependencyNode("f", "f"),
		buchheim.NewDependencyNode("g", "g"),
	}

	links := buchheim.LinkList{
		buchheim.NewLink("r", "a"),
		buchheim.NewLink("r", "b"),
		buchheim.NewLink("r", "c"),
		buchheim.NewLink("b", "e"),
		buchheim.NewLink("b", "d"),
		buchheim.NewLink("c", "d"),
		buchheim.NewLink("c", "f"),
		buchheim.NewLink("d", "g"),
	}

	g := buchheim.NewGraphService(root, nodes, links)
	g.Layout()

	for _, node := range nodes {
		fmt.Println(node)
	}

	file, err := os.Create("set3.svg")
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
