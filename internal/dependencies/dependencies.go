package dependencies

import (
	"go.uber.org/dig"
)

type HandlersContainer struct{}

func NewWire() *dig.Container {
	container := dig.New()

	container.Provide(func() *HandlersContainer {
		return &HandlersContainer{}
	})
	return container
}
