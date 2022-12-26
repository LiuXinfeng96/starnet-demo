/*! 
 Build based on gin-vue-admin 
 Time : 1671785957000 */
import{s as e,r as l,g as a,o as t,c as u,b as o,d as s,w as r,h as d,y as i,G as n,e as m,F as p,v,l as c,i as g}from"./index.860b48e6.js";import{a as b}from"./format.4fc21eae.js";import"./date.23f5a973.js";import"./dictionary.8dcfcd04.js";import"./sysDictionary.12ad4d06.js";const f={class:"gva-search-box"},y={class:"gva-table-box"},V={class:"gva-btn-list"},_={key:0},h={class:"gva-pagination"},w={class:"dialog-footer"},I={name:"satelliteStatus"},z=Object.assign(I,{setup(I){const z=(e,l,a)=>{if(!/^[a-zA-Z0-9._-]{3,30}$/.test(l))return a(new Error("编号须3-30位，由字母、数字以及._-组成"));a()},k=l({satelliteId:"",satelliteName:"",orbitId:"",meanAnomaly:null,speed:null,runState:""}),C=l([{value:"正常",label:"正常"},{value:"异常",label:"异常"}]);l("");const S=l({satelliteId:[{required:!0,message:"请输入卫星编码",trigger:"blur"},{validator:z,trigger:"blur"}],satelliteName:[{required:!0,message:"请输入卫星名称",trigger:"blur"}],orbitId:[{required:!0,message:"请输入所在轨道编码",trigger:"blur"},{validator:z,trigger:"blur"}],meanAnomaly:[{required:!0,message:"请输入平近点角",trigger:"blur"}],speed:[{required:!0,message:"请输入速度（km/s）",trigger:"blur"}],runState:[{required:!0,message:"请选择运行状态",trigger:"blur"}]}),x=l(1),A=l(0),U=l(10),j=l([]),N=l({}),q=()=>{x.value=1,U.value=10,D()},F=e=>{U.value=e,D()},T=e=>{x.value=e,D()},D=async()=>{let l={page:x.value,pageSize:U.value,...N.value};const a=await(l=>e({url:"/satellitebc/exec/getsatellitestatelist",method:"get",params:l}))(l);200===a.code&&(j.value=a.data,A.value=a.total)};D();const E=l(null),G=l("添加卫星轨道状态"),O=l(!1),Z=()=>{O.value=!0},$=()=>{E.value.resetFields(),k.value={satelliteId:"",satelliteName:"",orbitId:"",meanAnomaly:null,speed:null,runState:""},O.value=!1},B=async()=>{E.value.validate((async l=>{if(l){200===(await(a=k.value,e({url:"/satellitebc/exec/addsatellitestate",method:"post",data:a}))).code&&g({type:"success",message:"添加成功",showClose:!0}),D(),$()}var a}))};return(e,l)=>{const g=a("el-input"),I=a("el-form-item"),z=a("el-button"),D=a("el-form"),H=a("el-table-column"),J=a("el-table"),K=a("el-pagination"),L=a("el-input-number"),M=a("el-option"),P=a("el-select"),Q=a("el-dialog");return t(),u("div",null,[o("div",f,[s(D,{ref:"searchForm",inline:!0,model:N.value},{default:r((()=>[s(I,{label:"卫星名称"},{default:r((()=>[s(g,{modelValue:N.value.searchConditions,"onUpdate:modelValue":l[0]||(l[0]=e=>N.value.searchConditions=e),placeholder:"卫星名称"},null,8,["modelValue"])])),_:1}),s(I,null,{default:r((()=>[s(z,{size:"small",type:"primary",icon:"search",onClick:q},{default:r((()=>[d("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),o("div",y,[o("div",V,[s(z,{size:"small",type:"primary",icon:"plus",onClick:Z},{default:r((()=>[d("添加")])),_:1})]),s(J,{data:j.value},{default:r((()=>[s(H,{align:"left",label:"卫星名称",prop:"satelliteName"}),s(H,{align:"left",label:"卫星编码",prop:"satelliteId"}),s(H,{align:"left",label:"所在轨道编码",prop:"orbitId"}),s(H,{align:"left",label:"平近点角",prop:"meanAnomaly"}),s(H,{align:"left",label:"速度（km/s）",prop:"speed"}),s(H,{align:"left",label:"运行状态",prop:"runState"}),s(H,{label:"更新时间",width:"180"},{default:r((e=>[d(i(n(b)(e.row.lastTime,1)),1)])),_:1}),s(H,{label:"上链时间",width:"180"},{default:r((e=>[e.row.chainTime?(t(),u("div",_,i(n(b)(e.row.chainTime)),1)):m("",!0)])),_:1})])),_:1},8,["data"]),o("div",h,[s(K,{"current-page":x.value,"page-size":U.value,"page-sizes":[10,30,50,100],total:A.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:T,onSizeChange:F},null,8,["current-page","page-size","total"])])]),s(Q,{modelValue:O.value,"onUpdate:modelValue":l[7]||(l[7]=e=>O.value=e),"before-close":$,title:G.value},{footer:r((()=>[o("div",w,[s(z,{size:"small",onClick:$},{default:r((()=>[d("取 消")])),_:1}),s(z,{size:"small",type:"primary",onClick:B},{default:r((()=>[d("确 定")])),_:1})])])),default:r((()=>[s(D,{ref_key:"apiForm",ref:E,model:k.value,rules:S.value,"label-width":"120px"},{default:r((()=>[s(I,{label:"卫星编码",prop:"satelliteId"},{default:r((()=>[s(g,{modelValue:k.value.satelliteId,"onUpdate:modelValue":l[1]||(l[1]=e=>k.value.satelliteId=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),s(I,{label:"卫星名称",prop:"satelliteName"},{default:r((()=>[s(g,{modelValue:k.value.satelliteName,"onUpdate:modelValue":l[2]||(l[2]=e=>k.value.satelliteName=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),s(I,{label:"所在轨道编码",prop:"orbitId"},{default:r((()=>[s(g,{modelValue:k.value.orbitId,"onUpdate:modelValue":l[3]||(l[3]=e=>k.value.orbitId=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),s(I,{label:"平近点角",prop:"meanAnomaly"},{default:r((()=>[s(L,{modelValue:k.value.meanAnomaly,"onUpdate:modelValue":l[4]||(l[4]=e=>k.value.meanAnomaly=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),s(I,{label:"速度（km/s）",prop:"speed"},{default:r((()=>[s(L,{modelValue:k.value.speed,"onUpdate:modelValue":l[5]||(l[5]=e=>k.value.speed=e),controls:!1,style:{width:"100%"}},null,8,["modelValue"])])),_:1}),s(I,{label:"运行状态",prop:"runState"},{default:r((()=>[s(P,{modelValue:k.value.runState,"onUpdate:modelValue":l[6]||(l[6]=e=>k.value.runState=e),placeholder:"请选择",style:{width:"100%"}},{default:r((()=>[(t(!0),u(p,null,v(C.value,(e=>(t(),c(M,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue","title"])])}}});export{z as default};
