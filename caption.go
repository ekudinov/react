package react

import "github.com/gopherjs/gopherjs/js"

// CaptionDef is the React component definition corresponding to the HTML <caption> element
type CaptionDef struct {
	underlying *js.Object
}

// _CaptionProps defines the properties for the <caption> element
type _CaptionProps struct {
	*BasicHTMLElement
}

func (d *CaptionDef) reactElement() {}

// Caption creates a new instance of an <caption> element with the provided props
// and children
func Caption(props *CaptionProps, children ...Element) *CaptionDef {

	rProps := &_CaptionProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"caption", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &CaptionDef{underlying: underlying}
}
