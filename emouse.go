package react

type OnMouseOver interface {
	Event

	OnMouseOver(e *SyntheticMouseEvent)
}

type OnMouseOut interface {
	Event

	OnMouseOut(e *SyntheticMouseEvent)
}

type OnMouseDown interface {
	Event

	OnMouseDown(e *SyntheticMouseEvent)
}

type OnMouseMove interface {
	Event

	OnMouseMove(e *SyntheticMouseEvent)
}

type OnMouseUp interface {
	Event

	OnMouseUp(e *SyntheticMouseEvent)
}

type OnMouseEnter interface {
	Event

	OnMouseEnter(e *SyntheticMouseEvent)
}

type OnMouseLeave interface {
	Event

	OnMouseLeave(e *SyntheticMouseEvent)
}
