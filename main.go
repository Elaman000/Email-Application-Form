package main

import (
    "fmt"
    "net/smtp"
	"net/http"
	"html/template"
	"log"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
	r.HandleFunc("/", viewHandler)
	r.HandleFunc("/create", create)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = html.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("user-name")
	userNumber := r.FormValue("user-number")

	smtpHost := "smtp.gmail.com"
    smtpPort := 587
    smtpUsername := "kylychbekuuluelaman2002@gmail.com"
    smtpPassword := "iiboqtigqrqaocmk"

    to := []string{"kylychbekuuluelaman2003@gmail.com"}
    from := "kylychbekuuluelaman2002@gmail.com"
    subject := "Эй, я просто проверяю тебя."
    body := "Имя пользователя: "+userName +"\n"+"Номер телефона: "+userNumber

    header := make(map[string]string)
    header["From"] = from
    header["To"] = to[0]
    header["Subject"] = subject

    message := ""
    for k, v := range header {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
    err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, from, to, []byte(message))
    if err != nil {
        fmt.Println("Ошибка отправки электронной почты:", err)
    }
	fmt.Println("письмо успешно отправлено!")

	http.Redirect(w, r, "/", http.StatusFound)
}
