// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package buchheim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToTree(t *testing.T) {
	t.Parallel()
	t.Run("Should convert a node in a tree", func(t *testing.T) {
		t.Parallel()
		root := NewDependencyNode("a", "A")

		g := &graphService{
			nodes:    NodeList{root},
			links:    LinkList{},
			root:     root,
			layout:   Horizontal,
			distance: 1,
		}
		root.SetIDX(1)
		tree, err := g.convertToTree(root, 0, nil, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, tree.data.nodeIDX)
		assert.Equal(t, "a", tree.id)
	})
	t.Run("Should convert a node with dependancy in a tree with child", func(t *testing.T) {
		t.Parallel()
		root := NewDependencyNode("a", "A")
		dep := NewDependencyNode("b", "B")

		ab := NewLink("a", "b")
		ab.SetSource(root)
		ab.SetTarget(dep)
		root.AddOutgoingLink(ab)

		g := &graphService{
			nodes:    NodeList{root, dep},
			links:    LinkList{ab},
			root:     root,
			layout:   Horizontal,
			distance: 1,
		}
		root.SetIDX(1)
		tree, err := g.convertToTree(root, 0, nil, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, tree.data.nodeIDX)
		assert.Equal(t, "a", tree.id)
		assert.Len(t, tree.Children(), 1)
	})
	t.Run("Should convert a node with dependancy in a tree without child", func(t *testing.T) {
		t.Parallel()
		root := NewDependencyNode("a", "A")
		dep := NewDependencyNode("b", "B")

		ab := NewLink("a", "b")
		ab.SetSource(root)
		ab.SetTarget(dep)
		root.AddOutgoingLink(ab)

		g := &graphService{
			nodes:    NodeList{root},
			links:    LinkList{ab},
			root:     root,
			layout:   Horizontal,
			distance: 1,
		}
		root.SetIDX(1)
		tree, err := g.convertToTree(root, 0, nil, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, tree.data.nodeIDX)
		assert.Equal(t, "a", tree.id)
		assert.Len(t, tree.Children(), 0)
	})
}

func TestPrepareNodes(t *testing.T) {
	t.Run("Should set up node index", func(t *testing.T) {
		a := NewDependencyNode("a", "A")
		b := NewDependencyNode("b", "B")
		c := NewDependencyNode("c", "C")

		g := &graphService{
			nodes:    NodeList{a, b, c},
			links:    LinkList{},
			root:     a,
			layout:   Horizontal,
			distance: 1,
		}
		g.prepareNodes()
		assert.Equal(t, 0, a.IDX())
		assert.Equal(t, 1, b.IDX())
		assert.Equal(t, 2, c.IDX())
	})
}

func TestPrepareLinks(t *testing.T) {
	t.Run("Should set up links", func(t *testing.T) {
		a := NewDependencyNode("a", "A")
		b := NewDependencyNode("b", "B")
		c := NewDependencyNode("c", "C")

		ab := NewLink("a", "b")
		bc := NewLink("b", "c")
		cd := NewLink("c", "d")
		da := NewLink("d", "a")

		g := &graphService{
			nodes:    NodeList{a, b, c},
			links:    LinkList{ab, bc, cd, da},
			root:     a,
			layout:   Horizontal,
			distance: 1,
		}

		g.prepareLinks()
		assert.Equal(t, a, ab.Source())
		assert.Equal(t, b, ab.Target())
		assert.Equal(t, b, bc.Source())
		assert.Equal(t, c, bc.Target())

		assert.Len(t, g.links, 2)
	})
}

func TestFirstWalk(t *testing.T) {

	t.Run("should", func(t *testing.T) {
		r := NewDependencyNode("r", "Root")
		a := NewDependencyNode("a", "A")
		b := NewDependencyNode("b", "B")
		c := NewDependencyNode("c", "C")

		g := &graphService{
			nodes: NodeList{r, a, b, c},
			links: LinkList{
				NewLink("r", "a"),
				NewLink("r", "b"),
				NewLink("b", "c"),
			},
			root:     r,
			layout:   Horizontal,
			distance: 1,
		}

		g.prepareNodes()
		g.prepareLinks()
		root, err := g.convertToTree(r, 0, nil, 1)
		at := root.children[0]
		bt := root.children[0]
		ct := bt.children[0]
		assert.Nil(t, err)
		g.firstWalk(root)
		assert.Equal(t, float64(0), at.prelimX)
		assert.Equal(t, float64(0), bt.prelimX)
		assert.Equal(t, float64(0), ct.prelimX)
		assert.Equal(t, float64(0.5), root.prelimX)
	})
}
