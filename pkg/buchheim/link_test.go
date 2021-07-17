// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package buchheim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveLink(t *testing.T) {
	t.Run("Should return a new lnk list", func(t *testing.T) {
		link1 := NewLink("a", "b")
		link2 := NewLink("b", "c")
		link3 := NewLink("a", "c")
		link4 := NewLink("c", "z")

		list := LinkList{link1, link2, link3, link4}
		list = list.Remove(2)
		assert.Equal(t, len(list), 3)
		assert.Equal(t, list[0], link1)
		assert.Equal(t, list[1], link2)
		assert.Equal(t, list[2], link4)

	})
}
