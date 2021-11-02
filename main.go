package main

func main() {

	engine := Game{}
	engine.Init("Untitled", vector2{800, 400})
	engine.run()
	engine.statestack.Cleanup()

}
