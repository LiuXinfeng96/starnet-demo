/*! 
 Build based on gin-vue-admin 
 Time : 1671785957000 */
import{s as e}from"./index.vue_vue_type_style_index_0_scoped_944c2140_lang.da5f26a1.js";import{_ as a,B as s,D as n,r as l,o as t,c}from"./index.860b48e6.js";const o={key:0,class:"gvaIcon gvaIcon-fullscreen-expand"},d={key:1,class:"gvaIcon gvaIcon-fullscreen-shrink"},i={name:"Screenfull"},r=a(Object.assign(i,{props:{width:{type:Number,default:22},height:{type:Number,default:22},fill:{type:String,default:"#48576a"}},setup(a){s((()=>{e.isEnabled&&e.on("change",u)})),n((()=>{e.off("change")}));const i=()=>{e.isEnabled&&e.toggle()},r=l(!0),u=()=>{r.value=!e.isFullscreen};return(e,a)=>(t(),c("div",{onClick:i},[r.value?(t(),c("div",o)):(t(),c("div",d))]))}}),[["__scopeId","data-v-944c2140"]]);export{r as default};
