package ocache

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// protoMarshaller is a generic marshaller for protobuf messages.
// It's internal to the olric client package and used for cache storage.
type protoMarshaller[T protoreflect.ProtoMessage] struct {
	Msg T // Msg is the protobuf message to be marshaled or unmarshaled.
}

// MarshalBinary marshals the protobuf message to a binary format.
// It implements the encoding.BinaryMarshaler interface.
func (m *protoMarshaller[T]) MarshalBinary() ([]byte, error) {
	return proto.Marshal(m.Msg)
}

// UnmarshalBinary unmarshals a binary format into the protobuf message.
// It implements the encoding.BinaryUnmarshaler interface.
func (m *protoMarshaller[T]) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, m.Msg)
}

// newProtoMarshaller creates a new protoMarshaller for the given protobuf message.
func newProtoMarshaller[T protoreflect.ProtoMessage](protoMsg T) *protoMarshaller[T] {
	return &protoMarshaller[T]{
		Msg: protoMsg,
	}
}
