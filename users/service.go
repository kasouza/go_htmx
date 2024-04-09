package users

func maybeLoginUser(email string, password string) (*User, error) {
	user, err := FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// TODO: verify password

	return user, nil
}
