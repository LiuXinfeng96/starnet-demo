/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{g as e,d as l,u as a,e as t,f as o}from"./menu.85677336.js";import u from"./icon.d8539f86.js";import{W as d}from"./warningBar.9e6751a3.js";import{c as i}from"./authorityBtn.3ffb00a3.js";import{_ as n,a as m,r as p,g as r,o as s,c as v,b as c,d as f,w as h,h as y,y as b,l as w,m as g,e as V,Q as _,i as k,p as I,q as x}from"./index.32af5906.js";const D=e=>(I("data-v-956d4c6f"),e=e(),x(),e),U={class:"gva-table-box"},C={class:"gva-btn-list"},z={key:0,class:"icon-column"},B={style:{display:"inline-flex"}},N=D((()=>c("span",{style:{"font-size":"12px","margin-right":"12px"}},"如果菜单包含子菜单，请创建router-view二级路由页面或者",-1))),q=D((()=>c("span",null," 高亮菜单 ",-1))),T=D((()=>c("span",null," 是否为基础页面 ",-1))),A={class:"dialog-footer"},j={name:"Menus"},M=n(Object.assign(j,{setup(n){const I=m({path:[{required:!0,message:"请输入菜单name",trigger:"blur"}],component:[{required:!0,message:"请输入文件路径",trigger:"blur"}],"meta.title":[{required:!0,message:"请输入菜单展示名称",trigger:"blur"}]}),x=p(1),D=p(0),j=p(999),M=p([]),S=p({}),F=async()=>{const l=await e({page:x.value,pageSize:j.value,...S.value});0===l.code&&(M.value=l.data.list,D.value=l.data.total,x.value=l.data.page,j.value=l.data.pageSize)};F();const P=()=>{Q.value.component=Q.value.component.replace(/\\/g,"/")},Q=p({ID:0,path:"",name:"",hidden:!1,parentId:"",component:"",meta:{activeName:"",title:"",icon:"",defaultMenu:!1,closeTab:!1,keepAlive:!1},parameters:[],menuBtn:[]}),$=()=>{Q.value.path=Q.value.name},E=e=>{O(),e()},H=p(null),K=p(!1),O=()=>{K.value=!1,H.value.resetFields(),Q.value={ID:0,path:"",name:"",hidden:!1,parentId:"",component:"",meta:{title:"",icon:"",defaultMenu:!1,closeTab:!1,keepAlive:!1}}},W=p(!1),G=()=>{O(),W.value=!1},J=async()=>{H.value.validate((async e=>{if(e){let e;e=Y.value?await a(Q.value):await t(Q.value),0===e.code&&(k({type:"success",message:Y.value?"编辑成功":"添加成功!"}),F()),O(),W.value=!1}}))},L=p([{ID:"0",title:"根菜单"}]),R=()=>{L.value=[{ID:"0",title:"根目录"}],X(M.value,L.value,!1)},X=(e,l,a)=>{e&&e.forEach((e=>{if(e.children&&e.children.length){const t={title:e.meta.title,ID:String(e.ID),disabled:a||e.ID===Q.value.ID,children:[]};X(e.children,t.children,a||e.ID===Q.value.ID),l.push(t)}else{const t={title:e.meta.title,ID:String(e.ID),disabled:a||e.ID===Q.value.ID};l.push(t)}}))},Y=p(!1),Z=p("新增菜单"),ee=e=>{Z.value="新增菜单",Q.value.parentId=String(e),Y.value=!1,R(),W.value=!0};return(e,a)=>{const t=r("el-button"),n=r("el-table-column"),m=r("el-icon"),p=r("el-table"),D=r("el-input"),j=r("el-form-item"),S=r("el-checkbox"),O=r("el-option"),X=r("el-select"),le=r("el-cascader"),ae=r("QuestionFilled"),te=r("el-tooltip"),oe=r("el-form"),ue=r("el-dialog");return s(),v("div",null,[c("div",U,[c("div",C,[f(t,{size:"small",type:"primary",icon:"plus",onClick:a[0]||(a[0]=e=>ee("0"))},{default:h((()=>[y("新增根菜单")])),_:1})]),f(p,{data:M.value,"row-key":"ID"},{default:h((()=>[f(n,{align:"left",label:"ID","min-width":"100",prop:"ID"}),f(n,{align:"left",label:"展示名称","min-width":"120",prop:"authorityName"},{default:h((e=>[c("span",null,b(e.row.meta.title),1)])),_:1}),f(n,{align:"left",label:"图标","min-width":"140",prop:"authorityName"},{default:h((e=>[e.row.meta.icon?(s(),v("div",z,[f(m,null,{default:h((()=>[(s(),w(g(e.row.meta.icon)))])),_:2},1024),c("span",null,b(e.row.meta.icon),1)])):V("",!0)])),_:1}),f(n,{align:"left",label:"路由Name","show-overflow-tooltip":"","min-width":"160",prop:"name"}),f(n,{align:"left",label:"路由Path","show-overflow-tooltip":"","min-width":"160",prop:"path"}),f(n,{align:"left",label:"是否隐藏","min-width":"100",prop:"hidden"},{default:h((e=>[c("span",null,b(e.row.hidden?"隐藏":"显示"),1)])),_:1}),f(n,{align:"left",label:"父节点","min-width":"90",prop:"parentId"}),f(n,{align:"left",label:"排序","min-width":"70",prop:"sort"}),f(n,{align:"left",label:"文件路径","min-width":"360",prop:"component"}),f(n,{align:"left",fixed:"right",label:"操作",width:"300"},{default:h((e=>[f(t,{size:"small",type:"primary",link:"",icon:"plus",onClick:l=>ee(e.row.ID)},{default:h((()=>[y("添加子菜单")])),_:2},1032,["onClick"]),f(t,{size:"small",type:"primary",link:"",icon:"edit",onClick:l=>(async e=>{Z.value="编辑菜单";const l=await o({id:e});Q.value=l.data.menu,Y.value=!0,R(),W.value=!0})(e.row.ID)},{default:h((()=>[y("编辑")])),_:2},1032,["onClick"]),f(t,{size:"small",type:"primary",link:"",icon:"delete",onClick:a=>{return t=e.row.ID,void _.confirm("此操作将永久删除所有角色下该菜单, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{0===(await l({ID:t})).code&&(k({type:"success",message:"删除成功!"}),1===M.value.length&&x.value>1&&x.value--,F())})).catch((()=>{k({type:"info",message:"已取消删除"})}));var t}},{default:h((()=>[y("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"])]),f(ue,{modelValue:W.value,"onUpdate:modelValue":a[16]||(a[16]=e=>W.value=e),"before-close":E,title:Z.value},{footer:h((()=>[c("div",A,[f(t,{size:"small",onClick:G},{default:h((()=>[y("取 消")])),_:1}),f(t,{size:"small",type:"primary",onClick:J},{default:h((()=>[y("确 定")])),_:1})])])),default:h((()=>[f(d,{title:"新增菜单，需要在角色管理内配置权限才可使用"}),W.value?(s(),w(oe,{key:0,ref_key:"menuForm",ref:H,inline:!0,model:Q.value,rules:I,"label-position":"top","label-width":"85px"},{default:h((()=>[f(j,{label:"路由Name",prop:"path",style:{width:"30%"}},{default:h((()=>[f(D,{modelValue:Q.value.name,"onUpdate:modelValue":a[1]||(a[1]=e=>Q.value.name=e),autocomplete:"off",placeholder:"唯一英文字符串",onChange:$},null,8,["modelValue"])])),_:1}),f(j,{prop:"path",style:{width:"30%"}},{label:h((()=>[c("div",B,[y(" 路由Path "),f(S,{modelValue:K.value,"onUpdate:modelValue":a[2]||(a[2]=e=>K.value=e),style:{float:"right","margin-left":"20px"}},{default:h((()=>[y("添加参数")])),_:1},8,["modelValue"])])])),default:h((()=>[f(D,{modelValue:Q.value.path,"onUpdate:modelValue":a[3]||(a[3]=e=>Q.value.path=e),disabled:!K.value,autocomplete:"off",placeholder:"建议只在后方拼接参数"},null,8,["modelValue","disabled"])])),_:1}),f(j,{label:"是否隐藏",style:{width:"30%"}},{default:h((()=>[f(X,{modelValue:Q.value.hidden,"onUpdate:modelValue":a[4]||(a[4]=e=>Q.value.hidden=e),placeholder:"是否在列表隐藏"},{default:h((()=>[f(O,{value:!1,label:"否"}),f(O,{value:!0,label:"是"})])),_:1},8,["modelValue"])])),_:1}),f(j,{label:"父节点ID",style:{width:"30%"}},{default:h((()=>[f(le,{modelValue:Q.value.parentId,"onUpdate:modelValue":a[5]||(a[5]=e=>Q.value.parentId=e),style:{width:"100%"},disabled:!Y.value,options:L.value,props:{checkStrictly:!0,label:"title",value:"ID",disabled:"disabled",emitPath:!1},"show-all-levels":!1,filterable:""},null,8,["modelValue","disabled","options"])])),_:1}),f(j,{label:"文件路径",prop:"component",style:{width:"60%"}},{default:h((()=>[f(D,{modelValue:Q.value.component,"onUpdate:modelValue":a[6]||(a[6]=e=>Q.value.component=e),autocomplete:"off",placeholder:"页面:view/xxx/xx.vue 插件:plugin/xx/xx.vue",onBlur:P},null,8,["modelValue"]),N,f(t,{style:{"margin-top":"4px"},size:"small",onClick:a[7]||(a[7]=e=>Q.value.component="view/routerHolder.vue")},{default:h((()=>[y("点我设置")])),_:1})])),_:1}),f(j,{label:"展示名称",prop:"meta.title",style:{width:"30%"}},{default:h((()=>[f(D,{modelValue:Q.value.meta.title,"onUpdate:modelValue":a[8]||(a[8]=e=>Q.value.meta.title=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),f(j,{label:"图标",prop:"meta.icon",style:{width:"30%"}},{default:h((()=>[f(u,{meta:Q.value.meta,style:{width:"100%"}},null,8,["meta"])])),_:1}),f(j,{label:"排序标记",prop:"sort",style:{width:"30%"}},{default:h((()=>[f(D,{modelValue:Q.value.sort,"onUpdate:modelValue":a[9]||(a[9]=e=>Q.value.sort=e),modelModifiers:{number:!0},autocomplete:"off"},null,8,["modelValue"])])),_:1}),f(j,{prop:"meta.activeName",style:{width:"30%"}},{label:h((()=>[c("div",null,[q,f(te,{content:"注：当到达此路由时候，指定左侧菜单指定name会处于活跃状态（亮起），可为空，为空则为本路由Name。",placement:"top",effect:"light"},{default:h((()=>[f(m,null,{default:h((()=>[f(ae)])),_:1})])),_:1})])])),default:h((()=>[f(D,{modelValue:Q.value.meta.activeName,"onUpdate:modelValue":a[10]||(a[10]=e=>Q.value.meta.activeName=e),placeholder:Q.value.name,autocomplete:"off"},null,8,["modelValue","placeholder"])])),_:1}),f(j,{label:"KeepAlive",prop:"meta.keepAlive",style:{width:"30%"}},{default:h((()=>[f(X,{modelValue:Q.value.meta.keepAlive,"onUpdate:modelValue":a[11]||(a[11]=e=>Q.value.meta.keepAlive=e),style:{width:"100%"},placeholder:"是否keepAlive缓存页面"},{default:h((()=>[f(O,{value:!1,label:"否"}),f(O,{value:!0,label:"是"})])),_:1},8,["modelValue"])])),_:1}),f(j,{label:"CloseTab",prop:"meta.closeTab",style:{width:"30%"}},{default:h((()=>[f(X,{modelValue:Q.value.meta.closeTab,"onUpdate:modelValue":a[12]||(a[12]=e=>Q.value.meta.closeTab=e),style:{width:"100%"},placeholder:"是否自动关闭tab"},{default:h((()=>[f(O,{value:!1,label:"否"}),f(O,{value:!0,label:"是"})])),_:1},8,["modelValue"])])),_:1}),f(j,{style:{width:"30%"}},{label:h((()=>[c("div",null,[T,f(te,{content:"此项选择为是，则不会展示左侧菜单以及顶部信息。",placement:"top",effect:"light"},{default:h((()=>[f(m,null,{default:h((()=>[f(ae)])),_:1})])),_:1})])])),default:h((()=>[f(X,{modelValue:Q.value.meta.defaultMenu,"onUpdate:modelValue":a[13]||(a[13]=e=>Q.value.meta.defaultMenu=e),style:{width:"100%"},placeholder:"是否为基础页面"},{default:h((()=>[f(O,{value:!1,label:"否"}),f(O,{value:!0,label:"是"})])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])):V("",!0),c("div",null,[f(t,{size:"small",type:"primary",icon:"edit",onClick:a[14]||(a[14]=e=>{return(l=Q.value).parameters||(l.parameters=[]),void l.parameters.push({type:"query",key:"",value:""});var l})},{default:h((()=>[y("新增菜单参数")])),_:1}),f(p,{data:Q.value.parameters,style:{width:"100%"}},{default:h((()=>[f(n,{align:"left",prop:"type",label:"参数类型",width:"180"},{default:h((e=>[f(X,{modelValue:e.row.type,"onUpdate:modelValue":l=>e.row.type=l,placeholder:"请选择"},{default:h((()=>[f(O,{key:"query",value:"query",label:"query"}),f(O,{key:"params",value:"params",label:"params"})])),_:2},1032,["modelValue","onUpdate:modelValue"])])),_:1}),f(n,{align:"left",prop:"key",label:"参数key",width:"180"},{default:h((e=>[c("div",null,[f(D,{modelValue:e.row.key,"onUpdate:modelValue":l=>e.row.key=l},null,8,["modelValue","onUpdate:modelValue"])])])),_:1}),f(n,{align:"left",prop:"value",label:"参数值"},{default:h((e=>[c("div",null,[f(D,{modelValue:e.row.value,"onUpdate:modelValue":l=>e.row.value=l},null,8,["modelValue","onUpdate:modelValue"])])])),_:1}),f(n,{align:"left"},{default:h((e=>[c("div",null,[f(t,{type:"danger",size:"small",icon:"delete",onClick:l=>{return a=Q.value.parameters,t=e.$index,void a.splice(t,1);var a,t}},{default:h((()=>[y("删除")])),_:2},1032,["onClick"])])])),_:1})])),_:1},8,["data"]),f(t,{style:{"margin-top":"12px"},size:"small",type:"primary",icon:"edit",onClick:a[15]||(a[15]=e=>{return(l=Q.value).menuBtn||(l.menuBtn=[]),void l.menuBtn.push({name:"",desc:""});var l})},{default:h((()=>[y("新增可控按钮")])),_:1}),f(p,{data:Q.value.menuBtn,style:{width:"100%"}},{default:h((()=>[f(n,{align:"left",prop:"name",label:"按钮名称",width:"180"},{default:h((e=>[c("div",null,[f(D,{modelValue:e.row.name,"onUpdate:modelValue":l=>e.row.name=l},null,8,["modelValue","onUpdate:modelValue"])])])),_:1}),f(n,{align:"left",prop:"name",label:"备注",width:"180"},{default:h((e=>[c("div",null,[f(D,{modelValue:e.row.desc,"onUpdate:modelValue":l=>e.row.desc=l},null,8,["modelValue","onUpdate:modelValue"])])])),_:1}),f(n,{align:"left"},{default:h((e=>[c("div",null,[f(t,{type:"danger",size:"small",icon:"delete",onClick:l=>(async(e,l)=>{const a=e[l];if(0===a.ID)return void e.splice(l,1);0!==(await i({id:a.ID})).code||e.splice(l,1)})(Q.value.menuBtn,e.$index)},{default:h((()=>[y("删除")])),_:2},1032,["onClick"])])])),_:1})])),_:1},8,["data"])])])),_:1},8,["modelValue","title"])])}}}),[["__scopeId","data-v-956d4c6f"]]);export{M as default};