/*! 
 Build based on gin-vue-admin 
 Time : 1671785957000 */
import{s as e,i as a,_ as l,r as t,j as o,g as s,o as n,c as i,b as d,d as p,w as c,h as r,G as m,y as u}from"./index.860b48e6.js";import{g as b}from"./menu.e3568139.js";const x=(e,l)=>{if(void 0!==e.data){if("application/json"===e.data.type){const l=new FileReader;l.onload=function(){const e=JSON.parse(l.result).msg;a({showClose:!0,message:e,type:"error"})},l.readAsText(new Blob([e.data]))}}else{var t=window.URL.createObjectURL(new Blob([e])),o=document.createElement("a");o.style.display="none",o.href=t,o.download=l;var s=new MouseEvent("click");o.dispatchEvent(s)}},f=()=>e({url:"/excel/loadExcel",method:"get"}),w={class:"upload"},h={class:"gva-table-box"},v={class:"gva-btn-list"},g=l(Object.assign({name:"Excel"},{setup(a){const l=t("/api"),g=t(1),y=t(0),E=t(999),_=t([]),j=async(e=(()=>{}))=>{const a=await e({page:g.value,pageSize:E.value});0===a.code&&(_.value=a.data.list,y.value=a.data.total,g.value=a.data.page,E.value=a.data.pageSize)};j(b);const k=o(),z=a=>{a&&"string"==typeof a||(a="ExcelExport.xlsx"),((a,l)=>{e({url:"/excel/exportExcel",method:"post",data:{fileName:l,infoList:a},responseType:"blob"}).then((e=>{x(e,l)}))})(_.value,a)},I=()=>{j(f)},T=()=>{var a;e({url:"/excel/downloadTemplate",method:"get",params:{fileName:a="ExcelTemplate.xlsx"},responseType:"blob"}).then((e=>{x(e,a)}))};return(e,a)=>{const t=s("el-button"),o=s("el-upload"),b=s("el-table-column"),x=s("el-table");return n(),i("div",w,[d("div",h,[d("div",v,[p(o,{class:"excel-btn",action:`${l.value}/excel/importExcel`,headers:{"x-token":m(k).token},"on-success":I,"show-file-list":!1},{default:c((()=>[p(t,{size:"small",type:"primary",icon:"upload"},{default:c((()=>[r("导入")])),_:1})])),_:1},8,["action","headers"]),p(t,{class:"excel-btn",size:"small",type:"primary",icon:"download",onClick:a[0]||(a[0]=e=>z("ExcelExport.xlsx"))},{default:c((()=>[r("导出")])),_:1}),p(t,{class:"excel-btn",size:"small",type:"success",icon:"download",onClick:a[1]||(a[1]=e=>T())},{default:c((()=>[r("下载模板")])),_:1})]),p(x,{data:_.value,"row-key":"ID"},{default:c((()=>[p(b,{align:"left",label:"ID","min-width":"100",prop:"ID"}),p(b,{align:"left","show-overflow-tooltip":"",label:"路由Name","min-width":"160",prop:"name"}),p(b,{align:"left","show-overflow-tooltip":"",label:"路由Path","min-width":"160",prop:"path"}),p(b,{align:"left",label:"是否隐藏","min-width":"100",prop:"hidden"},{default:c((e=>[d("span",null,u(e.row.hidden?"隐藏":"显示"),1)])),_:1}),p(b,{align:"left",label:"父节点","min-width":"90",prop:"parentId"}),p(b,{align:"left",label:"排序","min-width":"70",prop:"sort"}),p(b,{align:"left",label:"文件路径","min-width":"360",prop:"component"})])),_:1},8,["data"])])])}}}),[["__scopeId","data-v-3bd577b3"]]);export{g as default};
