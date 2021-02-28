package bar

import (
	"fmt"
	//"os"
	//"os/exec"
	//"runtime"
	"time"
	"bar"
)

//Variables that are needed for creating the progress bar.
type Progress_bar struct {
	percent int64 
	current_iteration_index int64	  
	total int64	  
	rate string	  
	fill_progressbar_char string
	fill_counter int
}

func (bar *Progress_bar) Default_values(total int64) {
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
	return int64(float32(bar.current_iteration_index) / float32(bar.total) * 100)
}

func (bar *Progress_bar) Print(current_iteration_index int64) {
	bar.current_iteration_index = current_iteration_index
	last_percent := bar.percent
	bar.percent = bar.Count_percent()

	//check if the rate go out of frame.
	if bar.fill_counter < 50 {
		//filling the frame with characters.
		for i := int64(0); i < bar.percent - last_percent; i += 2 {
			bar.rate += bar.fill_progressbar_char
			bar.fill_counter++
		}
	}

	//Code for clean the terminal after each print.
	
	//time.Sleep(500*time.Millisecond)

	if(bar.percent != 100) {
		fmt.Printf("\r Loading!...")
	} else {
		fmt.Printf("\r Finished!!!")
	}

	fmt.Printf("\r╔══════════════════════════════════════════════════╗\n")
	fmt.Printf("\r║%-50s║ %3d %% %4d / %d\n", bar.rate, bar.percent, bar.current_iteration_index, bar.total)
	fmt.Printf("\r╚══════════════════════════════════════════════════╝\n")
	fmt.Println()

}

func progressBar(nrOfLoops int64, progress_char string) {
	var bar Progress_bar

	bar.Default_values(nrOfLoops);
	bar.Progressbar_sign_exists(progress_char)

	for i := 0; i <= int(bar.total); i++ {
		time.Sleep(500*time.Millisecond)
		fmt.Print("\033[2J")
		bar.Print(int64(i))
		
	}
}
