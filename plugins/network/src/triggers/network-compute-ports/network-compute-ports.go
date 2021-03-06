package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dokku/dokku/plugins/common"
	"github.com/dokku/dokku/plugins/config"
)

// computes the ports for a given app container
func main() {
	flag.Parse()
	appName := flag.Arg(0)
	procType := flag.Arg(1)
	isHerokuishContainer := common.ToBool(flag.Arg(2))

	if procType != "web" {
		return
	}

	var dockerfilePorts []string
	if !isHerokuishContainer {
		dockerfilePorts = strings.Split(config.GetWithDefault(appName, "DOKKU_DOCKERFILE_PORTS", ""), " ")
	}

	var ports []string
	if len(dockerfilePorts) == 0 {
		ports = append(ports, "5000")
	} else {
		for _, port := range dockerfilePorts {
			port = strings.TrimSuffix(strings.TrimSpace(port), "/tcp")
			if port == "" || strings.HasSuffix(port, "/udp") {
				continue
			}
			ports = append(ports, port)
		}
	}
	fmt.Fprint(os.Stdout, strings.Join(ports, " "))
}
