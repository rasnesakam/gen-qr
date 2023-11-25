package main

import qrcode "github.com/skip2/go-qrcode"

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gen-qr [message] [options]...")
		os.Exit(1)
	}

	var outfile string
	
	message := os.Args[1]
	
	flag.StringVar(&outfile, "o","qr.png","Out file")
	
	err := qrcode.WriteFile(message,qrcode.Medium, 256, outfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
		
}
