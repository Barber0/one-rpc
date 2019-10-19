package onerpc

import "github.com/golang/protobuf/protoc-gen-go/generator"

const (
	ctxPkgPath		=	"context"
)

var (
	ctxPkg	string
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
	o.P("// fffffffffffffffff")
}

func (o *onerpc) GenerateImports(file *generator.FileDescriptor) {

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