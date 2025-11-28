package interfaces

import (
	"zotes/shared/structs"

	"github.com/a-h/templ"
)

type CustomTable interface {
	Headers() templ.Component
	Overlay() templ.Component
	RowOverlay() templ.Component
	Footer() templ.Component
	Header() templ.Component
	Empty() templ.Component
}

type CustomTableConstraint[T any] interface {
	structs.Common
	CustomTable
	Row(T) templ.Component
	Data() []T
}
