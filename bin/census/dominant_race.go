package main

import (
	"encoding/json"
)

type Race struct {
	White    int `json:"White"`
	Black    int `json:"Black"`
	Native   int `json:"Native"`
	Asian    int `json:"Asian"`
	Pacific  int `json:"Pacific"`
	Hispanic int `json:"Hispanic"`
}

func maj(d census.Demographic) Race {
	r := Race{White: 1}
	num := d.WhiteMen + d.WhiteWomen
	if num < d.BlackMen+d.BlackWomen {
		r = Race{Black: 1}
		num = d.BlackMen + d.BlackWomen
	}
	if num < d.NativeMen+d.NativeWomen {
		r = Race{Native: 1}
		num = d.NativeMen + d.NativeWomen
	}
	if num < d.AsianMen+d.AsianWomen {
		r = Race{Asian: 1}
		num = d.AsianMen + d.AsianWomen
	}
	if num < d.PacificMen+d.PacificWomen {
		r = Race{Pacific: 1}
		num = d.PacificMen + d.PacificWomen
	}
	if num < d.HispanicMen+d.HispanicWomen {
		r = Race{Hispanic: 1}
		num = d.HispanicMen + d.HispanicWomen
	}
	return r
}

func majority(d census.Demographic) string {
	ans := "White"
	num := d.WhiteMen + d.WhiteWomen
	if num < d.BlackMen+d.BlackWomen {
		ans = "Black"
		num = d.BlackMen + d.BlackWomen
	}
	if num < d.NativeMen+d.NativeWomen {
		ans = "Native"
		num = d.NativeMen + d.NativeWomen
	}
	if num < d.AsianMen+d.AsianWomen {
		ans = "Asian"
		num = d.AsianMen + d.AsianWomen
	}
	if num < d.PacificMen+d.PacificWomen {
		ans = "Pacific"
		num = d.PacificMen + d.PacificWomen
	}
	if num < d.HispanicMen+d.HispanicWomen {
		ans = "Hispanic"
		num = d.HispanicMen + d.HispanicWomen
	}
	return ans
}

func howWhite(d census.Demographic) float64 {
	return (d.WhiteMen + d.WhiteWomen) / (d.TotalMen + d.TotalWomen)
}

func main() {
	directory := "/home/jeffrey/data/census"
	d2015, err := census.GetPerCountyDataFor(2015, 0, directory)
	if err != nil {
		fmt.Printf("2015: %v", err)
		return
	}
	/*
		dems2000 := make(map[int]float64)
		for _, d := range d2000 {
			dems2000[d.FipsID] = howWhite(d)
		}
		var moreDiverse, lessDiverse int
		for _, d := range d2015 {
			White := howWhite(d)
			if White < dems2000[d.FipsID] {
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
	*/
	for _, d := range d2015 {
		fmt.Printf(d.County+","+d.State+","+"%v"+"\n", maj(d))
	}
}
