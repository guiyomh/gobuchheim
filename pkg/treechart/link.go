// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package treechart

import "github.com/guiyomh/gobuchheim/pkg/buchheim"

type Link interface {
	Source() Node
	Target() Node
}

type link struct {
	source Node
	target Node
}

func (l link) Source() Node {
	return l.source
}

func (l link) Target() Node {
	return l.target
}

func ConvertBuchheimLinkList(links buchheim.LinkList) []Link {
	treeLinks := make([]Link, len(links))
	for i := range links {
		sn := node{
			x:     links[i].Source().X(),
			y:     links[i].Source().Y(),
			label: links[i].Source().Label(),
		}
		tn := node{
			x:     links[i].Target().X(),
			y:     links[i].Target().Y(),
			label: links[i].Target().Label(),
		}
		treeLinks[i] = link{
			source: sn,
			target: tn,
		}
	}

	return treeLinks
}
