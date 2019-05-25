package search

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
)

type Search struct {
	cli   *elastic.Client
	index string
}

func NewES(url, index string) *Search {
	cli, err := elastic.Dial(elastic.SetURL(url))
	if err != nil {
		panic("init elasticsearch  error : " + err.Error())
	}
	return &Search{
		cli:   cli,
		index: index,
	}
}

func (s *Search) AddDoc(typ, id string, body interface{}) error {
	_, err := s.cli.Index().Index(s.index).Type(typ).Id(id).BodyJson(body).Do(context.Background())
	return err
}

func (s *Search) DeleteDoc(typ, id string) error {
	_, err := s.cli.Delete().Index(s.index).Type(typ).Id(id).Do(context.Background())
	return err
}

func (s *Search) UpdateDoc(typ, id string, body interface{}) error {
	_, err := s.cli.Update().Index(s.index).Type(typ).Id(id).Do(context.Background())
	return err
}

func (s *Search) Search(queryName string, queryValue, result interface{}) error {
	query := elastic.NewMatchPhraseQuery(queryName, queryValue)
	searchResults, err := s.cli.Search().Index(s.index).Query(query).Sort(queryName, true).Pretty(true).Do(context.Background())
	if err != nil {
		return err
	}
	return resolveSearchResults(searchResults, result)
}

func (s *Search) MultiSearch(value, result interface{}, names ...string) error {
	query := elastic.NewMultiMatchQuery(value, names...).Type("phrase")
	searchResults, err := s.cli.Search().Index(s.index).Query(query).Pretty(true).Do(context.Background())
	if err != nil {
		return err
	}
	return resolveSearchResults(searchResults, result)
}

func (s *Search) SearchScope(queryName string, queryValue, result interface{}, from, size int) error {
	query := elastic.NewMatchPhraseQuery(queryName, queryValue)
	searchResults, err := s.cli.Search().Index(s.index).Query(query).Sort(queryName, true).From(from).Size(size).Pretty(true).Do(context.Background())
	if err != nil {
		return err
	}
	return resolveSearchResults(searchResults, result)
}

func resolveSearchResults(results *elastic.SearchResult, value interface{}) error {
	if results.Hits.TotalHits > 0 {
		for _, hit := range results.Hits.Hits {
			err := json.Unmarshal(*hit.Source, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
