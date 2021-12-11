package parallel_shooter

type weapon_1 struct {
	Weapon
	owner *Common
	attack int
	speed float64
	shotImages *ImageSet
}

type NewWeapon_1(o *Common, a int, s float64, i *ImageSet) {
	w := &weapon_1{}
	w.owner = o
	w.attack = a
	w.speed = s
	w.shotImages = i
	return w
}

func (w *weapon_1) shot() {
	game := w.owner.game
	point := w.owner.point
	phase := w.owner.phase
	direction := w.owner.direction
	o1 := Object{game: game, point: point, height: 5, width: 5, phase: phase, images: w.shotImages}
	s1 := newShot(o1, 1, NewVector(math.Cos(direction) * w.speed, math.Sin(direction) * w.speed))
	o2 := Object{game: game, point: point, height: 5, width: 5, phase: phase, images: w.shotImages}
	s2 := newShot(o2, 1, NewVector(math.Cos(direction - 1.0/12.0 * math.Pi) * w.speed, math.Sin(direction - 1.0/12.0 * math.Pi) * w.speed))
	o3 := Object{game: game, point: point, height: 5, width: 5, phase: phase, images: w.shotImages}
	s3 := newShot(o3, 1, NewVector(math.Cos(direction + 1.0/12.0 * math.Pi) * w.speed, math.Sin(direction + 1.0/12.0 * math.Pi) * w.speed))
	shots := []*Shot{s1,s2,s3}
	for _,shot := range shots {
		shot.setCenter(w.owner.Center())
		enemies := game.getEnemies()
		for _, e := range enemies {
			shot.addEnemy(e)
		}
		game.setObject(shot)
	}
}
