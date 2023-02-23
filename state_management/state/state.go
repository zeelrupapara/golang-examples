package state

const (
	Add      = "add"
	Subtract = "subtract"
	Multiply = "multiply"
	Divide   = "divide"
)

type WorkRequest struct {
	Op             string
	Value1, Value2 int
}

type WorkResponce struct {
	Wr     *WorkRequest
	Result int
	Err    error
}
