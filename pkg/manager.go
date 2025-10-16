package pkg

import (
	"sync/atomic"
	"time"
)

type WindowItem[T any] struct {
	Data       T
	InsertedAt time.Time
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
	StartedAt time.Time
	EndedAt   time.Time
}

type WindTrigger[T any, W any] func(wind *WindowManager[T, W]) error

type CompressedWindow[T any] struct {
	ID        string
	Items     [][]byte
	StartedAt time.Time
	EndedAt   time.Time
}

func (c *CompressedWindow[T]) GetItems() ([]WindowItem[*T], error) {
	items := make([]WindowItem[*T], len(c.Items))
	for _, item := range c.Items {
		windItem, err := DecompressWind[T](item)
		if err != nil {
			return nil, err
		}
		if windItem.InsertedAt.IsZero() {
			continue // skip empty items
		}
		items = append(items, windItem)
	}
	return items, nil
}
func (c *CompressedWindow[T]) AppendItem(item WindowItem[*T]) error {
	compressedItem, err := CompressWind(item)
	if err != nil {
		return err
	}
	c.Items = append(c.Items, compressedItem)
	return nil
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
