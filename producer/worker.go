package producer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func Run() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6373"})
	defer client.Close()

	playerID := "P123"
	gameID := "G456"

	// 建立 Refund 任務
	task, err := NewRefundTask(playerID, gameID)
	if err != nil {
		log.Fatal(err)
	}

	// 延遲 300 秒
	info, err := client.Enqueue(task,
		asynq.ProcessIn(1*time.Second),
		asynq.TaskID(fmt.Sprintf("refund:%s:%s", playerID, gameID)), // 用 TaskID 保證唯一性
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

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
