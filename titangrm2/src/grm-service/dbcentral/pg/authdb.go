package pg

type AuthCentralDB struct {
	Central
}

func ConnectAuthDB(host, user, password string) (AuthCentralDB, error) {
	central, err := ConnectDB(host, AuthDBName, user, password)
	if err != nil {
		return AuthCentralDB{}, err
	}
	return AuthCentralDB{central}, err
}
