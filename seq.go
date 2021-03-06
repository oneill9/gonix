package main

import "os"
import "fmt"

func main() {
	if len(os.Args) == 2 {
		var end int
		fmt.Sscanf(os.Args[1], "%d", &end)
		for i := 1; i <= end; i++ {
			fmt.Println(i)
		}
	} else if len(os.Args) == 3 {
		var start, end int
		fmt.Sscanf(os.Args[1], "%d", &start)
		fmt.Sscanf(os.Args[2], "%d", &end)
		for i := start; i <= end; i++ {
			fmt.Println(i)
		}
	} else if len(os.Args) == 4 {
		var start, inc, end int
		fmt.Sscanf(os.Args[1], "%d", &start)
		fmt.Sscanf(os.Args[2], "%d", &inc)
		fmt.Sscanf(os.Args[3], "%d", &end)
		for i := start; i <= end; i += inc {
			fmt.Println(i)
		}
	}
}
