// Package examples provides a collection of examples for this project
package internal

import (
	"fmt"

	"github.com/engmtcdrm/go-eggy"
	pp "github.com/engmtcdrm/go-prettyprint"
	"github.com/engmtcdrm/stringfolds"
)

var AllExamples = []eggy.Example{
	{
		Name: "CutPrefixFold",
		Fn:   ExampleCutPrefixFold,
	},
	{
		Name: "CutSuffixFold",
		Fn:   ExampleCutSuffixFold,
	},
	{
		Name: "HasPrefixFold",
		Fn:   ExampleHasPrefixFold,
	},
	{
		Name: "HasSuffixFold",
		Fn:   ExampleHasSuffixFold,
	},
	{
		Name: "TrimPrefixFold",
		Fn:   ExampleTrimPrefixFold,
	},
	{
		Name: "TrimSuffixFold",
		Fn:   ExampleTrimSuffixFold,
	},
}

func ExampleCutPrefixFold() {
	example1Result, example1Found := stringfolds.CutPrefixFold("Hello, World!", "hello")
	example2Result, example2Found := stringfolds.CutPrefixFold("Hello, World!", "HELLO")
	example3Result, example3Found := stringfolds.CutPrefixFold("Hello, World!", "world")
	example4Result, example4Found := stringfolds.CutPrefixFold("Hello, World!", "WORLD")

	fmt.Println(pp.Cyan("CutPrefixFold") + " examples:")
	fmt.Println()
	fmt.Printf("  1. %s returns %q and found = %v\n", pp.Green("stringfolds.CutPrefixFold(\"Hello, World!\", \"hello\")"), example1Result, example1Found)
	fmt.Printf("  2. %s returns %q and found = %v\n", pp.Green("stringfolds.CutPrefixFold(\"Hello, World!\", \"HELLO\")"), example2Result, example2Found)
	fmt.Printf("  3. %s returns %q and found = %v\n", pp.Green("stringfolds.CutPrefixFold(\"Hello, World!\", \"world\")"), example3Result, example3Found)
	fmt.Printf("  4. %s returns %q and found = %v\n", pp.Green("stringfolds.CutPrefixFold(\"Hello, World!\", \"WORLD\")"), example4Result, example4Found)
}

func ExampleCutSuffixFold() {
	example1Result, example1Found := stringfolds.CutSuffixFold("Hello, World!", "hello")
	example2Result, example2Found := stringfolds.CutSuffixFold("Hello, World!", "HELLO")
	example3Result, example3Found := stringfolds.CutSuffixFold("Hello, World!", "world!")
	example4Result, example4Found := stringfolds.CutSuffixFold("Hello, World!", "WORLD!")

	fmt.Println(pp.Cyan("CutSuffixFold") + " examples:")
	fmt.Println()
	fmt.Printf("  1. %s returns %q and found = %v\n", pp.Green("stringfolds.CutSuffixFold(\"Hello, World!\", \"hello\")"), example1Result, example1Found)
	fmt.Printf("  2. %s returns %q and found = %v\n", pp.Green("stringfolds.CutSuffixFold(\"Hello, World!\", \"HELLO\")"), example2Result, example2Found)
	fmt.Printf("  3. %s returns %q and found = %v\n", pp.Green("stringfolds.CutSuffixFold(\"Hello, World!\", \"world!\")"), example3Result, example3Found)
	fmt.Printf("  4. %s returns %q and found = %v\n", pp.Green("stringfolds.CutSuffixFold(\"Hello, World!\", \"WORLD!\")"), example4Result, example4Found)
}

func ExampleHasPrefixFold() {
	example1Result := stringfolds.HasPrefixFold("Hello, World!", "hello")
	example2Result := stringfolds.HasPrefixFold("Hello, World!", "HELLO")
	example3Result := stringfolds.HasPrefixFold("Hello, World!", "world")
	example4Result := stringfolds.HasPrefixFold("Hello, World!", "WORLD")

	fmt.Println(pp.Cyan("HasPrefixFold") + " examples:")
	fmt.Println()
	fmt.Println("  1. "+pp.Green("stringfolds.HasPrefixFold(\"Hello, World!\", \"hello\")"), "returns", pp.Yellow(example1Result))
	fmt.Println("  2. "+pp.Green("stringfolds.HasPrefixFold(\"Hello, World!\", \"HELLO\")"), "returns", pp.Yellow(example2Result))
	fmt.Println("  3. "+pp.Green("stringfolds.HasPrefixFold(\"Hello, World!\", \"world\")"), "returns", pp.Yellow(example3Result))
	fmt.Println("  4. "+pp.Green("stringfolds.HasPrefixFold(\"Hello, World!\", \"WORLD\")"), "returns", pp.Yellow(example4Result))
}

func ExampleHasSuffixFold() {
	example1Result := stringfolds.HasSuffixFold("Hello, World!", "hello")
	example2Result := stringfolds.HasSuffixFold("Hello, World!", "HELLO")
	example3Result := stringfolds.HasSuffixFold("Hello, World!", "world!")
	example4Result := stringfolds.HasSuffixFold("Hello, World!", "WORLD!")

	fmt.Println(pp.Cyan("HasSuffixFold") + " examples:")
	fmt.Println()
	fmt.Println("  1. "+pp.Green("stringfolds.HasSuffixFold(\"Hello, World!\", \"hello\")"), "returns", pp.Yellow(example1Result))
	fmt.Println("  2. "+pp.Green("stringfolds.HasSuffixFold(\"Hello, World!\", \"HELLO\")"), "returns", pp.Yellow(example2Result))
	fmt.Println("  3. "+pp.Green("stringfolds.HasSuffixFold(\"Hello, World!\", \"world!\")"), "returns", pp.Yellow(example3Result))
	fmt.Println("  4. "+pp.Green("stringfolds.HasSuffixFold(\"Hello, World!\", \"WORLD!\")"), "returns", pp.Yellow(example4Result))
}

// TrimPrefixFold
func ExampleTrimPrefixFold() {
	example1Result := stringfolds.TrimPrefixFold("Hello, World!", "hello")
	example2Result := stringfolds.TrimPrefixFold("Hello, World!", "HELLO")
	example3Result := stringfolds.TrimPrefixFold("Hello, World!", "world")
	example4Result := stringfolds.TrimPrefixFold("Hello, World!", "WORLD")

	fmt.Println(pp.Cyan("TrimPrefixFold") + " examples:")
	fmt.Println()
	fmt.Println("  1. "+pp.Green("stringfolds.TrimPrefixFold(\"Hello, World!\", \"hello\")"), "returns", pp.Yellow(example1Result))
	fmt.Println("  2. "+pp.Green("stringfolds.TrimPrefixFold(\"Hello, World!\", \"HELLO\")"), "returns", pp.Yellow(example2Result))
	fmt.Println("  3. "+pp.Green("stringfolds.TrimPrefixFold(\"Hello, World!\", \"world\")"), "returns", pp.Yellow(example3Result))
	fmt.Println("  4. "+pp.Green("stringfolds.TrimPrefixFold(\"Hello, World!\", \"WORLD\")"), "returns", pp.Yellow(example4Result))
}

func ExampleTrimSuffixFold() {
	example1Result := stringfolds.TrimSuffixFold("Hello, World!", "hello")
	example2Result := stringfolds.TrimSuffixFold("Hello, World!", "HELLO")
	example3Result := stringfolds.TrimSuffixFold("Hello, World!", "world!")
	example4Result := stringfolds.TrimSuffixFold("Hello, World!", "WORLD!")

	fmt.Println(pp.Cyan("TrimSuffixFold") + " examples:")
	fmt.Println()
	fmt.Println("  1. "+pp.Green("stringfolds.TrimSuffixFold(\"Hello, World!\", \"hello\")"), "returns", pp.Yellow(example1Result))
	fmt.Println("  2. "+pp.Green("stringfolds.TrimSuffixFold(\"Hello, World!\", \"HELLO\")"), "returns", pp.Yellow(example2Result))
	fmt.Println("  3. "+pp.Green("stringfolds.TrimSuffixFold(\"Hello, World!\", \"world!\")"), "returns", pp.Yellow(example3Result))
	fmt.Println("  4. "+pp.Green("stringfolds.TrimSuffixFold(\"Hello, World!\", \"WORLD!\")"), "returns", pp.Yellow(example4Result))
}
