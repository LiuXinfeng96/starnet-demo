/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{s as e,r as l,g as a,o as t,c as o,b as u,d as s,w as i,h as n,y as r,G as d,e as m,F as p,v as c,l as v,i as b}from"./index.32af5906.js";import{a as g}from"./format.fe46b3b1.js";import"./date.23f5a973.js";import"./dictionary.28290efa.js";import"./sysDictionary.b7a85079.js";const f={class:"gva-search-box"},N={class:"gva-table-box"},V={class:"gva-btn-list"},h={key:0},w={class:"gva-pagination"},y={class:"dialog-footer"},_={name:"starsignStatus"},U=Object.assign(_,{setup(_){const U=l({constellationId:"",constellationName:"",satelliteTotalNum:null,satelliteUpNum:null,satelliteDownNum:null,satelliteLinkState:""}),k=l([{value:"正常",label:"正常"},{value:"异常",label:"异常"}]),z=l({constellationId:[{required:!0,message:"请输入星座编码",trigger:"blur"},{validator:(e,l,a)=>{if(!/^[a-zA-Z0-9._-]{3,30}$/.test(l))return a(new Error("编号须3-30位，由字母、数字以及._-组成"));a()},trigger:"blur"}],constellationName:[{required:!0,message:"请输入星座名称",trigger:"blur"}],satelliteTotalNum:[{required:!0,message:"请输入包含卫星数",trigger:"blur"}],satelliteUpNum:[{required:!0,message:"请输入正常卫星数量",trigger:"blur"}],satelliteDownNum:[{required:!0,message:"请输入异常卫星数量",trigger:"blur"}],satelliteLinkState:[{required:!0,message:"请选择星间链路状态",trigger:"blur"}]}),C=l(1),S=l(0),T=l(10),D=l([]),j=l({}),I=()=>{C.value=1,T.value=10,x()},L=e=>{T.value=e,x()},q=e=>{C.value=e,x()},x=async()=>{let l={page:C.value,pageSize:T.value,...j.value};const a=await(l=>e({url:"/satellitebc/control/getconstellationlist",method:"get",params:l}))(l);200===a.code&&(D.value=a.data,S.value=a.total)};x();const F=l(null),A=l("添加星座状态"),E=l(!1),G=()=>{E.value=!0},O=()=>{F.value.resetFields(),U.value={constellationId:"",constellationName:"",satelliteTotalNum:null,satelliteUpNum:null,satelliteDownNum:null,satelliteLinkState:""},E.value=!1},Z=async()=>{F.value.validate((async l=>{if(l){200===(await(a=U.value,e({url:"/satellitebc/control/addconstellation",method:"post",data:a}))).code&&b({type:"success",message:"添加成功",showClose:!0}),x(),O()}var a}))};return(e,l)=>{const b=a("el-input"),_=a("el-form-item"),x=a("el-button"),$=a("el-form"),B=a("el-table-column"),H=a("el-table"),J=a("el-pagination"),K=a("el-input-number"),M=a("el-option"),P=a("el-select"),Q=a("el-dialog");return t(),o("div",null,[u("div",f,[s($,{ref:"searchForm",inline:!0,model:j.value},{default:i((()=>[s(_,{label:"星座名称"},{default:i((()=>[s(b,{modelValue:j.value.searchConditions,"onUpdate:modelValue":l[0]||(l[0]=e=>j.value.searchConditions=e),placeholder:"星座名称"},null,8,["modelValue"])])),_:1}),s(_,null,{default:i((()=>[s(x,{size:"small",type:"primary",icon:"search",onClick:I},{default:i((()=>[n("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),u("div",N,[u("div",V,[s(x,{size:"small",type:"primary",icon:"plus",onClick:G},{default:i((()=>[n("添加")])),_:1})]),s(H,{data:D.value},{default:i((()=>[s(B,{label:"星座编码",prop:"constellationId"}),s(B,{label:"星座名称",prop:"constellationName"}),s(B,{label:"包含卫星数",prop:"satelliteTotalNum"}),s(B,{label:"正常卫星数量",prop:"satelliteUpNum"}),s(B,{label:"异常卫星数量",prop:"satelliteDownNum"}),s(B,{label:"星间链路状态",prop:"satelliteLinkState"}),s(B,{label:"更新时间",width:"180"},{default:i((e=>[n(r(d(g)(e.row.lastTime,1)),1)])),_:1}),s(B,{label:"上链时间",width:"180"},{default:i((e=>[e.row.chainTime?(t(),o("div",h,r(d(g)(e.row.chainTime)),1)):m("",!0)])),_:1})])),_:1},8,["data"]),u("div",w,[s(J,{"current-page":C.value,"page-size":T.value,"page-sizes":[10,30,50,100],total:S.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:q,onSizeChange:L},null,8,["current-page","page-size","total"])])]),s(Q,{modelValue:E.value,"onUpdate:modelValue":l[7]||(l[7]=e=>E.value=e),"before-close":O,title:A.value},{footer:i((()=>[u("div",y,[s(x,{size:"small",onClick:O},{default:i((()=>[n("取 消")])),_:1}),s(x,{size:"small",type:"primary",onClick:Z},{default:i((()=>[n("确 定")])),_:1})])])),default:i((()=>[s($,{ref_key:"apiForm",ref:F,model:U.value,rules:z.value,"label-width":"120px"},{default:i((()=>[s(_,{label:"星座编码",prop:"constellationId"},{default:i((()=>[s(b,{modelValue:U.value.constellationId,"onUpdate:modelValue":l[1]||(l[1]=e=>U.value.constellationId=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),s(_,{label:"星座名称",prop:"constellationName"},{default:i((()=>[s(b,{modelValue:U.value.constellationName,"onUpdate:modelValue":l[2]||(l[2]=e=>U.value.constellationName=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),s(_,{label:"包含卫星数",prop:"satelliteTotalNum"},{default:i((()=>[s(K,{modelValue:U.value.satelliteTotalNum,"onUpdate:modelValue":l[3]||(l[3]=e=>U.value.satelliteTotalNum=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),s(_,{label:"正常卫星数量",prop:"satelliteUpNum"},{default:i((()=>[s(K,{modelValue:U.value.satelliteUpNum,"onUpdate:modelValue":l[4]||(l[4]=e=>U.value.satelliteUpNum=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),s(_,{label:"异常卫星数量",prop:"satelliteDownNum"},{default:i((()=>[s(K,{modelValue:U.value.satelliteDownNum,"onUpdate:modelValue":l[5]||(l[5]=e=>U.value.satelliteDownNum=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),s(_,{label:"星间链路状态",prop:"satelliteLinkState"},{default:i((()=>[s(P,{modelValue:U.value.satelliteLinkState,"onUpdate:modelValue":l[6]||(l[6]=e=>U.value.satelliteLinkState=e),placeholder:"请选择",style:{width:"100%"}},{default:i((()=>[(t(!0),o(p,null,c(k.value,(e=>(t(),v(M,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"])])}}});export{U as default};
