package applications

import (
	"github.com/urfave/cli/v2"
	"marathon-explorer/marathon"
	"strconv"
	"strings"
)

var headers = []string{"Name", "Project", "Image", "Instances", "CPU", "Memory"}

func GetApplications(context *cli.Context) [][]string {
	applications := marathon.GetApplications(context.String("url"), context.String("user"), context.String("password"))
	if context.IsSet("project") {
		applications = applyProjectFilter(&applications, context.String("project"))
	}
	if context.IsSet("instances") {
		applications = applyInstancesFilter(&applications, context.Int("instances"))
	}
	if context.IsSet("image") {
		applications = applyImageFilter(&applications, context.String("image"))
	}

	result := [][]string{headers}
	for _, application := range applications {
		result = append(result, []string{
			application.Name,
			application.Project,
			application.Image,
			strconv.Itoa(application.Instances),
			strconv.FormatFloat(application.Cpu, 'f', 1, 64),
			strconv.FormatFloat(application.Memory, 'f', 1, 64),
		})
	}
	return result
}

func applyProjectFilter(applications *[]marathon.Application, project string) []marathon.Application {
	var result []marathon.Application
	for _, application := range *applications {
		if application.Project == project {
			result = append(result, application)
		}
	}
	return result
}

func applyInstancesFilter(applications *[]marathon.Application, instances int) []marathon.Application {
	var result []marathon.Application
	for _, application := range *applications {
		if application.Instances == instances {
			result = append(result, application)
		}
	}
	return result
}

func applyImageFilter(applications *[]marathon.Application, image string) []marathon.Application {
	var result []marathon.Application
	for _, application := range *applications {
		if strings.Contains(application.Image, image) {
			result = append(result, application)
		}
	}
	return result
}
