/*! 
 Build based on gin-vue-admin 
 Time : 1672224415000 */
import{f as a}from"./sysDictionary.b7a85079.js";import{M as t,r as i}from"./index.32af5906.js";const s=t("dictionary",(()=>{const t=i({}),s=a=>{t.value={...t.value,...a}};return{dictionaryMap:t,setDictionaryMap:s,getDictionary:async i=>{if(t.value[i]&&t.value[i].length)return t.value[i];{const e=await a({type:i});if(0===e.code){const a={},r=[];return e.data.resysDictionary.sysDictionaryDetails&&e.data.resysDictionary.sysDictionaryDetails.forEach((a=>{r.push({label:a.label,value:a.value})})),a[e.data.resysDictionary.type]=r,s(a),t.value[i]}}}}}));export{s as u};
