package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strings"
)

const apiURL = "https://spoo.me/"

type ShortenRequest struct {
    URL string `json:"url"`
}

type ShortenResponse struct {
    ShortURL string `json:"short_url"`
}

func shortenURL(longURL string) (string, error) {
    encodedURL := url.QueryEscape(longURL)
    payload := strings.NewReader("url=" + encodedURL)

    req, err := http.NewRequest("POST", apiURL, payload)
    if err != nil {
        return "", err
    }

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Accept", "application/json")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return "", err
    }

    var response ShortenResponse
    if err := json.Unmarshal(body, &response); err != nil {
        return "", err
    }

    return response.ShortURL, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    var req ShortenRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    shortURL, err := shortenURL(req.URL)
    if err != nil {
        http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
}

func main() {
    http.HandleFunc("/shorten", handler)
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}
