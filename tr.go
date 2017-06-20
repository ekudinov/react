package react

import "github.com/gopherjs/gopherjs/js"

// TrDef is the React component definition corresponding to the HTML <tr> element
type TrDef struct {
	underlying *js.Object
}

// _TrProps defines the properties for the <tr> element
type _TrProps struct {
	*BasicHTMLElement
}

func (d *TrDef) reactElement() {}

// Tr creates a new instance of an <tr> element with the provided props
// and children as <td>
func Tr(props *TrProps, children ...Element) *TrDef {

	rProps := &_TrProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"tr", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &TrDef{underlying: underlying}
}
