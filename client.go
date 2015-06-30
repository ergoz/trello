package trello

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client interface {
	BoardService() BoardService
}

type BoardService interface {
	GetBoard(id string) (Board, error)
}

type Board interface {
	Name() string
	Lists() ([]List, error)
}

type List interface {
	Name() string
}

type client struct {
	key   string
	token string
}

type boardService struct {
	client *client
}

func NewClient(key, token string) Client {
	return &client{
		key:   key,
		token: token,
	}
}

func (c *client) BoardService() BoardService {
	return &boardService{
		client: c,
	}
}

const baseURL = "https://api.trello.com"

func (b *boardService) GetBoard(id string) (Board, error) {

	restURL := fmt.Sprintf("%s/1/boards/%s?key=%s", baseURL, id, b.client.key)
	if len(b.client.token) > 0 {
		restURL += fmt.Sprintf("&token=%s", b.client.token)
	}

	// TODO(ttacon)
	req, err := http.NewRequest(
		"GET",
		restURL,
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var d board
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	d.client = b.client

	return &d, nil
}

type board struct {
	client *client `json:"-"`

	ID             string                 `json:"id"`
	DescData       interface{}            `json:"descData"` // TODO(ttacon): identify the actual type
	Closed         bool                   `json:"closed"`
	IDOrganization interface{}            `json:"idOrganization"` // same as descData
	Pinned         bool                   `json:"pinned"`
	ShortURL       string                 `json:"shortUrl"`
	Desc           string                 `json:"desc"`
	BoardName      string                 `json:"name"`
	URL            string                 `json:"url"`
	Prefs          map[string]interface{} `json:"prefs"`      // TODO(ttacon): pull concrete struct out
	LabelNames     map[string]interface{} `json:"labelNames"` // same as prefs
}

func (b *board) Name() string {
	return b.BoardName
}

func (b *board) Lists() ([]List, error) {
	// TODO(ttacon)
	return nil, nil
}

type list struct {
	client *client `json:"-"`

	name string
}

func (l *list) Name() string {
	// TODO(ttacon)
	return ""
}
