package marathon

import (
	"github.com/gambol99/go-marathon"
	"log"
	"strings"
)

type Application struct {
	Name      string
	Project   string
	Image     string
	Instances int
	Memory    float64
	Cpu       float64
}

func GetApplications(url string, login string, passwd string) []Application {
	client := buildClient(url, login, passwd)
	apps, err := client.Applications(nil)
	if err != nil {
		log.Fatalf("Failed to get applications from marathon, error: %s", err)
	}
	var result []Application

	for _, a := range apps.Apps {
		id := strings.Split(a.ID, "/")

		result = append(result, Application{
			Name:      id[len(id)-1],
			Project:   id[1],
			Image:     a.Container.Docker.Image,
			Instances: *a.Instances,
			Memory:    *a.Mem,
			Cpu:       a.CPUs,
		})
	}
	return result
}

func buildClient(url string, login string, passwd string) marathon.Marathon {
	config := marathon.NewDefaultConfig()
	config.URL = url
	config.HTTPBasicAuthUser = login
	config.HTTPBasicPassword = passwd
	client, err := marathon.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create a client for marathon, error: %s", err)
	}
	return client
}
