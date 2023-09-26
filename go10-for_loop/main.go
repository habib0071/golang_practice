package main 

import "log"

func main() {
	i := 1
	for i <= 10  {
		log.Println(i)
		i = i  + 1
	}
	for j := 7; j <= 9; j++ {
		log.Println(j)
	}

	for {
		log.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		log.Println(n)
	}
	animals := []string {"dog", "cat", "kona", "cow", "horse"}
	for _/*i*/, animal := range animals {
		log.Println(animal)
	}

	things := make(map[string]string)
	things["flowers"] = "rose"
	things["animal"] = "cat"
	things["color"] = "pink"
	for i, thing := range things{
		log.Println(i, thing)
	}


	var firstline = "Once upon a midnight dreary"
	

	for i, l := range firstline{
		log.Println(i, " -> ", string(l))
	}

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{"habib", "khan", "habibkhan@gamil.com", 30 })
	users = append(users, User{"ohona", "khan", "habibkhan@gamil.com", 23 })
	users = append(users, User{"sara", "khan", "habibkhan@gamil.com", 8})

	for i, l := range users {
		log.Println(i, l.FirstName, l.LastName, l.Email, l.Age)
	}


}