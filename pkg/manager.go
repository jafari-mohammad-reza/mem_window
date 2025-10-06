package pkg

import "sync/atomic"

type WindowItem[T any] struct {
	Data       T
	InsertedAt int64 // timestamp in unix nano
}

type CreateWindOpts struct {
}

type DeleteWindsOpt struct {
}

type ListWindsOpt struct {
}
type NewWindowingOpts struct {
}
type WindowManagerOpts struct{}
type Window[T any] struct {
	ID        string
	Items     []WindowItem[*T]
	Len       atomic.Int32
	StartedAt int64 // timestamp in unix nano
	EndedAt   int64 // timestamp in unix nano
}

type WindTrigger[T any, W any] func(wind *WindowManager[T, W]) error

type CompressedWindow[T any] struct {
	ID        string
	Items     [][]byte
	StartedAt int64 // timestamp in unix nano
	EndedAt   int64 // timestamp in unix nano
}

type Manager[T any, W any] interface {
	NewWindowing(sourceChan chan WindowItem[*T], opts NewWindowingOpts) error
	CreateWind(opts CreateWindOpts) error
	DeleteWind(opts DeleteWindsOpt) error
	ListWinds(opts ListWindsOpt) error
	RegisterTrigger(trigger WindTrigger[T, W]) error
}

// set window one of type of Window or CompressedWindow
type WindowManager[T any, W any] struct {
	Windows  atomic.Pointer[[]*W]
	MaxWind  uint32
	Len      atomic.Uint32
	Triggers []WindTrigger[T, W] // trigger are methods that check to see if are needed to get called on each change event
}

func NewWindowManager[T any, W any](opts WindowManagerOpts) (Manager[T, W], error) {
	return &WindowManager[T, W]{}, nil
}

func (w *WindowManager[T, W]) NewWindowing(sourceChan chan WindowItem[*T], opts NewWindowingOpts) error {
	return nil
}

func (w *WindowManager[T, W]) CreateWind(opts CreateWindOpts) error {
	return nil
}
func (w *WindowManager[T, W]) DeleteWind(opts DeleteWindsOpt) error {
	return nil
}
func (w *WindowManager[T, W]) ListWinds(opts ListWindsOpt) error {
	return nil
}
func (w *WindowManager[T, W]) RegisterTrigger(trigger WindTrigger[T, W]) error {
	w.Triggers = append(w.Triggers, trigger)
	return nil
}
