package types

type BrokerAdd struct {
	Expr     string
	Cpus     float64
	Mem      float64
	Heap     float64
	Delay    string
	MaxDelay string
}
