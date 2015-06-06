package kinematics

import (
	"math"
	"testing"
)

func AssertFloat64Equal(actual, expected float64) bool {
	return math.Abs(actual-expected) > 1
}

func TestFreeFallTime(t *testing.T) {
	expected := 1.32
	actual := FreeFallTime(float64(-8.52))
	if AssertFloat64Equal(actual, expected) {
		t.Errorf("actual %v", actual)
	}
}

func TestDisplacement(t *testing.T) {
	expected := 0.0
	actual := Displacement(0.0, 0.0, Gravity)
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	expected = -4.9
	actual = Displacement(0.0, 1.0, Gravity)
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	expected = -19.6
	actual = Displacement(0.0, 2.0, Gravity)
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestFreeFall(t *testing.T) {
	expected := 4.0
	actual := FreeFall(78.4)
	if AssertFloat64Equal(actual, expected) {
		t.Errorf("actual %v", actual)
	}
}

func TestProjectile(t *testing.T) {
	expected := 17.7
	actual := Projectile(25.0, 45.0)
	if AssertFloat64Equal(actual.vx, expected) {
		t.Errorf("actual %v", actual.vx)
	}
	if AssertFloat64Equal(actual.vy, expected) {
		t.Errorf("actual %v", actual.vx)
	}

	expected = 3.61
	if AssertFloat64Equal(actual.GetFlightTime(), expected) {
		t.Errorf("actual %v", actual.vx)
	}

	actual.Launch()
}
