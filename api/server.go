package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	ListenPort	string
}

func NewAPIServer(listenPort string) *APIServer {
	return &APIServer {
		ListenPort:	listenPort,
	}
}

// Run permet d'exécuter le serveur http
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/auth/signup", makeHTTPHandlerFunc(s.handleSignup))
	router.HandleFunc("/auth/signin", makeHTTPHandlerFunc(s.handleSignin))

	// lancer le serveur
	if err := http.ListenAndServe(s.ListenPort, router); err != nil {
		log.Fatal("impossible de lancer le serveur api:", err)
	}
}

// APIFunc correspond au type de fonctions traitant les requêtes
// reçues via nos endpoints (handlers)
type APIFunc func(w http.ResponseWriter, r *http.Request) error

// makeHTTPHandlerFunc permet de traiter les erreurs renvoyées après le
// traitement de la requête http
func makeHTTPHandlerFunc(f APIFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			respondWithError(w, err)
		}
	}
}

// respondWithError permet de renvoyer une erreur au client en réponse à une requête http
func respondWithError(w http.ResponseWriter, err error) {
	errorStatus := http.StatusNotImplemented
	errorMessage := "error"
	
	respondWithJSON(w, errorStatus, errorMessage)
}

// respondWithJSON permet de renvoyer du contenu json en réponse à une requête http
func respondWithJSON(w http.ResponseWriter, status int, content any) error {
	// configurer le header de la réponse http
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(content)
}