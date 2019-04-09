package bitmex

import "github.com/chuckpreslar/emission"

//On adds a listener to a specific event
func (b *BitMEX) On(event interface{}, listener interface{}) *emission.Emitter {
	return b.emitter.On(event, listener)
}

//Emit emits an event
func (b *BitMEX) Emit(event interface{}, arguments ...interface{}) *emission.Emitter {
	return b.emitter.Emit(event, arguments...)
}

//Off removes a listener for an event
func (b *BitMEX) Off(event interface{}, listener interface{}) *emission.Emitter {
	return b.emitter.Off(event, listener)
}
