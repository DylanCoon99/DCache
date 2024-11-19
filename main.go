/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	//"fmt"
	"log"
	//"strings"
	//"strconv"
	repl "github.com/openengineer/go-repl"
	"github.com/DylanCoon99/DCache/handler"
)

func main() {
	
	// This is all the REPL shit

	var handler handler.MyHandler
	handler.R = repl.NewRepl(handler)

	// start the terminal loop
 	if err := handler.R.Loop(); err != nil {
    	log.Fatal(err)
  	}


  	////////




}
