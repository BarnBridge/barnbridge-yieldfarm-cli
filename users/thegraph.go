package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var log = logrus.WithField("module", "users")

func fetchActions() []Action {
	query := `{
		  stakingActions(first: 1000, skip: %d){
			type
			user
			token
			amount
			blockTimestamp
			txHash
		  }
		}`

	skip := 0
	var actions []Action

	req := gorequest.New()

	for {
		log.Info("fetching batch of 1000 actions from thegraph")
		start := time.Now()

		resp, body, errs := req.Post(viper.GetString("graph-url")).Send(map[string]string{
			"query": fmt.Sprintf(query, skip),
		}).End()

		if len(errs) > 0 {
			panic(errs)
		}

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Expected status 200, but got %d", resp.StatusCode)
		}

		type Response struct {
			Data struct {
				StakingActions []Action
			}
		}

		var gqlResp Response

		err := json.Unmarshal([]byte(body), &gqlResp)
		if err != nil {
			log.Fatal(err)
		}

		actions = append(actions, gqlResp.Data.StakingActions...)

		log.WithField("duration", time.Since(start)).Info("done fetching")

		if len(gqlResp.Data.StakingActions) < 1000 {
			break
		}
		skip += 1000
	}

	return actions
}

func (u *Users) getStableActions(actions []Action) []Action {
	var stableActions []Action

	for _, a := range actions {
		switch u.config.Tokens[strings.ToLower(a.Token)].ID {
		case "dai", "susd", "usdc":
			stableActions = append(stableActions, a)
		}
	}

	return stableActions
}

func (u *Users) getUniLPActions(actions []Action) []Action {
	var stableActions []Action

	for _, a := range actions {
		switch u.config.Tokens[strings.ToLower(a.Token)].ID {
		case "unilp":
			stableActions = append(stableActions, a)
		}
	}

	return stableActions
}
