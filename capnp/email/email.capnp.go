package email

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	util "zenhack.net/go/sandstorm/capnp/util"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type EmailAddress struct{ capnp.Struct }

// EmailAddress_TypeID is the unique identifier for the type EmailAddress.
const EmailAddress_TypeID = 0xacaddcee86563ee1

func NewEmailAddress(s *capnp.Segment) (EmailAddress, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailAddress{st}, err
}

func NewRootEmailAddress(s *capnp.Segment) (EmailAddress, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailAddress{st}, err
}

func ReadRootEmailAddress(msg *capnp.Message) (EmailAddress, error) {
	root, err := msg.RootPtr()
	return EmailAddress{root.Struct()}, err
}

func (s EmailAddress) String() string {
	str, _ := text.Marshal(0xacaddcee86563ee1, s.Struct)
	return str
}

func (s EmailAddress) Address() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s EmailAddress) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAddress) AddressBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s EmailAddress) SetAddress(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s EmailAddress) Name() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s EmailAddress) HasName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailAddress) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s EmailAddress) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// EmailAddress_List is a list of EmailAddress.
type EmailAddress_List struct{ capnp.List }

// NewEmailAddress creates a new list of EmailAddress.
func NewEmailAddress_List(s *capnp.Segment, sz int32) (EmailAddress_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return EmailAddress_List{l}, err
}

func (s EmailAddress_List) At(i int) EmailAddress { return EmailAddress{s.List.Struct(i)} }

func (s EmailAddress_List) Set(i int, v EmailAddress) error { return s.List.SetStruct(i, v.Struct) }

// EmailAddress_Promise is a wrapper for a EmailAddress promised by a client call.
type EmailAddress_Promise struct{ *capnp.Pipeline }

func (p EmailAddress_Promise) Struct() (EmailAddress, error) {
	s, err := p.Pipeline.Struct()
	return EmailAddress{s}, err
}

type EmailAttachment struct{ capnp.Struct }

// EmailAttachment_TypeID is the unique identifier for the type EmailAttachment.
const EmailAttachment_TypeID = 0xb309c51a9d28244f

func NewEmailAttachment(s *capnp.Segment) (EmailAttachment, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return EmailAttachment{st}, err
}

func NewRootEmailAttachment(s *capnp.Segment) (EmailAttachment, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return EmailAttachment{st}, err
}

func ReadRootEmailAttachment(msg *capnp.Message) (EmailAttachment, error) {
	root, err := msg.RootPtr()
	return EmailAttachment{root.Struct()}, err
}

func (s EmailAttachment) String() string {
	str, _ := text.Marshal(0xb309c51a9d28244f, s.Struct)
	return str
}

func (s EmailAttachment) ContentType() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s EmailAttachment) HasContentType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentTypeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s EmailAttachment) SetContentType(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s EmailAttachment) ContentDisposition() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s EmailAttachment) HasContentDisposition() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentDispositionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s EmailAttachment) SetContentDisposition(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s EmailAttachment) ContentId() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s EmailAttachment) HasContentId() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) ContentIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s EmailAttachment) SetContentId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s EmailAttachment) Content() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s EmailAttachment) HasContent() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s EmailAttachment) SetContent(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, d.List.ToPtr())
}

// EmailAttachment_List is a list of EmailAttachment.
type EmailAttachment_List struct{ capnp.List }

// NewEmailAttachment creates a new list of EmailAttachment.
func NewEmailAttachment_List(s *capnp.Segment, sz int32) (EmailAttachment_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return EmailAttachment_List{l}, err
}

func (s EmailAttachment_List) At(i int) EmailAttachment { return EmailAttachment{s.List.Struct(i)} }

func (s EmailAttachment_List) Set(i int, v EmailAttachment) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAttachment_Promise is a wrapper for a EmailAttachment promised by a client call.
type EmailAttachment_Promise struct{ *capnp.Pipeline }

func (p EmailAttachment_Promise) Struct() (EmailAttachment, error) {
	s, err := p.Pipeline.Struct()
	return EmailAttachment{s}, err
}

type EmailMessage struct{ capnp.Struct }

// EmailMessage_TypeID is the unique identifier for the type EmailMessage.
const EmailMessage_TypeID = 0xcff459e769562d2f

func NewEmailMessage(s *capnp.Segment) (EmailMessage, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 12})
	return EmailMessage{st}, err
}

func NewRootEmailMessage(s *capnp.Segment) (EmailMessage, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 12})
	return EmailMessage{st}, err
}

func ReadRootEmailMessage(msg *capnp.Message) (EmailMessage, error) {
	root, err := msg.RootPtr()
	return EmailMessage{root.Struct()}, err
}

func (s EmailMessage) String() string {
	str, _ := text.Marshal(0xcff459e769562d2f, s.Struct)
	return str
}

func (s EmailMessage) Date() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s EmailMessage) SetDate(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s EmailMessage) From() (EmailAddress, error) {
	p, err := s.Struct.Ptr(0)
	return EmailAddress{Struct: p.Struct()}, err
}

func (s EmailMessage) HasFrom() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetFrom(v EmailAddress) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewFrom sets the from field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailMessage) NewFrom() (EmailAddress, error) {
	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s EmailMessage) To() (EmailAddress_List, error) {
	p, err := s.Struct.Ptr(1)
	return EmailAddress_List{List: p.List()}, err
}

func (s EmailMessage) HasTo() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetTo(v EmailAddress_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewTo sets the to field to a newly
// allocated EmailAddress_List, preferring placement in s's segment.
func (s EmailMessage) NewTo(n int32) (EmailAddress_List, error) {
	l, err := NewEmailAddress_List(s.Struct.Segment(), n)
	if err != nil {
		return EmailAddress_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) Cc() (EmailAddress_List, error) {
	p, err := s.Struct.Ptr(2)
	return EmailAddress_List{List: p.List()}, err
}

func (s EmailMessage) HasCc() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetCc(v EmailAddress_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewCc sets the cc field to a newly
// allocated EmailAddress_List, preferring placement in s's segment.
func (s EmailMessage) NewCc(n int32) (EmailAddress_List, error) {
	l, err := NewEmailAddress_List(s.Struct.Segment(), n)
	if err != nil {
		return EmailAddress_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) Bcc() (EmailAddress_List, error) {
	p, err := s.Struct.Ptr(3)
	return EmailAddress_List{List: p.List()}, err
}

func (s EmailMessage) HasBcc() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetBcc(v EmailAddress_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewBcc sets the bcc field to a newly
// allocated EmailAddress_List, preferring placement in s's segment.
func (s EmailMessage) NewBcc(n int32) (EmailAddress_List, error) {
	l, err := NewEmailAddress_List(s.Struct.Segment(), n)
	if err != nil {
		return EmailAddress_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) ReplyTo() (EmailAddress, error) {
	p, err := s.Struct.Ptr(4)
	return EmailAddress{Struct: p.Struct()}, err
}

func (s EmailMessage) HasReplyTo() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetReplyTo(v EmailAddress) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewReplyTo sets the replyTo field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailMessage) NewReplyTo() (EmailAddress, error) {
	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s EmailMessage) MessageId() (string, error) {
	p, err := s.Struct.Ptr(5)
	return p.Text(), err
}

func (s EmailMessage) HasMessageId() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s EmailMessage) MessageIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return p.TextBytes(), err
}

func (s EmailMessage) SetMessageId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(5, t.List.ToPtr())
}

func (s EmailMessage) References() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(6)
	return capnp.TextList{List: p.List()}, err
}

func (s EmailMessage) HasReferences() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetReferences(v capnp.TextList) error {
	return s.Struct.SetPtr(6, v.List.ToPtr())
}

// NewReferences sets the references field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s EmailMessage) NewReferences(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(6, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) InReplyTo() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(7)
	return capnp.TextList{List: p.List()}, err
}

func (s EmailMessage) HasInReplyTo() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetInReplyTo(v capnp.TextList) error {
	return s.Struct.SetPtr(7, v.List.ToPtr())
}

// NewInReplyTo sets the inReplyTo field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s EmailMessage) NewInReplyTo(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(7, l.List.ToPtr())
	return l, err
}

func (s EmailMessage) Subject() (string, error) {
	p, err := s.Struct.Ptr(8)
	return p.Text(), err
}

func (s EmailMessage) HasSubject() bool {
	p, err := s.Struct.Ptr(8)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SubjectBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(8)
	return p.TextBytes(), err
}

func (s EmailMessage) SetSubject(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(8, t.List.ToPtr())
}

func (s EmailMessage) Text() (string, error) {
	p, err := s.Struct.Ptr(9)
	return p.Text(), err
}

func (s EmailMessage) HasText() bool {
	p, err := s.Struct.Ptr(9)
	return p.IsValid() || err != nil
}

func (s EmailMessage) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(9)
	return p.TextBytes(), err
}

func (s EmailMessage) SetText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(9, t.List.ToPtr())
}

func (s EmailMessage) Html() (string, error) {
	p, err := s.Struct.Ptr(10)
	return p.Text(), err
}

func (s EmailMessage) HasHtml() bool {
	p, err := s.Struct.Ptr(10)
	return p.IsValid() || err != nil
}

func (s EmailMessage) HtmlBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(10)
	return p.TextBytes(), err
}

func (s EmailMessage) SetHtml(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(10, t.List.ToPtr())
}

func (s EmailMessage) Attachments() (EmailAttachment_List, error) {
	p, err := s.Struct.Ptr(11)
	return EmailAttachment_List{List: p.List()}, err
}

func (s EmailMessage) HasAttachments() bool {
	p, err := s.Struct.Ptr(11)
	return p.IsValid() || err != nil
}

func (s EmailMessage) SetAttachments(v EmailAttachment_List) error {
	return s.Struct.SetPtr(11, v.List.ToPtr())
}

// NewAttachments sets the attachments field to a newly
// allocated EmailAttachment_List, preferring placement in s's segment.
func (s EmailMessage) NewAttachments(n int32) (EmailAttachment_List, error) {
	l, err := NewEmailAttachment_List(s.Struct.Segment(), n)
	if err != nil {
		return EmailAttachment_List{}, err
	}
	err = s.Struct.SetPtr(11, l.List.ToPtr())
	return l, err
}

// EmailMessage_List is a list of EmailMessage.
type EmailMessage_List struct{ capnp.List }

// NewEmailMessage creates a new list of EmailMessage.
func NewEmailMessage_List(s *capnp.Segment, sz int32) (EmailMessage_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 12}, sz)
	return EmailMessage_List{l}, err
}

func (s EmailMessage_List) At(i int) EmailMessage { return EmailMessage{s.List.Struct(i)} }

func (s EmailMessage_List) Set(i int, v EmailMessage) error { return s.List.SetStruct(i, v.Struct) }

// EmailMessage_Promise is a wrapper for a EmailMessage promised by a client call.
type EmailMessage_Promise struct{ *capnp.Pipeline }

func (p EmailMessage_Promise) Struct() (EmailMessage, error) {
	s, err := p.Pipeline.Struct()
	return EmailMessage{s}, err
}

func (p EmailMessage_Promise) From() EmailAddress_Promise {
	return EmailAddress_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p EmailMessage_Promise) ReplyTo() EmailAddress_Promise {
	return EmailAddress_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

type EmailSendPort struct{ Client capnp.Client }

func (c EmailSendPort) Send(ctx context.Context, params func(EmailSendPort_send_Params) error, opts ...capnp.CallOption) EmailSendPort_send_Results_Promise {
	if c.Client == nil {
		return EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailSendPort_send_Params{Struct: s}) }
	}
	return EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c EmailSendPort) HintAddress(ctx context.Context, params func(EmailSendPort_hintAddress_Params) error, opts ...capnp.CallOption) EmailSendPort_hintAddress_Results_Promise {
	if c.Client == nil {
		return EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailSendPort_hintAddress_Params{Struct: s}) }
	}
	return EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type EmailSendPort_Server interface {
	Send(EmailSendPort_send) error

	HintAddress(EmailSendPort_hintAddress) error
}

func EmailSendPort_ServerToClient(s EmailSendPort_Server) EmailSendPort {
	c, _ := s.(server.Closer)
	return EmailSendPort{Client: server.New(EmailSendPort_Methods(nil, s), c)}
}

func EmailSendPort_Methods(methods []server.Method, s EmailSendPort_Server) []server.Method {
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
			call := EmailSendPort_send{c, opts, EmailSendPort_send_Params{Struct: p}, EmailSendPort_send_Results{Struct: r}}
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
			call := EmailSendPort_hintAddress{c, opts, EmailSendPort_hintAddress_Params{Struct: p}, EmailSendPort_hintAddress_Results{Struct: r}}
			return s.HintAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// EmailSendPort_send holds the arguments for a server call to EmailSendPort.send.
type EmailSendPort_send struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailSendPort_send_Params
	Results EmailSendPort_send_Results
}

// EmailSendPort_hintAddress holds the arguments for a server call to EmailSendPort.hintAddress.
type EmailSendPort_hintAddress struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailSendPort_hintAddress_Params
	Results EmailSendPort_hintAddress_Results
}

type EmailSendPort_PowerboxTag struct{ capnp.Struct }

// EmailSendPort_PowerboxTag_TypeID is the unique identifier for the type EmailSendPort_PowerboxTag.
const EmailSendPort_PowerboxTag_TypeID = 0x90790c61fc899dd3

func NewEmailSendPort_PowerboxTag(s *capnp.Segment) (EmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailSendPort_PowerboxTag{st}, err
}

func NewRootEmailSendPort_PowerboxTag(s *capnp.Segment) (EmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailSendPort_PowerboxTag{st}, err
}

func ReadRootEmailSendPort_PowerboxTag(msg *capnp.Message) (EmailSendPort_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_PowerboxTag{root.Struct()}, err
}

func (s EmailSendPort_PowerboxTag) String() string {
	str, _ := text.Marshal(0x90790c61fc899dd3, s.Struct)
	return str
}

func (s EmailSendPort_PowerboxTag) FromHint() (EmailAddress, error) {
	p, err := s.Struct.Ptr(0)
	return EmailAddress{Struct: p.Struct()}, err
}

func (s EmailSendPort_PowerboxTag) HasFromHint() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailSendPort_PowerboxTag) SetFromHint(v EmailAddress) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewFromHint sets the fromHint field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailSendPort_PowerboxTag) NewFromHint() (EmailAddress, error) {
	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s EmailSendPort_PowerboxTag) ListIdHint() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s EmailSendPort_PowerboxTag) HasListIdHint() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailSendPort_PowerboxTag) ListIdHintBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s EmailSendPort_PowerboxTag) SetListIdHint(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

// EmailSendPort_PowerboxTag_List is a list of EmailSendPort_PowerboxTag.
type EmailSendPort_PowerboxTag_List struct{ capnp.List }

// NewEmailSendPort_PowerboxTag creates a new list of EmailSendPort_PowerboxTag.
func NewEmailSendPort_PowerboxTag_List(s *capnp.Segment, sz int32) (EmailSendPort_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return EmailSendPort_PowerboxTag_List{l}, err
}

func (s EmailSendPort_PowerboxTag_List) At(i int) EmailSendPort_PowerboxTag {
	return EmailSendPort_PowerboxTag{s.List.Struct(i)}
}

func (s EmailSendPort_PowerboxTag_List) Set(i int, v EmailSendPort_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_PowerboxTag_Promise is a wrapper for a EmailSendPort_PowerboxTag promised by a client call.
type EmailSendPort_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_PowerboxTag_Promise) Struct() (EmailSendPort_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_PowerboxTag{s}, err
}

func (p EmailSendPort_PowerboxTag_Promise) FromHint() EmailAddress_Promise {
	return EmailAddress_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EmailSendPort_send_Params struct{ capnp.Struct }

// EmailSendPort_send_Params_TypeID is the unique identifier for the type EmailSendPort_send_Params.
const EmailSendPort_send_Params_TypeID = 0xa5adb72b4ccc59ee

func NewEmailSendPort_send_Params(s *capnp.Segment) (EmailSendPort_send_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailSendPort_send_Params{st}, err
}

func NewRootEmailSendPort_send_Params(s *capnp.Segment) (EmailSendPort_send_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailSendPort_send_Params{st}, err
}

func ReadRootEmailSendPort_send_Params(msg *capnp.Message) (EmailSendPort_send_Params, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_send_Params{root.Struct()}, err
}

func (s EmailSendPort_send_Params) String() string {
	str, _ := text.Marshal(0xa5adb72b4ccc59ee, s.Struct)
	return str
}

func (s EmailSendPort_send_Params) Email() (EmailMessage, error) {
	p, err := s.Struct.Ptr(0)
	return EmailMessage{Struct: p.Struct()}, err
}

func (s EmailSendPort_send_Params) HasEmail() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailSendPort_send_Params) SetEmail(v EmailMessage) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewEmail sets the email field to a newly
// allocated EmailMessage struct, preferring placement in s's segment.
func (s EmailSendPort_send_Params) NewEmail() (EmailMessage, error) {
	ss, err := NewEmailMessage(s.Struct.Segment())
	if err != nil {
		return EmailMessage{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// EmailSendPort_send_Params_List is a list of EmailSendPort_send_Params.
type EmailSendPort_send_Params_List struct{ capnp.List }

// NewEmailSendPort_send_Params creates a new list of EmailSendPort_send_Params.
func NewEmailSendPort_send_Params_List(s *capnp.Segment, sz int32) (EmailSendPort_send_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailSendPort_send_Params_List{l}, err
}

func (s EmailSendPort_send_Params_List) At(i int) EmailSendPort_send_Params {
	return EmailSendPort_send_Params{s.List.Struct(i)}
}

func (s EmailSendPort_send_Params_List) Set(i int, v EmailSendPort_send_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_send_Params_Promise is a wrapper for a EmailSendPort_send_Params promised by a client call.
type EmailSendPort_send_Params_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_send_Params_Promise) Struct() (EmailSendPort_send_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_send_Params{s}, err
}

func (p EmailSendPort_send_Params_Promise) Email() EmailMessage_Promise {
	return EmailMessage_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EmailSendPort_send_Results struct{ capnp.Struct }

// EmailSendPort_send_Results_TypeID is the unique identifier for the type EmailSendPort_send_Results.
const EmailSendPort_send_Results_TypeID = 0xd063b4e6c91bf8d8

func NewEmailSendPort_send_Results(s *capnp.Segment) (EmailSendPort_send_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailSendPort_send_Results{st}, err
}

func NewRootEmailSendPort_send_Results(s *capnp.Segment) (EmailSendPort_send_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailSendPort_send_Results{st}, err
}

func ReadRootEmailSendPort_send_Results(msg *capnp.Message) (EmailSendPort_send_Results, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_send_Results{root.Struct()}, err
}

func (s EmailSendPort_send_Results) String() string {
	str, _ := text.Marshal(0xd063b4e6c91bf8d8, s.Struct)
	return str
}

// EmailSendPort_send_Results_List is a list of EmailSendPort_send_Results.
type EmailSendPort_send_Results_List struct{ capnp.List }

// NewEmailSendPort_send_Results creates a new list of EmailSendPort_send_Results.
func NewEmailSendPort_send_Results_List(s *capnp.Segment, sz int32) (EmailSendPort_send_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return EmailSendPort_send_Results_List{l}, err
}

func (s EmailSendPort_send_Results_List) At(i int) EmailSendPort_send_Results {
	return EmailSendPort_send_Results{s.List.Struct(i)}
}

func (s EmailSendPort_send_Results_List) Set(i int, v EmailSendPort_send_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_send_Results_Promise is a wrapper for a EmailSendPort_send_Results promised by a client call.
type EmailSendPort_send_Results_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_send_Results_Promise) Struct() (EmailSendPort_send_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_send_Results{s}, err
}

type EmailSendPort_hintAddress_Params struct{ capnp.Struct }

// EmailSendPort_hintAddress_Params_TypeID is the unique identifier for the type EmailSendPort_hintAddress_Params.
const EmailSendPort_hintAddress_Params_TypeID = 0x9c78c3c5de56e4d4

func NewEmailSendPort_hintAddress_Params(s *capnp.Segment) (EmailSendPort_hintAddress_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailSendPort_hintAddress_Params{st}, err
}

func NewRootEmailSendPort_hintAddress_Params(s *capnp.Segment) (EmailSendPort_hintAddress_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailSendPort_hintAddress_Params{st}, err
}

func ReadRootEmailSendPort_hintAddress_Params(msg *capnp.Message) (EmailSendPort_hintAddress_Params, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_hintAddress_Params{root.Struct()}, err
}

func (s EmailSendPort_hintAddress_Params) String() string {
	str, _ := text.Marshal(0x9c78c3c5de56e4d4, s.Struct)
	return str
}

func (s EmailSendPort_hintAddress_Params) Address() (EmailAddress, error) {
	p, err := s.Struct.Ptr(0)
	return EmailAddress{Struct: p.Struct()}, err
}

func (s EmailSendPort_hintAddress_Params) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailSendPort_hintAddress_Params) SetAddress(v EmailAddress) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAddress sets the address field to a newly
// allocated EmailAddress struct, preferring placement in s's segment.
func (s EmailSendPort_hintAddress_Params) NewAddress() (EmailAddress, error) {
	ss, err := NewEmailAddress(s.Struct.Segment())
	if err != nil {
		return EmailAddress{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// EmailSendPort_hintAddress_Params_List is a list of EmailSendPort_hintAddress_Params.
type EmailSendPort_hintAddress_Params_List struct{ capnp.List }

// NewEmailSendPort_hintAddress_Params creates a new list of EmailSendPort_hintAddress_Params.
func NewEmailSendPort_hintAddress_Params_List(s *capnp.Segment, sz int32) (EmailSendPort_hintAddress_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailSendPort_hintAddress_Params_List{l}, err
}

func (s EmailSendPort_hintAddress_Params_List) At(i int) EmailSendPort_hintAddress_Params {
	return EmailSendPort_hintAddress_Params{s.List.Struct(i)}
}

func (s EmailSendPort_hintAddress_Params_List) Set(i int, v EmailSendPort_hintAddress_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_hintAddress_Params_Promise is a wrapper for a EmailSendPort_hintAddress_Params promised by a client call.
type EmailSendPort_hintAddress_Params_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_hintAddress_Params_Promise) Struct() (EmailSendPort_hintAddress_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_hintAddress_Params{s}, err
}

func (p EmailSendPort_hintAddress_Params_Promise) Address() EmailAddress_Promise {
	return EmailAddress_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EmailSendPort_hintAddress_Results struct{ capnp.Struct }

// EmailSendPort_hintAddress_Results_TypeID is the unique identifier for the type EmailSendPort_hintAddress_Results.
const EmailSendPort_hintAddress_Results_TypeID = 0xbd727a009329aabc

func NewEmailSendPort_hintAddress_Results(s *capnp.Segment) (EmailSendPort_hintAddress_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailSendPort_hintAddress_Results{st}, err
}

func NewRootEmailSendPort_hintAddress_Results(s *capnp.Segment) (EmailSendPort_hintAddress_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailSendPort_hintAddress_Results{st}, err
}

func ReadRootEmailSendPort_hintAddress_Results(msg *capnp.Message) (EmailSendPort_hintAddress_Results, error) {
	root, err := msg.RootPtr()
	return EmailSendPort_hintAddress_Results{root.Struct()}, err
}

func (s EmailSendPort_hintAddress_Results) String() string {
	str, _ := text.Marshal(0xbd727a009329aabc, s.Struct)
	return str
}

// EmailSendPort_hintAddress_Results_List is a list of EmailSendPort_hintAddress_Results.
type EmailSendPort_hintAddress_Results_List struct{ capnp.List }

// NewEmailSendPort_hintAddress_Results creates a new list of EmailSendPort_hintAddress_Results.
func NewEmailSendPort_hintAddress_Results_List(s *capnp.Segment, sz int32) (EmailSendPort_hintAddress_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return EmailSendPort_hintAddress_Results_List{l}, err
}

func (s EmailSendPort_hintAddress_Results_List) At(i int) EmailSendPort_hintAddress_Results {
	return EmailSendPort_hintAddress_Results{s.List.Struct(i)}
}

func (s EmailSendPort_hintAddress_Results_List) Set(i int, v EmailSendPort_hintAddress_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailSendPort_hintAddress_Results_Promise is a wrapper for a EmailSendPort_hintAddress_Results promised by a client call.
type EmailSendPort_hintAddress_Results_Promise struct{ *capnp.Pipeline }

func (p EmailSendPort_hintAddress_Results_Promise) Struct() (EmailSendPort_hintAddress_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailSendPort_hintAddress_Results{s}, err
}

type VerifiedEmail struct{ Client capnp.Client }

type VerifiedEmail_Server interface {
}

func VerifiedEmail_ServerToClient(s VerifiedEmail_Server) VerifiedEmail {
	c, _ := s.(server.Closer)
	return VerifiedEmail{Client: server.New(VerifiedEmail_Methods(nil, s), c)}
}

func VerifiedEmail_Methods(methods []server.Method, s VerifiedEmail_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 0)
	}

	return methods
}

type VerifiedEmail_PowerboxTag struct{ capnp.Struct }

// VerifiedEmail_PowerboxTag_TypeID is the unique identifier for the type VerifiedEmail_PowerboxTag.
const VerifiedEmail_PowerboxTag_TypeID = 0x97469291ac5bb892

func NewVerifiedEmail_PowerboxTag(s *capnp.Segment) (VerifiedEmail_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return VerifiedEmail_PowerboxTag{st}, err
}

func NewRootVerifiedEmail_PowerboxTag(s *capnp.Segment) (VerifiedEmail_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return VerifiedEmail_PowerboxTag{st}, err
}

func ReadRootVerifiedEmail_PowerboxTag(msg *capnp.Message) (VerifiedEmail_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return VerifiedEmail_PowerboxTag{root.Struct()}, err
}

func (s VerifiedEmail_PowerboxTag) String() string {
	str, _ := text.Marshal(0x97469291ac5bb892, s.Struct)
	return str
}

func (s VerifiedEmail_PowerboxTag) VerifierId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s VerifiedEmail_PowerboxTag) HasVerifierId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerifiedEmail_PowerboxTag) SetVerifierId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s VerifiedEmail_PowerboxTag) Address() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s VerifiedEmail_PowerboxTag) HasAddress() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerifiedEmail_PowerboxTag) AddressBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s VerifiedEmail_PowerboxTag) SetAddress(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s VerifiedEmail_PowerboxTag) Domain() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s VerifiedEmail_PowerboxTag) HasDomain() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s VerifiedEmail_PowerboxTag) DomainBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s VerifiedEmail_PowerboxTag) SetDomain(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

// VerifiedEmail_PowerboxTag_List is a list of VerifiedEmail_PowerboxTag.
type VerifiedEmail_PowerboxTag_List struct{ capnp.List }

// NewVerifiedEmail_PowerboxTag creates a new list of VerifiedEmail_PowerboxTag.
func NewVerifiedEmail_PowerboxTag_List(s *capnp.Segment, sz int32) (VerifiedEmail_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return VerifiedEmail_PowerboxTag_List{l}, err
}

func (s VerifiedEmail_PowerboxTag_List) At(i int) VerifiedEmail_PowerboxTag {
	return VerifiedEmail_PowerboxTag{s.List.Struct(i)}
}

func (s VerifiedEmail_PowerboxTag_List) Set(i int, v VerifiedEmail_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// VerifiedEmail_PowerboxTag_Promise is a wrapper for a VerifiedEmail_PowerboxTag promised by a client call.
type VerifiedEmail_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p VerifiedEmail_PowerboxTag_Promise) Struct() (VerifiedEmail_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return VerifiedEmail_PowerboxTag{s}, err
}

type VerifiedEmailSendPort struct{ Client capnp.Client }

func (c VerifiedEmailSendPort) Send(ctx context.Context, params func(EmailSendPort_send_Params) error, opts ...capnp.CallOption) EmailSendPort_send_Results_Promise {
	if c.Client == nil {
		return EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailSendPort_send_Params{Struct: s}) }
	}
	return EmailSendPort_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c VerifiedEmailSendPort) HintAddress(ctx context.Context, params func(EmailSendPort_hintAddress_Params) error, opts ...capnp.CallOption) EmailSendPort_hintAddress_Results_Promise {
	if c.Client == nil {
		return EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
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
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailSendPort_hintAddress_Params{Struct: s}) }
	}
	return EmailSendPort_hintAddress_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type VerifiedEmailSendPort_Server interface {
	Send(EmailSendPort_send) error

	HintAddress(EmailSendPort_hintAddress) error
}

func VerifiedEmailSendPort_ServerToClient(s VerifiedEmailSendPort_Server) VerifiedEmailSendPort {
	c, _ := s.(server.Closer)
	return VerifiedEmailSendPort{Client: server.New(VerifiedEmailSendPort_Methods(nil, s), c)}
}

func VerifiedEmailSendPort_Methods(methods []server.Method, s VerifiedEmailSendPort_Server) []server.Method {
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
			call := EmailSendPort_send{c, opts, EmailSendPort_send_Params{Struct: p}, EmailSendPort_send_Results{Struct: r}}
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
			call := EmailSendPort_hintAddress{c, opts, EmailSendPort_hintAddress_Params{Struct: p}, EmailSendPort_hintAddress_Results{Struct: r}}
			return s.HintAddress(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

type VerifiedEmailSendPort_PowerboxTag struct{ capnp.Struct }

// VerifiedEmailSendPort_PowerboxTag_TypeID is the unique identifier for the type VerifiedEmailSendPort_PowerboxTag.
const VerifiedEmailSendPort_PowerboxTag_TypeID = 0x8f555bd4141fbb3b

func NewVerifiedEmailSendPort_PowerboxTag(s *capnp.Segment) (VerifiedEmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return VerifiedEmailSendPort_PowerboxTag{st}, err
}

func NewRootVerifiedEmailSendPort_PowerboxTag(s *capnp.Segment) (VerifiedEmailSendPort_PowerboxTag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return VerifiedEmailSendPort_PowerboxTag{st}, err
}

func ReadRootVerifiedEmailSendPort_PowerboxTag(msg *capnp.Message) (VerifiedEmailSendPort_PowerboxTag, error) {
	root, err := msg.RootPtr()
	return VerifiedEmailSendPort_PowerboxTag{root.Struct()}, err
}

func (s VerifiedEmailSendPort_PowerboxTag) String() string {
	str, _ := text.Marshal(0x8f555bd4141fbb3b, s.Struct)
	return str
}

func (s VerifiedEmailSendPort_PowerboxTag) Verification() (VerifiedEmail_PowerboxTag, error) {
	p, err := s.Struct.Ptr(0)
	return VerifiedEmail_PowerboxTag{Struct: p.Struct()}, err
}

func (s VerifiedEmailSendPort_PowerboxTag) HasVerification() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerifiedEmailSendPort_PowerboxTag) SetVerification(v VerifiedEmail_PowerboxTag) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewVerification sets the verification field to a newly
// allocated VerifiedEmail_PowerboxTag struct, preferring placement in s's segment.
func (s VerifiedEmailSendPort_PowerboxTag) NewVerification() (VerifiedEmail_PowerboxTag, error) {
	ss, err := NewVerifiedEmail_PowerboxTag(s.Struct.Segment())
	if err != nil {
		return VerifiedEmail_PowerboxTag{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerifiedEmailSendPort_PowerboxTag) Port() (EmailSendPort_PowerboxTag, error) {
	p, err := s.Struct.Ptr(1)
	return EmailSendPort_PowerboxTag{Struct: p.Struct()}, err
}

func (s VerifiedEmailSendPort_PowerboxTag) HasPort() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerifiedEmailSendPort_PowerboxTag) SetPort(v EmailSendPort_PowerboxTag) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPort sets the port field to a newly
// allocated EmailSendPort_PowerboxTag struct, preferring placement in s's segment.
func (s VerifiedEmailSendPort_PowerboxTag) NewPort() (EmailSendPort_PowerboxTag, error) {
	ss, err := NewEmailSendPort_PowerboxTag(s.Struct.Segment())
	if err != nil {
		return EmailSendPort_PowerboxTag{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// VerifiedEmailSendPort_PowerboxTag_List is a list of VerifiedEmailSendPort_PowerboxTag.
type VerifiedEmailSendPort_PowerboxTag_List struct{ capnp.List }

// NewVerifiedEmailSendPort_PowerboxTag creates a new list of VerifiedEmailSendPort_PowerboxTag.
func NewVerifiedEmailSendPort_PowerboxTag_List(s *capnp.Segment, sz int32) (VerifiedEmailSendPort_PowerboxTag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return VerifiedEmailSendPort_PowerboxTag_List{l}, err
}

func (s VerifiedEmailSendPort_PowerboxTag_List) At(i int) VerifiedEmailSendPort_PowerboxTag {
	return VerifiedEmailSendPort_PowerboxTag{s.List.Struct(i)}
}

func (s VerifiedEmailSendPort_PowerboxTag_List) Set(i int, v VerifiedEmailSendPort_PowerboxTag) error {
	return s.List.SetStruct(i, v.Struct)
}

// VerifiedEmailSendPort_PowerboxTag_Promise is a wrapper for a VerifiedEmailSendPort_PowerboxTag promised by a client call.
type VerifiedEmailSendPort_PowerboxTag_Promise struct{ *capnp.Pipeline }

func (p VerifiedEmailSendPort_PowerboxTag_Promise) Struct() (VerifiedEmailSendPort_PowerboxTag, error) {
	s, err := p.Pipeline.Struct()
	return VerifiedEmailSendPort_PowerboxTag{s}, err
}

func (p VerifiedEmailSendPort_PowerboxTag_Promise) Verification() VerifiedEmail_PowerboxTag_Promise {
	return VerifiedEmail_PowerboxTag_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerifiedEmailSendPort_PowerboxTag_Promise) Port() EmailSendPort_PowerboxTag_Promise {
	return EmailSendPort_PowerboxTag_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type EmailVerifier struct{ Client capnp.Client }

func (c EmailVerifier) GetId(ctx context.Context, params func(EmailVerifier_getId_Params) error, opts ...capnp.CallOption) EmailVerifier_getId_Results_Promise {
	if c.Client == nil {
		return EmailVerifier_getId_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "getId",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailVerifier_getId_Params{Struct: s}) }
	}
	return EmailVerifier_getId_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c EmailVerifier) VerifyEmail(ctx context.Context, params func(EmailVerifier_verifyEmail_Params) error, opts ...capnp.CallOption) EmailVerifier_verifyEmail_Results_Promise {
	if c.Client == nil {
		return EmailVerifier_verifyEmail_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "verifyEmail",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailVerifier_verifyEmail_Params{Struct: s}) }
	}
	return EmailVerifier_verifyEmail_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type EmailVerifier_Server interface {
	GetId(EmailVerifier_getId) error

	VerifyEmail(EmailVerifier_verifyEmail) error
}

func EmailVerifier_ServerToClient(s EmailVerifier_Server) EmailVerifier {
	c, _ := s.(server.Closer)
	return EmailVerifier{Client: server.New(EmailVerifier_Methods(nil, s), c)}
}

func EmailVerifier_Methods(methods []server.Method, s EmailVerifier_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "getId",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailVerifier_getId{c, opts, EmailVerifier_getId_Params{Struct: p}, EmailVerifier_getId_Results{Struct: r}}
			return s.GetId(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd458f7ca9d1ba9ff,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailVerifier",
			MethodName:    "verifyEmail",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailVerifier_verifyEmail{c, opts, EmailVerifier_verifyEmail_Params{Struct: p}, EmailVerifier_verifyEmail_Results{Struct: r}}
			return s.VerifyEmail(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// EmailVerifier_getId holds the arguments for a server call to EmailVerifier.getId.
type EmailVerifier_getId struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailVerifier_getId_Params
	Results EmailVerifier_getId_Results
}

// EmailVerifier_verifyEmail holds the arguments for a server call to EmailVerifier.verifyEmail.
type EmailVerifier_verifyEmail struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailVerifier_verifyEmail_Params
	Results EmailVerifier_verifyEmail_Results
}

type EmailVerifier_getId_Params struct{ capnp.Struct }

// EmailVerifier_getId_Params_TypeID is the unique identifier for the type EmailVerifier_getId_Params.
const EmailVerifier_getId_Params_TypeID = 0xe5927352f65eba5c

func NewEmailVerifier_getId_Params(s *capnp.Segment) (EmailVerifier_getId_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailVerifier_getId_Params{st}, err
}

func NewRootEmailVerifier_getId_Params(s *capnp.Segment) (EmailVerifier_getId_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailVerifier_getId_Params{st}, err
}

func ReadRootEmailVerifier_getId_Params(msg *capnp.Message) (EmailVerifier_getId_Params, error) {
	root, err := msg.RootPtr()
	return EmailVerifier_getId_Params{root.Struct()}, err
}

func (s EmailVerifier_getId_Params) String() string {
	str, _ := text.Marshal(0xe5927352f65eba5c, s.Struct)
	return str
}

// EmailVerifier_getId_Params_List is a list of EmailVerifier_getId_Params.
type EmailVerifier_getId_Params_List struct{ capnp.List }

// NewEmailVerifier_getId_Params creates a new list of EmailVerifier_getId_Params.
func NewEmailVerifier_getId_Params_List(s *capnp.Segment, sz int32) (EmailVerifier_getId_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return EmailVerifier_getId_Params_List{l}, err
}

func (s EmailVerifier_getId_Params_List) At(i int) EmailVerifier_getId_Params {
	return EmailVerifier_getId_Params{s.List.Struct(i)}
}

func (s EmailVerifier_getId_Params_List) Set(i int, v EmailVerifier_getId_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailVerifier_getId_Params_Promise is a wrapper for a EmailVerifier_getId_Params promised by a client call.
type EmailVerifier_getId_Params_Promise struct{ *capnp.Pipeline }

func (p EmailVerifier_getId_Params_Promise) Struct() (EmailVerifier_getId_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailVerifier_getId_Params{s}, err
}

type EmailVerifier_getId_Results struct{ capnp.Struct }

// EmailVerifier_getId_Results_TypeID is the unique identifier for the type EmailVerifier_getId_Results.
const EmailVerifier_getId_Results_TypeID = 0xc7e287c5d3518c34

func NewEmailVerifier_getId_Results(s *capnp.Segment) (EmailVerifier_getId_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailVerifier_getId_Results{st}, err
}

func NewRootEmailVerifier_getId_Results(s *capnp.Segment) (EmailVerifier_getId_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailVerifier_getId_Results{st}, err
}

func ReadRootEmailVerifier_getId_Results(msg *capnp.Message) (EmailVerifier_getId_Results, error) {
	root, err := msg.RootPtr()
	return EmailVerifier_getId_Results{root.Struct()}, err
}

func (s EmailVerifier_getId_Results) String() string {
	str, _ := text.Marshal(0xc7e287c5d3518c34, s.Struct)
	return str
}

func (s EmailVerifier_getId_Results) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s EmailVerifier_getId_Results) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailVerifier_getId_Results) SetId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// EmailVerifier_getId_Results_List is a list of EmailVerifier_getId_Results.
type EmailVerifier_getId_Results_List struct{ capnp.List }

// NewEmailVerifier_getId_Results creates a new list of EmailVerifier_getId_Results.
func NewEmailVerifier_getId_Results_List(s *capnp.Segment, sz int32) (EmailVerifier_getId_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailVerifier_getId_Results_List{l}, err
}

func (s EmailVerifier_getId_Results_List) At(i int) EmailVerifier_getId_Results {
	return EmailVerifier_getId_Results{s.List.Struct(i)}
}

func (s EmailVerifier_getId_Results_List) Set(i int, v EmailVerifier_getId_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailVerifier_getId_Results_Promise is a wrapper for a EmailVerifier_getId_Results promised by a client call.
type EmailVerifier_getId_Results_Promise struct{ *capnp.Pipeline }

func (p EmailVerifier_getId_Results_Promise) Struct() (EmailVerifier_getId_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailVerifier_getId_Results{s}, err
}

type EmailVerifier_verifyEmail_Params struct{ capnp.Struct }

// EmailVerifier_verifyEmail_Params_TypeID is the unique identifier for the type EmailVerifier_verifyEmail_Params.
const EmailVerifier_verifyEmail_Params_TypeID = 0x93ee926bb1bd4eea

func NewEmailVerifier_verifyEmail_Params(s *capnp.Segment) (EmailVerifier_verifyEmail_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailVerifier_verifyEmail_Params{st}, err
}

func NewRootEmailVerifier_verifyEmail_Params(s *capnp.Segment) (EmailVerifier_verifyEmail_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return EmailVerifier_verifyEmail_Params{st}, err
}

func ReadRootEmailVerifier_verifyEmail_Params(msg *capnp.Message) (EmailVerifier_verifyEmail_Params, error) {
	root, err := msg.RootPtr()
	return EmailVerifier_verifyEmail_Params{root.Struct()}, err
}

func (s EmailVerifier_verifyEmail_Params) String() string {
	str, _ := text.Marshal(0x93ee926bb1bd4eea, s.Struct)
	return str
}

func (s EmailVerifier_verifyEmail_Params) TabId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s EmailVerifier_verifyEmail_Params) HasTabId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailVerifier_verifyEmail_Params) SetTabId(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s EmailVerifier_verifyEmail_Params) Verification() VerifiedEmail {
	p, _ := s.Struct.Ptr(1)
	return VerifiedEmail{Client: p.Interface().Client()}
}

func (s EmailVerifier_verifyEmail_Params) HasVerification() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s EmailVerifier_verifyEmail_Params) SetVerification(v VerifiedEmail) error {
	if v.Client == nil {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// EmailVerifier_verifyEmail_Params_List is a list of EmailVerifier_verifyEmail_Params.
type EmailVerifier_verifyEmail_Params_List struct{ capnp.List }

// NewEmailVerifier_verifyEmail_Params creates a new list of EmailVerifier_verifyEmail_Params.
func NewEmailVerifier_verifyEmail_Params_List(s *capnp.Segment, sz int32) (EmailVerifier_verifyEmail_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return EmailVerifier_verifyEmail_Params_List{l}, err
}

func (s EmailVerifier_verifyEmail_Params_List) At(i int) EmailVerifier_verifyEmail_Params {
	return EmailVerifier_verifyEmail_Params{s.List.Struct(i)}
}

func (s EmailVerifier_verifyEmail_Params_List) Set(i int, v EmailVerifier_verifyEmail_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailVerifier_verifyEmail_Params_Promise is a wrapper for a EmailVerifier_verifyEmail_Params promised by a client call.
type EmailVerifier_verifyEmail_Params_Promise struct{ *capnp.Pipeline }

func (p EmailVerifier_verifyEmail_Params_Promise) Struct() (EmailVerifier_verifyEmail_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailVerifier_verifyEmail_Params{s}, err
}

func (p EmailVerifier_verifyEmail_Params_Promise) Verification() VerifiedEmail {
	return VerifiedEmail{Client: p.Pipeline.GetPipeline(1).Client()}
}

type EmailVerifier_verifyEmail_Results struct{ capnp.Struct }

// EmailVerifier_verifyEmail_Results_TypeID is the unique identifier for the type EmailVerifier_verifyEmail_Results.
const EmailVerifier_verifyEmail_Results_TypeID = 0xcc99614322e49040

func NewEmailVerifier_verifyEmail_Results(s *capnp.Segment) (EmailVerifier_verifyEmail_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailVerifier_verifyEmail_Results{st}, err
}

func NewRootEmailVerifier_verifyEmail_Results(s *capnp.Segment) (EmailVerifier_verifyEmail_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailVerifier_verifyEmail_Results{st}, err
}

func ReadRootEmailVerifier_verifyEmail_Results(msg *capnp.Message) (EmailVerifier_verifyEmail_Results, error) {
	root, err := msg.RootPtr()
	return EmailVerifier_verifyEmail_Results{root.Struct()}, err
}

func (s EmailVerifier_verifyEmail_Results) String() string {
	str, _ := text.Marshal(0xcc99614322e49040, s.Struct)
	return str
}

func (s EmailVerifier_verifyEmail_Results) Address() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s EmailVerifier_verifyEmail_Results) HasAddress() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailVerifier_verifyEmail_Results) AddressBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s EmailVerifier_verifyEmail_Results) SetAddress(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// EmailVerifier_verifyEmail_Results_List is a list of EmailVerifier_verifyEmail_Results.
type EmailVerifier_verifyEmail_Results_List struct{ capnp.List }

// NewEmailVerifier_verifyEmail_Results creates a new list of EmailVerifier_verifyEmail_Results.
func NewEmailVerifier_verifyEmail_Results_List(s *capnp.Segment, sz int32) (EmailVerifier_verifyEmail_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailVerifier_verifyEmail_Results_List{l}, err
}

func (s EmailVerifier_verifyEmail_Results_List) At(i int) EmailVerifier_verifyEmail_Results {
	return EmailVerifier_verifyEmail_Results{s.List.Struct(i)}
}

func (s EmailVerifier_verifyEmail_Results_List) Set(i int, v EmailVerifier_verifyEmail_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailVerifier_verifyEmail_Results_Promise is a wrapper for a EmailVerifier_verifyEmail_Results promised by a client call.
type EmailVerifier_verifyEmail_Results_Promise struct{ *capnp.Pipeline }

func (p EmailVerifier_verifyEmail_Results_Promise) Struct() (EmailVerifier_verifyEmail_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailVerifier_verifyEmail_Results{s}, err
}

type EmailAgent struct{ Client capnp.Client }

func (c EmailAgent) Send(ctx context.Context, params func(EmailAgent_send_Params) error, opts ...capnp.CallOption) EmailAgent_send_Results_Promise {
	if c.Client == nil {
		return EmailAgent_send_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x8b6f158d70cbc773,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailAgent",
			MethodName:    "send",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailAgent_send_Params{Struct: s}) }
	}
	return EmailAgent_send_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c EmailAgent) AddReceiver(ctx context.Context, params func(EmailAgent_addReceiver_Params) error, opts ...capnp.CallOption) EmailAgent_addReceiver_Results_Promise {
	if c.Client == nil {
		return EmailAgent_addReceiver_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x8b6f158d70cbc773,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailAgent",
			MethodName:    "addReceiver",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(EmailAgent_addReceiver_Params{Struct: s}) }
	}
	return EmailAgent_addReceiver_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type EmailAgent_Server interface {
	Send(EmailAgent_send) error

	AddReceiver(EmailAgent_addReceiver) error
}

func EmailAgent_ServerToClient(s EmailAgent_Server) EmailAgent {
	c, _ := s.(server.Closer)
	return EmailAgent{Client: server.New(EmailAgent_Methods(nil, s), c)}
}

func EmailAgent_Methods(methods []server.Method, s EmailAgent_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x8b6f158d70cbc773,
			MethodID:      0,
			InterfaceName: "email.capnp:EmailAgent",
			MethodName:    "send",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailAgent_send{c, opts, EmailAgent_send_Params{Struct: p}, EmailAgent_send_Results{Struct: r}}
			return s.Send(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x8b6f158d70cbc773,
			MethodID:      1,
			InterfaceName: "email.capnp:EmailAgent",
			MethodName:    "addReceiver",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := EmailAgent_addReceiver{c, opts, EmailAgent_addReceiver_Params{Struct: p}, EmailAgent_addReceiver_Results{Struct: r}}
			return s.AddReceiver(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// EmailAgent_send holds the arguments for a server call to EmailAgent.send.
type EmailAgent_send struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailAgent_send_Params
	Results EmailAgent_send_Results
}

// EmailAgent_addReceiver holds the arguments for a server call to EmailAgent.addReceiver.
type EmailAgent_addReceiver struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  EmailAgent_addReceiver_Params
	Results EmailAgent_addReceiver_Results
}

type EmailAgent_send_Params struct{ capnp.Struct }

// EmailAgent_send_Params_TypeID is the unique identifier for the type EmailAgent_send_Params.
const EmailAgent_send_Params_TypeID = 0xa8eb16da45ad8f97

func NewEmailAgent_send_Params(s *capnp.Segment) (EmailAgent_send_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_send_Params{st}, err
}

func NewRootEmailAgent_send_Params(s *capnp.Segment) (EmailAgent_send_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_send_Params{st}, err
}

func ReadRootEmailAgent_send_Params(msg *capnp.Message) (EmailAgent_send_Params, error) {
	root, err := msg.RootPtr()
	return EmailAgent_send_Params{root.Struct()}, err
}

func (s EmailAgent_send_Params) String() string {
	str, _ := text.Marshal(0xa8eb16da45ad8f97, s.Struct)
	return str
}

func (s EmailAgent_send_Params) Email() (EmailMessage, error) {
	p, err := s.Struct.Ptr(0)
	return EmailMessage{Struct: p.Struct()}, err
}

func (s EmailAgent_send_Params) HasEmail() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAgent_send_Params) SetEmail(v EmailMessage) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewEmail sets the email field to a newly
// allocated EmailMessage struct, preferring placement in s's segment.
func (s EmailAgent_send_Params) NewEmail() (EmailMessage, error) {
	ss, err := NewEmailMessage(s.Struct.Segment())
	if err != nil {
		return EmailMessage{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// EmailAgent_send_Params_List is a list of EmailAgent_send_Params.
type EmailAgent_send_Params_List struct{ capnp.List }

// NewEmailAgent_send_Params creates a new list of EmailAgent_send_Params.
func NewEmailAgent_send_Params_List(s *capnp.Segment, sz int32) (EmailAgent_send_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailAgent_send_Params_List{l}, err
}

func (s EmailAgent_send_Params_List) At(i int) EmailAgent_send_Params {
	return EmailAgent_send_Params{s.List.Struct(i)}
}

func (s EmailAgent_send_Params_List) Set(i int, v EmailAgent_send_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAgent_send_Params_Promise is a wrapper for a EmailAgent_send_Params promised by a client call.
type EmailAgent_send_Params_Promise struct{ *capnp.Pipeline }

func (p EmailAgent_send_Params_Promise) Struct() (EmailAgent_send_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailAgent_send_Params{s}, err
}

func (p EmailAgent_send_Params_Promise) Email() EmailMessage_Promise {
	return EmailMessage_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EmailAgent_send_Results struct{ capnp.Struct }

// EmailAgent_send_Results_TypeID is the unique identifier for the type EmailAgent_send_Results.
const EmailAgent_send_Results_TypeID = 0x81f33f1803485545

func NewEmailAgent_send_Results(s *capnp.Segment) (EmailAgent_send_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailAgent_send_Results{st}, err
}

func NewRootEmailAgent_send_Results(s *capnp.Segment) (EmailAgent_send_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return EmailAgent_send_Results{st}, err
}

func ReadRootEmailAgent_send_Results(msg *capnp.Message) (EmailAgent_send_Results, error) {
	root, err := msg.RootPtr()
	return EmailAgent_send_Results{root.Struct()}, err
}

func (s EmailAgent_send_Results) String() string {
	str, _ := text.Marshal(0x81f33f1803485545, s.Struct)
	return str
}

// EmailAgent_send_Results_List is a list of EmailAgent_send_Results.
type EmailAgent_send_Results_List struct{ capnp.List }

// NewEmailAgent_send_Results creates a new list of EmailAgent_send_Results.
func NewEmailAgent_send_Results_List(s *capnp.Segment, sz int32) (EmailAgent_send_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return EmailAgent_send_Results_List{l}, err
}

func (s EmailAgent_send_Results_List) At(i int) EmailAgent_send_Results {
	return EmailAgent_send_Results{s.List.Struct(i)}
}

func (s EmailAgent_send_Results_List) Set(i int, v EmailAgent_send_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAgent_send_Results_Promise is a wrapper for a EmailAgent_send_Results promised by a client call.
type EmailAgent_send_Results_Promise struct{ *capnp.Pipeline }

func (p EmailAgent_send_Results_Promise) Struct() (EmailAgent_send_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailAgent_send_Results{s}, err
}

type EmailAgent_addReceiver_Params struct{ capnp.Struct }

// EmailAgent_addReceiver_Params_TypeID is the unique identifier for the type EmailAgent_addReceiver_Params.
const EmailAgent_addReceiver_Params_TypeID = 0xfacf412b11767e9e

func NewEmailAgent_addReceiver_Params(s *capnp.Segment) (EmailAgent_addReceiver_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_addReceiver_Params{st}, err
}

func NewRootEmailAgent_addReceiver_Params(s *capnp.Segment) (EmailAgent_addReceiver_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_addReceiver_Params{st}, err
}

func ReadRootEmailAgent_addReceiver_Params(msg *capnp.Message) (EmailAgent_addReceiver_Params, error) {
	root, err := msg.RootPtr()
	return EmailAgent_addReceiver_Params{root.Struct()}, err
}

func (s EmailAgent_addReceiver_Params) String() string {
	str, _ := text.Marshal(0xfacf412b11767e9e, s.Struct)
	return str
}

func (s EmailAgent_addReceiver_Params) Port() EmailSendPort {
	p, _ := s.Struct.Ptr(0)
	return EmailSendPort{Client: p.Interface().Client()}
}

func (s EmailAgent_addReceiver_Params) HasPort() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAgent_addReceiver_Params) SetPort(v EmailSendPort) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// EmailAgent_addReceiver_Params_List is a list of EmailAgent_addReceiver_Params.
type EmailAgent_addReceiver_Params_List struct{ capnp.List }

// NewEmailAgent_addReceiver_Params creates a new list of EmailAgent_addReceiver_Params.
func NewEmailAgent_addReceiver_Params_List(s *capnp.Segment, sz int32) (EmailAgent_addReceiver_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailAgent_addReceiver_Params_List{l}, err
}

func (s EmailAgent_addReceiver_Params_List) At(i int) EmailAgent_addReceiver_Params {
	return EmailAgent_addReceiver_Params{s.List.Struct(i)}
}

func (s EmailAgent_addReceiver_Params_List) Set(i int, v EmailAgent_addReceiver_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAgent_addReceiver_Params_Promise is a wrapper for a EmailAgent_addReceiver_Params promised by a client call.
type EmailAgent_addReceiver_Params_Promise struct{ *capnp.Pipeline }

func (p EmailAgent_addReceiver_Params_Promise) Struct() (EmailAgent_addReceiver_Params, error) {
	s, err := p.Pipeline.Struct()
	return EmailAgent_addReceiver_Params{s}, err
}

func (p EmailAgent_addReceiver_Params_Promise) Port() EmailSendPort {
	return EmailSendPort{Client: p.Pipeline.GetPipeline(0).Client()}
}

type EmailAgent_addReceiver_Results struct{ capnp.Struct }

// EmailAgent_addReceiver_Results_TypeID is the unique identifier for the type EmailAgent_addReceiver_Results.
const EmailAgent_addReceiver_Results_TypeID = 0x8e8e3d68615d430c

func NewEmailAgent_addReceiver_Results(s *capnp.Segment) (EmailAgent_addReceiver_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_addReceiver_Results{st}, err
}

func NewRootEmailAgent_addReceiver_Results(s *capnp.Segment) (EmailAgent_addReceiver_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EmailAgent_addReceiver_Results{st}, err
}

func ReadRootEmailAgent_addReceiver_Results(msg *capnp.Message) (EmailAgent_addReceiver_Results, error) {
	root, err := msg.RootPtr()
	return EmailAgent_addReceiver_Results{root.Struct()}, err
}

func (s EmailAgent_addReceiver_Results) String() string {
	str, _ := text.Marshal(0x8e8e3d68615d430c, s.Struct)
	return str
}

func (s EmailAgent_addReceiver_Results) Handle() util.Handle {
	p, _ := s.Struct.Ptr(0)
	return util.Handle{Client: p.Interface().Client()}
}

func (s EmailAgent_addReceiver_Results) HasHandle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EmailAgent_addReceiver_Results) SetHandle(v util.Handle) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// EmailAgent_addReceiver_Results_List is a list of EmailAgent_addReceiver_Results.
type EmailAgent_addReceiver_Results_List struct{ capnp.List }

// NewEmailAgent_addReceiver_Results creates a new list of EmailAgent_addReceiver_Results.
func NewEmailAgent_addReceiver_Results_List(s *capnp.Segment, sz int32) (EmailAgent_addReceiver_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EmailAgent_addReceiver_Results_List{l}, err
}

func (s EmailAgent_addReceiver_Results_List) At(i int) EmailAgent_addReceiver_Results {
	return EmailAgent_addReceiver_Results{s.List.Struct(i)}
}

func (s EmailAgent_addReceiver_Results_List) Set(i int, v EmailAgent_addReceiver_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// EmailAgent_addReceiver_Results_Promise is a wrapper for a EmailAgent_addReceiver_Results promised by a client call.
type EmailAgent_addReceiver_Results_Promise struct{ *capnp.Pipeline }

func (p EmailAgent_addReceiver_Results_Promise) Struct() (EmailAgent_addReceiver_Results, error) {
	s, err := p.Pipeline.Struct()
	return EmailAgent_addReceiver_Results{s}, err
}

func (p EmailAgent_addReceiver_Results_Promise) Handle() util.Handle {
	return util.Handle{Client: p.Pipeline.GetPipeline(0).Client()}
}

const schema_dd10df585a82c6d8 = "x\xda\x9cWo\x8c\x1ce\x19\x7f\x9eywvv\xdb" +
	"\xdd\xee\x8d\xb3\x88\x1e!\x17\x9a#\xd2\xc3\xd6\xf6\xe8\x07" +
	"\xa9\xe2^\x0f\x8e\xb4\xa6\xd5\x9d\xdd\xa3i\x0f0\x9d\xdd" +
	"}\xdb\x9b\xba;s\xcc\xcc\x95[C\xadUk1X" +
	"h\xcb\x07\xd4P\x13#M\xd0\xf4\xc4F\x13\x0dB\x8c" +
	"\xd1\x13\xa9m0\xa5\x1f\x1aC\xa5\x90\"\xa0Pc\x04" +
	"%D3\xe6yw\xe7\xcf\xcd\xed\x1d\x8d\xdfv\x7f\xf3" +
	"\xcc\xef\xf9\xbd\xcf\xbf\xf7\x99\xb5vjDZ'?\xb5" +
	"\x0c@\x7f@N\xfbcwmb\x1f)\xfd\xf3\x00\xa8" +
	"\xd7\"@J\x01\xb8\xe56y%B\xcaw\x9f\xfb\xc3" +
	"\xd4\xe1k\xec\x87@\xcd1\xff\xc2\xef\xbe:\xb1\xfd\xe5" +
	"\xbe\x8b\x00\xa8\xdd(?\xad\xad\x96\x15\x00m\x95|H" +
	"\xdbG\xbf\xfc\xdc\xed\xf7\x1a\x93\xb7=\xfcp\x87FF" +
	"\xe2\xe1\xf2\x04\x02j\xf7\xc9%@\xffS\xbf\x1c(\x9e" +
	"\xbf\xfb\xaeG@]\x89\xfe\x9b\x17~<6\xfe\xe0\x99" +
	"\x1f\x80,\x91\xe1Q\xd9A\xed\x84`\xfc\xbe|?\xa0" +
	"\xff\xe2\xf1o\xfe\xc7\xc8\xb5\x8f\x80\xda\x8f\xfe\xe9g\x9e" +
	"\xdf\xf2\xab\xeb\xbf\xf6V\xd7\x18\xd3\xc3\xa8\xa9i2\xce" +
	"\xa7\xc9\xf8\xaf\x9f{\xf6\xd4\x17\x8f]y\x94\x8c\xa1k" +
	"t_z\x0f\xb9\xde'\x0c\x8e\xfd\xe2\xee\x93G\x8f\xdd" +
	"\xf9\x98`\x9bx\x7f\xeb\x9d\xd2?\x1ez\x0fdF\x86" +
	"\x17\x89\xedm\xc1\xf6f\xfa)@\xff\xfc\xe5m\x7f\x9e" +
	"\xfb\xcd\xcc\xe3]6q\x90\x83\x8a`;\xaa\xd0AB" +
	"\xed\x89\xb0\xdcrJ\xf9\x10j\xbfV\xae\x05\xd0\xce*" +
	"\x8avV\xf9\x18\x80\x7fe\xc7\x99-7\xff|\xf6D" +
	"\x9cnN\x19&\xba\xb3\x82\xee\xb1Gf\xc7\xfe\xf4\xe1" +
	"\xbf=\x19\x0f\xdc\xdbJ?\x19\xbc+\x0c^\xf9\xcc\xb6" +
	"o\\yi\xf6$\xa89\x8c\xfc\x89cj\xd7dN" +
	"k7d\xe8\xd7\xf5\x19:\xe9\xe7\x07o:\xde?\x97" +
	"\xfdi\xc2\x96\xb2\xaa\xb53\xafj\x07\x85\xed\x81\xcc\xeb" +
	"\x80\x7f\x7f\xe6G\xab\x1e\xfd\x92\xf3\xac\xd0%\xd2\xce\xb3" +
	"\x0e\xa5}\xfd\xb7\xf4\x17\xe7\x0e\xbd\xfa\\\\\xb1\x9e\x1d" +
	"%A;\xb2$h\xe4\xc8\xe5\x95\xb7\x1b\xdf9\x137" +
	"h\xd3\xbb\xa8\x1d\x14\x06\x9fX\xbd\xcd|}\xc7;/" +
	"\x80\x9e\xc3\xb8\x8c\x1c9?\x91=\xad\x9d\xca\xd2;\xb3" +
	"\xd9\x97%@\xff\xc2{\xd7=\xff\x97\x9f\xd5\xff\x08\x91" +
	"\x92|~\x03)\xf1\x7fx\xdd\xf1\xd3\xff\xde~~A" +
	"\x01\xbe\x9b;\xa7a\x9e\xd8\xfe\x9b;\xa4\xe9\xf4\xcb\xbf" +
	"\xe7\xe9/\xfc\xab\xe2\x1e{-Fsk\x87&\xac\x9f" +
	"$\xcd\x0d\xf9s\xda\xea<\xe5\xeb\xd6\xfc!\xed\xb0\xa0" +
	"\x09\xcb#i<\x9d?\xa7\x1d\x10\xc6\x87\xf3J\xd7\xf8" +
	"{_\xde\xab\xde\xbc\xf1\x85\xf7\xe3\xb9\x9b\xceW(\x12" +
	"\x07\xf2%\xd8\xe1\xf3\x96a6\xd7\xd4\x0d6eMm" +
	"\x18\xa3?\x1bws\xcb[\xe3r\xab1X\xe1\xee\xb4" +
	"\xd2\xf4\xdc\xd0\x0a\x03\xabR\xc7\xac\x8c\xa8g\x98\x0c\x10" +
	"\x16\x09\x06\xdd\xaa\xae\x1b\x02I\xbdQA\x0cE`\xd0" +
	"\x82\xeaGk \xa9\xaaR /#\xe8\x1b\x8dF\x85" +
	"\xd7\xb9\x09\xca^\xee\x8c`\x191\xf4\x98J\xe8\x0aL" +
	"\xf7rG\xc8k2\xcf\xd5S,\x05\x90B\x005\xbf" +
	"\x01@\xcf0\xd4\x8b\x12\x96&\x0d\xab\xd1\xe4\xa8\xfa\x97" +
	"Fw\xee<9\xf8\xce\xb7\x01\x10UH\x90o\xe3\x8e" +
	"\xb9\xcb\xe4\x0d\xe1\xa4\xca\xadF\xd9v\xbc5e\xfb~" +
	"\xee\xd4\x06\xec\x99qcw\xe7\x94\x81\x8bU{\x00\xf4" +
	"\x9b\x18\xea\xeb%T\x11\x8bH\xe0\xba!\x00\xfd\xe3\x0c" +
	"\xf5OJ\xe8\xef\x15\x8cu\x03\x0a\x9ei[\xd8\x17\xb5" +
	"7 \xf6\x01\x16\xa6l\xc7\xc3\xbeh\x86t\xe0\x1e\xb9" +
	"\xb8\x1a9\x9f\xed%g\x02@_\xcbP\xff\xb4\x84\xfe" +
	".\xc7nm2-\x0f\x00\xb0/j\xd6\xae\xd3\xa6\xe9" +
	"z\x9b\x1b\x9bL`\x96\x879\x900\x07\xbd\xa2\xdf\x8d" +
	"\x92\xb3F\x1c\xae-\xb0\xc1\xb2Qp\x8c\x96\x1bW3" +
	"\x0c\xa0\x0f2\xd4\xd7\xc6\xd4\xac\xde\x13\x05g\xc03j" +
	"\x9b\x1b\x98\x07\x09\xf3\xb0 TjT\xdd\x89T\xb1\x05" +
	"\xa9Z\x10\x93\\\xa8b\x8c\x8e\x7f\x07C\xbd\x1cS\xb1" +
	"u\x14@\xdf\xc4P\x1f\x97P\x95\xa4\"J\x00\xaaN" +
	"\xf5\xb2\x85\xa1\xbe=\xcc\x1bw\x80E\x0a\xf7\x1b\x8d\x86" +
	"\xc3]7\x08M\xa9a\xb7\x0c\xd3Z\"Ra\xce&" +
	"M\xcb\xdb\xd8y;\x8cT\xacRG\xa3J\x0d\x9d," +
	"L\xcf\xe2\xe7\xef\xfaAOOa\xfc\xf2\xc2\x9a\xdf\x89" +
	"\x8c=\x03\x0a\x85\x86\xc9z\x06c\xd7\x0a@4p\x00" +
	"\x96\xaa:1\x04\xca\x03FR\xfap$}@\xbc\x8e" +
	"}\xd1H]\xb4\x98c\x83\xa5l8\x06\xfb\xff9\xa5" +
	"\x88\xb3\x138\x80DS\x8c\xf6*\xc3\xa1\xa8S\x92i" +
	"-XF\x8b/Hj\xcc\x8f\xe7\x19\xf5\xc9\x96\xd2\x9d" +
	"y}\xa1'\xa3\x06\xa0\xefd\xa87c\x9e\xcc\xef\x02" +
	"\xe8M\x86\xfaL\xac\xd4\xa6+\x00\xba\xc7P\xff\x8a\x84" +
	"*cEd\x00\xea>\x12:\xc3P\xff\xba\x84~\xdd" +
	"\xb6<ny\xe3\xa0\xb4\xa7\"1]\xf4\x0e4\xdd)" +
	"\xdb5=\x93\xd9V\xf2\xe1f\xc0F\x80\xed\xefba" +
	"\x8f]e\x89V\xf8\x80;\x1d\x1f\xf6=\x9a\x7f7\xf7" +
	"6w\xae\x85\xa6\xe7B<}\xfdQ\xfa\x98\xd9X\xc2" +
	"w\xcfA\x12\xf8\xfe\xa0\xfeX<A[\xb9\xeb\x1a\xbb" +
	"y\xa7\x10\xd6\x07,\xda\xbd8\x04P\xdd\x8e\x0c\xab\x0d" +
	"\x94\xb0\x9b \xcd\x10\xf0=\x04O\"\xe5\x08E\x8e4" +
	"\x8e\xfd\x00\xd5\x9d\x84?@8\x93D\x9a\xb4\xb6\xc0=" +
	"\xc2\x8f\x10\x9ebEL\xd1E\x8b+\x01\xaa\x0f\x12\xfe" +
	"\x04\xe1r\xaa\x882\xed\x888\x0aP}\x9c\xf0'\x09" +
	"O\xcbEL\xd3j\x81\x15\x80\xea\x13\x84\xff\x84p%" +
	"]\xa4+Y\x9b\xc5\x09\x80\xeaI\xc2\x7fOxF)" +
	"b\x06@\x9b\x13\xf6\xbf%\xfc\x12\xe1\xd9L\x11\xb3\x00" +
	"\xdaE\xc1\x7f\x81\xf0\xcb\x84/\xcb\x16q\x19\x80\xf6\x8a" +
	"8\xd7K\x84\xbfA\xf8\xf2eE\\\x0e\xa0\xbd&\xf0" +
	"K\x84\xbfExny\x11s\xb4Pb\x0d\xa0\xfa\x06" +
	"\xe1)I\xc2B\xc3\xf08\xca \xa1\x0cX\xa0{c" +
	"\xe1Db\x9e\x8d+\x00\xcb\x0c\x13\xcfV\x00\xb2z}" +
	"\xd1gJm\x89\x87\xfb\x1d>\xd5l\x8f\xdb=\x06`" +
	"\xab\x93\xd8x\x85\xfb\x0e\xdf\xc5\x1dn\xd5\x81q7\xe0" +
	"\xa4g+\x00}\xd3\xaa\x08.@;\xf1h\xbf;]" +
	"\xdb\xc3\xeb\xe1-W\xf0\xf8L\xf4g\xd2k5C\x07" +
	"F\xa7\xdd9(\x96\xe7F\xaa\xc3\xe5\xb5\xa3\xfa\x03G" +
	"g\xa5\xc4\xe7\xb7\x94\x94l\x03t\xa2\x15*\xd8\x0f1" +
	"\xd8o\xd5u\xc3\xc1\x0a\x15|A`\xb0\xda\x06+\xd4" +
	"\x80h\xc9\x91\xee\xe5\xd5\x1ek\x81b\x98\xcd\xf9;\x14" +
	"[\xac\x91Ke1\xda{\xc8\x9b\x7f\xb9D\x1f;\xc9" +
	"\xcb\xa5+=\xf8\x86\xc0`U\x8em\x7f\xc1\xe7\x0a\xd2" +
	":\x0f\xb4\xcf'\xb6\xbf`\x08\x81\xc2]w\xber)" +
	"y\xebaS\x08\x8a\xbe\x97z\xdcvW\xbb>\x8a\xc3" +
	"\xe3\xbc\x993\x14\xcd\x9c\xce\x92\xa6\xc6\xaeJ\xb1\x90\xfc" +
	"/\x00\x00\xff\xff\xbe\x8a\xefT"

func init() {
	schemas.Register(schema_dd10df585a82c6d8,
		0x81f33f1803485545,
		0x8b6f158d70cbc773,
		0x8e8e3d68615d430c,
		0x8f555bd4141fbb3b,
		0x90790c61fc899dd3,
		0x93ee926bb1bd4eea,
		0x97469291ac5bb892,
		0x9c78c3c5de56e4d4,
		0xa3cc885445aed8e9,
		0xa5adb72b4ccc59ee,
		0xa8eb16da45ad8f97,
		0xacaddcee86563ee1,
		0xb309c51a9d28244f,
		0xbd727a009329aabc,
		0xc7e287c5d3518c34,
		0xcc99614322e49040,
		0xcff459e769562d2f,
		0xd063b4e6c91bf8d8,
		0xd458f7ca9d1ba9ff,
		0xe5927352f65eba5c,
		0xec831dbf4cc9bcca,
		0xf88bf102464dfa5a,
		0xfacf412b11767e9e)
}
