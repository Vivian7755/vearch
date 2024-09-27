// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package gamma_api

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FieldInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsFieldInfo(buf []byte, offset flatbuffers.UOffsetT) *FieldInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FieldInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FieldInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FieldInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FieldInfo) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FieldInfo) DataType() DataType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FieldInfo) MutateDataType(n DataType) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

func (rcv *FieldInfo) IsIndex() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *FieldInfo) MutateIsIndex(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func FieldInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func FieldInfoAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func FieldInfoAddDataType(builder *flatbuffers.Builder, dataType int8) {
	builder.PrependInt8Slot(1, dataType, 0)
}
func FieldInfoAddIsIndex(builder *flatbuffers.Builder, isIndex bool) {
	builder.PrependBoolSlot(2, isIndex, false)
}
func FieldInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
