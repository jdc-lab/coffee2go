package main

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
