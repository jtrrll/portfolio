package software

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/google/go-github/v76/github"
	"golang.org/x/sync/errgroup"
)

type RepositorySummary struct {
	Name        string
	Description string
	Thumbnail   string
	Topics      []string
	Languages   map[string]int
}

var (
	cacheMu    sync.RWMutex
	cachedData []RepositorySummary
	cacheReady bool
	cacheCond  = sync.NewCond(&sync.Mutex{})
)

// GetAllRepositorySummaries returns a channel that delivers cached repository data.
// Blocks until data is available if the initial fetch has not yet completed.
func GetAllRepositorySummaries(ctx context.Context) <-chan []RepositorySummary {
	ch := make(chan []RepositorySummary, 1)
	go func() {
		defer close(ch)

		cacheCond.L.Lock()
		for !cacheReady {
			cacheCond.Wait()
		}
		cacheCond.L.Unlock()

		cacheMu.RLock()
		data := cachedData
		cacheMu.RUnlock()

		select {
		case ch <- data:
		case <-ctx.Done():
		}
	}()
	return ch
}

// StartBackgroundRefresh fetches GitHub data immediately and then on the given interval.
// Blocks until ctx is cancelled.
func StartBackgroundRefresh(ctx context.Context, interval time.Duration) {
	slog.InfoContext(ctx, "fetching initial software data")
	if err := fetchAndCache(ctx); err != nil {
		slog.ErrorContext(ctx, "initial software data fetch failed", "error", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			slog.InfoContext(ctx, "stopping software data refresh")
			return
		case <-ticker.C:
			slog.InfoContext(ctx, "refreshing software data")
			if err := fetchAndCache(ctx); err != nil {
				slog.ErrorContext(ctx, "software data refresh failed", "error", err)
			}
		}
	}
}

func fetchAndCache(ctx context.Context) error {
	allRepos, err := ListRepositoriesForUser(ctx, "jtrrll")
	if err != nil {
		return err
	}

	var repos []*github.Repository
	for _, repo := range allRepos {
		if !repo.GetFork() {
			repos = append(repos, repo)
		}
	}

	summaries := make([]RepositorySummary, len(repos))
	g, ctx := errgroup.WithContext(ctx)
	for i, repo := range repos {
		g.Go(func() error {
			thumbnail, err := GetThumbnailForRepository(ctx, repo.GetOwner().GetLogin(), repo.GetName())
			if err != nil {
				return err
			}

			languages, err := ListLanguagesForRepository(ctx, repo.GetOwner().GetLogin(), repo.GetName())
			if err != nil {
				return err
			}

			summaries[i] = RepositorySummary{
				Name:        repo.GetName(),
				Description: repo.GetDescription(),
				Thumbnail:   thumbnail,
				Topics:      repo.Topics,
				Languages:   languages,
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	cacheMu.Lock()
	cachedData = summaries
	cacheMu.Unlock()

	cacheCond.L.Lock()
	cacheReady = true
	cacheCond.Broadcast()
	cacheCond.L.Unlock()

	return nil
}
