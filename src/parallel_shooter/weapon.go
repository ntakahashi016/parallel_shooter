package parallel_shooter

type WeaponStatus int
const (
	WEAPON_READY WeaponStatus = iota
	WEAPON_INTERVAL
	WEAPON_RELOADING
)

type Weapon interface {
	shot()
}
