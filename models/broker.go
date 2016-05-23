package models

type BrokerFailover struct {
	Delay       int64 `json:"delay"`
	MaxDelay    int64 `json:"maxDelay"`
	Failures    int   `json:"failures"`
	FailureTime int64 `json:"failureTime"`
}

type BrokerTask struct {
	//TODO(eddyzags): Attributes field
	ExecutorID string `json:"executorId"`
	SlaveID    string `json:"slaveId"`
	Hostname   string `json:"hostname"`
	ID         string `json:"id"`
	Endpoint   string `json:"endpoint"`
	State      string `json:"state"`
}

type Broker struct {
	ID       int64           `json:id`
	Mem      float64         `json:"mem"`
	Cpus     float64         `json:"cpus"`
	Heap     float64         `json:"heap"`
	Failover *BrokerFailover `json:"failover"`
	Active   bool            `json:"active"`
	Task     *BrokerTask     `json:"task"`
}
