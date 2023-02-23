package state

import "errors"

func Process(wr *WorkRequest) *WorkResponce {
	resp := WorkResponce{Wr: wr}

	switch wr.Op {
	case Add:
		resp.Result = wr.Value1 + wr.Value2
	case Subtract:
		resp.Result = wr.Value1 - wr.Value2
	case Multiply:
		resp.Result = wr.Value1 * wr.Value2
	case Divide:
		if wr.Value2 == 0 {
			resp.Err = errors.New("cannot divide by zero")
			break
		} else {
			resp.Result = wr.Value1 / wr.Value2
		}
	}
	return &resp
}
