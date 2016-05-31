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

type BrokerStickiness struct {
	Hostname string `json:"hostname"`
	Period   string `json:"period"`
}

type Broker struct {
	ID         string            `json:id`
	Mem        float64           `json:"mem"`
	Cpus       float64           `json:"cpus"`
	Heap       float64           `json:"heap"`
	Failover   *BrokerFailover   `json:"failover"`
	Stickiness *BrokerStickiness `json:"stickiness"`
	Active     bool              `json:"active"`
	Task       *BrokerTask       `json:"task"`
	Port       string            `json:"port"`
}
