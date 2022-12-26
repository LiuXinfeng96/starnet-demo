/*! 
 Build based on gin-vue-admin 
 Time : 1671785957000 */
import{s as l,r as e,g as a,o as t,c as u,b as n,d as i,w as r,h as o,y as s,G as d,e as c}from"./index.860b48e6.js";import{a as p}from"./format.4fc21eae.js";import"./date.23f5a973.js";import"./dictionary.8dcfcd04.js";import"./sysDictionary.12ad4d06.js";const m={class:"gva-search-box"},v={class:"gva-table-box"},g={key:0},h={class:"gva-pagination"},f={class:"starTableBox"},b=n("th",null,"指令编码:",-1),y=n("th",null,"指令类型:",-1),x=n("th",null,"指令生成时间:",-1),T=n("th",null,"执行时间:",-1),I=n("th",null,"指令来源:",-1),_=n("th",null,"执行状态:",-1),w=n("th",null,"碎片编号:",-1),C=n("th",null,"碎片名称:",-1),z=n("th",null,"卫星编码:",-1),j=n("th",null,"卫星名称:",-1),k=n("th",null,"指令内容:",-1),S={colspan:"3"},V=n("th",null,"区块链ID:",-1),G=n("th",null,"上链时间:",-1),q=n("th",null,"区块高度:",-1),N={colspan:"3"},D={class:"dialog-footer"},O={name:"commandResult"},U=Object.assign(O,{setup(O){e([]);const U=e({path:"",apiGroup:"",method:"",description:""});e([{value:"POST",label:"创建",type:"success"},{value:"GET",label:"查看",type:""}]),e({path:[{required:!0,message:"请输入api路径",trigger:"blur"}],apiGroup:[{required:!0,message:"请输入组名称",trigger:"blur"}],method:[{required:!0,message:"请选择请求方式",trigger:"blur"}],description:[{required:!0,message:"请输入api介绍",trigger:"blur"}]});const B=e(1),E=e(0),F=e(10),H=e([]),P=e({}),R=()=>{B.value=1,F.value=10,K()},A=l=>{F.value=l,K()},J=l=>{B.value=l,K()},K=async()=>{let e={page:B.value,pageSize:F.value,...P.value};const a=await(e=>l({url:"/satellitebc/control/getexecresultlist",method:"get",params:e}))(e);200===a.code&&(H.value=a.data,E.value=a.total)};K(),e(null);const L=e("指令执行详情"),M=e(!1),Q=()=>{U.value={path:"",apiGroup:"",method:"",description:""},M.value=!1},W=async e=>{const a=await(t={id:e.id},l({url:"/satellitebc/control/getexecresult",method:"get",params:t}));var t;U.value=a.data,M.value=!0};return(l,e)=>{const O=a("el-input"),K=a("el-form-item"),X=a("el-button"),Y=a("el-form"),Z=a("el-table-column"),$=a("el-table"),ll=a("el-pagination"),el=a("el-dialog");return t(),u("div",null,[n("div",m,[i(Y,{ref:"searchForm",inline:!0,model:P.value},{default:r((()=>[i(K,{label:"指令信息"},{default:r((()=>[i(O,{modelValue:P.value.searchConditions,"onUpdate:modelValue":e[0]||(e[0]=l=>P.value.searchConditions=l),placeholder:"指令信息"},null,8,["modelValue"])])),_:1}),i(K,null,{default:r((()=>[i(X,{size:"small",type:"primary",icon:"search",onClick:R},{default:r((()=>[o("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),n("div",v,[i($,{data:H.value},{default:r((()=>[i(Z,{align:"left",label:"指令编码",width:"90",prop:"instructionId"}),i(Z,{align:"left",label:"指令执行时间"},{default:r((l=>[o(s(d(p)(l.row.execInstructionTime)),1)])),_:1}),i(Z,{align:"left",label:"执行结果",prop:"execState"}),i(Z,{align:"left",label:"指令类型",prop:"instructionType"}),i(Z,{align:"left",label:"指令下发时间"},{default:r((l=>[o(s(d(p)(l.row.genInstructionTime)),1)])),_:1}),i(Z,{align:"left",label:"碎片编号",prop:"debrisId"}),i(Z,{align:"left",label:"卫星编码",prop:"satelliteId"}),i(Z,{align:"left",label:"卫星名称",prop:"satelliteName"}),i(Z,{align:"left",label:"上链时间",width:"180"},{default:r((l=>[l.row.chainTime?(t(),u("div",g,s(d(p)(l.row.chainTime)),1)):c("",!0)])),_:1}),i(Z,{align:"left",fixed:"right",label:"操作",width:"80"},{default:r((l=>[i(X,{icon:"edit",size:"small",type:"primary",link:"",onClick:e=>W(l.row)},{default:r((()=>[o("查看")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),n("div",h,[i(ll,{"current-page":B.value,"page-size":F.value,"page-sizes":[10,30,50,100],total:E.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:J,onSizeChange:A},null,8,["current-page","page-size","total"])])]),i(el,{modelValue:M.value,"onUpdate:modelValue":e[1]||(e[1]=l=>M.value=l),"before-close":Q,title:L.value},{footer:r((()=>[n("div",D,[i(X,{size:"small",onClick:Q},{default:r((()=>[o("关闭")])),_:1})])])),default:r((()=>[n("div",f,[n("table",null,[n("tr",null,[b,n("td",null,s(U.value.instructionId),1),y,n("td",null,s(U.value.instructionType),1)]),n("tr",null,[x,n("td",null,s(d(p)(U.value.genInstructionTime)),1),T,n("td",null,s(d(p)(U.value.execInstructionTime)),1)]),n("tr",null,[I,n("td",null,s(U.value.instructionSource),1),_,n("td",null,s(U.value.execState),1)]),n("tr",null,[w,n("td",null,s(U.value.debrisId),1),C,n("td",null,s(U.value.debrisName),1)]),n("tr",null,[z,n("td",null,s(U.value.satelliteId),1),j,n("td",null,s(U.value.satelliteName),1)]),n("tr",null,[k,n("td",S,s(U.value.instructionContent),1)]),n("tr",null,[V,n("td",null,s(U.value.chainId),1),G,n("td",null,s(U.value.chainTime?d(p)(U.value.chainTime):""),1)]),n("tr",null,[q,n("td",N,s(U.value.blockHeight),1)])])])])),_:1},8,["modelValue","title"])])}}});export{U as default};
