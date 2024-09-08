package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/samsyntax/textio/internal/database"
	"github.com/samsyntax/textio/internal/models"
	"github.com/samsyntax/textio/internal/utils"
)

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error parsing JSON %v", err))
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't follow feed %v", err))
	}

	utils.RespondWithJSON(w, 201, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *ApiConfig) HandlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't fetch followed feeds %v", err))
	}

	utils.RespondWithJSON(w, 200, models.DatabaseFeedFollowsToFeedFollows(feeds))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
  feedID := chi.URLParam(r, "feedFollowID")
  parsedID, err := uuid.Parse(feedID)
  if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't parse feed id %v", err))
  }

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
    ID: parsedID,
		UserID: user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("couldn't unfollow feed %v", err))
    return
	}

	utils.RespondWithJSON(w, 200, struct{}{})

}
