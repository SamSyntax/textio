package handlers 


import (
	"net/http"

	"github.com/samsyntax/textio/internal/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request){
  utils.RespondWithJSON(w, 200, struct{Key string}{Key: "Ok"})
}
