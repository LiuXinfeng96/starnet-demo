/*! 
 Build based on gin-vue-admin 
 Time : 1671785957000 */
!function(){function e(){"use strict";/*! regenerator-runtime -- Copyright (c) 2014-present, Facebook, Inc. -- license (MIT): https://github.com/facebook/regenerator/blob/main/LICENSE */e=function(){return t};var t={},n=Object.prototype,r=n.hasOwnProperty,a=Object.defineProperty||function(e,t,n){e[t]=n.value},l="function"==typeof Symbol?Symbol:{},o=l.iterator||"@@iterator",u=l.asyncIterator||"@@asyncIterator",i=l.toStringTag||"@@toStringTag";function c(e,t,n){return Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}),e[t]}try{c({},"")}catch(L){c=function(e,t,n){return e[t]=n}}function f(e,t,n,r){var l=t&&t.prototype instanceof p?t:p,o=Object.create(l.prototype),u=new O(r||[]);return a(o,"_invoke",{value:_(e,n,u)}),o}function d(e,t,n){try{return{type:"normal",arg:e.call(t,n)}}catch(L){return{type:"throw",arg:L}}}t.wrap=f;var s={};function p(){}function m(){}function v(){}var h={};c(h,o,(function(){return this}));var y=Object.getPrototypeOf,g=y&&y(y(j([])));g&&g!==n&&r.call(g,o)&&(h=g);var w=v.prototype=p.prototype=Object.create(h);function b(e){["next","throw","return"].forEach((function(t){c(e,t,(function(e){return this._invoke(t,e)}))}))}function x(e,t){function n(a,l,o,u){var i=d(e[a],e,l);if("throw"!==i.type){var c=i.arg,f=c.value;return f&&"object"==typeof f&&r.call(f,"__await")?t.resolve(f.__await).then((function(e){n("next",e,o,u)}),(function(e){n("throw",e,o,u)})):t.resolve(f).then((function(e){c.value=e,o(c)}),(function(e){return n("throw",e,o,u)}))}u(i.arg)}var l;a(this,"_invoke",{value:function(e,r){function a(){return new t((function(t,a){n(e,r,t,a)}))}return l=l?l.then(a,a):a()}})}function _(e,t,n){var r="suspendedStart";return function(a,l){if("executing"===r)throw new Error("Generator is already running");if("completed"===r){if("throw"===a)throw l;return D()}for(n.method=a,n.arg=l;;){var o=n.delegate;if(o){var u=V(o,n);if(u){if(u===s)continue;return u}}if("next"===n.method)n.sent=n._sent=n.arg;else if("throw"===n.method){if("suspendedStart"===r)throw r="completed",n.arg;n.dispatchException(n.arg)}else"return"===n.method&&n.abrupt("return",n.arg);r="executing";var i=d(e,t,n);if("normal"===i.type){if(r=n.done?"completed":"suspendedYield",i.arg===s)continue;return{value:i.arg,done:n.done}}"throw"===i.type&&(r="completed",n.method="throw",n.arg=i.arg)}}}function V(e,t){var n=e.iterator[t.method];if(void 0===n){if(t.delegate=null,"throw"===t.method){if(e.iterator.return&&(t.method="return",t.arg=void 0,V(e,t),"throw"===t.method))return s;t.method="throw",t.arg=new TypeError("The iterator does not provide a 'throw' method")}return s}var r=d(n,e.iterator,t.arg);if("throw"===r.type)return t.method="throw",t.arg=r.arg,t.delegate=null,s;var a=r.arg;return a?a.done?(t[e.resultName]=a.value,t.next=e.nextLoc,"return"!==t.method&&(t.method="next",t.arg=void 0),t.delegate=null,s):a:(t.method="throw",t.arg=new TypeError("iterator result is not an object"),t.delegate=null,s)}function k(e){var t={tryLoc:e[0]};1 in e&&(t.catchLoc=e[1]),2 in e&&(t.finallyLoc=e[2],t.afterLoc=e[3]),this.tryEntries.push(t)}function I(e){var t=e.completion||{};t.type="normal",delete t.arg,e.completion=t}function O(e){this.tryEntries=[{tryLoc:"root"}],e.forEach(k,this),this.reset(!0)}function j(e){if(e){var t=e[o];if(t)return t.call(e);if("function"==typeof e.next)return e;if(!isNaN(e.length)){var n=-1,a=function t(){for(;++n<e.length;)if(r.call(e,n))return t.value=e[n],t.done=!1,t;return t.value=void 0,t.done=!0,t};return a.next=a}}return{next:D}}function D(){return{value:void 0,done:!0}}return m.prototype=v,a(w,"constructor",{value:v,configurable:!0}),a(v,"constructor",{value:m,configurable:!0}),m.displayName=c(v,i,"GeneratorFunction"),t.isGeneratorFunction=function(e){var t="function"==typeof e&&e.constructor;return!!t&&(t===m||"GeneratorFunction"===(t.displayName||t.name))},t.mark=function(e){return Object.setPrototypeOf?Object.setPrototypeOf(e,v):(e.__proto__=v,c(e,i,"GeneratorFunction")),e.prototype=Object.create(w),e},t.awrap=function(e){return{__await:e}},b(x.prototype),c(x.prototype,u,(function(){return this})),t.AsyncIterator=x,t.async=function(e,n,r,a,l){void 0===l&&(l=Promise);var o=new x(f(e,n,r,a),l);return t.isGeneratorFunction(n)?o:o.next().then((function(e){return e.done?e.value:o.next()}))},b(w),c(w,i,"Generator"),c(w,o,(function(){return this})),c(w,"toString",(function(){return"[object Generator]"})),t.keys=function(e){var t=Object(e),n=[];for(var r in t)n.push(r);return n.reverse(),function e(){for(;n.length;){var r=n.pop();if(r in t)return e.value=r,e.done=!1,e}return e.done=!0,e}},t.values=j,O.prototype={constructor:O,reset:function(e){if(this.prev=0,this.next=0,this.sent=this._sent=void 0,this.done=!1,this.delegate=null,this.method="next",this.arg=void 0,this.tryEntries.forEach(I),!e)for(var t in this)"t"===t.charAt(0)&&r.call(this,t)&&!isNaN(+t.slice(1))&&(this[t]=void 0)},stop:function(){this.done=!0;var e=this.tryEntries[0].completion;if("throw"===e.type)throw e.arg;return this.rval},dispatchException:function(e){if(this.done)throw e;var t=this;function n(n,r){return o.type="throw",o.arg=e,t.next=n,r&&(t.method="next",t.arg=void 0),!!r}for(var a=this.tryEntries.length-1;a>=0;--a){var l=this.tryEntries[a],o=l.completion;if("root"===l.tryLoc)return n("end");if(l.tryLoc<=this.prev){var u=r.call(l,"catchLoc"),i=r.call(l,"finallyLoc");if(u&&i){if(this.prev<l.catchLoc)return n(l.catchLoc,!0);if(this.prev<l.finallyLoc)return n(l.finallyLoc)}else if(u){if(this.prev<l.catchLoc)return n(l.catchLoc,!0)}else{if(!i)throw new Error("try statement without catch or finally");if(this.prev<l.finallyLoc)return n(l.finallyLoc)}}}},abrupt:function(e,t){for(var n=this.tryEntries.length-1;n>=0;--n){var a=this.tryEntries[n];if(a.tryLoc<=this.prev&&r.call(a,"finallyLoc")&&this.prev<a.finallyLoc){var l=a;break}}l&&("break"===e||"continue"===e)&&l.tryLoc<=t&&t<=l.finallyLoc&&(l=null);var o=l?l.completion:{};return o.type=e,o.arg=t,l?(this.method="next",this.next=l.finallyLoc,s):this.complete(o)},complete:function(e,t){if("throw"===e.type)throw e.arg;return"break"===e.type||"continue"===e.type?this.next=e.arg:"return"===e.type?(this.rval=this.arg=e.arg,this.method="return",this.next="end"):"normal"===e.type&&t&&(this.next=t),s},finish:function(e){for(var t=this.tryEntries.length-1;t>=0;--t){var n=this.tryEntries[t];if(n.finallyLoc===e)return this.complete(n.completion,n.afterLoc),I(n),s}},catch:function(e){for(var t=this.tryEntries.length-1;t>=0;--t){var n=this.tryEntries[t];if(n.tryLoc===e){var r=n.completion;if("throw"===r.type){var a=r.arg;I(n)}return a}}throw new Error("illegal catch attempt")},delegateYield:function(e,t,n){return this.delegate={iterator:j(e),resultName:t,nextLoc:n},"next"===this.method&&(this.arg=void 0),s}},t}function t(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function n(e){for(var n=1;n<arguments.length;n++){var a=null!=arguments[n]?arguments[n]:{};n%2?t(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):t(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t,n,r,a,l,o){try{var u=e[l](o),i=u.value}catch(c){return void n(c)}u.done?t(i):Promise.resolve(i).then(r,a)}function l(e){return function(){var t=this,n=arguments;return new Promise((function(r,l){var o=e.apply(t,n);function u(e){a(o,r,l,u,i,"next",e)}function i(e){a(o,r,l,u,i,"throw",e)}u(void 0)}))}}System.register(["./menu-legacy.08397938.js","./icon-legacy.15111151.js","./warningBar-legacy.81832a14.js","./authorityBtn-legacy.702ca678.js","./index-legacy.94eed338.js"],(function(t,r){"use strict";var a,o,u,i,c,f,d,s,p,m,v,h,y,g,w,b,x,_,V,k,I,O,j,D,L,E,U=document.createElement("style");return U.textContent=".warning[data-v-956d4c6f]{color:#dc143c}.icon-column[data-v-956d4c6f]{display:flex;align-items:center}.icon-column .el-icon[data-v-956d4c6f]{margin-right:8px}\n",document.head.appendChild(U),{setters:[function(e){a=e.g,o=e.d,u=e.u,i=e.e,c=e.f},function(e){f=e.default},function(e){d=e.W},function(e){s=e.c},function(e){p=e._,m=e.a,v=e.r,h=e.g,y=e.o,g=e.c,w=e.b,b=e.d,x=e.w,_=e.h,V=e.y,k=e.l,I=e.m,O=e.e,j=e.Q,D=e.i,L=e.p,E=e.q}],execute:function(){var r=function(e){return L("data-v-956d4c6f"),e=e(),E(),e},U={class:"gva-table-box"},P={class:"gva-btn-list"},C={key:0,class:"icon-column"},N={style:{display:"inline-flex"}},S=r((function(){return w("span",{style:{"font-size":"12px","margin-right":"12px"}},"如果菜单包含子菜单，请创建router-view二级路由页面或者",-1)})),z=r((function(){return w("span",null," 高亮菜单 ",-1)})),T=r((function(){return w("span",null," 是否为基础页面 ",-1)})),B={class:"dialog-footer"},A={name:"Menus"},q=Object.assign(A,{setup:function(t){var r=m({path:[{required:!0,message:"请输入菜单name",trigger:"blur"}],component:[{required:!0,message:"请输入文件路径",trigger:"blur"}],"meta.title":[{required:!0,message:"请输入菜单展示名称",trigger:"blur"}]}),p=v(1),L=v(0),E=v(999),A=v([]),q=v({}),F=function(){var t=l(e().mark((function t(){var r;return e().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,a(n({page:p.value,pageSize:E.value},q.value));case 2:0===(r=e.sent).code&&(A.value=r.data.list,L.value=r.data.total,p.value=r.data.page,E.value=r.data.pageSize);case 4:case"end":return e.stop()}}),t)})));return function(){return t.apply(this,arguments)}}();F();var G=function(){Q.value.component=Q.value.component.replace(/\\/g,"/")},M=function(){var t=l(e().mark((function t(n,r){var a;return e().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(0!==(a=n[r]).ID){e.next=4;break}return n.splice(r,1),e.abrupt("return");case 4:return e.next=6,s({id:a.ID});case 6:if(0!==e.sent.code){e.next=10;break}return n.splice(r,1),e.abrupt("return");case 10:case"end":return e.stop()}}),t)})));return function(e,n){return t.apply(this,arguments)}}(),Q=v({ID:0,path:"",name:"",hidden:!1,parentId:"",component:"",meta:{activeName:"",title:"",icon:"",defaultMenu:!1,closeTab:!1,keepAlive:!1},parameters:[],menuBtn:[]}),Y=function(){Q.value.path=Q.value.name},$=function(e){W(),e()},H=v(null),K=v(!1),W=function(){K.value=!1,H.value.resetFields(),Q.value={ID:0,path:"",name:"",hidden:!1,parentId:"",component:"",meta:{title:"",icon:"",defaultMenu:!1,closeTab:!1,keepAlive:!1}}},J=v(!1),R=function(){W(),J.value=!1},X=function(){var t=l(e().mark((function t(){return e().wrap((function(t){for(;;)switch(t.prev=t.next){case 0:H.value.validate(function(){var t=l(e().mark((function t(n){var r;return e().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!n){e.next=13;break}if(!ne.value){e.next=7;break}return e.next=4,u(Q.value);case 4:r=e.sent,e.next=10;break;case 7:return e.next=9,i(Q.value);case 9:r=e.sent;case 10:0===r.code&&(D({type:"success",message:ne.value?"编辑成功":"添加成功!"}),F()),W(),J.value=!1;case 13:case"end":return e.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}());case 1:case"end":return t.stop()}}),t)})));return function(){return t.apply(this,arguments)}}(),Z=v([{ID:"0",title:"根菜单"}]),ee=function(){Z.value=[{ID:"0",title:"根目录"}],te(A.value,Z.value,!1)},te=function e(t,n,r){t&&t.forEach((function(t){if(t.children&&t.children.length){var a={title:t.meta.title,ID:String(t.ID),disabled:r||t.ID===Q.value.ID,children:[]};e(t.children,a.children,r||t.ID===Q.value.ID),n.push(a)}else{var l={title:t.meta.title,ID:String(t.ID),disabled:r||t.ID===Q.value.ID};n.push(l)}}))},ne=v(!1),re=v("新增菜单"),ae=function(e){re.value="新增菜单",Q.value.parentId=String(e),ne.value=!1,ee(),J.value=!0},le=function(){var t=l(e().mark((function t(n){var r;return e().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return re.value="编辑菜单",e.next=3,c({id:n});case 3:r=e.sent,Q.value=r.data.menu,ne.value=!0,ee(),J.value=!0;case 8:case"end":return e.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}();return function(t,n){var a=h("el-button"),u=h("el-table-column"),i=h("el-icon"),c=h("el-table"),s=h("el-input"),m=h("el-form-item"),v=h("el-checkbox"),L=h("el-option"),E=h("el-select"),q=h("el-cascader"),W=h("QuestionFilled"),ee=h("el-tooltip"),te=h("el-form"),oe=h("el-dialog");return y(),g("div",null,[w("div",U,[w("div",P,[b(a,{size:"small",type:"primary",icon:"plus",onClick:n[0]||(n[0]=function(e){return ae("0")})},{default:x((function(){return[_("新增根菜单")]})),_:1})]),b(c,{data:A.value,"row-key":"ID"},{default:x((function(){return[b(u,{align:"left",label:"ID","min-width":"100",prop:"ID"}),b(u,{align:"left",label:"展示名称","min-width":"120",prop:"authorityName"},{default:x((function(e){return[w("span",null,V(e.row.meta.title),1)]})),_:1}),b(u,{align:"left",label:"图标","min-width":"140",prop:"authorityName"},{default:x((function(e){return[e.row.meta.icon?(y(),g("div",C,[b(i,null,{default:x((function(){return[(y(),k(I(e.row.meta.icon)))]})),_:2},1024),w("span",null,V(e.row.meta.icon),1)])):O("",!0)]})),_:1}),b(u,{align:"left",label:"路由Name","show-overflow-tooltip":"","min-width":"160",prop:"name"}),b(u,{align:"left",label:"路由Path","show-overflow-tooltip":"","min-width":"160",prop:"path"}),b(u,{align:"left",label:"是否隐藏","min-width":"100",prop:"hidden"},{default:x((function(e){return[w("span",null,V(e.row.hidden?"隐藏":"显示"),1)]})),_:1}),b(u,{align:"left",label:"父节点","min-width":"90",prop:"parentId"}),b(u,{align:"left",label:"排序","min-width":"70",prop:"sort"}),b(u,{align:"left",label:"文件路径","min-width":"360",prop:"component"}),b(u,{align:"left",fixed:"right",label:"操作",width:"300"},{default:x((function(t){return[b(a,{size:"small",type:"primary",link:"",icon:"plus",onClick:function(e){return ae(t.row.ID)}},{default:x((function(){return[_("添加子菜单")]})),_:2},1032,["onClick"]),b(a,{size:"small",type:"primary",link:"",icon:"edit",onClick:function(e){return le(t.row.ID)}},{default:x((function(){return[_("编辑")]})),_:2},1032,["onClick"]),b(a,{size:"small",type:"primary",link:"",icon:"delete",onClick:function(n){return r=t.row.ID,void j.confirm("此操作将永久删除所有角色下该菜单, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(l(e().mark((function t(){return e().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,o({ID:r});case 2:0===e.sent.code&&(D({type:"success",message:"删除成功!"}),1===A.value.length&&p.value>1&&p.value--,F());case 4:case"end":return e.stop()}}),t)})))).catch((function(){D({type:"info",message:"已取消删除"})}));var r}},{default:x((function(){return[_("删除")]})),_:2},1032,["onClick"])]})),_:1})]})),_:1},8,["data"])]),b(oe,{modelValue:J.value,"onUpdate:modelValue":n[16]||(n[16]=function(e){return J.value=e}),"before-close":$,title:re.value},{footer:x((function(){return[w("div",B,[b(a,{size:"small",onClick:R},{default:x((function(){return[_("取 消")]})),_:1}),b(a,{size:"small",type:"primary",onClick:X},{default:x((function(){return[_("确 定")]})),_:1})])]})),default:x((function(){return[b(d,{title:"新增菜单，需要在角色管理内配置权限才可使用"}),J.value?(y(),k(te,{key:0,ref_key:"menuForm",ref:H,inline:!0,model:Q.value,rules:r,"label-position":"top","label-width":"85px"},{default:x((function(){return[b(m,{label:"路由Name",prop:"path",style:{width:"30%"}},{default:x((function(){return[b(s,{modelValue:Q.value.name,"onUpdate:modelValue":n[1]||(n[1]=function(e){return Q.value.name=e}),autocomplete:"off",placeholder:"唯一英文字符串",onChange:Y},null,8,["modelValue"])]})),_:1}),b(m,{prop:"path",style:{width:"30%"}},{label:x((function(){return[w("div",N,[_(" 路由Path "),b(v,{modelValue:K.value,"onUpdate:modelValue":n[2]||(n[2]=function(e){return K.value=e}),style:{float:"right","margin-left":"20px"}},{default:x((function(){return[_("添加参数")]})),_:1},8,["modelValue"])])]})),default:x((function(){return[b(s,{modelValue:Q.value.path,"onUpdate:modelValue":n[3]||(n[3]=function(e){return Q.value.path=e}),disabled:!K.value,autocomplete:"off",placeholder:"建议只在后方拼接参数"},null,8,["modelValue","disabled"])]})),_:1}),b(m,{label:"是否隐藏",style:{width:"30%"}},{default:x((function(){return[b(E,{modelValue:Q.value.hidden,"onUpdate:modelValue":n[4]||(n[4]=function(e){return Q.value.hidden=e}),placeholder:"是否在列表隐藏"},{default:x((function(){return[b(L,{value:!1,label:"否"}),b(L,{value:!0,label:"是"})]})),_:1},8,["modelValue"])]})),_:1}),b(m,{label:"父节点ID",style:{width:"30%"}},{default:x((function(){return[b(q,{modelValue:Q.value.parentId,"onUpdate:modelValue":n[5]||(n[5]=function(e){return Q.value.parentId=e}),style:{width:"100%"},disabled:!ne.value,options:Z.value,props:{checkStrictly:!0,label:"title",value:"ID",disabled:"disabled",emitPath:!1},"show-all-levels":!1,filterable:""},null,8,["modelValue","disabled","options"])]})),_:1}),b(m,{label:"文件路径",prop:"component",style:{width:"60%"}},{default:x((function(){return[b(s,{modelValue:Q.value.component,"onUpdate:modelValue":n[6]||(n[6]=function(e){return Q.value.component=e}),autocomplete:"off",placeholder:"页面:view/xxx/xx.vue 插件:plugin/xx/xx.vue",onBlur:G},null,8,["modelValue"]),S,b(a,{style:{"margin-top":"4px"},size:"small",onClick:n[7]||(n[7]=function(e){return Q.value.component="view/routerHolder.vue"})},{default:x((function(){return[_("点我设置")]})),_:1})]})),_:1}),b(m,{label:"展示名称",prop:"meta.title",style:{width:"30%"}},{default:x((function(){return[b(s,{modelValue:Q.value.meta.title,"onUpdate:modelValue":n[8]||(n[8]=function(e){return Q.value.meta.title=e}),autocomplete:"off"},null,8,["modelValue"])]})),_:1}),b(m,{label:"图标",prop:"meta.icon",style:{width:"30%"}},{default:x((function(){return[b(f,{meta:Q.value.meta,style:{width:"100%"}},null,8,["meta"])]})),_:1}),b(m,{label:"排序标记",prop:"sort",style:{width:"30%"}},{default:x((function(){return[b(s,{modelValue:Q.value.sort,"onUpdate:modelValue":n[9]||(n[9]=function(e){return Q.value.sort=e}),modelModifiers:{number:!0},autocomplete:"off"},null,8,["modelValue"])]})),_:1}),b(m,{prop:"meta.activeName",style:{width:"30%"}},{label:x((function(){return[w("div",null,[z,b(ee,{content:"注：当到达此路由时候，指定左侧菜单指定name会处于活跃状态（亮起），可为空，为空则为本路由Name。",placement:"top",effect:"light"},{default:x((function(){return[b(i,null,{default:x((function(){return[b(W)]})),_:1})]})),_:1})])]})),default:x((function(){return[b(s,{modelValue:Q.value.meta.activeName,"onUpdate:modelValue":n[10]||(n[10]=function(e){return Q.value.meta.activeName=e}),placeholder:Q.value.name,autocomplete:"off"},null,8,["modelValue","placeholder"])]})),_:1}),b(m,{label:"KeepAlive",prop:"meta.keepAlive",style:{width:"30%"}},{default:x((function(){return[b(E,{modelValue:Q.value.meta.keepAlive,"onUpdate:modelValue":n[11]||(n[11]=function(e){return Q.value.meta.keepAlive=e}),style:{width:"100%"},placeholder:"是否keepAlive缓存页面"},{default:x((function(){return[b(L,{value:!1,label:"否"}),b(L,{value:!0,label:"是"})]})),_:1},8,["modelValue"])]})),_:1}),b(m,{label:"CloseTab",prop:"meta.closeTab",style:{width:"30%"}},{default:x((function(){return[b(E,{modelValue:Q.value.meta.closeTab,"onUpdate:modelValue":n[12]||(n[12]=function(e){return Q.value.meta.closeTab=e}),style:{width:"100%"},placeholder:"是否自动关闭tab"},{default:x((function(){return[b(L,{value:!1,label:"否"}),b(L,{value:!0,label:"是"})]})),_:1},8,["modelValue"])]})),_:1}),b(m,{style:{width:"30%"}},{label:x((function(){return[w("div",null,[T,b(ee,{content:"此项选择为是，则不会展示左侧菜单以及顶部信息。",placement:"top",effect:"light"},{default:x((function(){return[b(i,null,{default:x((function(){return[b(W)]})),_:1})]})),_:1})])]})),default:x((function(){return[b(E,{modelValue:Q.value.meta.defaultMenu,"onUpdate:modelValue":n[13]||(n[13]=function(e){return Q.value.meta.defaultMenu=e}),style:{width:"100%"},placeholder:"是否为基础页面"},{default:x((function(){return[b(L,{value:!1,label:"否"}),b(L,{value:!0,label:"是"})]})),_:1},8,["modelValue"])]})),_:1})]})),_:1},8,["model","rules"])):O("",!0),w("div",null,[b(a,{size:"small",type:"primary",icon:"edit",onClick:n[14]||(n[14]=function(e){return function(e){e.parameters||(e.parameters=[]),e.parameters.push({type:"query",key:"",value:""})}(Q.value)})},{default:x((function(){return[_("新增菜单参数")]})),_:1}),b(c,{data:Q.value.parameters,style:{width:"100%"}},{default:x((function(){return[b(u,{align:"left",prop:"type",label:"参数类型",width:"180"},{default:x((function(e){return[b(E,{modelValue:e.row.type,"onUpdate:modelValue":function(t){return e.row.type=t},placeholder:"请选择"},{default:x((function(){return[b(L,{key:"query",value:"query",label:"query"}),b(L,{key:"params",value:"params",label:"params"})]})),_:2},1032,["modelValue","onUpdate:modelValue"])]})),_:1}),b(u,{align:"left",prop:"key",label:"参数key",width:"180"},{default:x((function(e){return[w("div",null,[b(s,{modelValue:e.row.key,"onUpdate:modelValue":function(t){return e.row.key=t}},null,8,["modelValue","onUpdate:modelValue"])])]})),_:1}),b(u,{align:"left",prop:"value",label:"参数值"},{default:x((function(e){return[w("div",null,[b(s,{modelValue:e.row.value,"onUpdate:modelValue":function(t){return e.row.value=t}},null,8,["modelValue","onUpdate:modelValue"])])]})),_:1}),b(u,{align:"left"},{default:x((function(e){return[w("div",null,[b(a,{type:"danger",size:"small",icon:"delete",onClick:function(t){return n=Q.value.parameters,r=e.$index,void n.splice(r,1);var n,r}},{default:x((function(){return[_("删除")]})),_:2},1032,["onClick"])])]})),_:1})]})),_:1},8,["data"]),b(a,{style:{"margin-top":"12px"},size:"small",type:"primary",icon:"edit",onClick:n[15]||(n[15]=function(e){return function(e){e.menuBtn||(e.menuBtn=[]),e.menuBtn.push({name:"",desc:""})}(Q.value)})},{default:x((function(){return[_("新增可控按钮")]})),_:1}),b(c,{data:Q.value.menuBtn,style:{width:"100%"}},{default:x((function(){return[b(u,{align:"left",prop:"name",label:"按钮名称",width:"180"},{default:x((function(e){return[w("div",null,[b(s,{modelValue:e.row.name,"onUpdate:modelValue":function(t){return e.row.name=t}},null,8,["modelValue","onUpdate:modelValue"])])]})),_:1}),b(u,{align:"left",prop:"name",label:"备注",width:"180"},{default:x((function(e){return[w("div",null,[b(s,{modelValue:e.row.desc,"onUpdate:modelValue":function(t){return e.row.desc=t}},null,8,["modelValue","onUpdate:modelValue"])])]})),_:1}),b(u,{align:"left"},{default:x((function(e){return[w("div",null,[b(a,{type:"danger",size:"small",icon:"delete",onClick:function(t){return M(Q.value.menuBtn,e.$index)}},{default:x((function(){return[_("删除")]})),_:2},1032,["onClick"])])]})),_:1})]})),_:1},8,["data"])])]})),_:1},8,["modelValue","title"])])}}});t("default",p(q,[["__scopeId","data-v-956d4c6f"]]))}}}))}();
