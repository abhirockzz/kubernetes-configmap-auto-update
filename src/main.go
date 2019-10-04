package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const CONFIG_DIR string = "/config/"

var appconfig map[string]string

func init() {
	appconfig = make(map[string]string)
	load()
}

func main() {

	http.HandleFunc("/readconfig/", read)
	http.HandleFunc("/reload/", reload)
	http.ListenAndServe(":8080", nil)
}

func load() {
	fmt.Println("loading config from file")

	files, err := ioutil.ReadDir(CONFIG_DIR)
	if err != nil {
		fmt.Println("cannot read dir "+CONFIG_DIR, err)
		return
	}
	for _, file := range files {
		key := file.Name()
		filename := CONFIG_DIR + file.Name()

		value, _ := ioutil.ReadFile(filename)
		if string(value) == "" {
			fmt.Println("Unable to read config value from", CONFIG_DIR+file.Name())
			continue
		}

		appconfig[key] = string(value)
	}
}

func read(rw http.ResponseWriter, req *http.Request) {
	key := strings.TrimPrefix(req.URL.Path, "/readconfig/") //http://host:8080/read/foo (foo is the key)
	value, there := appconfig[key]
	if !there {
		rw.Write([]byte("Configuration '" + key + "' does not exist"))
		return
	}
	fmt.Println("Got Value for config key " + key + " from map - " + value)
	rw.Write([]byte(value))
}

func reload(w http.ResponseWriter, r *http.Request) {
	load()
}
