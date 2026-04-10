package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/demo", demoHandler)

	log.Println("Khoi dong server tai port: 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Loi khoi dong server: ", err)
	}
}

func demoHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("%+v", res)

	if req.Method != http.MethodGet {
		http.Error(res, "Phuong thuc nay khong duoc chap nhan", http.StatusMethodNotAllowed)
		return 
	}

	response := map[string]string {
		"message" : "Chao mung ban",
		"author" : "dwchwang",
	}

	res.Header().Set("Content-Type", "application/json")

	// data, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(res, "Loi ma hoa Json: ", http.StatusInternalServerError)
	// 	return
	// }
	// res.Write(data)

	json.NewEncoder(res).Encode(response)
}
