/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{g as a}from"./networkChain.845c38c0.js";import{a as s}from"./format.fe46b3b1.js";import{V as l,r as e,g as i,o as t,c as v,b as n,d,w as c,h as r,F as o,v as m,l as u,y as h,G as g,t as C}from"./index.32af5906.js";import"./date.23f5a973.js";import"./dictionary.28290efa.js";import"./sysDictionary.b7a85079.js";const I={class:"starTraceDetail"},f={class:"titleBox"},k=n("div",{class:"title"},"网络状态追溯",-1),w={class:"con"},M={class:"left"},b=n("br",null,null,-1),p={class:"right"},y=n("div",{class:"sub-title"},"网络状态信息",-1),j={class:"starChainMsgBox"},x={class:"starChainMsgItem"},T=n("div",{class:"name"},"卫星编码：",-1),D={class:"val"},B={class:"starChainMsgItem"},S=n("div",{class:"name"},"卫星名称：",-1),_={class:"val"},N={class:"starChainMsgItem"},q=n("div",{class:"name"},"运行轨道：",-1),F={class:"val"},G={class:"starChainMsgItem"},H=n("div",{class:"name"},"网络状态：",-1),O={class:"val"},R={class:"starChainMsgItem"},V=n("div",{class:"name"},"所在网段：",-1),z={class:"val"},A={class:"starChainMsgItem"},E=n("div",{class:"name"},"网络带宽：",-1),J={class:"val"},K={class:"starChainMsgItem singRow"},L=n("div",{class:"name"},"更新时间：",-1),P={class:"val"},Q={class:"starChainMsgItem"},U=n("div",{class:"name"},"所属业务链：",-1),W={class:"val"},X={class:"starChainMsgItem"},Y=n("div",{class:"name"},"交易ID：",-1),Z={class:"val"},$={class:"starChainMsgItem"},aa=n("div",{class:"name"},"上链时间：",-1),sa={class:"val"},la={class:"starChainMsgItem"},ea=n("div",{class:"name"},"块高：",-1),ia={class:"val"},ta={name:"networkChainDetail"},va=Object.assign(ta,{setup(ta){const va=l(),na=e(va.query.id),da=e([]),ca=e({});(async()=>{const s=await a({satelliteId:na.value});200===s.code&&s.data&&s.data.length>0&&(da.value=s.data,ca.value=s.data[0])})();return(a,l)=>{const e=i("el-button"),ta=i("el-timeline-item"),va=i("el-timeline");return t(),v("div",I,[n("div",f,[d(e,{type:"info",icon:"arrow-left",onClick:l[0]||(l[0]=a=>{return s="networkChain",void C.push({name:s});var s})},{default:c((()=>[r("返回")])),_:1}),k]),n("div",w,[n("div",M,[d(va,null,{default:c((()=>[(t(!0),v(o,null,m(da.value,((a,l)=>(t(),u(ta,{key:l,color:l==da.value.length-1?"blue":"green",onClick:s=>(a=>{if(console.log(a,"id"),da.value.length>0){let s=da.value.filter((s=>s.id===a));s.length>0&&(ca.value=s[0])}})(a.id)},{default:c((()=>[r(h(g(s)(a.lastTime,1))+" ",1),b,r(" 【"+h(a.satelliteId)+"】"+h(a.satelliteName)+" 网络状态："+h(a.networkState),1)])),_:2},1032,["color","onClick"])))),128))])),_:1})]),n("div",p,[y,n("div",j,[n("div",x,[T,n("div",D,h(ca.value.satelliteId),1)]),n("div",B,[S,n("div",_,h(ca.value.satelliteName),1)]),n("div",N,[q,n("div",F,h(ca.value.orbitId),1)]),n("div",G,[H,n("div",O,h(ca.value.networkState),1)]),n("div",R,[V,n("div",z,h(ca.value.networkSegment),1)]),n("div",A,[E,n("div",J,h(ca.value.networkBandwidth),1)]),n("div",K,[L,n("div",P,h(g(s)(ca.value.lastTime,1)),1)]),n("div",Q,[U,n("div",W,h(ca.value.chainId),1)]),n("div",X,[Y,n("div",Z,h(ca.value.txId),1)]),n("div",$,[aa,n("div",sa,h(ca.value.chainTime?g(s)(ca.value.chainTime):""),1)]),n("div",la,[ea,n("div",ia,h(ca.value.blockHeight),1)])])])])])}}});export{va as default};
