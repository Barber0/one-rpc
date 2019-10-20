package onerpc

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"strings"
)

const (
	ctxPkgPath		=	"context"
	errPkgPath		=	"errors"
	rpcPkgPath		=	"github.com/Barber0/one-rpc"
	appProtoPkgPath	=	"github.com/Barber0/one-rpc/protocol"
	packetPkgPath	=	"github.com/Barber0/one-rpc/protocol/res/requestf"
)

var (
	ctxPkg		=	"context"
	errPkg		=	"errors"
	rpcPkg		=	"one"
	appProtoPkg	=	"protocol"
	packetPkg	=	"requestf"
)

type onerpc struct {
	gen *generator.Generator
}

func init() {
	generator.RegisterPlugin(new(onerpc))
}

func (o *onerpc) Name() string {
	return "onerpc"
}

func (o *onerpc) Init(gen *generator.Generator) {
	o.gen = gen
}

func (o *onerpc) Generate(file *generator.FileDescriptor) {
	for _,service := range file.GetService() {
		o.genService(service)
	}
}

func (o *onerpc) genService(svr *descriptor.ServiceDescriptorProto) {
	svrName := svr.GetName()
	privateSvrName := strings.Replace(svrName,svrName[:1],strings.ToLower(svrName[:1]),1)
	o.P("\ntype ",svrName," struct {")
	o.P("    c *",rpcPkg,".ServiceController")
	o.P("}\n")
	o.P("func New",svrName,"(addr ...string) *",svrName," {")
	o.P("    as := &",svrName,"{")
	o.P("        c:    ",rpcPkg,".NewServiceController(\"",svrName,"\",",appProtoPkg,".NewClientProtocol,addr...),")
	o.P("    }")
	o.P("    return as")
	o.P("}\n")
	for _,method := range svr.GetMethod() {
		o.genMethod(svrName,method)
	}
	o.P("type _",privateSvrName," interface {")
	for _,method := range svr.GetMethod() {
		o.P("    ",method.GetName(),"(",ctxPkg,".Context, *",o.typeName(method.GetInputType()),") (*",o.typeName(method.GetOutputType()),", error)")
	}
	o.P("}\n")

	o.P("func (c *",svrName,") RegisterServiceImp(name string, imp _",privateSvrName,") error {")
	o.P("    return ",appProtoPkg,".AddProxy(name, c, imp)")
	o.P("}\n")

	o.P("func (d *",svrName,") Dispatch(ctx ",ctxPkg,".Context, imp interface{}, req *",packetPkg,".ReqPacket, rsp *",packetPkg,".RspPacket) (err error) {")
	o.P("    switch req.FuncName {")
	for _,method := range svr.GetMethod() {
		inT := o.typeName(method.GetInputType())
		outT := o.typeName(method.GetOutputType())
		o.P("    case \"",method.GetName(),"\":")
		o.P("        payload := new(",inT,")")
		o.P("        if err = proto.Unmarshal(req.Content, payload); err != nil {")
		o.P("            return")
		o.P("        }")
		o.P("        var out *",outT)
		o.P("        out,err = imp.(_",privateSvrName,").",method.GetName(),"(ctx, payload)")
		o.P("        if err != nil {")
		o.P("            return")
		o.P("        }")
		o.P("        *rsp = ",packetPkg,".RspPacket{")
		o.P("            Version:",rpcPkg,".ONE_RPC_VERSION,")
		o.P("            ReqId:req.ReqId,")
		o.P("        }")
		o.P("        rsp.Content,_ = proto.Marshal(out)")
	}
	o.P("    default:")
	o.P("        err = errors.New(\"func mismatch, no such func\")")
	o.P("    }")
	o.P("    return")
	o.P("}")
}

func (o *onerpc) genMethod(svrName string, mth *descriptor.MethodDescriptorProto) {
	mthName := mth.GetName()
	inT := o.typeName(mth.GetInputType())
	outT := o.typeName(mth.GetOutputType())
	o.P("func (s *",svrName,") ",mthName,"(ctx ",ctxPkg,".Context, in *",inT,") (out *",outT,", err error) {")
	o.P("    var (")
	o.P("        reqPkg []byte")
	o.P("        rspPkg []byte")
	o.P("    )")
	o.P("    reqPkg,_ = proto.Marshal(in)")
	o.P("    if rspPkg,err = s.c.Send(ctx, \"",svrName,"\", \"",mthName,"\", reqPkg); err != nil {")
	o.P("        return")
	o.P("    }")
	o.P("    out = &",outT,"{}")
	o.P("    proto.Unmarshal(rspPkg,out)")
	o.P("    return")
	o.P("}\n")
}

func (o *onerpc) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		o.AddImport(rpcPkg,rpcPkgPath)
		o.AddImport(errPkg,errPkgPath)
		o.AddImport(appProtoPkg,appProtoPkgPath)
		o.AddImport(ctxPkg,ctxPkgPath)
		o.AddImport(packetPkg,packetPkgPath)
	}
}

func (o *onerpc) AddImport(pkg, path string)  {
	o.P("import ",pkg," \"",path,"\"")
}

func (o *onerpc) P(args ...interface{}) {
	o.gen.P(args...)
}

func (o *onerpc) objectName(name string) generator.Object {
	o.gen.RecordTypeUse(name)
	return o.gen.ObjectNamed(name)
}

func (o *onerpc) typeName(str string) string {
	return o.gen.TypeName(o.objectName(str))
}