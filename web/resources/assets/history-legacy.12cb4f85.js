/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
!function(){function e(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function t(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}System.register(["./index-legacy.763f76da.js"],(function(n,a){"use strict";var r,u,o,l,i,s,c,v,f,m,p,d,g,y,b,h,O,x,S,w,I,j,k,q,C,N=document.createElement("style");return N.textContent='.contextmenu{width:100px;margin:0;border:1px solid #ccc;background:#fff;z-index:3000;position:absolute;list-style-type:none;padding:5px 0;border-radius:4px;font-size:14px;color:#333;box-shadow:2px 2px 3px rgba(0,0,0,.2)}.el-tabs__item .el-icon-close{color:initial!important}.el-tabs__item .dot{content:"";width:9px;height:9px;margin-right:8px;display:inline-block;border-radius:50%;transition:background-color .2s}.contextmenu li{margin:0;padding:7px 16px}.contextmenu li:hover{background:#f2f2f2;cursor:pointer}\n',document.head.appendChild(N),{setters:[function(e){r=e.V,u=e.u,o=e.r,l=e.j,i=e.R,s=e.I,c=e.W,v=e.D,f=e.g,m=e.o,p=e.c,d=e.d,g=e.w,y=e.F,b=e.v,h=e.l,O=e.b,x=e.H,S=e.G,w=e.h,I=e.y,j=e.X,k=e.Y,q=e.J,C=e.K}],execute:function(){var a={class:"router-history"},N=["tab"],P={name:"HistoryComponent"};n("default",Object.assign(P,{setup:function(n){var P=r(),J=u(),E=function(e){return e.name+JSON.stringify(e.query)+JSON.stringify(e.params)},V=o([]),_=o(""),A=o(!1),D=l(),R=function(e){return e.name+JSON.stringify(e.query)+JSON.stringify(e.params)},T=o(0),L=o(0),z=o(!1),H=o(!1),K=o(""),X=i((function(){return D.userInfo.authority.defaultRouter})),Y=function(){V.value=[{name:X.value,meta:{title:"首页"},query:{},params:{}}],J.push({name:X.value}),A.value=!1,sessionStorage.setItem("historys",JSON.stringify(V.value))},F=function(){var e,t=V.value.findIndex((function(t){return E(t)===K.value&&(e=t),E(t)===K.value})),n=V.value.findIndex((function(e){return E(e)===_.value}));V.value.splice(0,t),t>n&&J.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},G=function(){var e,t=V.value.findIndex((function(t){return E(t)===K.value&&(e=t),E(t)===K.value})),n=V.value.findIndex((function(e){return E(e)===_.value}));V.value.splice(t+1,V.value.length),t<n&&J.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},U=function(){var e;V.value=V.value.filter((function(t){return E(t)===K.value&&(e=t),E(t)===K.value})),J.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},W=function(n){if(!V.value.some((function(e){return function(e,t){if(e.name!==t.name)return!1;if(Object.keys(e.query).length!==Object.keys(t.query).length||Object.keys(e.params).length!==Object.keys(t.params).length)return!1;for(var n in e.query)if(e.query[n]!==t.query[n])return!1;for(var a in e.params)if(e.params[a]!==t.params[a])return!1;return!0}(e,n)}))){var a={};a.name=n.name,a.meta=function(n){for(var a=1;a<arguments.length;a++){var r=null!=arguments[a]?arguments[a]:{};a%2?e(Object(r),!0).forEach((function(e){t(n,e,r[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(n,Object.getOwnPropertyDescriptors(r)):e(Object(r)).forEach((function(e){Object.defineProperty(n,e,Object.getOwnPropertyDescriptor(r,e))}))}return n}({},n.meta),delete a.meta.matched,a.query=n.query,a.params=n.params,V.value.push(a)}window.sessionStorage.setItem("activeValue",E(n))},$=o({}),B=function(e){var t=$.value[e];J.push({name:t.name,query:t.query,params:t.params})},M=function(e){var t=V.value.findIndex((function(t){return E(t)===e}));E(P)===e&&(1===V.value.length?J.push({name:X.value}):t<V.value.length-1?J.push({name:V.value[t+1].name,query:V.value[t+1].query,params:V.value[t+1].params}):J.push({name:V.value[t-1].name,query:V.value[t-1].query,params:V.value[t-1].params})),V.value.splice(t,1)};s((function(){return A.value}),(function(){A.value?document.body.addEventListener("click",(function(){A.value=!1})):document.body.removeEventListener("click",(function(){A.value=!1}))})),s((function(){return P}),(function(e,t){"Login"!==e.name&&"Reload"!==e.name&&(V.value=V.value.filter((function(e){return!e.meta.closeTab})),W(e),sessionStorage.setItem("historys",JSON.stringify(V.value)),_.value=window.sessionStorage.getItem("activeValue"))}),{deep:!0}),s((function(){return V.value}),(function(){sessionStorage.setItem("historys",JSON.stringify(V.value)),$.value={},V.value.forEach((function(e){$.value[E(e)]=e})),c.emit("setKeepAlive",V.value)}),{deep:!0});return function(){c.on("closeThisPage",(function(){M(R(P))})),c.on("closeAllPage",(function(){Y()})),c.on("mobile",(function(e){H.value=e})),c.on("collapse",(function(e){z.value=e}));var e=[{name:X.value,meta:{title:"首页"},query:{},params:{}}];V.value=JSON.parse(sessionStorage.getItem("historys"))||e,window.sessionStorage.getItem("activeValue")?_.value=window.sessionStorage.getItem("activeValue"):_.value=E(P),W(P),"true"===window.sessionStorage.getItem("needCloseAll")&&(Y(),window.sessionStorage.removeItem("needCloseAll"))}(),v((function(){c.off("collapse"),c.off("mobile")})),function(e,t){var n=f("el-tab-pane"),r=f("el-tabs");return m(),p("div",a,[d(r,{modelValue:_.value,"onUpdate:modelValue":t[0]||(t[0]=function(e){return _.value=e}),closable:!(1===V.value.length&&e.$route.name===S(X)),type:"card",onContextmenu:t[1]||(t[1]=k((function(e){return function(e){if(1===V.value.length&&P.name===X.value)return!1;var t,n="";(n="SPAN"===e.srcElement.nodeName?e.srcElement.offsetParent.id:e.srcElement.id)&&(A.value=!0,t=z.value?54:220,H.value&&(t=0),T.value=e.clientX-t,L.value=e.clientY+10,K.value=n.substring(4))}(e)}),["prevent"])),onTabChange:B,onTabRemove:M},{default:g((function(){return[(m(!0),p(y,null,b(V.value,(function(e){return m(),h(n,{key:R(e),label:e.meta.title,name:R(e),tab:e,class:"gva-tab"},{label:g((function(){return[O("span",{tab:e,style:x({color:_.value===R(e)?S(D).activeColor:"#333"})},[O("i",{class:"dot",style:x({backgroundColor:_.value===R(e)?S(D).activeColor:"#ddd"})},null,4),w(" "+I(S(j)(e.meta.title,e)),1)],12,N)]})),_:2},1032,["label","name","tab"])})),128))]})),_:1},8,["modelValue","closable"]),q(O("ul",{style:x({left:T.value+"px",top:L.value+"px"}),class:"contextmenu"},[O("li",{onClick:Y},"关闭所有"),O("li",{onClick:F},"关闭左侧"),O("li",{onClick:G},"关闭右侧"),O("li",{onClick:U},"关闭其他")],4),[[C,A.value]])])}}}))}}}))}();
