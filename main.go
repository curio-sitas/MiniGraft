package main

import (
	CORE "MiniGraft/core"
)

func main() {

	engine := CORE.InitEngine(CORE.GAME_SETTINGS{Width: 600, Height: 400, Title: "MiniGraft"})
	engine.Setup()
	engine.Run()
	engine.Cleanup()
	engine.Exit()

	return
}
