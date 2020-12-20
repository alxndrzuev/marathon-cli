package command

import (
	"github.com/urfave/cli/v2"
	"marathon-explorer/internal/marathon"
	"strconv"
	"strings"
)

var headers = []string{"Name", "Project", "Image", "Instances", "CPU", "Memory"}

func GetApplications(ctx *cli.Context) [][]string {
	apps := marathon.GetApplications(
		ctx.String("url"),
		ctx.String("user"),
		ctx.String("password"),
	)
	apps = filter(apps, func(a marathon.Application) bool {
		return !ctx.IsSet("project") || contains(ctx.StringSlice("project"), a.Project)
	})
	apps = filter(apps, func(a marathon.Application) bool {
		return !ctx.IsSet("image") || strings.Contains(a.Image, ctx.String("image"))
	})
	apps = filter(apps, func(a marathon.Application) bool {
		return !ctx.IsSet("instances") || a.Instances == ctx.Int("instances")
	})

	result := [][]string{headers}
	for _, a := range apps {
		result = append(result, []string{
			a.Name,
			a.Project,
			a.Image,
			strconv.Itoa(a.Instances),
			strconv.FormatFloat(a.Cpu, 'f', 1, 64),
			strconv.FormatFloat(a.Memory, 'f', 1, 64),
		})
	}
	return result
}

func filter(aa []marathon.Application, f func(marathon.Application) bool) []marathon.Application {
	result := make([]marathon.Application, 0)
	for _, a := range aa {
		if f(a) {
			result = append(result, a)
		}
	}
	return result
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
