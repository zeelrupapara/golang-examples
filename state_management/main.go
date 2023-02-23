package main

import (
	"context"
	"fmt"
	"state/state"
)

func main() {
	input := make(chan *state.WorkRequest, 2)
	output := make(chan *state.WorkResponce, 2)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go state.Processer(ctx, input, output)

	req1 := state.WorkRequest{state.Add, 1, 2}
	input <- &req1
	req2 := state.WorkRequest{state.Subtract, 3, 2}
	input <- &req2
	req3 := state.WorkRequest{state.Multiply, 3, 2}
	input <- &req3
	req4 := state.WorkRequest{state.Divide, 3, 2}
	input <- &req4
	req5 := state.WorkRequest{state.Divide, 3, 0}
	input <- &req5

	for i := 0; i < 5; i++ {
		resp := <-output
		fmt.Printf("Request: %v, Response: %v, Error: %v \n", resp.Wr, resp.Result, resp.Err)
	}
}
