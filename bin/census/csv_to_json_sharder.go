package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInt(number string) int {
	i, err := strconv.Atoi(number)
	if err != nil {
		log.Fatalf("Int parsed failed: %v", err)
	}
	return i
}

func getFloat(number string) float64 {
	f, err := strconv.ParseFloat(number, 64)
	if err != nil {
		log.Fatalf("Float parse failed: %v", err)
	}
	return f
}

// www.census.gov/popest/data/counties/asrh/2007/files/CC-EST2007-alldata.txt
func ParseDemographic(startYear int, line string) Demographic {
	var demographic Demographic
	l := strings.Split(line, ",")
	demographic.State = l[3]
	demographic.County = l[4]
	demographic.Year = startYear + getInt(l[5])
	demographic.AgeGroup = getInt(l[6])
	demographic.FipsID = 1000*getInt(l[1]) + getInt(l[2])
	if l[8] == "X" {
		log.Printf("No population data for county %s", l[4])
		return demographic
	}
	demographic.TotalMen = getFloat(l[8])
	demographic.TotalWomen = getFloat(l[9])
	demographic.WhiteMen = getFloat(l[34])
	demographic.WhiteWomen = getFloat(l[35])
	demographic.BlackMen = getFloat(l[36])
	demographic.BlackWomen = getFloat(l[37])
	demographic.NativeMen = getFloat(l[38])
	demographic.NativeWomen = getFloat(l[39])
	demographic.AsianMen = getFloat(l[40])
	demographic.AsianWomen = getFloat(l[41])
	demographic.PacificMen = getFloat(l[42])
	demographic.PacificWomen = getFloat(l[43])
	demographic.MixedMen = getFloat(l[44])
	demographic.MixedWomen = getFloat(l[45])
	demographic.HispanicMen = getFloat(l[56])
	demographic.HispanicWomen = getFloat(l[57])
	return demographic
}

func createFile(d Demographic) {
	os.MkdirAll(fmt.Sprintf("data/census/%d/%d", d.Year, d.FipsID), 0755)
	filename := fmt.Sprintf("data/census/%d/%d/%d.json",
		d.Year, d.FipsID, d.AgeGroup)
	json, err := json.Marshal(d)
	if err != nil {
		log.Printf("uh oh %v", err)
	}
	if err = ioutil.WriteFile(filename, json, 0755); err != nil {
		log.Printf("sad! %v", err)
	}
}

func main() {
	stateStrings := []string{"01", "02", "04", "05", "06", "08", "09", "10", "11", "12", "13", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "44", "45", "46", "47", "48", "49", "50", "51", "53", "54", "55", "56"}
	startNum := 3
	startYear := 1997
	for _, i := range stateStrings {
		file, _ := os.Open(fmt.Sprintf(
			"data/census/rawdata/CC-EST2009-ALLDATA-%s.csv", i))
		log.Print(i)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		for scanner.Scan() {
			d := ParseDemographic(startYear, scanner.Text())
			if d.Year < startYear+startNum {
				continue
			}
			createFile(d)
		}
	}
}
