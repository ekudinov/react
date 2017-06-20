package react

import "github.com/gopherjs/gopherjs/js"

// ImgDef is the React component definition corresponding to the HTML <img> element
type ImgDef struct {
	underlying *js.Object
}

// _ImgProps defines the properties for the <label> element
type _ImgProps struct {
	*BasicHTMLElement

	Align string `js:"align"`

	Alt string `js:"alt"`

	Border string `js:"border"`

	Height string `js:"height"`

	Hspace string `js:"hspace"`

	Ismap string `js:"ismap"`

	Longdesc string `js:"longdesc"`

	Lowsrc string `js:"lowsrc"`

	Src string `js:"src"`

	Vspace string `js:"vspace"`

	Width string `js:"string"`

	Usemap string `js:"string"`
}

func (d *ImgDef) reactElement() {}

// Img creates a new instance of a <img> element with the provided props and child
// element
func Img(props *ImgProps) *ImgDef {

	rProps := &_ImgProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	underlying := react.Call("createElement", "img", rProps)

	return &ImgDef{underlying: underlying}
}
