package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetCOnfigFile("./config/env.json)

	if error := viper.ReadInConfig(); err!=nil {
		fmt.Println("Error config file!, &s", err)
	}

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, youve requested:%s\n", r.URL.Path)
	})
	http.ListenAndServe(viper.GetString("server.port"), nil)
}