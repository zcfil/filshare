(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-customer-customer"],{"21df":function(t,e,a){"use strict";a.r(e);var s=a("c5eb"),i=a("9b7a");for(var o in i)["default"].indexOf(o)<0&&function(t){a.d(e,t,(function(){return i[t]}))}(o);a("341b");var n=a("f0c5"),r=Object(n["a"])(i["default"],s["b"],s["c"],!1,null,"7689a574",null,!1,s["a"],void 0);e["default"]=r.exports},"341b":function(t,e,a){"use strict";var s=a("60ce"),i=a.n(s);i.a},"60ce":function(t,e,a){var s=a("7d5d");s.__esModule&&(s=s.default),"string"===typeof s&&(s=[[t.i,s,""]]),s.locals&&(t.exports=s.locals);var i=a("4f06").default;i("1696e66b",s,!0,{sourceMap:!1,shadowMode:!1})},"7d5d":function(t,e,a){var s=a("24fb");e=s(!1),e.push([t.i,".tab[data-v-7689a574]{font-size:14px;color:#444;border-top:1px solid #ddd;padding:11px 20px 15px}.tab uni-view[data-v-7689a574]{margin-top:4px}",""]),t.exports=e},"9b7a":function(t,e,a){"use strict";a.r(e);var s=a("ff4c"),i=a.n(s);for(var o in s)["default"].indexOf(o)<0&&function(t){a.d(e,t,(function(){return s[t]}))}(o);e["default"]=i.a},c5eb:function(t,e,a){"use strict";a.d(e,"b",(function(){return s})),a.d(e,"c",(function(){return i})),a.d(e,"a",(function(){}));var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-uni-view",{staticClass:"customer"},t._l(t.data,(function(e,s){return a("v-uni-view",{key:s,staticClass:"tab"},[a("v-uni-view",[t._v("算力（TB）："+t._s(e.storage))]),a("v-uni-view",[t._v("备注："+t._s(e.remark))]),a("v-uni-view",[t._v("创建时间："+t._s(t.parseTime(e.create_time)))]),a("v-uni-view",[t._v("状态："),a("v-uni-text",{directives:[{name:"show",rawName:"v-show",value:0==e.status,expression:"item.status==0"}],staticStyle:{color:"rgb(31, 211, 0)"}},[t._v("进行中")]),a("v-uni-text",{directives:[{name:"show",rawName:"v-show",value:1==e.status,expression:"item.status==1"}],staticStyle:{color:"rgb(255, 80, 80)"}},[t._v("已终止")]),a("v-uni-text",{directives:[{name:"show",rawName:"v-show",value:2==e.status,expression:"item.status==2"}],staticStyle:{color:"rgb(31, 211, 0)"}},[t._v("线性释放中")]),a("v-uni-text",{directives:[{name:"show",rawName:"v-show",value:3==e.status,expression:"item.status==3"}],staticStyle:{color:"#007AFF"}},[t._v("已完成线性释放")])],1)],1)})),1)},i=[]},ff4c:function(t,e,a){"use strict";a("7a82"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0,a("99af");var s=a("e9e5"),i={data:function(){return{data:[],queryParams:{pageIndex:1,pageSize:10},loadStatus:"loading",isLoadMore:!1}},onShow:function(){this.data=[],this.auto()},onReachBottom:function(){this.isLoadMore||(this.isLoadMore=!0,this.queryParams.pageIndex+=1,this.auto())},onPullDownRefresh:function(){this.data=[],this.queryParams.pageIndex=1,this.auto(),uni.stopPullDownRefresh()},methods:{auto:function(){var t=this;(0,s.investmentList)(this.queryParams).then((function(e){e.list?(t.data=t.data.concat(e.list),e.list.length<t.queryParams.pageSize?(t.isLoadMore=!0,t.loadStatus="nomore"):t.isLoadMore=!1):(t.isLoadMore=!0,t.loadStatus="nomore")}))}}};e.default=i}}]);