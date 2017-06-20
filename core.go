// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

type BasicNode struct {
	o *js.Object
}

type BasicElement struct {
	*BasicNode
}

func newBasicElement() *BasicElement {
	return &BasicElement{
		BasicNode: &BasicNode{object.New()},
	}
}

type BasicHTMLElement struct {
	*BasicElement

	ID        string `js:"id" react:"omitempty"`
	Key       string `js:"key" react:"omitempty"`
	ClassName string `js:"className"`
	Role      string `js:"role"`
	Style     *CSS   `js:"style"`
	Title     string `js:"title"`

	OnChange `js:"onChange"`
	OnClick  `js:"onClick"`
	// mouse events
	OnMouseOut   `js:"onMouseOut"`
	OnMouseOver  `js:"onMouseOver"`
	OnMouseDown  `js:"onMouseDown"`
	OnMouseMove  `js:"onMouseMove"`
	OnMouseUp    `js:"onMouseUp"`
	OnMouseEnter `js:"onMouseEnter"`
	OnMouseLeave `js:"onMouseLeave"`

	DangerouslySetInnerHTML *DangerousInnerHTMLDef `js:"dangerouslySetInnerHTML"`
}

func newBasicHTMLElement() *BasicHTMLElement {
	return &BasicHTMLElement{
		BasicElement: newBasicElement(),
	}
}

// TODO complete the definition
type SyntheticEvent struct {
	o *js.Object

	PreventDefault  func() `js:"preventDefault"`
	StopPropagation func() `js:"stopPropagation"`
}

func (s *SyntheticEvent) Target() dom.HTMLElement {
	return dom.WrapHTMLElement(s.o.Get("target"))
}

type SyntheticMouseEvent struct {
	*SyntheticEvent

	ClientX int `js:"clientX"`
	PageX   int `js:"pageX"`
	PageY   int `js:"pageY"`
}
