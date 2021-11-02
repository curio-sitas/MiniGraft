package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type State interface {
	Render(screen *ebiten.Image)
	Update(dt float64)
	Init()
	Destroy()
}

type SState struct {
}

func (s *SState) Render(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 0, 15, 100})
}
func (s *SState) Update(dt float64) {}
func (s *SState) Init() {
	a := GameObject{}
	a.addComponent(nil)
	fmt.Println("Implementiation works")
}
func (s *SState) Destroy() {}

type StateStack struct {
	states [](*State)
}

func (stack *StateStack) InitStack(s State) {
	stack.states = []*State{}
	stack.PushState(s)
}

func (stack *StateStack) PushState(s State) {
	s.Init()
	stack.states = append(stack.states, &s)
}

func (stack *StateStack) PopState() {
	if len(stack.states) <= 1 {
		panic("State stack is too small to pop more")
	}
	state := *stack.states[len(stack.states)-1]
	state.Destroy()
	stack.states = stack.states[:len(stack.states)-1]
}

func (stack *StateStack) PeekState() *State {
	if len(stack.states) <= 0 {
		panic("State stack is empty")
	}
	return stack.states[len(stack.states)-1]
}

func (stack *StateStack) Cleanup() {
	for _, v := range stack.states {
		state := *v
		state.Destroy()
		stack.states = stack.states[:len(stack.states)-1]
	}
}
