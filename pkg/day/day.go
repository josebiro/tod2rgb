package day

import "time"

type Day struct {
	Dawn          time.Time
	Sunrise       time.Time
	SolarNoon     time.Time
	Dusk          time.Time
	Sunset        time.Time
	SolarMidnight time.Time
}

func (d *Day) SetDawn(t time.Time) {
	d.Dawn = t
	return
}

func (d *Day) SetSunrise(t time.Time) {
	d.Sunrise = t
}

func (d *Day) SetSolarNoon(t time.Time) {
	d.SolarNoon = t
}

func (d *Day) SetDusk(t time.Time) {
	d.Dusk = t
}

func (d *Day) SetSunset(t time.Time) {
	d.Sunset = t
}

func (d *Day) SetSolarMidnight(t time.Time) {
	d.SolarMidnight = t
}

func NewDay() *Day {
	d := Day{}
	return &d
}
