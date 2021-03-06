package webpublishing

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type WebSite struct{ Client capnp.Client }

func (c WebSite) GetUrl(ctx context.Context, params func(WebSite_getUrl_Params) error, opts ...capnp.CallOption) WebSite_getUrl_Results_Promise {
	if c.Client == nil {
		return WebSite_getUrl_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      0,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getUrl",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_getUrl_Params{Struct: s}) }
	}
	return WebSite_getUrl_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) GetEntities(ctx context.Context, params func(WebSite_getEntities_Params) error, opts ...capnp.CallOption) WebSite_getEntities_Results_Promise {
	if c.Client == nil {
		return WebSite_getEntities_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      1,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getEntities",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_getEntities_Params{Struct: s}) }
	}
	return WebSite_getEntities_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) GetNotFoundEntities(ctx context.Context, params func(WebSite_getNotFoundEntities_Params) error, opts ...capnp.CallOption) WebSite_getNotFoundEntities_Results_Promise {
	if c.Client == nil {
		return WebSite_getNotFoundEntities_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      2,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getNotFoundEntities",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_getNotFoundEntities_Params{Struct: s}) }
	}
	return WebSite_getNotFoundEntities_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) UploadBlob(ctx context.Context, params func(WebSite_uploadBlob_Params) error, opts ...capnp.CallOption) WebSite_uploadBlob_Results_Promise {
	if c.Client == nil {
		return WebSite_uploadBlob_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      3,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "uploadBlob",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_uploadBlob_Params{Struct: s}) }
	}
	return WebSite_uploadBlob_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) GetSubsite(ctx context.Context, params func(WebSite_getSubsite_Params) error, opts ...capnp.CallOption) WebSite_getSubsite_Results_Promise {
	if c.Client == nil {
		return WebSite_getSubsite_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      4,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getSubsite",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_getSubsite_Params{Struct: s}) }
	}
	return WebSite_getSubsite_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c WebSite) ListResources(ctx context.Context, params func(WebSite_listResources_Params) error, opts ...capnp.CallOption) WebSite_listResources_Results_Promise {
	if c.Client == nil {
		return WebSite_listResources_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      5,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "listResources",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(WebSite_listResources_Params{Struct: s}) }
	}
	return WebSite_listResources_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type WebSite_Server interface {
	GetUrl(WebSite_getUrl) error

	GetEntities(WebSite_getEntities) error

	GetNotFoundEntities(WebSite_getNotFoundEntities) error

	UploadBlob(WebSite_uploadBlob) error

	GetSubsite(WebSite_getSubsite) error

	ListResources(WebSite_listResources) error
}

func WebSite_ServerToClient(s WebSite_Server) WebSite {
	c, _ := s.(server.Closer)
	return WebSite{Client: server.New(WebSite_Methods(nil, s), c)}
}

func WebSite_Methods(methods []server.Method, s WebSite_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 6)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      0,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getUrl",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_getUrl{c, opts, WebSite_getUrl_Params{Struct: p}, WebSite_getUrl_Results{Struct: r}}
			return s.GetUrl(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      1,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getEntities",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_getEntities{c, opts, WebSite_getEntities_Params{Struct: p}, WebSite_getEntities_Results{Struct: r}}
			return s.GetEntities(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      2,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getNotFoundEntities",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_getNotFoundEntities{c, opts, WebSite_getNotFoundEntities_Params{Struct: p}, WebSite_getNotFoundEntities_Results{Struct: r}}
			return s.GetNotFoundEntities(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      3,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "uploadBlob",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_uploadBlob{c, opts, WebSite_uploadBlob_Params{Struct: p}, WebSite_uploadBlob_Results{Struct: r}}
			return s.UploadBlob(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      4,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "getSubsite",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_getSubsite{c, opts, WebSite_getSubsite_Params{Struct: p}, WebSite_getSubsite_Results{Struct: r}}
			return s.GetSubsite(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdddcca47803e2377,
			MethodID:      5,
			InterfaceName: "web-publishing.capnp:WebSite",
			MethodName:    "listResources",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := WebSite_listResources{c, opts, WebSite_listResources_Params{Struct: p}, WebSite_listResources_Results{Struct: r}}
			return s.ListResources(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// WebSite_getUrl holds the arguments for a server call to WebSite.getUrl.
type WebSite_getUrl struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_getUrl_Params
	Results WebSite_getUrl_Results
}

// WebSite_getEntities holds the arguments for a server call to WebSite.getEntities.
type WebSite_getEntities struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_getEntities_Params
	Results WebSite_getEntities_Results
}

// WebSite_getNotFoundEntities holds the arguments for a server call to WebSite.getNotFoundEntities.
type WebSite_getNotFoundEntities struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_getNotFoundEntities_Params
	Results WebSite_getNotFoundEntities_Results
}

// WebSite_uploadBlob holds the arguments for a server call to WebSite.uploadBlob.
type WebSite_uploadBlob struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_uploadBlob_Params
	Results WebSite_uploadBlob_Results
}

// WebSite_getSubsite holds the arguments for a server call to WebSite.getSubsite.
type WebSite_getSubsite struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_getSubsite_Params
	Results WebSite_getSubsite_Results
}

// WebSite_listResources holds the arguments for a server call to WebSite.listResources.
type WebSite_listResources struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  WebSite_listResources_Params
	Results WebSite_listResources_Results
}

type WebSite_Entity struct{ capnp.Struct }
type WebSite_Entity_body WebSite_Entity
type WebSite_Entity_body_Which uint16

const (
	WebSite_Entity_body_Which_bytes WebSite_Entity_body_Which = 0
	WebSite_Entity_body_Which_blob  WebSite_Entity_body_Which = 1
)

func (w WebSite_Entity_body_Which) String() string {
	const s = "bytesblob"
	switch w {
	case WebSite_Entity_body_Which_bytes:
		return s[0:5]
	case WebSite_Entity_body_Which_blob:
		return s[5:9]

	}
	return "WebSite_Entity_body_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// WebSite_Entity_TypeID is the unique identifier for the type WebSite_Entity.
const WebSite_Entity_TypeID = 0xd019dd3e3c0e7e68

func NewWebSite_Entity(s *capnp.Segment) (WebSite_Entity, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return WebSite_Entity{st}, err
}

func NewRootWebSite_Entity(s *capnp.Segment) (WebSite_Entity, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return WebSite_Entity{st}, err
}

func ReadRootWebSite_Entity(msg *capnp.Message) (WebSite_Entity, error) {
	root, err := msg.RootPtr()
	return WebSite_Entity{root.Struct()}, err
}

func (s WebSite_Entity) String() string {
	str, _ := text.Marshal(0xd019dd3e3c0e7e68, s.Struct)
	return str
}

func (s WebSite_Entity) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSite_Entity) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSite_Entity) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s WebSite_Entity) Language() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s WebSite_Entity) HasLanguage() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity) LanguageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s WebSite_Entity) SetLanguage(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s WebSite_Entity) Encoding() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s WebSite_Entity) HasEncoding() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity) EncodingBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s WebSite_Entity) SetEncoding(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s WebSite_Entity) Body() WebSite_Entity_body { return WebSite_Entity_body(s) }

func (s WebSite_Entity_body) Which() WebSite_Entity_body_Which {
	return WebSite_Entity_body_Which(s.Struct.Uint16(0))
}
func (s WebSite_Entity_body) Bytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s WebSite_Entity_body) HasBytes() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity_body) SetBytes(v []byte) error {
	s.Struct.SetUint16(0, 0)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, d.List.ToPtr())
}

func (s WebSite_Entity_body) Blob() util.Blob {
	p, _ := s.Struct.Ptr(3)
	return util.Blob{Client: p.Interface().Client()}
}

func (s WebSite_Entity_body) HasBlob() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity_body) SetBlob(v util.Blob) error {
	s.Struct.SetUint16(0, 1)
	if v.Client == nil {
		return s.Struct.SetPtr(3, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(3, in.ToPtr())
}

func (s WebSite_Entity) RedirectTo() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s WebSite_Entity) HasRedirectTo() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s WebSite_Entity) RedirectToBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s WebSite_Entity) SetRedirectTo(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, t.List.ToPtr())
}

// WebSite_Entity_List is a list of WebSite_Entity.
type WebSite_Entity_List struct{ capnp.List }

// NewWebSite_Entity creates a new list of WebSite_Entity.
func NewWebSite_Entity_List(s *capnp.Segment, sz int32) (WebSite_Entity_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return WebSite_Entity_List{l}, err
}

func (s WebSite_Entity_List) At(i int) WebSite_Entity { return WebSite_Entity{s.List.Struct(i)} }

func (s WebSite_Entity_List) Set(i int, v WebSite_Entity) error { return s.List.SetStruct(i, v.Struct) }

// WebSite_Entity_Promise is a wrapper for a WebSite_Entity promised by a client call.
type WebSite_Entity_Promise struct{ *capnp.Pipeline }

func (p WebSite_Entity_Promise) Struct() (WebSite_Entity, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_Entity{s}, err
}

func (p WebSite_Entity_Promise) Body() WebSite_Entity_body_Promise {
	return WebSite_Entity_body_Promise{p.Pipeline}
}

// WebSite_Entity_body_Promise is a wrapper for a WebSite_Entity_body promised by a client call.
type WebSite_Entity_body_Promise struct{ *capnp.Pipeline }

func (p WebSite_Entity_body_Promise) Struct() (WebSite_Entity_body, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_Entity_body{s}, err
}

func (p WebSite_Entity_body_Promise) Blob() util.Blob {
	return util.Blob{Client: p.Pipeline.GetPipeline(3).Client()}
}

type WebSite_getUrl_Params struct{ capnp.Struct }

// WebSite_getUrl_Params_TypeID is the unique identifier for the type WebSite_getUrl_Params.
const WebSite_getUrl_Params_TypeID = 0xdc072ae47872d061

func NewWebSite_getUrl_Params(s *capnp.Segment) (WebSite_getUrl_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_getUrl_Params{st}, err
}

func NewRootWebSite_getUrl_Params(s *capnp.Segment) (WebSite_getUrl_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_getUrl_Params{st}, err
}

func ReadRootWebSite_getUrl_Params(msg *capnp.Message) (WebSite_getUrl_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_getUrl_Params{root.Struct()}, err
}

func (s WebSite_getUrl_Params) String() string {
	str, _ := text.Marshal(0xdc072ae47872d061, s.Struct)
	return str
}

// WebSite_getUrl_Params_List is a list of WebSite_getUrl_Params.
type WebSite_getUrl_Params_List struct{ capnp.List }

// NewWebSite_getUrl_Params creates a new list of WebSite_getUrl_Params.
func NewWebSite_getUrl_Params_List(s *capnp.Segment, sz int32) (WebSite_getUrl_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return WebSite_getUrl_Params_List{l}, err
}

func (s WebSite_getUrl_Params_List) At(i int) WebSite_getUrl_Params {
	return WebSite_getUrl_Params{s.List.Struct(i)}
}

func (s WebSite_getUrl_Params_List) Set(i int, v WebSite_getUrl_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getUrl_Params_Promise is a wrapper for a WebSite_getUrl_Params promised by a client call.
type WebSite_getUrl_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_getUrl_Params_Promise) Struct() (WebSite_getUrl_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getUrl_Params{s}, err
}

type WebSite_getUrl_Results struct{ capnp.Struct }

// WebSite_getUrl_Results_TypeID is the unique identifier for the type WebSite_getUrl_Results.
const WebSite_getUrl_Results_TypeID = 0xb2855d483568639e

func NewWebSite_getUrl_Results(s *capnp.Segment) (WebSite_getUrl_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getUrl_Results{st}, err
}

func NewRootWebSite_getUrl_Results(s *capnp.Segment) (WebSite_getUrl_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getUrl_Results{st}, err
}

func ReadRootWebSite_getUrl_Results(msg *capnp.Message) (WebSite_getUrl_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_getUrl_Results{root.Struct()}, err
}

func (s WebSite_getUrl_Results) String() string {
	str, _ := text.Marshal(0xb2855d483568639e, s.Struct)
	return str
}

func (s WebSite_getUrl_Results) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSite_getUrl_Results) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getUrl_Results) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSite_getUrl_Results) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// WebSite_getUrl_Results_List is a list of WebSite_getUrl_Results.
type WebSite_getUrl_Results_List struct{ capnp.List }

// NewWebSite_getUrl_Results creates a new list of WebSite_getUrl_Results.
func NewWebSite_getUrl_Results_List(s *capnp.Segment, sz int32) (WebSite_getUrl_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getUrl_Results_List{l}, err
}

func (s WebSite_getUrl_Results_List) At(i int) WebSite_getUrl_Results {
	return WebSite_getUrl_Results{s.List.Struct(i)}
}

func (s WebSite_getUrl_Results_List) Set(i int, v WebSite_getUrl_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getUrl_Results_Promise is a wrapper for a WebSite_getUrl_Results promised by a client call.
type WebSite_getUrl_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_getUrl_Results_Promise) Struct() (WebSite_getUrl_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getUrl_Results{s}, err
}

type WebSite_getEntities_Params struct{ capnp.Struct }

// WebSite_getEntities_Params_TypeID is the unique identifier for the type WebSite_getEntities_Params.
const WebSite_getEntities_Params_TypeID = 0xfe1643d6b01e7cf4

func NewWebSite_getEntities_Params(s *capnp.Segment) (WebSite_getEntities_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getEntities_Params{st}, err
}

func NewRootWebSite_getEntities_Params(s *capnp.Segment) (WebSite_getEntities_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getEntities_Params{st}, err
}

func ReadRootWebSite_getEntities_Params(msg *capnp.Message) (WebSite_getEntities_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_getEntities_Params{root.Struct()}, err
}

func (s WebSite_getEntities_Params) String() string {
	str, _ := text.Marshal(0xfe1643d6b01e7cf4, s.Struct)
	return str
}

func (s WebSite_getEntities_Params) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSite_getEntities_Params) HasPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getEntities_Params) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSite_getEntities_Params) SetPath(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// WebSite_getEntities_Params_List is a list of WebSite_getEntities_Params.
type WebSite_getEntities_Params_List struct{ capnp.List }

// NewWebSite_getEntities_Params creates a new list of WebSite_getEntities_Params.
func NewWebSite_getEntities_Params_List(s *capnp.Segment, sz int32) (WebSite_getEntities_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getEntities_Params_List{l}, err
}

func (s WebSite_getEntities_Params_List) At(i int) WebSite_getEntities_Params {
	return WebSite_getEntities_Params{s.List.Struct(i)}
}

func (s WebSite_getEntities_Params_List) Set(i int, v WebSite_getEntities_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getEntities_Params_Promise is a wrapper for a WebSite_getEntities_Params promised by a client call.
type WebSite_getEntities_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_getEntities_Params_Promise) Struct() (WebSite_getEntities_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getEntities_Params{s}, err
}

type WebSite_getEntities_Results struct{ capnp.Struct }

// WebSite_getEntities_Results_TypeID is the unique identifier for the type WebSite_getEntities_Results.
const WebSite_getEntities_Results_TypeID = 0xc3739b6f070fb5ea

func NewWebSite_getEntities_Results(s *capnp.Segment) (WebSite_getEntities_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getEntities_Results{st}, err
}

func NewRootWebSite_getEntities_Results(s *capnp.Segment) (WebSite_getEntities_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getEntities_Results{st}, err
}

func ReadRootWebSite_getEntities_Results(msg *capnp.Message) (WebSite_getEntities_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_getEntities_Results{root.Struct()}, err
}

func (s WebSite_getEntities_Results) String() string {
	str, _ := text.Marshal(0xc3739b6f070fb5ea, s.Struct)
	return str
}

func (s WebSite_getEntities_Results) Entities() util.Assignable {
	p, _ := s.Struct.Ptr(0)
	return util.Assignable{Client: p.Interface().Client()}
}

func (s WebSite_getEntities_Results) HasEntities() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getEntities_Results) SetEntities(v util.Assignable) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// WebSite_getEntities_Results_List is a list of WebSite_getEntities_Results.
type WebSite_getEntities_Results_List struct{ capnp.List }

// NewWebSite_getEntities_Results creates a new list of WebSite_getEntities_Results.
func NewWebSite_getEntities_Results_List(s *capnp.Segment, sz int32) (WebSite_getEntities_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getEntities_Results_List{l}, err
}

func (s WebSite_getEntities_Results_List) At(i int) WebSite_getEntities_Results {
	return WebSite_getEntities_Results{s.List.Struct(i)}
}

func (s WebSite_getEntities_Results_List) Set(i int, v WebSite_getEntities_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getEntities_Results_Promise is a wrapper for a WebSite_getEntities_Results promised by a client call.
type WebSite_getEntities_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_getEntities_Results_Promise) Struct() (WebSite_getEntities_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getEntities_Results{s}, err
}

func (p WebSite_getEntities_Results_Promise) Entities() util.Assignable {
	return util.Assignable{Client: p.Pipeline.GetPipeline(0).Client()}
}

type WebSite_getNotFoundEntities_Params struct{ capnp.Struct }

// WebSite_getNotFoundEntities_Params_TypeID is the unique identifier for the type WebSite_getNotFoundEntities_Params.
const WebSite_getNotFoundEntities_Params_TypeID = 0xc3db68ea10a823b6

func NewWebSite_getNotFoundEntities_Params(s *capnp.Segment) (WebSite_getNotFoundEntities_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_getNotFoundEntities_Params{st}, err
}

func NewRootWebSite_getNotFoundEntities_Params(s *capnp.Segment) (WebSite_getNotFoundEntities_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_getNotFoundEntities_Params{st}, err
}

func ReadRootWebSite_getNotFoundEntities_Params(msg *capnp.Message) (WebSite_getNotFoundEntities_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_getNotFoundEntities_Params{root.Struct()}, err
}

func (s WebSite_getNotFoundEntities_Params) String() string {
	str, _ := text.Marshal(0xc3db68ea10a823b6, s.Struct)
	return str
}

// WebSite_getNotFoundEntities_Params_List is a list of WebSite_getNotFoundEntities_Params.
type WebSite_getNotFoundEntities_Params_List struct{ capnp.List }

// NewWebSite_getNotFoundEntities_Params creates a new list of WebSite_getNotFoundEntities_Params.
func NewWebSite_getNotFoundEntities_Params_List(s *capnp.Segment, sz int32) (WebSite_getNotFoundEntities_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return WebSite_getNotFoundEntities_Params_List{l}, err
}

func (s WebSite_getNotFoundEntities_Params_List) At(i int) WebSite_getNotFoundEntities_Params {
	return WebSite_getNotFoundEntities_Params{s.List.Struct(i)}
}

func (s WebSite_getNotFoundEntities_Params_List) Set(i int, v WebSite_getNotFoundEntities_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getNotFoundEntities_Params_Promise is a wrapper for a WebSite_getNotFoundEntities_Params promised by a client call.
type WebSite_getNotFoundEntities_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_getNotFoundEntities_Params_Promise) Struct() (WebSite_getNotFoundEntities_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getNotFoundEntities_Params{s}, err
}

type WebSite_getNotFoundEntities_Results struct{ capnp.Struct }

// WebSite_getNotFoundEntities_Results_TypeID is the unique identifier for the type WebSite_getNotFoundEntities_Results.
const WebSite_getNotFoundEntities_Results_TypeID = 0xff383cbabc241462

func NewWebSite_getNotFoundEntities_Results(s *capnp.Segment) (WebSite_getNotFoundEntities_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getNotFoundEntities_Results{st}, err
}

func NewRootWebSite_getNotFoundEntities_Results(s *capnp.Segment) (WebSite_getNotFoundEntities_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getNotFoundEntities_Results{st}, err
}

func ReadRootWebSite_getNotFoundEntities_Results(msg *capnp.Message) (WebSite_getNotFoundEntities_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_getNotFoundEntities_Results{root.Struct()}, err
}

func (s WebSite_getNotFoundEntities_Results) String() string {
	str, _ := text.Marshal(0xff383cbabc241462, s.Struct)
	return str
}

func (s WebSite_getNotFoundEntities_Results) Entities() util.Assignable {
	p, _ := s.Struct.Ptr(0)
	return util.Assignable{Client: p.Interface().Client()}
}

func (s WebSite_getNotFoundEntities_Results) HasEntities() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getNotFoundEntities_Results) SetEntities(v util.Assignable) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// WebSite_getNotFoundEntities_Results_List is a list of WebSite_getNotFoundEntities_Results.
type WebSite_getNotFoundEntities_Results_List struct{ capnp.List }

// NewWebSite_getNotFoundEntities_Results creates a new list of WebSite_getNotFoundEntities_Results.
func NewWebSite_getNotFoundEntities_Results_List(s *capnp.Segment, sz int32) (WebSite_getNotFoundEntities_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getNotFoundEntities_Results_List{l}, err
}

func (s WebSite_getNotFoundEntities_Results_List) At(i int) WebSite_getNotFoundEntities_Results {
	return WebSite_getNotFoundEntities_Results{s.List.Struct(i)}
}

func (s WebSite_getNotFoundEntities_Results_List) Set(i int, v WebSite_getNotFoundEntities_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getNotFoundEntities_Results_Promise is a wrapper for a WebSite_getNotFoundEntities_Results promised by a client call.
type WebSite_getNotFoundEntities_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_getNotFoundEntities_Results_Promise) Struct() (WebSite_getNotFoundEntities_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getNotFoundEntities_Results{s}, err
}

func (p WebSite_getNotFoundEntities_Results_Promise) Entities() util.Assignable {
	return util.Assignable{Client: p.Pipeline.GetPipeline(0).Client()}
}

type WebSite_uploadBlob_Params struct{ capnp.Struct }

// WebSite_uploadBlob_Params_TypeID is the unique identifier for the type WebSite_uploadBlob_Params.
const WebSite_uploadBlob_Params_TypeID = 0xe2ae317a2a41a9f6

func NewWebSite_uploadBlob_Params(s *capnp.Segment) (WebSite_uploadBlob_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_uploadBlob_Params{st}, err
}

func NewRootWebSite_uploadBlob_Params(s *capnp.Segment) (WebSite_uploadBlob_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return WebSite_uploadBlob_Params{st}, err
}

func ReadRootWebSite_uploadBlob_Params(msg *capnp.Message) (WebSite_uploadBlob_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_uploadBlob_Params{root.Struct()}, err
}

func (s WebSite_uploadBlob_Params) String() string {
	str, _ := text.Marshal(0xe2ae317a2a41a9f6, s.Struct)
	return str
}

// WebSite_uploadBlob_Params_List is a list of WebSite_uploadBlob_Params.
type WebSite_uploadBlob_Params_List struct{ capnp.List }

// NewWebSite_uploadBlob_Params creates a new list of WebSite_uploadBlob_Params.
func NewWebSite_uploadBlob_Params_List(s *capnp.Segment, sz int32) (WebSite_uploadBlob_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return WebSite_uploadBlob_Params_List{l}, err
}

func (s WebSite_uploadBlob_Params_List) At(i int) WebSite_uploadBlob_Params {
	return WebSite_uploadBlob_Params{s.List.Struct(i)}
}

func (s WebSite_uploadBlob_Params_List) Set(i int, v WebSite_uploadBlob_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_uploadBlob_Params_Promise is a wrapper for a WebSite_uploadBlob_Params promised by a client call.
type WebSite_uploadBlob_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_uploadBlob_Params_Promise) Struct() (WebSite_uploadBlob_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_uploadBlob_Params{s}, err
}

type WebSite_uploadBlob_Results struct{ capnp.Struct }

// WebSite_uploadBlob_Results_TypeID is the unique identifier for the type WebSite_uploadBlob_Results.
const WebSite_uploadBlob_Results_TypeID = 0xc63c6c15dd189744

func NewWebSite_uploadBlob_Results(s *capnp.Segment) (WebSite_uploadBlob_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSite_uploadBlob_Results{st}, err
}

func NewRootWebSite_uploadBlob_Results(s *capnp.Segment) (WebSite_uploadBlob_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return WebSite_uploadBlob_Results{st}, err
}

func ReadRootWebSite_uploadBlob_Results(msg *capnp.Message) (WebSite_uploadBlob_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_uploadBlob_Results{root.Struct()}, err
}

func (s WebSite_uploadBlob_Results) String() string {
	str, _ := text.Marshal(0xc63c6c15dd189744, s.Struct)
	return str
}

func (s WebSite_uploadBlob_Results) Blob() util.Blob {
	p, _ := s.Struct.Ptr(0)
	return util.Blob{Client: p.Interface().Client()}
}

func (s WebSite_uploadBlob_Results) HasBlob() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_uploadBlob_Results) SetBlob(v util.Blob) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s WebSite_uploadBlob_Results) Stream() util.ByteStream {
	p, _ := s.Struct.Ptr(1)
	return util.ByteStream{Client: p.Interface().Client()}
}

func (s WebSite_uploadBlob_Results) HasStream() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s WebSite_uploadBlob_Results) SetStream(v util.ByteStream) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// WebSite_uploadBlob_Results_List is a list of WebSite_uploadBlob_Results.
type WebSite_uploadBlob_Results_List struct{ capnp.List }

// NewWebSite_uploadBlob_Results creates a new list of WebSite_uploadBlob_Results.
func NewWebSite_uploadBlob_Results_List(s *capnp.Segment, sz int32) (WebSite_uploadBlob_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return WebSite_uploadBlob_Results_List{l}, err
}

func (s WebSite_uploadBlob_Results_List) At(i int) WebSite_uploadBlob_Results {
	return WebSite_uploadBlob_Results{s.List.Struct(i)}
}

func (s WebSite_uploadBlob_Results_List) Set(i int, v WebSite_uploadBlob_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_uploadBlob_Results_Promise is a wrapper for a WebSite_uploadBlob_Results promised by a client call.
type WebSite_uploadBlob_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_uploadBlob_Results_Promise) Struct() (WebSite_uploadBlob_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_uploadBlob_Results{s}, err
}

func (p WebSite_uploadBlob_Results_Promise) Blob() util.Blob {
	return util.Blob{Client: p.Pipeline.GetPipeline(0).Client()}
}

func (p WebSite_uploadBlob_Results_Promise) Stream() util.ByteStream {
	return util.ByteStream{Client: p.Pipeline.GetPipeline(1).Client()}
}

type WebSite_getSubsite_Params struct{ capnp.Struct }

// WebSite_getSubsite_Params_TypeID is the unique identifier for the type WebSite_getSubsite_Params.
const WebSite_getSubsite_Params_TypeID = 0xdcbe913db7496644

func NewWebSite_getSubsite_Params(s *capnp.Segment) (WebSite_getSubsite_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getSubsite_Params{st}, err
}

func NewRootWebSite_getSubsite_Params(s *capnp.Segment) (WebSite_getSubsite_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getSubsite_Params{st}, err
}

func ReadRootWebSite_getSubsite_Params(msg *capnp.Message) (WebSite_getSubsite_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_getSubsite_Params{root.Struct()}, err
}

func (s WebSite_getSubsite_Params) String() string {
	str, _ := text.Marshal(0xdcbe913db7496644, s.Struct)
	return str
}

func (s WebSite_getSubsite_Params) Prefix() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s WebSite_getSubsite_Params) HasPrefix() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getSubsite_Params) PrefixBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s WebSite_getSubsite_Params) SetPrefix(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// WebSite_getSubsite_Params_List is a list of WebSite_getSubsite_Params.
type WebSite_getSubsite_Params_List struct{ capnp.List }

// NewWebSite_getSubsite_Params creates a new list of WebSite_getSubsite_Params.
func NewWebSite_getSubsite_Params_List(s *capnp.Segment, sz int32) (WebSite_getSubsite_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getSubsite_Params_List{l}, err
}

func (s WebSite_getSubsite_Params_List) At(i int) WebSite_getSubsite_Params {
	return WebSite_getSubsite_Params{s.List.Struct(i)}
}

func (s WebSite_getSubsite_Params_List) Set(i int, v WebSite_getSubsite_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getSubsite_Params_Promise is a wrapper for a WebSite_getSubsite_Params promised by a client call.
type WebSite_getSubsite_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_getSubsite_Params_Promise) Struct() (WebSite_getSubsite_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getSubsite_Params{s}, err
}

type WebSite_getSubsite_Results struct{ capnp.Struct }

// WebSite_getSubsite_Results_TypeID is the unique identifier for the type WebSite_getSubsite_Results.
const WebSite_getSubsite_Results_TypeID = 0x90dec3f1a5d9b591

func NewWebSite_getSubsite_Results(s *capnp.Segment) (WebSite_getSubsite_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getSubsite_Results{st}, err
}

func NewRootWebSite_getSubsite_Results(s *capnp.Segment) (WebSite_getSubsite_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_getSubsite_Results{st}, err
}

func ReadRootWebSite_getSubsite_Results(msg *capnp.Message) (WebSite_getSubsite_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_getSubsite_Results{root.Struct()}, err
}

func (s WebSite_getSubsite_Results) String() string {
	str, _ := text.Marshal(0x90dec3f1a5d9b591, s.Struct)
	return str
}

func (s WebSite_getSubsite_Results) Site() WebSite {
	p, _ := s.Struct.Ptr(0)
	return WebSite{Client: p.Interface().Client()}
}

func (s WebSite_getSubsite_Results) HasSite() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_getSubsite_Results) SetSite(v WebSite) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// WebSite_getSubsite_Results_List is a list of WebSite_getSubsite_Results.
type WebSite_getSubsite_Results_List struct{ capnp.List }

// NewWebSite_getSubsite_Results creates a new list of WebSite_getSubsite_Results.
func NewWebSite_getSubsite_Results_List(s *capnp.Segment, sz int32) (WebSite_getSubsite_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_getSubsite_Results_List{l}, err
}

func (s WebSite_getSubsite_Results_List) At(i int) WebSite_getSubsite_Results {
	return WebSite_getSubsite_Results{s.List.Struct(i)}
}

func (s WebSite_getSubsite_Results_List) Set(i int, v WebSite_getSubsite_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_getSubsite_Results_Promise is a wrapper for a WebSite_getSubsite_Results promised by a client call.
type WebSite_getSubsite_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_getSubsite_Results_Promise) Struct() (WebSite_getSubsite_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_getSubsite_Results{s}, err
}

func (p WebSite_getSubsite_Results_Promise) Site() WebSite {
	return WebSite{Client: p.Pipeline.GetPipeline(0).Client()}
}

type WebSite_listResources_Params struct{ capnp.Struct }

// WebSite_listResources_Params_TypeID is the unique identifier for the type WebSite_listResources_Params.
const WebSite_listResources_Params_TypeID = 0x87a72be8e04db694

func NewWebSite_listResources_Params(s *capnp.Segment) (WebSite_listResources_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return WebSite_listResources_Params{st}, err
}

func NewRootWebSite_listResources_Params(s *capnp.Segment) (WebSite_listResources_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return WebSite_listResources_Params{st}, err
}

func ReadRootWebSite_listResources_Params(msg *capnp.Message) (WebSite_listResources_Params, error) {
	root, err := msg.RootPtr()
	return WebSite_listResources_Params{root.Struct()}, err
}

func (s WebSite_listResources_Params) String() string {
	str, _ := text.Marshal(0x87a72be8e04db694, s.Struct)
	return str
}

func (s WebSite_listResources_Params) Shallow() bool {
	return s.Struct.Bit(0)
}

func (s WebSite_listResources_Params) SetShallow(v bool) {
	s.Struct.SetBit(0, v)
}

// WebSite_listResources_Params_List is a list of WebSite_listResources_Params.
type WebSite_listResources_Params_List struct{ capnp.List }

// NewWebSite_listResources_Params creates a new list of WebSite_listResources_Params.
func NewWebSite_listResources_Params_List(s *capnp.Segment, sz int32) (WebSite_listResources_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return WebSite_listResources_Params_List{l}, err
}

func (s WebSite_listResources_Params_List) At(i int) WebSite_listResources_Params {
	return WebSite_listResources_Params{s.List.Struct(i)}
}

func (s WebSite_listResources_Params_List) Set(i int, v WebSite_listResources_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_listResources_Params_Promise is a wrapper for a WebSite_listResources_Params promised by a client call.
type WebSite_listResources_Params_Promise struct{ *capnp.Pipeline }

func (p WebSite_listResources_Params_Promise) Struct() (WebSite_listResources_Params, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_listResources_Params{s}, err
}

type WebSite_listResources_Results struct{ capnp.Struct }

// WebSite_listResources_Results_TypeID is the unique identifier for the type WebSite_listResources_Results.
const WebSite_listResources_Results_TypeID = 0xf5e0920223f0b273

func NewWebSite_listResources_Results(s *capnp.Segment) (WebSite_listResources_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_listResources_Results{st}, err
}

func NewRootWebSite_listResources_Results(s *capnp.Segment) (WebSite_listResources_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WebSite_listResources_Results{st}, err
}

func ReadRootWebSite_listResources_Results(msg *capnp.Message) (WebSite_listResources_Results, error) {
	root, err := msg.RootPtr()
	return WebSite_listResources_Results{root.Struct()}, err
}

func (s WebSite_listResources_Results) String() string {
	str, _ := text.Marshal(0xf5e0920223f0b273, s.Struct)
	return str
}

func (s WebSite_listResources_Results) Names() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s WebSite_listResources_Results) HasNames() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WebSite_listResources_Results) SetNames(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewNames sets the names field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s WebSite_listResources_Results) NewNames(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// WebSite_listResources_Results_List is a list of WebSite_listResources_Results.
type WebSite_listResources_Results_List struct{ capnp.List }

// NewWebSite_listResources_Results creates a new list of WebSite_listResources_Results.
func NewWebSite_listResources_Results_List(s *capnp.Segment, sz int32) (WebSite_listResources_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WebSite_listResources_Results_List{l}, err
}

func (s WebSite_listResources_Results_List) At(i int) WebSite_listResources_Results {
	return WebSite_listResources_Results{s.List.Struct(i)}
}

func (s WebSite_listResources_Results_List) Set(i int, v WebSite_listResources_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// WebSite_listResources_Results_Promise is a wrapper for a WebSite_listResources_Results promised by a client call.
type WebSite_listResources_Results_Promise struct{ *capnp.Pipeline }

func (p WebSite_listResources_Results_Promise) Struct() (WebSite_listResources_Results, error) {
	s, err := p.Pipeline.Struct()
	return WebSite_listResources_Results{s}, err
}

const schema_d5d3e63bd0a552b6 = "x\xda\xc4U]h\x1cU\x14>\xe7\xcel'\xd9\xdd" +
	"$\xbd\xec\xd66`\x0c4\x0b\x9a\x95\x84\xb6\xb1Pc" +
	"\xdaM[\xebO\xa1\x92I\x1b\xaa\x81\"3\xc9mv" +
	"`vg\xdd\x99%\xdd\xd0\xfaS\x0c\x15QL\x8a\xc5" +
	"T\xa9\xa0\x10k\x1f\xd4Zl\x04\xb1B-\x08-\x88" +
	"\x11Q\x08&R\xd4\xaa\x04\x04\x15\x15|\xe9\x95;\xd9" +
	"\xf9I\xda5\xcd\x93O\xc9\x9e\xf9\xee9\xdf\xbd\xe7;" +
	"\xdf\xd9\x90\x90\xbb\xc9\xc6\xc8\xcf\x14@\x9d\x8e\xac\xe2g" +
	"\xcb\x17\xdfJ\xef|\xef(\xa8)D\x9e}\xb2\xbek" +
	"\xdb\\\xe34\xf4E\x14$\x00\x1d35\xbb\x110q" +
	"\xadf\x18\x90\xbf<\xb5\xe7\xea/w\x9f>\x06j\x13" +
	"\"\x80\xac\x00tl\xaf=.\x00jm\x06\x90\x8f\x9f" +
	"\x9f\x99\xfc\xfd\xd2wc@\x9b\x10 \x82\x02\xf0D\xed" +
	"Q\x018\xe2\x02^\x1f\xc8n~\xe8\xc0\xe8\xb90\xe0" +
	"T\xad.\x00g\\\xc0\xfc\xf9\x06\xc5z\xcd\xbe\x14\x06" +
	"\\\xae}^\x00f\\\xc0T\xcb\xdb\xab\xe7\xb3\xdfV" +
	"\x00.\x85\xad\xd1O\x11d~\xff+\xeb\xe6\xd6\x98]" +
	"\x9fU\x8e\x12\xf1\xa9-\xea\x16\xbf7*\xe8\xfb\x97\x13" +
	"\xf4\xf9p\xcb\xb6\xa7\x1f\xbc2;\x07\x91\x88@\x9e\x88" +
	"\xae\xc7\xc4dT\xfc\xfbF\xb4\x19\x01\xb96]<\xf4" +
	"cZ\x99\x0dU\x9a\x8c\xf5\xbb\x95\x0e>\xfc\xe1\xd6\xf1" +
	"Of\xc3$\xc7c#\xa2\xd2\xa9\x98 \xe9\xe7\xa6k" +
	"$>\xd5;9}\xdfO_}\x0d\x80\x89\x0b\xb1\xf9" +
	"\xc4\xe5\xd8Z\x80\x8e/c\xc701\x13W\x00\xf8\xdf" +
	"g\xb6\xa7G6\xbe\xfb}\xa8\xd0\xc5\xf8\x88(d\x9f" +
	"\xfb\xad\x85\x1c\xbf\xfaW\xb8\xd0;\xf1WE\xa1\x0bq" +
	"Q\xe8\xcf\xc3w\x9c\xfdf\xe7m\xd7\xc3\x80_\xe3\xee" +
	"\x9d\xffq\x01z2\xf5\xf1G][x\x18\xd0Xw" +
	"E\x00\xda\xea2\xc0a\x82\x0f3\xbd\xadP\xd2M\xd9" +
	"\xb0\xb3F~\xa8}@+\xe4\x0b\x9d\xfb\x99\xbe\xd7p" +
	"X\xfb\xae\xbcc8\xe5v\xdd\x1a,\x03\xa85\x92|" +
	";\xe7R\x12%\x00\xda\xba\x09@MI\xa8n \xd8" +
	"\x84\xd7EX\x06\xa0mi\x00\xf5.\x09\xd5{\x086" +
	"\xebe\x87\xd9X\x07\x04\xeb\x00\x1bt\xd3\xd2\x91\xf2\x97" +
	"\x1a\xc7\x0e\x9c\xbes\xf35\x00D*\x9e\xab\xc2 r" +
	"S\x06\xa6a;\xbd\xcc\xb6J\xc5\x01f\xa7z\xb4\xa2" +
	"\x96C[\x95%\x19@F\x00Z\xb7\xc3%\x86j\x92" +
	"\xe0SvV3Mk\x18\x11\x08\xe2\xb2\xa9\x87\x98\xb3" +
	"\xb7\xa4\xdb\x86\xc3R\xbd\xcc.\x99\x8e\x0d\xe1\xc4\xe9 " +
	"q\x83\x00!\x0d:\xbb\x84\xba\\-\x7f_\xd1ts" +
	"+\xa6cW\xcb]\xd0\x9c,\xc6\x81`\xfcV\x18\xbb" +
	"\x1d1\x98\xedS^\xc4y7\x80\x1a\x97P}\x94 " +
	"g\x15$\x00 \xe5'\xc7\xd7O|\xd0\xf7\xc7<\x00" +
	"t#\xc5fU&\x18\x0eR\\\xab\xca\x88\x88=\x12" +
	"b=\xb8\x7fV\x07C\xb3\xe4\xc2\xf1j\xf4\x1e\xb1\x9c" +
	"\x07\xacR~\xd0\xa7\xe9v\xcc\x06\xef\xe0M\xcf\x95\x0a" +
	"\xa6\xa5\x0d\xee0-=hD\x8d\x7f\xa9\xd6t 4" +
	"\x8a\x98\x14\xe6C\xdb:\x03\x99U\x11V\xc6v\x8aL" +
	"\xcb!\xe5\x8f\xff\xf0E\xeb\xf0\x96\xfd\x9f/\xbd\x85T" +
	"M\xf3\x8a\xe1\x94{\x10\xd5\xa4O\xe2\x88x\xd9\xc3\x12" +
	"\xaa\xcf\x85H\x8c\x8a\xe0\xb3\x12\xaac\x04)!I\xe1" +
	"\x99\xf4E\x11|ABu\x82 J\x88\x81\xc7\xd2\x13" +
	"i T\x96\x93\x18\x11\x19\xfb\x83\x8c<g\xe4\xd8\xbe" +
	"r\x81\x89^yR0\xb5\xfcPI\x1bZ\x14c\xf9" +
	"\x01k\xd0\xc8\x0f\x85b\x0db8y\x91\x0d\x1aE6" +
	"\xe0\x80\xb4\xcf\xbaAL\xff)O\xd1 )g/\x0f" +
	"\xf6f%\xb3\xd0\xd2\xb0\xec:\x039g\x0aEv\xd0" +
	"8t\x03\x07\xb24m\x83\xc8+$\x17\x88\x8cbg" +
	"f\xc1q\xd4uR\x04\xc07a\xf46\x07}\xbf\x13" +
	"\x08\x9dT00?\xf4\x96\x06=\xa9\x03\xa1\xe3\x0aJ" +
	"\xfe\x9a@\xcf\x00\xe9\xe8\x9b@\xe83\x0a\xca\xbe\xdf\xa2" +
	"\xb71h\xa9\x1f\x08\xcd)\x18\xf1\xad\x1d\xbdUF5" +
	"\xf1\xed1\x05\x89\xbf\xfe\xd0\xb3e\xba\xa7\x08\x84\xeeR" +
	"2\x0b\x0f\xd9\x8d\xdc\x1bOP\x0cf/\xfcv\xe7\x01" +
	"\xbd\x81\xa8\xc4=\xbd\x83d\xe9\x0b0\xf7mArX" +
	"7r\xcf\xf1\xa0\xd9\xf5\xbcn\xec\xc1\xe5Z\x19\x1a\xa0" +
	"JwV\xe4\xaa\xee\xd0I\x8b\x1djS\xa5\xa5)\x82" +
	"\xcdy-\xc7l\xcf\x14Dc\xebW\xe4T\x15\x0bX" +
	"\x91\xff\xad\xbae\x83\xf9?|\xf0\xdf\x00\x00\x00\xff\xff" +
	"#\xdb\xa0/"

func init() {
	schemas.Register(schema_d5d3e63bd0a552b6,
		0x82af432aa6c179b0,
		0x87a72be8e04db694,
		0x90dec3f1a5d9b591,
		0xb2855d483568639e,
		0xc3739b6f070fb5ea,
		0xc3db68ea10a823b6,
		0xc63c6c15dd189744,
		0xd019dd3e3c0e7e68,
		0xdc072ae47872d061,
		0xdcbe913db7496644,
		0xdddcca47803e2377,
		0xe2ae317a2a41a9f6,
		0xf5e0920223f0b273,
		0xfe1643d6b01e7cf4,
		0xff383cbabc241462)
}
