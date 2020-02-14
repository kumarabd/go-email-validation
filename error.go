package evalidate

import (
	"github.com/kumarabd/appkit/errors"
)

var (
	ErrBadFormat = errors.New("Invalid.Format", "Invalid format for email")
)

func ErrUnresolvableHost(err error) error {
	return errors.New("Unresolvable.Host", err.Error())
}
func ErrTimeout(err error) error {
	return errors.New("Dial.Timeout.Error", err.Error())
}
func ErrNotReachable(err error) error {
	return errors.New("Client.Not.Reachable", err.Error())
}
func ErrClientEmail(err error) error {
	return errors.New("Client.Mail.Error", err.Error())
}
func ErrClientRct(err error) error {
	return errors.New("Client.Recipient.Error", err.Error())
}
