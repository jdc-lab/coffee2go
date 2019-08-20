package uc

func (i interactor) ConnectServer(host, username, password string) (user Chat, token string, err error) {
	/*	user, err := i.userRW.GetByEmailAndPassword(email, password)
		if err != nil {
			return nil, "", err
		}
		if user == nil {
			return nil, "", ErrNotFound
		}

		token, err := i.authHandler.GenUserToken(user.Name)
		if err != nil {
			return nil, "", err
		}
	*/
	panic("ConnectServer usecase not implemented yet")
	//return user, token, nil
}

func (i interactor) ConnectPreset() (host, username, password string) {
	return i.conf.GetConnectionPreset()
}
