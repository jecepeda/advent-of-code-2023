/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"time"

	"github.com/jecepeda/advent-of-code-2023/day1"
	"github.com/jecepeda/advent-of-code-2023/day2"
	"github.com/jecepeda/advent-of-code-2023/day3"
	"github.com/jecepeda/advent-of-code-2023/day4"
	"github.com/jecepeda/advent-of-code-2023/day5"
	"github.com/jecepeda/advent-of-code-2023/day6"
	"github.com/jecepeda/advent-of-code-2023/day7"
	"github.com/jecepeda/advent-of-code-2023/day8"
	"github.com/jecepeda/advent-of-code-2023/day9"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

type dayFunc func(filename string) (int, error)

type execFunc struct {
	filename    string
	function    dayFunc
	description string
	longRunning bool
}

var dayFunctions = map[string][]execFunc{
	"1": {
		execFunc{
			filename:    "day1/input.txt",
			function:    day1.Part1,
			description: "Day 1 Part 1",
		},
		execFunc{
			filename:    "day1/input.txt",
			function:    day1.Part2,
			description: "Day 1 Part 2",
		},
	},
	"2": {
		execFunc{
			filename:    "day2/input.txt",
			function:    day2.Part1,
			description: "Day 2 Part 1",
		},
		execFunc{
			filename:    "day2/input.txt",
			function:    day2.Part2,
			description: "Day 2 Part 2",
		},
	},
	"3": {
		execFunc{
			filename:    "day3/input.txt",
			function:    day3.Part1,
			description: "Day 3 Part 1",
		},
		execFunc{
			filename:    "day3/input.txt",
			function:    day3.Part2,
			description: "Day 3 Part 2",
		},
	},
	"4": {
		execFunc{
			filename:    "day4/input.txt",
			function:    day4.Part1,
			description: "Day 4 Part 1",
		},
		execFunc{
			filename:    "day4/input.txt",
			function:    day4.Part2,
			description: "Day 4 Part 2",
		},
	},
	"5": {
		execFunc{
			filename:    "day5/input.txt",
			function:    day5.Part1,
			description: "Day 5 Part 1",
		},
		execFunc{
			filename:    "day5/input.txt",
			function:    day5.Part2,
			description: "Day 5 Part 2",
		},
	},
	"6": {
		execFunc{
			filename:    "day6/input.txt",
			function:    day6.Part1,
			description: "Day 6 Part 1",
		},
		execFunc{
			filename:    "day6/input.txt",
			function:    day6.Part2,
			description: "Day 6 Part 2",
		},
	},
	"7": {
		execFunc{
			filename:    "day7/input.txt",
			function:    day7.Part1,
			description: "Day 7 Part 1",
		},
		execFunc{
			filename:    "day7/input.txt",
			function:    day7.Part2,
			description: "Day 7 Part 2",
		},
	},
	"8": {
		execFunc{
			filename:    "day8/input.txt",
			function:    day8.Part1,
			description: "Day 8 Part 1",
		},
		execFunc{
			filename:    "day8/input.txt",
			function:    day8.Part2,
			description: "Day 8 Part 2",
		},
	},
	"9": {
		execFunc{
			filename:    "day9/input.txt",
			function:    day9.Part1,
			description: "Day 9 Part 1",
		},
		execFunc{
			filename:    "day9/input.txt",
			function:    day9.Part2,
			description: "Day 9 Part 2",
		},
	},
}

func getDays() []string {
	keys := make([]string, 0, len(dayFunctions))
	for k := range dayFunctions {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "This command runs the solution for any given day",
	Run: func(cmd *cobra.Command, args []string) {
		var days []string
		if len(args) > 0 {
			if _, ok := dayFunctions[args[0]]; !ok {
				fmt.Printf("Invalid day: %s\n", args[0])
				return
			}
			days = []string{args[0]}
		} else {
			days = getDays()
		}
		force := cmd.Flag("force").Value.String() == "true"
		w := table.NewWriter()
		w.AppendHeader(table.Row{"Title", "Time", "Result", "Notes"})
		for _, day := range days {
			for _, f := range dayFunctions[day] {
				if !force && f.longRunning {
					w.AppendRow(table.Row{f.description, "0s", "SKIPPED", "use --force to run"})
					continue
				} else {
					now := time.Now()
					result, err := f.function(f.filename)
					if err != nil {
						fmt.Println(err)
						continue
					}
					w.AppendRow(table.Row{f.description, time.Since(now), result, ""})
				}
			}
		}
		fmt.Println(w.Render())
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolP("force", "f", false, "Forces long-running parts to run")
}
