package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	// "setup"

	// "encoding/json"
	"time"

	"github.com/gofrs/uuid/v5"

	// "encoding/json"
	"golang.org/x/crypto/bcrypt"
	// "io"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type UserData struct {
	Name string
}

type Post struct {
	Title   string
	Content string
}

func root(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "ui/templates/home.html")

	schemaPostGet := `SELECT title, content FROM posts`

	row, err := db.Query(schemaPostGet)
	if err != nil {
	}

	var post []Post

	for row.Next() {
		var title, content string
		row.Scan(&title, &content)
		post = append(post, Post{title, content})
	}
	fmt.Println(post[0])

	tmpl, err := template.ParseFiles("ui/templates/home.html")
	if err != nil {
		http.Error(w, "failed to update ui", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, post)
}

func ping(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userList := params["user"]

	var output bytes.Buffer

	user := "user"
	output.WriteString("Hello ")
	if len(userList) > 0 {
		user = userList[0]
	}
	output.WriteString(user)
	output.WriteString("!\n")

	_, err := w.Write(output.Bytes())
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "ui/templates/register.html")

	// Firstname := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confPassword := r.FormValue("confirmPassword")
	confPassError := r.FormValue("confirmPasswordError")

	// userList := params["name"]

	// fmt.Println(params)
	if confPassword == "" || username == "" || password == "" || email == "" {
		http.Error(w, "all fields are required", http.StatusBadRequest)
		return
	}

	if confPassword != password {
		http.Error(w, "passwords do not match", http.StatusBadRequest)
		// confPassError = "passwords do not match"
		return
	}
	// confPassError = r.FormValue("confirmPasswordError")

	schema := `
	INSERT INTO users (username, password_hash, email) VALUES (?, ?, ?)`

	user := "user"
	// name := "unknown"
	pass := "unknown"
	mail := "unknown"
	// var output bytes.Buffer

	// output.WriteString("Welcome ")

	// if len("unknown") > 0 {
	user = username
	pass = password
	mail = email
	// name = Firstname
	// }

	passByte := []byte(password)

	fmt.Println(passByte)

	hashedPassword, error := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)

	if error != nil {
		http.Error(w, "failed to hash the password", http.StatusInternalServerError)
		return
		// panic(error)
	}

	_, err := db.Exec(schema, username, string(hashedPassword), email)

	if err != nil {
		fmt.Println("DB error: ", err)
		http.Error(w, "failed to save data into databse username or email already exists", http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "data saved successfully into database", http.StatusOK)
	}

	// _, err := w.Write(output.Bytes())
	fmt.Println("this is the errror: ", confPassError)

	fmt.Fprintf(w, "Username: %s\nEmail: %s\nPassword: %s\n", user, mail, pass)
	// body, diode := io.ReadAll(r.Body)

	// var data UserData

	// err := json.Unmarshal([]byte(name), &data.Name)
	// diode := json.NewDecoder(r.Body).Decode(&data)

	// if err != nil {
	// 	fmt.Errorf("failed to get userdata %v", err)
	// 	return
	// }

	// fmt.Println(name)
	// fmt.Println(&db)
	// fmt.Println(data.Name)
}

func login(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "all fields must be filled", http.StatusBadRequest)
		return
	}

	schema := `
	SELECT id, email, password_hash FROM users WHERE email = ?
	`

	row := db.QueryRow(schema, email)

	// if row == "" {
	// 	http.Error(w, "user not found", http.StatusNotFound)
	// 	return
	// }

	// if err != nil {
	// 	http.Error(w, "user not found", http.StatusNotFound)
	// 	return
	// }
	var dbId, dbemail, dbpassword string

	row.Scan(&dbId, &dbemail, &dbpassword)

	err := bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(password))
	if err != nil {
		http.Error(w, "user unknown try again", http.StatusForbidden)
		return
	}
	fmt.Println(dbpassword)

	u4, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate uuid: %v", err)
	}

	expiry := time.Now().Add(24 * time.Hour)

	// expiry := Add(24 * time.Hour)

	fmt.Println(expiry)
	// fmt.Println(expiry)

	schemaSession := `INSERT INTO sessions(id, user_id, expires_at) VALUES (?, ?, ?)`

	// if dbpassword != password {
	// 	http.Error(w, "user unknown try again", http.StatusForbidden)
	// 	return
	// }
	_, err = db.Exec(schemaSession, u4, dbId, expiry)

	uid := u4.String()

	setUpCookie(w, uid, expiry)

	fmt.Println(dbemail, user)

	// for row.Next() {
	// 	// var username, password string
	// 	row.Scan(&username, &password)

	fmt.Fprintf(w, "Welcome back %v", dbemail)
	// }
}

func setUpCookie(w http.ResponseWriter, uid string, exp time.Time) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    uid,
		Expires:  exp,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	schema := `SELECT title, content FROM posts`

	row, err := db.Query(schema)
	if err != nil {
		http.Error(w, "failed to retrieve data from db", http.StatusConflict)
		return
	}

	for row.Next() {
		var title, content string
		row.Scan(&title, &content)

		fmt.Fprintf(w, "title: %v, content: %v\n", title, content)
	}
}

func handleLoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/templates/login.html")
}

func handleRegisterHtml(w http.ResponseWriter, r *http.Request) {
	// fs := http.FileServer(http.Dir("./ui/templates/register.html"))
	http.ServeFile(w, r, "ui/templates/signup.html")
	// http.Handle("registering", fs)
	// fs.ServeHTTP(w,r)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/ui/templates/home.html")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	schema := `
	SELECT username, password_hash, email FROM users`
	row, err := db.Query(schema)
	if err != nil {
		http.Error(w, "failed to retrieve data from the database", http.StatusInternalServerError)
		return
	}
	// defer db.Close()

	for row.Next() {
		var username, password, email string
		row.Scan(&username, &password, &email)

		fmt.Fprintf(w, "username: %v, password: %v, email: %v\n", username, password, email)
	}
}

func handlePostPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/templates/post.html")
}

func sendPost(w http.ResponseWriter, r *http.Request) {
	postTitle := r.FormValue("postitle")
	postContent := r.FormValue("postContent")
	// user_id := 1

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "not authenticated", http.StatusUnauthorized)
		return
	}

	schemaCookie := `SELECT id, user_id, expires_at FROM sessions WHERE id = ?`

	row := db.QueryRow(schemaCookie, cookie.Value)

	var dbId string

	var dbuser_id int

	var dbexpiry time.Time

	err = row.Scan(&dbId, &dbuser_id, &dbexpiry)
	if err != nil {
		http.Error(w, "failed to get user credents", http.StatusNotFound)
	}

	if dbexpiry.Before(time.Now()) {
		http.Error(w, "cookie expired", http.StatusUnauthorized)
		return
	}

	user_id := dbuser_id

	fmt.Println("post sent successfully")
	if postTitle == "" {
		http.Error(w, "post title is required", http.StatusBadRequest)
		return
	}
	if postContent == "" {
		http.Error(w, "post contentn is required", http.StatusBadRequest)
		return
	}

	schema := `INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)`

	_, err = db.Exec(schema, postTitle, postContent, user_id)

	if err != nil {
		fmt.Printf("failed to add post into the database: %v", err)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// mux := http.NewServeMux()

	// mux.HandleFunc("/{$}", root)

	// handleHomePage(w, r)
}

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/{$}", root)
	mux.HandleFunc("/{$}", root)
	mux.HandleFunc("/sendpost", sendPost)
	mux.HandleFunc("/registering", handleRegisterHtml)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/log", handleLoginPage)
	mux.HandleFunc("/getusers", getUsers)
	mux.HandleFunc("/getposts", getPosts)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/post", handlePostPage)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))
	fmt.Println("server running on port 8080")

	setup()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
