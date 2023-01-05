/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{s as e,_ as t,r as a,j as s,g as o,o as l,c as i,d as n,w as r,h as d,G as u,i as c}from"./index.32af5906.js";const h=t=>e({url:"/fileUploadAndDownload/getFileList",method:"post",data:t}),p=t=>e({url:"/fileUploadAndDownload/deleteFile",method:"post",data:t}),m=t=>e({url:"/fileUploadAndDownload/editFileName",method:"post",data:t});class f{constructor(e,t,a=1920){this.file=e,this.fileSize=t,this.maxWH=a}compress(){const e=this.file.type,t=this.file.size/1024;return new Promise((a=>{const s=new FileReader;s.readAsDataURL(this.file),s.onload=()=>{const o=document.createElement("canvas"),l=document.createElement("img");l.src=s.result,l.onload=()=>{const s=o.getContext("2d"),i=this.dWH(l.width,l.height,this.maxWH);o.width=i.width,o.height=i.height,s.clearRect(0,0,o.width,o.height),s.drawImage(l,0,0,o.width,o.height);const n=o.toDataURL(e,.9),r=this.fileSizeKB(n);r>this.fileSize&&console.log("图片尺寸太大!"+t+" >> "+r);const d=this.dataURLtoBlob(n,e),u=new File([d],this.file.name);a(u)}}}))}dWH(e,t,a){const s={width:e,height:t};return Math.max(e,t)>a?e>t?(s.width=a,s.height=Math.round(t*(a/e)),s):(s.height=a,s.width=Math.round(e*(a/t)),s):s}fileSizeKB(e){let t=0;return t=Math.round(3*e.split(",")[1].length/4/1024),t}dataURLtoBlob(e,t){const a=atob(e.split(",")[1]);let s=e.split(",")[0].split(":")[1].split(";")[0];const o=new ArrayBuffer(a.length),l=new Uint8Array(o);for(let i=0;i<a.length;i++)l[i]=a.charCodeAt(i);return t&&(s=t),new Blob([o],{type:s,lastModifiedDate:new Date})}}const g={name:"UploadImage",methods:{}},w=t(Object.assign(g,{props:{imageUrl:{type:String,default:""},fileSize:{type:Number,default:2048},maxWH:{type:Number,default:1920}},emits:["on-success"],setup(e,{emit:t}){const h=e,p=a("/satellitebc"),m=s(),g=e=>{const t="image/jpeg"===e.type,a="image/png"===e.type;if(!t&&!a)return c.error("上传头像图片只能是 jpg或png 格式!"),!1;const s=e.size/1024<h.fileSize;if(!s){return new f(e,h.fileSize,h.maxWH).compress()}return s},w=e=>{const{data:a}=e;a.file&&t("on-success",a.file.url)};return(e,t)=>{const a=o("el-button"),s=o("el-upload");return l(),i("div",null,[n(s,{action:`${p.value}/fileUploadAndDownload/upload`,headers:{"x-token":u(m).token},"show-file-list":!1,"on-success":w,"before-upload":g,multiple:!1},{default:r((()=>[n(a,{size:"small",type:"primary"},{default:r((()=>[d("压缩上传")])),_:1})])),_:1},8,["action","headers"])])}}}),[["__scopeId","data-v-b6c7875e"]]),b={name:"UploadCommon",methods:{}},y=Object.assign(b,{emits:["on-success"],setup(e,{emit:t}){const h=a("/satellitebc"),p=s(),m=a(!1),f=e=>{m.value=!0;const t="image/jpeg"===e.type,a="image/png"===e.type,s=e.size/1024/1024<.5;return t||a||(c.error("上传图片只能是 jpg或png 格式!"),m.value=!1),s||(c.error("未压缩未见上传图片大小不能超过 500KB，请使用压缩上传"),m.value=!1),(a||t)&&s},g=e=>{const{data:a}=e;a.file&&t("on-success",a.file.url)},w=()=>{c({type:"error",message:"上传失败"}),m.value=!1};return(e,t)=>{const a=o("el-button"),s=o("el-upload");return l(),i("div",null,[n(s,{action:`${h.value}/fileUploadAndDownload/upload`,"before-upload":f,headers:{"x-token":u(p).token},"on-error":w,"on-success":g,"show-file-list":!1,class:"upload-btn"},{default:r((()=>[n(a,{size:"small",type:"primary"},{default:r((()=>[d("普通上传")])),_:1})])),_:1},8,["action","headers"])])}}});export{w as U,y as _,p as d,m as e,h as g};
