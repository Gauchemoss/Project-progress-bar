# Project-progress-bar

In this project i have made a simple progressbar for golang.

## Installation

You have to use the package manager [go get](https://golang.org/cmd/go/) for installing the progress bar:

```bash
go get github.com/Gauchemoss/Project-progress-bar/v1
```

## Example for using:
```go
package main
 
import (
    "fmt"
    "github.com/Gauchemoss/Project-progress-bar/v1"
)
 
func main() {
   nr_of iterations := 3;
   progressBar(nr_of_iterations, "*")
  
   for i := 0; i < nr_of_iterations; i++ {
       fmt.Println("hello world")
   }
}
```

### Result:
```bash
╔══════════════════════════════════════════════════╗
║**************************************************║ 100 %    3 / 3
╚══════════════════════════════════════════════════╝
Finished!!!

hello world
hello world
hello world
```

## Functionalities

You could change the filling sign with what ever you want. If you just writing "", the program will use █ as default. A few examples could be seen below:

### Examples

```go
progressBar(nr_of_iterations, "")
```
#### Result
```bash
╔══════════════════════════════════════════════════╗
║██████████████████████████████████████████████████║ 100 %    3 / 3
╚══════════════════════════════════════════════════╝
```

```go
progressBar(nr_of_iterations, "#")
```
#### Result
```bash
╔══════════════════════════════════════════════════╗
║##################################################║ 100 %    3 / 3
╚══════════════════════════════════════════════════╝
```

