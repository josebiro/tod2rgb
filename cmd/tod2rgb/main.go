package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bradfitz/latlong"
	day "github.com/josebiro/tod2rgb/pkg/day"
	"github.com/josebiro/tod2rgb/pkg/kelvin"
	"github.com/sixdouglas/suncalc"
	flag "github.com/spf13/pflag"
)

// Flags
var debug bool
var lat float64
var long float64

func init() {
	flag.BoolVar(&debug, "debug", false, "Turn on debug logging")
	flag.Float64Var(&lat, "lat", 1234, "Lattitude of target location")
	flag.Float64Var(&long, "long", 1234, "Longitude of target location.")
}

func main() {
	// tod2rgb - take geo lat long as input and return the rgb of daylight color temp.
	/* 	The principle behind this comes fromm
	TODO: Link to blog (links in readme)
	*/

	var err error

	flag.Parse()

	if debug {
		fmt.Println("*** Debug output is on.")
		fmt.Println("*** Flags: ")
		fmt.Println("***    debug: ", debug)
		fmt.Println("***    lat: ", lat)
		fmt.Println("***    long: ", long)
	}

	if lat == 1234 {
		lat, err = strconv.ParseFloat(os.Getenv("LAT"), 64)
		if err != nil {
			fmt.Println("ERROR: No Lattitude defined; ", err)
		}
	}
	if long == 1234 {
		long, err = strconv.ParseFloat(os.Getenv("LONG"), 64)
		if err != nil {
			fmt.Println("ERROR: No longitude defined; ", err)
		}
	}

	if debug {
		fmt.Println("*** Lattitude: ", lat)
		fmt.Println("*** Longitude: ", long)
	}

	d := day.NewDay()

	// convert lat long to timezone local time
	targetTime, err := LocaltimeFromLatLong(lat, long)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(targetTime)

	// convert local time to RGB color value
	times := suncalc.GetTimes(targetTime, lat, long)
	//PrettyPrint(times)

	d.SetDawn(times["dawn"].Time)
	d.SetSunrise(times["sunrise"].Time)
	d.SetSolarNoon(times["solarNoon"].Time)
	d.SetSunset(times["sunset"].Time)
	d.SetDusk(times["dusk"].Time)
	d.SetSolarMidnight(times["nadir"].Time)
	d.SetCurrent(targetTime)

	if debug {
		PrettyPrint(d)
	}

	fmt.Println("Daytime: ", d.IsDaytime(targetTime))
	fmt.Println("Nighttime: ", d.IsNighttime(targetTime))

	fmt.Println("Current Phase: ", d.Between())

	//moon_illum := suncalc.GetMoonIllumination(target_time)
	//PrettyPrint(moon_illum)

	//moon_times := suncalc.GetMoonTimes(target_time, home_lat, home_long, false)
	//PrettyPrint(moon_times)

	// Map time of day to kelvin temp
	currentKelvin := d.CurrentKelvin()
	fmt.Println("Current Kelvin: ", currentKelvin)
	c := kelvin.KelvinToRGB(currentKelvin)
	PrettyPrint(c)

	// return rgb color values
}

func LocaltimeFromLatLong(lat, long float64) (time.Time, error) {
	tz := latlong.LookupZoneName(lat, long)
	time_utc := time.Now().UTC()

	location, err := time.LoadLocation(tz)
	if err != nil {
		return time_utc, err
	}
	target_time := time_utc.In(location)
	fmt.Println("TZ: ", location, " Time: ", target_time)
	return target_time, nil
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
