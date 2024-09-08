package handlers

import (
	"fmt"
	"net/http"
	"github.com/samsyntax/textio/internal/auth"
	"github.com/samsyntax/textio/internal/database"
	"github.com/samsyntax/textio/internal/utils"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, 400, fmt.Sprintf("couldn't fetch the user %v", err))
			return
		}
    handler(w, r, user)
	}
}
