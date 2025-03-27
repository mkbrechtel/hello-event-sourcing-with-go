package event

import (
	"fmt"
	"reflect"

	"github.com/gofrs/uuid/v5"
)

type Event[T any] struct {
	ID      uuid.UUID `json:"id"`
	Type    string    `json:"type"`
	Source  string    `json:"source"`
	Version int16     `json:"version"`
	Details T         `json:"details"`
}

func NewEvent[T any](details T, source string, version int16) Event[T] {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	e := Event[T]{
		ID:      id,
		Source:  source,
		Version: version,
		Details: details,
	}
	e.Type = getType(e)
	return e
}

func getType(v any) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		return t.Elem().PkgPath() + "." + t.Elem().Name()
	}
	return t.PkgPath() + "." + t.Name()
}

// String returns a text representation of the event
func (e Event[T]) String() string {
	details := reflect.ValueOf(e.Details)
	return "Event{ID: " + e.ID.String() + ", Source: " + e.Source + ", Version: " + fmt.Sprintf("%d", e.Version) + ", Type: " + e.Type + ", Details: " + fmt.Sprintf("%+v", details) + "}"
}
