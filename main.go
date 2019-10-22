package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"

	"github.com/vmihailenco/msgpack"
)

var indentation bool
var j2m bool

func init() {
	flag.BoolVar(&indentation, "i", false, "indentation")
	flag.BoolVar(&j2m, "j2m", false, "json to msgpack")
}

func main() {
	flag.Parse()

	reader := os.Stdin
	writer := os.Stdout

	if j2m {
		decoder := json.NewDecoder(reader)
		encoder := msgpack.NewEncoder(writer)

		for {
			var content interface{}
			err := decoder.Decode(&content)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			err = encoder.Encode(content)
			if err != nil {
				panic(err)
			}
		}
	} else {
		decoder := msgpack.NewDecoder(reader)
		encoder := json.NewEncoder(writer)
		if indentation {
			encoder.SetIndent("", "\t")
		}

		for {
			var content interface{}
			err := decoder.Decode(&content)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			err = encoder.Encode(content)
			if err != nil {
				panic(err)
			}
		}
	}
}
