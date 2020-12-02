package calc

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddRequest struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
	N2 int32 `protobuf:"varint,2,opt,name=n2" json:"n2,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *AddRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type AddReply struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

type SubtractRequest struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
	N2 int32 `protobuf:"varint,2,opt,name=n2" json:"n2,omitempty"`
}

func (m *SubtractRequest) Reset()                    { *m = SubtractRequest{} }
func (m *SubtractRequest) String() string            { return proto.CompactTextString(m) }
func (*SubtractRequest) ProtoMessage()               {}
func (*SubtractRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SubtractRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *SubtractRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type SubtractReply struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
}

func (m *SubtractReply) Reset()                    { *m = SubtractReply{} }
func (m *SubtractReply) String() string            { return proto.CompactTextString(m) }
func (*SubtractReply) ProtoMessage()               {}
func (*SubtractReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SubtractReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

type MultiplyRequest struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
	N2 int32 `protobuf:"varint,2,opt,name=n2" json:"n2,omitempty"`
}

func (m *MultiplyRequest) Reset()                    { *m = MultiplyRequest{} }
func (m *MultiplyRequest) String() string            { return proto.CompactTextString(m) }
func (*MultiplyRequest) ProtoMessage()               {}
func (*MultiplyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MultiplyRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *MultiplyRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type MultiplyReply struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
}

func (m *MultiplyReply) Reset()                    { *m = MultiplyReply{} }
func (m *MultiplyReply) String() string            { return proto.CompactTextString(m) }
func (*MultiplyReply) ProtoMessage()               {}
func (*MultiplyReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MultiplyReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

type DivideRequest struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
	N2 int32 `protobuf:"varint,2,opt,name=n2" json:"n2,omitempty"`
}

func (m *DivideRequest) Reset()                    { *m = DivideRequest{} }
func (m *DivideRequest) String() string            { return proto.CompactTextString(m) }
func (*DivideRequest) ProtoMessage()               {}
func (*DivideRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DivideRequest) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func (m *DivideRequest) GetN2() int32 {
	if m != nil {
		return m.N2
	}
	return 0
}

type DivideReply struct {
	N1 int32 `protobuf:"varint,1,opt,name=n1" json:"n1,omitempty"`
}

func (m *DivideReply) Reset()                    { *m = DivideReply{} }
func (m *DivideReply) String() string            { return proto.CompactTextString(m) }
func (*DivideReply) ProtoMessage()               {}
func (*DivideReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DivideReply) GetN1() int32 {
	if m != nil {
		return m.N1
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "calc.AddRequest")
	proto.RegisterType((*AddReply)(nil), "calc.AddReply")
	proto.RegisterType((*SubtractRequest)(nil), "calc.SubtractRequest")
	proto.RegisterType((*SubtractReply)(nil), "calc.SubtractReply")
	proto.RegisterType((*MultiplyRequest)(nil), "calc.MultiplyRequest")
	proto.RegisterType((*MultiplyReply)(nil), "calc.MultiplyReply")
	proto.RegisterType((*DivideRequest)(nil), "calc.DivideRequest")
	proto.RegisterType((*DivideReply)(nil), "calc.DivideReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calculator service

type CalculatorClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractReply, error)
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyReply, error)
	Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideReply, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/calc.Calculator/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractReply, error) {
	out := new(SubtractReply)
	err := grpc.Invoke(ctx, "/calc.Calculator/Subtract", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyReply, error) {
	out := new(MultiplyReply)
	err := grpc.Invoke(ctx, "/calc.Calculator/Multiply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideReply, error) {
	out := new(DivideReply)
	err := grpc.Invoke(ctx, "/calc.Calculator/Divide", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calculator service

type CalculatorServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
	Subtract(context.Context, *SubtractRequest) (*SubtractReply, error)
	Multiply(context.Context, *MultiplyRequest) (*MultiplyReply, error)
	Divide(context.Context, *DivideRequest) (*DivideReply, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calculator/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Subtract(ctx, req.(*SubtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calculator/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.Calculator/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Divide(ctx, req.(*DivideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _Calculator_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Calculator_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Calculator_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}

func init() { proto.RegisterFile("calc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0xcc, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x74, 0xb8, 0xb8, 0x1c, 0x53,
	0x52, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0xf2, 0x0c, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0xf2, 0x0c, 0xc1, 0x7c, 0x23, 0x09, 0x26, 0x28, 0xdf, 0x48,
	0x49, 0x8a, 0x8b, 0x03, 0xac, 0xba, 0x20, 0xa7, 0x12, 0x5d, 0xad, 0x92, 0x21, 0x17, 0x7f, 0x70,
	0x69, 0x52, 0x49, 0x51, 0x62, 0x72, 0x09, 0xb1, 0xc6, 0xc9, 0x73, 0xf1, 0x22, 0xb4, 0xe0, 0x30,
	0xd3, 0xb7, 0x34, 0xa7, 0x24, 0xb3, 0x20, 0xa7, 0x92, 0x04, 0x33, 0x11, 0x5a, 0xb0, 0x99, 0xa9,
	0xcf, 0xc5, 0xeb, 0x92, 0x59, 0x96, 0x99, 0x92, 0x4a, 0xac, 0x89, 0xb2, 0x5c, 0xdc, 0x30, 0x0d,
	0x58, 0xcc, 0x33, 0xba, 0xcf, 0xc8, 0xc5, 0xe5, 0x9c, 0x98, 0x93, 0x5c, 0x9a, 0x93, 0x58, 0x92,
	0x5f, 0x24, 0xa4, 0xc9, 0xc5, 0xec, 0x98, 0x92, 0x22, 0x24, 0xa0, 0x07, 0x0e, 0x6a, 0x44, 0xd8,
	0x4a, 0xf1, 0x21, 0x89, 0x14, 0xe4, 0x54, 0x2a, 0x31, 0x08, 0x59, 0x70, 0x71, 0xc0, 0xbc, 0x2f,
	0x24, 0x0a, 0x91, 0x45, 0x0b, 0x41, 0x29, 0x61, 0x74, 0x61, 0xb8, 0x4e, 0x98, 0x27, 0x61, 0x3a,
	0xd1, 0xc2, 0x09, 0xa6, 0x13, 0x25, 0x2c, 0x94, 0x18, 0x84, 0x8c, 0xb8, 0xd8, 0x20, 0x9e, 0x11,
	0x82, 0x2a, 0x40, 0x09, 0x0b, 0x29, 0x41, 0x54, 0x41, 0xb0, 0x9e, 0x24, 0x36, 0x70, 0x82, 0x31,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xa3, 0x15, 0x87, 0x3e, 0x02, 0x00, 0x00,
}
