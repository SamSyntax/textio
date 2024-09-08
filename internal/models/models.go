package models

import (
	"time"
	"github.com/google/uuid"
	"github.com/samsyntax/textio/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}
func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(feed))
	}
	return feeds
}

func DatabaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
  return FeedFollow{
    ID: dbFeedFollow.ID,
    CreatedAt: dbFeedFollow.CreatedAt,
    UpdatedAt: dbFeedFollow.UpdatedAt,
    UserID: dbFeedFollow.UserID,
    FeedID: dbFeedFollow.FeedID,
  }
}

func DatabaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
  feedFollows := []FeedFollow{}
  for _, dbFeedFollow := range dbFeedFollows {
    feedFollows = append(feedFollows, DatabaseFeedFollowToFeedFollow(dbFeedFollow))
  }
  return feedFollows
}

type Post struct {
  ID uuid.UUID `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Title string `json:"title"`
  Description string `json:"description"`
  PublishedAt time.Time `json:"published_at"`
  Url string `json:"url"`
  FeedID uuid.UUID `json:"feed_id"`
}

func DatabasePostToPost(dbPost database.Post) Post {
  return Post{
    ID: dbPost.ID,
    CreatedAt: dbPost.CreatedAt,
    UpdatedAt: dbPost.UpdatedAt,
    Title: dbPost.Title,
    Description: dbPost.Description.String, 
    PublishedAt: dbPost.PublishedAt,
    Url: dbPost.Url,
    FeedID: dbPost.FeedID,
  }
}

func DatabasePostsToPosts(dbPosts []database.Post) []Post {
  posts := []Post{}

  for _, post := range dbPosts {
    posts = append(posts, DatabasePostToPost(post))
  }

  return posts
}
