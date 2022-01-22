package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func init() {
	file := readFile()
	filterFile(file)
}

var confFileLocal string = "./GoodGoodStudySettings.txt"

// 进程黑名单
var blackList []string

// 黑名单时间
var times [][]string

// cron时间
var cronStr string

func readFile() []string {
	file, err := ioutil.ReadFile(confFileLocal)
	if err != nil {
		fmt.Println("程序出错: 找不到配置文件")
	}
	s := string(file)
	split := strings.Split(s, "\r\n")
	return split
}

func filterFile(strs []string) {
	index := make([]int, 0)
	for i, str := range strs {
		switch str {
		case "[blackList]":
			index = append(index, i)
		case "[time]":
			index = append(index, i)
		case "[trigger]":
			index = append(index, i)
		}
	}

	blackList = getBlackList(readNext(index[0], index[1], strs))
	times = getTime(readNext(index[1], index[2], strs))
	cronStr = getCron(readNext(index[2], len(strs), strs))
	fmt.Println("配置文件加载..成功")

}

func readNext(start int, end int, strs []string) []string {
	res := make([]string, 0)
	start++
	for start < end {
		res = append(res, strs[start])
		start++
	}
	return res
}

func getBlackList(str []string) []string {
	res := make([]string, 0)
	for _, v := range str {
		if v == "" {
			continue
		}
		res = append(res, v)
	}
	return res
}

func getTime(str []string) [][]string {
	res := make([][]string, 0)
	for _, v := range str {
		if v == "" {
			continue
		}
		splits := strings.Split(v, "-")
		res = append(res, splits)
	}
	return res
}

func getCron(str []string) string {
	if len(str) <= 0 || str[0] == "" {
		return "0 /1 * * * ? *"
	}
	return str[0]
}
