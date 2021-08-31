package geomimg

import (
	"fmt"
	"math/rand"
	"strings"
)

var AvailablePatterns = []string{
	"chevrons",
	"concentric-circles",
	"diamonds",
	"hexagons",
	"mosaic-squares",
	"nested-squares",
	"octagons",
	"overlapping-circles",
	"overlapping-rings",
	"plaid",
	"plus-signs",
	"sine-waves",
	"squares",
	"tessellation",
	"triangles",
	"xes",
}

func RandomPattern() string {
	return AvailablePatterns[rand.Intn(len(AvailablePatterns))]
}

func RandomString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteRune(rune(97 + rand.Intn(26)))
	}
	// log.Println(sb.String())
	return sb.String()
}

func RandomColorString() string {
	return fmt.Sprintf("#%x%x%x",
		rand.Intn(255), rand.Intn(255), rand.Intn(255),
	)
}
