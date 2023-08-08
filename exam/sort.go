package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
)

type PageInfo struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       []Detail `json:"data"`
}

type Detail struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Author      string `json:"author"`
	NumComments int    `json:"num_comments"`
	StoryTitle  string `json:"story_title"`
}

func main() {
	limit := 2
	result := []*Detail{}
	r := []string{}
	index := 1

	current := 0
	total := limit
	for current < total {
		req := &http.Request{}
		req.URL = &url.URL{
			Scheme:   "https",
			Host:     "jsonmock.hackerrank.com",
			Path:     "/api/articles",
			RawQuery: fmt.Sprintf("page=%d", index),
		}

		rsp, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		data := make([]byte, rsp.ContentLength)
		_, e := rsp.Body.Read(data)
		if e != nil {
			fmt.Println(e)
		}

		page := &PageInfo{}
		e = json.Unmarshal(data, &page)
		if e != nil {
			fmt.Println(e)
		}
		total = page.Total
		current += page.PerPage

		for _, da := range page.Data {
			if len(da.Title) == 0 && len(da.StoryTitle) == 0 {
				continue
			}
			title := da.Title
			if len(title) == 0 {
				title = da.StoryTitle
			}
			result = append(result, &Detail{
				Title:       title,
				NumComments: da.NumComments,
			})
		}
		index++
	}
	sort.Slice(result, func(i, j int) bool {
		a := result[i]
		b := result[j]
		if a.NumComments > b.NumComments {
			return true
		} else if a.NumComments < b.NumComments {
			return false
		}

		for k := 0; k < len(a.Author) && k < len(b.Author); k++ {
			if a.Author[k] > b.Author[k] {
				return false
			} else if a.Author[k] < b.Author[k] {
				return true
			}
		}

		if len(a.Author) < len(b.Author) {
			return true
		}
		return true
	})
	for i := 0; i < limit; i++ {
		r = append(r, result[i].Title)
	}
	fmt.Println(r)
}
