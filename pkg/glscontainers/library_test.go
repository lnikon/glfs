package glscontainers

import "testing" 

func TestRunContainers(t *testing.T) {
	url := "registry.hub.docker.com/library/alpine"
	RunContainer(url)
}

func TestListContainers(t *testing.T) {
	ListContainers()
}
