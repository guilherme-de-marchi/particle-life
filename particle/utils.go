package particle

func GenerateArrayOf(p *Particle, quantity, winWidth, winHeight int) []*Particle {
	arr := make([]*Particle, 0)
	for i := 0; i < quantity; i++ {
		pCopy := *p
		pCopy.SetRandomPosition(winWidth, winHeight)
		arr = append(arr, &pCopy)
	}
	return arr
}
