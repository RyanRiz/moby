package client

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

// DockerClient wraps the Docker client with our custom methods
type DockerClient struct {
	client *client.Client
}

// NewDockerClient creates a new Docker client instance
func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &DockerClient{client: cli}, nil
}

// Close closes the Docker client connection
func (dc *DockerClient) Close() error {
	return dc.client.Close()
}

// ListContainers returns a list of containers
func (dc *DockerClient) ListContainers(ctx context.Context, all bool) ([]types.Container, error) {
	return dc.client.ContainerList(ctx, container.ListOptions{All: all})
}

// StartContainer starts a container by ID or name
func (dc *DockerClient) StartContainer(ctx context.Context, containerID string) error {
	return dc.client.ContainerStart(ctx, containerID, container.StartOptions{})
}

// StopContainer stops a container by ID or name
func (dc *DockerClient) StopContainer(ctx context.Context, containerID string) error {
	return dc.client.ContainerStop(ctx, containerID, container.StopOptions{})
}

// InspectContainer returns detailed information about a container
func (dc *DockerClient) InspectContainer(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	return dc.client.ContainerInspect(ctx, containerID)
}

// RemoveContainer removes a container by ID or name
func (dc *DockerClient) RemoveContainer(ctx context.Context, containerID string, force bool) error {
	return dc.client.ContainerRemove(ctx, containerID, container.RemoveOptions{Force: force})
}

// PullImage pulls an image from registry
func (dc *DockerClient) PullImage(ctx context.Context, imageName string) (io.ReadCloser, error) {
	return dc.client.ImagePull(ctx, imageName, image.PullOptions{})
}

// GetServerVersion returns Docker server version info
func (dc *DockerClient) GetServerVersion(ctx context.Context) (types.Version, error) {
	return dc.client.ServerVersion(ctx)
}
