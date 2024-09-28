package models

import (
	"database/sql"
	"time"

	"github.com/Agrimaldo/blockchain-in-go/util"
	"github.com/uptrace/bun"
)

type Block struct {
	bun.BaseModel `bun:"table:tb_block,alias:b"`
	Id            int64 `bun:"Id,pk,autoincrement"`
	Height        int32
	PreviousHash  string
	Hash          string
	Creator       string
	CreatedAt     time.Time     `bun:"createdAt,nullzero,notnull,default:current_timestamp"`
	Transactions  []Transaction `bun:"rel:has-many,join:Id=BlockId"`
}

func (b *Block) CreateGenesis() {
	b.Height = 1
	b.PreviousHash = "-"
	b.Hash = util.NewHash()
	b.Creator = "System"
	b.CreatedAt = time.Date(2019, time.September, 10, 0, 0, 0, 0, time.Local)
	b.Transactions = append(b.Transactions, Transaction{Message: "Genesis Block created by P.HC  on 2019 10 24"})
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
