/*! 
 Build based on gin-vue-admin 
 Time : 1671785957000 */
import{a}from"./avoidChain.caaed5a2.js";import{r as e,g as l,o as t,c as i,b as o,d as n,w as r,h as s,y as d,G as u,e as p,t as c}from"./index.860b48e6.js";import{a as m}from"./format.4fc21eae.js";import"./date.23f5a973.js";import"./dictionary.8dcfcd04.js";import"./sysDictionary.12ad4d06.js";const v={class:"gva-search-box"},f={class:"gva-table-box"},g={key:0},h={class:"gva-pagination"},b={name:"avoidChain"},y=Object.assign(b,{setup(b){const y=e(1),C=e(0),w=e(10),_=e([]),j=e({}),z=()=>{y.value=1,w.value=10,I()},x=a=>{w.value=a,I()},k=a=>{y.value=a,I()},I=async()=>{let e={page:y.value,pageSize:w.value,...j.value};const l=await a(e);200===l.code&&(_.value=l.data,C.value=l.total)};I();return(a,e)=>{const b=l("el-input"),I=l("el-form-item"),T=l("el-button"),V=l("el-form"),D=l("el-table-column"),S=l("el-table"),q=l("el-pagination");return t(),i("div",null,[o("div",v,[n(V,{ref:"searchForm",inline:!0,model:j.value},{default:r((()=>[n(I,{label:"卫星名称/卫星编码"},{default:r((()=>[n(b,{modelValue:j.value.searchConditions,"onUpdate:modelValue":e[0]||(e[0]=a=>j.value.searchConditions=a),placeholder:"卫星名称/卫星编码"},null,8,["modelValue"])])),_:1}),n(I,null,{default:r((()=>[n(T,{size:"small",type:"primary",icon:"search",onClick:z},{default:r((()=>[s("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),o("div",f,[n(S,{data:_.value},{default:r((()=>[n(D,{align:"left",label:"操作者",prop:"operator"}),n(D,{label:"操作日期",width:"180"},{default:r((a=>[s(d(u(m)(a.row.operationTime)),1)])),_:1}),n(D,{align:"left",label:"IP",prop:"operationIp"}),n(D,{align:"left",label:"操作记录",prop:"operationRecord"}),n(D,{align:"left",label:"卫星编码",prop:"satelliteId"}),n(D,{align:"left",label:"卫星名称",prop:"satelliteName"}),n(D,{label:"上链时间",width:"180"},{default:r((a=>[a.row.chainTime?(t(),i("div",g,d(u(m)(a.row.chainTime)),1)):p("",!0)])),_:1}),n(D,{align:"left",fixed:"right",label:"操作",width:"80"},{default:r((a=>[n(T,{size:"small",type:"primary",link:"",onClick:e=>(a=>{let e=a.satelliteId;c.push({name:"avoidChainDetail",query:{id:e}})})(a.row)},{default:r((()=>[s("追溯")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),o("div",h,[n(q,{"current-page":y.value,"page-size":w.value,"page-sizes":[10,30,50,100],total:C.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:k,onSizeChange:x},null,8,["current-page","page-size","total"])])])])}}});export{y as default};
