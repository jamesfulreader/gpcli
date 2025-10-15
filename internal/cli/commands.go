package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/jamesfulreader/gpcli/internal/tsdb"
)

type AppOptions struct {
	APIKey string
	League string
	JSON   bool
}

type App struct {
	Options AppOptions
	Out     io.Writer
	Client  *tsdb.Client
}

func (a *App) ResolveLeagueID() (int, error) {
	switch a.Options.League {
	case "f1":
		return 4370, nil
	case "f3":
		return 4487, nil
	default:
		// try parsing numeric
		id, err := strconv.Atoi(a.Options.League)
		if err != nil {
			return 0, fmt.Errorf("invalid league %q: must be f1, f3, or a number", a.Options.League)
		}
		return id, nil
	}
}

// Below are stubs the Cobra subcommands will call.
// In Step 2, wire these to tsdb client and real formatting.

func ShowLeague(ctx context.Context, app *App, leagueID int) error {
	leagueRes, err := app.Client.GetLeague(ctx, leagueID)
	if err != nil {
		return fmt.Errorf("error fetching league: %w", err)
	}
	if len(leagueRes.Leagues) == 0 {
		return fmt.Errorf("no league found with ID %d", leagueID)
	}
	league := leagueRes.Leagues[0]
	fmt.Fprintf(app.Out, "%s (Sport: %s, Country: %s, id=%s)\n",
		league.StrLeague, league.StrSport, league.StrCountry, league.IDLeague)
	return nil
}

func ListSeasons(ctx context.Context, app *App, leagueID int) error {
	seasonsRes, err := app.Client.ListSeasons(ctx, leagueID)
	if err != nil {
		return fmt.Errorf("error fetching seasons: %w", err)
	}
	if len(seasonsRes.Seasons) == 0 {
		return fmt.Errorf("no seasons found for league ID %d", leagueID)
	}
	for _, s := range seasonsRes.Seasons {
		fmt.Fprintln(app.Out, s.StrSeason)
	}
	return nil
}

func ShowSchedule(
	ctx context.Context,
	app *App,
	leagueID int,
	season string,
) error {
	scheduleRes, err := app.Client.EventsBySeason(ctx, leagueID, season)
	if err != nil {
		return fmt.Errorf("error fetching schedule: %w", err)
	}
	for _, e := range scheduleRes.Events {
		fmt.Fprintf(app.Out,
			"Round %s | %s | %s %s | %s, %s\n",
			e.IntRound, e.StrEvent, e.DateEvent, e.StrTime,
			e.StrVenue, e.StrCountry,
		)
	}
	return nil
}

func ShowNextEvents(ctx context.Context, app *App, leagueID int) error {
	er, err := app.Client.EventsNext(ctx, leagueID, "2025")
	if err != nil {
		return fmt.Errorf("error fetching next events: %w", err)
	}
	for _, e := range er.Events {
		fmt.Fprintf(app.Out, "%s | %s %s | %s\n",
			e.StrEvent, e.DateEvent, e.StrTime, e.StrVenue)
	}
	return nil
}

func ListTeams(ctx context.Context, app *App, leagueID int) error {
	tr, err := app.Client.TeamsByLeague(ctx, leagueID)
	if err != nil {
		return err
	}
	for _, t := range tr.Teams {
		fmt.Fprintf(app.Out, "%s (Stadium: %s, Country: %s)\n",
			t.StrTeam, t.StrStadium, t.StrCountry)
	}
	return nil
}

func ShowEvent(ctx context.Context, app *App, eventID string) error {
	res, err := app.Client.GetEvent(ctx, eventID)
	if err != nil {
		return fmt.Errorf("error fetching event: %w", err)
	}
	if len(res.Events) == 0 {
		fmt.Fprintln(app.Out, "no event found")
		return nil
	}
	e := res.Events[0]

	if app.Options.JSON {
		data, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			return fmt.Errorf("json marshal: %w", err)
		}
		_, _ = app.Out.Write(append(data, '\n'))
		return nil
	}

	fmt.Fprintf(app.Out, "Event:  %s\nDate:   %s %s\nVenue:  %s, %s\nStatus: %s\n\n%s\n",
		e.StrEvent, e.DateEvent, e.StrTime,
		e.StrVenue, e.StrCountry,
		e.StrStatus,
		e.StrDescriptionEN,
	)
	return nil
}
