package day

import (
	"fmt"
	"time"

	"github.com/josebiro/tod2rgb/pkg/kelvin"
	log "github.com/sirupsen/logrus"
)

// Day - struct that carries times for day events
type Day struct {
	Dawn          time.Time
	Sunrise       time.Time
	SolarNoon     time.Time
	Sunset        time.Time
	Dusk          time.Time
	SolarMidnight time.Time
	Current       time.Time // pointer to current local time
}

func (d *Day) IsDaytime(t time.Time) bool {
	// if current time is between sunrise and sunset, it's daytime
	if t.After(d.GetSunrise()) && t.Before(d.GetSunset()) {
		return true
	}
	return false
}

func (d *Day) IsNighttime(t time.Time) bool {
	if t.After(d.GetSunset()) || t.Before(d.GetSunrise()) {
		return true
	}
	return false
}

func (d *Day) GetSunrise() time.Time {
	return d.Sunrise
}

func (d *Day) GetSunset() time.Time {
	return d.Sunset
}

func (d *Day) GetDusk() time.Time {
	return d.Dusk
}

func (d *Day) GetDawn() time.Time {
	return d.Dawn
}

func (d *Day) GetSolarMidnight() time.Time {
	return d.SolarMidnight
}
func (d *Day) GetNextSolarMidnight() time.Time {
	return d.SolarMidnight.Add(time.Hour * 24)
}

func (d *Day) GetSolarNoon() time.Time {
	return d.SolarNoon
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

func (d *Day) SetCurrent(t time.Time) {
	d.Current = t
}

func (d *Day) BetweenSolarMidnightAndDawn() bool {
	if d.Current.After(d.SolarMidnight) && d.Current.Before(d.Dawn) {
		return true
	}
	return false
}

func (d *Day) BetweenDawnAndSunrise() bool {
	if d.Current.After(d.Dawn) && d.Current.Before(d.Sunrise) {
		return true
	}
	return false
}

func (d *Day) BetweenSunriseAndSolarNoon() bool {
	if d.Current.After(d.Sunrise) && d.Current.Before(d.SolarNoon) {
		return true
	}
	return false
}

func (d *Day) BetweenSolarNoonAndSunset() bool {
	if d.Current.After(d.SolarNoon) && d.Current.Before(d.Sunset) {
		return true
	}
	return false
}

func (d *Day) BetweenSunsetAndDusk() bool {
	if d.Current.After(d.Sunset) && d.Current.Before(d.Dusk) {
		return true
	}
	return false
}

func (d *Day) BetweenDuskAndSolarMidnight() bool {
	if d.Current.After(d.Dusk) && d.Current.Before(d.GetNextSolarMidnight()) {
		return true
	}
	return false
}

func (d *Day) Between() string {
	if d.BetweenSolarMidnightAndDawn() {
		return "midnight-to-dawn"
	} else if d.BetweenDawnAndSunrise() {
		return "dawn-to-sunrise"
	} else if d.BetweenSunriseAndSolarNoon() {
		return "sunrise-to-noon"
	} else if d.BetweenSolarNoonAndSunset() {
		return "noon-to-sunset"
	} else if d.BetweenSunsetAndDusk() {
		return "sunset-to-dusk"
	} else if d.BetweenDuskAndSolarMidnight() {
		return "dusk-to-midnight"
	}
	return "unknown"
}

func (d *Day) CurrentKelvin() float64 {
	var kStart float64
	var kEnd float64

	switch cp := d.Between(); cp {
	case "midnight-to-dawn":
		return kelvin.Night
	case "dusk-to-midnight":
		return kelvin.Night
	case "dawn-to-sunrise":
		// dawn to sunrise == 2400k to 2700k
		kStart = 2400
		kEnd = 2700
		phaseStart := d.GetDawn()
		phaseEnd := d.GetSunrise()
		return d.GetKelvin(phaseStart, phaseEnd, kStart, kEnd)
	case "sunrise-to-noon":
		kStart = 2700
		kEnd = 6500
		phaseStart := d.GetSunrise()
		phaseEnd := d.GetSolarNoon()
		return d.GetKelvin(phaseStart, phaseEnd, kStart, kEnd)
	case "noon-to-sunset":
		kStart = 6500
		kEnd = 2700
		phaseStart := d.GetSunrise()
		phaseEnd := d.GetSolarNoon()
		return d.GetKelvin(phaseStart, phaseEnd, kStart, kEnd)
	case "sunset-to-dusk":
		kStart = 2700
		kEnd = 2400
		phaseStart := d.GetSunrise()
		phaseEnd := d.GetSolarNoon()
		return d.GetKelvin(phaseStart, phaseEnd, kStart, kEnd)
	default:
		log.Error("ERROR: should not have reached here. cp=", cp)
		return 0
	} // Daylight phases return kelvin temp for time of day
	return 0
}

func (d *Day) GetKelvin(st time.Time, et time.Time, sk float64, ek float64) float64 {
	diff := et.Sub(st)
	rDiff := diff.Round(time.Minute)
	fmt.Println("Phase Duration: ", diff)
	fmt.Println("Phase Duration in Minutes: ", rDiff.Minutes())
	kDiff := ek - sk
	fmt.Println("Phase Kelvin Diff: ", kDiff)
	cDiff := d.Current.Sub(st).Round(time.Minute)
	fmt.Println("Minutes since phase start: ", cDiff.Minutes())
	c2Diff := et.Sub(d.Current).Round(time.Minute)
	fmt.Println("Minutes until phase end: ", c2Diff.Minutes())
	percentPhaseComplete := cDiff.Minutes() / rDiff.Minutes()
	fmt.Println("Phase Complete (percent): ", percentPhaseComplete)
	kelvinPercent := float64(kDiff) * percentPhaseComplete
	fmt.Println("Kelvin Value: ", int(float64(sk)+kelvinPercent))
	return 0
}

func NewDay() *Day {
	d := Day{}
	return &d
}
