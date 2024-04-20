package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"time"
)

// Custom CLI flags
var (
	timeField  string
	timeFormat string
)

// DurationLine stores the duration and the corresponding line for sorting
type DurationLine struct {
	Duration time.Duration
	Line     string
}

func init() {
	// Setting up CLI flags
	flag.StringVar(&timeField, "field", "time", "JSON field to extract the timestamp from")
	flag.StringVar(&timeFormat, "format", time.RFC3339, "Format of the timestamp in the JSON field")
}

func main() {
	flag.Parse() // Parse any CLI flags

	scanner := bufio.NewScanner(os.Stdin)
	var previousTime time.Time
	firstLine := true
	var durationLines []DurationLine

	for scanner.Scan() {
		line := scanner.Text()
		currentTime := getTimeFromLine(line)

		if !firstLine {
			duration := currentTime.Sub(previousTime)
			fmt.Printf("%s\t%s\n", formatDuration(duration), line)
			durationLines = append(durationLines, DurationLine{Duration: duration, Line: line})
		} else {
			firstLine = false
			fmt.Printf("%s\t%s\n", formatDuration(0), line) // Print 0 for the first line as there is no previous line
		}
		previousTime = currentTime
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %v\n", err)
		os.Exit(1)
	}

	// Sort and print the top 5 durations
	sort.Slice(durationLines, func(i, j int) bool {
		return durationLines[i].Duration > durationLines[j].Duration
	})
	fmt.Println("Top 5 longest durations:")
	for i := 0; i < len(durationLines) && i < 5; i++ {
		fmt.Printf("%s\t%s\n", formatDuration(durationLines[i].Duration), durationLines[i].Line)
	}
}

func formatDuration(d time.Duration) string {
	h := d / time.Hour
	m := (d % time.Hour) / time.Minute
	s := (d % time.Minute) / time.Second
	ms := (d % time.Second) / time.Millisecond
	ns := d % time.Millisecond

	var result string
	if h > 0 {
		result = fmt.Sprintf("%3dh", h)
	} else {
		result += "    " // 4 spaces for hour
	}
	if m > 0 || h > 0 {
		result += fmt.Sprintf("%3dm", m)
	} else {
		result += "    " // 4 spaces for minute
	}
	if s > 0 || m > 0 || h > 0 {
		result += fmt.Sprintf("%3ds", s)
	} else {
		result += "    " // 4 spaces for second
	}
	if ms > 0 || s > 0 || m > 0 || h > 0 {
		result += fmt.Sprintf("%4dms", ms)
	} else {
		result += "      " // 6 spaces for ms if less than 1ms
	}
	result += fmt.Sprintf("%7dns", ns)

	return fmt.Sprintf("%-27s", result) // Ensure the total length is 29 characters
}

// getTimeFromLine attempts to extract and parse the timestamp from the line based on provided flags.
func getTimeFromLine(line string) time.Time {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(line), &data)
	if err == nil {
		if timestamp, ok := data[timeField]; ok {
			if tsStr, ok := timestamp.(string); ok {
				if t, err := time.Parse(timeFormat, tsStr); err == nil {
					return t
				}
			}
		}
	}
	return time.Now() // Default to current time if any step fails
}
