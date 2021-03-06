// Code generated by reactGen. DO NOT EDIT.

package react

// UlProps defines the properties for the <ul> element
type UlProps struct {
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

func (u *UlProps) assign(v *_UlProps) {

	v.ClassName = u.ClassName

	v.DangerouslySetInnerHTML = u.DangerouslySetInnerHTML

	if u.ID != "" {
		v.ID = u.ID
	}

	if u.Key != "" {
		v.Key = u.Key
	}

	if u.OnChange != nil {
		v.o.Set("onChange", u.OnChange.OnChange)
	}

	if u.OnClick != nil {
		v.o.Set("onClick", u.OnClick.OnClick)
	}

	if u.OnMouseDown != nil {
		v.o.Set("onMouseDown", u.OnMouseDown.OnMouseDown)
	}

	if u.OnMouseEnter != nil {
		v.o.Set("onMouseEnter", u.OnMouseEnter.OnMouseEnter)
	}

	if u.OnMouseLeave != nil {
		v.o.Set("onMouseLeave", u.OnMouseLeave.OnMouseLeave)
	}

	if u.OnMouseMove != nil {
		v.o.Set("onMouseMove", u.OnMouseMove.OnMouseMove)
	}

	if u.OnMouseOut != nil {
		v.o.Set("onMouseOut", u.OnMouseOut.OnMouseOut)
	}

	if u.OnMouseOver != nil {
		v.o.Set("onMouseOver", u.OnMouseOver.OnMouseOver)
	}

	if u.OnMouseUp != nil {
		v.o.Set("onMouseUp", u.OnMouseUp.OnMouseUp)
	}

	v.Role = u.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = u.Style.hack()

	v.Title = u.Title

}
