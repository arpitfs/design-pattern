package factory

type Config struct {
	DbType string
	Url    string
}

func SetUpDatabase(config Config) (Database, error) {
	if config.DbType == "Influx" {
		return &Influx{url: config.Url}, nil
	} else if config.DbType == "Postgres" {
		return &Postgres{url: config.Url}, nil
	} else {
		return nil, nil
	}
}
