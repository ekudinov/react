package react

import "github.com/gopherjs/gopherjs/js"

// TdDef is the React component definition corresponding to the HTML <td> element
type TdDef struct {
	underlying *js.Object
}

// _TdProps defines the properties for the <td> element
type _TdProps struct {
	*BasicHTMLElement
}

func (d *TdDef) reactElement() {}

// Td creates a new instance of an <td> element with the provided props
// and children
func Td(props *TdProps, children ...Element) *TdDef {

	rProps := &_TdProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"td", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &TdDef{underlying: underlying}
}
