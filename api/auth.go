package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthRequest struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type AuthResponse struct {
	JWT	string	`json:"jwt"`
}

// handleSignup permet d'inscrire un utilisateur dans notre application
func (s *APIServer) handleSignup(w http.ResponseWriter, r *http.Request) error {
	credentials, err := readAuthRequestCredentials(r)
	if err != nil {
		return err
	}

	fmt.Printf("demande d'inscription:\nusername: %s\npassword: %s\n", credentials.Username, credentials.Password)
	
	return respondWithJWT(w, "random-jwt-for-now")
}

// handleSignin permet de connecter un utilisateur déjà inscrit dans notre application
func (s *APIServer) handleSignin(w http.ResponseWriter, r *http.Request) error {
	credentials, err := readAuthRequestCredentials(r)
	if err != nil {
		return err
	}

	fmt.Printf("demande de connexion:\nusername: %s\npassword: %s\n", credentials.Username, credentials.Password)
	
	return respondWithJWT(w, "ramdom-jwt-for-now")
}

// readAuthRequestCredentials permet de récupérer les credentials d'une requête
// d'authentification
func readAuthRequestCredentials(r *http.Request) (*AuthRequest, error) {
	credentials := new(AuthRequest)
	
	if err := json.NewDecoder(r.Body).Decode(credentials); err != nil {
		return nil, err
	}

	return credentials, nil
}

// respondWithJWT permet d'envoyer un JWT à l'utilisateur en cas
// d'authentification réussie
func respondWithJWT(w http.ResponseWriter, jwt string) error {
	authResponse := AuthResponse{
		JWT:	jwt,
	}

	return respondWithJSON(w, http.StatusCreated, authResponse)
}

// TODO: fonction qui vérifie la méthode de la requête http dans les handlers