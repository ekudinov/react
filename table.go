package react

import "github.com/gopherjs/gopherjs/js"

// TableDef is the React component definition corresponding to the HTML <table> element
type TableDef struct {
	underlying *js.Object
}

// _TableProps defines the properties for the <table> element
type _TableProps struct {
	*BasicHTMLElement
}

func (d *TableDef) reactElement() {}

// Ul creates a new instance of a <table> element with the provided props and
// children
func Table(props *TableProps, children ...Element) *TableDef {

	rProps := &_TableProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	args := []interface{}{"table", rProps}

	for _, v := range children {
		args = append(args, elementToReactObj(v))
	}

	underlying := react.Call("createElement", args...)

	return &TableDef{underlying: underlying}
}
