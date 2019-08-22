package uc

func (i interactor) ConnectServer(host, username, password string) (sessionID string, err error) {
	serverConnection, err := i.connection.Connect(host, username, password)

	if err != nil {
		return "", err
	}

	serverConnection.Run()

	return i.session.Add(&serverConnection)
}

func (i interactor) ConnectPreset() (host, username, password string) {
	return i.conf.GetConnectionPreset()
}
