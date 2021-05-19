package faceit

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type Player struct {
	ID            string `json:"player_id"`
	Avatar        string `json:"avatar"`
	Country       string `json:"country"`
	FaceitURL     string `json:"faceit_url"`
	Nickname      string `json:"nickname"`
	SteamNickname string `json:"steam_nickname"`
	NewSteamID    string `json:"new_steam_id"`
	SteamID       string `json:"steam_id_64"`
}

type PlayerQuery struct {
	Nickname     string
	Game         string
	GamePlayerID string
}

// GetPlayer Retrieve player details by filter
func (fc *API) GetPlayer(params PlayerQuery) (*Player, error) {
	var player Player

	req, err := fc.request("GET", "/players", nil)
	if err != nil {
		return nil, err
	}

	queries := req.URL.Query()

	queries.Add("nickname", params.Nickname)
	queries.Add("game", params.Game)
	queries.Add("game_player_id", params.GamePlayerID)

	req.URL.RawQuery = queries.Encode()

	resp, err := fc.doRequest(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	err = json.Unmarshal(body, &player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

// GetPlayerByID Retrieve player details by ID
func (fc *API) GetPlayerByID(uuid string) (*Player, error) {
	var player Player
	path := fmt.Sprintf("/players/%v", uuid)

	req, err := fc.request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := fc.doRequest(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	err = json.Unmarshal(body, &player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}
