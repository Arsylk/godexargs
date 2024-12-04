package main

import (
	"flag"
	"log"

	"github.com/Arsylk/godexargs/provider"
	"github.com/gookit/goutil/dump"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "", "dex file")
	flag.StringVar(&file, "file", "", "dex file")
	flag.Parse()
	if len(file) == 0 {
		args := flag.Args()
		if len(args) > 0 {
			file = args[0]
		} else {
			log.Fatalln("error: no input file specified")
		}
	}

	p, err := provider.NewDexTkProvider(file)
	if err != nil {
		log.Fatalln(err)
	}
	classes := p.ListClasses()
	for _, clazz := 
}
