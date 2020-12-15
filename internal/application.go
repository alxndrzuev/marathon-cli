package internal

import (
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
)

var headers = []string{"Name", "Project", "Image", "Instances", "CPU", "Memory"}

func GetApplications(context *cli.Context) [][]string {
	applications := GetApplicationsFromMarathon(
		context.String("url"),
		context.String("user"),
		context.String("password"),
	)
	applications = filter(applications, func(application Application) bool {
		return !context.IsSet("project") || application.Project == context.String("project")
	})
	applications = filter(applications, func(application Application) bool {
		return !context.IsSet("image") || strings.Contains(application.Image, context.String("image"))
	})
	applications = filter(applications, func(application Application) bool {
		return !context.IsSet("instances") || application.Instances == context.Int("instances")
	})

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

func filter(applications []Application, f func(Application) bool) []Application {
	result := make([]Application, 0)
	for _, application := range applications {
		if f(application) {
			result = append(result, application)
		}
	}
	return result
}
