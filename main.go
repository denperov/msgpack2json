package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"

	"github.com/vmihailenco/msgpack"
)

var indentation bool
var mergeArrays bool

func init() {
	flag.BoolVar(&indentation, "i", false, "indentation")
	flag.BoolVar(&mergeArrays, "m", false, "merge-arrays")
}

func main() {
	flag.Parse()

	reader := os.Stdin
	writer := os.Stdout

	decoder := msgpack.NewDecoder(reader)
	encoder := json.NewEncoder(writer)
	if indentation {
		encoder.SetIndent("", "\t")
	}

	if !mergeArrays {
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
		begin := true
		for {
			var content []interface{}
			err := decoder.Decode(&content)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			for _, item := range content {
				if begin {
					begin = false
					_, err := writer.Write([]byte{'['})
					if err != nil {
						panic(err)
					}
				} else {
					_, err := writer.Write([]byte{','})
					if err != nil {
						panic(err)
					}
				}

				err := encoder.Encode(item)
				if err != nil {
					panic(err)
				}
			}
		}
		_, err := writer.Write([]byte{']'})
		if err != nil {
			panic(err)
		}
	}
}
