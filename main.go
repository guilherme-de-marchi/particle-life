package main

import (
	"fmt"

	"github.com/Guilherme-De-Marchi/particle-life/particle"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WIN_WIDTH    = 800
	WIN_HEIGHT   = 600
	MAX_ENTITIES = 300
)

var (
	proton   particle.Particle = *particle.NewParticle(10, [4]int{0, 255, 0, 255})
	electron particle.Particle = *particle.NewParticle(30, [4]int{255, 0, 0, 255})
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL: ", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Particle Life",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WIN_WIDTH, WIN_HEIGHT,
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		fmt.Println("Initializing WINDOW: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_ACCELERATED,
	)
	if err != nil {
		fmt.Println("Initializing RENDERER: ", err)
		return
	}
	defer renderer.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.Present()
	}
}
