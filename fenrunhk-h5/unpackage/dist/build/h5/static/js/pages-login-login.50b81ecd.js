(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-login-login"],{"3dd6":function(e,n,t){"use strict";t.d(n,"b",(function(){return a})),t.d(n,"c",(function(){return r})),t.d(n,"a",(function(){return o}));var o={uniForms:t("f95d").default,uniFormsItem:t("51a1").default,uniEasyinput:t("9957").default},a=function(){var e=this,n=e.$createElement,t=e._self._c||n;return t("v-uni-view",{staticClass:"login"},[t("v-uni-view",{staticClass:"welcome"},[e._v("用户登录")]),t("uni-forms",{ref:"form",staticClass:"content",attrs:{modelValue:e.loginForm,rules:e.rules}},[t("uni-forms-item",{attrs:{name:"phone"}},[t("uni-easyinput",{attrs:{type:"text",maxlength:"11",placeholder:"请输入手机号"},model:{value:e.loginForm.phone,callback:function(n){e.$set(e.loginForm,"phone",n)},expression:"loginForm.phone"}})],1),t("uni-forms-item",{attrs:{name:"password"}},[t("uni-easyinput",{attrs:{type:"password",maxlength:"18",placeholder:"请输入密码"},model:{value:e.loginForm.password,callback:function(n){e.$set(e.loginForm,"password",n)},expression:"loginForm.password"}})],1)],1),t("v-uni-button",{staticClass:"btn",on:{click:function(n){arguments[0]=n=e.$handleEvent(n),e.submit.apply(void 0,arguments)}}},[e._v("登录")])],1)},r=[]},"500b":function(e,n,t){"use strict";t("7a82"),Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var o=t("e9e5"),a={data:function(){return{loginForm:{phone:void 0,password:void 0},rules:{phone:{rules:[{required:!0,errorMessage:"请输入手机号"},{pattern:/^1[3|4|5|7|8|9][0-9]\d{8}$/,errorMessage:"请输入正确的手机号码"}]},password:{rules:[{required:!0,errorMessage:"请输入密码"},{pattern:/^\d{6,18}$/,errorMessage:"密码长度为6到18位的数字"}]}}}},methods:{submit:function(){var e=this;this.$refs.form.validate().then((function(n){(0,o.login)(e.loginForm).then((function(n){e.$tip.success("登录成功"),uni.setStorageSync("token",n.data),uni.switchTab({url:"/pages/index/index"})}))}))}}};n.default=a},5628:function(e,n,t){var o=t("5fff");o.__esModule&&(o=o.default),"string"===typeof o&&(o=[[e.i,o,""]]),o.locals&&(e.exports=o.locals);var a=t("4f06").default;a("7b7ad382",o,!0,{sourceMap:!1,shadowMode:!1})},"5fff":function(e,n,t){var o=t("24fb");n=o(!1),n.push([e.i,".login[data-v-9ae3953e]{position:relative;width:100vw;height:94vh;overflow:hidden;background:#f8f8f8}.welcome[data-v-9ae3953e]{margin:65px 0 95px 27px;font-size:25px;color:#333;text-shadow:1px 0 1px rgba(0,0,0,.3)}.content[data-v-9ae3953e]{padding:0 44px;font-size:20px}.btn[data-v-9ae3953e]{margin:20px 44px;background:#beac8b;color:#fff;height:40px;font-size:16px}",""]),e.exports=n},ad3a:function(e,n,t){"use strict";t.r(n);var o=t("3dd6"),a=t("c630");for(var r in a)["default"].indexOf(r)<0&&function(e){t.d(n,e,(function(){return a[e]}))}(r);t("c4e7");var i=t("f0c5"),s=Object(i["a"])(a["default"],o["b"],o["c"],!1,null,"9ae3953e",null,!1,o["a"],void 0);n["default"]=s.exports},c4e7:function(e,n,t){"use strict";var o=t("5628"),a=t.n(o);a.a},c630:function(e,n,t){"use strict";t.r(n);var o=t("500b"),a=t.n(o);for(var r in o)["default"].indexOf(r)<0&&function(e){t.d(n,e,(function(){return o[e]}))}(r);n["default"]=a.a}}]);