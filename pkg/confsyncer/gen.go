package confsyncer

import (
	"errors"
	"fmt"
	"github.com/Kuri-su/confSyncer/pkg/unit"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type (
	DCCmd string

	// struct of Docker-compose.yaml
	DockerComposeStruct struct {
		Version  string               `yaml:"version"`
		Services map[string]Container `yaml:"services"`
	}
	Container struct {
		Image   string   `yaml:"image"`
		Restart string   `yaml:"restart"`
		Volumes []string `yaml:"volumes"`
	}
)

var (
	// DC = Docker-Compose
	dcCmd    = DCCmd("")
	dcRunCmd = &cobra.Command{
		Use:   "run",
		Short: "run",
		Run:   dcCmd.Gen,
	}
	dcRestartCmd = &cobra.Command{
		Use:   "restart",
		Short: "restart",
		Run:   dcCmd.Restart,
	}
	dcStopCmd = &cobra.Command{
		Use:   "stop",
		Short: "stop",
		Run:   dcCmd.Stop,
	}
)

const (
	ContainerName      = "confsyncer"
	GenerateDCFileName = "docker-compose.yaml"

	// shell
	restartShell = `cd %s \
&& docker-compose pull \
&& docker-compose up -d `
	stopShell = `
cd %s \
&& docker-compose down
`
)

// ====================================== commands code ====================================
// ====================================== commands code ====================================

func (d *DCCmd) Gen(cmd *cobra.Command, args []string) {
	yamlContent, err := d.buildupDockerComposeYaml()
	if err != nil {
		panic(err)
	}
	println(string(yamlContent))

	genFilePath := fmt.Sprintf("%s/%s", dirPath, GenerateDCFileName)
	err = ioutil.WriteFile(genFilePath, yamlContent, os.FileMode(0544))
	if err != nil {
		panic(err)
	}

	// restart
	d.Restart(cmd, args)
}

func (d *DCCmd) Restart(cmd *cobra.Command, args []string) {
	output, err := unit.RunCommandInShellGetOutput(fmt.Sprintf(restartShell, dirPath))
	if err != nil {
		color.Red(fmt.Sprintf("Run restart command failed! err: %s \n", err.Error()))
	}
	fmt.Println(output)
}

func (d *DCCmd) Stop(cmd *cobra.Command, args []string) {
	output, err := unit.RunCommandInShellGetOutput(fmt.Sprintf(stopShell, dirPath))
	if err != nil {
		color.Red(fmt.Sprintf("Run stop command failed! err: %s \n", err.Error()))
	}
	fmt.Println(output)
}

// ====================================== private code ====================================
// ====================================== private code ====================================

func (d *DCCmd) buildupDockerComposeYaml() ([]byte, error) {

	maps, err := GetFilesMap()
	if err != nil {
		return nil, err
	} else if len(maps) == 0 {
		return nil, errors.New("")
	}

	dc := new(DockerComposeStruct)
	dc.Version = "3"
	dc.Services = make(map[string]Container)

	var volumes []string
	for _, m := range maps {
		containerPath := strings.Replace(m.Local, "~", "/root", 1)
		volumes = append(volumes, fmt.Sprintf("%s:%s", m.Local, containerPath))
	}

	dc.Services[ContainerName] = Container{
		Image:   "kurisux/conf-syncer:latest",
		Restart: "always",
		Volumes: volumes,
	}

	return yaml.Marshal(dc)
}
