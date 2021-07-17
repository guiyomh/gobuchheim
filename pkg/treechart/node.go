// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package treechart

import "github.com/guiyomh/gobuchheim/pkg/buchheim"

type Node interface {
	X() float64
	Y() float64
	Label() string
}

type node struct {
	x     float64
	y     float64
	label string
}

func (n node) X() float64 {
	return n.x
}
func (n node) Y() float64 {
	return n.y
}
func (n node) Label() string {
	return n.label
}

func ConvertBuchheimNodeList(nodes buchheim.NodeList) []Node {
	treeNodes := make([]Node, len(nodes))
	for i := range nodes {
		treeNodes[i] = node{
			x:     nodes[i].X(),
			y:     nodes[i].Y(),
			label: nodes[i].Label(),
		}
	}

	return treeNodes
}
