
<template>
  <div>
    <el-row :gutter="20" class="panel-group">
      <el-col :xs="12" :sm="12" :lg="10" class="card-panel-col">
        <div class="card-panel-icon-wrapper icon-money" @click="handleSetLineChartData('purchases')">
          当前区块高度 :  <count-to :start-val="0" :decimals="0" :end-val="this.BlockNumber" :duration="1000" class="card-panel-num" /> 
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="10" class="card-panel-col">
        <div class="card-panel-icon-wrapper icon-money" @click="handleSetLineChartData('purchases')">
           GOF 资金池 ：<count-to :start-val="0" :decimals="0" :end-val="this.GOF1" :duration="1000" class="card-panel-num" style="color:red" /> 
        </div>
         <div class="card-panel-icon-wrapper icon-money" @click="handleSetLineChartData('purchases')">
         USDT资金池：<count-to :start-val="0" :decimals="2" :end-val="this.USDT1" :duration="2000" class="card-panel-num" style="color:red" />
         </div>
      </el-col>
    </el-row>
    <el-row :gutter="40" class="panel-group">
      <el-col :xs="12" :sm="12" :lg="12" class="card-panel-col">
        <div class="card-panel" @click="handleSetLineChartData('purchases')">
          <div class="card-panel-icon-wrapper icon-money">
            <svg-icon icon-class="money" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              GOF资产总数
            </div>
            <count-to :start-val="0" :decimals="2" :end-val="this.TCO" :duration="2000" class="card-panel-num" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              USDT资产总数
            </div>
            <count-to :start-val="0" :decimals="2" :end-val="this.USDT" :duration="2000" class="card-panel-num" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              ETH资产总数
            </div>
            <count-to :start-val="0" :decimals="2" :end-val="this.ETH" :duration="2000" class="card-panel-num" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              彩票手续费总数
            </div>
            <count-to :start-val="0" :decimals="2" :end-val="this.Free" :duration="2000" class="card-panel-num" />
          </div>
        </div>
      </el-col>
      <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
        <div class="card-panel" @click="handleSetLineChartData('newVisitis')">
          <div class="card-panel-icon-wrapper icon-people">
            <svg-icon icon-class="peoples" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              注册总数
            </div>
            <count-to :start-val="0" :end-val="userCount.Usercount" :duration="2600" class="card-panel-num" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              今日注册
            </div>
            <count-to :start-val="0" :end-val="userCount.Todaycount" :duration="2600" class="card-panel-num" />
          </div>

        </div>
      </el-col>
    <!-- <el-col :xs="12" :sm="12" :lg="8" class="card-panel-col">
        <div class="card-panel" @click="handleSetLineChartData('messages')">
            <div class="card-panel-icon-wrapper icon-message">
                <svg-icon icon-class="shopping" class-name="card-panel-icon" />
            </div>
            <div class="card-panel-description">
                <div class="card-panel-text">
                    二元期权
                </div>
                <count-to :start-val="0" :end-val="gamesCount.brcount" :duration="3000" class="card-panel-num" />
            </div>
            <div class="card-panel-description">
                <div class="card-panel-text">
                    FOMO总数
                </div>
                <count-to :start-val="0" :end-val="gamesCount.fomocount" :duration="3000" class="card-panel-num" />
            </div>
            <div class="card-panel-description">
                <div class="card-panel-text">
                    DICE总数
                </div>
                <count-to :start-val="0" :end-val="gamesCount.dicecount" :duration="3000" class="card-panel-num" />
            </div>
            <div class="card-panel-description">
                <div class="card-panel-text">
                    彩票总数
                </div>
                <count-to :start-val="0" :end-val="gamesCount.betcount" :duration="3000" class="card-panel-num" />
            </div>

        </div>
    </el-col> -->

    </el-row>
    <el-row :gutter="40" class="panel-group">
      <el-col :xs="12" :sm="12" :lg="16" class="card-panel-col">
        <div class="card-panel" @click="handleSetLineChartData('messages')">
          <div class="card-panel-icon-wrapper icon-message">
            <svg-icon icon-class="shopping" class-name="card-panel-icon" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              二元期权参与总数
            </div>
            <count-to :start-val="0" :end-val="gamesCount.brcount" :duration="3000" class="card-panel-num1" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              FOMO下注总数
            </div>
            <count-to :start-val="0" :end-val="gamesCount.fomocount" :duration="3000" class="card-panel-num1" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              DICE下注总数
            </div>
            <count-to :start-val="0" :end-val="gamesCount.dicecount" :duration="3000" class="card-panel-num1" />
          </div>
          <div class="card-panel-description">
            <div class="card-panel-text">
              彩票下注总数
            </div>
            <count-to :start-val="0" :end-val="gamesCount.betcount" :duration="3000" class="card-panel-num1" />
          </div>

        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import CountTo from 'vue-count-to'
import {
  GetUserCount, GetGamesCount, GetUserAssetCcount, GetWalletInfo,GetLotteryFreeCount 
} from '@/api/admin/PanelGroup'
export default {
  components: { CountTo },
  data() {
    return {
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        status: '',
        keyword: ''
        // cid: undefined
      },
      userCount: [],
      gamesCount: [],
      userAssetCcount: [],
      BlockNumber: 0,
      USDT: 0,
      ETH: 0,
      TCO: 0,
      USDT1: 0,
      GOF1: 0,
      Free: 0
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      GetUserCount().then(response => {
        this.userCount = response.data.list

        //  this.total = response.data.count
        //  this.loading = false
      })

      GetGamesCount().then(response => {
        this.gamesCount = response.data.list

        //  this.total = response.data.count
        //  this.loading = false
      })
      GetLotteryFreeCount().then(response => {
        this.Free = response.data.list['freecount']
      })

      GetWalletInfo().then(response => {
        this.BlockNumber = response.data
      })

      GetUserAssetCcount().then(response => {
        this.userAssetCcount = response.data.list
        for (var i = 0; i < this.userAssetCcount.length; i++) {
          const name = this.userAssetCcount[i].symbol
          switch (name.toUpperCase()) {
            case 'USDT':
              this.USDT = Number(this.userAssetCcount[i].quantity)
              break
            case 'GOF':
              this.TCO = Number(this.userAssetCcount[i].quantity)
              break
            case 'ETH':
              this.ETH = Number(this.userAssetCcount[i].quantity)
              break
            case 'USDT1':
              this.USDT1 = Number(this.userAssetCcount[i].quantity)
              break
            case 'GOF1':
              this.GOF1 = Number(this.userAssetCcount[i].quantity)
              break
          }
        }
      })
    },
    handleSetLineChartData(type) {
      this.$emit('handleSetLineChartData', type)
    },
    components: {
      CountTo
    }
  }
}
</script>

<style lang="scss" scoped>
.panel-group {
    margin-top: 18px;

    .card-panel-col {
        margin-bottom: 32px;
    }

    .card-panel {
        height: 108px;
        cursor: pointer;
        font-size: 12px;
        position: relative;
        overflow: hidden;
        color: #666;
        background: #fff;
        box-shadow: 4px 4px 40px rgba(0, 0, 0, .05);
        border-color: rgba(0, 0, 0, .05);

        &:hover {
            .card-panel-icon-wrapper {
                color: #fff;
            }

            .icon-people {
                background: #40c9c6;
            }

            .icon-message {
                background: #36a3f7;
            }

            .icon-money {
                background: #f4516c;
            }

            .icon-shopping {
                background: #34bfa3
            }
        }

        .icon-people {
            color: #40c9c6;
        }

        .icon-message {
            color: #36a3f7;
        }

        .icon-money {
            color: #f4516c;
        }

        .icon-shopping {
            color: #34bfa3
        }

        .card-panel-icon-wrapper {
            float: left;
            margin: 14px 0 0 14px;
            padding: 16px;
            transition: all 0.38s ease-out;
            border-radius: 6px;
        }

        .card-panel-icon {
            float: left;
            font-size: 48px;
        }

        .card-panel-description {
            float: left;
            font-weight: bold;
            margin: 26px;
            margin-left: 20px;

            .card-panel-text {
                line-height: 18px;
                color: rgba(0, 0, 0, 0.45);
                font-size: 16px;
                margin-bottom: 12px;
            }
            .card-panel-num {
                font-size: 20px;
            }
            .card-panel-num1 {
                font-size: 20px;
                margin-left: 36px;
            }
        }
    }
}

@media (max-width:550px) {
    .card-panel-description {
        display: none;
    }

    .card-panel-icon-wrapper {
        float: none !important;
        width: 100%;
        height: 100%;
        margin: 0 !important;

        .svg-icon {
            display: block;
            margin: 14px auto !important;
            float: none !important;
        }
    }
}
</style>
