package state

import (
	"context"
	"fmt"
)

func Processer(ctx context.Context, input chan *WorkRequest, output chan *WorkResponce) {
	for {
		select {
		case <-ctx.Done():
			output := <-ctx.Done()
			fmt.Println(output)
			return
		case wr := <-input:
			output <- Process(wr)
		}
	}
}
