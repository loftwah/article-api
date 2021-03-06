package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Article struct (Model)
// {
// 	"id": "1",
// 	"title": "latest science shows that potato chips are better for you than sugar",
// 	"date" : "2016-09-22",
// 	"body" : "some text, potentially containing simple markup about how potato chips are great",
// 	"tags" : ["health", "fitness", "science"]
//   }

type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

// Init articles var as a slice Article struct
var articles []Article

// Get all articles
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// Get single article
func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through articles and find one with the id from the params
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

// Add function to return the list of articles that have the matching {tagName] against the {date} and return the tag, the count of what is returned, an array of the article ids, and an array of the related tags
func getArticleByTagAndDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tagName = params["tagName"]
	var date = params["date"]
	var articlesByTagAndDate []Article
	var articlesByTag []Article
	var articlesByDate []Article
	var relatedTags []string
	var articleCount int
	var articleIds []string

	for _, item := range articles {
		for _, tag := range item.Tags {
			if tag == tagName {
				articlesByTag = append(articlesByTag, item)
			}
		}
	}
	for _, item := range articlesByTag {
		if item.Date == date {
			articlesByDate = append(articlesByDate, item)
		}
	}
	for _, item := range articlesByDate {
		articlesByTagAndDate = append(articlesByTagAndDate, item)
	}
	for _, item := range articlesByTagAndDate {
		for _, tag := range item.Tags {
			if tag != tagName {
				relatedTags = append(relatedTags, tag)
			}
		}
	}
	for _, item := range articlesByTagAndDate {
		articleIds = append(articleIds, item.ID)
	}

	articleCount = len(articleIds)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tag":          tagName,
		"count":        articleCount,
		"articles":     articleIds,
		"related_tags": relatedTags,
	})
}

// Add new article
func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

// Update article
func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			var article Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.ID = params["id"]
			articles = append(articles, article)
			json.NewEncoder(w).Encode(article)
			return
		}
	}
}

// Delete article
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database if this was going to be a real app
	articles = append(articles, Article{ID: "1", Title: "Article One", Date: "2021-09-17", Body: "This is the body of article one", Tags: []string{"tag1", "tag2", "tag3"}})
	articles = append(articles, Article{ID: "2", Title: "Article Two", Date: "2020-09-17", Body: "This is the body of article two", Tags: []string{"tag1", "tag2", "tag3"}})

	// Route handles & endpoints
	r.HandleFunc("/articles", getArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", getArticle).Methods("GET")                     // This meets requirement 2
	r.HandleFunc("/tags/{tagName}/{date}", getArticleByTagAndDate).Methods("GET") // This meets requirement 3
	r.HandleFunc("/articles", createArticle).Methods("POST")                      // This meets requirement 1
	r.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
