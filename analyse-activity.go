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

const stats = "stats"

func main() {
	fileName := flag.String("file", "", "The FIT file containing the activity to analyse")
	analysis := flag.String("analysis", stats, fmt.Sprintf("The analysis to perform. Options are: %s", stats))
	flag.Parse()

	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	fitFile, _ := ioutil.ReadFile(*fileName)
	fit, _ := fit.Decode(bytes.NewReader(fitFile))
	activity, _ := fit.Activity()

	switch *analysis {
	case stats:
		analyseStats(activity)
	}

}

func analyseStats(activity *fit.ActivityFile) {
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
