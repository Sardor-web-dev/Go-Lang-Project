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

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Thread struct {
	ID        int    `json:"id"`
	Avatar    string `json:"avatar"`
	User      string `json:"user"`
	Time      string `json:"time"`
	Text      string `json:"text"`
	LikeCount int    `json:"likeCount"`
	Comments  int    `json:"comments"`
	Resends   int    `json:"resends"`
	Views     int    `json:"views"`
}

var threads = []Thread{
	{	1, "/danya.jpg", "Danya", "5ч", "My name is Danya, I am a Frontend Developer", 200, 200, 200, 200},
	{2, "/sardor.jpg", "Sardor", "2ч", "My name is Sardor, I am a Frontend Developer", 500, 245, 464, 654},
	{3, "/avatar.jpg", "Erkin", "1ч", "Hello world, it's Erkin here! Writing some code 🚀", 312, 98, 112, 430},
	{4, "/avatar.jpg", "Murodjon", "3ч", "Building something cool with React and Next.js 🔥", 220, 150, 133, 512},
	{5, "/avatar.jpg", "Ramizjon", "6ч", "Just finished my new portfolio website 🎉", 610, 342, 278, 800},
	{6, "/avatar.jpg", "Daler", "8ч", "Working on some TypeScript magic ✨", 180, 76, 90, 300},
	{7, "/avatar.jpg", "Ozodbek", "9ч", "I love using shadcn/ui in my projects!", 400, 210, 150, 550},
	{8, "/avatar.jpg", "Konstantin", "10ч", "Frontend is fun when you understand the logic 🧠", 275, 115, 134, 460},
	{9, "/avatar.jpg", "Shaxriyor", "11ч", "CSS can be tricky, but Tailwind makes it easier 🧩", 490, 198, 205, 620},
	{10, "/avatar.jpg", "Alex", "12ч", "Sharing my open-source project soon, stay tuned 💡", 305, 150, 160, 580},
}

func threadsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(threads)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/api/threads", threadsHandler)
	log.Println("Server running on http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func Hello() {
	print("Hello, World!")
}
