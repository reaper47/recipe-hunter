package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/reaper47/recipya/internal/models"
)

func scrapeKennyMcGovern(root *goquery.Document) (models.RecipeSchema, error) {
	rs, err := parseGraph(root)
	if err != nil {
		return rs, err
	}
	rs.Image.Value, _ = root.Find(".gridfeel-post-thumbnail-single-img").Attr("src")
	return rs, nil
}
