package parallel_shooter

import (
	"math"
	"time"
)

type weapon_1 struct {
	Weapon
	owner Owner
	attack int
	speed float64
	number int
	burst int
	width float64
	interval time.Duration
	reload time.Duration
	shotImages *ImageSet
	status WeaponStatus
}

func NewWeapon_1(o Owner, a int, s float64, n int, b int, wi float64, i,r time.Duration, img *ImageSet) Weapon {
	w := &weapon_1{}
	w.owner = o
	w.attack = a
	w.speed = s
	w.number = n
	w.burst = b
	w.width = wi
	w.interval = i
	w.reload = r
	w.shotImages = img
	w.status = WEAPON_READY
	return w
}

func (w *weapon_1) shot() {
	switch w.status {
	case WEAPON_READY:
		game := w.owner.Game()
		point := w.owner.Center()
		phase := w.owner.Phase()
		direction := w.owner.Direction()
		shots := []*Shot{}
		w.status = WEAPON_INTERVAL
		go func() {
			for b := 0; b < w.burst; b++ {
				for i := 0; i < w.number; i++ {
					o := Object{game: game, point: point, height: 5, width: 5, phase: phase, images: w.shotImages}
					s := newShot(o, w.attack, NewVector(math.Cos(direction) * w.speed, math.Sin(direction) * w.speed))
					shots = append(shots, s)
				}
				for _,shot := range shots {
					shot.setCenter(w.owner.Center())
					enemies := game.getEnemies()
					for _, e := range enemies {
						shot.addEnemy(e)
					}
					game.setObject(shot)
				}
				time.Sleep(w.interval * time.Millisecond)
			}
			w.status = WEAPON_RELOADING
			time.Sleep(w.reload * time.Millisecond)
			w.status = WEAPON_READY
		}()
	}
}
