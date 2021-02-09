package main

import (
		"fmt"
		"time"
)

type Bar struct {
	percent int64 // progress precentage
	cur int64	  // current progress
	total int64	  // total value for progress
	rate string	  // the actual progress bar to be printed
	graph string  // the fill value for progress bar
}

func (bar *Bar) NewOption(total int64)  {
	bar.cur = 0
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 1; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph // initial progress position
		fmt.Println(bar.rate)
	}
	
}

func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}

func (bar *Bar) NewOptionWithGraph(start, total int64, graph string) {
	bar.graph = graph
	bar.NewOption(total)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()

	if bar.percent != last {
		var i int64 = 0

		for ; i < bar.percent - last; i++ {
			bar.rate += bar.graph
		}
		fmt.Printf("\r[%-50s]%3d %% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
	}
	
}

func (bar *Bar) Finish() {
	fmt.Println()
}

func progressBar(nrOfLoops int64) {
	var bar Bar
	bar.NewOption(nrOfLoops)
	for i := 0; i <= int(nrOfLoops); i++ {
		time.Sleep(100*time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}

func main() {
	progressBar(25)
}