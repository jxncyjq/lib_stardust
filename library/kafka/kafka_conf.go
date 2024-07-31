package kafka

type ClientConfig struct {
	Topics   string
	Brokers  []string
	Group    string
	UserName string
	Password string
	mode     string
}
