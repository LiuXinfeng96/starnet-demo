/*! 
 Build based on gin-vue-admin 
 Time : 1672137460000 */
System.register(["./index-legacy.7e49b66d.js","./index-legacy.eb45a827.js","./menuItem-legacy.d14027ac.js","./asyncSubmenu-legacy.1e72810a.js"],(function(e,n){"use strict";var a,t,u,l,o,r,c,i,d,f,s,v,m,p,h,g,b,y,x,k,T,_=document.createElement("style");return _.textContent=".el-sub-menu__title:hover,.el-menu-item:hover{background:transparent}.el-scrollbar .el-scrollbar__view{height:100%}.menu-info .menu-contorl{line-height:52px;font-size:20px;display:table-cell;vertical-align:middle}\n",document.head.appendChild(_),{setters:[function(e){a=e.V,t=e.u,u=e.j,l=e.N,o=e.r,r=e.I,c=e.D,i=e.W,d=e.g,f=e.o,s=e.c,v=e.d,m=e.w,p=e.T,h=e.F,g=e.v,b=e.G,y=e.l,x=e.e,k=e.H},function(e){T=e.default},function(){},function(){}],execute:function(){var n={name:"Aside"};e("default",Object.assign(n,{setup:function(e){var n=a(),_=t(),j=u(),w=l(),F=o({}),M=function(){switch(j.sideMode){case"#fff":F.value={background:"#fff",activeBackground:"#4D70FF",activeText:"#fff",normalText:"#333",hoverBackground:"rgba(64, 158, 255, 0.08)",hoverText:"#333"};break;case"#191a23":F.value={background:"#191a23",activeBackground:"#4D70FF",activeText:"#fff",normalText:"#fff",hoverBackground:"rgba(64, 158, 255, 0.08)",hoverText:"#fff"}}};M();var B=o("");r((function(){return n}),(function(){B.value=n.meta.activeName||n.name}),{deep:!0}),r((function(){return j.sideMode}),(function(){M()}));var q=o(!1);B.value=n.meta.activeName||n.name,document.body.clientWidth<1e3&&(q.value=!q.value),i.on("collapse",(function(e){q.value=e})),c((function(){i.off("collapse")}));var D=function(e,a,t,u){var l,o,r={},c={};(null===(l=w.routeMap[e])||void 0===l?void 0:l.parameters)&&(null===(o=w.routeMap[e])||void 0===o||o.parameters.forEach((function(e){"query"===e.type?r[e.key]=e.value:c[e.key]=e.value}))),e!==n.name&&(e.indexOf("http://")>-1||e.indexOf("https://")>-1?window.open(e):_.push({name:e,query:r,params:c}))};return function(e,n){var a=d("el-menu"),t=d("el-scrollbar");return f(),s("div",{style:k({background:b(j).sideMode})},[v(t,{style:{height:"calc(100vh - 60px)"}},{default:m((function(){return[v(p,{duration:{enter:800,leave:100},mode:"out-in",name:"el-fade-in-linear"},{default:m((function(){return[v(a,{collapse:q.value,"collapse-transition":!1,"default-active":B.value,"background-color":F.value.background,"active-text-color":F.value.active,class:"el-menu-vertical","unique-opened":"",onSelect:D},{default:m((function(){return[(f(!0),s(h,null,g(b(w).asyncRouters[0].children,(function(e){return f(),s(h,null,[e.hidden?x("",!0):(f(),y(T,{key:e.name,"is-collapse":q.value,"router-info":e,theme:F.value},null,8,["is-collapse","router-info","theme"]))],64)})),256))]})),_:1},8,["collapse","default-active","background-color","active-text-color"])]})),_:1})]})),_:1})],4)}}}))}}}));