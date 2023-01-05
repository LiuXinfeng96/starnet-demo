/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{g as a}from"./fragmentChain.c2872f2e.js";import{a as s}from"./format.fe46b3b1.js";import{V as l,r as e,g as i,o as t,c as v,b as d,d as c,w as n,h as r,F as m,v as o,l as u,y as g,G as h,t as f}from"./index.32af5906.js";import"./date.23f5a973.js";import"./dictionary.28290efa.js";import"./sysDictionary.b7a85079.js";const C={class:"starTraceDetail"},I={class:"titleBox"},b=d("div",{class:"title"},"碎片信息追溯",-1),M={class:"con"},p={class:"left"},y={class:"right"},j=d("div",{class:"sub-title"},"碎片信息",-1),k={class:"starChainMsgBox"},x={class:"starChainMsgItem"},T=d("div",{class:"name"},"碎片编号：",-1),w={class:"val"},D={class:"starChainMsgItem"},_=d("div",{class:"name"},"碎片名称：",-1),B={class:"val"},N={class:"starChainMsgItem"},S=d("div",{class:"name"},"碎片信息来源：",-1),q={class:"val"},F={class:"starChainMsgItem"},G=d("div",{class:"name"},"运行角度：",-1),H={class:"val"},O={class:"starChainMsgItem"},R=d("div",{class:"name"},"运行速度：",-1),V={class:"val"},z={class:"starChainMsgItem"},A=d("div",{class:"name"},"运行高度：",-1),E={class:"val"},J={class:"starChainMsgItem"},K=d("div",{class:"name"},"碎片体积：",-1),L={class:"val"},P={class:"starChainMsgItem"},Q=d("div",{class:"name"},"碎片类型：",-1),U={class:"val"},W={class:"starChainMsgItem singRow"},X=d("div",{class:"name"},"更新时间：",-1),Y={class:"val"},Z={class:"starChainMsgItem"},$=d("div",{class:"name"},"所属业务链：",-1),aa={class:"val"},sa={class:"starChainMsgItem"},la=d("div",{class:"name"},"交易ID：",-1),ea={class:"val"},ia={class:"starChainMsgItem"},ta=d("div",{class:"name"},"上链时间：",-1),va={class:"val"},da={class:"starChainMsgItem"},ca=d("div",{class:"name"},"块高：",-1),na={class:"val"},ra={name:"fragmentChainDetail"},ma=Object.assign(ra,{setup(ra){const ma=l(),oa=e(ma.query.id),ua=e([]),ga=e({});(async()=>{const s=await a({debrisId:oa.value});200===s.code&&s.data&&s.data.length>0&&(ua.value=s.data,ga.value=s.data[0])})();return(a,l)=>{const e=i("el-button"),ra=i("el-timeline-item"),ma=i("el-timeline");return t(),v("div",C,[d("div",I,[c(e,{type:"info",icon:"arrow-left",onClick:l[0]||(l[0]=a=>{return s="fragmentChain",void f.push({name:s});var s})},{default:n((()=>[r("返回")])),_:1}),b]),d("div",M,[d("div",p,[c(ma,null,{default:n((()=>[(t(!0),v(m,null,o(ua.value,((a,l)=>(t(),u(ra,{key:l,color:l==ua.value.length-1?"blue":"green",onClick:s=>(a=>{if(console.log(a,"id"),ua.value.length>0){let s=ua.value.filter((s=>s.id===a));s.length>0&&(ga.value=s[0])}})(a.id)},{default:n((()=>[r(g(h(s)(a.lastTime,1))+" “"+g(a.debrisSource)+"”更新“"+g(a.debrisName)+"”碎片信息 ",1)])),_:2},1032,["color","onClick"])))),128))])),_:1})]),d("div",y,[j,d("div",k,[d("div",x,[T,d("div",w,g(ga.value.debrisId),1)]),d("div",D,[_,d("div",B,g(ga.value.debrisName),1)]),d("div",N,[S,d("div",q,g(ga.value.debrisSource),1)]),d("div",F,[G,d("div",H,g(ga.value.angle),1)]),d("div",O,[R,d("div",V,g(ga.value.speed),1)]),d("div",z,[A,d("div",E,g(ga.value.height),1)]),d("div",J,[K,d("div",L,g(ga.value.volume),1)]),d("div",P,[Q,d("div",U,g(ga.value.type),1)]),d("div",W,[X,d("div",Y,g(h(s)(ga.value.lastTime,1)),1)]),d("div",Z,[$,d("div",aa,g(ga.value.chainId),1)]),d("div",sa,[la,d("div",ea,g(ga.value.txId),1)]),d("div",ia,[ta,d("div",va,g(ga.value.chainTime?h(s)(ga.value.chainTime):""),1)]),d("div",da,[ca,d("div",na,g(ga.value.blockHeight),1)])])])])])}}});export{ma as default};
