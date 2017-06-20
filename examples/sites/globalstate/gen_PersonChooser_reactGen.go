// Code generated by reactGen. DO NOT EDIT.

package main

import "myitcv.io/react"

func (p *PersonChooserDef) ShouldComponentUpdateIntf(nextProps, prevState, nextState interface{}) bool {
	res := false

	{
		res = p.Props() != nextProps.(PersonChooserProps) || res
	}
	v := prevState.(PersonChooserState)
	res = !v.EqualsIntf(nextState) || res
	return res
}

// SetState is an auto-generated proxy proxy to update the state for the
// PersonChooser component.  SetState does not immediately mutate p.State()
// but creates a pending state transition.
func (p *PersonChooserDef) SetState(state PersonChooserState) {
	p.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the PersonChooser component
func (p *PersonChooserDef) State() PersonChooserState {
	return p.ComponentDef.State().(PersonChooserState)
}

// IsState is an auto-generated definition so that PersonChooserState implements
// the myitcv.io/react.State interface.
func (p PersonChooserState) IsState() {}

var _ react.State = PersonChooserState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (p *PersonChooserDef) GetInitialStateIntf() react.State {
	return PersonChooserState{}
}

func (p PersonChooserState) EqualsIntf(val interface{}) bool {
	return p == val.(PersonChooserState)
}

// Props is an auto-generated proxy to the current props of PersonChooser
func (p *PersonChooserDef) Props() PersonChooserProps {
	uprops := p.ComponentDef.Props()
	return uprops.(PersonChooserProps)
}

func (p PersonChooserProps) EqualsIntf(val interface{}) bool {
	return p == val.(PersonChooserProps)
}

var _ react.Equals = PersonChooserProps{}
