package main

import (
	"flag"
	"fmt" //Imports the standard flag package
)

var name = flag.String("name", "world", "a name to say hello to") //Creates a new variable from a flag
var spanish bool                                                  //New variable to store flag value

func init() {
	//Sets variable to the flag value
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
}

func main() {
	//Parses the flags, placing E values in variables
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
}
