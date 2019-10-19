package protocol

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

const ONE_RPC_VERSION = 1