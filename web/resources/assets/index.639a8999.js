/*! 
 Build based on gin-vue-admin 
 Time : 1672137460000 */
import{_ as e,r,j as a,R as s,g as c,o as i,c as t,F as p,G as l,l as n,e as u}from"./index.326c508a.js";const o=""+new URL("noBody.745c3d16.png",import.meta.url).href,d={class:"headerAvatar"},m=["src"],g=["src"],y=["src"],I={name:"CustomPic"},f=e(Object.assign(I,{props:{picType:{type:String,required:!1,default:"avatar"},picSrc:{type:String,required:!1,default:""}},setup(e){const I=e,f=r("/satellitebc/"),v=r(o),S=a(),h=s((()=>""===I.picSrc?""!==S.userInfo.headerImg&&"http"===S.userInfo.headerImg.slice(0,4)?S.userInfo.headerImg:f.value+S.userInfo.headerImg:""!==I.picSrc&&"http"===I.picSrc.slice(0,4)?I.picSrc:f.value+I.picSrc)),k=s((()=>I.picSrc&&"http"!==I.picSrc.slice(0,4)?f.value+I.picSrc:I.picSrc));return(r,a)=>{const s=c("el-avatar");return i(),t("span",d,["avatar"===e.picType?(i(),t(p,{key:0},[l(S).userInfo.headerImg?(i(),n(s,{key:0,size:30,src:l(h)},null,8,["src"])):(i(),n(s,{key:1,size:30,src:v.value},null,8,["src"]))],64)):u("",!0),"img"===e.picType?(i(),t(p,{key:1},[l(S).userInfo.headerImg?(i(),t("img",{key:0,src:l(h),class:"avatar"},null,8,m)):(i(),t("img",{key:1,src:v.value,class:"avatar"},null,8,g))],64)):u("",!0),"file"===e.picType?(i(),t("img",{key:2,src:l(k),class:"file"},null,8,y)):u("",!0)])}}}),[["__scopeId","data-v-9774a207"]]);export{f as C};
