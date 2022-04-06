package database

/*
Типа база данных
*/
type User struct {
	LastName    string
	FirtName    string
	Age         int
	Score       int
	NumberPhone string
	Raiting     float32 // 0.0-5.0
	Login       string
	Password    string
}
type Data struct {
	DataBaseName string
	Users        []User
}

//Тип база
func DB() Data {
	DataDB := Data{
		DataBaseName: "mssql: User",
		Users: []User{
			{
				Login:       "R2D2Dan",
				Password:    "qwerty1",
				LastName:    "Danil",
				FirtName:    "La",
				Age:         22,
				Score:       9999,
				NumberPhone: "79998885522",
				Raiting:     4.9,
			},

			{Login: "ArtLebedev", Password: "YaKryt3000"},
			{Login: "Vasya2020", Password: "test"},
			{Login: "test@test.com", Password: "test"},
		},
	}
	return DataDB

}
