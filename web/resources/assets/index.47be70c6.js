/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{s as a,r as e,g as l,o as t,c as o,b as s,d as i,w as r,h as n,y as u,G as p,e as d}from"./index.32af5906.js";import{a as c}from"./format.fe46b3b1.js";import"./date.23f5a973.js";import"./dictionary.28290efa.js";import"./sysDictionary.b7a85079.js";const m={class:"gva-search-box"},v={class:"gva-table-box"},g={key:0},b={class:"gva-pagination"},f={name:"starAvoidLog"},h=Object.assign(f,{setup(f){const h=e(1),y=e(0),j=e(10),w=e([]),z=e({}),_=()=>{h.value=1,j.value=10,T()},x=a=>{j.value=a,T()},C=a=>{h.value=a,T()},T=async()=>{let e={page:h.value,pageSize:j.value,...z.value};const l=await(e=>a({url:"/satellitebc/exec/getoperationlist",method:"get",params:e}))(e);200===l.code&&(w.value=l.data,y.value=l.total)};return T(),(a,e)=>{const f=l("el-input"),T=l("el-form-item"),V=l("el-button"),k=l("el-form"),S=l("el-table-column"),A=l("el-table"),D=l("el-pagination");return t(),o("div",null,[s("div",m,[i(k,{ref:"searchForm",inline:!0,model:z.value},{default:r((()=>[i(T,{label:"卫星名称/卫星编码"},{default:r((()=>[i(f,{modelValue:z.value.searchConditions,"onUpdate:modelValue":e[0]||(e[0]=a=>z.value.searchConditions=a),placeholder:"卫星名称/卫星编码"},null,8,["modelValue"])])),_:1}),i(T,null,{default:r((()=>[i(V,{size:"small",type:"primary",icon:"search",onClick:_},{default:r((()=>[n("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),s("div",v,[i(A,{data:w.value},{default:r((()=>[i(S,{label:"操作者",prop:"operator"}),i(S,{label:"操作日期",width:"180"},{default:r((a=>[n(u(p(c)(a.row.operationTime)),1)])),_:1}),i(S,{label:"操作记录",prop:"operationRecord"}),i(S,{label:"卫星编码",prop:"satelliteId"}),i(S,{label:"卫星名称",prop:"satelliteName"}),i(S,{label:"上链时间",width:"180"},{default:r((a=>[a.row.chainTime?(t(),o("div",g,u(p(c)(a.row.chainTime)),1)):d("",!0)])),_:1})])),_:1},8,["data"]),s("div",b,[i(D,{"current-page":h.value,"page-size":j.value,"page-sizes":[10,30,50,100],total:y.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:C,onSizeChange:x},null,8,["current-page","page-size","total"])])])])}}});export{h as default};