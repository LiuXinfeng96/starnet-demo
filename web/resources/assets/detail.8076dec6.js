/*! 
 Build based on gin-vue-admin 
 Time : 1672137460000 */
import{g as a}from"./trackChain.0eb1a064.js";import{a as s}from"./format.140daded.js";import{V as l,r as e,g as i,o as t,c as v,b as d,d as c,w as n,h as r,F as m,v as o,l as u,y as h,G as g,t as C}from"./index.326c508a.js";import"./date.23f5a973.js";import"./dictionary.5509472a.js";import"./sysDictionary.08a35ed5.js";const I={class:"starTraceDetail"},f={class:"titleBox"},M=d("div",{class:"title"},"卫星在轨状态追溯",-1),p={class:"con"},k={class:"left"},y={class:"right"},b=d("div",{class:"sub-title"},"轨道状态信息",-1),j={class:"starChainMsgBox"},x={class:"starChainMsgItem"},T=d("div",{class:"name"},"卫星编码：",-1),D={class:"val"},w={class:"starChainMsgItem"},_=d("div",{class:"name"},"卫星名称：",-1),B={class:"val"},S={class:"starChainMsgItem"},q=d("div",{class:"name"},"所在轨道编码：",-1),A={class:"val"},F={class:"starChainMsgItem"},G=d("div",{class:"name"},"运行状态：",-1),H={class:"val"},N={class:"starChainMsgItem"},O=d("div",{class:"name"},"平近点角：",-1),V={class:"val"},z={class:"starChainMsgItem"},E=d("div",{class:"name"},"速度（km/s）：",-1),J={class:"val"},K={class:"starChainMsgItem"},L=d("div",{class:"name"},"更新时间：",-1),P={class:"val"},Q={class:"starChainMsgItem"},R=d("div",{class:"name"},"所属业务链：",-1),U={class:"val"},W={class:"starChainMsgItem"},X=d("div",{class:"name"},"交易ID：",-1),Y={class:"val"},Z={class:"starChainMsgItem"},$=d("div",{class:"name"},"上链时间：",-1),aa={class:"val"},sa={class:"starChainMsgItem"},la=d("div",{class:"name"},"块高：",-1),ea={class:"val"},ia={name:"trackChainDetail"},ta=Object.assign(ia,{setup(ia){const ta=l(),va=e(ta.query.id),da=e([]),ca=e({});(async()=>{const s=await a({satelliteId:va.value});200===s.code&&s.data&&s.data.length>0&&(da.value=s.data,ca.value=s.data[0])})();return(a,l)=>{const e=i("el-button"),ia=i("el-timeline-item"),ta=i("el-timeline");return t(),v("div",I,[d("div",f,[c(e,{type:"info",icon:"arrow-left",onClick:l[0]||(l[0]=a=>{return s="trackChain",void C.push({name:s});var s})},{default:n((()=>[r("返回")])),_:1}),M]),d("div",p,[d("div",k,[c(ta,null,{default:n((()=>[(t(!0),v(m,null,o(da.value,((a,l)=>(t(),u(ia,{key:l,color:l==da.value.length-1?"blue":"green",onClick:s=>(a=>{if(console.log(a,"id"),da.value.length>0){let s=da.value.filter((s=>s.id===a));s.length>0&&(ca.value=s[0])}})(a.id)},{default:n((()=>[r(h(g(s)(a.lastTime,1))+" 状态："+h(a.runState),1)])),_:2},1032,["color","onClick"])))),128))])),_:1})]),d("div",y,[b,d("div",j,[d("div",x,[T,d("div",D,h(ca.value.satelliteId),1)]),d("div",w,[_,d("div",B,h(ca.value.satelliteName),1)]),d("div",S,[q,d("div",A,h(ca.value.orbitId),1)]),d("div",F,[G,d("div",H,h(ca.value.runState),1)]),d("div",N,[O,d("div",V,h(ca.value.meanAnomaly),1)]),d("div",z,[E,d("div",J,h(ca.value.speed),1)]),d("div",K,[L,d("div",P,h(g(s)(ca.value.lastTime,1)),1)]),d("div",Q,[R,d("div",U,h(ca.value.chainId),1)]),d("div",W,[X,d("div",Y,h(ca.value.txId),1)]),d("div",Z,[$,d("div",aa,h(ca.value.chainTime?g(s)(ca.value.chainTime):""),1)]),d("div",sa,[la,d("div",ea,h(ca.value.blockHeight),1)])])])])])}}});export{ta as default};