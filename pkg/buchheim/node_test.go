// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package buchheim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {

	a := NewDependencyNode("a", "A")
	b := NewDependencyNode("b", "B")
	c := NewDependencyNode("c", "C")
	d := NewDependencyNode("d", "D")

	nodes := NodeList{a, b, c, d}

	var testCases = []struct {
		description string
		searchID    string
		expected    Node
	}{
		{
			description: "should find the node a",
			searchID:    "a",
			expected:    a,
		},

		{
			description: "should find the node b",
			searchID:    "b",
			expected:    b,
		},

		{
			description: "should find the node c",
			searchID:    "c",
			expected:    c,
		},
	}

	for _, c := range testCases {
		t.Run(c.description, func(t *testing.T) {
			n, err := nodes.FindByID(c.searchID)
			assert.Nil(t, err)
			assert.Equal(t, c.expected, n)
		})
	}
}

func TestAddOutgoingLink(t *testing.T) {
	t.Run("Should have 0 out going link", func(t *testing.T) {
		a := NewDependencyNode("a", "A")
		assert.Len(t, a.OutgoingLink(), 0)
	})

	t.Run("Should have 1 out going link", func(t *testing.T) {
		link := NewLink("a", "b")
		a := NewDependencyNode("a", "A")
		a.AddOutgoingLink(link)
		assert.Len(t, a.OutgoingLink(), 1)
		assert.Equal(t, link, a.OutgoingLink()[0])
	})
}

func TestNodeListAdd(t *testing.T) {
	t.Run("Should add a node in the list", func(t *testing.T) {
		list := NodeList{}
		n := NewDependencyNode("a", "b")
		list.Add(n)
		assert.Len(t, list, 1)
	})

	t.Run("Should doesn't add a node in the list", func(t *testing.T) {
		list := NodeList{NewDependencyNode("a", "b"), NewDependencyNode("b", "c")}
		n := NewDependencyNode("a", "b")
		list.Add(n)
		assert.Len(t, list, 2)
	})
}
