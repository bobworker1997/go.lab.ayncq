package main

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeRefund = "player:refund"

type RefundPayload struct {
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
}

func NewRefundTask(playerID, gameID string) (*asynq.Task, error) {
	payload, err := json.Marshal(RefundPayload{PlayerID: playerID, GameID: gameID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeRefund, payload), nil
}
