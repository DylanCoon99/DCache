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


	// I think I will add the instance of the database to the handler as a ptr 

	// I will create the instance of the database here and start it


	// start the terminal loop
 	if err := handler.R.Loop(); err != nil {
    	log.Fatal(err)
  	}


  	////////




}
