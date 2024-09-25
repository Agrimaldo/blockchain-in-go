package models

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type Block struct {
	bun.BaseModel `bun:"table:tb_block,alias:b"`
	Id            int64 `bun:"Id,pk,autoincrement"`
	Height        int32
	PreviousHash  string
	Hash          string
	CreatedAt     time.Time     `bun:"createdAt,nullzero,notnull,default:current_timestamp"`
	Transactions  []Transaction `bun:"rel:has-many,join:Id=BlockId"`
}

type Transaction struct {
	bun.BaseModel `bun:"table:tb_transaction,alias:t"`
	Id            int64 `bun:"Id,pk,autoincrement"`
	Sender        string
	Receiver      string
	Amount        float64
	Fee           float64
	Message       string
	BlockId       sql.NullInt64
	CreatedAt     time.Time `bun:"createdAt,nullzero,notnull,default:current_timestamp"`
	Block         *Block    `bun:"rel:has-one,join:BlockId=Id"`
}
