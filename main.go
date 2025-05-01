// package main

// func main() {
// 	print("Hello, World!")
// 	print(add(220, 50))
// }

// func add(a, b int) int {
// 	return a + b
// }

// package main

// func main() {
// 	http.HandleFunc("/api/threads", handleThreads)
// 	log.Fatal(http.ListenAndServe(":4000", nil))
// }

// func handleThreads(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		Post {
// 			Title:   "Hello World",
// 			Content: "This is my first post",
// 			active: true,
// 		}
// 	case "POST":
// 		// Создать новый пост
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// type Thread struct {
// 	ID        int    `json:"id"`
// 	Avatar    string `json:"avatar"`
// 	User      string `json:"user"`
// 	Time      string `json:"time"`
// 	Text      string `json:"text"`
// 	LikeCount int    `json:"likeCount"`
// 	Comments  int    `json:"comments"`
// 	Resends   int    `json:"resends"`
// 	Views     int    `json:"views"`
// }

// var threads = []Thread{
// 	{	1, "/danya.jpg", "Danya", "5ч", "My name is Danya, I am a Frontend Developer", 200, 200, 200, 200},
// 	{2, "/sardor.jpg", "Sardor", "2ч", "My name is Sardor, I am a Frontend Developer", 500, 245, 464, 654},
// 	{3, "/avatar.jpg", "Erkin", "1ч", "Hello world, it's Erkin here! Writing some code 🚀", 312, 98, 112, 430},
// 	{4, "/avatar.jpg", "Murodjon", "3ч", "Building something cool with React and Next.js 🔥", 220, 150, 133, 512},
// 	{5, "/avatar.jpg", "Ramizjon", "6ч", "Just finished my new portfolio website 🎉", 610, 342, 278, 800},
// 	{6, "/avatar.jpg", "Daler", "8ч", "Working on some TypeScript magic ✨", 180, 76, 90, 300},
// 	{7, "/avatar.jpg", "Ozodbek", "9ч", "I love using shadcn/ui in my projects!", 400, 210, 150, 550},
// 	{8, "/avatar.jpg", "Konstantin", "10ч", "Frontend is fun when you understand the logic 🧠", 275, 115, 134, 460},
// 	{9, "/avatar.jpg", "Shaxriyor", "11ч", "CSS can be tricky, but Tailwind makes it easier 🧩", 490, 198, 205, 620},
// 	{10, "/avatar.jpg", "Alex", "12ч", "Sharing my open-source project soon, stay tuned 💡", 305, 150, 160, 580},
// }

// func threadsHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(threads)
// 		return
// 	}

// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// }

// func main() {
// 	http.HandleFunc("/api/threads", threadsHandler)
// 	log.Println("Server running on http://localhost:4000")
// 	log.Fatal(http.ListenAndServe(":4000", nil))
// }

// func Hello() {
// 	print("Hello, World!")
// }

// package main

// import "fmt"

// func main() {
// 	var name string = "Sardor"
// 	var age int = 15
// 	var isFrontend bool = true
// 	city := "Samarkand"
// 	var surname string
// 	fmt.Println("Введите вашу фамилию")
// 	fmt.Scan(&surname)
// 	fmt.Println("Ваша фамилия:", surname)
// 	fmt.Println("Ваше имя:", name)
// 	fmt.Println("Ваш возраст:", age)
// 	fmt.Println("Вы фронтендер", isFrontend)
// 	fmt.Println("Ваш город", city)

// 	if age < 18 {
// 		fmt.Println("Вы еще ребенок")
// 	} else {
// 		fmt.Println("Вы уже взрослый")
// 	}

// 	if isFrontend == true {
// 		fmt.Println("Ты програмист")
// 	}else {
// 		fmt.Println("Ты не програмист")

// 	}
// }

package main

import "fmt"

func main() {
	var name string
	fmt.Print("What's your name? ")
	fmt.Scan(&name)

	var age int
	fmt.Print("Enter your age ")
	fmt.Scan(&age)
	fmt.Println("Your name is:", name)
	fmt.Println("You are:", age, "years old")

	if age < 14 {
		fmt.Println("You are still a kid")
	} else if age > 14 && age < 18 {
		fmt.Println("You are teen")
	} else {
		fmt.Println("You are adult")
	}

	// Обычный цикл ничего не обычного
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Цикл как while или женская логика сначало делать потом думать
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Unlimited cycle && infiniy (не делать так ломает комп)
	// for{
	// 	print("Не делать так пожалуйста или комп взорвется просто знай про это закоментируй и забудь")
	// }

	
	// это массив
	var nubmers [3]int
	nubmers[0] = 1
	nubmers[1] = 2
	nubmers[2] = 3
	fmt.Println(nubmers)

	users := []string{"Sardor", "Danya", "Erkin", "Daler"}
	users = append(users, "Murodjon")
	fmt.Println(users)

	//Массив чисел это срез
	numbers := []int{1, 2, 3, 4, 5}

	// цикл + массив
	for i := 0; i < 5; i++ {
		fmt.Println("Number", numbers[i])
	}

	// цикл range
	for i, user := range users {
		fmt.Println("User", i, "is", user)
	}

	// Метод map
	people := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	// Цикл range
	for username, age := range people {
		fmt.Println(username, "is", age, "years old")
	}
	index()
	isAdult(age)
	seperate()
	// работа с map
	devs := map[string]int{
		"Sardor": 15,
		"Danya":  16,
		"Erkin":  16,
	}
	// работа с циклом и range
	for name, age := range devs {
		fmt.Println(name, "is", age, "years old")
	}

}

// типо функция которая принимает пропсы
func greet(name string) {
	fmt.Println("Hello ", name)
}

// функция котороя задает пропсы
func index() {
	var name string
	fmt.Print("What is Your name? ")
	fmt.Scan(&name)
	greet(name)
}

// длинная моя версия кода
// func isAdult(age int) bool {
// 	if age >= 18 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// короткая версия самого лучшего помощника в мире

func isAdult(age int) bool {
	return age >= 18
}

func seperate() {
	num := []int{1, 2, 3, 4}
	for n := 0; n < 4; n++ {
		fmt.Println(num[n], "sec")
	}
}
