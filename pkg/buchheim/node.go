// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package buchheim

import (
	"fmt"

	"github.com/pkg/errors"
)

// Node represents a graph node
//go:generate mockery --name=Node  --output=mocks --filename=node.go --outpkg=mocks
type Node interface {
	IDX() int
	SetIDX(int)
	SetX(float64)
	SetY(float64)
	X() float64
	Y() float64
	ID() string
	Label() string
	AddOutgoingLink(Link)
	OutgoingLink() LinkList
	AddIncomingLink(Link)
	IncomingLink() LinkList
	String() string
}

// DependencyNode implements the Node interface
type DependencyNode struct {
	id            string
	label         string
	idx           int
	x             float64
	y             float64
	outgoingLinks LinkList
	incomingLinks LinkList
}

// NewDependencyNode creates an new instance of a DependencyNode
func NewDependencyNode(id, label string) *DependencyNode {
	return &DependencyNode{
		id:            id,
		label:         label,
		outgoingLinks: make(LinkList, 0),
		incomingLinks: make(LinkList, 0),
	}
}

func (bn DependencyNode) ID() string {
	return bn.id
}

func (bn *DependencyNode) SetIDX(id int) {
	bn.idx = id
}

func (bn DependencyNode) IDX() int {
	return bn.idx
}

func (bn *DependencyNode) AddOutgoingLink(l Link) {
	bn.outgoingLinks = append(bn.outgoingLinks, l)
}

func (bn DependencyNode) OutgoingLink() LinkList {
	return bn.outgoingLinks
}

func (bn *DependencyNode) AddIncomingLink(l Link) {
	bn.incomingLinks = append(bn.incomingLinks, l)
}

func (bn DependencyNode) IncomingLink() LinkList {
	return bn.incomingLinks
}

func (bn DependencyNode) Label() string {
	return bn.label
}

func (bn *DependencyNode) SetX(x float64) {
	bn.x = x
}

func (bn DependencyNode) X() float64 {
	return bn.x
}

func (bn *DependencyNode) SetY(y float64) {
	bn.y = y
}

func (bn DependencyNode) Y() float64 {
	return bn.y
}

func (bn DependencyNode) String() string {
	out := fmt.Sprintf("# %s (x:%v, y:%v)\n", bn.id, bn.x, bn.y)
	for _, outlink := range bn.outgoingLinks {
		out += fmt.Sprintf("%s -> %s\n", outlink.SourceID(), outlink.TargetID())
	}

	return out + "\n"
}

// NodeList represents a collection of unordered Node
type NodeList []Node

// FindByID search a Node by id in the NodeList
func (nl NodeList) FindByID(id string) (Node, error) {
	for _, n := range nl {
		if n.ID() == id {
			return n, nil
		}
	}

	return nil, errors.New(`Node "` + id + `" not found`)
}
