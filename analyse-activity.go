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

type analysisOption struct {
	Name        string
	Description string
}

var stats = analysisOption{
	Name:        "stats",
	Description: "Display duration, laps with length, and total length",
}

var avgWatts = analysisOption{
	Name:        "avgWatts",
	Description: "Calculate average wattage for entire workout as well as for individual laps",
}

func main() {
	fileName := flag.String("file", "", "The FIT file containing the activity to analyse")
	analysis := flag.String("analysis", stats.Name, fmt.Sprintf("The analysis to perform. Options are:\n\t%s - %s\n\t%s - %s", stats.Name, stats.Description, avgWatts.Name, avgWatts.Description))
	flag.Parse()

	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	fitFile, _ := ioutil.ReadFile(*fileName)
	fit, _ := fit.Decode(bytes.NewReader(fitFile))
	activity, _ := fit.Activity()

	switch *analysis {
	case stats.Name:
		analyseStats(activity)
	case avgWatts.Name:
		analyseAvgWatts(activity)
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

func analyseAvgWatts(activity *fit.ActivityFile) {
	averageWattsForInterval(activity.Records)
	var totalAveragePower uint
	for lapNumber, lap := range activity.Laps {
		fmt.Println(lap.TotalElapsedTime)
		totalAveragePower += uint(lap.AvgPower) * uint(lap.TotalElapsedTime/1000)
		fmt.Println(lapNumber)
		fmt.Println(lap.AvgPower)
	}
	fmt.Println(totalAveragePower / uint(activity.Activity.GetTotalTimerTimeScaled()))
}

func averageWattsForInterval(records []*fit.RecordMsg) {
	duration := records[len(records)-1].Timestamp.Sub(records[0].Timestamp)
	numberOfSeconds := uint(duration.Seconds())

	allWattagesPerSecond := make([]uint16, numberOfSeconds)
	var totalWatts uint

	for i := 0; i < len(records)-1; i++ {
		firstRecord := records[i]
		secondRecord := records[i+1]

		durationCalculated := secondRecord.Timestamp.Sub(firstRecord.Timestamp)
		numberOfSecondsCalculated := int(durationCalculated.Seconds())
		averageWatts := (firstRecord.Power + secondRecord.Power) / 2
		for j := 0; j < numberOfSecondsCalculated; j++ {
			allWattagesPerSecond[i+j] = averageWatts
			totalWatts += uint(averageWatts)
		}
	}

	fmt.Println(totalWatts / numberOfSeconds)

	if numberOfSeconds >= 1200 {
		averageWattageOver20Minutes := make([]uint16, numberOfSeconds-1200+1)
		var highestAverageWattage uint16
		for i := uint(0); i < numberOfSeconds-uint(1200); i++ {
			averageWattageOver20Minutes[i] = sumWattagesOverInterval(allWattagesPerSecond[i : int(i)+1200])
			if averageWattageOver20Minutes[i] > highestAverageWattage {
				highestAverageWattage = averageWattageOver20Minutes[i]
			}
		}

		fmt.Println(highestAverageWattage)
	}
}

func sumWattagesOverInterval(secondWattages []uint16) uint16 {
	var totalWattage uint
	for _, singleSecondAverageWattage := range secondWattages {
		totalWattage += uint(singleSecondAverageWattage)
	}

	return uint16(totalWattage / uint(len(secondWattages)))
}
