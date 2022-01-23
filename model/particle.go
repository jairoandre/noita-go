package model

import "noita-go/utils"

type Particle struct {
	Weight float64
	Vel    utils.Vec2d
	Acc    utils.Vec2d
	Pos    utils.Point
}

func NewParticle(w float64, pos utils.Point) Particle {
	return Particle{
		Weight: w,
		Pos:    pos,
	}
}

func (p *Particle) Update() {
	p.Vel = p.Vel.Add(p.Acc)
	p.Pos = p.Pos.Add(p.Vel.ToPt())
	p.Acc = utils.ZeroVec2d()
}

func (p *Particle) ApplyForce(force utils.Vec2d) {
	acc := force.Div(p.Weight)
	p.Acc = p.Acc.Add(acc)
}
