package messaging

import "github.com/nats-io/nats.go"

type NATSClient struct {
	Conn *nats.Conn
}

func NewNATSClient(url string) (*NATSClient, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NATSClient{Conn: conn}, nil
}

func (n *NATSClient) Publish(subject string, message []byte) error {
	return n.Conn.Publish(subject, message)
}

func (n *NATSClient) Close() {
	n.Conn.Close()
}
