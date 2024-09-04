package internal

import (
	"net/http"

	internal "github.com/samsyntax/textio/internal/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request){
  internal.RespondWithJSON(w, 200, struct{Key string}{Key: "Ok"})
}
