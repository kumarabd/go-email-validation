package evalidate

import (
	"github.com/kumarabd/appkit/errors"
)

var (
	ErrBadFormat        = errors.New("invalid format")
	ErrUnresolvableHost = errors.New("unresolvable host")
	ErrTimeout          = errors.New("Dial.Timeout.Error", err.Error())
	ErrNotReachable     = errors.New("Client.Not.Reachable", err.Error())
	ErrClientEmail      = errors.New("Client.Mail.Error", err.Error())
	ErrClientRct        = errors.New("Client.Recipient.Error", err.Error())
)
