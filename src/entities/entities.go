package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	PiOv2 float32 = rl.Pi/2
	TwoPi float32 = rl.Pi * 2
)

var (
	SCREEN_W float32
	SCREEN_H float32

	Ent *Entities
)

type Entity interface { // Interface for body
  Draw()
  Update(float32)
}

type Projectile interface {
	Entity
	getTimeOfCreation() float32
	getTimeLimit() float32
}

type Entities struct {
	players []*Tank
	projec []Projectile			 // Projectiles
}

func (e *Entities) AddPlayer(pos rl.Vector2) {
	e.players = append(e.players, NewTank(len(e.players), pos))
}

func (e *Entities) DrawAllEntites() {
	for i := 0; i < len(e.players); i++ {
		e.players[i].Draw()
	}
	for i := 0; i < len(e.projec); i++ {
		e.projec[i].Draw()
	}
}

func (e *Entities) UpdateAllEntites(dt float32) {
	for i := 0; i < len(e.players); i++ {
		e.players[i].Update(dt)
	}
	for i := 0; i < len(e.projec); i++ {
		e.projec[i].Update(dt)
	}

	e.checkProjectileTimeouts()
}

func (e *Entities) checkProjectileTimeouts() {
	currTime := rl.GetTime()
	for i := 0; i < len(e.projec); i++ {
		if currTime - e.projec[i].getTimeOfCreation() > e.projec[i].getTimeLimit() {
			copy(e.projec[i:], e.projec[i+1:])
			e.projec[len(e.projec)-1] = nil
			e.projec = e.projec[:len(e.projec)-1]
		}
	}
}
