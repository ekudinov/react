// Code generated by reactGen. DO NOT EDIT.

package react

// FormProps defines the properties for the <form> element
type FormProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
	ID                      string
	Key                     string

	OnChange
	OnClick
	OnMouseDown
	OnMouseEnter
	OnMouseLeave
	OnMouseMove
	OnMouseOut
	OnMouseOver
	OnMouseUp

	Role  string
	Style *CSS
	Title string
}

func (f *FormProps) assign(v *_FormProps) {

	v.ClassName = f.ClassName

	v.DangerouslySetInnerHTML = f.DangerouslySetInnerHTML

	if f.ID != "" {
		v.ID = f.ID
	}

	if f.Key != "" {
		v.Key = f.Key
	}

	if f.OnChange != nil {
		v.o.Set("onChange", f.OnChange.OnChange)
	}

	if f.OnClick != nil {
		v.o.Set("onClick", f.OnClick.OnClick)
	}

	if f.OnMouseDown != nil {
		v.o.Set("onMouseDown", f.OnMouseDown.OnMouseDown)
	}

	if f.OnMouseEnter != nil {
		v.o.Set("onMouseEnter", f.OnMouseEnter.OnMouseEnter)
	}

	if f.OnMouseLeave != nil {
		v.o.Set("onMouseLeave", f.OnMouseLeave.OnMouseLeave)
	}

	if f.OnMouseMove != nil {
		v.o.Set("onMouseMove", f.OnMouseMove.OnMouseMove)
	}

	if f.OnMouseOut != nil {
		v.o.Set("onMouseOut", f.OnMouseOut.OnMouseOut)
	}

	if f.OnMouseOver != nil {
		v.o.Set("onMouseOver", f.OnMouseOver.OnMouseOver)
	}

	if f.OnMouseUp != nil {
		v.o.Set("onMouseUp", f.OnMouseUp.OnMouseUp)
	}

	v.Role = f.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = f.Style.hack()

	v.Title = f.Title

}
