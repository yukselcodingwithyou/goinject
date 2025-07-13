package main

import (
	"fmt"
	"github.com/user/goinject"
)

type repo interface {
	Get() string
}

type repoImplA struct{}

func (r *repoImplA) Get() string { return "data from repo A" }

type repoImplB struct{}

func (r *repoImplB) Get() string { return "data from repo B" }

type service struct {
	Repo repo `autowire:"b"`
}

func (s *service) Serve() {
	fmt.Println("Service got:", s.Repo.Get())
}

func main() {
	c := goinject.New()
	c.ProvideNamed("a", &repoImplA{})
	c.ProvideNamed("b", &repoImplB{})

	s := &service{}
	if err := c.Fill(s); err != nil {
		panic(err)
	}

	s.Serve()
}
