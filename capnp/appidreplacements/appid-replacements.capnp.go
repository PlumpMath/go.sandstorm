package appidreplacements

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in appid-replacements.capnp.
var (
	AppIdReplacementList = AppIdReplacement_List{List: capnp.MustUnmarshalRootPtr(x_a53cae3f717a1676[0:712]).List()}
)

type AppIdReplacement struct{ capnp.Struct }

// AppIdReplacement_TypeID is the unique identifier for the type AppIdReplacement.
const AppIdReplacement_TypeID = 0x888dcc6878baa07a

func NewAppIdReplacement(s *capnp.Segment) (AppIdReplacement, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return AppIdReplacement{st}, err
}

func NewRootAppIdReplacement(s *capnp.Segment) (AppIdReplacement, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return AppIdReplacement{st}, err
}

func ReadRootAppIdReplacement(msg *capnp.Message) (AppIdReplacement, error) {
	root, err := msg.RootPtr()
	return AppIdReplacement{root.Struct()}, err
}

func (s AppIdReplacement) String() string {
	str, _ := text.Marshal(0x888dcc6878baa07a, s.Struct)
	return str
}

func (s AppIdReplacement) Original() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s AppIdReplacement) HasOriginal() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s AppIdReplacement) OriginalBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s AppIdReplacement) SetOriginal(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s AppIdReplacement) Replacement() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s AppIdReplacement) HasReplacement() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s AppIdReplacement) ReplacementBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s AppIdReplacement) SetReplacement(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s AppIdReplacement) RevokeExceptPackageIds() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s AppIdReplacement) HasRevokeExceptPackageIds() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s AppIdReplacement) SetRevokeExceptPackageIds(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewRevokeExceptPackageIds sets the revokeExceptPackageIds field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s AppIdReplacement) NewRevokeExceptPackageIds(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// AppIdReplacement_List is a list of AppIdReplacement.
type AppIdReplacement_List struct{ capnp.List }

// NewAppIdReplacement creates a new list of AppIdReplacement.
func NewAppIdReplacement_List(s *capnp.Segment, sz int32) (AppIdReplacement_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return AppIdReplacement_List{l}, err
}

func (s AppIdReplacement_List) At(i int) AppIdReplacement { return AppIdReplacement{s.List.Struct(i)} }

func (s AppIdReplacement_List) Set(i int, v AppIdReplacement) error {
	return s.List.SetStruct(i, v.Struct)
}

// AppIdReplacement_Promise is a wrapper for a AppIdReplacement promised by a client call.
type AppIdReplacement_Promise struct{ *capnp.Pipeline }

func (p AppIdReplacement_Promise) Struct() (AppIdReplacement, error) {
	s, err := p.Pipeline.Struct()
	return AppIdReplacement{s}, err
}

const schema_a53cae3f717a1676 = "x\xda|\xd0\xbd\x8e\xddD\x14\x07\xf0s\xec5+E" +
	"\xab\xbd\xb16\x04\xa5\"\x0de\x90\xc7_3\x83\"\x85" +
	" \xb6X\x14\xa4\xbd\x89\x90B\x91b\xeexl\xaf'" +
	"v\xc6_sm?\x01\x144\x14tH44W\x02" +
	"*\x9a\xbc\x00\x05\xf0\x0eT\xbc\x03\x052ZD\x14h" +
	"h\x7f\xfa\xeb\x9c\xf3?\x81\x83\xef;\xc4\xfb\xc9\x01\xd8" +
	"\xde\xf5\xdeX\x97o^N\xe5/_|\x06\xfe\x1d\\" +
	"\xed\xed\xa5}\xf0\xfd\xfdo\xc1s\x8f\x01\xa2\x0b|\x8c" +
	"g\xcf\xf0\x18\xe0\xecS\xfc\x01p\xfd\xed\xad?\xbf\xfb" +
	"\xea\xcb\x9f\x7f\x07\xff\xce\xd1\xeb0`\xe4;\x1d\x9e\xbd" +
	"\xe3\x1c\x03<\xb9\xeb\xb8\xf8\xe4\xbe\xe3 |\xb0\x0ac" +
	"\xae\xb2{\x9d:2\xcf\x85T\xb5j\x86\xfe])L" +
	"c\xde{h\xccE\xf6X\xbdr\x1c.\x11\xb7'\xee" +
	"\x11\xc0\x11\x02\xf8\xe7\x1f\x01l?tq{\xe9\xa0\x8f" +
	"x\x0b\xaf\xf1\xe3\x1d\xc0\xf6\x91\x8b\xdb\xa7\x0e\xfa\x8es" +
	"\x0b\x1d\x00\xff\x93\x97\x00\xdb\xa7.n\x07\x07\xd7\x17\xdd" +
	"Uq\xd5\x88\xe7\x00\x80'\xe0\xe0\x09\xe0\xda\xfd\xb3\x04" +
	"\x8eU3\xfcK\xed\x0b\xad\xce'\x94\xca\x0c\x97B\xea" +
	"\x07\xa2P\x17Y\x8f\xa7\x80\x97.\xfe\x9d;\x05\xfc\x9f" +
	"\x06\xe2\xbf\x0d\xde\x1e\x1e]\xf5\xd7-^M\xb8\xf9\xfa" +
	"\xb5p\x8d\xe0c\xf1\xf0\xa6K\xee\x1d\x90\x9c\x1f\xd0\x7f" +
	"vJ\xda\x03\x92\xcf\x0f\x08@\xbe> \xf9\xf1\x80\xfe" +
	"\xaf\xb7\xc9\x1f\x07\x8c|\xbc\xd6\xd5VVi%K/" +
	"\x8b8\xcb\x1b2\x10\xdd\x90\xac\xc8\x9aZ(\xdd\xb6\xbc" +
	"\xd2\x95\x8d\x96\xbeX\xe6Z\xc6\x0b'\xd1\xa6Sy\xb0" +
	"\xee[\x9e\xb4\xf58xR\xcb`\xcek%m\xa0g" +
	"\x9e\xcav*\x0cI\xd2\xd1\xb0\xde22O\xb6V\xdd" +
	"\x9c\xb0v\xc3hU\x12\xbc\x81\xeb.\xd9\xedx\xc6\x02" +
	"\x97\xc4\"\xc8\xf9\x8ed)Q!Q\x94\xa7\x19e\x99" +
	"\x94\xd7WM\x8a\x89\xbd\x9c<;T\xe9<\x06\xb6\xa8" +
	"LM\x86^\xa8\x91N\xcc\xb2!7\x94\x0c\xb3m\xea" +
	"t\xd6Z\x99r3\xf26X\x1b&\x1bJ\xe2\xc0\xa3" +
	"M\\\xab\xa5\xa1u\x11\xe8D\xeb:$\xd58\x9a\xb2" +
	"TR\x86q\x99\xe5\\\xe7I\xba\xcc\xd3\xa6\x8eE\xb9" +
	"6y[N,\xb2\x9e]j\x16\xa8\xcc\x14\xda\x88\x89" +
	"\xd5ek\x08M\xdb*\xb4\xfb\"\xa5]\xd1&*\xd2" +
	"\x95L\xbaM,\xe7r\x15\xd1>\x09J\x12{QR" +
	"\xf4\x93\\\xc6\xa2&i\x1b\x88z\xaf\xdb\x9a\xe7\xb1\xec" +
	"f\xbd(\x91D\xbdM\x89\x19\xa8\xd9\x94\x9a\x95\xc4\xbb" +
	"\x81\xe4\xcd\x1b\xb8\x928\x09T \x85+B\xbeK\xb8" +
	"\x92<b\xbb\x88\xf2d\x97\x13*wA\x08\xb0\xd2\x88" +
	"\xe5\x81JRWp)\xe24T\x94\x86q\xa2\xa2<" +
	"\xe2a\xca\xd2\x8c\x02\xac5K\xdb \xe92\xcfV$" +
	"\x9emCYQ\x8ain\xda\x85\x8eao\xf7i\xd7" +
	"\x0c\x83\x19\xa6)\xe6\x05\xa1,\x91\x9b\xcc\x92rM\xab" +
	"\x85\x08\xb1\x1f\xbd\x8e\xeaB\xd0!\xeb+\xae\x0b3M" +
	"d^\xcatZH[w\xa6iG-\x0d\xe9\x946" +
	"\xdd&\xe3y\xf0W\x00\x00\x00\xff\xff\xd6\x83-H"

func init() {
	schemas.Register(schema_a53cae3f717a1676,
		0x888dcc6878baa07a,
		0xe6cb9296adfd17e0)
}

var x_a53cae3f717a1676 = []byte{
	0, 0, 0, 0, 88, 0, 0, 0,
	1, 0, 0, 0, 103, 0, 0, 0,
	16, 0, 0, 0, 0, 0, 3, 0,
	45, 0, 0, 0, 170, 1, 0, 0,
	69, 0, 0, 0, 170, 1, 0, 0,
	93, 0, 0, 0, 14, 0, 0, 0,
	113, 0, 0, 0, 170, 1, 0, 0,
	137, 0, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	157, 0, 0, 0, 170, 1, 0, 0,
	181, 0, 0, 0, 170, 1, 0, 0,
	205, 0, 0, 0, 22, 0, 0, 0,
	249, 0, 0, 0, 170, 1, 0, 0,
	17, 1, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	118, 106, 118, 101, 107, 101, 99, 104,
	100, 51, 57, 56, 102, 110, 49, 116,
	49, 107, 110, 49, 100, 103, 100, 110,
	109, 97, 101, 107, 113, 113, 57, 106,
	107, 106, 118, 51, 122, 115, 103, 122,
	121, 109, 99, 52, 122, 57, 49, 51,
	114, 101, 102, 48, 0, 0, 0, 0,
	119, 113, 57, 53, 113, 109, 117, 116,
	99, 107, 99, 48, 121, 102, 109, 101,
	99, 118, 48, 107, 121, 57, 54, 99,
	113, 120, 103, 112, 49, 53, 54, 117,
	112, 56, 115, 118, 56, 49, 121, 120,
	118, 109, 101, 114, 121, 53, 56, 113,
	56, 55, 106, 104, 0, 0, 0, 0,
	1, 0, 0, 0, 10, 1, 0, 0,
	98, 53, 98, 98, 57, 100, 56, 48,
	49, 52, 97, 48, 102, 57, 98, 49,
	100, 54, 49, 101, 50, 49, 101, 55,
	57, 54, 100, 55, 56, 100, 99, 99,
	0, 0, 0, 0, 0, 0, 0, 0,
	118, 120, 101, 56, 97, 119, 99, 120,
	118, 116, 106, 54, 121, 117, 48, 118,
	103, 106, 112, 109, 49, 116, 115, 97,
	101, 117, 55, 120, 56, 118, 56, 116,
	102, 112, 55, 49, 116, 121, 118, 110,
	109, 54, 121, 107, 107, 101, 112, 104,
	117, 57, 113, 48, 0, 0, 0, 0,
	110, 56, 99, 110, 55, 49, 52, 48,
	55, 110, 52, 109, 101, 122, 110, 55,
	109, 103, 48, 107, 53, 107, 107, 109,
	50, 49, 106, 117, 117, 112, 104, 104,
	101, 99, 99, 50, 52, 104, 100, 102,
	57, 107, 102, 53, 54, 122, 121, 120,
	109, 52, 97, 104, 0, 0, 0, 0,
	110, 102, 113, 104, 120, 56, 51, 118,
	118, 122, 109, 56, 48, 101, 100, 112,
	103, 107, 112, 97, 120, 56, 109, 104,
	113, 112, 49, 55, 54, 113, 106, 50,
	118, 119, 103, 54, 55, 114, 103, 113,
	53, 101, 51, 107, 106, 99, 53, 114,
	52, 99, 121, 104, 0, 0, 0, 0,
	97, 51, 119, 53, 48, 104, 49, 52,
	51, 53, 103, 115, 120, 99, 122, 117,
	103, 109, 49, 54, 113, 48, 97, 109,
	119, 107, 113, 109, 57, 102, 52, 99,
	114, 121, 107, 122, 101, 97, 53, 51,
	115, 118, 54, 49, 112, 116, 55, 112,
	104, 107, 56, 104, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 1, 0, 0,
	21, 0, 0, 0, 10, 1, 0, 0,
	49, 52, 53, 48, 101, 48, 99, 97,
	97, 50, 57, 98, 53, 57, 101, 99,
	57, 51, 56, 98, 51, 55, 57, 53,
	98, 102, 49, 55, 99, 98, 48, 50,
	0, 0, 0, 0, 0, 0, 0, 0,
	55, 51, 56, 102, 48, 101, 53, 54,
	97, 57, 99, 97, 52, 54, 50, 101,
	55, 55, 50, 52, 53, 101, 51, 102,
	51, 57, 50, 54, 56, 54, 100, 55,
	0, 0, 0, 0, 0, 0, 0, 0,
	109, 56, 54, 113, 48, 53, 114, 100,
	118, 106, 49, 52, 121, 118, 110, 55,
	56, 103, 104, 97, 120, 121, 110, 113,
	122, 55, 117, 50, 115, 118, 119, 54,
	114, 110, 116, 116, 112, 116, 120, 120,
	52, 57, 103, 49, 55, 56, 53, 99,
	100, 118, 49, 104, 0, 0, 0, 0,
	54, 106, 122, 49, 97, 97, 119, 117,
	114, 55, 107, 103, 97, 55, 116, 100,
	115, 106, 57, 107, 103, 112, 120, 120,
	49, 121, 122, 104, 54, 120, 122, 49,
	113, 109, 114, 112, 110, 113, 117, 107,
	99, 112, 49, 114, 101, 107, 112, 114,
	100, 57, 102, 48, 0, 0, 0, 0,
}
