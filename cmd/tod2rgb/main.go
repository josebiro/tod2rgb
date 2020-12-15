package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	day "github.com/josebiro/tod2rgb/pkg/day"

	"github.com/bradfitz/latlong"
	"github.com/sixdouglas/suncalc"
)

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

func main() {
	// tod2rgb - take geo lat long as input and return the rgb of daylight color temp.
	/* 	The principle behind this comes fromm
	TODO: Link to blog
	*/

	d := day.NewDay()

	//input lat long
	lat, err := strconv.ParseFloat(os.Getenv("LAT"), 64)
	long, err := strconv.ParseFloat(os.Getenv("LONG"), 64)

	// convert lat long to timezone local time
	target_time, err := LocaltimeFromLatLong(lat, long)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(target_time)

	// convert local time to RGB color value
	times := suncalc.GetTimes(target_time, lat, long)
	//PrettyPrint(times)

	d.SetDawn(times["dawn"].Time)
	d.SetSunrise(times["sunrise"].Time)
	d.SetSolarNoon(times["solarNoon"].Time)
	d.SetSunset(times["sunset"].Time)
	d.SetDusk(times["dusk"].Time)
	d.SetSolarMidnight(times["nadir"].Time)
	d.SetCurrent(target_time)

	PrettyPrint(d)

	fmt.Println("Daytime: ", d.IsDaytime(target_time))
	fmt.Println("Nighttime: ", d.IsNighttime(target_time))

	fmt.Println("Current Phase: ", d.Between())

	//moon_illum := suncalc.GetMoonIllumination(target_time)
	//PrettyPrint(moon_illum)

	//moon_times := suncalc.GetMoonTimes(target_time, home_lat, home_long, false)
	//PrettyPrint(moon_times)

	// Map time of day to kelvin temp
	fmt.Println("Current Kelvin: ", d.Kelvin())

	// return rgb color values
}
