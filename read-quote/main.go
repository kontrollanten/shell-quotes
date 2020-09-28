package main

import (
	"bufio"
	"github.com/fatih/color"
	"github.com/mitchellh/go-wordwrap"
	"math/rand"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result
}

func main() {
	filePath := os.Args[1]

	f := LinesInFile(filePath)

	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(len(f))

	red := color.New(color.FgBlack)
	whiteBackground := red.Add(color.BgCyan)
	whiteBackground.Println(wordwrap.WrapString(f[rnd], 60))
}
