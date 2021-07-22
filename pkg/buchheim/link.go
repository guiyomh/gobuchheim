// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package buchheim

import "github.com/pkg/errors"

// Link represents a link between 2 Nodes
//go:generate mockery --name=Link  --output=mocks --filename=link.go --outpkg=mocks
type Link interface {
	SourceID() string
	TargetID() string
	SetSource(Node)
	SetTarget(Node)
	Source() Node
	Target() Node
}

type link struct {
	source Node
	target Node
	sID    string
	tID    string
}

// NewLink creates a Link instance
func NewLink(sID, tID string) Link {
	return &link{
		sID: sID,
		tID: tID,
	}
}

func (bl link) SourceID() string {
	return bl.sID
}

func (bl link) TargetID() string {
	return bl.tID
}

func (bl *link) SetSource(n Node) {
	bl.source = n
}

func (bl *link) SetTarget(n Node) {
	bl.target = n
}

func (bl link) Source() Node {
	return bl.source
}

func (bl link) Target() Node {
	return bl.target
}

// LinkList represents a collection of Link
type LinkList []Link

// Remove removes the link from the list at the specified index
func (ll LinkList) Remove(i int) LinkList {
	ll[i] = ll[len(ll)-1]

	return ll[:len(ll)-1]
}

// FindByIDs search a link in the list
func (ll LinkList) FindByIDs(sID, tID string) (Link, error) {
	for _, l := range ll {
		if l.SourceID() == sID && l.TargetID() == tID {
			return l, nil
		}
	}

	return nil, errors.New(`Link("` + sID + `","` + tID + `") not found`)
}

// Add appends a link if isn't already append
func (ll *LinkList) Add(l Link) {
	_, err := ll.FindByIDs(l.SourceID(), l.TargetID())
	if err == nil {
		return
	}
	*ll = append(*ll, l)
}
