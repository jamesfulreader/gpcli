package tsdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://www.thesportsdb.com/api/v1/json"

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 15 * time.Second},
		apiKey:     apiKey,
	}
}

func (c *Client) getJSON(ctx context.Context, path string, out any) error {
	url := fmt.Sprintf("%s/%s/%s", baseURL, c.apiKey, path)

	fmt.Fprintln(os.Stderr, "DEBUG: GET", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http do: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(res.Body, 4096))
		return fmt.Errorf("status %d: %s", res.StatusCode, string(body))
	}
	dec := json.NewDecoder(res.Body)

	if err := dec.Decode(out); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}
	return nil
}

func (c *Client) GetLeague(ctx context.Context, leagueID int) (*LeagueResponse, error) {
	var leagueRes LeagueResponse
	path := fmt.Sprintf("lookupleague.php?id=%d", leagueID)
	if err := c.getJSON(ctx, path, &leagueRes); err != nil {
		return nil, err
	}
	return &leagueRes, nil
}

func (c *Client) ListSeasons(ctx context.Context, leagueID int) (*SeasonsResponse, error) {
	var listSeasonsRes SeasonsResponse
	path := fmt.Sprintf("search_all_seasons.php?id=%d", leagueID)
	if err := c.getJSON(ctx, path, &listSeasonsRes); err != nil {
		return nil, fmt.Errorf("error fetching seasons: %w", err)
	}
	return &listSeasonsRes, nil
}

func (c *Client) EventsBySeason(ctx context.Context, leagueID int, season string) (*EventsResponse, error) {
	var er EventsResponse
	path := fmt.Sprintf("eventsseason.php?id=%d&s=%s", leagueID, season)
	if err := c.getJSON(ctx, path, &er); err != nil {
		return nil, err
	}
	return &er, nil
}

func (c *Client) EventsNext(ctx context.Context, leagueID int, season string) (*EventsResponse, error) {
	var nextEventsRes EventsResponse
	path := fmt.Sprintf("eventsseason.php?id=%d&s=%s", leagueID, season)
	if err := c.getJSON(ctx, path, &nextEventsRes); err != nil {
		return nil, err
	}
	return &nextEventsRes, nil
}

func (c *Client) TeamsByLeague(ctx context.Context, leagueID int) (*TeamsResponse, error) {
	var teamsRes TeamsResponse
	path := fmt.Sprintf("lookup_all_teams.php?id=%d", leagueID)
	if err := c.getJSON(ctx, path, &teamsRes); err != nil {
		return nil, err
	}
	return &teamsRes, nil
}

func (c *Client) EventByID(ctx context.Context, eventID string) (*EventsResponse, error) {
	var er EventsResponse
	path := fmt.Sprintf("lookupevent.php?id=%s", eventID)
	if err := c.getJSON(ctx, path, &er); err != nil {
		return nil, err
	}
	return &er, nil
}

func (c *Client) GetEvent(ctx context.Context, eventID string) (*EventsResponse, error) {
	var er EventsResponse
	path := fmt.Sprintf("lookupevent.php?id=%s", eventID)
	if err := c.getJSON(ctx, path, &er); err != nil {
		return nil, err
	}
	return &er, nil
}
