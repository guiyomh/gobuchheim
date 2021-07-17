// Copyright (c) 2021 Guillaume CAMUS
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// Package treechart generates SVG that represents a tree of nodes and links
package treechart

import (
	"io"

	svg "github.com/ajstarks/svgo/float"
)

const (
	defaultHeight float64 = 500
	defaultWidth  float64 = 500
	yScale        float64 = 20  // 60
	xScale        float64 = 100 // 150
	xModifier     float64 = 100
	labelMargin   float64 = 6
	dotSize       float64 = 6
)

type graphOption func(*TreeChart)

// TreeChart is a service that draws the tree
type TreeChart struct {
	canvas  *svg.SVG
	nodes   []Node
	links   []Link
	height  float64
	width   float64
	yScale  float64
	xScale  float64
	offset  float64
	dotSize float64
}

func WithHeight(h int) graphOption {
	return func(t *TreeChart) {
		t.height = float64(h)
	}
}

func WithWidth(w int) graphOption {
	return func(t *TreeChart) {
		t.width = float64(w)
	}
}

func WithXScale(x int) graphOption {
	return func(t *TreeChart) {
		t.xScale = float64(x)
	}
}

func WithYScale(y int) graphOption {
	return func(t *TreeChart) {
		t.yScale = float64(y)
	}
}

func WithDotSize(s int) graphOption {
	return func(t *TreeChart) {
		t.dotSize = float64(s)
	}
}

// NewTreeChart creates a instance of Treechart service
func NewTreeChart(nodes []Node, links []Link, w io.Writer, opts ...graphOption) Chart {
	t := &TreeChart{
		nodes:   nodes,
		links:   links,
		canvas:  svg.New(w),
		height:  defaultHeight,
		width:   defaultWidth,
		yScale:  yScale,
		xScale:  xScale,
		dotSize: dotSize,
	}
	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (tc *TreeChart) Draw() {
	tc.canvas.Start(tc.width, tc.height)
	defer tc.canvas.End()
	tc.drawLinks()
	tc.drawNodes()
}

func (tc TreeChart) drawNodes() {
	for _, n := range tc.nodes {
		tc.drawNode(n)
	}
}

func (tc TreeChart) drawNode(n Node) {
	x := tc.scaleX(n.X())
	y := tc.scaleY(n.Y())
	tc.canvas.Circle(
		x,
		y,
		tc.dotSize,
		"fill:gray;stroke:black;stroke-weight:1",
	)
	tc.canvas.Text(x+labelMargin, y+labelMargin, n.Label(), "fill:white;font-size:12")
}

func (tc TreeChart) drawLinks() {
	for _, link := range tc.links {
		tc.drawLink(link.Source(), link.Target())
	}
}

func (tc TreeChart) drawLink(source, target Node) {

	var quarterDivider float64 = 4
	var treeQuarter float64 = 3

	x3 := target.X()
	y3 := source.Y() + (target.Y()-source.Y())/quarterDivider

	x4 := source.X()
	y4 := source.Y() + (target.Y()-source.Y())/quarterDivider*treeQuarter

	// tc.canvas.Circle(x3, y3, 3, "fill:purple")
	// tc.canvas.Circle(x4, y4, 3, "fill:green")
	tc.canvas.Bezier(
		tc.scaleX(source.X()),
		tc.scaleY(source.Y()),
		tc.scaleX(x3),
		tc.scaleY(y3),
		tc.scaleX(x4),
		tc.scaleY(y4),
		tc.scaleX(target.X()),
		tc.scaleY(target.Y()),
		"fill:transparent;stroke-width:2; stroke:red",
	)
}

func (tc TreeChart) scaleX(x float64) float64 {
	return x*tc.xScale + xModifier
}

func (tc TreeChart) scaleY(y float64) float64 {
	return y*tc.yScale + tc.offset
}
