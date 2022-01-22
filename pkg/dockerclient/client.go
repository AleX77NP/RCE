package dockerclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

var (
	buf bytes.Buffer
)

type DockerClient struct {
	ID string
	image string
	command []string
	file string
	hostDir string
}

func NewDockerClient(image string, command []string, file, hostDir string) *DockerClient{
	return &DockerClient{
		ID: "",
		image: image,
		command: command,
		hostDir: hostDir,
	}
}

// Make things dynamic - image 
func (dockerClient *DockerClient) RunContainer() (error, string) {
	ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return err, err.Error()
    }

    reader, err := cli.ImagePull(ctx, dockerClient.image, types.ImagePullOptions{})
    if err != nil {
        return err, err.Error()
    }
    io.Copy(os.Stdout, reader)

	mount := dockerClient.hostDir + ":" + "/code"

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "python",
        Cmd:   dockerClient.command,
    }, 
	&container.HostConfig{
        Binds: []string{
            mount,
        },
    },
	 nil, nil, "")
    
	if err != nil {
        return err, err.Error()
    }

	dockerClient.ID = resp.ID

	fmt.Printf("Start Container...")
    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err, err.Error()
    }

    statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
    select {
    case err := <-errCh:
        if err != nil {
            return err, err.Error()
        }
    case <-statusCh:
    }

    out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
    if err != nil {
        return err, err.Error()
    }

	r := dockerClient.generateResponse(out)

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	// remove container here using ID

	return nil, r
}

func (dockerClient *DockerClient) generateResponse(out io.ReadCloser) string{
	b := new(bytes.Buffer)
	b.ReadFrom(out)
	response := b.String()
	return response
}