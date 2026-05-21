package main

import (
	"github.com/engmtcdrm/go-eggy"
	pp "github.com/engmtcdrm/go-prettyprint"

	"github.com/engmtcdrm/stringfolds/examples/internal"
)

func main() {
	eggy.NewExamplePrompt(internal.AllExamples).
		Title(pp.Yellow("Examples of stringfolds package")).
		Show()
}
