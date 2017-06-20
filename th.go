package react

import "github.com/gopherjs/gopherjs/js"

// ThDef is the React component definition corresponding to the HTML <th> element
type ThDef struct {
	underlying *js.Object
}

// _ThProps defines the properties for the <th> element
type _ThProps struct {
	*BasicHTMLElement
}

func (d *ThDef) reactElement() {}

// Th creates a new instance of an <td> element with the provided props
// and children
func Th(props *ThProps, children ...Element) *ThDef {

	rProps := &_ThProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"th", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &ThDef{underlying: underlying}
}
