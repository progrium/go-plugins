package main

import (
	"strings"

	"github.com/progrium/go-extpoints/tests/extpoints"
)

func init() {
	extpoints.Register(new(noop), "noop")                  // Noop
	extpoints.Register(new(noop2), "noop2")                // Noop
	extpoints.Register(new(uppercaseTransformer), "upper") // StringTransformer
	extpoints.NoopFactories.Register(noopFactory, "")
}

func noopFactory() extpoints.Noop {
	return new(noop)
}

type noop struct{}

func (n *noop) Noop() string {
	return "noop"
}

type noop2 struct{}

func (n *noop2) Noop() string {
	return "noop2"
}

type uppercaseTransformer struct{}

func (t *uppercaseTransformer) Transform(input string) string {
	return strings.ToUpper(input)
}
