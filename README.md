# Project-progress-bar
In this project i have made a simple progressbar for golang.

Installation
You have to implement the third-party module by typeing "go get github.com/Gauchemoss/Project-progress-bar" in your terminal window.

Example for using:

package main
 
import (
    "fmt"
	  "github.com/Gauchemoss/Project-progress-bar"
)
 
func main() {
  nr_of iterations := 3;
	progressBar(nr_of_iterations, "*")
  
   for i := 0; i < nr_of_iterations; i++ {
       fmt.Println("hello world")
   }
}

