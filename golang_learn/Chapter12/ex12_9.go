package main

import (
	"flag" // command line option parser
	"os"
	"fmt"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	fmt.Println(flag.NArg())
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
		fmt.Println("BBB")
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				fmt.Println("AAA")
				s += Newline
			}
		}
		
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}