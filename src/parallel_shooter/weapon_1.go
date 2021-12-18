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
	bullet int
	width float64
	interval time.Duration
	reload time.Duration
	shotImages *ImageSet
	status WeaponStatus
	lastShot time.Time
}

func NewWeapon_1(o Owner, a int, s float64, n int, b int, wi float64, i,r time.Duration, img *ImageSet) Weapon {
	w := &weapon_1{}
	w.owner = o
	w.attack = a
	w.speed = s
	w.number = n
	w.burst = b
	w.bullet = n * b
	w.width = wi
	w.interval = i
	w.reload = r
	w.shotImages = img
	w.status = WEAPON_READY
	w.lastShot = time.Time{}
	return w
}

func (w *weapon_1) shot() {
	switch w.status {
	case WEAPON_INTERVAL:
		if !w.isElapsed(w.interval)  {
			return
		}
		w.status = WEAPON_READY
	case WEAPON_RELOADING:
		if !w.isElapsed(w.reload)  {
			return
		}
		w.bullet = w.number * w.burst
		w.status = WEAPON_READY
	}
	switch w.status {
	case WEAPON_READY:
		game := w.owner.Game()
		point := w.owner.Center()
		phase := w.owner.Phase()
		direction := w.owner.Direction()
		shots := []*Shot{}
		for i := 0; i < w.number; i++ {
			o := Object{game: game, point: point, height: 5, width: 5, phase: phase, images: w.shotImages}
			s := newShot(o, w.attack, NewVector(math.Cos(direction) * w.speed, math.Sin(direction) * w.speed))
			shots = append(shots, s)
		}
		for i,shot := range shots {
			offsetPoint := point.offset(NewVector(10*float64(i),0))
			shot.setCenter(offsetPoint)
			enemies := game.getEnemies()
			for _, e := range enemies {
				shot.addEnemy(e)
			}
			game.setObject(shot)
		}
		w.lastShot = time.Now()
		w.bullet -= w.number
		if w.bullet > 0 {
			w.status = WEAPON_INTERVAL
		} else {
			w.status = WEAPON_RELOADING
		}
	}
}

func (w *weapon_1) isElapsed(t time.Duration) bool {
	return time.Now().Sub(w.lastShot.Add(t * time.Millisecond)) > 0
}
