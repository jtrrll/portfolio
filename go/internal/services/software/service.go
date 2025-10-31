package software

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type RepositorySummary struct {
	Name        string
	Description string
	Thumbnail   string
	Languages   map[string]int
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
				Languages:   languages,
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return summaries, nil
}
