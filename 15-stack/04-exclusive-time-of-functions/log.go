package main

import (
	"strconv"
	"strings"
)

type Log struct {
	id      int
	isStart bool
	time    int
}

func (this *Log) LogInit(content string) {
    strs := strings.Split(content, ":")
    this.id, _ = strconv.Atoi(strs[0])
    this.isStart = strs[1] == "start"
    this.time, _ = strconv.Atoi(strs[2])
}

func LogInitWithString(content string) *Log {
    log := new(Log)
    log.LogInit(content)
    return log
}
