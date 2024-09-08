package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/samsyntax/textio/internal/database"
	"github.com/samsyntax/textio/internal/models"
	"github.com/samsyntax/textio/internal/utils"
)

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error parsing JSON %v", err))
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't create user %v", err))
	}
	utils.RespondWithJSON(w, 201, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  20,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't get posts for user %v", err))
		return
	}
	utils.RespondWithJSON(w, 200, models.DatabasePostsToPosts(posts))
}
