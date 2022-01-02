package main

import (
	"fmt"
	flesch_go "github.com/NICEXAI/flesch-go"
)

func main() {
	tContent := `Immune response, MT and HSP70 gene expression, and bioaccumulation induced by lead exposure of the marine crab, Charybdis japonica.`

	document, err := flesch_go.ParseString(tContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get score
	fmt.Println(document.Score())

	// get word count
	fmt.Println(document.WordCount())
}
