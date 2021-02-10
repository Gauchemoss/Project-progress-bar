package main

import (
		"fmt"
		"time"
)

//Variables that are needed for creating the progress bar.
type Progress_bar struct {
	percent int64 
	current_iteration_index int64	  
	total int64	  
	rate string	  
	fill_progressbar_char string  
}

//
func (bar *Progress_bar) Default(total int64) {
	bar.current_iteration_index = 0
	bar.percent = 0
	bar.total = total
}

func (bar *Progress_bar) Progressbar_sign_exists(progress_char string)  {
	if progress_char == "" {
		bar.fill_progressbar_char = "█"
	}else{
		bar.fill_progressbar_char = progress_char
	}
}

func (bar *Progress_bar) Count_percent() int64 {
	fmt.Println(int64(float32(bar.current_iteration_index) / float32(bar.total) * 100))
	return int64(float32(bar.current_iteration_index) / float32(bar.total) * 100)
}

func (bar *Progress_bar) Print(current_iteration_index int64) {
	bar.current_iteration_index = current_iteration_index
	last := bar.percent
	bar.percent = bar.Count_percent()

	var i int64 = 0

	for ; i < bar.percent - last; i++ {
		bar.rate += bar.fill_progressbar_char
	}

	fmt.Printf("\r╔══════════╗")
	fmt.Printf("\r║%-100s║ %3d %8d / %d", bar.rate, bar.percent, bar.current_iteration_index, bar.total)
	fmt.Printf("\r╚══════════╝")
	fmt.Println()
}

func progressBar(nrOfLoops int64, progress_char string) {
	var bar Progress_bar

	bar.Default(nrOfLoops);
	bar.Progressbar_sign_exists(progress_char)

	for i := 0; i <= int(bar.total); i++ {
		time.Sleep(50*time.Millisecond)
		bar.Print(int64(i))
	}
}

func main() {
	progressBar(25, "▌")
}