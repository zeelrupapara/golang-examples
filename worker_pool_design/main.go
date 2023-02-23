package main

import (
	"fmt"
	"worker-pool/pool"
)

func main() {
	cancel, input, output := pool.Dispatch(10)
	defer cancel()

	for i := 0; i < 10; i++ {
		input <- pool.WorkRequest{Operation: pool.Decrypt, Text: []byte("password"), Compare: []byte("$2a$10$kwz3HPzuFD2yk25lAzA1Zed58fxzRNXPYFPOkvqd6CowXOJavyLuO")}
	}

	for i := 0; i < 10; i++ {
		wr := <-output
		fmt.Printf("string %s match %v\n", wr.Wr.Text, wr.Matched)
	}
}
