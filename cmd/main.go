package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/limoges/sstats"
)

func main() {
	// Check that we have something connected to stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Input is pipe
	} else if stat.Size() == 0 {
		// No input pipe, no arg
		fmt.Fprintln(os.Stderr, "error: no input")
		return
	}

	var s []sstats.Statistic
	s = append(s, sstats.NewCount(0), new(sstats.Minimum), new(sstats.Maximum), sstats.NewSum(0),
		new(sstats.Stddev), new(sstats.Variance), new(sstats.Median))
	sstats := sstats.NewSimpleStats(s...)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		v, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			continue
		}
		sstats.Push(v)
	}

	stats := sstats.String()
	stats = strings.Replace(stats, ":", ":\t", -1)
	stats = strings.Replace(stats, ",", "\n", -1)
	fmt.Println(stats)
}
