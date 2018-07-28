package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

// curl -id

func headers(w http.ResponseWriter, r *http.Request) {

	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h)

}

func bodyRecv(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	if _, err := r.Body.Read(body); err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))
	fmt.Fprintf(w, string(body))

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	pass := r.PostFormValue("pass")
	fmt.Println(name, pass)
	fmt.Fprintf(w, name+pass)

}

func fileRecv(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	file, _, err := r.FormFile("upload")
	if err != nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
			return
		}
	}

	fmt.Fprintln(w, err.Error())

}

func jsonRecv(w http.ResponseWriter, r *http.Request) {
	data := make([]byte, 1024)
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	s := struct {
		name string
		pass string
		age  int `json:"age,string"`
	}{}

	if err = json.Unmarshal(data, &s); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Println(s)

	fmt.Fprintln(w, string(data))
}

func writeHtml(w http.ResponseWriter, r *http.Request) {
	s := struct {
		Name string
		Pass string
		Age  int `json:"age,string"`
	}{

		Name:"kk",
		Pass:"123",
		Age:22,
	}


	js,err:=json.Marshal(&s)






	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(js)
}


func writeJson(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("a.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(data)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/redirect")
	w.WriteHeader(http.StatusFound)

}

func status(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusNotImplemented)
	w.Write()
}
func main() {

	http.HandleFunc("/header", headers)
	http.HandleFunc("/body", bodyRecv)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/upload", fileRecv)
	http.HandleFunc("json", jsonRecv)

	http.HandleFunc("/redirect",redirect)
	log.Fatal(http.ListenAndServe(":7000", nil))

}

//curl -i -H "Content-type:application/json"
