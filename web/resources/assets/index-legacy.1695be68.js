/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
System.register(["./menuItem-legacy.275f0d25.js","./asyncSubmenu-legacy.d3d75a22.js","./index-legacy.763f76da.js"],(function(e,n){"use strict";var t,r,o,u,i,f,l,c,a,s,d,h,m;return{setters:[function(e){t=e.default},function(e){r=e.default},function(e){o=e.R,u=e.g,i=e.o,f=e.l,l=e.w,c=e.c,a=e.F,s=e.v,d=e.e,h=e.m,m=e.G}],execute:function(){var n={name:"AsideComponent"};e("default",Object.assign(n,{props:{routerInfo:{type:Object,default:function(){return null}},isCollapse:{default:function(){return!1},type:Boolean},theme:{default:function(){return{}},type:Object}},setup:function(e){var n=e,p=o((function(){return n.routerInfo.children&&n.routerInfo.children.filter((function(e){return!e.hidden})).length?r:t}));return function(n,t){var r=u("AsideComponent");return e.routerInfo.hidden?d("",!0):(i(),f(h(m(p)),{key:0,"is-collapse":e.isCollapse,theme:e.theme,"router-info":e.routerInfo},{default:l((function(){return[e.routerInfo.children&&e.routerInfo.children.length?(i(!0),c(a,{key:0},s(e.routerInfo.children,(function(n){return i(),f(r,{key:n.name,"is-collapse":!1,"router-info":n,theme:e.theme},null,8,["router-info","theme"])})),128)):d("",!0)]})),_:1},8,["is-collapse","theme","router-info"]))}}}))}}}));
