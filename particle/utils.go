package particle

import "math"

type distance struct {
	x float64
	y float64
}

func GenerateArrayOf(p *Particle, quantity, winWidth, winHeight int) []*Particle {
	arr := make([]*Particle, 0)
	for i := 0; i < quantity; i++ {
		arr = append(arr, p.GetClone())
	}

	return arr
}

func GetDistanceBetween(p1, p2 *Particle) distance {
	return distance{
		math.Abs(float64(p1.Xpos - p2.Xpos)),
		math.Abs(float64(p1.Ypos - p2.Ypos)),
	}
}
