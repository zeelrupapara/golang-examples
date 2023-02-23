package pool

import (
	"context"
	"fmt"
)

func Dispatch(numWorkers int) (context.CancelFunc, chan WorkRequest, chan WorkResponce) {
	input := make(chan WorkRequest)
	output := make(chan WorkResponce)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	for i := 0; i < numWorkers; i++ {
		go Worker(ctx, i, input, output)
	}
	return cancel, input, output
}

func Worker(ctx context.Context, id int, input chan WorkRequest, output chan WorkResponce) {
	for {
		select {
		case <-ctx.Done():
			return
		case wr := <-input:
			fmt.Printf("working id : %d, performing %s work\n", id, wr.Operation)
			output <- Process(wr)
		}
	}
}
