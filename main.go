package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"sync"
)

func checkDomain(d string, wg *sync.WaitGroup) {
	if !validateDomainTLD(d) {
		color.New(color.BgHiRed, color.FgBlack).Print(" ")
		fmt.Println(d, "INVALID TLD")

		wg.Done()
		return
	}

	exist, err := domainExists(d)

	if err != nil {
		color.New(color.BgHiRed, color.FgBlack).Print(" ")
		fmt.Println(d, "ERR", err)
	} else {
		if exist {
			color.New(color.BgRed, color.FgBlack).Print(" ")
			fmt.Println(d, "TAKEN")
		} else {
			color.New(color.BgGreen, color.FgBlack).Print(" ")
			fmt.Println(d, "AVAIL")
		}
	}

	wg.Done()
}

func main() {
	domains := os.Args[1:]

	wg := &sync.WaitGroup{}

	for _, d := range domains {
		wg.Add(1)

		go checkDomain(d, wg)
	}

	wg.Wait()
}