// Code generated by reactGen. DO NOT EDIT.

package react

// OptionProps defines the properties for the <option> element
type OptionProps struct {
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
	Value string
}

func (o *OptionProps) assign(v *_OptionProps) {

	v.ClassName = o.ClassName

	v.DangerouslySetInnerHTML = o.DangerouslySetInnerHTML

	if o.ID != "" {
		v.ID = o.ID
	}

	if o.Key != "" {
		v.Key = o.Key
	}

	if o.OnChange != nil {
		v.o.Set("onChange", o.OnChange.OnChange)
	}

	if o.OnClick != nil {
		v.o.Set("onClick", o.OnClick.OnClick)
	}

	if o.OnMouseDown != nil {
		v.o.Set("onMouseDown", o.OnMouseDown.OnMouseDown)
	}

	if o.OnMouseEnter != nil {
		v.o.Set("onMouseEnter", o.OnMouseEnter.OnMouseEnter)
	}

	if o.OnMouseLeave != nil {
		v.o.Set("onMouseLeave", o.OnMouseLeave.OnMouseLeave)
	}

	if o.OnMouseMove != nil {
		v.o.Set("onMouseMove", o.OnMouseMove.OnMouseMove)
	}

	if o.OnMouseOut != nil {
		v.o.Set("onMouseOut", o.OnMouseOut.OnMouseOut)
	}

	if o.OnMouseOver != nil {
		v.o.Set("onMouseOver", o.OnMouseOver.OnMouseOver)
	}

	if o.OnMouseUp != nil {
		v.o.Set("onMouseUp", o.OnMouseUp.OnMouseUp)
	}

	v.Role = o.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = o.Style.hack()

	v.Title = o.Title

	v.Value = o.Value

}
