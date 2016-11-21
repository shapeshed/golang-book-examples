package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func flagUsage() {
	usageText := `example06 is an example cli tool.
        
Usage:

example06 command [arguments]

The commands are:

uppercase  uppercase a string 
lowercase  lowercase a string

Use "example06 [command] --help" for more information about a command.`
	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func uppercaseCmdUsage() {
	usageText := `usage: uppercase [-s]
        
    -s
        A string to be uppercased`
	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func lowercaseCmdUsage() {
	usageText := `usage: lowercase [-s]
        
    -s
        A string to be lowercased`
	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func main() {

	flag.Usage = flagUsage

	uppercaseCmd := flag.NewFlagSet("uppercase", flag.ExitOnError)
	uppercaseCmd.Usage = uppercaseCmdUsage
	lowercaseCmd := flag.NewFlagSet("lowercase", flag.ExitOnError)
	lowercaseCmd.Usage = lowercaseCmdUsage

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "uppercase":
		s := uppercaseCmd.String("s", "", "A string of text to be uppercased")
		uppercaseCmd.Parse(os.Args[2:])
		if *s == "" {
			uppercaseCmd.Usage()
			return
		}
		fmt.Println(strings.ToUpper(*s))
	case "lowercase":
		s := lowercaseCmd.String("s", "", "A string of text to be lowercased")
		lowercaseCmd.Parse(os.Args[2:])
		if *s == "" {
			lowercaseCmd.Usage()
			return
		}
		fmt.Println(strings.ToLower(*s))
	default:
		flag.Usage()
	}

}
