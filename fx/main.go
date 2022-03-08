package main

import (
	"io"
	"os"

	lib "github.com/nmasse-itix/golang-fx"
	"go.uber.org/fx"
)

func childs() ([]lib.Child, error) {
	c1, err := lib.NewChild(2)
	if err != nil {
		return nil, err
	}

	c2, err := lib.NewChild(5)
	if err != nil {
		return nil, err
	}

	return []lib.Child{c1, c2}, nil
}

func main() {
	fx.New(
		fx.Supply(
			fx.Annotated{Name: "house_name", Target: "New-York"},
			fx.Annotated{Name: "john_name", Target: "John"},
			fx.Annotated{Name: "jane_name", Target: "Jane"},
			fx.Annotated{Name: "cat1_name", Target: "Gros Minet"},
			fx.Annotated{Name: "cat2_name", Target: "Isidore"},
		),
		fx.Provide(
			fx.Annotate(lib.NewCat, fx.ParamTags(`name:"cat1_name"`), fx.ResultTags(`group:"cats"`)),
			fx.Annotate(lib.NewCat, fx.ParamTags(`name:"cat2_name"`), fx.ResultTags(`group:"cats"`)),
			fx.Annotate(lib.NewAdult, fx.ParamTags(`name:"john_name"`, `group:"childs"`), fx.ResultTags(`group:"adults"`)),
			fx.Annotate(lib.NewAdult, fx.ParamTags(`name:"jane_name"`, `group:"childs"`), fx.ResultTags(`group:"adults"`)),
			fx.Annotate(childs, fx.ResultTags(`group:"childs"`)),
			fx.Annotate(lib.NewHouse, fx.ParamTags(`name:"house_name"`, `group:"adults"`, `group:"cats"`)),
			func() io.Writer { return os.Stdout },
		),
		fx.Invoke(lib.House.Present),
		fx.NopLogger,
	).Run()
}
