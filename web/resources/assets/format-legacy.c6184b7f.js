/*! 
 Build based on gin-vue-admin 
 Time : 1672137460000 */
System.register(["./date-legacy.431857fb.js","./dictionary-legacy.5a7b76de.js"],(function(e,t){"use strict";var n;return{setters:[function(e){n=e.f},function(){}],execute:function(){e("b",(function(e){return null!==e?e?"是":"否":""})),e("f",(function(e){if(null!==e&&""!==e){var t=new Date(e);return n(t,"yyyy-MM-dd hh:mm:ss")}return""})),e("a",(function(e,t){if(0===e)return"";var n=new Date(1e3*e);return 1==t&&(n=new Date(e)),n.getFullYear()+"-"+(parseInt(n.getMonth())<9?"0"+(n.getMonth()+1):n.getMonth()+1)+"-"+(parseInt(n.getDate())<10?"0"+n.getDate():n.getDate())+" "+(parseInt(n.getHours())<10?"0"+n.getHours():n.getHours())+":"+(parseInt(n.getMinutes())<10?"0"+n.getMinutes():n.getMinutes())+":"+(parseInt(n.getSeconds())<10?"0"+n.getSeconds():n.getSeconds())}))}}}));
