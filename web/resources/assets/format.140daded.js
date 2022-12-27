/*! 
 Build based on gin-vue-admin 
 Time : 1672137460000 */
import{f as e}from"./date.23f5a973.js";import"./dictionary.5509472a.js";const t=e=>null!==e?e?"是":"否":"",n=t=>{if(null!==t&&""!==t){var n=new Date(t);return e(n,"yyyy-MM-dd hh:mm:ss")}return""},r=(e,t)=>{if(0===e)return"";var n=new Date(1e3*e);return 1==t&&(n=new Date(e)),n.getFullYear()+"-"+(parseInt(n.getMonth())<9?"0"+(n.getMonth()+1):n.getMonth()+1)+"-"+(parseInt(n.getDate())<10?"0"+n.getDate():n.getDate())+" "+(parseInt(n.getHours())<10?"0"+n.getHours():n.getHours())+":"+(parseInt(n.getMinutes())<10?"0"+n.getMinutes():n.getMinutes())+":"+(parseInt(n.getSeconds())<10?"0"+n.getSeconds():n.getSeconds())};export{r as a,t as b,n as f};
