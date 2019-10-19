package one

import "time"

const (
	ONE_RPC_VERSION		=	1
	PROTOCOL			=	"tcp"

	SvrAcceptTimeout		=	500 * time.Millisecond

	CltDialTimeout			=	500 * time.Millisecond
	CltReadTimeout			=	500 * time.Millisecond
	CltWriteTimeout			=	500 * time.Millisecond
	CltIdleTimeout			=	500 * time.Millisecond

	QueueCap			=	10000

	TCPReadBuf			=	128 * 1024 * 1024
	TCPWriteBuf			=	128 * 1024 * 1024

	CLT_REQ_TIMEOUT		=	3 * time.Second

	NORMAL_BALANCE		=	"normal"
)
