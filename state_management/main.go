package main

import (
	"context"
	"fmt"
	"state/state"
)

func main() {
	input := make(chan *state.WorkRequest, 100000000)
	output := make(chan *state.WorkResponce, 100000000)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go state.Processer(ctx, input, output)

	// req1 := state.WorkRequest{state.Add, 1, 2}
	// input <- &req1
	// req2 := state.WorkRequest{state.Subtract, 3, 2}
	// input <- &req2
	// req3 := state.WorkRequest{state.Multiply, 3, 2}
	// input <- &req3
	// req4 := state.WorkRequest{state.Divide, 3, 2}
	// input <- &

	for j := 0; j < 10000000; j++ {
		req := state.WorkRequest{state.Add, 1, 2}
		input <- &req
	}

	for i := 0; i < 100000000; i++ {
		resp := <-output
		fmt.Printf("Request: %v, Response: %v, Error: %v \n", resp.Wr, resp.Result, resp.Err)
	}
}
