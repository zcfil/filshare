<template>
  <div class="in">
    <div v-if="data.height!=undefined" class="block">
      <div>区块高度</div>
      <div>{{ data.height }}</div>
    </div>
    <div v-if="data.total_valid_hashrate!=undefined" class="block">
      <div>全网有效算力</div>
      <div>{{ data.total_valid_hashrate }}</div>
    </div>
    <div v-if="data.active_provider!=undefined" class="block">
      <div>活跃提供者</div>
      <div>{{ data.active_provider }}</div>
    </div>
    <div v-if="data.every_award!=undefined" class="block">
      <div>每区块奖励</div>
      <div>{{ data.every_award }}</div>
    </div>
    <div v-if="data.income_24ft!=undefined" class="block">
      <div>24h平均提供服务收益</div>
      <div>{{ data.income_24ft }}</div>
    </div>
    <div v-if="data.output_24f!=undefined" class="block">
      <div>近24h产出量</div>
      <div>{{ data.output_24f }}</div>
    </div>
    <div v-if="data.cur_sector_pledge!=undefined" class="block">
      <div>当前扇区质押量</div>
      <div>{{ data.cur_sector_pledge }}</div>
    </div>
    <div v-if="data.fil_pledge!=undefined" class="block">
      <div>FIL质押量</div>
      <div>{{ data.fil_pledge }}</div>
    </div>
    <div v-if="data.message_count24!=undefined" class="block">
      <div>24h消息数</div>
      <div>{{ data.message_count24 }}</div>
    </div>
    <div v-if="data.fil_turnover!=undefined" class="block">
      <div>FIL流通量</div>
      <div>{{ data.fil_turnover }}</div>
    </div>
    <div v-if="data.total_account!=undefined" class="block">
      <div>总帐户数</div>
      <div>{{ data.total_account }}</div>
    </div>
    <div v-if="data.average_interval!=undefined" class="block">
      <div>平均区块间隔</div>
      <div>{{ data.average_interval }}</div>
    </div>
    <div v-if="data.average_height_count!=undefined" class="block">
      <div>平均高度区块数量</div>
      <div>{{ data.average_height_count }}</div>
    </div>
    <div v-if="data.new_add_hashrate_ft!=undefined" class="block">
      <div>新增算力成本</div>
      <div>{{ data.new_add_hashrate_ft }}</div>
    </div>
    <div v-if="data.cur_basic_hashrate!=undefined" class="block">
      <div>当前基础费率</div>
      <div>{{ data.cur_basic_hashrate }}</div>
    </div>
    <div v-if="data.fil_destroy!=undefined" class="block">
      <div>FIL销毁量</div>
      <div>{{ data.fil_destroy }}</div>
    </div>
    <div v-if="data.fil_supply!=undefined" class="block">
      <div>FIL提供量</div>
      <div>{{ data.fil_supply }}</div>
    </div>
    <div v-if="data.flow_rate!=undefined" class="block">
      <div>FIL流通率</div>
      <div>{{ data.flow_rate }}</div>
    </div>
    <div v-if="data.price!=undefined" class="block">
      <div>最新价格</div>
      <div>{{ data.price }}</div>
    </div>
    <div v-if="data.date!=undefined" class="block">
      <div>日期</div>
      <div>{{ parseTime(data.date) }}</div>
    </div>
  </div>
</template>

<script>
// import { officialInfo } from '@/api/customer/customer'
export default {
  name: 'Dashboard',
  data() {
    return {
      data: {}
    }
  },
  computed: {},
  created() {
    this.auto()
  },
  methods: {
    auto() {
      // officialInfo().then(res=> {
      //   this.data=res.data
      // })
    }
  }
}
</script>
<style scoped>
.in {
  padding: 20px;
  overflow: hidden;
}
.block {
  float: left;
  background-color: rgba(240,246,251);
  width: 18%;
  height: 110px;
  margin: 10px 1%;
  text-align: center;
}
.block div:nth-child(1) {
  color: #888888;
  font-size: 12px;
  margin-top: 34px;
}
.block div:nth-child(2) {
  color: #000000;
  font-size: 24px;
  margin-top: 4px;
}
</style>
