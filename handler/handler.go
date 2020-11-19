package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DavidNazareno/h_prueba/config"
	"github.com/DavidNazareno/h_prueba/graph"
	"github.com/DavidNazareno/h_prueba/graph/generated"
	"github.com/DavidNazareno/h_prueba/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//HandlerController controlador de rutas 
func HandlerController() {

	port := os.Getenv("PORT")
	if port == "" {
		port = config.PORT
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router := mux.NewRouter()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/status", routes.Status).Methods(("GET"))
	router.HandleFunc("/update", routes.AddScore).Methods(("POST"))

	//router.HandleFunc("/orders")
	handler := cors.AllowAll().Handler(router)
	log.Println("Server started in localhost:" + port)

	log.Fatalln(http.ListenAndServe(":"+port, handler))

}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
