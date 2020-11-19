package routes

import (
	"html/template"
	"net/http"
)

type Param struct {
	Token  string
	Status string
	Tech   string
}

//Status web static para mostrar el estado
func Status(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query()["token"][0]
	status := r.URL.Query()["status"][0]
	tech := r.URL.Query()["tech"][0]

	if status =="1"{
		status = "CREADA"
	}else{
		status = "FINALIZADA"

	}

	P := Param{Token: token, Status: status, Tech: tech}

	//	td := Todo{"Test templates", "Let's test a template to see the magic."}

	t, _ := template.ParseGlob("./templates/*.html")

	t.ExecuteTemplate(w, "status.html", P)

}
