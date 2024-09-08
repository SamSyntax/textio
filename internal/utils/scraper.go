package utils

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/samsyntax/textio/internal/database"
	"github.com/theritikchoure/logx"
)

func StartScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	logx.LogWithTimestamp(fmt.Sprintf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest), logx.FGCYAN, "")
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			logx.LogWithTimestamp(fmt.Sprintf("Error getting next feeds to fetch %v", err), logx.ERROR, "")
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MakrFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		logx.LogWithTimestamp(fmt.Sprintf("Error marking feed as fetched %v", err), logx.FGRED, "")
		return
	}

	rssFeed, err := UrlToFeed(feed.Url)
	if err != nil {
		logx.LogWithTimestamp(fmt.Sprintf("Error fetching feed %v", err), logx.FGRED, "")
	}

	for _, item := range rssFeed.Channel.Item {
		desc := sql.NullString{}
		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			logx.LogWithTimestamp(fmt.Sprintf("Error parsing publish date: %v", err), logx.FGRED, "")
		}
    if item.Title == "" {
			logx.LogWithTimestamp(fmt.Sprintf("Title of %v is empty, skipping", item.Link), logx.FGCYAN, "")
    }
    _, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: desc,
			PublishedAt: t,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
    if err != nil {
      if strings.Contains(err.Error(), "duplicate key") {
        continue
      }
			logx.LogWithTimestamp(fmt.Sprintf("Couldn't create a post: %v", err), logx.FGRED, "")
    }
	}

	logx.LogWithTimestamp(fmt.Sprintf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item)), logx.FGGREEN, "")

}
