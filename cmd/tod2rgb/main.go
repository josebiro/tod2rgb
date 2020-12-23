package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bradfitz/latlong"
	"github.com/josebiro/tod2rgb/pkg/color"
	day "github.com/josebiro/tod2rgb/pkg/day"
	"github.com/josebiro/tod2rgb/pkg/kelvin"
	log "github.com/sirupsen/logrus"
	"github.com/sixdouglas/suncalc"
	flag "github.com/spf13/pflag"
)

// Flags
var debug bool
var lat float64
var long float64
var host string
var interval int

func init() {
	flag.BoolVar(&debug, "debug", false, "Turn on debug logging")
	flag.Float64Var(&lat, "lat", 1234, "Lattitude of target location")
	flag.Float64Var(&long, "long", 1234, "Longitude of target location.")
	flag.StringVar(&host, "host", "1.2.3.4", "WLED Controller IP addr or host name.")
	flag.IntVar(&interval, "interval", 0, "Update interval in minutes.")
}

func main() {
	// tod2rgb - take geo lat long as input and return the rgb of daylight color temp.
	/* 	The principle behind this comes fromm
	TODO: Link to blog (links in readme)
	*/

	var err error

	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug output is on.")
		log.Debug("Flags: ")
		log.Debug("--debug: ", debug)
		log.Debug("--lat: ", lat)
		log.Debug("--long: ", long)
		log.Debug("--interval: ", interval)
	}

	if lat == 1234 {
		lat, err = strconv.ParseFloat(os.Getenv("LAT"), 64)
		if err != nil {
			log.Fatal("ERROR: No Lattitude defined; ", err)
		}
	}
	if long == 1234 {
		long, err = strconv.ParseFloat(os.Getenv("LONG"), 64)
		if err != nil {
			log.Fatal("ERROR: No longitude defined; ", err)
		}
	}

	if debug {
		log.Debug("Lattitude: ", lat)
		log.Debug("Longitude: ", long)
	}

	if interval == 0 {
		c := GetKelvinColor(lat, long)
		url := fmt.Sprintf("http://%s/win&R=%v&G=%v&B=%v", host, c.Red, c.Green, c.Blue)
		log.Debug(url)
		err := UpdateWled(url)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Info("Running continually at ", interval, "minute interval.")
		for {
			c := GetKelvinColor(lat, long)
			url := fmt.Sprintf("http://%s/win&R=%v&G=%v&B=%v", host, c.Red, c.Green, c.Blue)
			log.Debug(url)
			err := UpdateWled(url)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Duration(interval) * time.Minute)
		}
	}
}

func UpdateWled(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		log.Debug(bodyString)
	}
	return nil
}

func GetKelvinColor(lat, long float64) *color.Color {
	d := day.NewDay()

	// convert lat long to timezone local time
	targetTime, err := LocaltimeFromLatLong(lat, long)
	if err != nil {
		log.Debug(err)
		log.Fatal(err)
	}
	if debug {
		log.Debug(targetTime)
	}

	// convert local time to RGB color value
	times := suncalc.GetTimes(targetTime, lat, long)

	d.SetDawn(times["dawn"].Time)
	d.SetSunrise(times["sunrise"].Time)
	d.SetSolarNoon(times["solarNoon"].Time)
	d.SetSunset(times["sunset"].Time)
	d.SetDusk(times["dusk"].Time)
	d.SetSolarMidnight(times["nadir"].Time)
	d.SetCurrent(targetTime)

	if debug {
		log.Debug(d)
		log.Debug("Daytime: ", d.IsDaytime(targetTime))
		log.Debug("Nighttime: ", d.IsNighttime(targetTime))
		log.Debug("Current Phase: ", d.Between())
	}

	/*
		 TODO: eventually add moon phases and brightness controls
		moon_illum := suncalc.GetMoonIllumination(target_time)
		moon_times := suncalc.GetMoonTimes(target_time, home_lat, home_long, false)
	*/

	currentKelvin := d.CurrentKelvin()
	c := kelvin.KelvinToRGB(currentKelvin)
	if debug {
		log.Debug("Current Kelvin: ", currentKelvin)
	}
	log.Debug(c)
	return c
}

func LocaltimeFromLatLong(lat, long float64) (time.Time, error) {
	tz := latlong.LookupZoneName(lat, long)
	time_utc := time.Now().UTC()

	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Debug("time_utc: ", time_utc, " err: ", err)
		return time_utc, err
	}
	target_time := time_utc.In(location)
	if debug {
		log.Debug("TZ: ", location, " Time: ", target_time)
	}
	return target_time, nil
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		log.Error(string(b))
	}
	return
}
