package test

import "errors"

type Email struct {
	sendTo []string
	cc     []string
}

type IEmail interface {
	Send(sendTo []string, cc []string) error
}

func (e *Email) Send(sendTo []string, cc []string) error {
	if len(sendTo) == 0 {
		return errors.New("sendTo cannot be empty")
	}
	if len(cc) == 0 {
		return errors.New("cc cannot be empty")
	}

	e.sendTo = sendTo
	e.cc = cc

	return nil
}

func DoSendEmail(emailI IEmail) {
	if err := emailI.Send([]string{"tiki@gmail.com"}, []string{"taka@gmail.com"}); err != nil {
		panic(err)
	}
}