package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

// SSHExecutor 封装SSH命令执行
type SSHExecutor struct {
	timeout time.Duration
	port    int // 添加端口字段
}

// NewSSHExecutor 创建新的SSH执行器
func NewSSHExecutor(timeout time.Duration, port int) *SSHExecutor {
	if port <= 0 {
		port = 22 // 默认SSH端口
	}
	return &SSHExecutor{
		timeout: timeout,
		port:    port,
	}
}

// ExecuteSSHCommand 执行SSH命令并提供完整的错误处理
func (e *SSHExecutor) ExecuteSSHCommand(host string, command string) error {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	// 构建SSH命令参数，添加端口
	args := []string{
		"-p", fmt.Sprintf("%d", e.port), // 添加端口参数
		host,
		command,
	}

	// 也可以添加其他SSH选项
	sshOptions := []string{
		"-o", "ConnectTimeout=10", // 连接超时时间
		"-o", "StrictHostKeyChecking=no", // 关闭主机密钥检查（注意：在生产环境中要谨慎使用）
	}

	// 合并所有参数
	args = append(sshOptions, args...)

	cmd := exec.CommandContext(ctx, "ssh", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("启动命令失败: %v", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		if err := cmd.Process.Kill(); err != nil {
			return fmt.Errorf("命令超时且无法杀死进程: %v", err)
		}
		return fmt.Errorf("命令执行超时")
	case err := <-done:
		if err != nil {
			return fmt.Errorf("命令执行失败: %v\n错误输出: %s", err, stderr.String())
		}
	}

	if !cmd.ProcessState.Success() {
		return fmt.Errorf("命令执行失败，退出状态非0\n错误输出: %s", stderr.String())
	}

	return nil
}

// 使用示例
func main() {
	// 创建一个SSH执行器，设置30秒超时和自定义端口（例如2222）
	executor := NewSSHExecutor(30*time.Second, 22022)

	// 使用用户名@主机的形式
	err := executor.ExecuteSSHCommand("root@10.19.90.15", "ll")
	if err != nil {
		fmt.Printf("SSH命令执行失败1: %v\n", err)
		return
	}

	fmt.Println("SSH命令执行成功")

	// 如果要使用默认端口22
	executor2 := NewSSHExecutor(30*time.Second, 0) // 传入0会使用默认端口22
	err = executor2.ExecuteSSHCommand("root@10.19.90.15", "ls -l")
	if err != nil {
		fmt.Printf("SSH命令执行失败2: %v\n", err)
		return
	}
}
