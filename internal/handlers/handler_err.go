
package internal

import (
	"net/http"

	internal "github.com/samsyntax/textio/internal/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request){
  internal.RespondWithError(w, 200, "Something went wrong")
}
