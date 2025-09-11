package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func Run() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6373"},
		asynq.Config{
			Concurrency: 10, // worker 數量
			//TaskCheckInterval: 500 * time.Millisecond,
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeRefund, HandleRefundTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func HandleRefundTask(ctx context.Context, t *asynq.Task) error {
	var p RefundPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	// TODO: 在這裡實作真正的 Refund 邏輯
	log.Printf("收到消息 - 送出時間: [%v] - 收到時間: [%v] - 延遲秒數: [%v]", p.SendTime, time.Now().Format("2006-01-02 15:04:05"), p.DelaySec)
	return nil
}

const TypeRefund = "player:refund"

type RefundPayload struct {
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
	SendTime string `json:"send_time"`
	DelaySec int    `json:"delay_sec"`
}
