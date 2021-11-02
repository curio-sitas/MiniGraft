package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	_title     string
	_size      vector2
	statestack StateStack
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	(*g.statestack.PeekState()).Update(0.0)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	(*g.statestack.PeekState()).Render(screen)
	//screen.Clear()
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Init(title string, size vector2) {
	g.statestack = StateStack{}
	g.statestack.InitStack(&SState{}) // initial state for testing purpose
	g._size = size
	g._title = title
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(int(g._size.x), int(g._size.y))
	ebiten.SetWindowTitle(g._title)
	ebiten.SetRunnableOnUnfocused(true)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetWindowDecorated(true)
	// Call ebiten.RunGame to start your game loop.

}

func (g *Game) run() {
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
