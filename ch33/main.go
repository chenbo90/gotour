package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func parseYAMLFiles(path string) ([]string, error) {
	// 判断path路径下是否存在.files目录
	_, err1 := os.Stat(filepath.Join(path, ".files"))
	if err1 == nil {
		fmt.Println("files directory already exists")
		// 删除.files 目录
		os.RemoveAll(filepath.Join(path, ".files"))
	}
	// 创建.files目录
	fmt.Println("Creating .files directory")
	filesDir := filepath.Join(path, ".files")
	err := os.MkdirAll(filesDir, 0755)
	if err != nil {
		return nil, err
	}

	// 读取目录下的所有 YAML 文件
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var outputFiles []string

	// 生成字母序列
	letters := make([]string, 26)
	for i := 0; i < 26; i++ {
		letters[i] = string(rune('z' - i))
	}

	// 遍历文件
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") { // 跳过隐藏文件
			continue
		}
		if strings.HasSuffix(file.Name(), ".yaml") {
			data, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
			if err != nil {
				return nil, err
			}

			// 以 "---" 分割多个文档
			docs := strings.Split(string(data), "---")
			letterIndex := 0 // 用于追踪当前使用的字母索引

			for _, doc := range docs {
				doc = strings.TrimSpace(doc)
				if doc == "" {
					continue // 跳过空文档
				}

				// 生成新的文件名
				var outputFileName string
				if len(docs) > 1 {
					if letterIndex >= len(letters) {
						return nil, fmt.Errorf("too many documents in file %s: maximum is 26", file.Name())
					}
					baseName := strings.TrimSuffix(file.Name(), ".yaml")
					outputFileName = fmt.Sprintf("%s-%s.yaml", baseName, letters[letterIndex])
					letterIndex++
				} else {
					outputFileName = file.Name() // 直接使用原文件名
				}

				// 将文件写入 files 目录
				err := ioutil.WriteFile(filepath.Join(filesDir, outputFileName), []byte(doc), 0644)
				if err != nil {
					return nil, err
				}
				outputFiles = append(outputFiles, outputFileName)
			}
		}
	}

	// 将文件名倒序排序
	sort.Sort(sort.Reverse(sort.StringSlice(outputFiles)))

	return outputFiles, nil
}

func parseYAMLFiles1(path string) ([]string, error) {
	//判断path路径下是否存在。files目录
	_, err1 := os.Stat(filepath.Join(path, ".files"))
	if err1 == nil {
		fmt.Println("files directory already exists")
		//删除.files 目录
		os.RemoveAll(filepath.Join(path, ".files"))
	}
	//创建.files目录
	fmt.Println("Creating .files directory ")
	filesDir := filepath.Join(path, ".files")
	err := os.MkdirAll(filesDir, 0755)
	if err != nil {
		return nil, err
	}

	// 读取目录下的所有 YAML 文件
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var outputFiles []string

	// 遍历文件
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") { // 跳过隐藏文件
			continue
		}
		if strings.HasSuffix(file.Name(), ".yaml") {
			data, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
			if err != nil {
				return nil, err
			}

			// 以 "---" 分割多个文档
			docs := strings.Split(string(data), "---")
			for i, doc := range docs {
				doc = strings.TrimSpace(doc)
				if doc == "" {
					continue // 跳过空文档
				}

				// 生成新的文件名
				var outputFileName string
				if len(docs) > 1 {
					outputFileName = fmt.Sprintf("%s%d.yaml", strings.TrimSuffix(file.Name(), ".yaml"), i+1)
				} else {
					outputFileName = file.Name() // 直接使用原文件名
				}

				// 将文件写入 files 目录
				err := ioutil.WriteFile(filepath.Join(filesDir, outputFileName), []byte(doc), 0644)
				if err != nil {
					return nil, err
				}
				outputFiles = append(outputFiles, outputFileName)
			}
		}
	}

	// 将文件名倒序排序
	sort.Sort(sort.Reverse(sort.StringSlice(outputFiles)))

	return outputFiles, nil
}

func main() {
	path := "/Users/pipi/haha" // 替换为你的目录路径
	files, err := parseYAMLFiles(path)
	if err != nil {
		log.Fatalf("Error parsing YAML files: %v", err)
	}

	// 打印最终结果
	for _, file := range files {
		fmt.Println(file)
	}
}
