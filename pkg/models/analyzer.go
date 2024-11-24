package models

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

type Analyzer struct {
	Command        string `yaml:"command"`
	Flags          string `yaml:"flags"`
	OutputFlag     string `yaml:"output_flag"`
	Stdout         string `yaml:"stdout"`
	Level          string `yaml:"level"`
	InstallCommand string `yaml:"install_command"`
	As             string `yaml:"as"`
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

func recreateFile(filename string) {
	if _, err := os.Stat(filename); err == nil {
		// File exists, delete it
		err = os.Remove(filename)
		if err != nil {
			fmt.Printf("Error deleting file: %v\n", err)
			return
		}
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Printf("Error checking file: %v\n", err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
}

func (a *Analyzer) getOutputFlag() string {
	if a.OutputFlag != "" {
		var outputFile = ProjectRoot + "/output/" + a.As + ".txt"
		recreateFile(outputFile)
		return a.OutputFlag + "=" + outputFile
	}
	return a.OutputFlag
}

func (a *Analyzer) PrepareCommandArgs(path string) []string {
	var result []string
	args := []string{a.Flags, a.getOutputFlag(), a.Level, path}
	for _, arg := range args {
		if arg != "" {
			result = append(result, arg)
		}
	}
	return result
}

func (a *Analyzer) Analyze(path string) string {
	cmd := exec.Command(a.Command, a.PrepareCommandArgs(path)...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	fmt.Printf("Executing %s\n", cmd.String())
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\nStderr: %s\n", err, stderr.String())
		return ""
	}
	return string(stdout)
}
