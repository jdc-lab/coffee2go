package uc

import "github.com/jdc-lab/coffee2go/domain"

func (i interactor) ChatSend(message domain.Message) (err error) {
	panic("ChatSend usecase not implemented yet")
}

func (i interactor) ChatRegisterForPush(sessionID string) (pushToken string, err error) {
	_, err = i.session.Get(sessionID)

	if err != nil {
		return "", err
	}

	if pushToken, err := i.push.Register(sessionID); err != nil {
		return "", err
	} else {
		return pushToken, nil
	}
}

func (i interactor) ChatPushNewMessage() (err error) {
	panic("ChatPushNewMessage usecase not implemented yet")
}

func (i interactor) ChatGetContacts() (err error) {
	panic("ChatGetContacts usecase not implemented yet")
}

func (i interactor) ChatSwitch() (err error) {
	panic("ChatGetContacts usecase not implemented yet")
}
