// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dockeriface

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
)

// ResponseError is containing a response error
type ResponseError error

// ListContainersResponse encapsulates the response from the docker client for listing containers
type ListContainersResponse struct {
	// DockerIDs is the list of container IDs from the ListContainers call
	DockerIDs []string
	// Error contains any error returned when listing containers
	Error ResponseError
}

// type ContainerChangEvent struct {
// 	// Status represents the container's status in the event
// 	Status apicontainerstatus.ContainerStatus
// 	// DockerContainerMetadata is the metadata of the container in the event
// 	DockerContainerMetadata
// 	// Type is the event type received from docker events
// 	Type apicontainer.DockerEventType
// }

// ContainerMetadata is a type for metadata about Docker containers
// type ContainerMetadata struct {
// 	// DockerID is the contianer's id generated by Docker
// 	DockerID string
// 	// ExitCode contains container's exit code if it has stopped
// 	ExitCode *int
// 	// PortBindings is the list of port binding information of the container
// 	PortBindings []apicontainer.PortBinding
// 	// Error wraps various container transition errors and is set if engine
// 	// is unable to perform any of the required container transitions
// 	// Error apierrors.NamedError
// 	// Volumes contains volume informaton for the container
// 	Volumes []docker.Mount
// 	// Labels contains labels set for the container
// 	Labels map[string]string
// 	// CreatedAt is the timestamp of container creation
// 	CreatedAt time.Time
// 	// StartedAt is the timestamp of container start
// 	StartedAt time.Time
// 	// FinishedAt is the timestamp of container stop
// 	FinishedAt time.Time
// 	// Health contains the result of a container health check
// 	// Health apicontainer.HealthStatus
// }

// DockerContainerChangeEvent is a type for container change events
type DockerContainerChangeEvent struct {
}

// Client interface to implement tests
type Client interface {
	// Version returns the version of the Docker daemon.
	Version(ctx context.Context, timeout time.Duration) (version types.Version, err error)

	// Stop is stopping the client
	Stop() (err error)

	// ListContainers returns the containers known to the daemon
	ListContainers(ctx context.Context, timeout time.Duration, all bool) ListContainersResponse

	// AttachContainerEvents returns a channel to listen to the Docker event stream
	ContainerEvents(ctx context.Context) (<-chan events.Message, <-chan error)
}
