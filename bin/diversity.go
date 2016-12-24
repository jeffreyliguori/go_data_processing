package main

import (
	"census"
	"fmt"
)

func howWhite(d census.Demographic) float64 {
	return (d.WhiteMen + d.WhiteWomen) / (d.TotalMen + d.TotalWomen)
}

func main() {
	directory := "/home/jeffrey/data/census"
	d2000, err := census.GetPerCountyDataFor(2000, 0, directory)
	if err != nil {
		fmt.Printf("200: %v", err)
		return
	}
	d2015, err := census.GetPerCountyDataFor(2015, 0, directory)
	if err != nil {
		fmt.Printf("2015: %v", err)
		return
	}
	dems2000 := make(map[int]float64)
	for _, d := range d2000 {
		dems2000[d.FipsID] = howWhite(d)
	}
	var moreDiverse, lessDiverse int
	for _, d := range d2015 {
		white := howWhite(d)
		if white < dems2000[d.FipsID] {
			moreDiverse++
		} else {
			fmt.Printf("%s, %s: 2000: %f, 2015: %f\n", d.County, d.State, dems2000[d.FipsID], howWhite(d))
			lessDiverse++
		}
	}
	fmt.Printf("%d %d", moreDiverse, lessDiverse)
	leastDiverse := 0.0
	place := ""
	for _, d := range d2015 {
		if howWhite(d) > leastDiverse {
			place = d.County + "," + d.State
			leastDiverse = howWhite(d)
		}
	}
	fmt.Printf(place+": %f", leastDiverse)
}
