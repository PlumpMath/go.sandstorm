package hacksession

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	email "zenhack.net/go/sandstorm/capnp/email"
	grain "zenhack.net/go/sandstorm/capnp/grain"
	identity "zenhack.net/go/sandstorm/capnp/identity"
	ip "zenhack.net/go/sandstorm/capnp/ip"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type HackSessionContext struct{ Client capnp.Client }

func (c HackSessionContext) GetPublicId(ctx context.Context, params func(HackSessionContext_getPublicId_Params) error, opts ...capnp.CallOption) HackSessionContext_getPublicId_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_getPublicId_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      0,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getPublicId",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_getPublicId_Params{Struct: s}) }
	}
	return HackSessionContext_getPublicId_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) HttpGet(ctx context.Context, params func(HackSessionContext_httpGet_Params) error, opts ...capnp.CallOption) HackSessionContext_httpGet_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_httpGet_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      1,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "httpGet",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_httpGet_Params{Struct: s}) }
	}
	return HackSessionContext_httpGet_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) GetUserAddress(ctx context.Context, params func(HackSessionContext_getUserAddress_Params) error, opts ...capnp.CallOption) email.EmailAddress_Promise {
	if c.Client == nil {
		return email.EmailAddress_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      2,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getUserAddress",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_getUserAddress_Params{Struct: s}) }
	}
	return email.EmailAddress_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteGenerateApiToken(ctx context.Context, params func(HackSessionContext_obsoleteGenerateApiToken_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteGenerateApiToken_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteGenerateApiToken_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      3,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGenerateApiToken",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error {
			return params(HackSessionContext_obsoleteGenerateApiToken_Params{Struct: s})
		}
	}
	return HackSessionContext_obsoleteGenerateApiToken_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteListApiTokens(ctx context.Context, params func(HackSessionContext_obsoleteListApiTokens_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteListApiTokens_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteListApiTokens_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      4,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteListApiTokens",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_obsoleteListApiTokens_Params{Struct: s}) }
	}
	return HackSessionContext_obsoleteListApiTokens_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteRevokeApiToken(ctx context.Context, params func(HackSessionContext_obsoleteRevokeApiToken_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteRevokeApiToken_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteRevokeApiToken_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      5,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteRevokeApiToken",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_obsoleteRevokeApiToken_Params{Struct: s}) }
	}
	return HackSessionContext_obsoleteRevokeApiToken_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteGetIpNetwork(ctx context.Context, params func(HackSessionContext_obsoleteGetIpNetwork_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteGetIpNetwork_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteGetIpNetwork_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      6,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGetIpNetwork",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_obsoleteGetIpNetwork_Params{Struct: s}) }
	}
	return HackSessionContext_obsoleteGetIpNetwork_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ObsoleteGetIpInterface(ctx context.Context, params func(HackSessionContext_obsoleteGetIpInterface_Params) error, opts ...capnp.CallOption) HackSessionContext_obsoleteGetIpInterface_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_obsoleteGetIpInterface_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      7,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGetIpInterface",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_obsoleteGetIpInterface_Params{Struct: s}) }
	}
	return HackSessionContext_obsoleteGetIpInterface_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) GetUiViewForEndpoint(ctx context.Context, params func(HackSessionContext_getUiViewForEndpoint_Params) error, opts ...capnp.CallOption) HackSessionContext_getUiViewForEndpoint_Results_Promise {
	if c.Client == nil {
		return HackSessionContext_getUiViewForEndpoint_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      8,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getUiViewForEndpoint",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HackSessionContext_getUiViewForEndpoint_Params{Struct: s}) }
	}
	return HackSessionContext_getUiViewForEndpoint_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) GetSharedPermissions(ctx context.Context, params func(grain.SessionContext_getSharedPermissions_Params) error, opts ...capnp.CallOption) grain.SessionContext_getSharedPermissions_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_getSharedPermissions_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      0,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "getSharedPermissions",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_getSharedPermissions_Params{Struct: s}) }
	}
	return grain.SessionContext_getSharedPermissions_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) TieToUser(ctx context.Context, params func(grain.SessionContext_tieToUser_Params) error, opts ...capnp.CallOption) grain.SessionContext_tieToUser_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_tieToUser_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      1,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "tieToUser",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_tieToUser_Params{Struct: s}) }
	}
	return grain.SessionContext_tieToUser_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Offer(ctx context.Context, params func(grain.SessionContext_offer_Params) error, opts ...capnp.CallOption) grain.SessionContext_offer_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_offer_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      2,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "offer",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_offer_Params{Struct: s}) }
	}
	return grain.SessionContext_offer_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Request(ctx context.Context, params func(grain.SessionContext_request_Params) error, opts ...capnp.CallOption) grain.SessionContext_request_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_request_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      3,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "request",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_request_Params{Struct: s}) }
	}
	return grain.SessionContext_request_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) FulfillRequest(ctx context.Context, params func(grain.SessionContext_fulfillRequest_Params) error, opts ...capnp.CallOption) grain.SessionContext_fulfillRequest_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_fulfillRequest_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "fulfillRequest",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 4}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_fulfillRequest_Params{Struct: s}) }
	}
	return grain.SessionContext_fulfillRequest_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Close(ctx context.Context, params func(grain.SessionContext_close_Params) error, opts ...capnp.CallOption) grain.SessionContext_close_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_close_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      5,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "close",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_close_Params{Struct: s}) }
	}
	return grain.SessionContext_close_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) OpenView(ctx context.Context, params func(grain.SessionContext_openView_Params) error, opts ...capnp.CallOption) grain.SessionContext_openView_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_openView_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      6,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "openView",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_openView_Params{Struct: s}) }
	}
	return grain.SessionContext_openView_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) ClaimRequest(ctx context.Context, params func(grain.SessionContext_claimRequest_Params) error, opts ...capnp.CallOption) grain.SessionContext_claimRequest_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_claimRequest_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      7,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "claimRequest",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_claimRequest_Params{Struct: s}) }
	}
	return grain.SessionContext_claimRequest_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Activity(ctx context.Context, params func(grain.SessionContext_activity_Params) error, opts ...capnp.CallOption) grain.SessionContext_activity_Results_Promise {
	if c.Client == nil {
		return grain.SessionContext_activity_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      8,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "activity",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(grain.SessionContext_activity_Params{Struct: s}) }
	}
	return grain.SessionContext_activity_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) Send(ctx context.Context, params func(email.EmailSendPort_send_Params) error, opts ...capnp.CallOption) email.EmailSendPort_send_Results_Promise {
	if c.Client == nil {
		return email.EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailSendPort_send_Params{Struct: s}) }
	}
	return email.EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackSessionContext) HintAddress(ctx context.Context, params func(email.EmailSendPort_hintAddress_Params) error, opts ...capnp.CallOption) email.EmailSendPort_hintAddress_Results_Promise {
	if c.Client == nil {
		return email.EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailSendPort_hintAddress_Params{Struct: s}) }
	}
	return email.EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type HackSessionContext_Server interface {
	GetPublicId(HackSessionContext_getPublicId) error

	HttpGet(HackSessionContext_httpGet) error

	GetUserAddress(HackSessionContext_getUserAddress) error

	ObsoleteGenerateApiToken(HackSessionContext_obsoleteGenerateApiToken) error

	ObsoleteListApiTokens(HackSessionContext_obsoleteListApiTokens) error

	ObsoleteRevokeApiToken(HackSessionContext_obsoleteRevokeApiToken) error

	ObsoleteGetIpNetwork(HackSessionContext_obsoleteGetIpNetwork) error

	ObsoleteGetIpInterface(HackSessionContext_obsoleteGetIpInterface) error

	GetUiViewForEndpoint(HackSessionContext_getUiViewForEndpoint) error

	GetSharedPermissions(grain.SessionContext_getSharedPermissions) error

	TieToUser(grain.SessionContext_tieToUser) error

	Offer(grain.SessionContext_offer) error

	Request(grain.SessionContext_request) error

	FulfillRequest(grain.SessionContext_fulfillRequest) error

	Close(grain.SessionContext_close) error

	OpenView(grain.SessionContext_openView) error

	ClaimRequest(grain.SessionContext_claimRequest) error

	Activity(grain.SessionContext_activity) error

	Send(email.EmailSendPort_send) error

	HintAddress(email.EmailSendPort_hintAddress) error
}

func HackSessionContext_ServerToClient(s HackSessionContext_Server) HackSessionContext {
	c, _ := s.(server.Closer)
	return HackSessionContext{Client: server.New(HackSessionContext_Methods(nil, s), c)}
}

func HackSessionContext_Methods(methods []server.Method, s HackSessionContext_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 20)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      0,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getPublicId",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_getPublicId{c, opts, HackSessionContext_getPublicId_Params{Struct: p}, HackSessionContext_getPublicId_Results{Struct: r}}
			return s.GetPublicId(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 3},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      1,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "httpGet",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_httpGet{c, opts, HackSessionContext_httpGet_Params{Struct: p}, HackSessionContext_httpGet_Results{Struct: r}}
			return s.HttpGet(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      2,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getUserAddress",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_getUserAddress{c, opts, HackSessionContext_getUserAddress_Params{Struct: p}, email.EmailAddress{Struct: r}}
			return s.GetUserAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      3,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGenerateApiToken",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteGenerateApiToken{c, opts, HackSessionContext_obsoleteGenerateApiToken_Params{Struct: p}, HackSessionContext_obsoleteGenerateApiToken_Results{Struct: r}}
			return s.ObsoleteGenerateApiToken(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 3},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      4,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteListApiTokens",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteListApiTokens{c, opts, HackSessionContext_obsoleteListApiTokens_Params{Struct: p}, HackSessionContext_obsoleteListApiTokens_Results{Struct: r}}
			return s.ObsoleteListApiTokens(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      5,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteRevokeApiToken",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteRevokeApiToken{c, opts, HackSessionContext_obsoleteRevokeApiToken_Params{Struct: p}, HackSessionContext_obsoleteRevokeApiToken_Results{Struct: r}}
			return s.ObsoleteRevokeApiToken(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      6,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGetIpNetwork",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteGetIpNetwork{c, opts, HackSessionContext_obsoleteGetIpNetwork_Params{Struct: p}, HackSessionContext_obsoleteGetIpNetwork_Results{Struct: r}}
			return s.ObsoleteGetIpNetwork(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      7,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "obsoleteGetIpInterface",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_obsoleteGetIpInterface{c, opts, HackSessionContext_obsoleteGetIpInterface_Params{Struct: p}, HackSessionContext_obsoleteGetIpInterface_Results{Struct: r}}
			return s.ObsoleteGetIpInterface(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe14c1f5321159b8f,
			MethodID:      8,
			InterfaceName: "hack-session.capnp:HackSessionContext",
			MethodName:    "getUiViewForEndpoint",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HackSessionContext_getUiViewForEndpoint{c, opts, HackSessionContext_getUiViewForEndpoint_Params{Struct: p}, HackSessionContext_getUiViewForEndpoint_Results{Struct: r}}
			return s.GetUiViewForEndpoint(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      0,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "getSharedPermissions",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_getSharedPermissions{c, opts, grain.SessionContext_getSharedPermissions_Params{Struct: p}, grain.SessionContext_getSharedPermissions_Results{Struct: r}}
			return s.GetSharedPermissions(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      1,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "tieToUser",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_tieToUser{c, opts, grain.SessionContext_tieToUser_Params{Struct: p}, grain.SessionContext_tieToUser_Results{Struct: r}}
			return s.TieToUser(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      2,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "offer",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_offer{c, opts, grain.SessionContext_offer_Params{Struct: p}, grain.SessionContext_offer_Results{Struct: r}}
			return s.Offer(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      3,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "request",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_request{c, opts, grain.SessionContext_request_Params{Struct: p}, grain.SessionContext_request_Results{Struct: r}}
			return s.Request(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 2},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      4,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "fulfillRequest",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_fulfillRequest{c, opts, grain.SessionContext_fulfillRequest_Params{Struct: p}, grain.SessionContext_fulfillRequest_Results{Struct: r}}
			return s.FulfillRequest(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      5,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "close",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_close{c, opts, grain.SessionContext_close_Params{Struct: p}, grain.SessionContext_close_Results{Struct: r}}
			return s.Close(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      6,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "openView",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_openView{c, opts, grain.SessionContext_openView_Params{Struct: p}, grain.SessionContext_openView_Results{Struct: r}}
			return s.OpenView(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      7,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "claimRequest",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_claimRequest{c, opts, grain.SessionContext_claimRequest_Params{Struct: p}, grain.SessionContext_claimRequest_Results{Struct: r}}
			return s.ClaimRequest(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbf3e401d5a63f336,
			MethodID:      8,
			InterfaceName: "grain.capnp:SessionContext",
			MethodName:    "activity",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := grain.SessionContext_activity{c, opts, grain.SessionContext_activity_Params{Struct: p}, grain.SessionContext_activity_Results{Struct: r}}
			return s.Activity(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailSendPort_send{c, opts, email.EmailSendPort_send_Params{Struct: p}, email.EmailSendPort_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailSendPort_hintAddress{c, opts, email.EmailSendPort_hintAddress_Params{Struct: p}, email.EmailSendPort_hintAddress_Results{Struct: r}}
			return s.HintAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// HackSessionContext_getPublicId holds the arguments for a server call to HackSessionContext.getPublicId.
type HackSessionContext_getPublicId struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_getPublicId_Params
	Results HackSessionContext_getPublicId_Results
}

// HackSessionContext_httpGet holds the arguments for a server call to HackSessionContext.httpGet.
type HackSessionContext_httpGet struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_httpGet_Params
	Results HackSessionContext_httpGet_Results
}

// HackSessionContext_getUserAddress holds the arguments for a server call to HackSessionContext.getUserAddress.
type HackSessionContext_getUserAddress struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_getUserAddress_Params
	Results email.EmailAddress
}

// HackSessionContext_obsoleteGenerateApiToken holds the arguments for a server call to HackSessionContext.obsoleteGenerateApiToken.
type HackSessionContext_obsoleteGenerateApiToken struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteGenerateApiToken_Params
	Results HackSessionContext_obsoleteGenerateApiToken_Results
}

// HackSessionContext_obsoleteListApiTokens holds the arguments for a server call to HackSessionContext.obsoleteListApiTokens.
type HackSessionContext_obsoleteListApiTokens struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteListApiTokens_Params
	Results HackSessionContext_obsoleteListApiTokens_Results
}

// HackSessionContext_obsoleteRevokeApiToken holds the arguments for a server call to HackSessionContext.obsoleteRevokeApiToken.
type HackSessionContext_obsoleteRevokeApiToken struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteRevokeApiToken_Params
	Results HackSessionContext_obsoleteRevokeApiToken_Results
}

// HackSessionContext_obsoleteGetIpNetwork holds the arguments for a server call to HackSessionContext.obsoleteGetIpNetwork.
type HackSessionContext_obsoleteGetIpNetwork struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteGetIpNetwork_Params
	Results HackSessionContext_obsoleteGetIpNetwork_Results
}

// HackSessionContext_obsoleteGetIpInterface holds the arguments for a server call to HackSessionContext.obsoleteGetIpInterface.
type HackSessionContext_obsoleteGetIpInterface struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_obsoleteGetIpInterface_Params
	Results HackSessionContext_obsoleteGetIpInterface_Results
}

// HackSessionContext_getUiViewForEndpoint holds the arguments for a server call to HackSessionContext.getUiViewForEndpoint.
type HackSessionContext_getUiViewForEndpoint struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HackSessionContext_getUiViewForEndpoint_Params
	Results HackSessionContext_getUiViewForEndpoint_Results
}

type HackSessionContext_TokenInfo struct{ capnp.Struct }

// HackSessionContext_TokenInfo_TypeID is the unique identifier for the type HackSessionContext_TokenInfo.
const HackSessionContext_TokenInfo_TypeID = 0xf910658ae8c6674d

func NewHackSessionContext_TokenInfo(s *capnp.Segment) (HackSessionContext_TokenInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HackSessionContext_TokenInfo{st}, err
}

func NewRootHackSessionContext_TokenInfo(s *capnp.Segment) (HackSessionContext_TokenInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HackSessionContext_TokenInfo{st}, err
}

func ReadRootHackSessionContext_TokenInfo(msg *capnp.Message) (HackSessionContext_TokenInfo, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_TokenInfo{root.Struct()}, err
}

func (s HackSessionContext_TokenInfo) String() string {
	str, _ := text.Marshal(0xf910658ae8c6674d, s.Struct)
	return str
}

func (s HackSessionContext_TokenInfo) TokenId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_TokenInfo) HasTokenId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_TokenInfo) SetTokenId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_TokenInfo) Petname() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s HackSessionContext_TokenInfo) HasPetname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) PetnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s HackSessionContext_TokenInfo) SetPetname(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s HackSessionContext_TokenInfo) UserInfo() (identity.UserInfo, error) {
	p, err := s.Struct.Ptr(2)
	return identity.UserInfo{Struct: p.Struct()}, err
}

func (s HackSessionContext_TokenInfo) HasUserInfo() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_TokenInfo) SetUserInfo(v identity.UserInfo) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated identity.UserInfo struct, preferring placement in s's segment.
func (s HackSessionContext_TokenInfo) NewUserInfo() (identity.UserInfo, error) {
	ss, err := identity.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return identity.UserInfo{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// HackSessionContext_TokenInfo_List is a list of HackSessionContext_TokenInfo.
type HackSessionContext_TokenInfo_List struct{ capnp.List }

// NewHackSessionContext_TokenInfo creates a new list of HackSessionContext_TokenInfo.
func NewHackSessionContext_TokenInfo_List(s *capnp.Segment, sz int32) (HackSessionContext_TokenInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return HackSessionContext_TokenInfo_List{l}, err
}

func (s HackSessionContext_TokenInfo_List) At(i int) HackSessionContext_TokenInfo {
	return HackSessionContext_TokenInfo{s.List.Struct(i)}
}

func (s HackSessionContext_TokenInfo_List) Set(i int, v HackSessionContext_TokenInfo) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_TokenInfo_Promise is a wrapper for a HackSessionContext_TokenInfo promised by a client call.
type HackSessionContext_TokenInfo_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_TokenInfo_Promise) Struct() (HackSessionContext_TokenInfo, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_TokenInfo{s}, err
}

func (p HackSessionContext_TokenInfo_Promise) UserInfo() identity.UserInfo_Promise {
	return identity.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type HackSessionContext_getPublicId_Params struct{ capnp.Struct }

// HackSessionContext_getPublicId_Params_TypeID is the unique identifier for the type HackSessionContext_getPublicId_Params.
const HackSessionContext_getPublicId_Params_TypeID = 0xe706d835e9cd64af

func NewHackSessionContext_getPublicId_Params(s *capnp.Segment) (HackSessionContext_getPublicId_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_getPublicId_Params{st}, err
}

func NewRootHackSessionContext_getPublicId_Params(s *capnp.Segment) (HackSessionContext_getPublicId_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_getPublicId_Params{st}, err
}

func ReadRootHackSessionContext_getPublicId_Params(msg *capnp.Message) (HackSessionContext_getPublicId_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getPublicId_Params{root.Struct()}, err
}

func (s HackSessionContext_getPublicId_Params) String() string {
	str, _ := text.Marshal(0xe706d835e9cd64af, s.Struct)
	return str
}

// HackSessionContext_getPublicId_Params_List is a list of HackSessionContext_getPublicId_Params.
type HackSessionContext_getPublicId_Params_List struct{ capnp.List }

// NewHackSessionContext_getPublicId_Params creates a new list of HackSessionContext_getPublicId_Params.
func NewHackSessionContext_getPublicId_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getPublicId_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_getPublicId_Params_List{l}, err
}

func (s HackSessionContext_getPublicId_Params_List) At(i int) HackSessionContext_getPublicId_Params {
	return HackSessionContext_getPublicId_Params{s.List.Struct(i)}
}

func (s HackSessionContext_getPublicId_Params_List) Set(i int, v HackSessionContext_getPublicId_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getPublicId_Params_Promise is a wrapper for a HackSessionContext_getPublicId_Params promised by a client call.
type HackSessionContext_getPublicId_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getPublicId_Params_Promise) Struct() (HackSessionContext_getPublicId_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getPublicId_Params{s}, err
}

type HackSessionContext_getPublicId_Results struct{ capnp.Struct }

// HackSessionContext_getPublicId_Results_TypeID is the unique identifier for the type HackSessionContext_getPublicId_Results.
const HackSessionContext_getPublicId_Results_TypeID = 0xe96193c522f6c05d

func NewHackSessionContext_getPublicId_Results(s *capnp.Segment) (HackSessionContext_getPublicId_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return HackSessionContext_getPublicId_Results{st}, err
}

func NewRootHackSessionContext_getPublicId_Results(s *capnp.Segment) (HackSessionContext_getPublicId_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return HackSessionContext_getPublicId_Results{st}, err
}

func ReadRootHackSessionContext_getPublicId_Results(msg *capnp.Message) (HackSessionContext_getPublicId_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getPublicId_Results{root.Struct()}, err
}

func (s HackSessionContext_getPublicId_Results) String() string {
	str, _ := text.Marshal(0xe96193c522f6c05d, s.Struct)
	return str
}

func (s HackSessionContext_getPublicId_Results) PublicId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_getPublicId_Results) HasPublicId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) PublicIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_getPublicId_Results) SetPublicId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_getPublicId_Results) Hostname() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s HackSessionContext_getPublicId_Results) HasHostname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) HostnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s HackSessionContext_getPublicId_Results) SetHostname(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s HackSessionContext_getPublicId_Results) AutoUrl() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s HackSessionContext_getPublicId_Results) HasAutoUrl() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getPublicId_Results) AutoUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s HackSessionContext_getPublicId_Results) SetAutoUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s HackSessionContext_getPublicId_Results) IsDemoUser() bool {
	return s.Struct.Bit(0)
}

func (s HackSessionContext_getPublicId_Results) SetIsDemoUser(v bool) {
	s.Struct.SetBit(0, v)
}

// HackSessionContext_getPublicId_Results_List is a list of HackSessionContext_getPublicId_Results.
type HackSessionContext_getPublicId_Results_List struct{ capnp.List }

// NewHackSessionContext_getPublicId_Results creates a new list of HackSessionContext_getPublicId_Results.
func NewHackSessionContext_getPublicId_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_getPublicId_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return HackSessionContext_getPublicId_Results_List{l}, err
}

func (s HackSessionContext_getPublicId_Results_List) At(i int) HackSessionContext_getPublicId_Results {
	return HackSessionContext_getPublicId_Results{s.List.Struct(i)}
}

func (s HackSessionContext_getPublicId_Results_List) Set(i int, v HackSessionContext_getPublicId_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getPublicId_Results_Promise is a wrapper for a HackSessionContext_getPublicId_Results promised by a client call.
type HackSessionContext_getPublicId_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getPublicId_Results_Promise) Struct() (HackSessionContext_getPublicId_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getPublicId_Results{s}, err
}

type HackSessionContext_httpGet_Params struct{ capnp.Struct }

// HackSessionContext_httpGet_Params_TypeID is the unique identifier for the type HackSessionContext_httpGet_Params.
const HackSessionContext_httpGet_Params_TypeID = 0xe54437a7b8f52843

func NewHackSessionContext_httpGet_Params(s *capnp.Segment) (HackSessionContext_httpGet_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_httpGet_Params{st}, err
}

func NewRootHackSessionContext_httpGet_Params(s *capnp.Segment) (HackSessionContext_httpGet_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_httpGet_Params{st}, err
}

func ReadRootHackSessionContext_httpGet_Params(msg *capnp.Message) (HackSessionContext_httpGet_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_httpGet_Params{root.Struct()}, err
}

func (s HackSessionContext_httpGet_Params) String() string {
	str, _ := text.Marshal(0xe54437a7b8f52843, s.Struct)
	return str
}

func (s HackSessionContext_httpGet_Params) Url() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_httpGet_Params) HasUrl() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_httpGet_Params) UrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_httpGet_Params) SetUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// HackSessionContext_httpGet_Params_List is a list of HackSessionContext_httpGet_Params.
type HackSessionContext_httpGet_Params_List struct{ capnp.List }

// NewHackSessionContext_httpGet_Params creates a new list of HackSessionContext_httpGet_Params.
func NewHackSessionContext_httpGet_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_httpGet_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_httpGet_Params_List{l}, err
}

func (s HackSessionContext_httpGet_Params_List) At(i int) HackSessionContext_httpGet_Params {
	return HackSessionContext_httpGet_Params{s.List.Struct(i)}
}

func (s HackSessionContext_httpGet_Params_List) Set(i int, v HackSessionContext_httpGet_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_httpGet_Params_Promise is a wrapper for a HackSessionContext_httpGet_Params promised by a client call.
type HackSessionContext_httpGet_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_httpGet_Params_Promise) Struct() (HackSessionContext_httpGet_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_httpGet_Params{s}, err
}

type HackSessionContext_httpGet_Results struct{ capnp.Struct }

// HackSessionContext_httpGet_Results_TypeID is the unique identifier for the type HackSessionContext_httpGet_Results.
const HackSessionContext_httpGet_Results_TypeID = 0xb44df810894a44f4

func NewHackSessionContext_httpGet_Results(s *capnp.Segment) (HackSessionContext_httpGet_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return HackSessionContext_httpGet_Results{st}, err
}

func NewRootHackSessionContext_httpGet_Results(s *capnp.Segment) (HackSessionContext_httpGet_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return HackSessionContext_httpGet_Results{st}, err
}

func ReadRootHackSessionContext_httpGet_Results(msg *capnp.Message) (HackSessionContext_httpGet_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_httpGet_Results{root.Struct()}, err
}

func (s HackSessionContext_httpGet_Results) String() string {
	str, _ := text.Marshal(0xb44df810894a44f4, s.Struct)
	return str
}

func (s HackSessionContext_httpGet_Results) MimeType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_httpGet_Results) HasMimeType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_httpGet_Results) MimeTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_httpGet_Results) SetMimeType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_httpGet_Results) Content() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s HackSessionContext_httpGet_Results) HasContent() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_httpGet_Results) SetContent(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

// HackSessionContext_httpGet_Results_List is a list of HackSessionContext_httpGet_Results.
type HackSessionContext_httpGet_Results_List struct{ capnp.List }

// NewHackSessionContext_httpGet_Results creates a new list of HackSessionContext_httpGet_Results.
func NewHackSessionContext_httpGet_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_httpGet_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return HackSessionContext_httpGet_Results_List{l}, err
}

func (s HackSessionContext_httpGet_Results_List) At(i int) HackSessionContext_httpGet_Results {
	return HackSessionContext_httpGet_Results{s.List.Struct(i)}
}

func (s HackSessionContext_httpGet_Results_List) Set(i int, v HackSessionContext_httpGet_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_httpGet_Results_Promise is a wrapper for a HackSessionContext_httpGet_Results promised by a client call.
type HackSessionContext_httpGet_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_httpGet_Results_Promise) Struct() (HackSessionContext_httpGet_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_httpGet_Results{s}, err
}

type HackSessionContext_getUserAddress_Params struct{ capnp.Struct }

// HackSessionContext_getUserAddress_Params_TypeID is the unique identifier for the type HackSessionContext_getUserAddress_Params.
const HackSessionContext_getUserAddress_Params_TypeID = 0xa15e445037d1834c

func NewHackSessionContext_getUserAddress_Params(s *capnp.Segment) (HackSessionContext_getUserAddress_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_getUserAddress_Params{st}, err
}

func NewRootHackSessionContext_getUserAddress_Params(s *capnp.Segment) (HackSessionContext_getUserAddress_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_getUserAddress_Params{st}, err
}

func ReadRootHackSessionContext_getUserAddress_Params(msg *capnp.Message) (HackSessionContext_getUserAddress_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getUserAddress_Params{root.Struct()}, err
}

func (s HackSessionContext_getUserAddress_Params) String() string {
	str, _ := text.Marshal(0xa15e445037d1834c, s.Struct)
	return str
}

// HackSessionContext_getUserAddress_Params_List is a list of HackSessionContext_getUserAddress_Params.
type HackSessionContext_getUserAddress_Params_List struct{ capnp.List }

// NewHackSessionContext_getUserAddress_Params creates a new list of HackSessionContext_getUserAddress_Params.
func NewHackSessionContext_getUserAddress_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getUserAddress_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_getUserAddress_Params_List{l}, err
}

func (s HackSessionContext_getUserAddress_Params_List) At(i int) HackSessionContext_getUserAddress_Params {
	return HackSessionContext_getUserAddress_Params{s.List.Struct(i)}
}

func (s HackSessionContext_getUserAddress_Params_List) Set(i int, v HackSessionContext_getUserAddress_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getUserAddress_Params_Promise is a wrapper for a HackSessionContext_getUserAddress_Params promised by a client call.
type HackSessionContext_getUserAddress_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getUserAddress_Params_Promise) Struct() (HackSessionContext_getUserAddress_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getUserAddress_Params{s}, err
}

type HackSessionContext_obsoleteGenerateApiToken_Params struct{ capnp.Struct }

// HackSessionContext_obsoleteGenerateApiToken_Params_TypeID is the unique identifier for the type HackSessionContext_obsoleteGenerateApiToken_Params.
const HackSessionContext_obsoleteGenerateApiToken_Params_TypeID = 0x837afa61d880beb6

func NewHackSessionContext_obsoleteGenerateApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return HackSessionContext_obsoleteGenerateApiToken_Params{st}, err
}

func NewRootHackSessionContext_obsoleteGenerateApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return HackSessionContext_obsoleteGenerateApiToken_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteGenerateApiToken_Params(msg *capnp.Message) (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGenerateApiToken_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) String() string {
	str, _ := text.Marshal(0x837afa61d880beb6, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) Petname() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) HasPetname() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) PetnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) SetPetname(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) UserInfo() (identity.UserInfo, error) {
	p, err := s.Struct.Ptr(1)
	return identity.UserInfo{Struct: p.Struct()}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) HasUserInfo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) SetUserInfo(v identity.UserInfo) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewUserInfo sets the userInfo field to a newly
// allocated identity.UserInfo struct, preferring placement in s's segment.
func (s HackSessionContext_obsoleteGenerateApiToken_Params) NewUserInfo() (identity.UserInfo, error) {
	ss, err := identity.NewUserInfo(s.Struct.Segment())
	if err != nil {
		return identity.UserInfo{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) Expires() uint64 {
	return s.Struct.Uint64(0)
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params) SetExpires(v uint64) {
	s.Struct.SetUint64(0, v)
}

// HackSessionContext_obsoleteGenerateApiToken_Params_List is a list of HackSessionContext_obsoleteGenerateApiToken_Params.
type HackSessionContext_obsoleteGenerateApiToken_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGenerateApiToken_Params creates a new list of HackSessionContext_obsoleteGenerateApiToken_Params.
func NewHackSessionContext_obsoleteGenerateApiToken_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGenerateApiToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return HackSessionContext_obsoleteGenerateApiToken_Params_List{l}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params_List) At(i int) HackSessionContext_obsoleteGenerateApiToken_Params {
	return HackSessionContext_obsoleteGenerateApiToken_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGenerateApiToken_Params_List) Set(i int, v HackSessionContext_obsoleteGenerateApiToken_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGenerateApiToken_Params_Promise is a wrapper for a HackSessionContext_obsoleteGenerateApiToken_Params promised by a client call.
type HackSessionContext_obsoleteGenerateApiToken_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGenerateApiToken_Params_Promise) Struct() (HackSessionContext_obsoleteGenerateApiToken_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGenerateApiToken_Params{s}, err
}

func (p HackSessionContext_obsoleteGenerateApiToken_Params_Promise) UserInfo() identity.UserInfo_Promise {
	return identity.UserInfo_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HackSessionContext_obsoleteGenerateApiToken_Results struct{ capnp.Struct }

// HackSessionContext_obsoleteGenerateApiToken_Results_TypeID is the unique identifier for the type HackSessionContext_obsoleteGenerateApiToken_Results.
const HackSessionContext_obsoleteGenerateApiToken_Results_TypeID = 0xc9973f31a90887f9

func NewHackSessionContext_obsoleteGenerateApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HackSessionContext_obsoleteGenerateApiToken_Results{st}, err
}

func NewRootHackSessionContext_obsoleteGenerateApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HackSessionContext_obsoleteGenerateApiToken_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteGenerateApiToken_Results(msg *capnp.Message) (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGenerateApiToken_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) String() string {
	str, _ := text.Marshal(0xc9973f31a90887f9, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) Token() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) HasToken() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) TokenBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetToken(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) EndpointUrl() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) HasEndpointUrl() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) EndpointUrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetEndpointUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) TokenId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) HasTokenId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results) SetTokenId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// HackSessionContext_obsoleteGenerateApiToken_Results_List is a list of HackSessionContext_obsoleteGenerateApiToken_Results.
type HackSessionContext_obsoleteGenerateApiToken_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGenerateApiToken_Results creates a new list of HackSessionContext_obsoleteGenerateApiToken_Results.
func NewHackSessionContext_obsoleteGenerateApiToken_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGenerateApiToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return HackSessionContext_obsoleteGenerateApiToken_Results_List{l}, err
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results_List) At(i int) HackSessionContext_obsoleteGenerateApiToken_Results {
	return HackSessionContext_obsoleteGenerateApiToken_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGenerateApiToken_Results_List) Set(i int, v HackSessionContext_obsoleteGenerateApiToken_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGenerateApiToken_Results_Promise is a wrapper for a HackSessionContext_obsoleteGenerateApiToken_Results promised by a client call.
type HackSessionContext_obsoleteGenerateApiToken_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGenerateApiToken_Results_Promise) Struct() (HackSessionContext_obsoleteGenerateApiToken_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGenerateApiToken_Results{s}, err
}

type HackSessionContext_obsoleteListApiTokens_Params struct{ capnp.Struct }

// HackSessionContext_obsoleteListApiTokens_Params_TypeID is the unique identifier for the type HackSessionContext_obsoleteListApiTokens_Params.
const HackSessionContext_obsoleteListApiTokens_Params_TypeID = 0xfe777c71f871f413

func NewHackSessionContext_obsoleteListApiTokens_Params(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteListApiTokens_Params{st}, err
}

func NewRootHackSessionContext_obsoleteListApiTokens_Params(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteListApiTokens_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteListApiTokens_Params(msg *capnp.Message) (HackSessionContext_obsoleteListApiTokens_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteListApiTokens_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Params) String() string {
	str, _ := text.Marshal(0xfe777c71f871f413, s.Struct)
	return str
}

// HackSessionContext_obsoleteListApiTokens_Params_List is a list of HackSessionContext_obsoleteListApiTokens_Params.
type HackSessionContext_obsoleteListApiTokens_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteListApiTokens_Params creates a new list of HackSessionContext_obsoleteListApiTokens_Params.
func NewHackSessionContext_obsoleteListApiTokens_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteListApiTokens_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_obsoleteListApiTokens_Params_List{l}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Params_List) At(i int) HackSessionContext_obsoleteListApiTokens_Params {
	return HackSessionContext_obsoleteListApiTokens_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteListApiTokens_Params_List) Set(i int, v HackSessionContext_obsoleteListApiTokens_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteListApiTokens_Params_Promise is a wrapper for a HackSessionContext_obsoleteListApiTokens_Params promised by a client call.
type HackSessionContext_obsoleteListApiTokens_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteListApiTokens_Params_Promise) Struct() (HackSessionContext_obsoleteListApiTokens_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteListApiTokens_Params{s}, err
}

type HackSessionContext_obsoleteListApiTokens_Results struct{ capnp.Struct }

// HackSessionContext_obsoleteListApiTokens_Results_TypeID is the unique identifier for the type HackSessionContext_obsoleteListApiTokens_Results.
const HackSessionContext_obsoleteListApiTokens_Results_TypeID = 0xe9e4890dae20b03c

func NewHackSessionContext_obsoleteListApiTokens_Results(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteListApiTokens_Results{st}, err
}

func NewRootHackSessionContext_obsoleteListApiTokens_Results(s *capnp.Segment) (HackSessionContext_obsoleteListApiTokens_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteListApiTokens_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteListApiTokens_Results(msg *capnp.Message) (HackSessionContext_obsoleteListApiTokens_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteListApiTokens_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Results) String() string {
	str, _ := text.Marshal(0xe9e4890dae20b03c, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteListApiTokens_Results) Tokens() (HackSessionContext_TokenInfo_List, error) {
	p, err := s.Struct.Ptr(0)
	return HackSessionContext_TokenInfo_List{List: p.List()}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Results) HasTokens() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteListApiTokens_Results) SetTokens(v HackSessionContext_TokenInfo_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTokens sets the tokens field to a newly
// allocated HackSessionContext_TokenInfo_List, preferring placement in s's segment.
func (s HackSessionContext_obsoleteListApiTokens_Results) NewTokens(n int32) (HackSessionContext_TokenInfo_List, error) {
	l, err := NewHackSessionContext_TokenInfo_List(s.Struct.Segment(), n)
	if err != nil {
		return HackSessionContext_TokenInfo_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HackSessionContext_obsoleteListApiTokens_Results_List is a list of HackSessionContext_obsoleteListApiTokens_Results.
type HackSessionContext_obsoleteListApiTokens_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteListApiTokens_Results creates a new list of HackSessionContext_obsoleteListApiTokens_Results.
func NewHackSessionContext_obsoleteListApiTokens_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteListApiTokens_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_obsoleteListApiTokens_Results_List{l}, err
}

func (s HackSessionContext_obsoleteListApiTokens_Results_List) At(i int) HackSessionContext_obsoleteListApiTokens_Results {
	return HackSessionContext_obsoleteListApiTokens_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteListApiTokens_Results_List) Set(i int, v HackSessionContext_obsoleteListApiTokens_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteListApiTokens_Results_Promise is a wrapper for a HackSessionContext_obsoleteListApiTokens_Results promised by a client call.
type HackSessionContext_obsoleteListApiTokens_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteListApiTokens_Results_Promise) Struct() (HackSessionContext_obsoleteListApiTokens_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteListApiTokens_Results{s}, err
}

type HackSessionContext_obsoleteRevokeApiToken_Params struct{ capnp.Struct }

// HackSessionContext_obsoleteRevokeApiToken_Params_TypeID is the unique identifier for the type HackSessionContext_obsoleteRevokeApiToken_Params.
const HackSessionContext_obsoleteRevokeApiToken_Params_TypeID = 0x845e8081686d8a0f

func NewHackSessionContext_obsoleteRevokeApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteRevokeApiToken_Params{st}, err
}

func NewRootHackSessionContext_obsoleteRevokeApiToken_Params(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteRevokeApiToken_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteRevokeApiToken_Params(msg *capnp.Message) (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteRevokeApiToken_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) String() string {
	str, _ := text.Marshal(0x845e8081686d8a0f, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) TokenId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) HasTokenId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) TokenIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params) SetTokenId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// HackSessionContext_obsoleteRevokeApiToken_Params_List is a list of HackSessionContext_obsoleteRevokeApiToken_Params.
type HackSessionContext_obsoleteRevokeApiToken_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteRevokeApiToken_Params creates a new list of HackSessionContext_obsoleteRevokeApiToken_Params.
func NewHackSessionContext_obsoleteRevokeApiToken_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteRevokeApiToken_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_obsoleteRevokeApiToken_Params_List{l}, err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params_List) At(i int) HackSessionContext_obsoleteRevokeApiToken_Params {
	return HackSessionContext_obsoleteRevokeApiToken_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteRevokeApiToken_Params_List) Set(i int, v HackSessionContext_obsoleteRevokeApiToken_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteRevokeApiToken_Params_Promise is a wrapper for a HackSessionContext_obsoleteRevokeApiToken_Params promised by a client call.
type HackSessionContext_obsoleteRevokeApiToken_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteRevokeApiToken_Params_Promise) Struct() (HackSessionContext_obsoleteRevokeApiToken_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteRevokeApiToken_Params{s}, err
}

type HackSessionContext_obsoleteRevokeApiToken_Results struct{ capnp.Struct }

// HackSessionContext_obsoleteRevokeApiToken_Results_TypeID is the unique identifier for the type HackSessionContext_obsoleteRevokeApiToken_Results.
const HackSessionContext_obsoleteRevokeApiToken_Results_TypeID = 0xaea67098dc479ce6

func NewHackSessionContext_obsoleteRevokeApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteRevokeApiToken_Results{st}, err
}

func NewRootHackSessionContext_obsoleteRevokeApiToken_Results(s *capnp.Segment) (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteRevokeApiToken_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteRevokeApiToken_Results(msg *capnp.Message) (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteRevokeApiToken_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Results) String() string {
	str, _ := text.Marshal(0xaea67098dc479ce6, s.Struct)
	return str
}

// HackSessionContext_obsoleteRevokeApiToken_Results_List is a list of HackSessionContext_obsoleteRevokeApiToken_Results.
type HackSessionContext_obsoleteRevokeApiToken_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteRevokeApiToken_Results creates a new list of HackSessionContext_obsoleteRevokeApiToken_Results.
func NewHackSessionContext_obsoleteRevokeApiToken_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteRevokeApiToken_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_obsoleteRevokeApiToken_Results_List{l}, err
}

func (s HackSessionContext_obsoleteRevokeApiToken_Results_List) At(i int) HackSessionContext_obsoleteRevokeApiToken_Results {
	return HackSessionContext_obsoleteRevokeApiToken_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteRevokeApiToken_Results_List) Set(i int, v HackSessionContext_obsoleteRevokeApiToken_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteRevokeApiToken_Results_Promise is a wrapper for a HackSessionContext_obsoleteRevokeApiToken_Results promised by a client call.
type HackSessionContext_obsoleteRevokeApiToken_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteRevokeApiToken_Results_Promise) Struct() (HackSessionContext_obsoleteRevokeApiToken_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteRevokeApiToken_Results{s}, err
}

type HackSessionContext_obsoleteGetIpNetwork_Params struct{ capnp.Struct }

// HackSessionContext_obsoleteGetIpNetwork_Params_TypeID is the unique identifier for the type HackSessionContext_obsoleteGetIpNetwork_Params.
const HackSessionContext_obsoleteGetIpNetwork_Params_TypeID = 0x97f9d7a51ffe0741

func NewHackSessionContext_obsoleteGetIpNetwork_Params(s *capnp.Segment) (HackSessionContext_obsoleteGetIpNetwork_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteGetIpNetwork_Params{st}, err
}

func NewRootHackSessionContext_obsoleteGetIpNetwork_Params(s *capnp.Segment) (HackSessionContext_obsoleteGetIpNetwork_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteGetIpNetwork_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteGetIpNetwork_Params(msg *capnp.Message) (HackSessionContext_obsoleteGetIpNetwork_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGetIpNetwork_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGetIpNetwork_Params) String() string {
	str, _ := text.Marshal(0x97f9d7a51ffe0741, s.Struct)
	return str
}

// HackSessionContext_obsoleteGetIpNetwork_Params_List is a list of HackSessionContext_obsoleteGetIpNetwork_Params.
type HackSessionContext_obsoleteGetIpNetwork_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGetIpNetwork_Params creates a new list of HackSessionContext_obsoleteGetIpNetwork_Params.
func NewHackSessionContext_obsoleteGetIpNetwork_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGetIpNetwork_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_obsoleteGetIpNetwork_Params_List{l}, err
}

func (s HackSessionContext_obsoleteGetIpNetwork_Params_List) At(i int) HackSessionContext_obsoleteGetIpNetwork_Params {
	return HackSessionContext_obsoleteGetIpNetwork_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGetIpNetwork_Params_List) Set(i int, v HackSessionContext_obsoleteGetIpNetwork_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGetIpNetwork_Params_Promise is a wrapper for a HackSessionContext_obsoleteGetIpNetwork_Params promised by a client call.
type HackSessionContext_obsoleteGetIpNetwork_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGetIpNetwork_Params_Promise) Struct() (HackSessionContext_obsoleteGetIpNetwork_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGetIpNetwork_Params{s}, err
}

type HackSessionContext_obsoleteGetIpNetwork_Results struct{ capnp.Struct }

// HackSessionContext_obsoleteGetIpNetwork_Results_TypeID is the unique identifier for the type HackSessionContext_obsoleteGetIpNetwork_Results.
const HackSessionContext_obsoleteGetIpNetwork_Results_TypeID = 0xa9502e5fdabb8b07

func NewHackSessionContext_obsoleteGetIpNetwork_Results(s *capnp.Segment) (HackSessionContext_obsoleteGetIpNetwork_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteGetIpNetwork_Results{st}, err
}

func NewRootHackSessionContext_obsoleteGetIpNetwork_Results(s *capnp.Segment) (HackSessionContext_obsoleteGetIpNetwork_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteGetIpNetwork_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteGetIpNetwork_Results(msg *capnp.Message) (HackSessionContext_obsoleteGetIpNetwork_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGetIpNetwork_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results) String() string {
	str, _ := text.Marshal(0xa9502e5fdabb8b07, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results) Network() ip.IpNetwork {
	p, _ := s.Struct.Ptr(0)
	return ip.IpNetwork{Client: p.Interface().Client()}
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results) HasNetwork() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results) SetNetwork(v ip.IpNetwork) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// HackSessionContext_obsoleteGetIpNetwork_Results_List is a list of HackSessionContext_obsoleteGetIpNetwork_Results.
type HackSessionContext_obsoleteGetIpNetwork_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGetIpNetwork_Results creates a new list of HackSessionContext_obsoleteGetIpNetwork_Results.
func NewHackSessionContext_obsoleteGetIpNetwork_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGetIpNetwork_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_obsoleteGetIpNetwork_Results_List{l}, err
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results_List) At(i int) HackSessionContext_obsoleteGetIpNetwork_Results {
	return HackSessionContext_obsoleteGetIpNetwork_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGetIpNetwork_Results_List) Set(i int, v HackSessionContext_obsoleteGetIpNetwork_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGetIpNetwork_Results_Promise is a wrapper for a HackSessionContext_obsoleteGetIpNetwork_Results promised by a client call.
type HackSessionContext_obsoleteGetIpNetwork_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGetIpNetwork_Results_Promise) Struct() (HackSessionContext_obsoleteGetIpNetwork_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGetIpNetwork_Results{s}, err
}

func (p HackSessionContext_obsoleteGetIpNetwork_Results_Promise) Network() ip.IpNetwork {
	return ip.IpNetwork{Client: p.Pipeline.GetPipeline(0).Client()}
}

type HackSessionContext_obsoleteGetIpInterface_Params struct{ capnp.Struct }

// HackSessionContext_obsoleteGetIpInterface_Params_TypeID is the unique identifier for the type HackSessionContext_obsoleteGetIpInterface_Params.
const HackSessionContext_obsoleteGetIpInterface_Params_TypeID = 0xecebff1662ba10a1

func NewHackSessionContext_obsoleteGetIpInterface_Params(s *capnp.Segment) (HackSessionContext_obsoleteGetIpInterface_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteGetIpInterface_Params{st}, err
}

func NewRootHackSessionContext_obsoleteGetIpInterface_Params(s *capnp.Segment) (HackSessionContext_obsoleteGetIpInterface_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HackSessionContext_obsoleteGetIpInterface_Params{st}, err
}

func ReadRootHackSessionContext_obsoleteGetIpInterface_Params(msg *capnp.Message) (HackSessionContext_obsoleteGetIpInterface_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGetIpInterface_Params{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGetIpInterface_Params) String() string {
	str, _ := text.Marshal(0xecebff1662ba10a1, s.Struct)
	return str
}

// HackSessionContext_obsoleteGetIpInterface_Params_List is a list of HackSessionContext_obsoleteGetIpInterface_Params.
type HackSessionContext_obsoleteGetIpInterface_Params_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGetIpInterface_Params creates a new list of HackSessionContext_obsoleteGetIpInterface_Params.
func NewHackSessionContext_obsoleteGetIpInterface_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGetIpInterface_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HackSessionContext_obsoleteGetIpInterface_Params_List{l}, err
}

func (s HackSessionContext_obsoleteGetIpInterface_Params_List) At(i int) HackSessionContext_obsoleteGetIpInterface_Params {
	return HackSessionContext_obsoleteGetIpInterface_Params{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGetIpInterface_Params_List) Set(i int, v HackSessionContext_obsoleteGetIpInterface_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGetIpInterface_Params_Promise is a wrapper for a HackSessionContext_obsoleteGetIpInterface_Params promised by a client call.
type HackSessionContext_obsoleteGetIpInterface_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGetIpInterface_Params_Promise) Struct() (HackSessionContext_obsoleteGetIpInterface_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGetIpInterface_Params{s}, err
}

type HackSessionContext_obsoleteGetIpInterface_Results struct{ capnp.Struct }

// HackSessionContext_obsoleteGetIpInterface_Results_TypeID is the unique identifier for the type HackSessionContext_obsoleteGetIpInterface_Results.
const HackSessionContext_obsoleteGetIpInterface_Results_TypeID = 0xb9147a48c12c807d

func NewHackSessionContext_obsoleteGetIpInterface_Results(s *capnp.Segment) (HackSessionContext_obsoleteGetIpInterface_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteGetIpInterface_Results{st}, err
}

func NewRootHackSessionContext_obsoleteGetIpInterface_Results(s *capnp.Segment) (HackSessionContext_obsoleteGetIpInterface_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_obsoleteGetIpInterface_Results{st}, err
}

func ReadRootHackSessionContext_obsoleteGetIpInterface_Results(msg *capnp.Message) (HackSessionContext_obsoleteGetIpInterface_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_obsoleteGetIpInterface_Results{root.Struct()}, err
}

func (s HackSessionContext_obsoleteGetIpInterface_Results) String() string {
	str, _ := text.Marshal(0xb9147a48c12c807d, s.Struct)
	return str
}

func (s HackSessionContext_obsoleteGetIpInterface_Results) Interface() ip.IpInterface {
	p, _ := s.Struct.Ptr(0)
	return ip.IpInterface{Client: p.Interface().Client()}
}

func (s HackSessionContext_obsoleteGetIpInterface_Results) HasInterface() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_obsoleteGetIpInterface_Results) SetInterface(v ip.IpInterface) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// HackSessionContext_obsoleteGetIpInterface_Results_List is a list of HackSessionContext_obsoleteGetIpInterface_Results.
type HackSessionContext_obsoleteGetIpInterface_Results_List struct{ capnp.List }

// NewHackSessionContext_obsoleteGetIpInterface_Results creates a new list of HackSessionContext_obsoleteGetIpInterface_Results.
func NewHackSessionContext_obsoleteGetIpInterface_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_obsoleteGetIpInterface_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_obsoleteGetIpInterface_Results_List{l}, err
}

func (s HackSessionContext_obsoleteGetIpInterface_Results_List) At(i int) HackSessionContext_obsoleteGetIpInterface_Results {
	return HackSessionContext_obsoleteGetIpInterface_Results{s.List.Struct(i)}
}

func (s HackSessionContext_obsoleteGetIpInterface_Results_List) Set(i int, v HackSessionContext_obsoleteGetIpInterface_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_obsoleteGetIpInterface_Results_Promise is a wrapper for a HackSessionContext_obsoleteGetIpInterface_Results promised by a client call.
type HackSessionContext_obsoleteGetIpInterface_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_obsoleteGetIpInterface_Results_Promise) Struct() (HackSessionContext_obsoleteGetIpInterface_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_obsoleteGetIpInterface_Results{s}, err
}

func (p HackSessionContext_obsoleteGetIpInterface_Results_Promise) Interface() ip.IpInterface {
	return ip.IpInterface{Client: p.Pipeline.GetPipeline(0).Client()}
}

type HackSessionContext_getUiViewForEndpoint_Params struct{ capnp.Struct }

// HackSessionContext_getUiViewForEndpoint_Params_TypeID is the unique identifier for the type HackSessionContext_getUiViewForEndpoint_Params.
const HackSessionContext_getUiViewForEndpoint_Params_TypeID = 0xb45bb2d206694fe6

func NewHackSessionContext_getUiViewForEndpoint_Params(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_getUiViewForEndpoint_Params{st}, err
}

func NewRootHackSessionContext_getUiViewForEndpoint_Params(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_getUiViewForEndpoint_Params{st}, err
}

func ReadRootHackSessionContext_getUiViewForEndpoint_Params(msg *capnp.Message) (HackSessionContext_getUiViewForEndpoint_Params, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getUiViewForEndpoint_Params{root.Struct()}, err
}

func (s HackSessionContext_getUiViewForEndpoint_Params) String() string {
	str, _ := text.Marshal(0xb45bb2d206694fe6, s.Struct)
	return str
}

func (s HackSessionContext_getUiViewForEndpoint_Params) Url() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HackSessionContext_getUiViewForEndpoint_Params) HasUrl() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getUiViewForEndpoint_Params) UrlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HackSessionContext_getUiViewForEndpoint_Params) SetUrl(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// HackSessionContext_getUiViewForEndpoint_Params_List is a list of HackSessionContext_getUiViewForEndpoint_Params.
type HackSessionContext_getUiViewForEndpoint_Params_List struct{ capnp.List }

// NewHackSessionContext_getUiViewForEndpoint_Params creates a new list of HackSessionContext_getUiViewForEndpoint_Params.
func NewHackSessionContext_getUiViewForEndpoint_Params_List(s *capnp.Segment, sz int32) (HackSessionContext_getUiViewForEndpoint_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_getUiViewForEndpoint_Params_List{l}, err
}

func (s HackSessionContext_getUiViewForEndpoint_Params_List) At(i int) HackSessionContext_getUiViewForEndpoint_Params {
	return HackSessionContext_getUiViewForEndpoint_Params{s.List.Struct(i)}
}

func (s HackSessionContext_getUiViewForEndpoint_Params_List) Set(i int, v HackSessionContext_getUiViewForEndpoint_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getUiViewForEndpoint_Params_Promise is a wrapper for a HackSessionContext_getUiViewForEndpoint_Params promised by a client call.
type HackSessionContext_getUiViewForEndpoint_Params_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getUiViewForEndpoint_Params_Promise) Struct() (HackSessionContext_getUiViewForEndpoint_Params, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getUiViewForEndpoint_Params{s}, err
}

type HackSessionContext_getUiViewForEndpoint_Results struct{ capnp.Struct }

// HackSessionContext_getUiViewForEndpoint_Results_TypeID is the unique identifier for the type HackSessionContext_getUiViewForEndpoint_Results.
const HackSessionContext_getUiViewForEndpoint_Results_TypeID = 0xfdc944999b9296df

func NewHackSessionContext_getUiViewForEndpoint_Results(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_getUiViewForEndpoint_Results{st}, err
}

func NewRootHackSessionContext_getUiViewForEndpoint_Results(s *capnp.Segment) (HackSessionContext_getUiViewForEndpoint_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HackSessionContext_getUiViewForEndpoint_Results{st}, err
}

func ReadRootHackSessionContext_getUiViewForEndpoint_Results(msg *capnp.Message) (HackSessionContext_getUiViewForEndpoint_Results, error) {
	root, err := msg.RootPtr()
	return HackSessionContext_getUiViewForEndpoint_Results{root.Struct()}, err
}

func (s HackSessionContext_getUiViewForEndpoint_Results) String() string {
	str, _ := text.Marshal(0xfdc944999b9296df, s.Struct)
	return str
}

func (s HackSessionContext_getUiViewForEndpoint_Results) View() grain.UiView {
	p, _ := s.Struct.Ptr(0)
	return grain.UiView{Client: p.Interface().Client()}
}

func (s HackSessionContext_getUiViewForEndpoint_Results) HasView() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HackSessionContext_getUiViewForEndpoint_Results) SetView(v grain.UiView) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// HackSessionContext_getUiViewForEndpoint_Results_List is a list of HackSessionContext_getUiViewForEndpoint_Results.
type HackSessionContext_getUiViewForEndpoint_Results_List struct{ capnp.List }

// NewHackSessionContext_getUiViewForEndpoint_Results creates a new list of HackSessionContext_getUiViewForEndpoint_Results.
func NewHackSessionContext_getUiViewForEndpoint_Results_List(s *capnp.Segment, sz int32) (HackSessionContext_getUiViewForEndpoint_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HackSessionContext_getUiViewForEndpoint_Results_List{l}, err
}

func (s HackSessionContext_getUiViewForEndpoint_Results_List) At(i int) HackSessionContext_getUiViewForEndpoint_Results {
	return HackSessionContext_getUiViewForEndpoint_Results{s.List.Struct(i)}
}

func (s HackSessionContext_getUiViewForEndpoint_Results_List) Set(i int, v HackSessionContext_getUiViewForEndpoint_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// HackSessionContext_getUiViewForEndpoint_Results_Promise is a wrapper for a HackSessionContext_getUiViewForEndpoint_Results promised by a client call.
type HackSessionContext_getUiViewForEndpoint_Results_Promise struct{ *capnp.Pipeline }

func (p HackSessionContext_getUiViewForEndpoint_Results_Promise) Struct() (HackSessionContext_getUiViewForEndpoint_Results, error) {
	s, err := p.Pipeline.Struct()
	return HackSessionContext_getUiViewForEndpoint_Results{s}, err
}

func (p HackSessionContext_getUiViewForEndpoint_Results_Promise) View() grain.UiView {
	return grain.UiView{Client: p.Pipeline.GetPipeline(0).Client()}
}

type HackEmailSession struct{ Client capnp.Client }

func (c HackEmailSession) Send(ctx context.Context, params func(email.EmailSendPort_send_Params) error, opts ...capnp.CallOption) email.EmailSendPort_send_Results_Promise {
	if c.Client == nil {
		return email.EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailSendPort_send_Params{Struct: s}) }
	}
	return email.EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c HackEmailSession) HintAddress(ctx context.Context, params func(email.EmailSendPort_hintAddress_Params) error, opts ...capnp.CallOption) email.EmailSendPort_hintAddress_Results_Promise {
	if c.Client == nil {
		return email.EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(email.EmailSendPort_hintAddress_Params{Struct: s}) }
	}
	return email.EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type HackEmailSession_Server interface {
	Send(email.EmailSendPort_send) error

	HintAddress(email.EmailSendPort_hintAddress) error
}

func HackEmailSession_ServerToClient(s HackEmailSession_Server) HackEmailSession {
	c, _ := s.(server.Closer)
	return HackEmailSession{Client: server.New(HackEmailSession_Methods(nil, s), c)}
}

func HackEmailSession_Methods(methods []server.Method, s HackEmailSession_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailSendPort_send{c, opts, email.EmailSendPort_send_Params{Struct: p}, email.EmailSendPort_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec831dbf4cc9bcca,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailSendPort",
			MethodName:    "hintAddress",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := email.EmailSendPort_hintAddress{c, opts, email.EmailSendPort_hintAddress_Params{Struct: p}, email.EmailSendPort_hintAddress_Results{Struct: r}}
			return s.HintAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

const schema_bf6889795837d1e0 = "x\xda\xa4Vml\x14U\x17>\xe7\xce\xb6S\x92\xf6" +
	"m\x87\xed\x0b\xbbMK\xa1ix\xa1)\xbc\xb6\xa0(" +
	"1t\x0b\xade\xb1\xd5\x9d\xf2!`@\xa7\xdbK;" +
	"\xe9\xee\xccvgJ\xa1\x91\xb4\x05?R\x8c1\xa8\x08" +
	"\x04\xff\xf0\x03\x0cDA\x10\x8cE\x11\x89\xa9II\xd0" +
	"\x14\x83\x10\x85((b\x95`\x8cH\x80\x04\xc6\xdc\xd9" +
	"\x9d\xe9\xa5PJm\xf6\xcff\xce9\xcf=\x9f\xcf9" +
	"\x0f\xadK\x0b\x90\x92\x94\xaf\xff\x0f\xb0\xb0\x98\xa4\xa4Z" +
	"\x1f}\xd6qF\xb9\xd9\xb6\x01\xe4\xc9\x88\x00)D\x04" +
	"\x98qb|)\x01\xf4^\x1c\xbf\x0f\xd0\xca\xdc\x18m" +
	"\xec\xecX\xf9\"H\x93\x99\x022\x85N_\x01S\xd8" +
	"\xe6+\x03\xb4\xca\xc5\xdb\xf9;O\xdf\xd8\x92P\xf00" +
	"y\xb7o,\x01\x8fU\xbd\xa1oV\xa8b\xe5\x0eN" +
	"\xb2\xdfw\x01\xc1c\x89\xaf~\xf2\xdds\xd3C\xbby" +
	"\xd0\x9d\xbe\x1c\x06z\xc4\x06\xfd\xe5\x9d\xaa\xb3[c\xbb" +
	"\xf6r\xa6\xe7|E\x0c\xf4j\xc5\x82\xae\xac\xeb5\x07" +
	"\x93\xa6\xb6\xc3\xa7|\x07\x909\xecke\xa6O\xab\xa9" +
	"'\x0f<{\x90\xc7\xae\xf4\x8fe\xd8+\xfc\x0c{]" +
	"G\xf1\xb1\xf9m\xd9\xdd\xbc\xc2:\x7f\x11S\xd8l+" +
	"\xec\xf2<9\xf3\xf4W\x87\xbe\x00\xc9+X?\xf6\xcd" +
	"Z\xba\xb6\xab\xf1(\x00\xce\xe8\xf6\x17\xa0\xb7\xd7/\x02" +
	"x{\xfc\xa2\xb7\xc7\xff?\x00\xeb\xc6+i\xbbK\xca" +
	"\xb6\xf4&\xe1\x04\x06w\xc4?\x9b\xc1\x9d\xf2\xb3\x0c\xbe" +
	"\xbe\xfd\xbf\x93\x16\xe6W\x9f\xbf\x0b\xae9\xa7\x14\xbd\x9d" +
	"9\xe3\x01f\xbc\x96SE\xbc%\xb9\x0co\xde\x94\xbf" +
	"?~wV\xc5E\xde\xbd\xa9\xb9{X\x80\x8f\xe52" +
	"\xf7\xf6\xd5\x9f\xe8\x7f\xf8L\xea%.7\xcbr\x8f\xb3" +
	"\xb4\xae\xf8\xfcZA\xcf\x9bJ\xbfSM\xdb\x97\x9a\xdc" +
	"\x93\xccV\xc9\xbd\x04h=\xfe\xc1\xc4\xbd\x19]?\xf7" +
	"\xf3\xe0\xe5yv5\x97\xe51\xf0\x1dY\x87\xeb\xc6Y" +
	"\xbf_\xe6\xc0\xbb\x98\xdcc\xd54|\xf9\xebF\x9au" +
	"\x83I\x06bJ\xbc\xd1\x99\xb7\x1e\xbd\x9b\xf3Xb6" +
	"\xe5\xb1\x98\x7fx\xfb\x8d\xed\xdb*zo\xf1\xefL\x9b" +
	"`\x17\xb8r\x02{\xc7{\xb5\xf9z\xf3\x0b\xad\xb7\xb9" +
	"w(\x93[\xeeo\x92\xd5\xa8\x84\x9b\xa6\x19\xd4\x10\x0d" +
	"U\xd7\xa6\x87\x95\x98\x16\x9b=_\x097-\xa4\x06\xfb" +
	"2O\xd7L\xba\xc6\x9c\xae\xd7\x19z\x84\x9a\xb4\x8aj" +
	"4\xae\x98\xb4<\xa6.\xd2\x9b\xa8V\x18\xcaW\xe2J" +
	"\xd4\x90\xd3\x05\x0f\x80\x07\x01\xa4\xca\xb9\x00r@@\xb9" +
	"\x9a\xa0\x84\x98\xcd\xb2$\x05\x17\x00\xc8\xf3\x05\x94\x17\x11" +
	"D\x92\x8d\x04Q\x92\x99b\xb5\x80\xf2R\x82\xed1j" +
	"jJ\x94b:\x10L\x07\xb4Z\x0c\x1a\x0fj\xabt" +
	"\x00\xc0,\xeb\x1bi\xc5\x87\x7f\xf4u\xbf\x05\x80\x98\x05" +
	"\xd8N\xd7\xc4\xd485p\x0c\x10\x1c\x038\xd2(j" +
	"\xe9j\xbd\x89\x8bA\x89\x8b,\x06\x8f\x1bC\x06s-" +
	"M@9\x9b`\xbb\xc9\x94\x82\xf5\xaek#N\x99\x19" +
	"\x8c=E\xcdV=\xde\xc4\x9eR\xa2h\xb8\x18\xa9\xf7" +
	"\xc7h\xa0\xe6b\x83\xc6\xcb\xeb\xeb\xe3\xd40\x1cGG" +
	"\xe5A-5Z\"\x829d\xb4ZB\x0f%+0" +
	"\xb1\xa8\xad\xe9\x99\xf5\xbbY\xd2\xa5Q'\xb9\x96\x1a\x99" +
	"-\x11s\xc0\xf9\x94\xfb\xc34\x9af\xac\x8a\x9a\x85\xb5" +
	"4\xdf`vr\x9a\xeb\xefT\xd6LS\x04\x94gr" +
	"\x1dV\xc2\x82(\x16P~\x94\xa0\x15U\xa3t\xd1\xda" +
	"\x18e\xdd\x93,[{\x98\x01k&f\x00\xc1\x8c\x07" +
	"\x0f\x87\x95@]\xa2\xd2\xd6'\xf4x\xa5V\x1f\xd3U" +
	"\xcdt\xca\xc8\xe7\xb0` \x87bK<2\xaan\x09" +
	"j&\x8d\xafR\xc2\xd4\xcd\x1a\xffR-\x80\x9c.\xa0" +
	"\xec#h\xa9IM@\x8a\x92u\xf3\xdb9\xfdZ\xa8" +
	"\xf8\xa7\xc1\x15\x13\x06\xbd]\x19U\xd4\x08s@Tu" +
	"-\x84\x18\x12R\xe44DK\x1fwh\xe9\x95\xf2\xd2" +
	"k\x00`\x1d\xff\xb4\xb7\xfah\xde\x86\xcb\xec\xffhI" +
	"\xa2\xb6\x8c&j\xc8\xb1D\xe9\xbdX\xa2n\x80%$" +
	"\xc2h\x02\xe0\x0e\x9a\xc8\xb7g\xd1\xcd-M\xd6\x03\xc4" +
	"\xc5\x03\x19\x1fr^\x85\xa1\xbc\xcf\xb7\xdd\x97=\x88\x1c" +
	"\x01c\xade;\x1f\xd4V\x01\xear\xa1\x90\x02\xe0\xae" +
	"\x05t\xb6\x80W\xc2: \xde1(\"\xba[\x05\x9d" +
	"\xfd)\xdd\x9a\x0bD\xfaSD\xe2\xeei<?g\xc9" +
	"\xcbW\xce\xbe\xff\x9et\xb1\x0d\x88tND\xc1=\x10" +
	"\xd0YsR\xdfq R\x9f\x88\x1e\x97\xc3\xd1\xd9*" +
	"R\xcf\x01 \xd21\x11S\xdc\xb3\x01\x9dM.\x1d:" +
	"\x0cD\xda/\xa2\xe8^\x0c\xe8\x9c\x01\xd2\xce=@\xa4" +
	"\x1d\"\xa6\xb9\xfb\x07\x9d--mfv\x9bDLu" +
	"7;:\xdbEz\x89\xd9u\x8aV\x035C-u" +
	"\x11\x15\xc4p\xb0>\x80\xed\xc9\x01\x0d\xa0\xe5\xb0\x14\x94" +
	"%x*\x80\x96\xd3\x0adp/\x00pR\xacV\x0d" +
	"\x93I\xf2\x99\x88\xb7C\x87=\xca\x12f\xbc\xc8\xa1\xb3" +
	"L\xc6Sw\x09\x82\x9aY\x96\x18\x9e\xa4glx\xd1" +
	"\x99\xdeL\xd6.\x01\xb4\xbb\xfd\x91\xbf\xc2\xcb\xf3\x02s" +
	"\x8e\x0e\xd5\xed\x0fHP!%3>h{\x0c\xc7\x05" +
	"\xc3\xb3\xbe\x9d\xe9p\xb0>\xc14\x06\xe7\xd5\x08lm" +
	"\xaa7\xd1\x90\xb3\\\xd7\x14F\x9d\xcf\x0b(G\xb8\xb1" +
	"S\xd9\xc7F\x01e\x93\x1b\xbbf6v\x11\x01\xe55" +
	"\x04Q\xc8F\x01@jY\x0e \x9b\x02\xca\x1d\x04\xad" +
	"X\xf2\x19\x8ec\xadF\xdd\xb0\xd78\xcf\xbbJ\x8b\xa9" +
	"s\xe3i\xa9F\x05\x8d\xea\x8b\x0d\x10h\x1c\x11\x08\xe2" +
	"\xc8Y\xd2\xe9\x1b\xbbm\xec8\xc5A$9;Y\x82" +
	"B\x82e6!\x18\xf8\x1f\xc0\x90\x80\x9850\xe4\x80" +
	"\xec\xe3()z\xf0R\x1e\xa6m\x1cR\xd1!\x848" +
	"\xec\xd9\xc4>V\x08(\x87\xb8\xca\xd4,\xe0\xee\xa6A" +
	"d7\xb2;jT;\xf0^\x87D\xd1@\xe3g\xae" +
	"Vi+J\xd6\xa5\x0b\x0d\xbfm=}\xf0\xfb\x7fy" +
	"E\xdcYi6\x0dB\xd4\xf8'\x00\x00\xff\xff\x81N" +
	"\x0f|"

func init() {
	schemas.Register(schema_bf6889795837d1e0,
		0x837afa61d880beb6,
		0x845e8081686d8a0f,
		0x97f9d7a51ffe0741,
		0xa15e445037d1834c,
		0xa9502e5fdabb8b07,
		0xaea67098dc479ce6,
		0xb44df810894a44f4,
		0xb45bb2d206694fe6,
		0xb9147a48c12c807d,
		0xc3b5ced7344b04a6,
		0xc9973f31a90887f9,
		0xe14c1f5321159b8f,
		0xe54437a7b8f52843,
		0xe706d835e9cd64af,
		0xe96193c522f6c05d,
		0xe9e4890dae20b03c,
		0xecebff1662ba10a1,
		0xf910658ae8c6674d,
		0xfdc944999b9296df,
		0xfe777c71f871f413)
}
