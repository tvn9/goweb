package typeconv

import (
	"fmt"
	"log"
	"strconv"
)

func FloatToString(p float64) (s string) {
	if p < 0.0 {
		log.Fatalln("invalid value")
	} else {
		s = fmt.Sprintf("%.2f", p)
	}
	return s
}

func StringToFloat(s string) (f float64) {
	if len(s) <= 0 {
		log.Fatalln("invalid value")
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
