package auth

import (
	"context"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"math/big"
	"time"
	"xAdmin/config"
	log "xAdmin/logrus"
	"xAdmin/models"
	"xAdmin/utils"
)

// 获取消息
func GetMessage(ctx context.Context) {
	go func() {
		for {
			var t models.Transfer
			ts, err := t.GetTransferList()
			if err != nil {
				log.Error("获取转账记录失败！", err.Error())
			}
			for _, v := range ts {
				Cid, err := cid.Decode(v.Cid)
				if err != nil {
					continue
				}
				msg, err := config.FullAPI.Rpc.StateReplay(ctx, types.TipSetKey{}, Cid)
				if err != nil {
					log.Error("获取消息失败：", err)
					continue
				}
				totalstr := new(big.Int).Add(msg.GasCost.GasUsed.Int, msg.GasCost.TotalCost.Int).String()
				log.Info("手续费：", totalstr)
				total := utils.NanoOrAttoToFIL(totalstr, utils.AttoFIL)
				//totalf := strconv.FormatFloat(total,'f',-1,64)
				err = v.SetServiceCharge(total)
				log.Info("手续费：", total)
			}
			time.Sleep(time.Minute)
		}
	}()
}
