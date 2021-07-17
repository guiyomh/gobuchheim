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
		buchheim.NewDependencyNode("h", "h"),
		buchheim.NewDependencyNode("i", "i"),
		buchheim.NewDependencyNode("j", "j"),
		buchheim.NewDependencyNode("k", "k"),
		buchheim.NewDependencyNode("l", "l"),
		buchheim.NewDependencyNode("m", "m"),
		buchheim.NewDependencyNode("n", "n"),
		buchheim.NewDependencyNode("o", "o"),
		buchheim.NewDependencyNode("p", "p"),
		buchheim.NewDependencyNode("q", "q"),
		buchheim.NewDependencyNode("s", "s"),
		buchheim.NewDependencyNode("t", "t"),
	}

	links := buchheim.LinkList{
		buchheim.NewLink("r", "a"),
		buchheim.NewLink("r", "d"),
		buchheim.NewLink("r", "n"),
		buchheim.NewLink("r", "p"),
		buchheim.NewLink("a", "b"),
		buchheim.NewLink("a", "c"),
		buchheim.NewLink("d", "e"),
		buchheim.NewLink("d", "k"),
		buchheim.NewLink("e", "f"),
		buchheim.NewLink("e", "i"),
		buchheim.NewLink("f", "g"),
		buchheim.NewLink("f", "h"),
		buchheim.NewLink("i", "j"),
		buchheim.NewLink("k", "l"),
		buchheim.NewLink("l", "m"),
		buchheim.NewLink("n", "o"),
		buchheim.NewLink("p", "q"),
		buchheim.NewLink("q", "s"),
		buchheim.NewLink("q", "t"),
	}

	g := buchheim.NewGraphService(root, nodes, links)
	g.Layout()

	for _, node := range nodes {
		fmt.Println(node)
	}

	file, err := os.Create("set2.svg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// var treeNodes []treechart.Node = nodes

	graph := treechart.NewTreeChart(
		treechart.ConvertBuchheimNodeList(nodes),
		treechart.ConvertBuchheimLinkList(links),
		file,
		treechart.WithHeight(500),
		treechart.WithWidth(500),
		treechart.WithXScale(50),
		treechart.WithYScale(20),
	)
	graph.Draw()

}
