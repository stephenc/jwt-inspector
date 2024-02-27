package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	encoded := os.Getenv("JWT_TOKEN")
	if len(encoded) == 0 {
		fmt.Println("${JWT_TOKEN} is empty")
		return
	}
	split := strings.Split(encoded, ".")
	for i := 0; i < len(split); i++ {
		tokenBytes, err := base64.RawStdEncoding.DecodeString(split[i])
		if err != nil {
			return
		}
		var decoded interface{}
		if err := json.Unmarshal(tokenBytes, &decoded); err != nil {
			var sToken = string(tokenBytes)
			fmt.Printf("%s\n", sToken)
		} else {
			bytes, _ := json.MarshalIndent(decoded, "", "  ")
			var sToken = string(bytes)
			fmt.Printf("%s\n", sToken)
		}
	}
}
