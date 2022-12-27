/*! 
 Build based on gin-vue-admin 
 Time : 1672137460000 */
import{s as l,r as e,g as a,o as t,c as n,b as u,d as i,w as o,h as s,y as r,G as d,e as c}from"./index.326c508a.js";import{a as v}from"./format.140daded.js";import"./date.23f5a973.js";import"./dictionary.5509472a.js";import"./sysDictionary.08a35ed5.js";const p={class:"gva-search-box"},m={class:"gva-table-box"},g={key:0},h={class:"gva-pagination"},f={class:"starTableBox"},b=u("th",null,"指令编码:",-1),y=u("th",null,"指令类型:",-1),w=u("th",null,"指令来源:",-1),C=u("th",null,"指令生成时间:",-1),I=u("th",null,"碎片编号:",-1),_=u("th",null,"碎片名称:",-1),x=u("th",null,"卫星编码:",-1),z=u("th",null,"卫星名称:",-1),T=u("th",null,"指令内容:",-1),j={colspan:"3"},k=u("th",null,"所属业务链:",-1),V=u("th",null,"交易ID:",-1),S=u("th",null,"上链时间:",-1),N=u("th",null,"区块高度:",-1),D={class:"dialog-footer"},U={name:"commandReception"},B=Object.assign(U,{setup(U){const B=e({}),F=e(1),G=e(0),H=e(10),O=e([]),R=e({}),q=()=>{F.value=1,H.value=10,J()},A=l=>{H.value=l,J()},E=l=>{F.value=l,J()},J=async()=>{let e={page:F.value,pageSize:H.value,...R.value};const a=await(e=>l({url:"/satellitebc/exec/getinstructionlist",method:"get",params:e}))(e);200===a.code&&(O.value=a.data,G.value=a.total)};J();const K=e("指令详情"),L=e(!1),M=()=>{B.value={},L.value=!1},P=async e=>{const a=await(t={id:e.id},l({url:"/satellitebc/exec/getinstruction",method:"get",params:t}));var t;B.value=a.data,L.value=!0};return(l,e)=>{const U=a("el-input"),J=a("el-form-item"),Q=a("el-button"),W=a("el-form"),X=a("el-table-column"),Y=a("el-table"),Z=a("el-pagination"),$=a("el-dialog");return t(),n("div",null,[u("div",p,[i(W,{ref:"searchForm",inline:!0,model:R.value},{default:o((()=>[i(J,{label:"指令信息"},{default:o((()=>[i(U,{modelValue:R.value.searchConditions,"onUpdate:modelValue":e[0]||(e[0]=l=>R.value.searchConditions=l),placeholder:"指令信息"},null,8,["modelValue"])])),_:1}),i(J,null,{default:o((()=>[i(Q,{size:"small",type:"primary",icon:"search",onClick:q},{default:o((()=>[s("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),u("div",m,[i(Y,{data:O.value},{default:o((()=>[i(X,{align:"left",label:"指令编码",prop:"instructionId"}),i(X,{align:"left",label:"指令类型",prop:"instructionType"}),i(X,{align:"left",label:"指令内容",prop:"instructionContent"}),i(X,{label:"指令生成时间",width:"180"},{default:o((l=>[s(r(d(v)(l.row.genInstructionTime)),1)])),_:1}),i(X,{align:"left",label:"指令来源",prop:"instructionSource"}),i(X,{align:"left",label:"碎片编号",prop:"debrisId"}),i(X,{align:"left",label:"卫星编码",prop:"satelliteId"}),i(X,{align:"left",label:"卫星名称",prop:"satelliteName"}),i(X,{label:"上链时间",width:"180"},{default:o((l=>[l.row.chainTime?(t(),n("div",g,r(d(v)(l.row.chainTime)),1)):c("",!0)])),_:1}),i(X,{align:"left",fixed:"right",label:"操作",width:"80"},{default:o((l=>[i(Q,{icon:"view",size:"small",type:"primary",link:"",onClick:e=>P(l.row)},{default:o((()=>[s("查看")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),u("div",h,[i(Z,{"current-page":F.value,"page-size":H.value,"page-sizes":[10,30,50,100],total:G.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:E,onSizeChange:A},null,8,["current-page","page-size","total"])])]),i($,{modelValue:L.value,"onUpdate:modelValue":e[1]||(e[1]=l=>L.value=l),"before-close":M,title:K.value,width:820},{footer:o((()=>[u("div",D,[i(Q,{size:"small",onClick:M},{default:o((()=>[s("关闭")])),_:1})])])),default:o((()=>[u("div",f,[u("table",null,[u("tr",null,[b,u("td",null,r(B.value.instructionId),1),y,u("td",null,r(B.value.instructionType),1)]),u("tr",null,[w,u("td",null,r(B.value.instructionSource),1),C,u("td",null,r(d(v)(B.value.genInstructionTime)),1)]),u("tr",null,[I,u("td",null,r(B.value.debrisId),1),_,u("td",null,r(B.value.debrisName),1)]),u("tr",null,[x,u("td",null,r(B.value.satelliteId),1),z,u("td",null,r(B.value.satelliteName),1)]),u("tr",null,[T,u("td",j,r(B.value.instructionContent),1)]),u("tr",null,[k,u("td",null,r(B.value.chainId),1),V,u("td",null,r(B.value.txId),1)]),u("tr",null,[S,u("td",null,r(B.value.chainTime?d(v)(B.value.chainTime):""),1),N,u("td",null,r(B.value.blockHeight),1)])])])])),_:1},8,["modelValue","title"])])}}});export{B as default};
