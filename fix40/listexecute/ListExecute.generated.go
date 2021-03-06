package listexecute

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//ListExecute is the fix40 ListExecute type, MsgType = L
type ListExecute struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ListExecute from a quickfix.Message instance
func FromMessage(m quickfix.Message) ListExecute {
	return ListExecute{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ListExecute) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ListExecute initialized with the required fields for ListExecute
func New(listid field.ListIDField) (m ListExecute) {
	m.Header = fix40.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("L"))
	m.Set(listid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ListExecute, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "L", r
}

//SetText sets Text, Tag 58
func (m ListExecute) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetListID sets ListID, Tag 66
func (m ListExecute) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetWaveNo sets WaveNo, Tag 105
func (m ListExecute) SetWaveNo(v string) {
	m.Set(field.NewWaveNo(v))
}

//GetText gets Text, Tag 58
func (m ListExecute) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m ListExecute) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetWaveNo gets WaveNo, Tag 105
func (m ListExecute) GetWaveNo() (f field.WaveNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m ListExecute) HasText() bool {
	return m.Has(tag.Text)
}

//HasListID returns true if ListID is present, Tag 66
func (m ListExecute) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasWaveNo returns true if WaveNo is present, Tag 105
func (m ListExecute) HasWaveNo() bool {
	return m.Has(tag.WaveNo)
}
