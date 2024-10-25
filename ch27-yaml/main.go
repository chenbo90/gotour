package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type People struct {
	Name string            `yaml:"name" json: "name"`
	Age  int               `yaml:"age" json:"age""`
	Work map[string]string `yaml:"work" json:"work`
}

func main() {
	// 待解析数据
	yamlContent := `
field1: 小韩说课
field2:
  field3: value
  field4: [42, 1024]
`
	// 存储解析数据
	result := make(map[string]interface{})
	// 执行解析
	err := yaml.Unmarshal([]byte(yamlContent), &result)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(result)

	//解析文件
	file, err := os.ReadFile("ch27-yaml/hello.yaml")
	fmt.Println(file)

	People1 := &People{}

	err = yaml.Unmarshal(file, &People1)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(People1)
	}

}
