<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import resize from './mixins/resize'
import {
  GetStatGamesmoth
} from '@/api/admin/PanelGroup'
export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '350px'
    },
    autoResize: {
      type: Boolean,
      default: true
    },
    chartData: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      chart: null,
      monthDays: [],
      StatGamesMoth: [],
      br: [],
      fomo: [],
      dice: [],
      bet: []
    }
  },
  watch: {
    chartData: {
      deep: true,
      handler(val) {

      }
    }
  },
  created() {
    this.GetMothDate()
    this.getList()
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart()
    })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    getList() {
      GetStatGamesmoth().then(response => {
        this.StatGamesMoth = response.data.list
        for (var i = 0; i < this.monthDays.length; i++) {
          for (var j = 0; j < this.StatGamesMoth.length; j++) {
            var date = this.StatGamesMoth[j].date1
            // console.log(this.monthDays[i],":",date.substr(date.length - 2, date.length))
            if (this.monthDays[i] === date.substr(date.length - 2, date.length)) {
              this.bet[i] = Number(this.StatGamesMoth[j].betcount)
              this.br[i] = Number(this.StatGamesMoth[j].brcount)
              this.dice[i] = Number(this.StatGamesMoth[j].dicecount)
              this.fomo[i] = Number(this.StatGamesMoth[j].fomocount)
              break
            } else {
              this.bet[i] = 0
              this.br[i] = 0
              this.dice[i] = 0
              this.fomo[i] = 0
            }
          }
        }
        this.setOptions()
      })
    },
    initChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      // this.getList()
    },
    GetMothDate() {
      var date = new Date()
      var year = date.getFullYear()
      var month = date.getMonth() + 1
      var d = new Date(year, month, 0)
      /* 返回当月的天数 */
      for (var i = 0; i < d.getDate(); i++) {
        this.monthDays[i] = i + 1 + ''
      }
    },
    /**
         *  @obj {expectedData,actualData}
         * this.setOptions({expectedData:123,actualData:456});
         */
    setOptions() {
      // var expectedData=obj.expectedDat
      this.chart.setOption({
        xAxis: {
          data: this.monthDays,
          nameTextStyle: {
            fontStyle: 'italic'
          },
          boundaryGap: false,
          axisTick: {
            show: false
          }
        },
        grid: {
          left: 10,
          right: 10,
          bottom: 20,
          top: 30,
          containLabel: true
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          },
          padding: [5, 10]
        },
        yAxis: {
          axisTick: {
            show: false
          }
        },
        legend: {
          data: ['二元期权', 'DICE', 'FOMO', '彩票']
        },
        series: [{
          name: 'DICE',
          itemStyle: {
            normal: {
              color: '#FF005A',
              lineStyle: {
                color: '#FF005A',
                width: 2
              }
            }
          },
          smooth: true,
          type: 'line',
          data: this.dice,
          animationDuration: 2800,
          animationEasing: 'cubicInOut'
        },
        {
          name: '彩票',
          smooth: true,
          type: 'line',
          itemStyle: {
            normal: {
              color: '#3888fa',
              lineStyle: {
                color: '#3888fa',
                width: 2
              },
              areaStyle: {
                color: '#f3f8ff'
              }
            }
          },
          data: this.bet,
          animationDuration: 2800,
          animationEasing: 'quadraticOut'
        },
        {
          name: '二元期权',
          smooth: true,
          type: 'line',
          itemStyle: {
            normal: {
              color: '#c23531',
              lineStyle: {
                color: '#c23531',
                width: 2
              },
              areaStyle: {
                color: '#f3f8ff'
              }
            }
          },
          data: this.br,
          animationDuration: 2800,
          animationEasing: 'quadraticOut'
        },
        {
          name: 'FOMO',
          smooth: true,
          type: 'line',
          itemStyle: {
            normal: {
              color: '#2f4554',
              lineStyle: {
                color: '#2f4554',
                width: 2
              },
              areaStyle: {
                color: '#f3f8ff'
              }
            }
          },
          data: this.fomo,
          animationDuration: 2800,
          animationEasing: 'quadraticOut'
        }

        ]
      })
    }
  }
}
</script>
