package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func findAllPid() []string {
	cmd := exec.Command("cmd", "/C", "tasklist")
	output, _ := cmd.Output()
	// output 是类似下面这样的字符串的byte，从System开始的字符串才是有用的
	// 映像名称                       PID 会话名              会话#       内存使用
	//========================= ======== ================ =========== ============
	//System Idle Process              0 Services                   0          8 K
	//System                           4 Services                   0      2,028 K

	// 获取System出现的位置
	n := strings.Index(string(output), "System")
	if n == -1 {
		fmt.Println("no find")
		os.Exit(1)
	}

	// 切割字符串接下来都是 System 4 Services 0 2,028 Kddde 一行一行的数据
	data := string(output)[n:]
	fields := strings.Fields(data)
	return fields
}

func killProcess(appName string) bool {
	command := exec.Command("TASKKILL", "/F", "/T", "/IM", appName)
	output, err := command.Output()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(output))
	return true
}
