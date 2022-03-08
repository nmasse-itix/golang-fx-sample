package lib

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type Cat interface {
	String() string
}

type MyCat struct {
	Name string
}

func (c *MyCat) String() string {
	return fmt.Sprintf("Cat: %s", c.Name)
}

func NewCat(name string) Cat {
	log.Default().Printf("NewCat(%s)", name)
	return &MyCat{Name: name}
}

type Child interface {
	String() string
}

type MyChild struct {
	Age int
}

func NewChild(age int) (Child, error) {
	log.Default().Printf("NewChild(%d)", age)
	if age < 0 {
		return nil, fmt.Errorf("wrong age")
	}

	return &MyChild{Age: age}, nil
}

func (c *MyChild) String() string {
	return fmt.Sprintf("%d year old child", c.Age)
}

type Adult interface {
	String() string
}

type MyAdult struct {
	Name   string
	Childs []Child
}

func NewAdult(name string, childs []Child) Adult {
	log.Default().Printf("NewAdult(%s, %d childs)", name, len(childs))
	return &MyAdult{Name: name, Childs: childs}
}

func (a *MyAdult) String() string {
	var b strings.Builder
	b.WriteString(a.Name)
	if len(a.Childs) > 0 {
		b.WriteString(" (with ")
		for i, c := range a.Childs {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(c.String())
		}
		b.WriteString(")")
	}
	return b.String()
}

type House interface {
	Present(w io.Writer)
}

type MyHouse struct {
	Address string
	Adults  []Adult
	Cats    []Cat
}

func (h *MyHouse) Present(w io.Writer) {
	fmt.Fprintf(w, "House at: %s\n", h.Address)
	fmt.Fprintln(w, "with:")
	for _, a := range h.Adults {
		fmt.Fprintf(w, "- %s\n", a)
	}
	fmt.Fprintln(w, "and:")
	for _, c := range h.Cats {
		fmt.Fprintf(w, "- %s\n", c)
	}
	fmt.Fprintln(w)
}

func NewHouse(address string, adults []Adult, cats []Cat) House {
	log.Default().Printf("NewHouse(%s, %d adults, %d cats)", address, len(adults), len(cats))
	return &MyHouse{
		Address: address,
		Adults:  adults,
		Cats:    cats,
	}
}
