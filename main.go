package main

import (
	"flag"
	"github.com/krig/go-sox"
	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	// All libSoX applications must start by initializing the SoX library
	if !sox.Init() {
		log.Fatal("Failed to initialize SoX")
	}
	// Make sure to call Quit before terminating
	defer sox.Quit()

	// Open the input file (with default parameters)
	in := sox.OpenRead(flag.Arg(0))
	if in == nil {
		log.Fatal("Failed to open input file")
	}
	// Close the file before exiting
	defer in.Release()

	log.Println(in.Type())
	log.Println("length:", in.Signal().Length(), "channels:", in.Signal().Channels(), "rate:", in.Signal().Rate())
	log.Println(in.Signal().Length() / (uint64(in.Signal().Rate()) * uint64(in.Signal().Channels())))
}
