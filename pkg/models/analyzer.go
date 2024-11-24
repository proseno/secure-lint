package models

import (
	"fmt"
	"github.com/fatih/color"
	"os/exec"
)

type Analyzer struct {
	Command        string `yaml:"command"`
	Flags          string `yaml:"flags"`
	InstallCommand string `yaml:"install_command"`
}

func (a *Analyzer) GetCommand() string {
	if a.Flags != "" {
		return a.Command + " " + a.Flags
	}
	return a.Command
}

func (a *Analyzer) Install() ([]byte, error) {
	installCmd := exec.Command(a.InstallCommand)
	return installCmd.Output()
}

func (a *Analyzer) InstallWithOutput() bool {
	color.Red("No %s executable found in $PATH\n", a.Command)
	color.Yellow("Installing %s\n", a.Command)
	color.Yellow("Running %s\n", a.InstallCommand)

	stdout, err := a.Install()

	if err != nil {
		color.Red("Failed to install %s\n", a.Command)
		color.Red("Error %s\n", err.Error())
		return false
	}
	fmt.Println(string(stdout))
	color.Green("%s installed. Continue analyzing\n", a.Command)
	return true
}

func (a *Analyzer) CheckExecutable() bool {
	_, err := exec.LookPath(a.Command)
	success := true
	if err != nil {
		success = a.InstallWithOutput()
	}
	return success
}

func (a *Analyzer) Analyze(path string) string {
	cmd := exec.Command(a.GetCommand(), path)
	cmd.Stdout = nil
	//ignore err
	//in case stdout have analysis then err isn't empty
	fmt.Printf("Executing %s\n", cmd.String())
	stdout, _ := cmd.Output()

	return string(stdout)
}
