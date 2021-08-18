package core

import (

	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type GAME_SETTINGS struct {
	Width  int
	Height int
	Title  string
}

type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			screen.Set(i, j, color.RGBA{255, 200, 15, 255})
		}
	}

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

type GameEngine struct {

	_settings GAME_SETTINGS
}

func (e GameEngine) Cleanup() {

}

func (e GameEngine) Run() {

	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(e._settings.Width, e._settings.Height)
	ebiten.SetWindowTitle(e._settings.Title)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}

func (e GameEngine) Setup() {

}

func (e GameEngine) Exit() {
	return
}

func InitEngine(setting GAME_SETTINGS) *GameEngine {
	return &GameEngine{_settings: setting}
}
