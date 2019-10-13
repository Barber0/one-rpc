package one

import "time"

const (
	ONE_RPC_VERSION		=	1
	PROTOCOL			=	"tcp"

	AcceptTimeout		=	500 * time.Millisecond
	QueueCap			=	10000

	TCPReadBuf			=	128 * 1024 * 1024
	TCPWriteBuf			=	128 * 1024 * 1024
)
