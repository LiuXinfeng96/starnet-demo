/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{a}from"./communicationChain.1c56808f.js";import{r as e,g as l,o as t,c as i,b as o,d as n,w as s,h as r,y as u,G as m,e as p,t as d}from"./index.32af5906.js";import{a as c}from"./format.fe46b3b1.js";import"./date.23f5a973.js";import"./dictionary.28290efa.js";import"./sysDictionary.b7a85079.js";const f={class:"gva-search-box"},g={class:"gva-table-box"},v={key:0},b={class:"gva-pagination"},h={name:"communicationChain"},y=Object.assign(h,{setup(h){const y=e(1),C=e(0),w=e(10),_=e([]),j=e({}),z=()=>{y.value=1,w.value=10,D()},k=a=>{w.value=a,D()},x=a=>{y.value=a,D()},D=async()=>{let e={page:y.value,pageSize:w.value,...j.value};const l=await a(e);200===l.code&&(_.value=l.data,C.value=l.total)};D();return(a,e)=>{const h=l("el-input"),D=l("el-form-item"),I=l("el-button"),S=l("el-form"),T=l("el-table-column"),V=l("el-table"),q=l("el-pagination");return t(),i("div",null,[o("div",f,[n(S,{ref:"searchForm",inline:!0,model:j.value},{default:s((()=>[n(D,{label:"卫星名称"},{default:s((()=>[n(h,{modelValue:j.value.searchConditions,"onUpdate:modelValue":e[0]||(e[0]=a=>j.value.searchConditions=a),placeholder:"卫星名称"},null,8,["modelValue"])])),_:1}),n(D,null,{default:s((()=>[n(I,{size:"small",type:"primary",icon:"search",onClick:z},{default:s((()=>[r("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),o("div",g,[n(V,{data:_.value},{default:s((()=>[n(T,{align:"left",label:"卫星编码",prop:"satelliteId"}),n(T,{align:"left",label:"卫星名称",prop:"satelliteName"}),n(T,{align:"left",label:"运行轨道",prop:"orbitId"}),n(T,{align:"left",label:"通信状态",prop:"commState"}),n(T,{align:"left",label:"通信端口",prop:"commPort"}),n(T,{align:"left",label:"通信带宽",prop:"commBandwidth"}),n(T,{align:"left",label:"延时",prop:"commDelay"}),n(T,{align:"left",label:"链路负荷",prop:"linkLoad"}),n(T,{label:"更新时间",width:"180"},{default:s((a=>[r(u(m(c)(a.row.lastTime,1)),1)])),_:1}),n(T,{label:"上链时间",width:"180"},{default:s((a=>[a.row.chainTime?(t(),i("div",v,u(m(c)(a.row.chainTime)),1)):p("",!0)])),_:1}),n(T,{align:"left",fixed:"right",label:"操作",width:"80"},{default:s((a=>[n(I,{size:"small",type:"primary",link:"",onClick:e=>(a=>{let e=a.satelliteId;d.push({name:"communicationChainDetail",query:{id:e}})})(a.row)},{default:s((()=>[r("追溯")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),o("div",b,[n(q,{"current-page":y.value,"page-size":w.value,"page-sizes":[10,30,50,100],total:C.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:x,onSizeChange:k},null,8,["current-page","page-size","total"])])])])}}});export{y as default};
