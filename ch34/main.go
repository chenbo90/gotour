package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

// SSHConfig 存储SSH连接配置
type SSHConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	PrivateKey string        // 私钥文件路径
	Timeout    time.Duration // 连接超时时间
}

// Result 存储命令执行结果
type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
	Error    error
}

// ExecuteCommand 执行远程命令并返回结果
func ExecuteCommand(config SSHConfig, command string) Result {
	result := Result{}

	// 创建SSH客户端配置
	clientConfig := &ssh.ClientConfig{
		User:            config.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 注意：生产环境应该验证主机密钥
		Timeout:         config.Timeout,
	}

	// 设置认证方式
	//if config.Password != "" {
	//	clientConfig.Auth = []ssh.AuthMethod{
	//		ssh.Password(config.Password),
	//	}
	//} else if config.PrivateKey != "" {
	//	key, err := ioutil.ReadFile(config.PrivateKey)
	//	if err != nil {
	//		result.Error = fmt.Errorf("读取私钥失败: %v", err)
	//		return result
	//	}
	//
	//	signer, err := ssh.ParsePrivateKey(key)
	//	if err != nil {
	//		result.Error = fmt.Errorf("解析私钥失败: %v", err)
	//		return result
	//	}
	//
	//	clientConfig.Auth = []ssh.AuthMethod{
	//		ssh.PublicKeys(signer),
	//	}
	//} else {
	//	result.Error = fmt.Errorf("未提供认证方式")
	//	return result
	//}

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

	// 处理命令执行结果
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
func ExampleUsage() {
	// 使用密码认证
	config := SSHConfig{
		Host:     "remote.example.com",
		Port:     22,
		Username: "user",
		Password: "password",
		Timeout:  30 * time.Second,
	}

	// 使用私钥认证
	//configWithKey := SSHConfig{
	//	Host:       "remote.example.com",
	//	Port:       22,
	//	Username:   "user",
	//	PrivateKey: "/path/to/private/key",
	//	Timeout:    30 * time.Second,
	//}

	// 执行命令
	result := ExecuteCommand(config, "ls -la")

	// 检查执行结果
	if result.Error != nil {
		fmt.Printf("错误: %v\n", result.Error)
		fmt.Printf("退出码: %d\n", result.ExitCode)
		fmt.Printf("标准错误: %s\n", result.Stderr)
	} else {
		fmt.Printf("命令执行成功:\n%s", result.Stdout)
	}

}

func main() {
	config := SSHConfig{
		Host:     "10.19.90.15",
		Port:     22022,
		Username: "root",
		//Password: "xdjr0lxGu",
		//PrivateKey: "/Users/pipi/.ssh/id_rsa",
		Timeout: 10 * time.Second,
	}

	// 执行单个命令
	result := ExecuteCommand(config, "ll")

	if result.Error != nil {
		fmt.Printf("执行命令出错: %v\n", result.Error)
		fmt.Printf("错误输出: %s\n", result.Stderr)
		return
	}

	fmt.Printf("命令输出:\n%s\n", result.Stdout)
	fmt.Printf("退出码: %d\n", result.ExitCode)
}
