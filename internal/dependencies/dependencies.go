package dependencies

import (
	"go.uber.org/dig"
)

type HanldersContainer struct{}

func NewWire() *dig.Container {
	container := dig.New()
	return container
}
