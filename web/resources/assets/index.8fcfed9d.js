/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{a}from"./starsignChain.ea1600d2.js";import{r as e,g as l,o as t,c as i,b as s,d as o,w as n,h as r,y as u,G as p,e as d,t as m}from"./index.32af5906.js";import{a as c}from"./format.fe46b3b1.js";import"./date.23f5a973.js";import"./dictionary.28290efa.js";import"./sysDictionary.b7a85079.js";const v={class:"gva-search-box"},f={class:"gva-table-box"},g={key:0},b={class:"gva-pagination"},h={name:"starsignChain"},y=Object.assign(h,{setup(h){const y=e(1),C=e(0),j=e(10),w=e([]),z=e({}),_=()=>{y.value=1,j.value=10,N()},k=a=>{j.value=a,N()},x=a=>{y.value=a,N()},N=async()=>{let e={page:y.value,pageSize:j.value,...z.value};const l=await a(e);200===l.code&&(w.value=l.data,C.value=l.total)};N();return(a,e)=>{const h=l("el-input"),N=l("el-form-item"),D=l("el-button"),S=l("el-form"),T=l("el-table-column"),V=l("el-table"),I=l("el-pagination");return t(),i("div",null,[s("div",v,[o(S,{ref:"searchForm",inline:!0,model:z.value},{default:n((()=>[o(N,{label:"星座名称"},{default:n((()=>[o(h,{modelValue:z.value.searchConditions,"onUpdate:modelValue":e[0]||(e[0]=a=>z.value.searchConditions=a),placeholder:"星座名称"},null,8,["modelValue"])])),_:1}),o(N,null,{default:n((()=>[o(D,{size:"small",type:"primary",icon:"search",onClick:_},{default:n((()=>[r("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),s("div",f,[o(V,{data:w.value},{default:n((()=>[o(T,{label:"星座名称",prop:"constellationName"}),o(T,{label:"星座编码",prop:"constellationId"}),o(T,{label:"包含卫星数",prop:"satelliteTotalNum"}),o(T,{label:"正常卫星数量",prop:"satelliteUpNum"}),o(T,{label:"异常卫星数量",prop:"satelliteDownNum"}),o(T,{label:"星间链路状态",prop:"satelliteLinkState"}),o(T,{label:"上链时间",width:"180"},{default:n((a=>[a.row.chainTime?(t(),i("div",g,u(p(c)(a.row.chainTime)),1)):d("",!0)])),_:1}),o(T,{align:"left",fixed:"right",label:"操作",width:"80"},{default:n((a=>[o(D,{size:"small",type:"primary",link:"",onClick:e=>(a=>{let e=a.constellationId;m.push({name:"starsignChainDetail",query:{id:e}})})(a.row)},{default:n((()=>[r("追溯")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),s("div",b,[o(I,{"current-page":y.value,"page-size":j.value,"page-sizes":[10,30,50,100],total:C.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:x,onSizeChange:k},null,8,["current-page","page-size","total"])])])])}}});export{y as default};
