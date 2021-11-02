package main

type Entity struct {
	id         int
	components []Component
}

type Component interface {
}
