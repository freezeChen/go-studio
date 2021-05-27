package transport

const (
	TypeGRPC = "grpc"
	TypeHTTP = "http"
)

type Server interface {
	Endpoint() (string, string, error)
	Start() error
	Stop() error
}
