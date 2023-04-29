## DESCRIPTION

This is a simple project designed to study and improve skills in the GoLang language.

This project is a simple `Contact` Rest API

**End Points**

| Method | Endpoint | Description |
| --- | --- | --- |
| GET | /contacts | Gets all contacts |
| GET | /contacts/{id} | Get contact search by id |
| POST | /contacts | Create an contact from request body |
| DELETE | /contacts/{id} | Delete an contact by id | 

Request Exemple

```json
{
	"name": "Thomas",
	"email": "thomas.matheus60@gmail.com",
	"phone_number": "98888-8888"
}
```

 ----
## INSTALLATION

Clone this project.

```sh 
$ git clone git@github.com:Thomas-Matheus/contacts.git
```

---

Install dependencies

```sh 
$ go install
```

Compile project 

```sh 
$ go build
```

---

Goals

- Go report A+
- Code coverage 100%
- Makefile
- Docker

--- 

[GoLang Installation](https://golang.org/doc/install)

[GoLang Doc](https://golang.org/doc/)
