package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"strings"
	"time"
)

func main() {
	c := newWithSeconds()
	c.AddFunc(cronStr, func() {
		checkPid()
	})
	c.Start()
	select {}
}

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func checkPid() {
	clock, min, sec := time.Now().Clock()
	fmt.Printf("%d:%d:%d 开始执行", clock, min, sec)
	if !isTime() {
		return
	}
	process := findAllPid()
	canDels := make([]string, 0)
	for _, v := range process {
		for _, v2 := range blackList {
			if v == v2 {
				canDels = append(canDels, v)
			}
		}
	}
	fmt.Printf("需要关闭的程序有：")
	for _, r := range canDels {
		fmt.Print(r, ", ")
	}
	fmt.Println()
	for _, del := range canDels {
		b := killProcess(del)
		if b {
			fmt.Println(del, " 关闭成功")
		} else {
			fmt.Println(del, " 关闭失败")
		}
	}
}

func isTime() bool {
	clock, min, _ := time.Now().Clock()
	mtime := NewMytimeByInt(clock, min)
	for _, v := range times {
		start := NewMytime(v[0])
		end := NewMytime(v[1])
		if mtime.inner(*start, *end) {
			return true
		}
	}
	return false
}

type mytime struct {
	hour int
	min  int
}

func NewMytime(str string) *mytime {
	split := strings.Split(str, ":")
	m := new(mytime)
	if len(split) < 2 {
		return m
	}
	m.hour = str2Int(split[0])
	m.min = str2Int(split[1])
	return m
}

func NewMytimeByInt(hour int, min int) *mytime {
	m := new(mytime)
	m.hour = hour
	m.min = min
	return m
}

func (this *mytime) inner(start mytime, end mytime) bool {
	return this.bigger(start) && end.bigger(*this)
}

func (this *mytime) bigger(target mytime) bool {
	if this.hour > target.hour {
		return true
	} else if this.hour == target.hour {
		return this.min >= target.min
	}
	return false
}

func str2Int(str string) int {
	sum := 0
	for _, v := range str {
		sum = sum*10 + int(v-'0')
	}
	return sum
}
