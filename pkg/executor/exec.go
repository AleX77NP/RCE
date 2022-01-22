package executor

import (
	dc "rce.amopdev/m/v2/pkg/dockerclient"
)

type Executor struct {
	language string
	client *dc.DockerClient
}

func NewExecutor(language, code, file, hostDir string) *Executor {
	image := selectImage(language)
	command := selectCommand(language, file)
	return &Executor{
		language: language,
		client: dc.NewDockerClient(image, command, file, hostDir),
	}
}

func (executor *Executor) ExecuteCode() (error, string) {
	return executor.client.RunContainer()
}

// Make these things dynamic
func selectImage(language string) string{
	return "docker.io/library/python"
}

func selectCommand(language, file string) []string {
	return []string{"python", file}
}