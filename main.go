package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	in    = flag.String("in", "", "base64 encoded input or embedded JSON with base64 field")
	err   error
	input []byte
)

func main() {
	flag.Parse()
	if *in != "" {
		input = []byte(*in)
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	}

	// first attempt to base64 decode the input
	decoded, decodeErr := base64.StdEncoding.DecodeString(string(input))
	if decodeErr == nil {
		input = decoded
	}

	var (
		out    map[string]interface{}
		output []byte
	)
	jsonErr := json.Unmarshal(input, &out)
	if jsonErr == nil {
		for key, val := range out {
			switch val.(type) {
			case string:
				decoded, decodeErr := base64.StdEncoding.DecodeString(val.(string))
				if decodeErr == nil {
					var i map[string]interface{}
					iErr := json.Unmarshal(decoded, &i)
					if iErr == nil {
						out[key] = i
					} else {
						out[key] = string(decoded)
					}
				}
			}
		}
		output, jsonErr = json.MarshalIndent(out, "", "  ")
		if jsonErr != nil {
			output = input
		}
	} else {
		output = input
	}

	fmt.Fprint(os.Stdout, string(output))
}
