package factory

type Database interface {
	Connect() error
}

func Factory() {
	config := Config{
		DbType: "Postgres",
		Url:    "localhost",
	}

	db, _ := SetUpDatabase(config)
	db.Connect()
}
