package progress_bar

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
	fill_counter int
}

func (bar *Progress_bar) Default_values(total int64) {
	bar.current_iteration_index = 0	//Tells the program to start at iteration 0.
	bar.percent = 0					//Tells the program to start at 0%.
	bar.total = total				//Tells the program what the total amount of iterations is.
}


//This function is used to let the user chanfe the filling sign in the progressbar.
//If none is selected, an █ will be used.
func (bar *Progress_bar) Progressbar_sign_exists(progress_char string)  {
	if progress_char == "" {
		bar.fill_progressbar_char = "█"
	}else{
		bar.fill_progressbar_char = progress_char
	}
}

//This function count what the current percentage should be, depending what the current iteration index and total amont of iterations.
func (bar *Progress_bar) Count_percent() int64 {
	return int64(float32(bar.current_iteration_index) / float32(bar.total) * 100)
}

//This function is used for printing to the terminal.
func (bar *Progress_bar) Print(current_iteration_index int64) {
	bar.current_iteration_index = current_iteration_index
	last_percent := bar.percent
	bar.percent = bar.Count_percent()

	
	
		//filling the frame with characters.
		for i := int64(0); i < bar.percent - last_percent; i += 2 {

			//check if the rate go out of frame.
			if bar.fill_counter < 50 {
				bar.rate += bar.fill_progressbar_char
				bar.fill_counter++
			}
		}

	//Print the progress bar on the terminal.
	fmt.Printf("\r╔══════════════════════════════════════════════════╗\n")
	fmt.Printf("\r║%-50s║ %3d %% %4d / %d\n", bar.rate, bar.percent, bar.current_iteration_index, bar.total)
	fmt.Printf("\r╚══════════════════════════════════════════════════╝\n")
	if(bar.percent == 100) {
		fmt.Println("Finished!!!")
	}else {
		fmt.Println("Loading!!!")
	}
	fmt.Println()

}

func progressBar(nrOfLoops int, progress_char string) {
	var bar Progress_bar

	bar.Default_values(int64(nrOfLoops));
	bar.Progressbar_sign_exists(progress_char)

	for i := 0; i <= int(bar.total); i++ {
		time.Sleep(50*time.Millisecond)

		//Code for clean the terminal after each print.
		fmt.Print("\033[2J")

		bar.Print(int64(i))
	}
	time.Sleep(50*time.Millisecond)
}
