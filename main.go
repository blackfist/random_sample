// Package main is a program that takes a list of things
// from stdin and returns a randomly-selected sublist
package main

import (
	"os"
	"fmt"
	"bufio"
	"math/rand"
	"time"
)

func main() {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeNamedPipe != 0) {
		values := make([]string, 0)
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			values = append(values, reader.Text())
		}
		
		// initialize random and get an array of random indexes
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for _, index := range r.Perm(len(values)) {
			fmt.Println(values[index])
		}

	} else {
		fmt.Println("you gotta pipe in some shit, fool")
		fmt.Println("e.g. cat somefile | random_sample")
		fmt.Println("e.g. pbpaste | cut -d\" \" -f2 | random_sample")

	}
}