package particle

type distance struct {
	x float64
	y float64
}

func GenerateArrayOf(p *Particle, quantity, winWidth, winHeight int) []*Particle {
	arr := make([]*Particle, 0)
	for i := 0; i < quantity; i++ {
		pCopy := *p
		pCopy.SetRandomPosition(winWidth, winHeight)
		arr = append(arr, &pCopy)
	}
	return arr
}

func GetDistanceBetween(p1, p2 *Particle) distance {
	return distance{
		float64(uint(p1.Xpos - p2.Xpos)),
		float64(uint(p1.Ypos - p2.Xpos)),
	}
}
