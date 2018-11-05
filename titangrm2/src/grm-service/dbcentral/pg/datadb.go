package pg

type DataCentralDB struct {
	Central
}

func ConnectDataDB(host, user, password string) (DataCentralDB, error) {
	central, err := ConnectDB(host, DataDBName, user, password)
	if err != nil {
		return DataCentralDB{}, err
	}
	return DataCentralDB{central}, err
}

func ConnectDataDBUrl(url string) (DataCentralDB, error) {
	//	central, err := ConnectDB(host, MetaDBName, user, password)
	central, err := ConnectDBUrl(url)
	if err != nil {
		return DataCentralDB{}, err
	}
	return DataCentralDB{central}, err
}
