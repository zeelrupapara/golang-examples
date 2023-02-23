package pool

import (
	"golang.org/x/crypto/bcrypt"
)

func HashWork(wr WorkRequest) WorkResponce {
	var result []byte
	var err error
	result, err = bcrypt.GenerateFromPassword(wr.Text, bcrypt.DefaultCost)
	return WorkResponce{Wr: wr, Result: result, Err: err}
}

func CompareWork(wr WorkRequest) WorkResponce {
	var matched bool
	err := bcrypt.CompareHashAndPassword(wr.Compare, wr.Text)
	if err == nil {
		matched = true
	}
	return WorkResponce{Wr: wr, Matched: matched, Err: err}
}
