package metrics

import (
	"net/http"
	"time"

	metric "github.com/rcrowley/go-metrics"
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	c := metric.GetOrRegisterCounter("counterhandler.counter", nil)
	c.Inc(1)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func TimerHandler(w http.ResponseWriter, r *http.Request) {
	currt := time.Now()
	t := metric.GetOrRegisterTimer("timerhandler.timer", nil)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
	t.UpdateSince(currt)
}
