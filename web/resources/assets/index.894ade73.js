/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{s as e,r as l,g as a,o as u,c as t,b as o,d as r,w as d,h as s,y as i,G as n,e as m,F as p,v,l as g,i as b}from"./index.32af5906.js";import{a as c}from"./format.fe46b3b1.js";import"./date.23f5a973.js";import"./dictionary.28290efa.js";import"./sysDictionary.b7a85079.js";const f={class:"gva-search-box"},h={class:"gva-table-box"},y={class:"gva-btn-list"},V={key:0},_={class:"gva-pagination"},w={class:"dialog-footer"},z={name:"fragmentInfo"},C=Object.assign(z,{setup(z){const C=l({debrisId:"",debrisName:"",angle:null,speed:null,height:null,volume:null,type:""}),U=l([{value:"Small",label:"Small"},{value:"Medium",label:"Medium"},{value:"Large",label:"Large"}]),x=l({debrisId:[{required:!0,message:"请输入碎片编号",trigger:"blur"},{validator:(e,l,a)=>{if(!/^[a-zA-Z0-9._-]{3,30}$/.test(l))return a(new Error("编号须3-30位，由字母、数字以及._-组成"));a()},trigger:"blur"}],debrisName:[{required:!0,message:"请输入碎片名称",trigger:"blur"}],angle:[{required:!0,message:"请输入运行角度",trigger:"blur"}],speed:[{required:!0,message:"请输入运行速度",trigger:"blur"}],height:[{required:!0,message:"请输入运行高度",trigger:"blur"}],volume:[{required:!0,message:"请输入碎片体积",trigger:"blur"}],type:[{required:!0,message:"请选择碎片类型",trigger:"blur"}]}),I=l(1),j=l(0),k=l(10),q=l([]),N=l({}),F=()=>{I.value=1,k.value=10,L()},S=e=>{k.value=e,L()},T=e=>{I.value=e,L()},L=async()=>{let l={page:I.value,pageSize:k.value,...N.value};const a=await(l=>e({url:"/satellitebc/exec/getdebrislist",method:"get",params:l}))(l);200===a.code&&(q.value=a.data,j.value=a.total)};L();const M=l(null),A=l("添加碎片信息"),D=l(!1),E=()=>{D.value=!0},G=()=>{M.value.resetFields(),C.value={debrisId:"",debrisName:"",angle:null,speed:null,height:null,volume:null,type:""},D.value=!1},O=async()=>{M.value.validate((async l=>{if(l){200===(await(a=C.value,e({url:"/satellitebc/exec/adddebris",method:"post",data:a}))).code&&b({type:"success",message:"添加成功",showClose:!0}),L(),G()}var a}))};return(e,l)=>{const b=a("el-input"),z=a("el-form-item"),L=a("el-button"),Z=a("el-form"),$=a("el-table-column"),B=a("el-table"),H=a("el-pagination"),J=a("el-input-number"),K=a("el-option"),P=a("el-select"),Q=a("el-dialog");return u(),t("div",null,[o("div",f,[r(Z,{ref:"searchForm",inline:!0,model:N.value},{default:d((()=>[r(z,{label:"碎片名称"},{default:d((()=>[r(b,{modelValue:N.value.searchConditions,"onUpdate:modelValue":l[0]||(l[0]=e=>N.value.searchConditions=e),placeholder:"碎片名称"},null,8,["modelValue"])])),_:1}),r(z,null,{default:d((()=>[r(L,{size:"small",type:"primary",icon:"search",onClick:F},{default:d((()=>[s("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),o("div",h,[o("div",y,[r(L,{size:"small",type:"primary",icon:"plus",onClick:E},{default:d((()=>[s("自主发现")])),_:1})]),r(B,{data:q.value},{default:d((()=>[r($,{align:"left",label:"碎片编号",prop:"debrisId"}),r($,{align:"left",label:"碎片名称",prop:"debrisName"}),r($,{align:"left",label:"运行角度",prop:"angle"}),r($,{align:"left",label:"运行速度",prop:"speed"}),r($,{align:"left",label:"运行高度",prop:"height"}),r($,{align:"left",label:"碎片体积",prop:"volume"}),r($,{align:"left",label:"碎片类型",prop:"type"}),r($,{label:"更新时间",width:"180"},{default:d((e=>[s(i(n(c)(e.row.lastTime,1)),1)])),_:1}),r($,{label:"上链时间",width:"180"},{default:d((e=>[e.row.chainTime?(u(),t("div",V,i(n(c)(e.row.chainTime)),1)):m("",!0)])),_:1})])),_:1},8,["data"]),o("div",_,[r(H,{"current-page":I.value,"page-size":k.value,"page-sizes":[10,30,50,100],total:j.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:T,onSizeChange:S},null,8,["current-page","page-size","total"])])]),r(Q,{modelValue:D.value,"onUpdate:modelValue":l[8]||(l[8]=e=>D.value=e),"before-close":G,title:A.value},{footer:d((()=>[o("div",w,[r(L,{size:"small",onClick:G},{default:d((()=>[s("取 消")])),_:1}),r(L,{size:"small",type:"primary",onClick:O},{default:d((()=>[s("确 定")])),_:1})])])),default:d((()=>[r(Z,{ref_key:"apiForm",ref:M,model:C.value,rules:x.value,"label-width":"80px"},{default:d((()=>[r(z,{label:"碎片编号",prop:"debrisId"},{default:d((()=>[r(b,{modelValue:C.value.debrisId,"onUpdate:modelValue":l[1]||(l[1]=e=>C.value.debrisId=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),r(z,{label:"碎片名称",prop:"debrisName"},{default:d((()=>[r(b,{modelValue:C.value.debrisName,"onUpdate:modelValue":l[2]||(l[2]=e=>C.value.debrisName=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),r(z,{label:"运行角度",prop:"angle"},{default:d((()=>[r(J,{modelValue:C.value.angle,"onUpdate:modelValue":l[3]||(l[3]=e=>C.value.angle=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),r(z,{label:"运行速度",prop:"speed"},{default:d((()=>[r(J,{modelValue:C.value.speed,"onUpdate:modelValue":l[4]||(l[4]=e=>C.value.speed=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),r(z,{label:"运行高度",prop:"height"},{default:d((()=>[r(J,{modelValue:C.value.height,"onUpdate:modelValue":l[5]||(l[5]=e=>C.value.height=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),r(z,{label:"碎片体积",prop:"volume"},{default:d((()=>[r(J,{modelValue:C.value.volume,"onUpdate:modelValue":l[6]||(l[6]=e=>C.value.volume=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),r(z,{label:"碎片类型",prop:"type"},{default:d((()=>[r(P,{modelValue:C.value.type,"onUpdate:modelValue":l[7]||(l[7]=e=>C.value.type=e),placeholder:"请选择",style:{width:"100%"}},{default:d((()=>[(u(!0),t(p,null,v(U.value,(e=>(u(),g(K,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"])])}}});export{C as default};