package main

import (
	"bytes"
	"flag"
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

var sink string

func init() {
	flag.StringVar(&sink, "sink", "", "where to sink events to")
}

func main() {
	flag.Parse()
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 1 * time.Minute
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    sink,
		Body:   []byte("hello from Vegeta!"),
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	var b bytes.Buffer

	reporter := vegeta.NewTextReporter(&metrics)
	reporter.Report(&b)

	fmt.Println(b.String())
}
