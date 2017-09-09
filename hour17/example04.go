package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		usageText := `Usage example04 [OPTION]
An example of customising usage output

  -s, --s         example string argument, default: String help text
  -i, --i         example integer argument, default: Int help text
  -b, --b         example boolean argument, default: Bool help text`
		fmt.Fprintf(os.Stderr, "%s\n", usageText)

	}
	s := flag.String("s", "Hello world", "String help text")
	i := flag.Int("i", 1, "Int help text")
	b := flag.Bool("b", true, "Bool help text")
	flag.Parse()
	fmt.Println("value of s:", *s)
	fmt.Println("value of i:", *i)
	fmt.Println("value of b:", *b)
}
