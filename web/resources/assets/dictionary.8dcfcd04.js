/*! 
 Build based on gin-vue-admin 
 Time : 1671785957000 */
import{f as a}from"./sysDictionary.12ad4d06.js";import{M as t,r as i}from"./index.860b48e6.js";const e=t("dictionary",(()=>{const t=i({}),e=a=>{t.value={...t.value,...a}};return{dictionaryMap:t,setDictionaryMap:e,getDictionary:async i=>{if(t.value[i]&&t.value[i].length)return t.value[i];{const s=await a({type:i});if(0===s.code){const a={},r=[];return s.data.resysDictionary.sysDictionaryDetails&&s.data.resysDictionary.sysDictionaryDetails.forEach((a=>{r.push({label:a.label,value:a.value})})),a[s.data.resysDictionary.type]=r,e(a),t.value[i]}}}}}));export{e as u};
