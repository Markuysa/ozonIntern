package main

import (
	"context"
	"flag"
)

func main() {

	repositoryMode := flag.Bool("db", false, "repository mode: if -db is specified => the program will use PostgreSQL")
	flag.Parse()
	ctx := context.Background()

}
