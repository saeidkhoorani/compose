package containers

import (
	"context"
	"io"
)

// Container represents a created container
type Container struct {
	ID          string
	Status      string
	Image       string
	Command     string
	CPUTime     uint64
	MemoryUsage uint64
	MemoryLimit uint64
	PidsCurrent uint64
	PidsLimit   uint64
	Labels      []string
}

// Port represents a published port of a container
type Port struct {
	// Source is the source port
	Source uint32
	// Destination is the destination port
	Destination uint32
}

// ContainerConfig contains the configuration data about a container
type ContainerConfig struct {
	// ID uniquely identifies the container
	ID string
	// Image specifies the iamge reference used for a container
	Image string
	// Ports provide a list of published ports
	Ports []Port
}

// LogsRequest contains configuration about a log request
type LogsRequest struct {
	Follow bool
	Tail   string
	Writer io.Writer
}

// ContainerService interacts with the underlying container backend
type ContainerService interface {
	// List returns all the containers
	List(ctx context.Context) ([]Container, error)
	// Run creates and starts a container
	Run(ctx context.Context, config ContainerConfig) error
	// Exec executes a command inside a running container
	Exec(ctx context.Context, containerName string, command string, reader io.Reader, writer io.Writer) error
	// Logs returns all the logs of a container
	Logs(ctx context.Context, containerName string, request LogsRequest) error
}