package software

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v76/github"
	"golang.org/x/net/html"
)

var client = github.NewClient(nil)

func GetThumbnailForRepository(ctx context.Context, owner string, repo string) (string, error) {
	url := fmt.Sprintf("https://github.com/%s/%s", owner, repo)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("fetching repo page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		switch tokenizer.Next() {
		case html.ErrorToken:
			if tokenizer.Err() != nil {
				return "", fmt.Errorf("parsing HTML: %w", tokenizer.Err())
			}
			return "", fmt.Errorf("og:image not found")
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data != "meta" {
				continue
			}

			var prop, content string
			for _, attr := range token.Attr {
				switch attr.Key {
				case "property":
					prop = attr.Val
				case "content":
					content = attr.Val
				}
			}

			if prop == "og:image" && content != "" {
				return content, nil
			}
		case html.EndTagToken:
			t := tokenizer.Token()
			if t.Data == "head" {
				return "", fmt.Errorf("og:image not found in <head>")
			}
		}
	}
}

func ListRepositoriesForUser(ctx context.Context, user string) ([]*github.Repository, error) {
	repos, _, err := client.Repositories.ListByUser(ctx, user, &github.RepositoryListByUserOptions{Sort: "pushed"})
	return repos, err
}

func ListLanguagesForRepository(ctx context.Context, owner string, repo string) (map[string]int, error) {
	if languages, _, err := client.Repositories.ListLanguages(ctx, owner, repo); err != nil {
		return nil, err
	} else {
		return languages, nil
	}
}
