// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package buchheim

type treeData struct {
	id      string
	nodeIDX int
}

type Tree struct {
	id              string
	mod             float64
	shift           float64
	change          float64
	thread          *Tree
	x               float64
	y               float64
	children        []*Tree
	parent          *Tree
	leftMostSibling *Tree
	number          int
	data            treeData
	prelimX         float64
	ancestor        *Tree
}

func NewTree(node Node, depth int, parent *Tree, familyNumber int) *Tree {
	return &Tree{
		id:              node.ID(),
		y:               float64(depth),
		children:        make([]*Tree, 0),
		parent:          parent,
		leftMostSibling: nil,
		number:          familyNumber,
		data: treeData{
			id:      node.ID(),
			nodeIDX: node.IDX(),
		},
	}
}

func (t *Tree) AddChild(child *Tree) {
	t.children = append(t.children, child)
}

func (t Tree) Children() []*Tree {
	return t.children
}

func (t Tree) GetLeftMostSibling() *Tree {
	if t.leftMostSibling == nil && t.parent != nil && len(t.parent.Children()) > 0 && !t.IsEqual(*t.parent.Children()[0]) {
		t.leftMostSibling = t.parent.Children()[0]
	}

	return t.leftMostSibling
}

func (t Tree) GetLeftSibling() *Tree {
	var w *Tree
	if t.parent != nil {
		for _, c := range t.parent.Children() {
			if c.IsEqual(t) {
				return w
			} else {
				w = c
			}
		}
	}

	return w
}

func (t Tree) IsEqual(p Tree) bool {
	return t.data.id == p.data.id
}

func (t *Tree) SetPrelimX(x float64) {
	t.prelimX = x
}

func (t Tree) PrelimX() float64 {
	return t.prelimX
}

func (t *Tree) SetMod(mod float64) {
	t.mod = mod
}

func (t Tree) Mod() float64 {
	return t.mod
}

func (t Tree) Thread() *Tree {
	return t.thread
}

func (t *Tree) SetThread(th *Tree) {
	t.thread = th
}

func (t Tree) Change() float64 {
	return t.change
}

func (t Tree) Shift() float64 {
	return t.shift
}

func (t *Tree) SetAncestor(a *Tree) {
	t.ancestor = a
}

func (t Tree) Ancestor() *Tree {
	return t.ancestor
}

func (t Tree) Parent() *Tree {
	return t.parent
}

func (t Tree) Number() int {
	return t.number
}

func (t *Tree) SetChange(c float64) {
	t.change = c
}

func (t *Tree) SetShift(s float64) {
	t.shift = s
}

func (t *Tree) SetX(x float64) {
	t.x = x
}

func (t Tree) X() float64 {
	return t.x
}

func (t Tree) MidPoint() float64 {
	firstX := t.Children()[0].PrelimX()
	lastX := t.Children()[len(t.Children())-1].PrelimX()
	var half float64 = 2

	return (firstX + lastX) / half
}

func (t Tree) nextLeft() *Tree {
	if len(t.Children()) > 0 {
		return t.Children()[0]
	} else {
		return t.thread
	}
}

func (t Tree) nextRight() *Tree {
	if len(t.Children()) > 0 {
		return t.Children()[len(t.Children())-1]
	} else {
		return t.thread
	}
}
