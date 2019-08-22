package main

import "github.com/jdc-lab/coffee2go/uc"

type dummyPush struct {
}

func (d dummyPush) Register() (err error) {
	panic("implement me")
}

func (d dummyPush) Send() (err error) {
	panic("implement me")
}

type dummyConf struct {
}

func (d dummyConf) GetConnectionPreset() (host, username, password string) {
	panic("implement me")
}

type dummySession struct {
}

func (d dummySession) Add(session *uc.Chat) (sessionID string, err error) {
	return "dummyID", nil
}

func (d dummySession) Get(sessionId string) (session *uc.Chat, err error) {
	panic("implement me")
}
