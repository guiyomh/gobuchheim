// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package buchheim

import "math"

// Grapher represents a service that evaluate X and Y for the list of nodes
// see https://github.com/BlueInt32/buchheim-js/blob/852eaf3de20b3852450c4237aa29ab18cd212ee2/service.js#L245
type Grapher interface {
	Layout() error
}

// graphService implement the buchheim algorithm
// This algorithm is O(n) with 2 passes (bottom up recursive then modifiers applications)
// <!> Don't forget your Doliprane <!>
type graphService struct {
	nodes    NodeList
	links    LinkList
	root     Node
	layout   Layout
	maxY     float64
	minY     float64
	maxX     float64
	minX     float64
	distance float64
}

// NewGraphService creates a instance of Grapher
func NewGraphService(root Node, nodes []Node, links []Link) Grapher {
	return &graphService{
		nodes:    nodes,
		links:    links,
		root:     root,
		layout:   Horizontal,
		distance: 1,
	}
}

func (g *graphService) Layout() error {
	g.prepareNodes()
	g.prepareLinks()
	tree, err := g.convertToTree(g.root, 0, nil, 1)
	if err != nil {
		return err
	}
	g.firstWalk(tree)
	g.secondWalk(tree, tree.PrelimX())

	return nil
}

func (g *graphService) prepareNodes() {
	for nodeIdx, node := range g.nodes {
		node.SetIDX(nodeIdx)
	}
}

func (g *graphService) prepareLinks() {
	for i := len(g.links) - 1; i >= 0; i-- {
		link := g.links[i]
		source, err := g.nodes.FindByID(link.SourceID())
		if err != nil {
			g.links = g.links.Remove(i)

			continue
		}
		link.SetSource(source)
		target, err := g.nodes.FindByID(link.TargetID())
		if err != nil {
			g.links = g.links.Remove(i)

			continue
		}
		link.SetTarget(target)
		link.Source().AddOutgoingLink(link)
	}
}

func (g *graphService) convertToTree(node Node, depth int, parent *Tree, familyNumber int) (*Tree, error) {
	tree := NewTree(node, depth, parent, familyNumber)
	err := g.setNodeY(tree, float64(depth))
	if err != nil {
		return nil, err
	}
	for i, link := range node.OutgoingLink() {
		subtree, err := g.convertToTree(link.Target(), depth+1, tree, i+1)
		if err != nil {
			continue
		}
		tree.AddChild(subtree)
	}

	return tree, nil
}

func (g *graphService) firstWalk(v *Tree) {

	// if tree is leaf
	if len(v.Children()) == 0 {
		if left := v.GetLeftMostSibling(); left != nil {
			v.SetPrelimX(v.GetLeftSibling().PrelimX() + g.distance)
		}
	} else {
		// the left most child of v
		defaultAncestor := v.Children()[0]
		for _, w := range v.Children() {
			g.firstWalk(w)
			g.apportion(w, defaultAncestor)
		}
		g.executeShifts(v)
		// midpoint is at the middle of leftmost and rightmost children of v
		midpoint := v.MidPoint()

		if w := v.GetLeftSibling(); w != nil {
			v.SetPrelimX(w.PrelimX() + g.distance)
			v.SetMod(v.PrelimX() - midpoint)
		} else {
			v.SetPrelimX(midpoint)
		}
	}
}

// apportion compares two subtrees v and its left sibling w, using their contour and
// sets some offsets to be sumed up in the secondWalk
// TODO reduce the complexity
func (g *graphService) apportion(tree *Tree, defaultAncestor *Tree) *Tree {
	if w := tree.GetLeftSibling(); w != nil {
		vri := tree                      // right inner : represents the inner contour of the right subtree
		vro := tree                      // right outer : represents the outer contour of the right subtree
		vli := w                         // left inner : represents the inner contour of the left subtree
		vlo := tree.GetLeftMostSibling() // left outer : represents the outer contour of the left subtree
		sri := vri.Mod()
		sro := vro.Mod()
		sli := vli.Mod()
		slo := vlo.Mod()
		for vli.nextRight() != nil && vri.nextLeft() != nil {
			var vriPrelimX, vriMod, vroMod float64 = 0, 0, 0
			var vliPrelimX, vliMod, vloMod float64 = 0, 0, 0

			if vli = vli.nextRight(); vli != nil {
				vliPrelimX = vli.PrelimX()
				vliMod = vli.Mod()
			}
			if vri = vri.nextLeft(); vri != nil {
				vriPrelimX = vri.PrelimX()
				vriMod = vri.Mod()
			}
			if vlo = vli.nextLeft(); vlo != nil {
				vloMod = vlo.Mod()
			}
			if vro = vro.nextRight(); vro != nil {
				vroMod = vro.Mod()
			}

			vro.ancestor = tree
			shift := vliPrelimX + sli - (vriPrelimX + sri) + g.distance
			if shift > 0 {
				g.moveSubTree(g.ancestor(vli, tree, defaultAncestor), tree, shift)
				sri += shift
				sro += shift
			}
			sli += vliMod
			sri += vriMod
			slo += vloMod
			sro += vroMod

		}
		if vli != nil && vro != nil && vli.nextRight() != nil && vro.nextRight() == nil {
			vro.SetThread(vli.nextRight())
			vro.SetMod(vro.Mod() + sli - sro)
		}
		if vri != nil && vlo != nil && vri.nextLeft() != nil && vlo.nextLeft() == nil {
			vlo.SetThread(vri.nextLeft())
			vlo.SetMod(vlo.Mod() + sri - slo)
			defaultAncestor = tree
		}
	}

	return defaultAncestor
}

func (g *graphService) moveSubTree(wl, wr *Tree, shift float64) {
	if shift == 0 {
		return
	}
	subtrees := float64(wr.Number() - wl.Number())
	wr.SetChange(wr.Change() - shift/subtrees)
	wr.SetShift(wr.Shift() + shift)
	wl.SetChange(wl.Change() + shift/subtrees)
	wr.SetPrelimX(wr.PrelimX() + shift)
	wr.SetMod(wr.Mod() + shift)
}

func (g *graphService) ancestor(vli, tree, defaultAncestor *Tree) *Tree {
	if g.areSiblings(vli.Ancestor(), tree) {
		return vli.Ancestor()
	}

	return defaultAncestor
}

func (g *graphService) secondWalk(v *Tree, m float64) {
	v.SetX(v.PrelimX() + m)
	g.setNodeX(v, v.X())
	for _, child := range v.Children() {
		g.secondWalk(child, m+v.Mod())
	}

}

func (g *graphService) areSiblings(v, w *Tree) bool {
	if v == nil || v.Parent() == nil {
		return false
	}
	for _, child := range v.Parent().Children() {
		if w.IsEqual(*child) {
			return true
		}
	}

	return false
}

func (g *graphService) setNodeX(v *Tree, x float64) {
	n, err := g.nodes.FindByID(v.data.id)
	if err != nil {
		return
	}
	if g.layout == Vertical {
		n.SetX(x)
	} else {
		n.SetY(x)
	}

	g.maxX = math.Max(g.maxX, x)
	g.minX = math.Min(g.minX, x)

}

func (g *graphService) setNodeY(tree *Tree, y float64) error {
	n, err := g.nodes.FindByID(tree.data.id)
	if err != nil {
		return err
	}
	if g.layout == Vertical {
		n.SetY(y)
	} else {
		n.SetX(y)
	}
	g.maxY = math.Max(g.maxY, y)
	g.minY = math.Min(g.minY, y)

	return nil
}

func (g *graphService) executeShifts(tree *Tree) {
	var shift float64 = 0
	var change float64 = 0
	for _, w := range tree.Children() {
		w.SetPrelimX(w.PrelimX() + shift)
		w.SetMod(w.Mod() + shift)
		change += w.Change()
		shift += w.Shift() + change
	}
}
