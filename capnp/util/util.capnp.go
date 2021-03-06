package util

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type KeyValue struct{ capnp.Struct }

// KeyValue_TypeID is the unique identifier for the type KeyValue.
const KeyValue_TypeID = 0x94a081e4abb13424

func NewKeyValue(s *capnp.Segment) (KeyValue, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return KeyValue{st}, err
}

func NewRootKeyValue(s *capnp.Segment) (KeyValue, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return KeyValue{st}, err
}

func ReadRootKeyValue(msg *capnp.Message) (KeyValue, error) {
	root, err := msg.RootPtr()
	return KeyValue{root.Struct()}, err
}

func (s KeyValue) String() string {
	str, _ := text.Marshal(0x94a081e4abb13424, s.Struct)
	return str
}

func (s KeyValue) Key() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s KeyValue) HasKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s KeyValue) KeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s KeyValue) SetKey(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s KeyValue) Value() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s KeyValue) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s KeyValue) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s KeyValue) SetValue(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// KeyValue_List is a list of KeyValue.
type KeyValue_List struct{ capnp.List }

// NewKeyValue creates a new list of KeyValue.
func NewKeyValue_List(s *capnp.Segment, sz int32) (KeyValue_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return KeyValue_List{l}, err
}

func (s KeyValue_List) At(i int) KeyValue { return KeyValue{s.List.Struct(i)} }

func (s KeyValue_List) Set(i int, v KeyValue) error { return s.List.SetStruct(i, v.Struct) }

// KeyValue_Promise is a wrapper for a KeyValue promised by a client call.
type KeyValue_Promise struct{ *capnp.Pipeline }

func (p KeyValue_Promise) Struct() (KeyValue, error) {
	s, err := p.Pipeline.Struct()
	return KeyValue{s}, err
}

type LocalizedText struct{ capnp.Struct }

// LocalizedText_TypeID is the unique identifier for the type LocalizedText.
const LocalizedText_TypeID = 0x8b5db772377be249

func NewLocalizedText(s *capnp.Segment) (LocalizedText, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return LocalizedText{st}, err
}

func NewRootLocalizedText(s *capnp.Segment) (LocalizedText, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return LocalizedText{st}, err
}

func ReadRootLocalizedText(msg *capnp.Message) (LocalizedText, error) {
	root, err := msg.RootPtr()
	return LocalizedText{root.Struct()}, err
}

func (s LocalizedText) String() string {
	str, _ := text.Marshal(0x8b5db772377be249, s.Struct)
	return str
}

func (s LocalizedText) DefaultText() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s LocalizedText) HasDefaultText() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s LocalizedText) DefaultTextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s LocalizedText) SetDefaultText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s LocalizedText) Localizations() (LocalizedText_Localization_List, error) {
	p, err := s.Struct.Ptr(1)
	return LocalizedText_Localization_List{List: p.List()}, err
}

func (s LocalizedText) HasLocalizations() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s LocalizedText) SetLocalizations(v LocalizedText_Localization_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewLocalizations sets the localizations field to a newly
// allocated LocalizedText_Localization_List, preferring placement in s's segment.
func (s LocalizedText) NewLocalizations(n int32) (LocalizedText_Localization_List, error) {
	l, err := NewLocalizedText_Localization_List(s.Struct.Segment(), n)
	if err != nil {
		return LocalizedText_Localization_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// LocalizedText_List is a list of LocalizedText.
type LocalizedText_List struct{ capnp.List }

// NewLocalizedText creates a new list of LocalizedText.
func NewLocalizedText_List(s *capnp.Segment, sz int32) (LocalizedText_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return LocalizedText_List{l}, err
}

func (s LocalizedText_List) At(i int) LocalizedText { return LocalizedText{s.List.Struct(i)} }

func (s LocalizedText_List) Set(i int, v LocalizedText) error { return s.List.SetStruct(i, v.Struct) }

// LocalizedText_Promise is a wrapper for a LocalizedText promised by a client call.
type LocalizedText_Promise struct{ *capnp.Pipeline }

func (p LocalizedText_Promise) Struct() (LocalizedText, error) {
	s, err := p.Pipeline.Struct()
	return LocalizedText{s}, err
}

type LocalizedText_Localization struct{ capnp.Struct }

// LocalizedText_Localization_TypeID is the unique identifier for the type LocalizedText_Localization.
const LocalizedText_Localization_TypeID = 0xa4f5ae06dd1b7791

func NewLocalizedText_Localization(s *capnp.Segment) (LocalizedText_Localization, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return LocalizedText_Localization{st}, err
}

func NewRootLocalizedText_Localization(s *capnp.Segment) (LocalizedText_Localization, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return LocalizedText_Localization{st}, err
}

func ReadRootLocalizedText_Localization(msg *capnp.Message) (LocalizedText_Localization, error) {
	root, err := msg.RootPtr()
	return LocalizedText_Localization{root.Struct()}, err
}

func (s LocalizedText_Localization) String() string {
	str, _ := text.Marshal(0xa4f5ae06dd1b7791, s.Struct)
	return str
}

func (s LocalizedText_Localization) Locale() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s LocalizedText_Localization) HasLocale() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s LocalizedText_Localization) LocaleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s LocalizedText_Localization) SetLocale(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s LocalizedText_Localization) Text() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s LocalizedText_Localization) HasText() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s LocalizedText_Localization) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s LocalizedText_Localization) SetText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// LocalizedText_Localization_List is a list of LocalizedText_Localization.
type LocalizedText_Localization_List struct{ capnp.List }

// NewLocalizedText_Localization creates a new list of LocalizedText_Localization.
func NewLocalizedText_Localization_List(s *capnp.Segment, sz int32) (LocalizedText_Localization_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return LocalizedText_Localization_List{l}, err
}

func (s LocalizedText_Localization_List) At(i int) LocalizedText_Localization {
	return LocalizedText_Localization{s.List.Struct(i)}
}

func (s LocalizedText_Localization_List) Set(i int, v LocalizedText_Localization) error {
	return s.List.SetStruct(i, v.Struct)
}

// LocalizedText_Localization_Promise is a wrapper for a LocalizedText_Localization promised by a client call.
type LocalizedText_Localization_Promise struct{ *capnp.Pipeline }

func (p LocalizedText_Localization_Promise) Struct() (LocalizedText_Localization, error) {
	s, err := p.Pipeline.Struct()
	return LocalizedText_Localization{s}, err
}

type Handle struct{ Client capnp.Client }

type Handle_Server interface {
}

func Handle_ServerToClient(s Handle_Server) Handle {
	c, _ := s.(server.Closer)
	return Handle{Client: server.New(Handle_Methods(nil, s), c)}
}

func Handle_Methods(methods []server.Method, s Handle_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 0)
	}

	return methods
}

type ByteStream struct{ Client capnp.Client }

func (c ByteStream) Write(ctx context.Context, params func(ByteStream_write_Params) error, opts ...capnp.CallOption) ByteStream_write_Results_Promise {
	if c.Client == nil {
		return ByteStream_write_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      0,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "write",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ByteStream_write_Params{Struct: s}) }
	}
	return ByteStream_write_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c ByteStream) Done(ctx context.Context, params func(ByteStream_done_Params) error, opts ...capnp.CallOption) ByteStream_done_Results_Promise {
	if c.Client == nil {
		return ByteStream_done_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      1,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "done",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ByteStream_done_Params{Struct: s}) }
	}
	return ByteStream_done_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c ByteStream) ExpectSize(ctx context.Context, params func(ByteStream_expectSize_Params) error, opts ...capnp.CallOption) ByteStream_expectSize_Results_Promise {
	if c.Client == nil {
		return ByteStream_expectSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      2,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "expectSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ByteStream_expectSize_Params{Struct: s}) }
	}
	return ByteStream_expectSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type ByteStream_Server interface {
	Write(ByteStream_write) error

	Done(ByteStream_done) error

	ExpectSize(ByteStream_expectSize) error
}

func ByteStream_ServerToClient(s ByteStream_Server) ByteStream {
	c, _ := s.(server.Closer)
	return ByteStream{Client: server.New(ByteStream_Methods(nil, s), c)}
}

func ByteStream_Methods(methods []server.Method, s ByteStream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      0,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "write",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ByteStream_write{c, opts, ByteStream_write_Params{Struct: p}, ByteStream_write_Results{Struct: r}}
			return s.Write(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      1,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "done",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ByteStream_done{c, opts, ByteStream_done_Params{Struct: p}, ByteStream_done_Results{Struct: r}}
			return s.Done(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xcd57387729cfe35f,
			MethodID:      2,
			InterfaceName: "util.capnp:ByteStream",
			MethodName:    "expectSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ByteStream_expectSize{c, opts, ByteStream_expectSize_Params{Struct: p}, ByteStream_expectSize_Results{Struct: r}}
			return s.ExpectSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// ByteStream_write holds the arguments for a server call to ByteStream.write.
type ByteStream_write struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ByteStream_write_Params
	Results ByteStream_write_Results
}

// ByteStream_done holds the arguments for a server call to ByteStream.done.
type ByteStream_done struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ByteStream_done_Params
	Results ByteStream_done_Results
}

// ByteStream_expectSize holds the arguments for a server call to ByteStream.expectSize.
type ByteStream_expectSize struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ByteStream_expectSize_Params
	Results ByteStream_expectSize_Results
}

type ByteStream_write_Params struct{ capnp.Struct }

// ByteStream_write_Params_TypeID is the unique identifier for the type ByteStream_write_Params.
const ByteStream_write_Params_TypeID = 0x97ed122121126ff2

func NewByteStream_write_Params(s *capnp.Segment) (ByteStream_write_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ByteStream_write_Params{st}, err
}

func NewRootByteStream_write_Params(s *capnp.Segment) (ByteStream_write_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ByteStream_write_Params{st}, err
}

func ReadRootByteStream_write_Params(msg *capnp.Message) (ByteStream_write_Params, error) {
	root, err := msg.RootPtr()
	return ByteStream_write_Params{root.Struct()}, err
}

func (s ByteStream_write_Params) String() string {
	str, _ := text.Marshal(0x97ed122121126ff2, s.Struct)
	return str
}

func (s ByteStream_write_Params) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s ByteStream_write_Params) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ByteStream_write_Params) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// ByteStream_write_Params_List is a list of ByteStream_write_Params.
type ByteStream_write_Params_List struct{ capnp.List }

// NewByteStream_write_Params creates a new list of ByteStream_write_Params.
func NewByteStream_write_Params_List(s *capnp.Segment, sz int32) (ByteStream_write_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ByteStream_write_Params_List{l}, err
}

func (s ByteStream_write_Params_List) At(i int) ByteStream_write_Params {
	return ByteStream_write_Params{s.List.Struct(i)}
}

func (s ByteStream_write_Params_List) Set(i int, v ByteStream_write_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_write_Params_Promise is a wrapper for a ByteStream_write_Params promised by a client call.
type ByteStream_write_Params_Promise struct{ *capnp.Pipeline }

func (p ByteStream_write_Params_Promise) Struct() (ByteStream_write_Params, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_write_Params{s}, err
}

type ByteStream_write_Results struct{ capnp.Struct }

// ByteStream_write_Results_TypeID is the unique identifier for the type ByteStream_write_Results.
const ByteStream_write_Results_TypeID = 0xecde2a9c6f3f84c9

func NewByteStream_write_Results(s *capnp.Segment) (ByteStream_write_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_write_Results{st}, err
}

func NewRootByteStream_write_Results(s *capnp.Segment) (ByteStream_write_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_write_Results{st}, err
}

func ReadRootByteStream_write_Results(msg *capnp.Message) (ByteStream_write_Results, error) {
	root, err := msg.RootPtr()
	return ByteStream_write_Results{root.Struct()}, err
}

func (s ByteStream_write_Results) String() string {
	str, _ := text.Marshal(0xecde2a9c6f3f84c9, s.Struct)
	return str
}

// ByteStream_write_Results_List is a list of ByteStream_write_Results.
type ByteStream_write_Results_List struct{ capnp.List }

// NewByteStream_write_Results creates a new list of ByteStream_write_Results.
func NewByteStream_write_Results_List(s *capnp.Segment, sz int32) (ByteStream_write_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ByteStream_write_Results_List{l}, err
}

func (s ByteStream_write_Results_List) At(i int) ByteStream_write_Results {
	return ByteStream_write_Results{s.List.Struct(i)}
}

func (s ByteStream_write_Results_List) Set(i int, v ByteStream_write_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_write_Results_Promise is a wrapper for a ByteStream_write_Results promised by a client call.
type ByteStream_write_Results_Promise struct{ *capnp.Pipeline }

func (p ByteStream_write_Results_Promise) Struct() (ByteStream_write_Results, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_write_Results{s}, err
}

type ByteStream_done_Params struct{ capnp.Struct }

// ByteStream_done_Params_TypeID is the unique identifier for the type ByteStream_done_Params.
const ByteStream_done_Params_TypeID = 0xbc1426493658b76e

func NewByteStream_done_Params(s *capnp.Segment) (ByteStream_done_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_done_Params{st}, err
}

func NewRootByteStream_done_Params(s *capnp.Segment) (ByteStream_done_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_done_Params{st}, err
}

func ReadRootByteStream_done_Params(msg *capnp.Message) (ByteStream_done_Params, error) {
	root, err := msg.RootPtr()
	return ByteStream_done_Params{root.Struct()}, err
}

func (s ByteStream_done_Params) String() string {
	str, _ := text.Marshal(0xbc1426493658b76e, s.Struct)
	return str
}

// ByteStream_done_Params_List is a list of ByteStream_done_Params.
type ByteStream_done_Params_List struct{ capnp.List }

// NewByteStream_done_Params creates a new list of ByteStream_done_Params.
func NewByteStream_done_Params_List(s *capnp.Segment, sz int32) (ByteStream_done_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ByteStream_done_Params_List{l}, err
}

func (s ByteStream_done_Params_List) At(i int) ByteStream_done_Params {
	return ByteStream_done_Params{s.List.Struct(i)}
}

func (s ByteStream_done_Params_List) Set(i int, v ByteStream_done_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_done_Params_Promise is a wrapper for a ByteStream_done_Params promised by a client call.
type ByteStream_done_Params_Promise struct{ *capnp.Pipeline }

func (p ByteStream_done_Params_Promise) Struct() (ByteStream_done_Params, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_done_Params{s}, err
}

type ByteStream_done_Results struct{ capnp.Struct }

// ByteStream_done_Results_TypeID is the unique identifier for the type ByteStream_done_Results.
const ByteStream_done_Results_TypeID = 0xd0d8d935ee30b219

func NewByteStream_done_Results(s *capnp.Segment) (ByteStream_done_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_done_Results{st}, err
}

func NewRootByteStream_done_Results(s *capnp.Segment) (ByteStream_done_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_done_Results{st}, err
}

func ReadRootByteStream_done_Results(msg *capnp.Message) (ByteStream_done_Results, error) {
	root, err := msg.RootPtr()
	return ByteStream_done_Results{root.Struct()}, err
}

func (s ByteStream_done_Results) String() string {
	str, _ := text.Marshal(0xd0d8d935ee30b219, s.Struct)
	return str
}

// ByteStream_done_Results_List is a list of ByteStream_done_Results.
type ByteStream_done_Results_List struct{ capnp.List }

// NewByteStream_done_Results creates a new list of ByteStream_done_Results.
func NewByteStream_done_Results_List(s *capnp.Segment, sz int32) (ByteStream_done_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ByteStream_done_Results_List{l}, err
}

func (s ByteStream_done_Results_List) At(i int) ByteStream_done_Results {
	return ByteStream_done_Results{s.List.Struct(i)}
}

func (s ByteStream_done_Results_List) Set(i int, v ByteStream_done_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_done_Results_Promise is a wrapper for a ByteStream_done_Results promised by a client call.
type ByteStream_done_Results_Promise struct{ *capnp.Pipeline }

func (p ByteStream_done_Results_Promise) Struct() (ByteStream_done_Results, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_done_Results{s}, err
}

type ByteStream_expectSize_Params struct{ capnp.Struct }

// ByteStream_expectSize_Params_TypeID is the unique identifier for the type ByteStream_expectSize_Params.
const ByteStream_expectSize_Params_TypeID = 0x8c9a3c7674c761d3

func NewByteStream_expectSize_Params(s *capnp.Segment) (ByteStream_expectSize_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return ByteStream_expectSize_Params{st}, err
}

func NewRootByteStream_expectSize_Params(s *capnp.Segment) (ByteStream_expectSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return ByteStream_expectSize_Params{st}, err
}

func ReadRootByteStream_expectSize_Params(msg *capnp.Message) (ByteStream_expectSize_Params, error) {
	root, err := msg.RootPtr()
	return ByteStream_expectSize_Params{root.Struct()}, err
}

func (s ByteStream_expectSize_Params) String() string {
	str, _ := text.Marshal(0x8c9a3c7674c761d3, s.Struct)
	return str
}

func (s ByteStream_expectSize_Params) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s ByteStream_expectSize_Params) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// ByteStream_expectSize_Params_List is a list of ByteStream_expectSize_Params.
type ByteStream_expectSize_Params_List struct{ capnp.List }

// NewByteStream_expectSize_Params creates a new list of ByteStream_expectSize_Params.
func NewByteStream_expectSize_Params_List(s *capnp.Segment, sz int32) (ByteStream_expectSize_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return ByteStream_expectSize_Params_List{l}, err
}

func (s ByteStream_expectSize_Params_List) At(i int) ByteStream_expectSize_Params {
	return ByteStream_expectSize_Params{s.List.Struct(i)}
}

func (s ByteStream_expectSize_Params_List) Set(i int, v ByteStream_expectSize_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_expectSize_Params_Promise is a wrapper for a ByteStream_expectSize_Params promised by a client call.
type ByteStream_expectSize_Params_Promise struct{ *capnp.Pipeline }

func (p ByteStream_expectSize_Params_Promise) Struct() (ByteStream_expectSize_Params, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_expectSize_Params{s}, err
}

type ByteStream_expectSize_Results struct{ capnp.Struct }

// ByteStream_expectSize_Results_TypeID is the unique identifier for the type ByteStream_expectSize_Results.
const ByteStream_expectSize_Results_TypeID = 0xf35749d82a51479b

func NewByteStream_expectSize_Results(s *capnp.Segment) (ByteStream_expectSize_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_expectSize_Results{st}, err
}

func NewRootByteStream_expectSize_Results(s *capnp.Segment) (ByteStream_expectSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ByteStream_expectSize_Results{st}, err
}

func ReadRootByteStream_expectSize_Results(msg *capnp.Message) (ByteStream_expectSize_Results, error) {
	root, err := msg.RootPtr()
	return ByteStream_expectSize_Results{root.Struct()}, err
}

func (s ByteStream_expectSize_Results) String() string {
	str, _ := text.Marshal(0xf35749d82a51479b, s.Struct)
	return str
}

// ByteStream_expectSize_Results_List is a list of ByteStream_expectSize_Results.
type ByteStream_expectSize_Results_List struct{ capnp.List }

// NewByteStream_expectSize_Results creates a new list of ByteStream_expectSize_Results.
func NewByteStream_expectSize_Results_List(s *capnp.Segment, sz int32) (ByteStream_expectSize_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ByteStream_expectSize_Results_List{l}, err
}

func (s ByteStream_expectSize_Results_List) At(i int) ByteStream_expectSize_Results {
	return ByteStream_expectSize_Results{s.List.Struct(i)}
}

func (s ByteStream_expectSize_Results_List) Set(i int, v ByteStream_expectSize_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ByteStream_expectSize_Results_Promise is a wrapper for a ByteStream_expectSize_Results promised by a client call.
type ByteStream_expectSize_Results_Promise struct{ *capnp.Pipeline }

func (p ByteStream_expectSize_Results_Promise) Struct() (ByteStream_expectSize_Results, error) {
	s, err := p.Pipeline.Struct()
	return ByteStream_expectSize_Results{s}, err
}

type Blob struct{ Client capnp.Client }

func (c Blob) GetSize(ctx context.Context, params func(Blob_getSize_Params) error, opts ...capnp.CallOption) Blob_getSize_Results_Promise {
	if c.Client == nil {
		return Blob_getSize_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      0,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "getSize",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Blob_getSize_Params{Struct: s}) }
	}
	return Blob_getSize_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Blob) WriteTo(ctx context.Context, params func(Blob_writeTo_Params) error, opts ...capnp.CallOption) Blob_writeTo_Results_Promise {
	if c.Client == nil {
		return Blob_writeTo_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      1,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "writeTo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Blob_writeTo_Params{Struct: s}) }
	}
	return Blob_writeTo_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Blob) GetSlice(ctx context.Context, params func(Blob_getSlice_Params) error, opts ...capnp.CallOption) Blob_getSlice_Results_Promise {
	if c.Client == nil {
		return Blob_getSlice_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      2,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "getSlice",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 16, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Blob_getSlice_Params{Struct: s}) }
	}
	return Blob_getSlice_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Blob_Server interface {
	GetSize(Blob_getSize) error

	WriteTo(Blob_writeTo) error

	GetSlice(Blob_getSlice) error
}

func Blob_ServerToClient(s Blob_Server) Blob {
	c, _ := s.(server.Closer)
	return Blob{Client: server.New(Blob_Methods(nil, s), c)}
}

func Blob_Methods(methods []server.Method, s Blob_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      0,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "getSize",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Blob_getSize{c, opts, Blob_getSize_Params{Struct: p}, Blob_getSize_Results{Struct: r}}
			return s.GetSize(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      1,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "writeTo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Blob_writeTo{c, opts, Blob_writeTo_Params{Struct: p}, Blob_writeTo_Results{Struct: r}}
			return s.WriteTo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe53527a75d90198f,
			MethodID:      2,
			InterfaceName: "util.capnp:Blob",
			MethodName:    "getSlice",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Blob_getSlice{c, opts, Blob_getSlice_Params{Struct: p}, Blob_getSlice_Results{Struct: r}}
			return s.GetSlice(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Blob_getSize holds the arguments for a server call to Blob.getSize.
type Blob_getSize struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Blob_getSize_Params
	Results Blob_getSize_Results
}

// Blob_writeTo holds the arguments for a server call to Blob.writeTo.
type Blob_writeTo struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Blob_writeTo_Params
	Results Blob_writeTo_Results
}

// Blob_getSlice holds the arguments for a server call to Blob.getSlice.
type Blob_getSlice struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Blob_getSlice_Params
	Results Blob_getSlice_Results
}

type Blob_getSize_Params struct{ capnp.Struct }

// Blob_getSize_Params_TypeID is the unique identifier for the type Blob_getSize_Params.
const Blob_getSize_Params_TypeID = 0x8ee5f62e1fab915d

func NewBlob_getSize_Params(s *capnp.Segment) (Blob_getSize_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Blob_getSize_Params{st}, err
}

func NewRootBlob_getSize_Params(s *capnp.Segment) (Blob_getSize_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Blob_getSize_Params{st}, err
}

func ReadRootBlob_getSize_Params(msg *capnp.Message) (Blob_getSize_Params, error) {
	root, err := msg.RootPtr()
	return Blob_getSize_Params{root.Struct()}, err
}

func (s Blob_getSize_Params) String() string {
	str, _ := text.Marshal(0x8ee5f62e1fab915d, s.Struct)
	return str
}

// Blob_getSize_Params_List is a list of Blob_getSize_Params.
type Blob_getSize_Params_List struct{ capnp.List }

// NewBlob_getSize_Params creates a new list of Blob_getSize_Params.
func NewBlob_getSize_Params_List(s *capnp.Segment, sz int32) (Blob_getSize_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Blob_getSize_Params_List{l}, err
}

func (s Blob_getSize_Params_List) At(i int) Blob_getSize_Params {
	return Blob_getSize_Params{s.List.Struct(i)}
}

func (s Blob_getSize_Params_List) Set(i int, v Blob_getSize_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_getSize_Params_Promise is a wrapper for a Blob_getSize_Params promised by a client call.
type Blob_getSize_Params_Promise struct{ *capnp.Pipeline }

func (p Blob_getSize_Params_Promise) Struct() (Blob_getSize_Params, error) {
	s, err := p.Pipeline.Struct()
	return Blob_getSize_Params{s}, err
}

type Blob_getSize_Results struct{ capnp.Struct }

// Blob_getSize_Results_TypeID is the unique identifier for the type Blob_getSize_Results.
const Blob_getSize_Results_TypeID = 0x8e48cb1497f3d6f4

func NewBlob_getSize_Results(s *capnp.Segment) (Blob_getSize_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Blob_getSize_Results{st}, err
}

func NewRootBlob_getSize_Results(s *capnp.Segment) (Blob_getSize_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Blob_getSize_Results{st}, err
}

func ReadRootBlob_getSize_Results(msg *capnp.Message) (Blob_getSize_Results, error) {
	root, err := msg.RootPtr()
	return Blob_getSize_Results{root.Struct()}, err
}

func (s Blob_getSize_Results) String() string {
	str, _ := text.Marshal(0x8e48cb1497f3d6f4, s.Struct)
	return str
}

func (s Blob_getSize_Results) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Blob_getSize_Results) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Blob_getSize_Results_List is a list of Blob_getSize_Results.
type Blob_getSize_Results_List struct{ capnp.List }

// NewBlob_getSize_Results creates a new list of Blob_getSize_Results.
func NewBlob_getSize_Results_List(s *capnp.Segment, sz int32) (Blob_getSize_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Blob_getSize_Results_List{l}, err
}

func (s Blob_getSize_Results_List) At(i int) Blob_getSize_Results {
	return Blob_getSize_Results{s.List.Struct(i)}
}

func (s Blob_getSize_Results_List) Set(i int, v Blob_getSize_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_getSize_Results_Promise is a wrapper for a Blob_getSize_Results promised by a client call.
type Blob_getSize_Results_Promise struct{ *capnp.Pipeline }

func (p Blob_getSize_Results_Promise) Struct() (Blob_getSize_Results, error) {
	s, err := p.Pipeline.Struct()
	return Blob_getSize_Results{s}, err
}

type Blob_writeTo_Params struct{ capnp.Struct }

// Blob_writeTo_Params_TypeID is the unique identifier for the type Blob_writeTo_Params.
const Blob_writeTo_Params_TypeID = 0x9f0719e9a9dccc4b

func NewBlob_writeTo_Params(s *capnp.Segment) (Blob_writeTo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Blob_writeTo_Params{st}, err
}

func NewRootBlob_writeTo_Params(s *capnp.Segment) (Blob_writeTo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Blob_writeTo_Params{st}, err
}

func ReadRootBlob_writeTo_Params(msg *capnp.Message) (Blob_writeTo_Params, error) {
	root, err := msg.RootPtr()
	return Blob_writeTo_Params{root.Struct()}, err
}

func (s Blob_writeTo_Params) String() string {
	str, _ := text.Marshal(0x9f0719e9a9dccc4b, s.Struct)
	return str
}

func (s Blob_writeTo_Params) Stream() ByteStream {
	p, _ := s.Struct.Ptr(0)
	return ByteStream{Client: p.Interface().Client()}
}

func (s Blob_writeTo_Params) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Blob_writeTo_Params) SetStream(v ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s Blob_writeTo_Params) StartAtOffset() uint64 {
	return s.Struct.Uint64(0)
}

func (s Blob_writeTo_Params) SetStartAtOffset(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Blob_writeTo_Params_List is a list of Blob_writeTo_Params.
type Blob_writeTo_Params_List struct{ capnp.List }

// NewBlob_writeTo_Params creates a new list of Blob_writeTo_Params.
func NewBlob_writeTo_Params_List(s *capnp.Segment, sz int32) (Blob_writeTo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Blob_writeTo_Params_List{l}, err
}

func (s Blob_writeTo_Params_List) At(i int) Blob_writeTo_Params {
	return Blob_writeTo_Params{s.List.Struct(i)}
}

func (s Blob_writeTo_Params_List) Set(i int, v Blob_writeTo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_writeTo_Params_Promise is a wrapper for a Blob_writeTo_Params promised by a client call.
type Blob_writeTo_Params_Promise struct{ *capnp.Pipeline }

func (p Blob_writeTo_Params_Promise) Struct() (Blob_writeTo_Params, error) {
	s, err := p.Pipeline.Struct()
	return Blob_writeTo_Params{s}, err
}

func (p Blob_writeTo_Params_Promise) Stream() ByteStream {
	return ByteStream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Blob_writeTo_Results struct{ capnp.Struct }

// Blob_writeTo_Results_TypeID is the unique identifier for the type Blob_writeTo_Results.
const Blob_writeTo_Results_TypeID = 0xdb3152bd3bc2aa40

func NewBlob_writeTo_Results(s *capnp.Segment) (Blob_writeTo_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Blob_writeTo_Results{st}, err
}

func NewRootBlob_writeTo_Results(s *capnp.Segment) (Blob_writeTo_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Blob_writeTo_Results{st}, err
}

func ReadRootBlob_writeTo_Results(msg *capnp.Message) (Blob_writeTo_Results, error) {
	root, err := msg.RootPtr()
	return Blob_writeTo_Results{root.Struct()}, err
}

func (s Blob_writeTo_Results) String() string {
	str, _ := text.Marshal(0xdb3152bd3bc2aa40, s.Struct)
	return str
}

func (s Blob_writeTo_Results) Handle() Handle {
	p, _ := s.Struct.Ptr(0)
	return Handle{Client: p.Interface().Client()}
}

func (s Blob_writeTo_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Blob_writeTo_Results) SetHandle(v Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Blob_writeTo_Results_List is a list of Blob_writeTo_Results.
type Blob_writeTo_Results_List struct{ capnp.List }

// NewBlob_writeTo_Results creates a new list of Blob_writeTo_Results.
func NewBlob_writeTo_Results_List(s *capnp.Segment, sz int32) (Blob_writeTo_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Blob_writeTo_Results_List{l}, err
}

func (s Blob_writeTo_Results_List) At(i int) Blob_writeTo_Results {
	return Blob_writeTo_Results{s.List.Struct(i)}
}

func (s Blob_writeTo_Results_List) Set(i int, v Blob_writeTo_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_writeTo_Results_Promise is a wrapper for a Blob_writeTo_Results promised by a client call.
type Blob_writeTo_Results_Promise struct{ *capnp.Pipeline }

func (p Blob_writeTo_Results_Promise) Struct() (Blob_writeTo_Results, error) {
	s, err := p.Pipeline.Struct()
	return Blob_writeTo_Results{s}, err
}

func (p Blob_writeTo_Results_Promise) Handle() Handle {
	return Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Blob_getSlice_Params struct{ capnp.Struct }

// Blob_getSlice_Params_TypeID is the unique identifier for the type Blob_getSlice_Params.
const Blob_getSlice_Params_TypeID = 0x8edb5f3937d96b8a

func NewBlob_getSlice_Params(s *capnp.Segment) (Blob_getSlice_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Blob_getSlice_Params{st}, err
}

func NewRootBlob_getSlice_Params(s *capnp.Segment) (Blob_getSlice_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Blob_getSlice_Params{st}, err
}

func ReadRootBlob_getSlice_Params(msg *capnp.Message) (Blob_getSlice_Params, error) {
	root, err := msg.RootPtr()
	return Blob_getSlice_Params{root.Struct()}, err
}

func (s Blob_getSlice_Params) String() string {
	str, _ := text.Marshal(0x8edb5f3937d96b8a, s.Struct)
	return str
}

func (s Blob_getSlice_Params) Offset() uint64 {
	return s.Struct.Uint64(0)
}

func (s Blob_getSlice_Params) SetOffset(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Blob_getSlice_Params) Size() uint32 {
	return s.Struct.Uint32(8)
}

func (s Blob_getSlice_Params) SetSize(v uint32) {
	s.Struct.SetUint32(8, v)
}

// Blob_getSlice_Params_List is a list of Blob_getSlice_Params.
type Blob_getSlice_Params_List struct{ capnp.List }

// NewBlob_getSlice_Params creates a new list of Blob_getSlice_Params.
func NewBlob_getSlice_Params_List(s *capnp.Segment, sz int32) (Blob_getSlice_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return Blob_getSlice_Params_List{l}, err
}

func (s Blob_getSlice_Params_List) At(i int) Blob_getSlice_Params {
	return Blob_getSlice_Params{s.List.Struct(i)}
}

func (s Blob_getSlice_Params_List) Set(i int, v Blob_getSlice_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_getSlice_Params_Promise is a wrapper for a Blob_getSlice_Params promised by a client call.
type Blob_getSlice_Params_Promise struct{ *capnp.Pipeline }

func (p Blob_getSlice_Params_Promise) Struct() (Blob_getSlice_Params, error) {
	s, err := p.Pipeline.Struct()
	return Blob_getSlice_Params{s}, err
}

type Blob_getSlice_Results struct{ capnp.Struct }

// Blob_getSlice_Results_TypeID is the unique identifier for the type Blob_getSlice_Results.
const Blob_getSlice_Results_TypeID = 0xc65caf9a2d389078

func NewBlob_getSlice_Results(s *capnp.Segment) (Blob_getSlice_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Blob_getSlice_Results{st}, err
}

func NewRootBlob_getSlice_Results(s *capnp.Segment) (Blob_getSlice_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Blob_getSlice_Results{st}, err
}

func ReadRootBlob_getSlice_Results(msg *capnp.Message) (Blob_getSlice_Results, error) {
	root, err := msg.RootPtr()
	return Blob_getSlice_Results{root.Struct()}, err
}

func (s Blob_getSlice_Results) String() string {
	str, _ := text.Marshal(0xc65caf9a2d389078, s.Struct)
	return str
}

func (s Blob_getSlice_Results) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Blob_getSlice_Results) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Blob_getSlice_Results) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// Blob_getSlice_Results_List is a list of Blob_getSlice_Results.
type Blob_getSlice_Results_List struct{ capnp.List }

// NewBlob_getSlice_Results creates a new list of Blob_getSlice_Results.
func NewBlob_getSlice_Results_List(s *capnp.Segment, sz int32) (Blob_getSlice_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Blob_getSlice_Results_List{l}, err
}

func (s Blob_getSlice_Results_List) At(i int) Blob_getSlice_Results {
	return Blob_getSlice_Results{s.List.Struct(i)}
}

func (s Blob_getSlice_Results_List) Set(i int, v Blob_getSlice_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Blob_getSlice_Results_Promise is a wrapper for a Blob_getSlice_Results promised by a client call.
type Blob_getSlice_Results_Promise struct{ *capnp.Pipeline }

func (p Blob_getSlice_Results_Promise) Struct() (Blob_getSlice_Results, error) {
	s, err := p.Pipeline.Struct()
	return Blob_getSlice_Results{s}, err
}

type Assignable struct{ Client capnp.Client }

func (c Assignable) Get(ctx context.Context, params func(Assignable_get_Params) error, opts ...capnp.CallOption) Assignable_get_Results_Promise {
	if c.Client == nil {
		return Assignable_get_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "get",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_get_Params{Struct: s}) }
	}
	return Assignable_get_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Assignable) AsGetter(ctx context.Context, params func(Assignable_asGetter_Params) error, opts ...capnp.CallOption) Assignable_asGetter_Results_Promise {
	if c.Client == nil {
		return Assignable_asGetter_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      1,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "asGetter",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_asGetter_Params{Struct: s}) }
	}
	return Assignable_asGetter_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Assignable) AsSetter(ctx context.Context, params func(Assignable_asSetter_Params) error, opts ...capnp.CallOption) Assignable_asSetter_Results_Promise {
	if c.Client == nil {
		return Assignable_asSetter_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      2,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "asSetter",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_asSetter_Params{Struct: s}) }
	}
	return Assignable_asSetter_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Assignable_Server interface {
	Get(Assignable_get) error

	AsGetter(Assignable_asGetter) error

	AsSetter(Assignable_asSetter) error
}

func Assignable_ServerToClient(s Assignable_Server) Assignable {
	c, _ := s.(server.Closer)
	return Assignable{Client: server.New(Assignable_Methods(nil, s), c)}
}

func Assignable_Methods(methods []server.Method, s Assignable_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "get",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_get{c, opts, Assignable_get_Params{Struct: p}, Assignable_get_Results{Struct: r}}
			return s.Get(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      1,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "asGetter",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_asGetter{c, opts, Assignable_asGetter_Params{Struct: p}, Assignable_asGetter_Results{Struct: r}}
			return s.AsGetter(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xeaf255b498229199,
			MethodID:      2,
			InterfaceName: "util.capnp:Assignable",
			MethodName:    "asSetter",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_asSetter{c, opts, Assignable_asSetter_Params{Struct: p}, Assignable_asSetter_Results{Struct: r}}
			return s.AsSetter(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Assignable_get holds the arguments for a server call to Assignable.get.
type Assignable_get struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_get_Params
	Results Assignable_get_Results
}

// Assignable_asGetter holds the arguments for a server call to Assignable.asGetter.
type Assignable_asGetter struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_asGetter_Params
	Results Assignable_asGetter_Results
}

// Assignable_asSetter holds the arguments for a server call to Assignable.asSetter.
type Assignable_asSetter struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_asSetter_Params
	Results Assignable_asSetter_Results
}

type Assignable_Getter struct{ Client capnp.Client }

func (c Assignable_Getter) Get(ctx context.Context, params func(Assignable_Getter_get_Params) error, opts ...capnp.CallOption) Assignable_Getter_get_Results_Promise {
	if c.Client == nil {
		return Assignable_Getter_get_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x80f2f65360d64224,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable.Getter",
			MethodName:    "get",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_Getter_get_Params{Struct: s}) }
	}
	return Assignable_Getter_get_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Assignable_Getter) Subscribe(ctx context.Context, params func(Assignable_Getter_subscribe_Params) error, opts ...capnp.CallOption) Assignable_Getter_subscribe_Results_Promise {
	if c.Client == nil {
		return Assignable_Getter_subscribe_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x80f2f65360d64224,
			MethodID:      1,
			InterfaceName: "util.capnp:Assignable.Getter",
			MethodName:    "subscribe",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_Getter_subscribe_Params{Struct: s}) }
	}
	return Assignable_Getter_subscribe_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Assignable_Getter_Server interface {
	Get(Assignable_Getter_get) error

	Subscribe(Assignable_Getter_subscribe) error
}

func Assignable_Getter_ServerToClient(s Assignable_Getter_Server) Assignable_Getter {
	c, _ := s.(server.Closer)
	return Assignable_Getter{Client: server.New(Assignable_Getter_Methods(nil, s), c)}
}

func Assignable_Getter_Methods(methods []server.Method, s Assignable_Getter_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x80f2f65360d64224,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable.Getter",
			MethodName:    "get",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_Getter_get{c, opts, Assignable_Getter_get_Params{Struct: p}, Assignable_Getter_get_Results{Struct: r}}
			return s.Get(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x80f2f65360d64224,
			MethodID:      1,
			InterfaceName: "util.capnp:Assignable.Getter",
			MethodName:    "subscribe",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_Getter_subscribe{c, opts, Assignable_Getter_subscribe_Params{Struct: p}, Assignable_Getter_subscribe_Results{Struct: r}}
			return s.Subscribe(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Assignable_Getter_get holds the arguments for a server call to Assignable_Getter.get.
type Assignable_Getter_get struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_Getter_get_Params
	Results Assignable_Getter_get_Results
}

// Assignable_Getter_subscribe holds the arguments for a server call to Assignable_Getter.subscribe.
type Assignable_Getter_subscribe struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_Getter_subscribe_Params
	Results Assignable_Getter_subscribe_Results
}

type Assignable_Getter_get_Params struct{ capnp.Struct }

// Assignable_Getter_get_Params_TypeID is the unique identifier for the type Assignable_Getter_get_Params.
const Assignable_Getter_get_Params_TypeID = 0xb19fdbd356844119

func NewAssignable_Getter_get_Params(s *capnp.Segment) (Assignable_Getter_get_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_Getter_get_Params{st}, err
}

func NewRootAssignable_Getter_get_Params(s *capnp.Segment) (Assignable_Getter_get_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_Getter_get_Params{st}, err
}

func ReadRootAssignable_Getter_get_Params(msg *capnp.Message) (Assignable_Getter_get_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_Getter_get_Params{root.Struct()}, err
}

func (s Assignable_Getter_get_Params) String() string {
	str, _ := text.Marshal(0xb19fdbd356844119, s.Struct)
	return str
}

// Assignable_Getter_get_Params_List is a list of Assignable_Getter_get_Params.
type Assignable_Getter_get_Params_List struct{ capnp.List }

// NewAssignable_Getter_get_Params creates a new list of Assignable_Getter_get_Params.
func NewAssignable_Getter_get_Params_List(s *capnp.Segment, sz int32) (Assignable_Getter_get_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_Getter_get_Params_List{l}, err
}

func (s Assignable_Getter_get_Params_List) At(i int) Assignable_Getter_get_Params {
	return Assignable_Getter_get_Params{s.List.Struct(i)}
}

func (s Assignable_Getter_get_Params_List) Set(i int, v Assignable_Getter_get_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Getter_get_Params_Promise is a wrapper for a Assignable_Getter_get_Params promised by a client call.
type Assignable_Getter_get_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_Getter_get_Params_Promise) Struct() (Assignable_Getter_get_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Getter_get_Params{s}, err
}

type Assignable_Getter_get_Results struct{ capnp.Struct }

// Assignable_Getter_get_Results_TypeID is the unique identifier for the type Assignable_Getter_get_Results.
const Assignable_Getter_get_Results_TypeID = 0x97ef2da226123492

func NewAssignable_Getter_get_Results(s *capnp.Segment) (Assignable_Getter_get_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_get_Results{st}, err
}

func NewRootAssignable_Getter_get_Results(s *capnp.Segment) (Assignable_Getter_get_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_get_Results{st}, err
}

func ReadRootAssignable_Getter_get_Results(msg *capnp.Message) (Assignable_Getter_get_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_Getter_get_Results{root.Struct()}, err
}

func (s Assignable_Getter_get_Results) String() string {
	str, _ := text.Marshal(0x97ef2da226123492, s.Struct)
	return str
}

func (s Assignable_Getter_get_Results) Value() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Assignable_Getter_get_Results) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Getter_get_Results) ValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Assignable_Getter_get_Results) SetValue(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Assignable_Getter_get_Results) SetValuePtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// Assignable_Getter_get_Results_List is a list of Assignable_Getter_get_Results.
type Assignable_Getter_get_Results_List struct{ capnp.List }

// NewAssignable_Getter_get_Results creates a new list of Assignable_Getter_get_Results.
func NewAssignable_Getter_get_Results_List(s *capnp.Segment, sz int32) (Assignable_Getter_get_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_Getter_get_Results_List{l}, err
}

func (s Assignable_Getter_get_Results_List) At(i int) Assignable_Getter_get_Results {
	return Assignable_Getter_get_Results{s.List.Struct(i)}
}

func (s Assignable_Getter_get_Results_List) Set(i int, v Assignable_Getter_get_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Getter_get_Results_Promise is a wrapper for a Assignable_Getter_get_Results promised by a client call.
type Assignable_Getter_get_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_Getter_get_Results_Promise) Struct() (Assignable_Getter_get_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Getter_get_Results{s}, err
}

func (p Assignable_Getter_get_Results_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Assignable_Getter_subscribe_Params struct{ capnp.Struct }

// Assignable_Getter_subscribe_Params_TypeID is the unique identifier for the type Assignable_Getter_subscribe_Params.
const Assignable_Getter_subscribe_Params_TypeID = 0xf02783ef982ecea9

func NewAssignable_Getter_subscribe_Params(s *capnp.Segment) (Assignable_Getter_subscribe_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_subscribe_Params{st}, err
}

func NewRootAssignable_Getter_subscribe_Params(s *capnp.Segment) (Assignable_Getter_subscribe_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_subscribe_Params{st}, err
}

func ReadRootAssignable_Getter_subscribe_Params(msg *capnp.Message) (Assignable_Getter_subscribe_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_Getter_subscribe_Params{root.Struct()}, err
}

func (s Assignable_Getter_subscribe_Params) String() string {
	str, _ := text.Marshal(0xf02783ef982ecea9, s.Struct)
	return str
}

func (s Assignable_Getter_subscribe_Params) Setter() Assignable_Setter {
	p, _ := s.Struct.Ptr(0)
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_Getter_subscribe_Params) HasSetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Getter_subscribe_Params) SetSetter(v Assignable_Setter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Assignable_Getter_subscribe_Params_List is a list of Assignable_Getter_subscribe_Params.
type Assignable_Getter_subscribe_Params_List struct{ capnp.List }

// NewAssignable_Getter_subscribe_Params creates a new list of Assignable_Getter_subscribe_Params.
func NewAssignable_Getter_subscribe_Params_List(s *capnp.Segment, sz int32) (Assignable_Getter_subscribe_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_Getter_subscribe_Params_List{l}, err
}

func (s Assignable_Getter_subscribe_Params_List) At(i int) Assignable_Getter_subscribe_Params {
	return Assignable_Getter_subscribe_Params{s.List.Struct(i)}
}

func (s Assignable_Getter_subscribe_Params_List) Set(i int, v Assignable_Getter_subscribe_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Getter_subscribe_Params_Promise is a wrapper for a Assignable_Getter_subscribe_Params promised by a client call.
type Assignable_Getter_subscribe_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_Getter_subscribe_Params_Promise) Struct() (Assignable_Getter_subscribe_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Getter_subscribe_Params{s}, err
}

func (p Assignable_Getter_subscribe_Params_Promise) Setter() Assignable_Setter {
	return Assignable_Setter{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Assignable_Getter_subscribe_Results struct{ capnp.Struct }

// Assignable_Getter_subscribe_Results_TypeID is the unique identifier for the type Assignable_Getter_subscribe_Results.
const Assignable_Getter_subscribe_Results_TypeID = 0x84e0f802c9af605b

func NewAssignable_Getter_subscribe_Results(s *capnp.Segment) (Assignable_Getter_subscribe_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_subscribe_Results{st}, err
}

func NewRootAssignable_Getter_subscribe_Results(s *capnp.Segment) (Assignable_Getter_subscribe_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Getter_subscribe_Results{st}, err
}

func ReadRootAssignable_Getter_subscribe_Results(msg *capnp.Message) (Assignable_Getter_subscribe_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_Getter_subscribe_Results{root.Struct()}, err
}

func (s Assignable_Getter_subscribe_Results) String() string {
	str, _ := text.Marshal(0x84e0f802c9af605b, s.Struct)
	return str
}

func (s Assignable_Getter_subscribe_Results) Handle() Handle {
	p, _ := s.Struct.Ptr(0)
	return Handle{Client: p.Interface().Client()}
}

func (s Assignable_Getter_subscribe_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Getter_subscribe_Results) SetHandle(v Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Assignable_Getter_subscribe_Results_List is a list of Assignable_Getter_subscribe_Results.
type Assignable_Getter_subscribe_Results_List struct{ capnp.List }

// NewAssignable_Getter_subscribe_Results creates a new list of Assignable_Getter_subscribe_Results.
func NewAssignable_Getter_subscribe_Results_List(s *capnp.Segment, sz int32) (Assignable_Getter_subscribe_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_Getter_subscribe_Results_List{l}, err
}

func (s Assignable_Getter_subscribe_Results_List) At(i int) Assignable_Getter_subscribe_Results {
	return Assignable_Getter_subscribe_Results{s.List.Struct(i)}
}

func (s Assignable_Getter_subscribe_Results_List) Set(i int, v Assignable_Getter_subscribe_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Getter_subscribe_Results_Promise is a wrapper for a Assignable_Getter_subscribe_Results promised by a client call.
type Assignable_Getter_subscribe_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_Getter_subscribe_Results_Promise) Struct() (Assignable_Getter_subscribe_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Getter_subscribe_Results{s}, err
}

func (p Assignable_Getter_subscribe_Results_Promise) Handle() Handle {
	return Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Assignable_Setter struct{ Client capnp.Client }

func (c Assignable_Setter) Set(ctx context.Context, params func(Assignable_Setter_set_Params) error, opts ...capnp.CallOption) Assignable_Setter_set_Results_Promise {
	if c.Client == nil {
		return Assignable_Setter_set_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd5256a3f93589d2f,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable.Setter",
			MethodName:    "set",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Assignable_Setter_set_Params{Struct: s}) }
	}
	return Assignable_Setter_set_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Assignable_Setter_Server interface {
	Set(Assignable_Setter_set) error
}

func Assignable_Setter_ServerToClient(s Assignable_Setter_Server) Assignable_Setter {
	c, _ := s.(server.Closer)
	return Assignable_Setter{Client: server.New(Assignable_Setter_Methods(nil, s), c)}
}

func Assignable_Setter_Methods(methods []server.Method, s Assignable_Setter_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd5256a3f93589d2f,
			MethodID:      0,
			InterfaceName: "util.capnp:Assignable.Setter",
			MethodName:    "set",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Assignable_Setter_set{c, opts, Assignable_Setter_set_Params{Struct: p}, Assignable_Setter_set_Results{Struct: r}}
			return s.Set(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// Assignable_Setter_set holds the arguments for a server call to Assignable_Setter.set.
type Assignable_Setter_set struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Assignable_Setter_set_Params
	Results Assignable_Setter_set_Results
}

type Assignable_Setter_set_Params struct{ capnp.Struct }

// Assignable_Setter_set_Params_TypeID is the unique identifier for the type Assignable_Setter_set_Params.
const Assignable_Setter_set_Params_TypeID = 0x98d0372787b787d1

func NewAssignable_Setter_set_Params(s *capnp.Segment) (Assignable_Setter_set_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Setter_set_Params{st}, err
}

func NewRootAssignable_Setter_set_Params(s *capnp.Segment) (Assignable_Setter_set_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_Setter_set_Params{st}, err
}

func ReadRootAssignable_Setter_set_Params(msg *capnp.Message) (Assignable_Setter_set_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_Setter_set_Params{root.Struct()}, err
}

func (s Assignable_Setter_set_Params) String() string {
	str, _ := text.Marshal(0x98d0372787b787d1, s.Struct)
	return str
}

func (s Assignable_Setter_set_Params) Value() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Assignable_Setter_set_Params) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_Setter_set_Params) ValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Assignable_Setter_set_Params) SetValue(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Assignable_Setter_set_Params) SetValuePtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// Assignable_Setter_set_Params_List is a list of Assignable_Setter_set_Params.
type Assignable_Setter_set_Params_List struct{ capnp.List }

// NewAssignable_Setter_set_Params creates a new list of Assignable_Setter_set_Params.
func NewAssignable_Setter_set_Params_List(s *capnp.Segment, sz int32) (Assignable_Setter_set_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_Setter_set_Params_List{l}, err
}

func (s Assignable_Setter_set_Params_List) At(i int) Assignable_Setter_set_Params {
	return Assignable_Setter_set_Params{s.List.Struct(i)}
}

func (s Assignable_Setter_set_Params_List) Set(i int, v Assignable_Setter_set_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Setter_set_Params_Promise is a wrapper for a Assignable_Setter_set_Params promised by a client call.
type Assignable_Setter_set_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_Setter_set_Params_Promise) Struct() (Assignable_Setter_set_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Setter_set_Params{s}, err
}

func (p Assignable_Setter_set_Params_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Assignable_Setter_set_Results struct{ capnp.Struct }

// Assignable_Setter_set_Results_TypeID is the unique identifier for the type Assignable_Setter_set_Results.
const Assignable_Setter_set_Results_TypeID = 0xdbfbb635d3e6abab

func NewAssignable_Setter_set_Results(s *capnp.Segment) (Assignable_Setter_set_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_Setter_set_Results{st}, err
}

func NewRootAssignable_Setter_set_Results(s *capnp.Segment) (Assignable_Setter_set_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_Setter_set_Results{st}, err
}

func ReadRootAssignable_Setter_set_Results(msg *capnp.Message) (Assignable_Setter_set_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_Setter_set_Results{root.Struct()}, err
}

func (s Assignable_Setter_set_Results) String() string {
	str, _ := text.Marshal(0xdbfbb635d3e6abab, s.Struct)
	return str
}

// Assignable_Setter_set_Results_List is a list of Assignable_Setter_set_Results.
type Assignable_Setter_set_Results_List struct{ capnp.List }

// NewAssignable_Setter_set_Results creates a new list of Assignable_Setter_set_Results.
func NewAssignable_Setter_set_Results_List(s *capnp.Segment, sz int32) (Assignable_Setter_set_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_Setter_set_Results_List{l}, err
}

func (s Assignable_Setter_set_Results_List) At(i int) Assignable_Setter_set_Results {
	return Assignable_Setter_set_Results{s.List.Struct(i)}
}

func (s Assignable_Setter_set_Results_List) Set(i int, v Assignable_Setter_set_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_Setter_set_Results_Promise is a wrapper for a Assignable_Setter_set_Results promised by a client call.
type Assignable_Setter_set_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_Setter_set_Results_Promise) Struct() (Assignable_Setter_set_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_Setter_set_Results{s}, err
}

type Assignable_get_Params struct{ capnp.Struct }

// Assignable_get_Params_TypeID is the unique identifier for the type Assignable_get_Params.
const Assignable_get_Params_TypeID = 0xbbfd27b5d2515662

func NewAssignable_get_Params(s *capnp.Segment) (Assignable_get_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_get_Params{st}, err
}

func NewRootAssignable_get_Params(s *capnp.Segment) (Assignable_get_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_get_Params{st}, err
}

func ReadRootAssignable_get_Params(msg *capnp.Message) (Assignable_get_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_get_Params{root.Struct()}, err
}

func (s Assignable_get_Params) String() string {
	str, _ := text.Marshal(0xbbfd27b5d2515662, s.Struct)
	return str
}

// Assignable_get_Params_List is a list of Assignable_get_Params.
type Assignable_get_Params_List struct{ capnp.List }

// NewAssignable_get_Params creates a new list of Assignable_get_Params.
func NewAssignable_get_Params_List(s *capnp.Segment, sz int32) (Assignable_get_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_get_Params_List{l}, err
}

func (s Assignable_get_Params_List) At(i int) Assignable_get_Params {
	return Assignable_get_Params{s.List.Struct(i)}
}

func (s Assignable_get_Params_List) Set(i int, v Assignable_get_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_get_Params_Promise is a wrapper for a Assignable_get_Params promised by a client call.
type Assignable_get_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_get_Params_Promise) Struct() (Assignable_get_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_get_Params{s}, err
}

type Assignable_get_Results struct{ capnp.Struct }

// Assignable_get_Results_TypeID is the unique identifier for the type Assignable_get_Results.
const Assignable_get_Results_TypeID = 0xb351b437cd426a4f

func NewAssignable_get_Results(s *capnp.Segment) (Assignable_get_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Assignable_get_Results{st}, err
}

func NewRootAssignable_get_Results(s *capnp.Segment) (Assignable_get_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Assignable_get_Results{st}, err
}

func ReadRootAssignable_get_Results(msg *capnp.Message) (Assignable_get_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_get_Results{root.Struct()}, err
}

func (s Assignable_get_Results) String() string {
	str, _ := text.Marshal(0xb351b437cd426a4f, s.Struct)
	return str
}

func (s Assignable_get_Results) Value() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Assignable_get_Results) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_get_Results) ValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Assignable_get_Results) SetValue(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Assignable_get_Results) SetValuePtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s Assignable_get_Results) Setter() Assignable_Setter {
	p, _ := s.Struct.Ptr(1)
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_get_Results) HasSetter() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Assignable_get_Results) SetSetter(v Assignable_Setter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// Assignable_get_Results_List is a list of Assignable_get_Results.
type Assignable_get_Results_List struct{ capnp.List }

// NewAssignable_get_Results creates a new list of Assignable_get_Results.
func NewAssignable_get_Results_List(s *capnp.Segment, sz int32) (Assignable_get_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Assignable_get_Results_List{l}, err
}

func (s Assignable_get_Results_List) At(i int) Assignable_get_Results {
	return Assignable_get_Results{s.List.Struct(i)}
}

func (s Assignable_get_Results_List) Set(i int, v Assignable_get_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_get_Results_Promise is a wrapper for a Assignable_get_Results promised by a client call.
type Assignable_get_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_get_Results_Promise) Struct() (Assignable_get_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_get_Results{s}, err
}

func (p Assignable_get_Results_Promise) Value() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p Assignable_get_Results_Promise) Setter() Assignable_Setter {
	return Assignable_Setter{Client: p.Pipeline.GetPipeline(1).Client()}
}

type Assignable_asGetter_Params struct{ capnp.Struct }

// Assignable_asGetter_Params_TypeID is the unique identifier for the type Assignable_asGetter_Params.
const Assignable_asGetter_Params_TypeID = 0xf907945b872b26cf

func NewAssignable_asGetter_Params(s *capnp.Segment) (Assignable_asGetter_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_asGetter_Params{st}, err
}

func NewRootAssignable_asGetter_Params(s *capnp.Segment) (Assignable_asGetter_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_asGetter_Params{st}, err
}

func ReadRootAssignable_asGetter_Params(msg *capnp.Message) (Assignable_asGetter_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_asGetter_Params{root.Struct()}, err
}

func (s Assignable_asGetter_Params) String() string {
	str, _ := text.Marshal(0xf907945b872b26cf, s.Struct)
	return str
}

// Assignable_asGetter_Params_List is a list of Assignable_asGetter_Params.
type Assignable_asGetter_Params_List struct{ capnp.List }

// NewAssignable_asGetter_Params creates a new list of Assignable_asGetter_Params.
func NewAssignable_asGetter_Params_List(s *capnp.Segment, sz int32) (Assignable_asGetter_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_asGetter_Params_List{l}, err
}

func (s Assignable_asGetter_Params_List) At(i int) Assignable_asGetter_Params {
	return Assignable_asGetter_Params{s.List.Struct(i)}
}

func (s Assignable_asGetter_Params_List) Set(i int, v Assignable_asGetter_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_asGetter_Params_Promise is a wrapper for a Assignable_asGetter_Params promised by a client call.
type Assignable_asGetter_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_asGetter_Params_Promise) Struct() (Assignable_asGetter_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_asGetter_Params{s}, err
}

type Assignable_asGetter_Results struct{ capnp.Struct }

// Assignable_asGetter_Results_TypeID is the unique identifier for the type Assignable_asGetter_Results.
const Assignable_asGetter_Results_TypeID = 0x8c3d547ef2930e96

func NewAssignable_asGetter_Results(s *capnp.Segment) (Assignable_asGetter_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_asGetter_Results{st}, err
}

func NewRootAssignable_asGetter_Results(s *capnp.Segment) (Assignable_asGetter_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_asGetter_Results{st}, err
}

func ReadRootAssignable_asGetter_Results(msg *capnp.Message) (Assignable_asGetter_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_asGetter_Results{root.Struct()}, err
}

func (s Assignable_asGetter_Results) String() string {
	str, _ := text.Marshal(0x8c3d547ef2930e96, s.Struct)
	return str
}

func (s Assignable_asGetter_Results) Getter() Assignable_Getter {
	p, _ := s.Struct.Ptr(0)
	return Assignable_Getter{Client: p.Interface().Client()}
}

func (s Assignable_asGetter_Results) HasGetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_asGetter_Results) SetGetter(v Assignable_Getter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Assignable_asGetter_Results_List is a list of Assignable_asGetter_Results.
type Assignable_asGetter_Results_List struct{ capnp.List }

// NewAssignable_asGetter_Results creates a new list of Assignable_asGetter_Results.
func NewAssignable_asGetter_Results_List(s *capnp.Segment, sz int32) (Assignable_asGetter_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_asGetter_Results_List{l}, err
}

func (s Assignable_asGetter_Results_List) At(i int) Assignable_asGetter_Results {
	return Assignable_asGetter_Results{s.List.Struct(i)}
}

func (s Assignable_asGetter_Results_List) Set(i int, v Assignable_asGetter_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_asGetter_Results_Promise is a wrapper for a Assignable_asGetter_Results promised by a client call.
type Assignable_asGetter_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_asGetter_Results_Promise) Struct() (Assignable_asGetter_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_asGetter_Results{s}, err
}

func (p Assignable_asGetter_Results_Promise) Getter() Assignable_Getter {
	return Assignable_Getter{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Assignable_asSetter_Params struct{ capnp.Struct }

// Assignable_asSetter_Params_TypeID is the unique identifier for the type Assignable_asSetter_Params.
const Assignable_asSetter_Params_TypeID = 0xa01f603357f3b349

func NewAssignable_asSetter_Params(s *capnp.Segment) (Assignable_asSetter_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_asSetter_Params{st}, err
}

func NewRootAssignable_asSetter_Params(s *capnp.Segment) (Assignable_asSetter_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Assignable_asSetter_Params{st}, err
}

func ReadRootAssignable_asSetter_Params(msg *capnp.Message) (Assignable_asSetter_Params, error) {
	root, err := msg.RootPtr()
	return Assignable_asSetter_Params{root.Struct()}, err
}

func (s Assignable_asSetter_Params) String() string {
	str, _ := text.Marshal(0xa01f603357f3b349, s.Struct)
	return str
}

// Assignable_asSetter_Params_List is a list of Assignable_asSetter_Params.
type Assignable_asSetter_Params_List struct{ capnp.List }

// NewAssignable_asSetter_Params creates a new list of Assignable_asSetter_Params.
func NewAssignable_asSetter_Params_List(s *capnp.Segment, sz int32) (Assignable_asSetter_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Assignable_asSetter_Params_List{l}, err
}

func (s Assignable_asSetter_Params_List) At(i int) Assignable_asSetter_Params {
	return Assignable_asSetter_Params{s.List.Struct(i)}
}

func (s Assignable_asSetter_Params_List) Set(i int, v Assignable_asSetter_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_asSetter_Params_Promise is a wrapper for a Assignable_asSetter_Params promised by a client call.
type Assignable_asSetter_Params_Promise struct{ *capnp.Pipeline }

func (p Assignable_asSetter_Params_Promise) Struct() (Assignable_asSetter_Params, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_asSetter_Params{s}, err
}

type Assignable_asSetter_Results struct{ capnp.Struct }

// Assignable_asSetter_Results_TypeID is the unique identifier for the type Assignable_asSetter_Results.
const Assignable_asSetter_Results_TypeID = 0xc6cbc10181c4f397

func NewAssignable_asSetter_Results(s *capnp.Segment) (Assignable_asSetter_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_asSetter_Results{st}, err
}

func NewRootAssignable_asSetter_Results(s *capnp.Segment) (Assignable_asSetter_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Assignable_asSetter_Results{st}, err
}

func ReadRootAssignable_asSetter_Results(msg *capnp.Message) (Assignable_asSetter_Results, error) {
	root, err := msg.RootPtr()
	return Assignable_asSetter_Results{root.Struct()}, err
}

func (s Assignable_asSetter_Results) String() string {
	str, _ := text.Marshal(0xc6cbc10181c4f397, s.Struct)
	return str
}

func (s Assignable_asSetter_Results) Setter() Assignable_Setter {
	p, _ := s.Struct.Ptr(0)
	return Assignable_Setter{Client: p.Interface().Client()}
}

func (s Assignable_asSetter_Results) HasSetter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Assignable_asSetter_Results) SetSetter(v Assignable_Setter) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Assignable_asSetter_Results_List is a list of Assignable_asSetter_Results.
type Assignable_asSetter_Results_List struct{ capnp.List }

// NewAssignable_asSetter_Results creates a new list of Assignable_asSetter_Results.
func NewAssignable_asSetter_Results_List(s *capnp.Segment, sz int32) (Assignable_asSetter_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Assignable_asSetter_Results_List{l}, err
}

func (s Assignable_asSetter_Results_List) At(i int) Assignable_asSetter_Results {
	return Assignable_asSetter_Results{s.List.Struct(i)}
}

func (s Assignable_asSetter_Results_List) Set(i int, v Assignable_asSetter_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Assignable_asSetter_Results_Promise is a wrapper for a Assignable_asSetter_Results promised by a client call.
type Assignable_asSetter_Results_Promise struct{ *capnp.Pipeline }

func (p Assignable_asSetter_Results_Promise) Struct() (Assignable_asSetter_Results, error) {
	s, err := p.Pipeline.Struct()
	return Assignable_asSetter_Results{s}, err
}

func (p Assignable_asSetter_Results_Promise) Setter() Assignable_Setter {
	return Assignable_Setter{Client: p.Pipeline.GetPipeline(0).Client()}
}

type StaticAsset struct{ Client capnp.Client }

func (c StaticAsset) GetUrl(ctx context.Context, params func(StaticAsset_getUrl_Params) error, opts ...capnp.CallOption) StaticAsset_getUrl_Results_Promise {
	if c.Client == nil {
		return StaticAsset_getUrl_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xfabb5e621fa9a23f,
			MethodID:      0,
			InterfaceName: "util.capnp:StaticAsset",
			MethodName:    "getUrl",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(StaticAsset_getUrl_Params{Struct: s}) }
	}
	return StaticAsset_getUrl_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type StaticAsset_Server interface {
	GetUrl(StaticAsset_getUrl) error
}

func StaticAsset_ServerToClient(s StaticAsset_Server) StaticAsset {
	c, _ := s.(server.Closer)
	return StaticAsset{Client: server.New(StaticAsset_Methods(nil, s), c)}
}

func StaticAsset_Methods(methods []server.Method, s StaticAsset_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xfabb5e621fa9a23f,
			MethodID:      0,
			InterfaceName: "util.capnp:StaticAsset",
			MethodName:    "getUrl",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := StaticAsset_getUrl{c, opts, StaticAsset_getUrl_Params{Struct: p}, StaticAsset_getUrl_Results{Struct: r}}
			return s.GetUrl(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 1},
	})

	return methods
}

// StaticAsset_getUrl holds the arguments for a server call to StaticAsset.getUrl.
type StaticAsset_getUrl struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  StaticAsset_getUrl_Params
	Results StaticAsset_getUrl_Results
}

type StaticAsset_Protocol uint16

// Values of StaticAsset_Protocol.
const (
	StaticAsset_Protocol_https StaticAsset_Protocol = 0
	StaticAsset_Protocol_http  StaticAsset_Protocol = 1
)

// String returns the enum's constant name.
func (c StaticAsset_Protocol) String() string {
	switch c {
	case StaticAsset_Protocol_https:
		return "https"
	case StaticAsset_Protocol_http:
		return "http"

	default:
		return ""
	}
}

// StaticAsset_ProtocolFromString returns the enum value with a name,
// or the zero value if there's no such value.
func StaticAsset_ProtocolFromString(c string) StaticAsset_Protocol {
	switch c {
	case "https":
		return StaticAsset_Protocol_https
	case "http":
		return StaticAsset_Protocol_http

	default:
		return 0
	}
}

type StaticAsset_Protocol_List struct{ capnp.List }

func NewStaticAsset_Protocol_List(s *capnp.Segment, sz int32) (StaticAsset_Protocol_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return StaticAsset_Protocol_List{l.List}, err
}

func (l StaticAsset_Protocol_List) At(i int) StaticAsset_Protocol {
	ul := capnp.UInt16List{List: l.List}
	return StaticAsset_Protocol(ul.At(i))
}

func (l StaticAsset_Protocol_List) Set(i int, v StaticAsset_Protocol) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type StaticAsset_getUrl_Params struct{ capnp.Struct }

// StaticAsset_getUrl_Params_TypeID is the unique identifier for the type StaticAsset_getUrl_Params.
const StaticAsset_getUrl_Params_TypeID = 0xa75ecf12570b2966

func NewStaticAsset_getUrl_Params(s *capnp.Segment) (StaticAsset_getUrl_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return StaticAsset_getUrl_Params{st}, err
}

func NewRootStaticAsset_getUrl_Params(s *capnp.Segment) (StaticAsset_getUrl_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return StaticAsset_getUrl_Params{st}, err
}

func ReadRootStaticAsset_getUrl_Params(msg *capnp.Message) (StaticAsset_getUrl_Params, error) {
	root, err := msg.RootPtr()
	return StaticAsset_getUrl_Params{root.Struct()}, err
}

func (s StaticAsset_getUrl_Params) String() string {
	str, _ := text.Marshal(0xa75ecf12570b2966, s.Struct)
	return str
}

// StaticAsset_getUrl_Params_List is a list of StaticAsset_getUrl_Params.
type StaticAsset_getUrl_Params_List struct{ capnp.List }

// NewStaticAsset_getUrl_Params creates a new list of StaticAsset_getUrl_Params.
func NewStaticAsset_getUrl_Params_List(s *capnp.Segment, sz int32) (StaticAsset_getUrl_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return StaticAsset_getUrl_Params_List{l}, err
}

func (s StaticAsset_getUrl_Params_List) At(i int) StaticAsset_getUrl_Params {
	return StaticAsset_getUrl_Params{s.List.Struct(i)}
}

func (s StaticAsset_getUrl_Params_List) Set(i int, v StaticAsset_getUrl_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// StaticAsset_getUrl_Params_Promise is a wrapper for a StaticAsset_getUrl_Params promised by a client call.
type StaticAsset_getUrl_Params_Promise struct{ *capnp.Pipeline }

func (p StaticAsset_getUrl_Params_Promise) Struct() (StaticAsset_getUrl_Params, error) {
	s, err := p.Pipeline.Struct()
	return StaticAsset_getUrl_Params{s}, err
}

type StaticAsset_getUrl_Results struct{ capnp.Struct }

// StaticAsset_getUrl_Results_TypeID is the unique identifier for the type StaticAsset_getUrl_Results.
const StaticAsset_getUrl_Results_TypeID = 0xa5c3aa75d6b648e2

func NewStaticAsset_getUrl_Results(s *capnp.Segment) (StaticAsset_getUrl_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return StaticAsset_getUrl_Results{st}, err
}

func NewRootStaticAsset_getUrl_Results(s *capnp.Segment) (StaticAsset_getUrl_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return StaticAsset_getUrl_Results{st}, err
}

func ReadRootStaticAsset_getUrl_Results(msg *capnp.Message) (StaticAsset_getUrl_Results, error) {
	root, err := msg.RootPtr()
	return StaticAsset_getUrl_Results{root.Struct()}, err
}

func (s StaticAsset_getUrl_Results) String() string {
	str, _ := text.Marshal(0xa5c3aa75d6b648e2, s.Struct)
	return str
}

func (s StaticAsset_getUrl_Results) Protocol() StaticAsset_Protocol {
	return StaticAsset_Protocol(s.Struct.Uint16(0))
}

func (s StaticAsset_getUrl_Results) SetProtocol(v StaticAsset_Protocol) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s StaticAsset_getUrl_Results) HostPath() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s StaticAsset_getUrl_Results) HasHostPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s StaticAsset_getUrl_Results) HostPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s StaticAsset_getUrl_Results) SetHostPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// StaticAsset_getUrl_Results_List is a list of StaticAsset_getUrl_Results.
type StaticAsset_getUrl_Results_List struct{ capnp.List }

// NewStaticAsset_getUrl_Results creates a new list of StaticAsset_getUrl_Results.
func NewStaticAsset_getUrl_Results_List(s *capnp.Segment, sz int32) (StaticAsset_getUrl_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return StaticAsset_getUrl_Results_List{l}, err
}

func (s StaticAsset_getUrl_Results_List) At(i int) StaticAsset_getUrl_Results {
	return StaticAsset_getUrl_Results{s.List.Struct(i)}
}

func (s StaticAsset_getUrl_Results_List) Set(i int, v StaticAsset_getUrl_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// StaticAsset_getUrl_Results_Promise is a wrapper for a StaticAsset_getUrl_Results promised by a client call.
type StaticAsset_getUrl_Results_Promise struct{ *capnp.Pipeline }

func (p StaticAsset_getUrl_Results_Promise) Struct() (StaticAsset_getUrl_Results, error) {
	s, err := p.Pipeline.Struct()
	return StaticAsset_getUrl_Results{s}, err
}

const schema_ecd50d792c3d9992 = "x\xda\xacX}l\x14\xd7\x11\x9f\xd9\xbb\xcb\xde\xa7\xcf" +
	"\xcfk\x059mu2\xb2qp\xf1\x15\xdbP\xe3+" +
	"\xe8|W\xb5\x80\x03\xca\xad? iB\xe2\xb5\xbd\x98" +
	"#g\x9f{\xbb\x17\xb0\xab6\xa5\xa5Bi\xa34\x90" +
	"\xa4\x86\xb4Q\x12\xd4*\"D\x8dB\xd2 5\xed\x1f" +
	"\xfd\x0a\xa2)%\x14\x19aQ\xaa\x90\x96JEq\xd3" +
	"@\x88\x0aU\xb4\xd5{\xbbow\xcf\xbe\x838\xed\x1f" +
	" \xdb7o\xdeo~\xf3\x9by3\xb7\xbc\xde\xd7%" +
	"\xb4\xfaNV\x01\xc8/\xfan1\x1a\xd2g\x06z?" +
	"\xbc\xfcM \xb7z\x8c\x03{\x17\xef\x7f\xb5\xff\xf2%" +
	"\x00\xa8Fi\xa3xI\xba[\x14\x01\xa4~q\x8ft" +
	"\x8d\xfed\xdc3\xf0\xd2q\xe1\xdfo\xef\x06\xf2\x19\x04" +
	"\xf0\xa1X\x8d\xed\x17\xc4I\x04\x94f\xc5$\xa0\xb1\xfe" +
	"\x9d\xafu\x14\x8en\xf9\x1e\x90\x10\x1a\xfb\x0e\xacY6" +
	"\x11\x99~\x17|\x02u\x13\xf1\xff^\xaa\xf3/\x02\x90" +
	"\x1a\xfd;\x00\x8d\x1fT=~\xf9\x1b}k\x1e\x01r" +
	"\xab\xed\xec\x80?A\x9d=\xe7\xa7\xceN+\xc7\xf4\x07" +
	"W?\xf5\x08\xc8\xb7\"\x02xE\x80\xf6i\x7f\x9a\x1a" +
	"\x9cg\x06\x1f\x9c\xb92U\xfb\xe6\xbaGA\xae\xb6\x0d" +
	">\xf2\x0b\xd4\x00\x03\xd4\xe0\xbb\x0f\xcctt\xde\x7f\x8e" +
	"\x19\x08\x96A}\x80\x194\x06(\x86-{\x0f\xc7\xe2" +
	"\x1f^|\x14H\xb5\xe5@\xda\x15\xb8\x0e^\xa3a\xc5" +
	"\xcb\x87\xff\xb6\xeb\xd9'\xca\x05\xa2\x06\x0eJ\xa3\x01\xfa" +
	"S\x969\xe9\x0c\x9f\xfdr\xf1\x8b\xb5S@\x16\x09F" +
	"\xf2\xe0\xa1\xd8\xe0}\xaf_\x07\xc0\xf6\xdf\x05\x04\x94N" +
	"1\xcb\x13\x81\x0e@\xe3r\xbe\xa6\xbe\xbefv\xca\x0e" +
	"\x19\xa0\xfdT\xe06\x8ag\x86\x01\xde\xb7\xa2f\xc9\xc1" +
	"\x96\xf7\xa6\xdc\x04_\x0btS\x03_\x90\x1a\x9c\xdas" +
	"tOS\xc7[\xfb\xdd\x06\x8dA\xc6I\x0b3x;" +
	"=0\xf0b\xc3\x07\xfb\x81\x84<\x0ep@icp" +
	"\x9f\xd4\x1f\xa4X\xe4\xa0H\xff\x01\x18w\xfc\xe1\xcf\x87" +
	"\xfeQ'>c\x11\xc8\x00I\x9d\xc1\xeb\x80\xd2\x9a " +
	"\x0dm\xfd+W6\xb7\x0f\xc4\x9e5\x01{\xe9mO" +
	"\x07\xdb\x10\xbc\xc6\xde\x1d\x9f:\x7f\xcbO\xaf\xfe\x18H" +
	"\x9d+\xed\x8c\xa0\xf6\xef\x04\xdbPz\x92\xdd\xb5\x97\xb9" +
	"yg\xddkg\x8a/\xfc\xe6' /\xb2\xefi\x9f" +
	"\xa5\x8eP\xba\xca,\xb6.\x0dm\xae9y\xdf\xf3@" +
	"\x16\xf1L\xde\x1dj\xa6\x17\xd5\xa5vo:}\xee\x99" +
	"\x97\xcd\x88\x19\x84T(M?\xb9s{\xfaD\xc7\xab" +
	"\xf2+\x16\x9b\x02\xfdhi\xa8\x86:m\x0dQ\xa7\x83" +
	"\x9b\xe4?\xfd\xac\xe9\xa3\xd7]\xe8\x9f\x0b\x05\xe9\xd1\xb1" +
	"\xa3w}~\xfd\x92\xda_\xf0O\x00\xda\x1f\xa6'\xbd" +
	"\xc6\xce\xc7V\xb5<\xf5\xd2\xbdo\x98\x8a0\x91~\x95" +
	"\x1eBi\"D\x09\x9e\xba\xf2\xdb]\xf8\xab7\xdf(" +
	"\x91m\xc8\x94-3\xb8\xff\xaf'\x97\xeeX\xb5\xf9\xc4" +
	"\xbc\x0cL\x87\x8eH\xe7C\x94\x95\x99\xd01\xa9?L" +
	"3Pwd\xf9?W\xce\x9c}\xcb\x05dM\xf86" +
	"\x0a\xe4sO\xdf\xf5xr{\xe3t\x99\xd2l\x0c_" +
	"\x92Z\xe9q\xa9%\xbcV\xfa:s\xd4\xf5\xc2\xaf\xbf" +
	"\xf0\xcb\x9e\xd6sn\xdcj\x98I=\x1b\xa6\xb0\x0e\x1f" +
	"\xfe\xfb\xe9\x95\xaf\xfd\xe7\x9c\x8b\xc7\x87\xc3\xdd\xf4\xa6\xef" +
	"\xd7=\xb6\xe5\xf9\xa6\x95\x17\xe7\x01.\x86\xbf%M\xb0" +
	"{\x8a\xe1c\xd2\xfb\xec\x1e\x1bI\xa91\x855\x13>" +
	"\"]\x087\x01HW\xc3\xc7\xa4\xd9\x88(\xcdF\xa2" +
	"\xc6\xf1\xdd\xc9\xfc\x8f\x9a\xff\xf2\xae+\xc2\x0b\x91\xc5\xf4" +
	"\xdeC\x7f\x8c\xef\x7f\xef\xdbM\xffrk\xf9x\xa4@" +
	"!OG(\xe4\x1f\xae\x95\x9b\xcf\xae\xdf|\xc5u4" +
	"P\xc5 \x9f\\\xf2\xd9=\xf7<!^se\xf6j" +
	"\x84\xe9\xd2\xae\xbf\xb9\xc1\x9c\x8f\xfc\\\xba\x18\xa1\xed\xe7" +
	"\xfd\xc8Z\xa9\xaeJ\x84\xb4Q\xd4\xb3\xb9\xf8\x902." +
	"\x8c\x8d'R\x9a\x96\x1d\x19S\x06sj|\xad\x1a\xd5" +
	"u\xb5\x90A\x94\xfd\x1e\x1f\xcd\x92\xa5A\xe4\xf5IZ" +
	"\x17\x03\xa4\x96aj5\x92\x8d\"\xa2\x1d\x0b\xf2\x16I" +
	"R=\x00\xa9.Le\x90(\xa28\xa2\xea\x04c\xb2" +
	"W@\x87AZ\x0ae\xff\xd8\x85\x86V\x1c\xd4\x86\x0a" +
	"\xd9A@u!\xe72\x88vP\xde\xb9A\xd1\x98\xe2" +
	"\xdc\xb1\xda\xd0\x93T\xb5bN\xd7d\xaf\xc7\x0b\xe0\xa5" +
	"N#\x09\x00\xd9\xefA\xb9V\xc0\xe46el8\xa7" +
	"\"qZ\x0a \x12\xc0\x12\xd66\xe4\x87\x94\\vR" +
	"\x1d\xeeSw\xea\x00\xb2\x17\xd1\xd5\x1ap\xbba\x19(" +
	"\x10\xd5\xb3\xf91\xd9o\xdf\xb5t\x10@\xbe\xdd\x83\xf2" +
	"\x0a\x01\x09b-\x8b\xaa\xb5\x00 /\xf7\xa0\xbcA@" +
	"cX\xdd\xaa\x14sz\x1f\x88\xeaN\x1d\xc3 `\x18" +
	"\xd0\xc8q\x871\xeaP\xc3*\xc0\x8c\x07\xb1\xda\xb9\x16" +
	"\x90\xfe\xd1\x86\xe9)\xe5A\xd1L&n\x18\xff\xed\x02" +
	"&G\x98\x19\x12\xe7\x91\xa4\x0c\x97%\xbe\x84\x16\xca{" +
	"zBW{\xf5\x82\xaa\x8c\xc6\xd5\x9d\xe3\xea\x90\xde\x9b" +
	"\x9dT\x1b2JA\x19\xd5\xc0}a\xb3CxT\xcb" +
	"N\xaa\x18\x00\x01\x03s\xd0\xa7s\xf9\xc1\xf8\x88j:" +
	"\xe91Q\xff\x0f^r\xd9!\x07\x8b+!4\xf8\x06" +
	"\x0f\xca\xcbiB\x043!-\xcdN\x96\x92\xf9\xad[" +
	"5U\xe7\xce\xcd\x9b\xfc \xa0\x7f\x8e(J\xf0&\xcd" +
	"\x9bl\x03\x1c\x1bO\xdc\xa1NlR\xc4\\Q5K" +
	"\xccF\xb0\xd8\x8d\xc0\x92DK\x9b\x83@|@\x9d\xe0" +
	"B\x88=\xa8\xe4\x8a\xaa-\x0bw\xa4\xbd\xba\xa2g\x87" +
	"R\x9a\xa6\xea\xf1L!\xaf\xe7\x87\xf29`W\xd1)" +
	"\x80\x906\xa6\xe4@3@l\x9b\xae\x8fkQ\xfa\x7f" +
	")WN\xfev\x14\xb2\xbaI\x97gT\xabD\xfa\xb0" +
	"\xa2+\x18\x01\x01#p\xb3\x02\x1cQu;\x87%I" +
	"ls\xfcY\xb1\xd5 \xba;?\xd6Tv\xdekU" +
	"\xb7\xaa\x97S\xd9\xc2\\\xd3\x0c\xad\xa3\xc5\x8f4?\x19" +
	"\x8f\xcf\xddU\xec\xfc2^\xfa\xf2<\xbf\x15\x84\xc4\xb2" +
	"\x88HZha/\xf3\xa0\xbcJ\xc0\xa4\xc6\x98E\xe2" +
	"<\x95Vg\xd1t\xa5\xa0\xa7\xf4;!V\xa2\xb4\xca" +
	"\xa5l\x86\xdd\x90\x89\x95j\xcc3\xb73\xc5\xad\xdfb" +
	"\x0am\x1asT\x97(\xa7:\xb7\xeeY\xcb\xb1\xa5\x16" +
	"\xd5\xdd\xed\xa8\x92\xeeFT\xbd\xbf\x90k\xe8Qcf" +
	"\x93q\xdd\xd7\xed\xb8\xb6\xfb^\xb7\xd5\xf7V\x0bh\x8c" +
	"s\xc9\x02`\xd4\x99.\x011\x0ahl\xcbkzF" +
	"\xd1\xb7\xd1O?&\x88\x8c\x12-\xa1\xc7_Q\x98\x96" +
	"v\xb8a\xa9\x9d\xa3\\,\x89\xa7\xad\x1c\x7f\x09+\xc8" +
	"u7\x92\\R\xe3=\xd6\x9ev>^\x8f\xf5\xccG" +
	"fA\x87Ju<\x9c\x1f\xb3\xba\x1ej7\xe8\x8be" +
	"K\xf3\xa6\xa5^I\x987{c>i\xfc\xc8c\x8b" +
	"\xb1\xe0\xa8\xa0\xc3lR\xe1\x1b\x06\xf2\xb1\x8b\xc8m " +
	"\x90/\xd1\x19\x85\x0f\xbd\xc8\x87N\xd2\xd9\x0c\x02i\x11" +
	"Q\xb0w-\xe43\x17\xa9\xff\x0a\x08\xa4N\x8c\xb1:" +
	"\xef\xc2(\xe5\xaf\x0b\x0d\xfe\x9c\x81gR-\x9d7\xca" +
	"\xf1\xcd\xd8\xf4\xe8Z\xa5I\xab\xd7\x99\xb4\xbc\x0c?\xdf" +
	"o\x90\x8f\xab\x84\xd0I+\x8c\xa9O#i\x11Em" +
	"as\xd4<|\xee\xceU\xee%]\xd0\x00T\xb9\x07" +
	"\xdb**\xcdX.?\x08N\xae\xf8\xf2\x89|\x8f%" +
	"r\x9a\xe7\x8aof\xc8\xe7z\xd2\x99\xe6\xb9\xe2[-" +
	"\xf2]\x85\xd4w\xb3\\=d\xbd\xb9]\xf8\x90\x15c" +
	"\x17\x1a\\\xd90\x97\x0f\xe4\xe8c\x0c\xbe\xecGt}" +
	"\x1f\x10H\xb86\x10_\"i6\x89\xa4\x19\xa3\x15\x00" +
	"\xdf\xaf\x90obD\xa6\xc9\xda\x80\xa9{\x91\x8c\xd20" +
	"\xf8\x9c\x8e|\xd9'J7@j\x00S\xe3Hv\xd1" +
	"`\xf8\x86\x89|\xaf\"Ej\xa1cj7\x92'\x17" +
	">8\xf3\xf1\x0e*\x98T>\xd7\xfbI\xcee\x10S" +
	"^$X\x83}7\x1e \xa8\"\xc4\x9c\xab\x12n6" +
	"\x9e[o\xda\xff\xbbsT\x9eN\xe7k\xb6\xd2\xe4<" +
	"\xf7\xb9E\xfe\xec$\xcdw\x87m\x01\xce\xb7\"\xd8m" +
	"d\x9c'\xcd*t\xbe\xf0#\xffn\x80\x90\x04\x08\xc4" +
	"'&\xcdG\x8bQ\xfb\xdf\x00\x00\x00\xff\xff\xb1WG" +
	"\xb8"

func init() {
	schemas.Register(schema_ecd50d792c3d9992,
		0x80f2f65360d64224,
		0x84e0f802c9af605b,
		0x8b5db772377be249,
		0x8c3d547ef2930e96,
		0x8c9a3c7674c761d3,
		0x8e48cb1497f3d6f4,
		0x8edb5f3937d96b8a,
		0x8ee5f62e1fab915d,
		0x94a081e4abb13424,
		0x9714437546d80c39,
		0x97ed122121126ff2,
		0x97ef2da226123492,
		0x98d0372787b787d1,
		0x98f424ac606042e0,
		0x9f0719e9a9dccc4b,
		0xa01f603357f3b349,
		0xa4f5ae06dd1b7791,
		0xa5c3aa75d6b648e2,
		0xa75ecf12570b2966,
		0xb19fdbd356844119,
		0xb351b437cd426a4f,
		0xbbfd27b5d2515662,
		0xbc1426493658b76e,
		0xc65caf9a2d389078,
		0xc6cbc10181c4f397,
		0xcd57387729cfe35f,
		0xd0d8d935ee30b219,
		0xd5256a3f93589d2f,
		0xdb3152bd3bc2aa40,
		0xdbfbb635d3e6abab,
		0xe53527a75d90198f,
		0xeaf255b498229199,
		0xecde2a9c6f3f84c9,
		0xf02783ef982ecea9,
		0xf35749d82a51479b,
		0xf907945b872b26cf,
		0xfabb5e621fa9a23f)
}
