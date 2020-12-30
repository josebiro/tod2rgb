# tod2rgb

## Time of Day to RGB

Takes input of lattitude and longitude and converts the local time of day using sunrise/sunset and astronomical noon to determine the RGB value of daylight for things like automation of daylight cycles for aquariums or plants or just plain old LED llghts. Outputs via http client to a defined host usinng [WLED's http api](https://github.com/Aircoookie/WLED/wiki/HTTP-request-API) for setting brightness and color values.

The core kelvin to rgb algorithm is adapted from
https://gist.github.com/paulkaplan/5184275 which itseelf was derived from
http://www.tannerhelland.com/4435/convert-temperature-rgb-algorithm-code/

This application runs in a loop every minute and calculates an RGB value from a kelvin temperature inferred from the current time of day and sunrise/solarnoon/sunset times with kelvin temp fades from 2400K to 2700K for dawn to sunrise and from sunset to dusk. Night is pure blue.

## Dependencies
* github.com/bradfitz/latlong -- to get the timezone from lat/long
* github.com/sixdouglas/suncalc -- to calculate the time of day solar events.

## How To

### Run Locally:
* `make build` will create binaries in `./bin/<platform>/`
* create or edit `./env.sh` and add environment variables `LAT=<your latitiude>` and `LONG=<your longitude>` then `source env.sh`
* `./bin/<platform>/tod2rgb`
* alternatively, you can define lat/long on the command line:
* `./bin/<platform>/tod2rgb --lat=<your lattitude> --long=<your longitude>`

### Kubernetes (coming soon)
* TODO: Run on k8s using kustomize.

