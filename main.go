package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//TODO Cache Access token

var MyPlaylistID = "03tfjQyEbooihXvHinrWqU"

type MyTokens struct {
	AccessToken string `json:"access_token"`
	AuthToken   string `json:"auth_token"`
	TokenType   string `json:"token_type"`
}

type Login_Info struct {
	CLIENT_ID      string
	CLIENT_SECRET  string
	CLIENT_REFRESH string
}

func AddSong(query string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Opening .env file")
	}

	My_Info := Login_Info{os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REFRESH_TOKEN")}

	GetClientFlowAccessToken(My_Info)               //Caching
	GetAccessToken(My_Info.CLIENT_REFRESH, My_Info) //Caching
	MySongID, _ := SpotifySearch(AccessToken, query, "TO_IMPLEMENT")
	AddToPlaylist(AccessToken.AuthToken, MySongID)

}

func GetClientFlowAccessToken(MyInfo Login_Info) error {
	if AccessToken != nil && Expiration != nil {
		if time.Now().Before(*Expiration) {
			fmt.Println("Client Access Token Already Exists")
			return nil
		}
	}

	url := "https://accounts.spotify.com/api/token"
	body := []byte("grant_type=client_credentials")
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", MyInfo.CLIENT_ID, MyInfo.CLIENT_SECRET)))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to get token: %s", body)
	}

	err = json.NewDecoder(resp.Body).Decode(&AccessToken)

	exp := time.Now().Add(time.Hour)
	Expiration = &exp
	if err != nil {
		log.Fatalf("failed to parse the json responnse: %v", err)
	}

	return nil
}

func SpotifySearch(AccessToken *MyTokens, SearchQuery, QueryType string) (string, error) {
	//Only works for songs, doesnt work with special characters like &
	//TODO Add artists, albums and more
	//TODO Fix Special Characters
	Queried := QueryFormatter(SearchQuery)
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track", Queried)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+AccessToken.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//print(string(body))

	type SongID struct {
		ID string `json:"id"`
	}

	type Tracks struct {
		Items []SongID `json:"items"`
	}

	type Response struct {
		Tracks Tracks `json:"tracks"`
	}

	var response Response
	firstID := ""
	json.Unmarshal([]byte(body), &response)
	if len(response.Tracks.Items) > 0 {
		firstID = response.Tracks.Items[0].ID
	} else {
		fmt.Println("Song Not Found")
	}

	return firstID, nil
}

func AddToPlaylist(AuthToken string, TRACK_ID string) error {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", MyPlaylistID)

	requestBody := map[string][]string{
		"uris": {fmt.Sprintf("spotify:track:%s", TRACK_ID)},
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error creating request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+AuthToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	fmt.Println(string(body))
	return nil

}

func GetAccessToken(RefreshToken string, MyInfo Login_Info) error {
	if AccessToken.AuthToken != "" && Expiration != nil {
		if time.Now().Before(*Expiration) {
			fmt.Println("Auth Token Already Exists")
			return nil
		}
	}

	Myurl := "https://accounts.spotify.com/api/token"

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", RefreshToken)

	req, err := http.NewRequest("POST", Myurl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.SetBasicAuth(MyInfo.CLIENT_ID, MyInfo.CLIENT_SECRET)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	type TokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	var tokenResp TokenResponse
	json.Unmarshal(bodyBytes, &tokenResp)

	AccessToken.AuthToken = tokenResp.AccessToken
	return nil
}

func QueryFormatter(Query string) string {
	NewString := strings.ReplaceAll(Query, " ", "+")
	return NewString
}

func GetPlaylistByID(AccessToken *MyTokens, PID string) error {
	MyPID := PID
	url := "https://api.spotify.com/v1/playlists/" + MyPID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("error creating request: %v", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+AccessToken.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil

}
