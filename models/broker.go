package models

type BrokerFailover struct {
	Delay       string `json:"delay"`
	MaxDelay    string `json:"maxDelay"`
	Failures    int    `json:"failures"`
	FailureTime string `json:"failureTime"`
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
	ID       string          `json:id`
	Mem      float64         `json:"mem"`
	Cpus     float64         `json:"cpus"`
	Heap     float64         `json:"heap"`
	Failover *BrokerFailover `json:"failover"`
	Active   bool            `json:"active"`
	Task     *BrokerTask     `json:"task"`
}
