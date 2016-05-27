package types

type BrokerAdd struct {
	BindAddress      string
	Constraints      string
	Cpus             float64
	Expr             string
	FailoverDelay    string
	FailoverMaxDelay string
	FailoverMaxTries string
	Heap             float64
	JvmOptions       string
	Log4jOptions     string
	Mem              float64
	Options          string
	Port             string
	StickinessPeriod string
	Volume           string
}
