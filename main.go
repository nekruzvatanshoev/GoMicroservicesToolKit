package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//)
//
//type User struct {
//	ID string `json:"id"`
//	Name string `json:"name"`
//}
//
//var Users [] User = []User {
//	User{
//		ID:   "1",
//		Name: "Nekruz",
//	},
//	User{
//		ID:   "2",
//		Name: "Tom",
//	},
//	User {
//		ID:   "3",
//		Name: "Jack",
//	},
//}
//
//
//
//func main() {
//	router := http.NewServeMux()
//
//	router.HandleFunc("/users", retrieveAllUsers)
//	router.HandleFunc("/user", findUserByID)
//	if err := http.ListenAndServe(":8000", router); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Hello World!")
//}
//
//
//func retrieveAllUsers(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	err := json.NewEncoder(w).Encode(Users)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	return
//}
//
//type ID struct {
//	ID string `json:"id"`
//}
//
//func findUserByID(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	var id ID
//	err := json.NewDecoder(r.Body).Decode(&id)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//
//	for _, user := range Users {
//
//		if user.ID == id.ID {
//			err := json.NewEncoder(w).Encode(user)
//			if err != nil {
//				w.WriteHeader(http.StatusInternalServerError)
//			}
//			fmt.Println("W")
//			return
//		}
//	}
//
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound)
//		return
//	}
//
//	return
//}