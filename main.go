package main

import (
	_ "embed"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	pg_query "github.com/pganalyze/pg_query_go/v5"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: pgmodelgen <file>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	compiler := NewCompiler()
	parse, err := pg_query.Parse(string(b))
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	err = compiler.ParseStatements(parse)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	spew.Dump(compiler.Catalog)
}
