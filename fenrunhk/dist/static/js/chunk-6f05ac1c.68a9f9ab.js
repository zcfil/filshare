(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-6f05ac1c"],{"293f":function(e,t,n){},"2e77":function(e,t,n){"use strict";n("293f")},"790d":function(e,t,n){"use strict";n("87f0")},"87e0":function(e,t,n){},"87f0":function(e,t,n){},"9ed6":function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"login-container"},[n("el-form",{ref:"loginForm",staticClass:"login-form",attrs:{model:e.loginForm,rules:e.loginRules,autocomplete:"on","label-position":"left"}},[n("div",{staticClass:"title-container"},[n("h3",{staticClass:"title"},[e._v("分润系统")])]),n("el-form-item",{attrs:{prop:"username"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"user"}})],1),n("el-input",{ref:"username",attrs:{placeholder:"用户名",name:"username",type:"text",tabindex:"1",autocomplete:"on"},model:{value:e.loginForm.username,callback:function(t){e.$set(e.loginForm,"username",t)},expression:"loginForm.username"}})],1),n("el-tooltip",{attrs:{content:"Caps lock is On",placement:"right",manual:""},model:{value:e.capsTooltip,callback:function(t){e.capsTooltip=t},expression:"capsTooltip"}},[n("el-form-item",{attrs:{prop:"password"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"password"}})],1),n("el-input",{key:e.passwordType,ref:"password",attrs:{type:e.passwordType,placeholder:"密码",name:"password",tabindex:"2",autocomplete:"on"},on:{blur:function(t){e.capsTooltip=!1}},nativeOn:{keyup:[function(t){return e.checkCapslock(t)},function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleLogin(t)}]},model:{value:e.loginForm.password,callback:function(t){e.$set(e.loginForm,"password",t)},expression:"loginForm.password"}}),n("span",{staticClass:"show-pwd",on:{click:e.showPwd}},[n("svg-icon",{attrs:{"icon-class":"password"===e.passwordType?"eye":"eye-open"}})],1)],1)],1),n("el-form-item",{staticStyle:{width:"60%",float:"left"},attrs:{prop:"code"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"validCode"}})],1),n("el-input",{ref:"username",staticStyle:{width:"75%"},attrs:{placeholder:"验证码",name:"username",type:"text",tabindex:"3",autocomplete:"off"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleLogin(t)}},model:{value:e.loginForm.code,callback:function(t){e.$set(e.loginForm,"code",t)},expression:"loginForm.code"}})],1),n("div",{staticClass:"login-code",staticStyle:{width:"38%",height:"48px",float:"right"}},[n("img",{staticStyle:{height:"48px",width:"100%"},attrs:{src:e.codeUrl},on:{click:e.getCode}})]),n("el-button",{staticStyle:{width:"100%","margin-bottom":"30px"},attrs:{loading:e.loading,type:"primary"},nativeOn:{click:function(t){return t.preventDefault(),e.handleLogin(t)}}},[e.loading?n("span",[e._v("登 录 中...")]):n("span",[e._v("登 录")])])],1),n("el-dialog",{attrs:{title:"Or connect with",visible:e.showDialog},on:{"update:visible":function(t){e.showDialog=t}}},[e._v(" Can not be simulated on local, so please combine you own business simulation! ! ! "),n("br"),n("br"),n("br"),n("social-sign")],1)],1)},s=[],i=(n("ac6a"),n("456d"),n("b775"));function a(){return Object(i["a"])({url:"/api/v1/getCaptcha",method:"get"})}var r=n("c62e"),c={name:"Login",components:{SocialSign:r["default"]},data:function(){return{codeUrl:"",cookiePassword:"",loginForm:{username:"",password:"",rememberMe:!1,code:"",uuid:""},loginRules:{username:[{required:!0,trigger:"blur",message:"用户名不能为空"}],password:[{required:!0,trigger:"blur",message:"密码不能为空"}],code:[{required:!0,trigger:"change",message:"验证码不能为空"}]},passwordType:"password",capsTooltip:!1,loading:!1,showDialog:!1,redirect:void 0,otherQuery:{}}},watch:{$route:{handler:function(e){var t=e.query;t&&(this.redirect=t.redirect,this.otherQuery=this.getOtherQuery(t))},immediate:!0}},created:function(){this.getCode()},mounted:function(){""===this.loginForm.username?this.$refs.username.focus():""===this.loginForm.password&&this.$refs.password.focus()},destroyed:function(){},methods:{getCode:function(){var e=this;a().then((function(t){console.log(t),e.codeUrl=t.data,e.loginForm.uuid=t.id}))},checkCapslock:function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},t=e.shiftKey,n=e.key;n&&1===n.length&&(this.capsTooltip=!!(t&&n>="a"&&n<="z"||!t&&n>="A"&&n<="Z")),"CapsLock"===n&&!0===this.capsTooltip&&(this.capsTooltip=!1)},showPwd:function(){var e=this;"password"===this.passwordType?this.passwordType="":this.passwordType="password",this.$nextTick((function(){e.$refs.password.focus()}))},handleLogin:function(){var e=this;this.$refs.loginForm.validate((function(t){if(!t)return console.log("error submit!!"),!1;e.loading=!0,e.$store.dispatch("user/login",e.loginForm).then((function(){e.$router.push({path:"/",query:e.otherQuery}),e.loading=!1})).catch((function(){e.loading=!1,e.getCode()}))}))},getOtherQuery:function(e){return Object.keys(e).reduce((function(t,n){return"redirect"!==n&&(t[n]=e[n]),t}),{})}}},l=c,u=(n("790d"),n("2e77"),n("2877")),d=Object(u["a"])(l,o,s,!1,null,"2776acdc",null);t["default"]=d.exports},bfec:function(e,t,n){"use strict";n("87e0")},c62e:function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"social-signup-container"},[n("div",{staticClass:"sign-btn",on:{click:function(t){return e.wechatHandleClick("wechat")}}},[n("span",{staticClass:"wx-svg-container"},[n("svg-icon",{staticClass:"icon",attrs:{"icon-class":"wechat"}})],1),e._v(" WeChat ")]),n("div",{staticClass:"sign-btn",on:{click:function(t){return e.tencentHandleClick("tencent")}}},[n("span",{staticClass:"qq-svg-container"},[n("svg-icon",{staticClass:"icon",attrs:{"icon-class":"qq"}})],1),e._v(" QQ ")])])},s=[],i={name:"SocialSignin",methods:{wechatHandleClick:function(e){alert("ok")},tencentHandleClick:function(e){alert("ok")}}},a=i,r=(n("bfec"),n("2877")),c=Object(r["a"])(a,o,s,!1,null,"7309fbbb",null);t["default"]=c.exports}}]);