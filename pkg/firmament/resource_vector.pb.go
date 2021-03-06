/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_vector.proto

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ResourceVector stores all the resources which will be take in account during scheduling.
type ResourceVector struct {
	// cpu_cores is the cpu millicores of node..
	CpuCores float32 `protobuf:"fixed32,1,opt,name=cpu_cores,json=cpuCores" json:"cpu_cores,omitempty"`
	RamBw    uint64  `protobuf:"varint,2,opt,name=ram_bw,json=ramBw" json:"ram_bw,omitempty"`
	// ram_cap is the memory capacity of node.
	RamCap uint64 `protobuf:"varint,3,opt,name=ram_cap,json=ramCap" json:"ram_cap,omitempty"`
	DiskBw uint64 `protobuf:"varint,4,opt,name=disk_bw,json=diskBw" json:"disk_bw,omitempty"`
	// disk_cap is the disk capacity of node.
	DiskCap uint64 `protobuf:"varint,5,opt,name=disk_cap,json=diskCap" json:"disk_cap,omitempty"`
	// net_tx_bw is transmit network packets in KB.
	NetTxBw uint64 `protobuf:"varint,6,opt,name=net_tx_bw,json=netTxBw" json:"net_tx_bw,omitempty"`
	// net_tx_bw is receive network packets in KB.
	NetRxBw uint64 `protobuf:"varint,7,opt,name=net_rx_bw,json=netRxBw" json:"net_rx_bw,omitempty"`
}

func (m *ResourceVector) Reset()                    { *m = ResourceVector{} }
func (m *ResourceVector) String() string            { return proto.CompactTextString(m) }
func (*ResourceVector) ProtoMessage()               {}
func (*ResourceVector) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *ResourceVector) GetCpuCores() float32 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *ResourceVector) GetRamBw() uint64 {
	if m != nil {
		return m.RamBw
	}
	return 0
}

func (m *ResourceVector) GetRamCap() uint64 {
	if m != nil {
		return m.RamCap
	}
	return 0
}

func (m *ResourceVector) GetDiskBw() uint64 {
	if m != nil {
		return m.DiskBw
	}
	return 0
}

func (m *ResourceVector) GetDiskCap() uint64 {
	if m != nil {
		return m.DiskCap
	}
	return 0
}

func (m *ResourceVector) GetNetTxBw() uint64 {
	if m != nil {
		return m.NetTxBw
	}
	return 0
}

func (m *ResourceVector) GetNetRxBw() uint64 {
	if m != nil {
		return m.NetRxBw
	}
	return 0
}

func init() {
	proto.RegisterType((*ResourceVector)(nil), "firmament.ResourceVector")
}

func init() { proto.RegisterFile("resource_vector.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcf, 0xbd, 0x6e, 0xc3, 0x20,
	0x14, 0x05, 0x60, 0xe1, 0xfa, 0x97, 0xa1, 0x03, 0x92, 0x55, 0xb7, 0x5d, 0xac, 0x4e, 0x9e, 0xba,
	0xf4, 0x0d, 0xf0, 0x1b, 0xa0, 0xaa, 0x2b, 0xc2, 0x94, 0x48, 0x56, 0x84, 0x41, 0xd7, 0x38, 0xf8,
	0x15, 0xf3, 0x56, 0xd1, 0xc5, 0x49, 0xc6, 0x7b, 0xbe, 0x73, 0x86, 0x4b, 0x5b, 0x30, 0xab, 0xdb,
	0x40, 0x1b, 0x79, 0x31, 0x3a, 0x38, 0xf8, 0xf6, 0xe0, 0x82, 0x63, 0xcd, 0x69, 0x06, 0xab, 0xac,
	0x59, 0xc2, 0xd7, 0x95, 0xd0, 0x57, 0x71, 0x2f, 0xfd, 0xa5, 0x0e, 0xfb, 0xa4, 0x8d, 0xf6, 0x9b,
	0xd4, 0x0e, 0xcc, 0xda, 0x91, 0x9e, 0x0c, 0x99, 0xa8, 0xb5, 0xdf, 0x46, 0xbc, 0x59, 0x4b, 0x4b,
	0x50, 0x56, 0x4e, 0xb1, 0xcb, 0x7a, 0x32, 0xe4, 0xa2, 0x00, 0x65, 0x79, 0x64, 0x6f, 0xb4, 0xc2,
	0x58, 0x2b, 0xdf, 0xbd, 0xa4, 0x1c, 0x5b, 0xa3, 0xf2, 0x08, 0xff, 0xf3, 0x7a, 0xc6, 0x41, 0x7e,
	0x00, 0x9e, 0x3c, 0xb2, 0x77, 0x5a, 0x27, 0xc0, 0x49, 0x91, 0x24, 0x15, 0x71, 0xf3, 0x41, 0x9b,
	0xc5, 0x04, 0x19, 0x76, 0x5c, 0x95, 0x87, 0x2d, 0x26, 0xfc, 0xee, 0x3c, 0x3e, 0x0c, 0x92, 0x55,
	0x4f, 0x13, 0x3b, 0x8f, 0x53, 0x99, 0xbe, 0xfb, 0xb9, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xaa,
	0xd5, 0x72, 0xf6, 0x00, 0x00, 0x00,
}
