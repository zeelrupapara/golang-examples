package pool

import "fmt"

const (
	Encrypt = "encrypt"
	Decrypt = "decrypt"
)

type WorkRequest struct {
	Operation string
	Text      []byte
	Compare   []byte
}

type WorkResponce struct {
	Wr      WorkRequest
	Result  []byte
	Matched bool
	Err     error
}

func Process(wr WorkRequest) WorkResponce {
	switch wr.Operation {
	case Encrypt:
		return HashWork(wr)
	case Decrypt:
		return CompareWork(wr)
	default:
		return WorkResponce{Err: fmt.Errorf("unknown operation: %s", wr.Operation)}
	}
}
