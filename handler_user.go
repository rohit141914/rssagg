
package main

import(
	"net/http"
	"fmt"
	"time"
	"github.com/google/uuid"
	"encoding/json"
	"github.com/rohit141914/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed parsing JSON: ", err))	
	return 
	}
	user,err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreateAt: time.Now().UTC(),
		UpdateAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err!=nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: ", err))
		return 
	}
		respondWithJSON(w,200,user)
}