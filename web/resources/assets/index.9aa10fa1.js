/*! 
 Build based on gin-vue-admin 
 Time : 1672137460000 */
import{a}from"./faultChain.312f8d38.js";import{r as e,g as l,o as t,c as i,b as o,d as s,w as n,h as r,y as u,G as d,e as p,t as f}from"./index.326c508a.js";import{a as m}from"./format.140daded.js";import"./date.23f5a973.js";import"./dictionary.5509472a.js";import"./sysDictionary.08a35ed5.js";const c={class:"gva-search-box"},g={class:"gva-table-box"},v={key:0},b={class:"gva-pagination"},h={name:"faultChain"},y=Object.assign(h,{setup(h){const y=e(1),C=e(0),_=e(10),j=e([]),w=e({}),z=()=>{y.value=1,_.value=10,T()},x=a=>{_.value=a,T()},k=a=>{y.value=a,T()},T=async()=>{let e={page:y.value,pageSize:_.value,...w.value};const l=await a(e);200===l.code&&(j.value=l.data,C.value=l.total)};T();return(a,e)=>{const h=l("el-input"),T=l("el-form-item"),D=l("el-button"),I=l("el-form"),S=l("el-table-column"),V=l("el-table"),q=l("el-pagination");return t(),i("div",null,[o("div",c,[s(I,{ref:"searchForm",inline:!0,model:w.value},{default:n((()=>[s(T,{label:"卫星名称"},{default:n((()=>[s(h,{modelValue:w.value.searchConditions,"onUpdate:modelValue":e[0]||(e[0]=a=>w.value.searchConditions=a),placeholder:"卫星名称"},null,8,["modelValue"])])),_:1}),s(T,null,{default:n((()=>[s(D,{size:"small",type:"primary",icon:"search",onClick:z},{default:n((()=>[r("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),o("div",g,[s(V,{data:j.value},{default:n((()=>[s(S,{align:"left",label:"卫星编码",prop:"satelliteId"}),s(S,{align:"left",label:"卫星名称",prop:"satelliteName"}),s(S,{align:"left",label:"运行轨道",prop:"orbitId"}),s(S,{align:"left",label:"故障类型",prop:"faultType"}),s(S,{align:"left",label:"故障时间"},{default:n((a=>[r(u(d(m)(a.row.faultTime)),1)])),_:1}),s(S,{align:"left",label:"故障描述",prop:"faultDescription"}),s(S,{align:"left",label:"故障修复状态",prop:"repairState"}),s(S,{label:"上链时间",width:"180"},{default:n((a=>[a.row.chainTime?(t(),i("div",v,u(d(m)(a.row.chainTime)),1)):p("",!0)])),_:1}),s(S,{align:"left",fixed:"right",label:"操作",width:"80"},{default:n((a=>[s(D,{size:"small",type:"primary",link:"",onClick:e=>(a=>{let e=a.satelliteId;f.push({name:"faultChainDetail",query:{id:e}})})(a.row)},{default:n((()=>[r("追溯")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),o("div",b,[s(q,{"current-page":y.value,"page-size":_.value,"page-sizes":[10,30,50,100],total:C.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:k,onSizeChange:x},null,8,["current-page","page-size","total"])])])])}}});export{y as default};