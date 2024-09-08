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

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string    `json:"name"`
		Url  string    `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error parsing JSON %v", err))
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't create feed %v", err))
	}

	utils.RespondWithJSON(w, 201, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
  feeds, err := apiCfg.DB.GetFeeds(r.Context())
  if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't fetch feeds %v", err))
  }

  utils.RespondWithJSON(w, 200, models.DatabaseFeedsToFeeds(feeds))
}

func (apiCfg *ApiConfig) HandlerGetNextFetchedAtFeed(w http.ResponseWriter, r *http.Request) {

}
