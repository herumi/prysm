// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/eth/v1alpha1/slasher.proto

package eth

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ValidatorIDToIdxAtt struct {
	Indices              []uint64 `protobuf:"varint,1,rep,packed,name=indices,proto3" json:"indices,omitempty"`
	DataRoot             []byte   `protobuf:"bytes,2,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidatorIDToIdxAtt) Reset()         { *m = ValidatorIDToIdxAtt{} }
func (m *ValidatorIDToIdxAtt) String() string { return proto.CompactTextString(m) }
func (*ValidatorIDToIdxAtt) ProtoMessage()    {}
func (*ValidatorIDToIdxAtt) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3db2cc39857595b, []int{0}
}
func (m *ValidatorIDToIdxAtt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorIDToIdxAtt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorIDToIdxAtt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorIDToIdxAtt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorIDToIdxAtt.Merge(m, src)
}
func (m *ValidatorIDToIdxAtt) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorIDToIdxAtt) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorIDToIdxAtt.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorIDToIdxAtt proto.InternalMessageInfo

func (m *ValidatorIDToIdxAtt) GetIndices() []uint64 {
	if m != nil {
		return m.Indices
	}
	return nil
}

func (m *ValidatorIDToIdxAtt) GetDataRoot() []byte {
	if m != nil {
		return m.DataRoot
	}
	return nil
}

func (m *ValidatorIDToIdxAtt) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type ValidatorIDToIdxAttList struct {
	IndicesList          []*ValidatorIDToIdxAtt `protobuf:"bytes,1,rep,name=indicesList,proto3" json:"indicesList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ValidatorIDToIdxAttList) Reset()         { *m = ValidatorIDToIdxAttList{} }
func (m *ValidatorIDToIdxAttList) String() string { return proto.CompactTextString(m) }
func (*ValidatorIDToIdxAttList) ProtoMessage()    {}
func (*ValidatorIDToIdxAttList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3db2cc39857595b, []int{1}
}
func (m *ValidatorIDToIdxAttList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorIDToIdxAttList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorIDToIdxAttList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorIDToIdxAttList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorIDToIdxAttList.Merge(m, src)
}
func (m *ValidatorIDToIdxAttList) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorIDToIdxAttList) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorIDToIdxAttList.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorIDToIdxAttList proto.InternalMessageInfo

func (m *ValidatorIDToIdxAttList) GetIndicesList() []*ValidatorIDToIdxAtt {
	if m != nil {
		return m.IndicesList
	}
	return nil
}

type ProposerSlashingRequest struct {
	BlockHeader          *BeaconBlockHeader `protobuf:"bytes,1,opt,name=block_header,json=blockHeader,proto3" json:"block_header,omitempty"`
	ValidatorIndex       uint64             `protobuf:"varint,2,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ProposerSlashingRequest) Reset()         { *m = ProposerSlashingRequest{} }
func (m *ProposerSlashingRequest) String() string { return proto.CompactTextString(m) }
func (*ProposerSlashingRequest) ProtoMessage()    {}
func (*ProposerSlashingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3db2cc39857595b, []int{2}
}
func (m *ProposerSlashingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposerSlashingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposerSlashingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposerSlashingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposerSlashingRequest.Merge(m, src)
}
func (m *ProposerSlashingRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProposerSlashingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposerSlashingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProposerSlashingRequest proto.InternalMessageInfo

func (m *ProposerSlashingRequest) GetBlockHeader() *BeaconBlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

func (m *ProposerSlashingRequest) GetValidatorIndex() uint64 {
	if m != nil {
		return m.ValidatorIndex
	}
	return 0
}

type ProposerSlashingResponse struct {
	ProposerSlashing     []*ProposerSlashing `protobuf:"bytes,1,rep,name=proposer_slashing,json=proposerSlashing,proto3" json:"proposer_slashing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProposerSlashingResponse) Reset()         { *m = ProposerSlashingResponse{} }
func (m *ProposerSlashingResponse) String() string { return proto.CompactTextString(m) }
func (*ProposerSlashingResponse) ProtoMessage()    {}
func (*ProposerSlashingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3db2cc39857595b, []int{3}
}
func (m *ProposerSlashingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposerSlashingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposerSlashingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposerSlashingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposerSlashingResponse.Merge(m, src)
}
func (m *ProposerSlashingResponse) XXX_Size() int {
	return m.Size()
}
func (m *ProposerSlashingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposerSlashingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProposerSlashingResponse proto.InternalMessageInfo

func (m *ProposerSlashingResponse) GetProposerSlashing() []*ProposerSlashing {
	if m != nil {
		return m.ProposerSlashing
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorIDToIdxAtt)(nil), "ethereum.eth.v1alpha1.ValidatorIDToIdxAtt")
	proto.RegisterType((*ValidatorIDToIdxAttList)(nil), "ethereum.eth.v1alpha1.ValidatorIDToIdxAttList")
	proto.RegisterType((*ProposerSlashingRequest)(nil), "ethereum.eth.v1alpha1.ProposerSlashingRequest")
	proto.RegisterType((*ProposerSlashingResponse)(nil), "ethereum.eth.v1alpha1.ProposerSlashingResponse")
}

func init() { proto.RegisterFile("proto/eth/v1alpha1/slasher.proto", fileDescriptor_c3db2cc39857595b) }

var fileDescriptor_c3db2cc39857595b = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0xa4, 0xa2, 0x74, 0x53, 0x41, 0x59, 0xd4, 0xd6, 0x4a, 0x51, 0x14, 0x45, 0xa0,
	0x44, 0x1c, 0xd6, 0x6d, 0x10, 0x27, 0x4e, 0x0d, 0x20, 0x11, 0xd1, 0x03, 0x72, 0x2b, 0x90, 0x7a,
	0xb1, 0xd6, 0xf6, 0x60, 0x2f, 0x38, 0x5e, 0xb3, 0x3b, 0xae, 0xda, 0xa7, 0xe0, 0xce, 0x13, 0x71,
	0xe4, 0x11, 0x50, 0x9e, 0x04, 0x79, 0xed, 0x34, 0x51, 0xe3, 0xa0, 0x70, 0x9c, 0x7f, 0xfe, 0x9d,
	0x7f, 0xfd, 0x79, 0x87, 0xf6, 0x72, 0xad, 0x50, 0xb9, 0x80, 0x89, 0x7b, 0x75, 0x22, 0xd2, 0x3c,
	0x11, 0x27, 0xae, 0x49, 0x85, 0x49, 0x40, 0x73, 0xdb, 0x62, 0xfb, 0x80, 0x09, 0x68, 0x28, 0xa6,
	0x1c, 0x30, 0xe1, 0x73, 0x53, 0xe7, 0x28, 0x56, 0x2a, 0x4e, 0xc1, 0xb5, 0xa6, 0xa0, 0xf8, 0xe2,
	0xc2, 0x34, 0xc7, 0x9b, 0xea, 0x4c, 0xe7, 0x79, 0xc3, 0xd4, 0x00, 0x44, 0xa8, 0x32, 0x3f, 0x48,
	0x55, 0xf8, 0xad, 0xb6, 0x3d, 0x6b, 0xb0, 0x09, 0x44, 0x30, 0x28, 0x50, 0xaa, 0xac, 0x72, 0xf5,
	0xbf, 0xd2, 0x27, 0x9f, 0x44, 0x2a, 0x23, 0x81, 0x4a, 0x4f, 0xde, 0x5e, 0xa8, 0x49, 0x74, 0x7d,
	0x8a, 0xc8, 0x1c, 0xba, 0x2d, 0xb3, 0x48, 0x86, 0x60, 0x1c, 0xd2, 0x6b, 0x0d, 0xb7, 0xbc, 0x79,
	0xc9, 0x8e, 0xe8, 0x4e, 0x24, 0x50, 0xf8, 0x5a, 0x29, 0x74, 0xee, 0xf5, 0xc8, 0x70, 0xd7, 0x7b,
	0x50, 0x0a, 0x9e, 0x52, 0xc8, 0x9e, 0xd2, 0x1d, 0x23, 0xe3, 0x4c, 0x60, 0xa1, 0xc1, 0x69, 0xd9,
	0xe6, 0x42, 0xe8, 0xc7, 0xf4, 0xb0, 0x21, 0xeb, 0x4c, 0x1a, 0x64, 0x67, 0xb4, 0x5d, 0x07, 0x94,
	0xa5, 0xcd, 0x6c, 0x8f, 0x5e, 0xf0, 0x46, 0x3a, 0xbc, 0x61, 0x88, 0xb7, 0x7c, 0xbc, 0xff, 0x83,
	0xd0, 0xc3, 0x8f, 0x5a, 0xe5, 0xca, 0x80, 0x3e, 0x2f, 0x79, 0xcb, 0x2c, 0xf6, 0xe0, 0x7b, 0x01,
	0x06, 0xd9, 0x07, 0xba, 0x6b, 0x29, 0xf9, 0x09, 0x88, 0x08, 0xb4, 0x43, 0x7a, 0x64, 0xd8, 0x1e,
	0x0d, 0xd7, 0x44, 0x8d, 0x2d, 0xd7, 0x71, 0x79, 0xe0, 0xbd, 0xf5, 0x7b, 0xed, 0x60, 0x51, 0xb0,
	0x01, 0x7d, 0x74, 0x35, 0xbf, 0x8c, 0x2f, 0xb3, 0x08, 0xae, 0x2d, 0x92, 0x2d, 0xef, 0xe1, 0xad,
	0x3c, 0x29, 0xd5, 0x7e, 0x4e, 0x9d, 0xd5, 0x0b, 0x99, 0x5c, 0x65, 0x06, 0xd8, 0x05, 0x7d, 0x9c,
	0xd7, 0x3d, 0xdf, 0xd4, 0xcd, 0x9a, 0xc0, 0x60, 0xcd, 0xb5, 0x56, 0x66, 0xed, 0xe5, 0x77, 0x94,
	0xd1, 0xcf, 0x16, 0xdd, 0x3e, 0xaf, 0xde, 0x1a, 0x03, 0x7a, 0x30, 0x31, 0xb6, 0x10, 0x41, 0x0a,
	0xa7, 0x8b, 0x47, 0xc0, 0xfa, 0x6b, 0x02, 0x96, 0x3c, 0x9d, 0xc1, 0x3f, 0x3d, 0x8b, 0x48, 0x66,
	0xe8, 0xde, 0x52, 0x8c, 0x85, 0xc6, 0xf8, 0xa6, 0x5f, 0x50, 0xfd, 0x9e, 0x8e, 0xbb, 0xb1, 0xbf,
	0xa6, 0xf7, 0x99, 0xb2, 0xdb, 0xc8, 0xca, 0x24, 0x52, 0xc3, 0x0e, 0x78, 0xb5, 0x41, 0x7c, 0xbe,
	0x41, 0xfc, 0x5d, 0xb9, 0x41, 0x9d, 0x4d, 0x81, 0x1e, 0x13, 0x76, 0x49, 0xf7, 0x9b, 0x90, 0xfd,
	0xff, 0xec, 0xbb, 0x9c, 0x8e, 0xc9, 0xf8, 0xcd, 0xaf, 0x59, 0x97, 0xfc, 0x9e, 0x75, 0xc9, 0x9f,
	0x59, 0x97, 0x5c, 0xbe, 0x8a, 0x25, 0x26, 0x45, 0xc0, 0x43, 0x35, 0x75, 0x73, 0x7d, 0x63, 0xa6,
	0x02, 0x65, 0x98, 0x8a, 0xc0, 0x54, 0x95, 0xbb, 0xba, 0xc8, 0xaf, 0x01, 0x93, 0xe0, 0xbe, 0xd5,
	0x5f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x04, 0xd4, 0x21, 0x51, 0x66, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SlasherClient is the client API for Slasher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SlasherClient interface {
	IsSlashableAttestation(ctx context.Context, in *Attestation, opts ...grpc.CallOption) (*AttesterSlashing, error)
	IsSlashableBlock(ctx context.Context, in *ProposerSlashingRequest, opts ...grpc.CallOption) (*ProposerSlashingResponse, error)
	SlashableProposals(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (Slasher_SlashableProposalsClient, error)
	SlashableAttestations(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (Slasher_SlashableAttestationsClient, error)
}

type slasherClient struct {
	cc *grpc.ClientConn
}

func NewSlasherClient(cc *grpc.ClientConn) SlasherClient {
	return &slasherClient{cc}
}

func (c *slasherClient) IsSlashableAttestation(ctx context.Context, in *Attestation, opts ...grpc.CallOption) (*AttesterSlashing, error) {
	out := new(AttesterSlashing)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Slasher/IsSlashableAttestation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slasherClient) IsSlashableBlock(ctx context.Context, in *ProposerSlashingRequest, opts ...grpc.CallOption) (*ProposerSlashingResponse, error) {
	out := new(ProposerSlashingResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Slasher/IsSlashableBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slasherClient) SlashableProposals(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (Slasher_SlashableProposalsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Slasher_serviceDesc.Streams[0], "/ethereum.eth.v1alpha1.Slasher/SlashableProposals", opts...)
	if err != nil {
		return nil, err
	}
	x := &slasherSlashableProposalsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Slasher_SlashableProposalsClient interface {
	Recv() (*ProposerSlashing, error)
	grpc.ClientStream
}

type slasherSlashableProposalsClient struct {
	grpc.ClientStream
}

func (x *slasherSlashableProposalsClient) Recv() (*ProposerSlashing, error) {
	m := new(ProposerSlashing)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *slasherClient) SlashableAttestations(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (Slasher_SlashableAttestationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Slasher_serviceDesc.Streams[1], "/ethereum.eth.v1alpha1.Slasher/SlashableAttestations", opts...)
	if err != nil {
		return nil, err
	}
	x := &slasherSlashableAttestationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Slasher_SlashableAttestationsClient interface {
	Recv() (*AttesterSlashing, error)
	grpc.ClientStream
}

type slasherSlashableAttestationsClient struct {
	grpc.ClientStream
}

func (x *slasherSlashableAttestationsClient) Recv() (*AttesterSlashing, error) {
	m := new(AttesterSlashing)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SlasherServer is the server API for Slasher service.
type SlasherServer interface {
	IsSlashableAttestation(context.Context, *Attestation) (*AttesterSlashing, error)
	IsSlashableBlock(context.Context, *ProposerSlashingRequest) (*ProposerSlashingResponse, error)
	SlashableProposals(*types.Empty, Slasher_SlashableProposalsServer) error
	SlashableAttestations(*types.Empty, Slasher_SlashableAttestationsServer) error
}

func RegisterSlasherServer(s *grpc.Server, srv SlasherServer) {
	s.RegisterService(&_Slasher_serviceDesc, srv)
}

func _Slasher_IsSlashableAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attestation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlasherServer).IsSlashableAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Slasher/IsSlashableAttestation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlasherServer).IsSlashableAttestation(ctx, req.(*Attestation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Slasher_IsSlashableBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposerSlashingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlasherServer).IsSlashableBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Slasher/IsSlashableBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlasherServer).IsSlashableBlock(ctx, req.(*ProposerSlashingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Slasher_SlashableProposals_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(types.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SlasherServer).SlashableProposals(m, &slasherSlashableProposalsServer{stream})
}

type Slasher_SlashableProposalsServer interface {
	Send(*ProposerSlashing) error
	grpc.ServerStream
}

type slasherSlashableProposalsServer struct {
	grpc.ServerStream
}

func (x *slasherSlashableProposalsServer) Send(m *ProposerSlashing) error {
	return x.ServerStream.SendMsg(m)
}

func _Slasher_SlashableAttestations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(types.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SlasherServer).SlashableAttestations(m, &slasherSlashableAttestationsServer{stream})
}

type Slasher_SlashableAttestationsServer interface {
	Send(*AttesterSlashing) error
	grpc.ServerStream
}

type slasherSlashableAttestationsServer struct {
	grpc.ServerStream
}

func (x *slasherSlashableAttestationsServer) Send(m *AttesterSlashing) error {
	return x.ServerStream.SendMsg(m)
}

var _Slasher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1alpha1.Slasher",
	HandlerType: (*SlasherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsSlashableAttestation",
			Handler:    _Slasher_IsSlashableAttestation_Handler,
		},
		{
			MethodName: "IsSlashableBlock",
			Handler:    _Slasher_IsSlashableBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SlashableProposals",
			Handler:       _Slasher_SlashableProposals_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SlashableAttestations",
			Handler:       _Slasher_SlashableAttestations_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/eth/v1alpha1/slasher.proto",
}

func (m *ValidatorIDToIdxAtt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorIDToIdxAtt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Indices) > 0 {
		dAtA2 := make([]byte, len(m.Indices)*10)
		var j1 int
		for _, num := range m.Indices {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintSlasher(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.DataRoot) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSlasher(dAtA, i, uint64(len(m.DataRoot)))
		i += copy(dAtA[i:], m.DataRoot)
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSlasher(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ValidatorIDToIdxAttList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorIDToIdxAttList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IndicesList) > 0 {
		for _, msg := range m.IndicesList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSlasher(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ProposerSlashingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposerSlashingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BlockHeader != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSlasher(dAtA, i, uint64(m.BlockHeader.Size()))
		n3, err := m.BlockHeader.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.ValidatorIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSlasher(dAtA, i, uint64(m.ValidatorIndex))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ProposerSlashingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposerSlashingResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProposerSlashing) > 0 {
		for _, msg := range m.ProposerSlashing {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSlasher(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSlasher(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ValidatorIDToIdxAtt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Indices) > 0 {
		l = 0
		for _, e := range m.Indices {
			l += sovSlasher(uint64(e))
		}
		n += 1 + sovSlasher(uint64(l)) + l
	}
	l = len(m.DataRoot)
	if l > 0 {
		n += 1 + l + sovSlasher(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovSlasher(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ValidatorIDToIdxAttList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IndicesList) > 0 {
		for _, e := range m.IndicesList {
			l = e.Size()
			n += 1 + l + sovSlasher(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProposerSlashingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeader != nil {
		l = m.BlockHeader.Size()
		n += 1 + l + sovSlasher(uint64(l))
	}
	if m.ValidatorIndex != 0 {
		n += 1 + sovSlasher(uint64(m.ValidatorIndex))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProposerSlashingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ProposerSlashing) > 0 {
		for _, e := range m.ProposerSlashing {
			l = e.Size()
			n += 1 + l + sovSlasher(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSlasher(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSlasher(x uint64) (n int) {
	return sovSlasher(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorIDToIdxAtt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlasher
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidatorIDToIdxAtt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorIDToIdxAtt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSlasher
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Indices = append(m.Indices, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSlasher
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSlasher
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthSlasher
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Indices) == 0 {
					m.Indices = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSlasher
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Indices = append(m.Indices, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Indices", wireType)
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlasher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlasher
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlasher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRoot = append(m.DataRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.DataRoot == nil {
				m.DataRoot = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlasher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlasher
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlasher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlasher(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlasher
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlasher
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValidatorIDToIdxAttList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlasher
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ValidatorIDToIdxAttList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorIDToIdxAttList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndicesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlasher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlasher
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlasher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IndicesList = append(m.IndicesList, &ValidatorIDToIdxAtt{})
			if err := m.IndicesList[len(m.IndicesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlasher(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlasher
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlasher
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProposerSlashingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlasher
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProposerSlashingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposerSlashingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlasher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlasher
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlasher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockHeader == nil {
				m.BlockHeader = &BeaconBlockHeader{}
			}
			if err := m.BlockHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorIndex", wireType)
			}
			m.ValidatorIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlasher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlasher(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlasher
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlasher
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProposerSlashingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlasher
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProposerSlashingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposerSlashingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerSlashing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlasher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlasher
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlasher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerSlashing = append(m.ProposerSlashing, &ProposerSlashing{})
			if err := m.ProposerSlashing[len(m.ProposerSlashing)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlasher(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlasher
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlasher
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSlasher(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlasher
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSlasher
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSlasher
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSlasher
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSlasher
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSlasher
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSlasher(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSlasher
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSlasher = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlasher   = fmt.Errorf("proto: integer overflow")
)
