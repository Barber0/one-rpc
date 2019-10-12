package transport

import (
	"context"
	"encoding/binary"
	"log"
	"net"
	"runtime/debug"
)

const iMaxLen = 10485760

var OLOG Logger

type SvrProtocol interface {
	Invoke(ctx context.Context, req []byte) []byte
	ParsePkg(pkg []byte) (int,int)
	InvokeTimeout(ctx context.Context, req []byte) []byte
}

type CltProtocol interface {
	Recv(pkg []byte)
	ParsePkg(pkg []byte) (int,int)
}

type Logger interface {
	Debug(msg...interface{})
	Info(msg...interface{})
	Warn(msg...interface{})
	Error(msg...interface{})

	Debugf(tpl string, args...interface{})
	Infof(tpl string, args...interface{})
	Warnf(tpl string, args...interface{})
	Errorf(tpl string, args...interface{})
}

type task func()
type GPool interface {
	AddTask(t task)
	Release()
}

func NewGPool(maxInvoke, queueCap int) GPool {
	return nil
}

const (
	PKG_LESS	=	iota
	PKG_FULL
	PKG_ERR
)

func ParsePkg(req []byte) (pkgLen, status int) {
	reqLen := len(req)
	if reqLen < 4 {
		status = PKG_LESS
		return
	}
	iHeaderLen := int(binary.BigEndian.Uint32(req[:4]))
	if iHeaderLen > iMaxLen || iHeaderLen < 4 {
		status = PKG_ERR
		return
	}
	if reqLen < iHeaderLen {
		status = PKG_LESS
		return
	}
	pkgLen = reqLen
	status = PKG_FULL
	return
}

func isNoDataErr(err error) bool {
	if e, ok := err.(net.Error); ok {
		return e.Temporary() || e.Timeout()
	}else {
		return ok
	}
}

func checkPanic() {
	if pa := recover(); pa != nil {
		if OLOG != nil {
			OLOG.Errorf("[PANIC] %v",pa)
		}else {
			log.Printf("[PANIC] %v",pa)
		}
		debug.PrintStack()
	}
}