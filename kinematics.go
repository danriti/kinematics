package kinematics

import (
	"fmt"
	"math"
)

const (
	Gravity     = float64(-9.8)
	GravityHalf = float64(Gravity / 2)
	DegToRad    = math.Pi / 180
)

// Return the displacement of an object at a given moment in time.
// d = (vi * t) + (1/2 * a * t^2)
func Displacement(vi float64, t float64, a float64) float64 {
	x := float64(0.5 * a)
	return float64((vi * t) + (x * math.Pow(t, 2.0)))
}

// Return degrees converted to radians.
func DegreeToRadian(d float64) float64 {
	return d * DegToRad
}

// Return the time required for the object to reach the ground from rest.
// d = (vi * t) + (1/2 * a * t^2)
func FreeFallTime(displacement float64) float64 {
	return math.Sqrt(displacement / GravityHalf)
}

// Simulate a free falling object until it reaches the ground from rest.
func FreeFall(displacement float64) float64 {
	time := float64(0)
	current := float64(0)
	for math.Abs(current) < math.Abs(displacement) {
		current = Displacement(0.0, time, Gravity) // current displacement
		fmt.Printf("time %0.2f, displacement %0.3f\n", time, current)
		time = time + 0.1
	}
	return time
}

// Simulate the physics of projectile motion.
type projectile struct {
	angle float64
	vi    float64
	vx    float64
	vy    float64
}

// Factory function for a projectile.
func Projectile(velocity, angle float64) *projectile {
	p := &projectile{
		vi:    velocity,
		angle: angle,
	}
	p.SetVelocityComponents()
	return p
}

// Set the x and y velocity components.
func (p *projectile) SetVelocityComponents() {
	radian := DegreeToRadian(p.angle)
	p.vx = p.vi * math.Cos(radian)
	p.vy = p.vi * math.Sin(radian)
}

// Get the total flight time of the projectile.
// t = (vfy - viy) / ay
func (p *projectile) GetFlightTime() float64 {
	viy := p.vy
	vfy := -viy
	return float64((vfy - viy) / Gravity)
}

// Launch the projectile and print its flight path until it reaches the ground.
func (p *projectile) Launch() {
	step := float64(0.1)
	time := float64(0)
	vcx := float64(0)
	vcy := float64(0)
	inFlight := true
	for inFlight {
		vcx = Displacement(p.vx, time, 0.0)     // current x displacement
		vcy = Displacement(p.vy, time, Gravity) // current y displacement
		fmt.Printf("time %0.2f, vx %2.3f, vy %2.3f\n", time, vcx, vcy)
		time = time + step
		if vcy < 0 {
			inFlight = false
		}
	}
}
