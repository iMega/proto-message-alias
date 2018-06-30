package main

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type Plugin struct {
	g *generator.Generator
}

func New() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Name() string {
	return "Message alias"
}

func (p *Plugin) Init(g *generator.Generator) {
	p.g = g
}

func (p *Plugin) Generate(file *generator.FileDescriptor) {
	p.g.P("//sdfsfsdf")
}

func (p *Plugin) GenerateImports(file *generator.FileDescriptor) {
	p.g.P("// RPC Imports")
}

func init() {
	generator.RegisterPlugin(new(Plugin))
}