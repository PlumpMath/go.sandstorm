package appidreplacementstest

// AUTO GENERATED - DO NOT EDIT

import (
	appidreplacements "zenhack.net/go/sandstorm/capnp/appidreplacements"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in appid-replacements-test.capnp.
const (
	TestIds_unusedApp = "6pm4ujs8f5f5wugc87uhuhvhs57he09u10rv8qd2jgdup9f69yzh"
	TestIds_app1      = "vjvekechd398fn1t1kn1dgdnmaekqq9jkjv3zsgzymc4z913ref0"
	TestIds_app2      = "wq95qmutckc0yfmecv0ky96cqxgp156up8sv81yxvmery58q87jh"
	TestIds_app3      = "302t6c6kf8hjer1kh3469d4ch10d936g7wkwtxcs12pwh9u5axqh"
	TestIds_app4      = "5ddk4uqnstnsqvp3thc2tyed41c7wp4x5ygt20zrh3u0tnv5jqd0"
	TestIds_app5      = "jkz6yhywhp4uk5sgkc5ugwnee57a5h5wu4rfmujtahny5r8g3ych"
	TestIds_app6      = "adk6syfj42fpp3xhgqrrheqgfxkhaw8e1t11vug44ys6pzaxqugh"
	TestIds_unusedPkg = "7300e3448dd2b53e075d0a8481c2bc06"
	TestIds_pkg1      = "b5bb9d8014a0f9b1d61e21e796d78dcc"
	TestIds_pkg2      = "8613a11b8ac365cb36775a6b8ca6176c"
	TestIds_pkg3      = "77c4f45aee83e376d31a5680cdb841a2"
)

// Constants defined in appid-replacements-test.capnp.
var (
	TestAppIdReplacementList = appidreplacements.AppIdReplacement_List{List: capnp.MustUnmarshalRootPtr(x_bee445adfb01a777[0:712]).List()}
)

type TestIds struct{ capnp.Struct }

// TestIds_TypeID is the unique identifier for the type TestIds.
const TestIds_TypeID = 0x9440399ec56efc9b

func NewTestIds(s *capnp.Segment) (TestIds, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TestIds{st}, err
}

func NewRootTestIds(s *capnp.Segment) (TestIds, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TestIds{st}, err
}

func ReadRootTestIds(msg *capnp.Message) (TestIds, error) {
	root, err := msg.RootPtr()
	return TestIds{root.Struct()}, err
}

func (s TestIds) String() string {
	str, _ := text.Marshal(0x9440399ec56efc9b, s.Struct)
	return str
}

// TestIds_List is a list of TestIds.
type TestIds_List struct{ capnp.List }

// NewTestIds creates a new list of TestIds.
func NewTestIds_List(s *capnp.Segment, sz int32) (TestIds_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return TestIds_List{l}, err
}

func (s TestIds_List) At(i int) TestIds { return TestIds{s.List.Struct(i)} }

func (s TestIds_List) Set(i int, v TestIds) error { return s.List.SetStruct(i, v.Struct) }

// TestIds_Promise is a wrapper for a TestIds promised by a client call.
type TestIds_Promise struct{ *capnp.Pipeline }

func (p TestIds_Promise) Struct() (TestIds, error) {
	s, err := p.Pipeline.Struct()
	return TestIds{s}, err
}

const schema_bee445adfb01a777 = "x\xda\xac\xd5]\x88\x1bU\x1b\x07\xf0sv2\x9bR" +
	"\xda\xa6\xfb\x86\xbeW\xa5\x8b\x05E\x82[\xe7\xcc\x993" +
	"\x1f\xa2t\xf7b\x91E\x0b\xbb\x9b\x0a\xb5\xb4\xe2d\xce" +
	"\xc9L3f8\x93\xf9H&7\xa5\xd5\xa2\x05\x05\x85" +
	"\"~\xa0TP\x8a\xa9\x8b\xf6\xa2\xb4\xd4V\xbc\xb0J" +
	"\xa1^\x88\x8a\x1fPP\xe8\x95\xe8\x85\x88Zi!r" +
	"\x92jc\x82\xba\x1b\xbcy\x18\xc2\x8f\xf3\xcfy\xces" +
	"f\x94\xfd\x13\xb3\x13H\xde\x99\x07`i\xb7<\xd9\xfd" +
	"\xe5\xddC\x0fE\x07\xaf<\x0e\xa6\xee\xc8u_\xbe\x11" +
	"\\|\xd5\x9a=\x06\x00\xc4G\xa4\xbd\xb0\xf8\xa2\x94\x07" +
	"\xa0|L\x92`\xf9\xb84\x01\x01\xb8E\xa6\xb6\xc1n" +
	"\xf3\x04\xbc\xbe2\x7f\xf5=\x90\xcb\x03\x80W$\x15\x16" +
	"/Hg\x80\xd4\xfdy\xfa\xe4\xbd\xd7N\xe5\x9f\x1fY" +
	"\xf5\xf6\xdc^X\xb4\x04/k9\x09\x96gs\xbdU" +
	"O?\x90IG^\x92^\x1f\xf1\xbbr\x87a\xd1\xee" +
	"\xf9}\xc2{}\xff0\x09\xdeH7\xde\xfd\xf6\x88O" +
	"\xc4\xfaG{\xfe\x09\xe1\x9f\xed{o\xe5\xbb\xd6\x99\xf3" +
	"O\x8e\xfaW\x84?\xd5\xf3o\x09\x7f\xb6\xefOnT" +
	"\xa6W^x\xe7\xfc\x88\xbf(\xfcW=\xff\x99\xf0\xdf" +
	"\xf4\xfd\x17O/D\xd3\xfe\xe1\xf7G\xfc\x0f\xc2CY" +
	"\xf8\x1b\xc2\xaf\x93{~\xdd\xd7\xcfm\xfd\xdf\xa3\x9f\x7f" +
	"8\xe27\xca\x87a\xf1\xb6\x9e\xdf*K\xb0|g\xdf" +
	"o\xbf\xfa\xfd\xa6\xeb[\xba\x97\xc0\xd4\xb6\xdc\xad\xae\x03" +
	"\x88g\xe4s\xb08\xdf\xf3\xb3\xc2\xef\xeb\xfb\x9f\xb6_" +
	"\xde\x9dH\x87>\x1dY\xff\xcd\xc9\xbd\xb0xaR\xf8" +
	"\xb3\x93\x12,\x7f0\xd9\xf3\x8f|b\x1f,\x7ft\xff" +
	"\xaf#\xfe\x92\xf0Wz\xfeK\xe1\xaf\xf6\xfd:\x9f\x9f" +
	"\xd8\x93\x1a\xd7F\xfc\x8f\xc2\xcb\xf9<\x00\xcby\x09\x96" +
	"7\xe4' \xd8\xd3\xb59?@g\x1a,\xc7\x1f\xb3" +
	"\x1dVgA\x1c\xcd\xc4,\x8aw86\x0f\xf8=\xbb" +
	"Y\x14/\xd0h\x87\xcd%N\x16!\x84\x1b\xc0\x04\xdc" +
	"\x00\x00\x82\x1d\xd8\xad\xf9m=\xf3\xb2\xa6\xecq-\xf1" +
	"I\xe4\xfa\x0eI\xdcf\xc0\x181l\xe2\x91f\xa25" +
	"\xaa\xf5\xa4\x16\xdb^\x90\x91\x86\xe9\x16p\xe6x\x7fF" +
	"J\x7f\x179\xdd\xcb\\\xba\x0b\xc2\x81\xe9[Z\x1e\x18" +
	"\xad\xa5\xd2@_v\x95\x06\x86z\xa140Q\xf3\xa5" +
	"\x81K4W\x1a\x98\x9d\xfbJ\x03\x07m-\x0ft\xcd" +
	"*\x0d\x8c\x0c)\x0d\x9c\x17*u\x93 \x89\x18\x9d\xe3" +
	"\x00\xf2\x82\xcd9\x12E\x15\x05\x8b\xa2\x89BD\xd1o" +
	"\xcaE\x1f@\xb7\xc0}\x17\x89\xa2\x8a\x82\xd7\xd2r<" +
	"\xdcr\xac\xa8\xb1\xee\xe8\xbe\\5\xbd\x1ak \xdf\xc3" +
	"\x9anQ\xcd\xf1\x90B-\xac\xbbF\xd3o\xc6-'" +
	"B*ozVB\x0av+\xbc\xd5r\xf9\x9f#\xff" +
	"\xd8\x1e\x07\xc3\xb9:\xafkI-2\xe5*\xa9\x92f" +
	"\xe2:\xa6\x91x\x89\x97z\x111<\xa6X\x09R\x1a" +
	"\xa9\x19R\xb5\xe6\xd2\x84[U\xbd`emo-[" +
	"E\xc3\x91i-e>s<\x99b\xcb\xac\x06(F" +
	"~\x80\xa8K\x83\xba\xcd\xfc0\xb4j~-\xc5\xed\xc8" +
	"mguGk[\x08\x17\x1a\xac\xaa\xac%R\x1b\x8e" +
	"$\x94\xfaZ\x12\x06r\x14\x07Q\x98r\x1c{\x8e\x1a" +
	"g\x8cj\xc81\x9a\\k\x91\xcc\x8dU\xa5\xdd\xf0p" +
	"\xa2\xc4AJ\x0a\xb5\x90\xae)R\x1f\x8e\xb4\xa9\xafG" +
	"Y\xb5&kj\x95s\xdc\xf2\xdc\xb0\xd1\xf0X\xe8V" +
	"[\xbeg7M\x86b\x84\xd2\xc4\xd5\xb4,\xd2y\xdb" +
	"n\x15\xc2\xc4]uc\xb9/\xb9\xea_\"\xd7\xc3\xae" +
	"\xa9#l#T\x91L\xdb\xc1:q*X7\x0cb" +
	"\xeb\x15\xd3\xb1ud\xe8\x0e\x00k\x1b\x98E\xdf\x05\xc3" +
	"!\x06V\x14\x865M2)U+\x043\xc5 T" +
	"\xb1M\xcdD\x8eZq\x14\xfd\xdfC\xc4\xe3\x1c\xe7\x0b" +
	"t\x99\xdd\x04;\x83\xf8\xc1\x03Q,\xb26\x01\xb8(" +
	"A\xb8\xb9\xdb>~\xae\xe5]~\xe6)\x00\xc4\x8f`" +
	"\x0a\xbas\x9b%4\xd3\x81h\xbe\x03\xa7\xf6\xff\x1f\x1d" +
	"\xed@\xf4Z\x07\x02\x80Nw \xfa\xb8\xf7\xf4m\x07" +
	"\xa2\xdf:\xb08\x057\x8d;g\xcd\xd0\"a=\x89" +
	"e\xc7w\x94\xacZgN\xaa\xf8\x99\xa5;a\xcb\xe5" +
	"\x88\xe8\x097\xa3\xd4DY+\xad\xb3FF\xcc\xb0`" +
	"\x1a5\x0f\xc9\xeb!\xda\xb2\x1ev+\xa4R\xb1\xa8\xa9" +
	"HH\xb3\x95\xaaUATGLE\xcc\xb0tj\x98" +
	"\xd4\x11\x87\xb0\x9a\x83\x1a\xefo\x8c\xfb\x1e\x19\xf3\x82\x8c" +
	"\xf9\xa1\x18\xf7\xfb2\xd6\x95\xeaO\xad\xe1hU\x8d\xd8" +
	"\x12c&f\xd8\xd0)F6\xd1M\xc5\xa1\x15SC" +
	"\xb6:0\xb5\xab\xb8zx\xe4V\xfcw\xeb\x8b\xb7\x89" +
	":\xfc6\x19s\x1a\xd6\xb0%4\xbc\xa5U\x0c\xf2\xef" +
	"\x01\x00\x00\xff\xffR\xda\x85p"

func init() {
	schemas.Register(schema_bee445adfb01a777,
		0x83dd7f735581bbf6,
		0x9440399ec56efc9b,
		0x9607b1f83cab1ff5,
		0xa4039a8503794bb5,
		0xaf2f0d76a56e3559,
		0xaf87bcb778eaad68,
		0xbcb098ad1f300dab,
		0xc0826b1f73498cd7,
		0xc6d560121c91da08,
		0xc9ff15fb0eece422,
		0xd381037554cc22f3,
		0xf747c7537f61d15e,
		0xf8377658a7706b08)
}

var x_bee445adfb01a777 = []byte{
	0, 0, 0, 0, 88, 0, 0, 0,
	1, 0, 0, 0, 103, 0, 0, 0,
	16, 0, 0, 0, 0, 0, 3, 0,
	45, 0, 0, 0, 170, 1, 0, 0,
	69, 0, 0, 0, 170, 1, 0, 0,
	93, 0, 0, 0, 22, 0, 0, 0,
	137, 0, 0, 0, 170, 1, 0, 0,
	161, 0, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	181, 0, 0, 0, 170, 1, 0, 0,
	205, 0, 0, 0, 170, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	225, 0, 0, 0, 170, 1, 0, 0,
	249, 0, 0, 0, 170, 1, 0, 0,
	17, 1, 0, 0, 14, 0, 0, 0,
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
	5, 0, 0, 0, 10, 1, 0, 0,
	21, 0, 0, 0, 10, 1, 0, 0,
	98, 53, 98, 98, 57, 100, 56, 48,
	49, 52, 97, 48, 102, 57, 98, 49,
	100, 54, 49, 101, 50, 49, 101, 55,
	57, 54, 100, 55, 56, 100, 99, 99,
	0, 0, 0, 0, 0, 0, 0, 0,
	56, 54, 49, 51, 97, 49, 49, 98,
	56, 97, 99, 51, 54, 53, 99, 98,
	51, 54, 55, 55, 53, 97, 54, 98,
	56, 99, 97, 54, 49, 55, 54, 99,
	0, 0, 0, 0, 0, 0, 0, 0,
	119, 113, 57, 53, 113, 109, 117, 116,
	99, 107, 99, 48, 121, 102, 109, 101,
	99, 118, 48, 107, 121, 57, 54, 99,
	113, 120, 103, 112, 49, 53, 54, 117,
	112, 56, 115, 118, 56, 49, 121, 120,
	118, 109, 101, 114, 121, 53, 56, 113,
	56, 55, 106, 104, 0, 0, 0, 0,
	51, 48, 50, 116, 54, 99, 54, 107,
	102, 56, 104, 106, 101, 114, 49, 107,
	104, 51, 52, 54, 57, 100, 52, 99,
	104, 49, 48, 100, 57, 51, 54, 103,
	55, 119, 107, 119, 116, 120, 99, 115,
	49, 50, 112, 119, 104, 57, 117, 53,
	97, 120, 113, 104, 0, 0, 0, 0,
	53, 100, 100, 107, 52, 117, 113, 110,
	115, 116, 110, 115, 113, 118, 112, 51,
	116, 104, 99, 50, 116, 121, 101, 100,
	52, 49, 99, 55, 119, 112, 52, 120,
	53, 121, 103, 116, 50, 48, 122, 114,
	104, 51, 117, 48, 116, 110, 118, 53,
	106, 113, 100, 48, 0, 0, 0, 0,
	106, 107, 122, 54, 121, 104, 121, 119,
	104, 112, 52, 117, 107, 53, 115, 103,
	107, 99, 53, 117, 103, 119, 110, 101,
	101, 53, 55, 97, 53, 104, 53, 119,
	117, 52, 114, 102, 109, 117, 106, 116,
	97, 104, 110, 121, 53, 114, 56, 103,
	51, 121, 99, 104, 0, 0, 0, 0,
	106, 107, 122, 54, 121, 104, 121, 119,
	104, 112, 52, 117, 107, 53, 115, 103,
	107, 99, 53, 117, 103, 119, 110, 101,
	101, 53, 55, 97, 53, 104, 53, 119,
	117, 52, 114, 102, 109, 117, 106, 116,
	97, 104, 110, 121, 53, 114, 56, 103,
	51, 121, 99, 104, 0, 0, 0, 0,
	97, 100, 107, 54, 115, 121, 102, 106,
	52, 50, 102, 112, 112, 51, 120, 104,
	103, 113, 114, 114, 104, 101, 113, 103,
	102, 120, 107, 104, 97, 119, 56, 101,
	49, 116, 49, 49, 118, 117, 103, 52,
	52, 121, 115, 54, 112, 122, 97, 120,
	113, 117, 103, 104, 0, 0, 0, 0,
	1, 0, 0, 0, 10, 1, 0, 0,
	55, 55, 99, 52, 102, 52, 53, 97,
	101, 101, 56, 51, 101, 51, 55, 54,
	100, 51, 49, 97, 53, 54, 56, 48,
	99, 100, 98, 56, 52, 49, 97, 50,
	0, 0, 0, 0, 0, 0, 0, 0,
}
