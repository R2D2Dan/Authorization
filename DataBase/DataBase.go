package database

import "log"

/*
Типа база данных
*/
type Users struct {
	/*
		LastName    string
		FirtName    string
		Age         int
		Score       int
		NumberPhone string
		Raiting     float32 // 0.0-5.0
		Login       string
		Password    string
	*/
	Id_User int
	User    map[int]map[string]interface{}
}
type Data struct {
	DataBaseName string
	DATA         []Users
}

//Тип база
func DB() *Data {
	user1 := make(map[int]map[string]interface{})

	user1[1] = make(map[string]interface{})
	user1[1]["Login"] = "R2D2Dan"
	user1[1]["Password"] = "qwerty1"
	user1[1]["LastName"] = "Su"
	user1[1]["FirtName"] = "Danil"
	user1[1]["Score"] = 9999
	user1[1]["Age"] = 22
	user1[1]["NumberPhone"] = 9998885522
	user1[1]["Raiting"] = 4.9

	user2 := make(map[int]map[string]interface{})
	user2[2] = make(map[string]interface{})
	user2[2]["Login"] = "Bob_303"
	user2[2]["Password"] = "Bob3000"
	user2[2]["LastName"] = "Tr"
	user2[2]["FirtName"] = "Bob"
	user2[2]["Score"] = 3405
	user2[2]["NumberPhone"] = 9998542641
	user2[2]["Raiting"] = 4.6

	data := Data{
		DataBaseName: "mssql",
		DATA: []Users{
			{
				Id_User: 1,
				User:    user1,
			},
			{
				Id_User: 2,
				User:    user2,
			},
		},
	}
	/*
		t := data.DATA
		log.Println(t[0].User[1]["Login"])

		for i := 0; i < len(data.DATA); i++ {
			log.Println(t[i].User[i+1]["Login"])
		}
		log.Println("111")
		for i, k := range data.DATA {
			log.Println(k.User[i+1]["Login"])
		}
	*/
	return &data

}

//Проверка клиента в базе
func IsClient(u, p string) (int, *Users, bool) {
	db := DB().DATA

	for i, k := range db {
		login, oklg := k.User[i+1]["Login"]
		pass, okps := k.User[i+1]["Password"]

		if !oklg && !okps {
			log.Println("Ключ не найден")
			return 0, &db[0], false
		}
		if login == u && pass == p {
			log.Println("Пользователь найден")
			return DB().DATA[i].Id_User, &db[i], true
		}

	}

	return 0, &db[0], false

}
