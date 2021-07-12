package registry

import "context"

// Registrar is service registrar.
type Registrar interface {
	// Register the registration.
	Register(ctx context.Context, service *ServiceInstance) error
	// Deregister the registration.
	Deregister(ctx context.Context, service *ServiceInstance) error
}

// ServiceInstance is an instance of a service in a discovery system.
type ServiceInstance struct {
	// ID is the unique instance ID as registered.
	ID string `json:"id"`
	// Name is the service name as registered.
	Name string `json:"name"`
	// Version is the version of the compiled.
	Version string `json:"version"`
	// Metadata is the kv pair metadata associated with the service instance.
	Metadata map[string]string `json:"metadata"`
	// Endpoints is endpoint addresses of the service instance.
	// schema:
	//   http://127.0.0.1:8000?isSecure=false
	//   grpc://127.0.0.1:9000?isSecure=false
	Endpoints []string `json:"endpoints"`
}
