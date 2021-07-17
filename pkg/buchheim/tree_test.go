// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package buchheim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	t.Parallel()
	t.Run("Should create a tree", func(t *testing.T) {
		t.Parallel()
		root := NewDependencyNode("r", "Root")
		tree := NewTree(root, 0, nil, 1)
		assert.Equal(t, "r", tree.id)
	})

	t.Run("Should create a tree with parent node", func(t *testing.T) {
		t.Parallel()
		root := NewDependencyNode("r", "Root")
		tr := NewTree(root, 0, nil, 1)
		child := NewDependencyNode("c", "Child")
		tc := NewTree(child, 0, tr, 2)
		assert.Equal(t, "c", tc.id)
	})
}

func TestIsEqual(t *testing.T) {
	t.Parallel()
	t.Run("a and b should be equal", func(t *testing.T) {
		t.Parallel()
		a := NewDependencyNode("a", "A")
		b := NewDependencyNode("a", "A")

		at := NewTree(a, 0, nil, 1)
		bt := NewTree(b, 0, nil, 1)

		assert.True(t, at.IsEqual(*bt))
		assert.True(t, bt.IsEqual(*at))
	})
	t.Run("a and b should not be equal", func(t *testing.T) {
		t.Parallel()
		a := NewDependencyNode("a", "A")
		b := NewDependencyNode("b", "A")

		at := NewTree(a, 0, nil, 1)
		bt := NewTree(b, 0, nil, 1)

		assert.False(t, at.IsEqual(*bt))
		assert.False(t, bt.IsEqual(*at))
	})
}

func TestSibling(t *testing.T) {
	t.Parallel()
	r := NewDependencyNode("r", "Root")
	a := NewDependencyNode("a", "A")
	b := NewDependencyNode("b", "B")
	c := NewDependencyNode("c", "C")
	d := NewDependencyNode("d", "D")

	rt := NewTree(r, 0, nil, 1)
	at := NewTree(a, 1, rt, 2)
	bt := NewTree(b, 1, rt, 2)
	ct := NewTree(c, 1, rt, 2)
	dt := NewTree(d, 1, rt, 2)
	rt.AddChild(at)
	rt.AddChild(bt)
	rt.AddChild(ct)
	rt.AddChild(dt)

	testCases := []struct {
		description  string
		selectedTree *Tree
		expectedTree *Tree
	}{
		// TODO To inspect if that is normal
		// {
		// 	description:  "Should A like the most left sibling for A",
		// 	selectedTree: at,
		// 	expectedTree: at,
		// },
		{
			description:  "Should A like the most left sibling for B",
			selectedTree: bt,
			expectedTree: at,
		},
		{
			description:  "Should A like the most left sibling for C",
			selectedTree: ct,
			expectedTree: at,
		},
		{
			description:  "Should A like the most left sibling for D",
			selectedTree: dt,
			expectedTree: at,
		},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			mls := test.selectedTree.GetLeftMostSibling()
			assert.Equal(t, test.expectedTree, mls)
		})
	}

	testCases = []struct {
		description  string
		selectedTree *Tree
		expectedTree *Tree
	}{

		{
			description:  "Should return NIL like the left sibling for A",
			selectedTree: at,
			expectedTree: nil,
		},
		{
			description:  "Should return A like the left sibling for B",
			selectedTree: bt,
			expectedTree: at,
		},
		{
			description:  "Should return B like the left sibling for C",
			selectedTree: ct,
			expectedTree: bt,
		},
		{
			description:  "Should return C like the left sibling for D",
			selectedTree: dt,
			expectedTree: ct,
		},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			ls := test.selectedTree.GetLeftSibling()
			assert.Equal(t, test.expectedTree, ls)
		})
	}
}

func TestMidPoint(t *testing.T) {
	r := NewDependencyNode("r", "Root")
	a := NewDependencyNode("a", "A")
	b := NewDependencyNode("b", "B")
	c := NewDependencyNode("c", "C")
	d := NewDependencyNode("d", "D")

	rt := NewTree(r, 0, nil, 1)
	at := NewTree(a, 1, rt, 2)
	at.SetPrelimX(0)
	bt := NewTree(b, 1, rt, 2)
	ct := NewTree(c, 1, rt, 2)
	dt := NewTree(d, 1, rt, 2)
	dt.SetPrelimX(4)
	rt.AddChild(at)
	rt.AddChild(bt)
	rt.AddChild(ct)
	rt.AddChild(dt)

	midpoint := rt.MidPoint()
	assert.Equal(t, 2.0, midpoint)
}

func TestNextLeftAndRight(t *testing.T) {
	t.Parallel()
	a := NewDependencyNode("a", "A")
	b := NewDependencyNode("b", "B")
	c := NewDependencyNode("c", "C")
	d := NewDependencyNode("d", "D")

	at := NewTree(a, 0, nil, 1)
	bt := NewTree(b, 1, at, 2)
	ct := NewTree(c, 1, at, 2)
	dt := NewTree(d, 1, at, 2)
	at.AddChild(bt)
	at.AddChild(ct)
	at.AddChild(dt)

	t.Run("Should return B the next left tree", func(t *testing.T) {
		t.Parallel()
		left := at.nextLeft()
		assert.Equal(t, bt, left)
	})

	t.Run("Should return D the next right tree", func(t *testing.T) {
		t.Parallel()
		right := at.nextRight()
		assert.Equal(t, dt, right)
	})

	t.Run("Should return nil the next right tree", func(t *testing.T) {
		t.Parallel()
		right := ct.nextRight()
		assert.Nil(t, right)
	})

	t.Run("Should return nil the next left tree", func(t *testing.T) {
		t.Parallel()
		left := ct.nextLeft()
		assert.Nil(t, left)
	})
}
