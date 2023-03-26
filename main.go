package main

import (
    "fmt"
    "net/smtp"
	"net/http"
	"html/template"
	"log"
)

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/create", create)
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal(err)
	}
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
    smtpUsername := ""  // Google e-mail Отправителя
    smtpPassword := "" //  Пароль из Google пароль и приложения

    to := []string{""}// Каму написать
    from := ""        // От кого письмо
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
	fmt.Println("Письмо успешно отправлено!")

	http.Redirect(w, r, "/", http.StatusFound)
}
