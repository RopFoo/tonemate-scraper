package scrape

import (
	"drilbur/internal/sites"
	"drilbur/pkg/model"
)

func ScrapeSites(parameters model.Parameters) map[string][]model.Teaser {
	sites.Musicstore.SetParameters(parameters)
	sites.Musikersucht.SetParameters(parameters)

	musikersuchtTeasers := scrapeTeasers(&sites.Musikersucht)
	musicstoreTeasers := scrapeTeasers(&sites.Musicstore)

	return map[string][]model.Teaser{
		"musikersucht": musikersuchtTeasers,
		"musicstore":   musicstoreTeasers,
	}
}
