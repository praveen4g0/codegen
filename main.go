package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/praveen4g0/codegen/pkg/codegen"
)

func main() {
	var u codegen.User
	in, _ := os.Open("user.json")
	b, _ := io.ReadAll(in)
	json.Unmarshal(b, &u)
	fmt.Printf("website: %+v\n", u.Website)
}
