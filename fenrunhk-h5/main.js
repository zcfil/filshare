
// #ifndef VUE3
import Vue from 'vue'
import App from './App'
import tip from'./common/tip.js'
import { parseTime,moneyFormat } from './common/costum.js'
Vue.prototype.parseTime = parseTime
Vue.prototype.moneyFormat = moneyFormat

Vue.config.productionTip = false
Vue.prototype.$tip=tip;

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()
// #endif

// #ifdef VUE3
// import { createSSRApp } from 'vue'
// import App from './App.vue'
// export function createApp() {
//   const app = createSSRApp(App)
//   return {
//     app
//   }
// }
// #endif