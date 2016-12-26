package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data/census/rawdata/CC-EST2015-ALLDATA.csv")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	files := make(map[int][]string)
	for scanner.Scan() {
		yearNum, err :=
			strconv.Atoi(strings.Split(scanner.Text(), ",")[5])
		if err != nil {
			log.Printf("invalid csv line: %s", scanner.Text())
		}
		if yearNum >= 3 {
			files[yearNum] = append(files[yearNum], scanner.Text())
		}
	}
	for num, entries := range files {
		f, err := os.Create(fmt.Sprintf("data/census/by_year/%d", num+2007))
		if err != nil {
			log.Fatalf("Could not output to file, err: %v", err)
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		for _, entry := range entries {
			w.WriteString(entry + "\n")
		}
	}
}
