package main

import (
	"fmt"
	"strings"
)

// This awesome example is from @davedev. From his slide deck preso here:
//
// 	Read This URI >>
// 	TODO: https://meetup.euggo.org/talks/objectively_harmful/objectively_harmful.slide#47

type human struct {
	name string
}

type wolf struct {
	freq int
}

type greeter interface {
	greeting(name string) string
}

func (h human) greeting(name string) string {
	return fmt.Sprintf("Hello, %s. I'm %s.", name, h.name)
}

func (w wolf) greeting(string) string {
	msg := strings.Repeat("woof ", w.freq)
	return fmt.Sprintf("%s!", strings.TrimSpace(msg))
}

func meet(name string, gs ...greeter) {
	for _, g := range gs {
		fmt.Println(g.greeting(name))
	}
}

func main() {
	a := human{"Alice"}
	b := wolf{6}
	username := "Dan"
	meet(username, a, b)
}

