package requestf

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
)

func Req2Bytes(req *ReqPacket) (res []byte) {
	pkg,_ := proto.Marshal(req)
	buf := bytes.NewBuffer(make([]byte,4))
	buf.Write(pkg)
	res = buf.Bytes()
	binary.BigEndian.PutUint32(res,uint32(buf.Len()))
	return
}

func Rsp2Bytes(rsp *RspPacket) (res []byte) {
	pkg,_ := proto.Marshal(rsp)
	buf := bytes.NewBuffer(make([]byte,4))
	buf.Write(pkg)
	res = buf.Bytes()
	binary.BigEndian.PutUint32(res,uint32(buf.Len()))
	return
}