package internal

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

func GetApplicationsFromMarathon(marathonUrl string, marathonLogin string, marathonPassword string) []Application {
	client := buildClient(marathonUrl, marathonLogin, marathonPassword)
	apps, err := client.Applications(nil)
	if err != nil {
		log.Fatalf("Failed to get applications from marathon, error: %s", err)
	}
	var result []Application

	for _, application := range apps.Apps {
		id := strings.Split(application.ID, "/")

		result = append(result, Application{
			Name:      id[len(id)-1],
			Project:   id[1],
			Image:     application.Container.Docker.Image,
			Instances: *application.Instances,
			Memory:    *application.Mem,
			Cpu:       application.CPUs,
		})
	}
	return result
}

func buildClient(marathonUrl string, marathonLogin string, marathonPassword string) marathon.Marathon {
	config := marathon.NewDefaultConfig()
	config.URL = marathonUrl
	config.HTTPBasicAuthUser = marathonLogin
	config.HTTPBasicPassword = marathonPassword
	client, err := marathon.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create a client for marathon, error: %s", err)
	}
	return client
}
