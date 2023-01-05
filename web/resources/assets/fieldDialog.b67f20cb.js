/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{a as e,t as l}from"./stringFun.2b3a18f6.js";import{g as a}from"./sysDictionary.b7a85079.js";import{W as u}from"./warningBar.9e6751a3.js";import{_ as d,r as o,g as t,o as i,c as m,d as r,w as n,F as p,v as s,l as f,p as v,q as b,b as c}from"./index.32af5906.js";const y=(e=>(v("data-v-c6bc08a3"),e=e(),b(),e))((()=>c("span",{style:{"font-size":"12px"}},"自动填充",-1))),V={name:"FieldDialog"},g=d(Object.assign(V,{props:{dialogMiddle:{type:Object,default:function(){return{}}}},setup(d,{expose:v}){const b=d,c=o({}),V=o([]),g=o([{label:"=",value:"="},{label:"<>",value:"<>"},{label:">",value:">"},{label:"<",value:"<"},{label:"LIKE",value:"LIKE"},{label:"BETWEEN",value:"BETWEEN"},{label:"NOT BETWEEN",value:"NOT BETWEEN"}]),T=o([{label:"字符串",value:"string"},{label:"整型",value:"int"},{label:"布尔值",value:"bool"},{label:"浮点型",value:"float64"},{label:"时间",value:"time.Time"},{label:"枚举",value:"enum"}]),_=o({fieldName:[{required:!0,message:"请输入field英文名",trigger:"blur"}],fieldDesc:[{required:!0,message:"请输入field中文名",trigger:"blur"}],fieldJson:[{required:!0,message:"请输入field格式化json",trigger:"blur"}],columnName:[{required:!0,message:"请输入数据库字段",trigger:"blur"}],fieldType:[{required:!0,message:"请选择field数据类型",trigger:"blur"}]});(async()=>{c.value=b.dialogMiddle;const e=await a({page:1,pageSize:999999});V.value=e.data.list})();const E=()=>{c.value.fieldJson=e(c.value.fieldName),c.value.columnName=l(c.value.fieldJson)},N=()=>{c.value.fieldSearchType="",c.value.dictType=""},h=o(null);return v({fieldDialogFrom:h}),(e,l)=>{const a=t("el-input"),d=t("el-button"),o=t("el-form-item"),v=t("el-option"),b=t("el-select"),U=t("el-switch"),F=t("el-form");return i(),m("div",null,[r(u,{title:"id , created_at , updated_at , deleted_at 会自动生成请勿重复创建。搜索时如果条件为LIKE只支持字符串"}),r(F,{ref_key:"fieldDialogFrom",ref:h,model:c.value,"label-width":"120px","label-position":"right",rules:_.value,class:"grid-form"},{default:n((()=>[r(o,{label:"Field名称",prop:"fieldName"},{default:n((()=>[r(a,{modelValue:c.value.fieldName,"onUpdate:modelValue":l[0]||(l[0]=e=>c.value.fieldName=e),autocomplete:"off",style:{width:"80%"}},null,8,["modelValue"]),r(d,{size:"small",style:{width:"18%","margin-left":"2%"},onClick:E},{default:n((()=>[y])),_:1})])),_:1}),r(o,{label:"Field中文名",prop:"fieldDesc"},{default:n((()=>[r(a,{modelValue:c.value.fieldDesc,"onUpdate:modelValue":l[1]||(l[1]=e=>c.value.fieldDesc=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),r(o,{label:"FieldJSON",prop:"fieldJson"},{default:n((()=>[r(a,{modelValue:c.value.fieldJson,"onUpdate:modelValue":l[2]||(l[2]=e=>c.value.fieldJson=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),r(o,{label:"数据库字段名",prop:"columnName"},{default:n((()=>[r(a,{modelValue:c.value.columnName,"onUpdate:modelValue":l[3]||(l[3]=e=>c.value.columnName=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),r(o,{label:"数据库字段描述",prop:"comment"},{default:n((()=>[r(a,{modelValue:c.value.comment,"onUpdate:modelValue":l[4]||(l[4]=e=>c.value.comment=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),r(o,{label:"Field数据类型",prop:"fieldType"},{default:n((()=>[r(b,{modelValue:c.value.fieldType,"onUpdate:modelValue":l[5]||(l[5]=e=>c.value.fieldType=e),style:{width:"100%"},placeholder:"请选择field数据类型",clearable:"",onChange:N},{default:n((()=>[(i(!0),m(p,null,s(T.value,(e=>(i(),f(v,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),r(o,{label:"enum"===c.value.fieldType?"枚举值":"类型长度",prop:"dataTypeLong"},{default:n((()=>[r(a,{modelValue:c.value.dataTypeLong,"onUpdate:modelValue":l[6]||(l[6]=e=>c.value.dataTypeLong=e),placeholder:"enum"===c.value.fieldType?"例:'北京','天津'":"数据库类型长度"},null,8,["modelValue","placeholder"])])),_:1},8,["label"]),r(o,{label:"Field查询条件",prop:"fieldSearchType"},{default:n((()=>[r(b,{modelValue:c.value.fieldSearchType,"onUpdate:modelValue":l[7]||(l[7]=e=>c.value.fieldSearchType=e),style:{width:"100%"},placeholder:"请选择Field查询条件",clearable:""},{default:n((()=>[(i(!0),m(p,null,s(g.value,(e=>(i(),f(v,{key:e.value,label:e.label,value:e.value,disabled:"string"!==c.value.fieldType&&"LIKE"===e.value||"int"!==c.value.fieldType&&"time.Time"!==c.value.fieldType&&"float64"!==c.value.fieldType&&("BETWEEN"===e.value||"NOT BETWEEN"===e.value)},null,8,["label","value","disabled"])))),128))])),_:1},8,["modelValue"])])),_:1}),r(o,{label:"关联字典",prop:"dictType"},{default:n((()=>[r(b,{modelValue:c.value.dictType,"onUpdate:modelValue":l[8]||(l[8]=e=>c.value.dictType=e),style:{width:"100%"},disabled:"int"!==c.value.fieldType,placeholder:"请选择字典",clearable:""},{default:n((()=>[(i(!0),m(p,null,s(V.value,(e=>(i(),f(v,{key:e.type,label:`${e.type}(${e.name})`,value:e.type},null,8,["label","value"])))),128))])),_:1},8,["modelValue","disabled"])])),_:1}),r(o,{label:"是否排序"},{default:n((()=>[r(U,{modelValue:c.value.sort,"onUpdate:modelValue":l[9]||(l[9]=e=>c.value.sort=e)},null,8,["modelValue"])])),_:1}),r(o,{label:"是否必填"},{default:n((()=>[r(U,{modelValue:c.value.require,"onUpdate:modelValue":l[10]||(l[10]=e=>c.value.require=e)},null,8,["modelValue"])])),_:1}),r(o,{label:"是否可清空"},{default:n((()=>[r(U,{modelValue:c.value.clearable,"onUpdate:modelValue":l[11]||(l[11]=e=>c.value.clearable=e)},null,8,["modelValue"])])),_:1}),r(o,{label:"校验失败文案"},{default:n((()=>[r(a,{modelValue:c.value.errorText,"onUpdate:modelValue":l[12]||(l[12]=e=>c.value.errorText=e)},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])}}}),[["__scopeId","data-v-c6bc08a3"]]);export{g as default};
