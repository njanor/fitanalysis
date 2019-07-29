package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/tormoder/fit"

	"time"
)

func main() {
	fmt.Println("Hello world")
	fitFile, _ := ioutil.ReadFile("2019-07-27.EnduranceRide.fit")
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
