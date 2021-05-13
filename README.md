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
    v1 "github.com/Gauchemoss/Project-progress-bar/v1"
)

func main() {
    nr_of_iterations := 3
    v1.ProgressBar(nr_of_iterations, "*")

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
v1.ProgressBar(nr_of_iterations, "")
```

#### Result
```bash
╔══════════════════════════════════════════════════╗
║██████████████████████████████████████████████████║ 100 %    3 / 3
╚══════════════════════════════════════════════════╝
```

```go
v1.ProgressBar(nr_of_iterations, "#")
```
#### Result
```bash
╔══════════════════════════════════════════════════╗
║##################################################║ 100 %    3 / 3
╚══════════════════════════════════════════════════╝
```
### OBS!!!
You must set nr_of_iterations to an iteration, or your program will crash. 
Another thing is that you have to set nr_of_iterations to an positive. If you set it to a negative integer, the program will write out "Nr of loops must be positive integer!!!".

