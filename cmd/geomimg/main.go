package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/pravj/geopattern"
)

var patterns = []string{
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

var (
	pattern = flag.String("p", "", fmt.Sprintf("pattern to use: %s", patterns))
	phrase  = flag.String("s", "", "seed phrase")
	color   = flag.String("c", "", "color, like #3b5998")
	output  = flag.String("o", "", "output file, write to stdout if empty")
)

func randomString() string {
	var sb strings.Builder
	for i := 0; i < 32; i++ {
		sb.WriteRune(rune(97 + rand.Intn(26)))
	}
	log.Println(sb.String())
	return sb.String()
}

func randomColorString() string {
	return fmt.Sprintf("#%x%x%x",
		rand.Intn(255), rand.Intn(255), rand.Intn(255),
	)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	if *pattern == "" {
		*pattern = patterns[rand.Intn(len(patterns))]
	}
	if *phrase == "" {
		*phrase = randomString()
	}
	if *color == "" {
		*color = randomColorString()
	}
	args := map[string]string{
		"phrase":    *phrase,
		"generator": *pattern,
		"color":     *color,
	}
	gp := geopattern.Generate(args)
	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		fmt.Fprintf(f, gp)
	} else {
		fmt.Println(gp)
	}
}
