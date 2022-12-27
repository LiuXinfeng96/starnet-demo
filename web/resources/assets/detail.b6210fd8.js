/*! 
 Build based on gin-vue-admin 
 Time : 1672137460000 */
import{g as a}from"./avoidChain.94ab817d.js";import{a as s}from"./format.140daded.js";import{V as l,r as i,g as e,o as t,c as v,b as d,d as c,w as n,h as o,F as r,v as m,l as u,y as h,G as g,t as p}from"./index.326c508a.js";import"./date.23f5a973.js";import"./dictionary.5509472a.js";import"./sysDictionary.08a35ed5.js";const C={class:"starTraceDetail"},I={class:"titleBox"},f=d("div",{class:"title"},"避障操作追溯",-1),M={class:"con"},y={class:"left"},b={class:"right"},j=d("div",{class:"sub-title"},"避障操作信息",-1),k={class:"starChainMsgBox"},x={class:"starChainMsgItem"},w=d("div",{class:"name"},"卫星编码：",-1),D={class:"val"},T={class:"starChainMsgItem"},R=d("div",{class:"name"},"卫星名称：",-1),_={class:"val"},B={class:"starChainMsgItem"},q=d("div",{class:"name"},"操作IP：",-1),F={class:"val"},G={class:"starChainMsgItem"},H=d("div",{class:"name"},"操作者：",-1),N={class:"val"},O={class:"starChainMsgItem singRow"},P=d("div",{class:"name"},"操作记录：",-1),V={class:"val"},z={class:"starChainMsgItem"},A=d("div",{class:"name"},"所属业务链：",-1),E={class:"val"},J={class:"starChainMsgItem"},K=d("div",{class:"name"},"交易ID：",-1),L={class:"val"},Q={class:"starChainMsgItem"},S=d("div",{class:"name"},"上链时间：",-1),U={class:"val"},W={class:"starChainMsgItem"},X=d("div",{class:"name"},"块高：",-1),Y={class:"val"},Z={name:"avoidChainDetail"},$=Object.assign(Z,{setup(Z){const $=l(),aa=i($.query.id),sa=i([]),la=i({});(async()=>{const s=await a({satelliteId:aa.value});200===s.code&&s.data&&s.data.length>0&&(sa.value=s.data,la.value=s.data[0])})();return(a,l)=>{const i=e("el-button"),Z=e("el-timeline-item"),$=e("el-timeline");return t(),v("div",C,[d("div",I,[c(i,{type:"info",icon:"arrow-left",onClick:l[0]||(l[0]=a=>{return s="avoidChain",void p.push({name:s});var s})},{default:n((()=>[o("返回")])),_:1}),f]),d("div",M,[d("div",y,[c($,null,{default:n((()=>[(t(!0),v(r,null,m(sa.value,((a,l)=>(t(),u(Z,{key:l,color:l==sa.value.length-1?"blue":"green",onClick:s=>(a=>{if(console.log(a,"id"),sa.value.length>0){let s=sa.value.filter((s=>s.id===a));s.length>0&&(la.value=s[0])}})(a.id)},{default:n((()=>[o(h(g(s)(a.lastTime,1))+" "+h(a.operator)+"用户进行了“"+h(a.operationRecord)+"”操作 ",1)])),_:2},1032,["color","onClick"])))),128))])),_:1})]),d("div",b,[j,d("div",k,[d("div",x,[w,d("div",D,h(la.value.satelliteId),1)]),d("div",T,[R,d("div",_,h(la.value.satelliteName),1)]),d("div",B,[q,d("div",F,h(la.value.operationIp),1)]),d("div",G,[H,d("div",N,h(la.value.operator),1)]),d("div",O,[P,d("div",V,h(la.value.operationRecord),1)]),d("div",z,[A,d("div",E,h(la.value.chainId),1)]),d("div",J,[K,d("div",L,h(la.value.txId),1)]),d("div",Q,[S,d("div",U,h(la.value.chainTime?g(s)(la.value.chainTime):""),1)]),d("div",W,[X,d("div",Y,h(la.value.blockHeight),1)])])])])])}}});export{$ as default};