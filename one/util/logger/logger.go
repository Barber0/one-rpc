package logger

import (
	"fmt"
	"log"
)

const (
	DEBUG	byte	=	iota
	INFO
	WARN
	ERROR
)

var (
	loggerMap 	=	make(map[string]OneLogger)
	levelMap	=	map[byte]string{
		DEBUG:"DEBUG",
		INFO:"INFO",
		WARN:"WARN",
		ERROR:"ERROR",
	}
)

type OneLogger struct {
	name	string
}

func GetOneLogger(name string) *OneLogger {
	var (
		ol OneLogger
		ok bool
	)
	if ol, ok = loggerMap[name]; !ok {
		ol = OneLogger{name:name}
		loggerMap[name] = ol
	}
	return &ol
}

func (l *OneLogger) println(level byte, msg ...interface{}) {
	log.Println(append([]interface{}{"|"+l.name+"|"+levelMap[level]+"|"},msg...)...)
}

func (l *OneLogger) printf(level byte, tpl string, args ...interface{}) {
	log.Printf("|%s|%s| "+tpl+"\n",append([]interface{}{l.name,levelMap[level]},args...)...)
}

func (l *OneLogger) Debug(msg ...interface{}) {
	fmt.Println(loggerMap)
	l.println(DEBUG,msg...)
}

func (l *OneLogger) Info(msg ...interface{}) {
	l.println(INFO,msg...)
}

func (l *OneLogger) Warn(msg ...interface{}) {
	l.println(WARN,msg...)
}

func (l *OneLogger) Error(msg ...interface{}) {
	l.println(ERROR,msg...)
}

func (l *OneLogger) Debugf(tpl string, args ...interface{}) {
	l.printf(DEBUG,tpl,args...)
}

func (l *OneLogger) Infof(tpl string, args ...interface{}) {
	l.printf(INFO,tpl,args...)
}

func (l *OneLogger) Warnf(tpl string, args ...interface{}) {
	l.printf(WARN,tpl,args...)
}

func (l *OneLogger) Errorf(tpl string, args ...interface{}) {
	l.printf(ERROR,tpl,args...)
}