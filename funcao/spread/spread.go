package main

import "fmt"

func media(medias ...float64) float64 {
	total := 0.0
	for _, num := range medias {
		total += num
	}

	return total / float64(len(medias))
}

func main() {
	fmt.Print(media(2, 5, 7, 9, 0))
}
