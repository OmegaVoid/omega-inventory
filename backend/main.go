//go:generate go get -d github.com/99designs/gqlgen
//go:generate go run github.com/99designs/gqlgen --config configs/gqlgen.yml
//go:generate go get -d github.com/99designs/gqlgen

package main

import (
	"github.com/OmegaVoid/omega-inventory/cmd"
)

func main() {
	cmd.Execute()
}
