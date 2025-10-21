package software

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type RepositorySummary struct {
	Name        string
	Description string
	Thumbnail   string
	Languages   map[string]float32
}

func GetAllRepositorySummaries(ctx context.Context) ([]RepositorySummary, error) {
	repos, err := ListRepositoriesForUser(ctx, "jtrrll")
	if err != nil {
		return nil, err
	}

	summaries := make([]RepositorySummary, len(repos))
	g, ctx := errgroup.WithContext(ctx)
	for i, repo := range repos {
		g.Go(func() error {
			thumbnail, err := GetThumbnailForRepository(ctx, repo.GetOwner().GetLogin(), repo.GetName())
			if err != nil {
				thumbnail = ""
			}

			summaries[i] = RepositorySummary{
				Name:        repo.GetName(),
				Description: repo.GetDescription(),
				Thumbnail:   thumbnail,
				Languages:   nil,
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return summaries, nil
}

type RepositoryDetails struct {
}

func GetRepositoryDetails() (*RepositoryDetails, error) {
	return nil, nil
}
