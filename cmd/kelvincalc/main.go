package main

import (
	"encoding/json"
	"fmt"

	"github.com/josebiro/tod2rgb/pkg/kelvin"
	flag "github.com/spf13/pflag"
)

// Flags
var k float64

func init() {
	flag.Float64Var(&k, "kelvin", 2400, "Kelvin temp to convert to rgb")
}

func main() {
	//var err error
	flag.Parse()
	fmt.Println("Converting kelvin: ", k)

	c := kelvin.KelvinToRGB(k)
	//fmt.Println(c)
	PrettyPrint(c)

}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
