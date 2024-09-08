package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/samsyntax/textio/internal/database"
	"github.com/samsyntax/textio/internal/handlers"
	"github.com/samsyntax/textio/internal/utils"
	v1 "github.com/samsyntax/textio/internal/v1"
)
func StartServer() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading environment variables:", err)
  }
  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT string not found in the environment")
  }
  dbURL := os.Getenv("DB_URL")
  if dbURL == "" {
    log.Fatal("DB connection string not found in the environment")
  }

  conn, err := sql.Open("postgres", dbURL)
  if err != nil {
    log.Fatal("Failed to connect to a database:", err)
  }
  db := database.New(conn)
  apiCfg := handlers.ApiConfig{
    DB: db,
  }
  routes := v1.InitRoutes(apiCfg)
  log.Printf("Server starting on port %v", portString)
   srv := &http.Server{
    Addr:":"+portString ,
    Handler: routes,

  }
  go utils.StartScraping(db, 10, time.Minute)

  defer conn.Close()

  srv.ListenAndServe()

  
}
