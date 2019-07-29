package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/tormoder/fit"
)

func main() {
	fileName := flag.String("file", "", "The FIT file to analyze")
	flag.Parse()

	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	fitFile, _ := ioutil.ReadFile(*fileName)
	fit, _ := fit.Decode(bytes.NewReader(fitFile))
	activity, _ := fit.Activity()

	duration, _ := time.ParseDuration(fmt.Sprintf("%dms", activity.Activity.TotalTimerTime))

	fmt.Println(duration)
	fmt.Println(activity.Activity.GetTotalTimerTimeScaled())
	fmt.Println(len(activity.Laps))

	totalDistance := 0.

	for lapNumber, lap := range activity.Laps {
		fmt.Println(lapNumber)
		lapDistance := lap.GetTotalDistanceScaled()
		fmt.Println(lapDistance)
		totalDistance += lapDistance
	}

	fmt.Println(totalDistance)
}
