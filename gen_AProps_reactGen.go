// Code generated by reactGen. DO NOT EDIT.

package react

// APropsDef defines the properties for the <a> element
type AProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
	Href                    string
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

	Role   string
	Style  *CSS
	Target string
	Title  string
}

func (a *AProps) assign(v *_AProps) {

	v.ClassName = a.ClassName

	v.DangerouslySetInnerHTML = a.DangerouslySetInnerHTML

	v.Href = a.Href

	if a.ID != "" {
		v.ID = a.ID
	}

	if a.Key != "" {
		v.Key = a.Key
	}

	if a.OnChange != nil {
		v.o.Set("onChange", a.OnChange.OnChange)
	}

	if a.OnClick != nil {
		v.o.Set("onClick", a.OnClick.OnClick)
	}

	if a.OnMouseDown != nil {
		v.o.Set("onMouseDown", a.OnMouseDown.OnMouseDown)
	}

	if a.OnMouseEnter != nil {
		v.o.Set("onMouseEnter", a.OnMouseEnter.OnMouseEnter)
	}

	if a.OnMouseLeave != nil {
		v.o.Set("onMouseLeave", a.OnMouseLeave.OnMouseLeave)
	}

	if a.OnMouseMove != nil {
		v.o.Set("onMouseMove", a.OnMouseMove.OnMouseMove)
	}

	if a.OnMouseOut != nil {
		v.o.Set("onMouseOut", a.OnMouseOut.OnMouseOut)
	}

	if a.OnMouseOver != nil {
		v.o.Set("onMouseOver", a.OnMouseOver.OnMouseOver)
	}

	if a.OnMouseUp != nil {
		v.o.Set("onMouseUp", a.OnMouseUp.OnMouseUp)
	}

	v.Role = a.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = a.Style.hack()

	v.Target = a.Target

	v.Title = a.Title

}
