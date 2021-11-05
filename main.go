package main

import (
	"fmt"

	"github.com/Guilherme-De-Marchi/particle-life/particle"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WIN_WIDTH    = 800
	WIN_HEIGHT   = 500
	MAX_ENTITIES = 300
)

var (
	proton   *particle.Particle = particle.NewParticle(500, [4]uint8{0, 255, 0, 255})
	electron *particle.Particle = particle.NewParticle(-1200, [4]uint8{255, 0, 0, 255})
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

	var environment []*particle.Particle
	environment = append(environment, particle.GenerateArrayOf(proton, 80, WIN_WIDTH, WIN_HEIGHT)...)
	environment = append(environment, particle.GenerateArrayOf(electron, 400, WIN_WIDTH, WIN_HEIGHT)...)

	for _, p := range environment {
		p.SetRandomPosition(WIN_WIDTH, WIN_HEIGHT)
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, p := range environment {
			for _, subp := range environment {
				p.InteractWith(subp, WIN_WIDTH, WIN_HEIGHT)
			}
			renderer.SetDrawColor(
				p.Color[0],
				p.Color[1],
				p.Color[2],
				p.Color[3],
			)
			renderer.DrawRect(p.Rect)
			renderer.FillRect(p.Rect)
		}

		renderer.Present()
	}
}
