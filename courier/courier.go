package courier

type Courier interface {
	send(level string, message string)
}
