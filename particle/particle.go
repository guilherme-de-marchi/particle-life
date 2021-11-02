package particle

import (
	"math"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	RECT_WIDTH  = 10
	RECT_HEIGHT = 10
)

type Particle struct {
	Energy     int
	Xpos, Ypos int
	Rect       sdl.Rect
	Color      [4]int
}

func NewParticle(energy int, color [4]int) *Particle {
	return &Particle{
		Energy: energy,
		Color:  color,
	}
}

func (p *Particle) SetPosition(x, y int) {
	p.Xpos = x
	p.Ypos = y

	p.Rect.X = int32(x) - RECT_WIDTH/2
	p.Rect.Y = int32(y) - RECT_HEIGHT/2
}

func (p *Particle) SetRandomPosition(winWidth, winHeight int) {
	rand.Seed(time.Now().UnixNano())
	p.SetPosition(rand.Intn(winWidth), rand.Intn(winHeight))
}

func (p *Particle) CalcEnergyFieldIn(pTarget *Particle) float64 {
	dist := GetDistanceBetween(p, pTarget)
	radius := math.Sqrt(math.Pow(dist.x, 2) + math.Pow(dist.y, 2)) // Hipotenuse calc

	return float64(p.Energy) / (2*math.Pi*radius + 1) // resultantEnergyField = energy / (2*PI*r+1)
}

func (p *Particle) CalcResultantMovimentByAxle(pAxlePos, pTargetAxlePos int, energyField float64) int {
	if pTargetAxlePos >= pAxlePos {
		return pTargetAxlePos - int(energyField)
	} else {
		return pTargetAxlePos + int(energyField)
	}
}

func (p *Particle) InteractWith(pTarget *Particle) {
	energyField := p.CalcEnergyFieldIn(pTarget)
	pTarget.SetPosition(
		p.CalcResultantMovimentByAxle(p.Xpos, pTarget.Xpos, energyField),
		p.CalcResultantMovimentByAxle(p.Ypos, pTarget.Ypos, energyField),
	)
}
