package updatetool

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in update-tool.capnp.
var (
	UpdatePublicKeys = PublicSigningKey_List{List: capnp.MustUnmarshalRootPtr(x_96c3fff3f4beb8fe[0:56]).List()}
)

type PublicSigningKey struct{ capnp.Struct }

// PublicSigningKey_TypeID is the unique identifier for the type PublicSigningKey.
const PublicSigningKey_TypeID = 0x9b54bbec5de99f09

func NewPublicSigningKey(s *capnp.Segment) (PublicSigningKey, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	return PublicSigningKey{st}, err
}

func NewRootPublicSigningKey(s *capnp.Segment) (PublicSigningKey, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	return PublicSigningKey{st}, err
}

func ReadRootPublicSigningKey(msg *capnp.Message) (PublicSigningKey, error) {
	root, err := msg.RootPtr()
	return PublicSigningKey{root.Struct()}, err
}

func (s PublicSigningKey) String() string {
	str, _ := text.Marshal(0x9b54bbec5de99f09, s.Struct)
	return str
}

func (s PublicSigningKey) Key0() uint64 {
	return s.Struct.Uint64(0)
}

func (s PublicSigningKey) SetKey0(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s PublicSigningKey) Key1() uint64 {
	return s.Struct.Uint64(8)
}

func (s PublicSigningKey) SetKey1(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s PublicSigningKey) Key2() uint64 {
	return s.Struct.Uint64(16)
}

func (s PublicSigningKey) SetKey2(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s PublicSigningKey) Key3() uint64 {
	return s.Struct.Uint64(24)
}

func (s PublicSigningKey) SetKey3(v uint64) {
	s.Struct.SetUint64(24, v)
}

// PublicSigningKey_List is a list of PublicSigningKey.
type PublicSigningKey_List struct{ capnp.List }

// NewPublicSigningKey creates a new list of PublicSigningKey.
func NewPublicSigningKey_List(s *capnp.Segment, sz int32) (PublicSigningKey_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0}, sz)
	return PublicSigningKey_List{l}, err
}

func (s PublicSigningKey_List) At(i int) PublicSigningKey { return PublicSigningKey{s.List.Struct(i)} }

func (s PublicSigningKey_List) Set(i int, v PublicSigningKey) error {
	return s.List.SetStruct(i, v.Struct)
}

// PublicSigningKey_Promise is a wrapper for a PublicSigningKey promised by a client call.
type PublicSigningKey_Promise struct{ *capnp.Pipeline }

func (p PublicSigningKey_Promise) Struct() (PublicSigningKey, error) {
	s, err := p.Pipeline.Struct()
	return PublicSigningKey{s}, err
}

type Signature struct{ capnp.Struct }

// Signature_TypeID is the unique identifier for the type Signature.
const Signature_TypeID = 0xc4d0558d33210637

func NewSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
	return Signature{st}, err
}

func NewRootSignature(s *capnp.Segment) (Signature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0})
	return Signature{st}, err
}

func ReadRootSignature(msg *capnp.Message) (Signature, error) {
	root, err := msg.RootPtr()
	return Signature{root.Struct()}, err
}

func (s Signature) String() string {
	str, _ := text.Marshal(0xc4d0558d33210637, s.Struct)
	return str
}

func (s Signature) Sig0() uint64 {
	return s.Struct.Uint64(0)
}

func (s Signature) SetSig0(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Signature) Sig1() uint64 {
	return s.Struct.Uint64(8)
}

func (s Signature) SetSig1(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Signature) Sig2() uint64 {
	return s.Struct.Uint64(16)
}

func (s Signature) SetSig2(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Signature) Sig3() uint64 {
	return s.Struct.Uint64(24)
}

func (s Signature) SetSig3(v uint64) {
	s.Struct.SetUint64(24, v)
}

func (s Signature) Sig4() uint64 {
	return s.Struct.Uint64(32)
}

func (s Signature) SetSig4(v uint64) {
	s.Struct.SetUint64(32, v)
}

func (s Signature) Sig5() uint64 {
	return s.Struct.Uint64(40)
}

func (s Signature) SetSig5(v uint64) {
	s.Struct.SetUint64(40, v)
}

func (s Signature) Sig6() uint64 {
	return s.Struct.Uint64(48)
}

func (s Signature) SetSig6(v uint64) {
	s.Struct.SetUint64(48, v)
}

func (s Signature) Sig7() uint64 {
	return s.Struct.Uint64(56)
}

func (s Signature) SetSig7(v uint64) {
	s.Struct.SetUint64(56, v)
}

// Signature_List is a list of Signature.
type Signature_List struct{ capnp.List }

// NewSignature creates a new list of Signature.
func NewSignature_List(s *capnp.Segment, sz int32) (Signature_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 64, PointerCount: 0}, sz)
	return Signature_List{l}, err
}

func (s Signature_List) At(i int) Signature { return Signature{s.List.Struct(i)} }

func (s Signature_List) Set(i int, v Signature) error { return s.List.SetStruct(i, v.Struct) }

// Signature_Promise is a wrapper for a Signature promised by a client call.
type Signature_Promise struct{ *capnp.Pipeline }

func (p Signature_Promise) Struct() (Signature, error) {
	s, err := p.Pipeline.Struct()
	return Signature{s}, err
}

type UpdateSignature struct{ capnp.Struct }

// UpdateSignature_TypeID is the unique identifier for the type UpdateSignature.
const UpdateSignature_TypeID = 0x9c805f76ef46e6c0

func NewUpdateSignature(s *capnp.Segment) (UpdateSignature, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UpdateSignature{st}, err
}

func NewRootUpdateSignature(s *capnp.Segment) (UpdateSignature, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return UpdateSignature{st}, err
}

func ReadRootUpdateSignature(msg *capnp.Message) (UpdateSignature, error) {
	root, err := msg.RootPtr()
	return UpdateSignature{root.Struct()}, err
}

func (s UpdateSignature) String() string {
	str, _ := text.Marshal(0x9c805f76ef46e6c0, s.Struct)
	return str
}

func (s UpdateSignature) Signatures() (Signature_List, error) {
	p, err := s.Struct.Ptr(0)
	return Signature_List{List: p.List()}, err
}

func (s UpdateSignature) HasSignatures() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UpdateSignature) SetSignatures(v Signature_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewSignatures sets the signatures field to a newly
// allocated Signature_List, preferring placement in s's segment.
func (s UpdateSignature) NewSignatures(n int32) (Signature_List, error) {
	l, err := NewSignature_List(s.Struct.Segment(), n)
	if err != nil {
		return Signature_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// UpdateSignature_List is a list of UpdateSignature.
type UpdateSignature_List struct{ capnp.List }

// NewUpdateSignature creates a new list of UpdateSignature.
func NewUpdateSignature_List(s *capnp.Segment, sz int32) (UpdateSignature_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return UpdateSignature_List{l}, err
}

func (s UpdateSignature_List) At(i int) UpdateSignature { return UpdateSignature{s.List.Struct(i)} }

func (s UpdateSignature_List) Set(i int, v UpdateSignature) error {
	return s.List.SetStruct(i, v.Struct)
}

// UpdateSignature_Promise is a wrapper for a UpdateSignature promised by a client call.
type UpdateSignature_Promise struct{ *capnp.Pipeline }

func (p UpdateSignature_Promise) Struct() (UpdateSignature, error) {
	s, err := p.Pipeline.Struct()
	return UpdateSignature{s}, err
}

const schema_96c3fff3f4beb8fe = "x\xda|\x93OH\x14Q\x1c\xc7\x7f\xdf\xf7{\xb3\xab" +
	"P\xe9\xdb\x99K\x9e\xec\x94\x04\x86:\x95\xd0\xa5?\x87" +
	".^\xf6a^\xa4\xc8M\xa7u\xc8tpv\x13\x09" +
	"\"\xe9\x9aA\x97(\xfaG\x97\xa0C\x87\xa0C\x91H" +
	"\x91\x10\x81\x95A\x81\x91\x91\x89\x96\x82R\x94\xb7\xa0\x89" +
	"\xb7[\xbb/\x88\x8e\x9f\x0f\xbf\xf9\xbd\xef|\xe7MK" +
	"\x1f\xf6\x8aVg\x84\x89t\x93\x93Jjo\xac\x1c^" +
	"}x\xf0\x0a\xe9\x0cd\xf2\xf3\xfe\xe4\xfa\xf7\xe4\xc9E" +
	"\x92i\"\xff&\x1a\xe0\xdeE\x9a\xc8\xbd\x83\xcf\x84\xe4" +
	"\xd1\xa7\x03_N\x1e9s\x95T\x06\xd5Y\xc7L\xf8" +
	"c\"\x03\xf7\x820\xc3\xe3b\x0f!iOm\xf1\xc7" +
	"\xbbf\xa6\xcc\xe6\x9a\xbf7\xbb\x8f\xc5\x82\xfb\xdc\xcc\xfa" +
	"\xcf\xc4S\x10\x92\x1f\xe7{\x96&\x1a\x1f|#\x95\xb1" +
	"R\x10\xfcq\xd9\x00\xf7\x9ay\xa8\xf3\x92dt\xde\x93" +
	"\x02\xb4?)F}\xb9B\xd0\\\xe0\xa1\xa1\x81\xed\xbd" +
	"\xb9h0\xda\x9d-\x1e\x1d\x08{;\xc3\xfc`8\x98" +
	"\xef\xe0`4\x0b\xe8z\x96D\x12D*\xb7\x8dH\x1f" +
	"b\xe8~\x01\x05x020\xb2\x87\xa1\x07\x04\x94\x10" +
	"\x1e\x04\x91\x0a\x8d\xecc\xe8H@1{`\"u\xc2" +
	"\xc8~\x86.\x08\xd4\x1d\x0fF[PK\x02\xb5T\x82" +
	"V\x1b\xdal\xf0\xff\xc0\xbf\x12w\x95\x94I\x9c+\x14" +
	"\x87\x11\x98\xc0\xb2\x12xc7\x91\xde\xc0\xd0M\x02I" +
	"\xfc{\x888\x88\xb1\x89\x90e\xa0\xbeZ1\xc1\xc8\xca" +
	"\x19\xa2zFy{\xba8\\\xda\xdeX\xd9\xfe\xca\xbc" +
	"\xcf4C\xcfZu\xbc1r\x86\xa1\xe7\xac:\xde\x1a" +
	"\xf9\x9a\xa1\xe7\xad:\xde\x1b9\xcb\xd0\x8b\x02JJ\x0f" +
	"\x92H}4r\x8e\xa1\x97\x05\x94\xe3xp\x88\xd4\x92" +
	"\x91\xf3\x0c\xbd*\xa0R)\x0f)\"\xb5b\xe4\"C" +
	"\x7f\x15P\xe9\xb4gn\x90Z3r\x99\xa1\xd7\x05\xea" +
	"\xe20_\xad8\x0e\xf3\xad6\xb4\xd9\xe0\xdb\xb0\xc3\x86" +
	"\x9d6\xec\xb2\xa1\xfd\x7f\x9f\xa5\xac\xca\xd7\xa9#\xe0\xd1" +
	"8\x0bT;\xaf\xfc0\xa5\xceIa\xab\x922\x19\xb9" +
	"}j\xf8\xfa\xe5\xe6n\xee}y.\x9a\xbf\xf5\xe2\xec" +
	"\xc4\xe6\xc9\xa9\xb5\xe9\xe8\xc3\xb1x\xdf\xbb\x85\xd35c" +
	"\xbf\x02\x00\x00\xff\xffZ\x94\xdd\x02"

func init() {
	schemas.Register(schema_96c3fff3f4beb8fe,
		0x9b54bbec5de99f09,
		0x9c805f76ef46e6c0,
		0xc4d0558d33210637,
		0xf2b920bce5608efb)
}

var x_96c3fff3f4beb8fe = []byte{
	0, 0, 0, 0, 6, 0, 0, 0,
	1, 0, 0, 0, 39, 0, 0, 0,
	4, 0, 0, 0, 4, 0, 0, 0,
	119, 169, 123, 114, 158, 153, 45, 90,
	99, 207, 140, 112, 224, 166, 206, 131,
	188, 25, 190, 196, 237, 204, 112, 223,
	102, 115, 65, 219, 226, 126, 8, 129,
}
