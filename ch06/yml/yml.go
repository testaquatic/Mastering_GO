package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

var yamlfile = `
image: Golang
matrix:
  docker: python
  version: [2.7, 3.9]
`

type Mat struct {
	DockerImage string    `yaml:"docker"`
	Version     []float32 `yaml:",flow"`
}

type Yaml struct {
	Image  string
	Matrix Mat
}

func main() {
	var yamlUnmarshal Yaml
	err := yaml.Unmarshal([]byte(yamlfile), &yamlUnmarshal)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("After Unmarshal (Structure):")
	fmt.Println(yamlUnmarshal)

	yamlMarshal, err := yaml.Marshal(&yamlUnmarshal)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("After Marshal (YAML code)")
	fmt.Println(string(yamlMarshal))
}
