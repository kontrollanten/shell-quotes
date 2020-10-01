package main

import (
	"fmt"
	"os"
)

type Quote struct {
	Author string
	Text   string
}

func main() {
	DownloadCatV()
	DownloadAurelius()
}

func WriteToFile(quotes []Quote, filename string) {
	targetPath := os.Args[1]

	fmt.Println("targetPath:")
	fmt.Println(targetPath)
	f, e := os.Create(fmt.Sprintf("%s/%s", targetPath, filename))

	if e != nil {
		panic(e)
	}

	defer f.Close()

	for _, q := range quotes {
		row := fmt.Sprintf("%s,%s\n", q.Text, q.Author)
		_, e2 := f.WriteString(row)

		if e2 != nil {
			panic(e2)
		}
	}
}
