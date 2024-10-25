package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

// SSHConfig 存储SSH连接配置
type SSHConfig struct {
	Host    string
	Port    int
	User    string        // 如果为空则使用当前用户
	Timeout time.Duration // 连接超时时间
}

// Result 存储命令执行结果
type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
	Error    error
}

// getDefaultKeyFiles 获取默认的SSH私钥文件列表
func getDefaultKeyFiles() []string {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	// 常见的私钥文件名
	keyFiles := []string{
		"id_rsa",
		"id_ed25519",
		"id_ecdsa",
		"id_dsa",
	}

	var paths []string
	for _, file := range keyFiles {
		path := filepath.Join(home, ".ssh", file)
		if _, err := os.Stat(path); err == nil {
			paths = append(paths, path)
		}
	}

	return paths
}

// GetCurrentUser 获取当前用户名，返回用户名和可能的错误
func GetCurrentUser() (string, error) {
	// 方法1: 使用 user.Current()
	if currentUser, err := user.Current(); err == nil && currentUser.Username != "" {
		// 在 macOS 中，Current().Username 可能返回形如 "name@domain"的格式
		// 我们只需要用户名部分
		username := currentUser.Username
		if i := strings.IndexByte(username, '@'); i != -1 {
			username = username[:i]
		}
		return username, nil
	}

	// 方法2: 使用环境变量
	if username := os.Getenv("USER"); username != "" {
		return username, nil
	}

	// 方法3: 使用 LOGNAME 环境变量
	if username := os.Getenv("LOGNAME"); username != "" {
		return username, nil
	}

	// 如果所有方法都失败了
	return "", fmt.Errorf("无法获取当前用户名")
}

// loadPrivateKey 加载私钥文件
func loadPrivateKey(file string) (ssh.Signer, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// 首先尝试无密码加载
	key, err := ssh.ParsePrivateKey(buffer)
	if err == nil {
		return key, nil
	}

	return nil, fmt.Errorf("无法加载私钥 %s: %v", file, err)
}

// ExecuteCommand 执行远程命令并返回结果
func ExecuteCommand(config SSHConfig, command string) Result {
	result := Result{}

	// 如果用户名为空，使用当前用户
	if config.User == "" {
		//currentUser, _ := GetCurrentUser()
		config.User = "root"
		if config.User == "" {
			result.Error = fmt.Errorf("无法获取当前用户名")
			return result
		}
	}

	// 获取默认的私钥文件
	keyFiles := getDefaultKeyFiles()
	if len(keyFiles) == 0 {
		result.Error = fmt.Errorf("在 ~/.ssh 目录下未找到私钥文件")
		return result
	}

	// 尝试加载所有可用的私钥
	var signers []ssh.Signer
	for _, keyFile := range keyFiles {
		if signer, err := loadPrivateKey(keyFile); err == nil {
			signers = append(signers, signer)
		}
	}

	if len(signers) == 0 {
		result.Error = fmt.Errorf("没有可用的私钥")
		return result
	}

	// 创建SSH客户端配置
	clientConfig := &ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signers...),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 注意：生产环境应该验证主机密钥
		Timeout:         config.Timeout,
	}

	// 建立SSH连接
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		result.Error = fmt.Errorf("SSH连接失败: %v", err)
		return result
	}
	defer client.Close()

	// 创建新会话
	session, err := client.NewSession()
	if err != nil {
		result.Error = fmt.Errorf("创建会话失败: %v", err)
		return result
	}
	defer session.Close()

	// 获取标准输出和标准错误输出
	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	// 执行命令
	err = session.Run(command)
	result.Stdout = stdout.String()
	result.Stderr = stderr.String()

	// 处理命令执行结果j
	if err != nil {
		if exitErr, ok := err.(*ssh.ExitError); ok {
			result.ExitCode = exitErr.ExitStatus()
		}
		result.Error = fmt.Errorf("命令执行失败: %v", err)
	} else {
		result.ExitCode = 0
	}

	return result
}

// 使用示例
func main() {
	// 最简配置，自动使用当前用户和默认私钥
	config := SSHConfig{
		Host:    "10.19.90.15",
		Port:    22022,
		Timeout: 30 * time.Second,
	}

	// 执行命令
	result := ExecuteCommand(config, "ll")

	// 检查执行结果
	if result.Error != nil {
		fmt.Printf("错误: %v\n", result.Error)
		fmt.Printf("退出码: %d\n", result.ExitCode)
		fmt.Printf("标准错误: %s\n", result.Stderr)
	} else {
		fmt.Printf("命令执行成功:\n%s", result.Stdout)
	}
}
