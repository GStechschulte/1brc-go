package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	Station string
	Temp    float64
}

type Stats struct {
	Min, Mean, Max float64
}

func main() {

	file, err := os.Open("data/measurements.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stationTemps := make(map[string][]float64)

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024) // 64KB initial buffer size
	scanner.Buffer(buf, 1024*1024)  // Max buffer size 1MB

	for scanner.Scan() {
		// Parse each line into the stationTemps map
		parts := strings.Split(scanner.Text(), ";")
		temp, _ := strconv.ParseFloat(parts[1], 64)
		stationTemps[parts[0]] = append(stationTemps[parts[0]], temp)
	}

	// Calculate the min, mean, and max for each station
	results := make(map[string]Stats)
	for station, temps := range stationTemps {
		min, max, sum := temps[0], temps[0], 0.0

		for _, temp := range temps {
			if temp < min {
				min = temp
			}
			if temp > max {
				max = temp
			}
			sum += temp
		}
		mean := sum / float64(len(temps))
		results[station] = Stats{Min: min, Mean: mean, Max: max}
	}
}
