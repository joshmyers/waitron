package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

// Machine configuration
type Machine struct {
	Hostname        string
	OperatingSystem string
	Finish          string
	Preseed         string
	Network         []Interface
	Vars            map[string]string
}

// Interface Configuration
type Interface struct {
	Name       string
	IPAddress  string
	MacAddress string
	Gateway    string
	Netmask    string
}

func machineDefinition(hostname string) (Machine, error) {
	var m Machine
	data, err := ioutil.ReadFile("conf/" + hostname + ".yaml")
	if err != nil {
		return Machine{}, err
	}
	yaml.Unmarshal(data, &m)
	return m, nil
}

func (m Machine) renderTemplate(templateName string) (string, error) {
	var tpl = pongo2.Must(pongo2.FromFile(path.Join("templates", templateName)))
	result, err := tpl.Execute(pongo2.Context{"machine": m})
	if err != nil {
		return "", err
	}
	return result, err
}

func main() {
	m, _ := machineDefinition("example.com")
	template, _ := m.renderTemplate(m.Finish)
	fmt.Println(template)
}