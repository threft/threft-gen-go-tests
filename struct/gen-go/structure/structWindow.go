package structure

import ()

type Window struct {
	Width  int32
	Height int32
	Broken bool
}

func NewWindow() *Window {
	return &Window{}
}

func (p *Window) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	//++ read code that /knows/ the struct's metadata
	// calls readField_Width, readField_Height, readField_Broken
}

func (p *Window) readField_Width(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	//++ read code that /knows/ the field's metadata
}

func (p *Window) readField_Height(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	//++ read code that /knows/ the field's metadata
}

func (p *Window) readField_Broken(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	//++ read code that /knows/ the field's metadata
}

func (p *Window) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	//++ write code that /knows/ the struct's metadata
	// calls writeField_Width, writeField_Height, writeField_Broken
}

func (p *Window) writeField_Width(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	//++ write code that /knows/ the field's metadata
}

func (p *Window) writeField_Height(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	//++ write code that /knows/ the field's metadata
}

func (p *Window) writeField_Broken(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	//++ write code that /knows/ the field's metadata
}
