package goinject

import "testing"

type repo interface {
	Get() string
}

type repoImpl struct{}

func (r *repoImpl) Get() string { return "repo" }

type service struct {
	R repo `autowire:""`
}

func TestFill(t *testing.T) {
	c := New()
	c.Provide(&repoImpl{})

	s := &service{}
	if err := c.Fill(s); err != nil {
		t.Fatalf("Fill error: %v", err)
	}
	if s.R == nil {
		t.Fatalf("expected repo to be injected")
	}
	if s.R.Get() != "repo" {
		t.Fatalf("wrong repo implementation")
	}
}

type repoImplA struct{}

func (r *repoImplA) Get() string { return "A" }

type repoImplB struct{}

func (r *repoImplB) Get() string { return "B" }

type serviceQual struct {
	R repo `autowire:"b"`
}

func TestFillWithQualifier(t *testing.T) {
	c := New()
	c.ProvideNamed("a", &repoImplA{})
	c.ProvideNamed("b", &repoImplB{})

	s := &serviceQual{}
	if err := c.Fill(s); err != nil {
		t.Fatalf("Fill error: %v", err)
	}
	if s.R == nil {
		t.Fatalf("expected repo to be injected")
	}
	if s.R.Get() != "B" {
		t.Fatalf("wrong repo implementation: %s", s.R.Get())
	}
}
