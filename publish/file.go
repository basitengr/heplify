package publish

import (
	"encoding/json"

	"github.com/negbie/heplify/decoder"
	"github.com/negbie/heplify/logp"
)

type FileOutputer struct {
}

func (fo *FileOutputer) Output(pkt *decoder.Packet) {
	jsonPkt, err := json.MarshalIndent(pkt, "", "  ")
	if err != nil {
		logp.Err("json %v", err)
		return
	}
	logp.Info("%s", jsonPkt)
}

func NewFileOutputer() (*FileOutputer, error) {
	fo := &FileOutputer{}
	return fo, nil
}
