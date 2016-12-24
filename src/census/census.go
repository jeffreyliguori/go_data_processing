package census

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
)

type Demographic struct {
	State         string  `json:"state"`
	County        string  `json:"county"`
	Year          int     `json:"year"`
	FipsID        int     `json:"fips_id"`
	AgeGroup      int     `json:"age_group"`
	TotalMen      float64 `json:"total_men"`
	TotalWomen    float64 `json:"total_women"`
	WhiteMen      float64 `json:"white_men"`
	WhiteWomen    float64 `json:"white_women"`
	BlackMen      float64 `json:"black_men"`
	BlackWomen    float64 `json:"black_women"`
	NativeMen     float64 `json:"native_men"`
	NativeWomen   float64 `json:"native_women"`
	AsianMen      float64 `json:"asian_men"`
	AsianWomen    float64 `json:"asian_women"`
	PacificMen    float64 `json:"pacific_men"`
	PacificWomen  float64 `json:"pacific_women"`
	MixedMen      float64 `json:"mixed_men"`
	MixedWomen    float64 `json:"mixed_women"`
	HispanicMen   float64 `json:"hispanic_men"`
	HispanicWomen float64 `json:"hispanic_women"`
}

func GetPerCountyDataForCounty(year, ageGroup, county int, dir string) (Demographic, error) {
	pathSprintf := path.Join(dir, "%d", "%d", "%d.json")
	path := fmt.Sprintf(pathSprintf, year, county, ageGroup)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return Demographic{}, err
	}
	var d Demographic
	err = json.Unmarshal(contents, &d)
	return d, err
}

func GetPerCountyDataFor(year, ageGroup int, dir string) ([]Demographic, error) {
	countiesDir := path.Join(dir, fmt.Sprintf("%d", year))
	files, err := ioutil.ReadDir(countiesDir)
	d := []Demographic{}
	if err != nil {
		return d, err
	}
	for _, file := range files {
		county, err := strconv.Atoi(file.Name())
		if err != nil {
			return d, err
		}
		data, err := GetPerCountyDataForCounty(year, ageGroup, county, dir)
		if err != nil {
			return d, err
		}
		d = append(d, data)
	}
	return d, nil
}
