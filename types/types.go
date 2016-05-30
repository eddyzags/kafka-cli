package types

type BrokerAdd struct {
	BindAddress      string
	Constraints      string
	Cpus             string
	Expr             string
	FailoverDelay    string
	FailoverMaxDelay string
	FailoverMaxTries string
	Heap             string
	JvmOptions       string
	Log4jOptions     string
	Mem              string
	Options          string
	Port             string
	StickinessPeriod string
	Volume           string
}
