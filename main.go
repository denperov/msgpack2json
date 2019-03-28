package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"

	"github.com/vmihailenco/msgpack"
)

var indentation bool

func init() {
	flag.BoolVar(&indentation, "i", false, "indentation")
}

func main() {
	flag.Parse()

	decoder := msgpack.NewDecoder(os.Stdin)
 	encoder := json.NewEncoder(os.Stdout)
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
