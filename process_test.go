package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestFidAll(t *testing.T) {
	findAllPid()
}

func TestNew(t *testing.T) {
	s := "123213"[3:]
	println(s)
}

func TestReadFile(t *testing.T) {
	file := readFile()
	filterFile(file)
}

func TestTime(t *testing.T) {

	clock, min := 20, 00
	mtime := NewMytimeByInt(clock, min)
	for _, v := range times {
		start := NewMytime(v[0])
		end := NewMytime(v[1])
		if mtime.inner(*start, *end) {
			println(true)
			return
		}
	}
	println(false)
}

func TestKill(t *testing.T) {
	//findAllPid()
	str := "Fallout4.exe"
	command := exec.Command("TASKKILL", "/F", "/T", "/IM", str)
	output, err := command.Output()
	if err != nil {
	}
	fmt.Println(string(output))
}
