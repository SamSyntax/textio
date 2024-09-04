package sever 

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	v1 "github.com/samsyntax/textio/internal/v1"
)
func StartServer() {
  routes := v1.InitRoutes()
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading environment variables:", err)
  }
  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT string not found in the environment")
  }
  log.Printf("Server starting on port %v", portString)
   srv := &http.Server{
    Addr:":"+portString ,
    Handler: routes,

  }

  srv.ListenAndServe()

  
}
