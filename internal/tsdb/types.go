package tsdb

type LeagueResponse struct {
	Leagues []League `json:"leagues"`
}

type League struct {
	IDLeague   string `json:"idLeague"`
	StrLeague  string `json:"strLeague"`
	StrSport   string `json:"strSport"`
	StrCountry string `json:"strCountry"`
}

type SeasonsResponse struct {
	Seasons []Season `json:"seasons"`
}

type Season struct {
	StrSeason string `json:"strSeason"`
}

type EventsResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	IDEvent          string  `json:"idEvent"`
	StrEvent         string  `json:"strEvent"`
	StrSeason        string  `json:"strSeason"`
	IntRound         string  `json:"intRound"`
	DateEvent        string  `json:"dateEvent"`
	StrTime          string  `json:"strTime"`
	StrVenue         string  `json:"strVenue"`
	StrCountry       string  `json:"strCountry"`
	StrThumb         *string `json:"strThumb"`
	StrStatus        string  `json:"strStatus"`
	StrDescriptionEN string  `json:"strDescriptionEN"`
}

type TeamsResponse struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	IDTeam     string `json:"idTeam"`
	StrTeam    string `json:"strTeam"`
	StrStadium string `json:"strStadium"`
	StrCountry string `json:"strCountry"`
}
