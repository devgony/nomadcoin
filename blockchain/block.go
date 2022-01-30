package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/devgony/nomadcoin/db"
	"github.com/devgony/nomadcoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"Hash"`
	PrevHash string `json:"PrevHash:omitempty"`
	Height   int    `json:"Height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}
