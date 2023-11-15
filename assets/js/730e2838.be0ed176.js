"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[1069],{3905:(e,t,n)=>{n.d(t,{Zo:()=>s,kt:()=>g});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var p=r.createContext({}),u=function(e){var t=r.useContext(p),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},s=function(e){var t=u(e.components);return r.createElement(p.Provider,{value:t},e.children)},c="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,p=e.parentName,s=i(e,["components","mdxType","originalType","parentName"]),c=u(n),m=a,g=c["".concat(p,".").concat(m)]||c[m]||d[m]||l;return n?r.createElement(g,o(o({ref:t},s),{},{components:n})):r.createElement(g,o({ref:t},s))}));function g(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,o=new Array(l);o[0]=m;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i[c]="string"==typeof e?e:a,o[1]=i;for(var u=2;u<l;u++)o[u]=n[u];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},5495:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>o,default:()=>d,frontMatter:()=>l,metadata:()=>i,toc:()=>u});var r=n(7462),a=(n(7294),n(3905));const l={slug:"/zh-CN/global-flags"},o="\u5168\u5c40\u53c2\u6570",i={unversionedId:"global-flags.zh-CN",id:"version-v0.16.0/global-flags.zh-CN",title:"\u5168\u5c40\u53c2\u6570",description:"\u53c2\u6570\u8bf4\u660e",source:"@site/versioned_docs/version-v0.16.0/global-flags.zh-CN.md",sourceDirName:".",slug:"/zh-CN/global-flags",permalink:"/ack-ram-tool/zh-CN/global-flags",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.16.0/global-flags.zh-CN.md",tags:[],version:"v0.16.0",frontMatter:{slug:"/zh-CN/global-flags"}},p={},u=[{value:"\u53c2\u6570\u8bf4\u660e",id:"\u53c2\u6570\u8bf4\u660e",level:2}],s={toc:u},c="wrapper";function d(e){let{components:t,...n}=e;return(0,a.kt)(c,(0,r.Z)({},s,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"\u5168\u5c40\u53c2\u6570"},"\u5168\u5c40\u53c2\u6570"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'Global Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively (env: "ACK_RAM_TOOL_ASSUME_YES")\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli (env: "ACK_RAM_TOOL_IGNORE_ALIYUN_CLI_CREDENTIALS")\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables (env: "ACK_RAM_TOOL_IGNORE_ENV_CREDENTIALS")\n      --log-level string                log level: info, debug, error (default "info") (env: "ACK_RAM_TOOL_LOG_LEVEL")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials) (env: "ACK_RAM_TOOL_PROFILE_FILE")\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli (env: "ACK_RAM_TOOL_PROFIL_ENAME")\n      --region-id string                The region to use (default "cn-hangzhou") (env: "ACK_RAM_TOOL_REGION_ID")\n  -v, --verbose                         Make the operation more talkative\n')),(0,a.kt)("h2",{id:"\u53c2\u6570\u8bf4\u660e"},"\u53c2\u6570\u8bf4\u660e"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"\u53c2\u6570\u540d\u79f0"),(0,a.kt)("th",{parentName:"tr",align:null},"\u9ed8\u8ba4\u503c"),(0,a.kt)("th",{parentName:"tr",align:null},"\u8bf4\u660e"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-y, --assume-yes"),(0,a.kt)("td",{parentName:"tr",align:null},"false"),(0,a.kt)("td",{parentName:"tr",align:null},"\u4e3a true \u65f6\uff0c\u7a0b\u5e8f\u5c06\u81ea\u52a8\u6267\u884c\uff0c\u4e0d\u518d\u8be2\u95ee\u662f\u5426\u7ee7\u7eed\u64cd\u4f5c")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--log-level"),(0,a.kt)("td",{parentName:"tr",align:null},"info"),(0,a.kt)("td",{parentName:"tr",align:null},"\u65e5\u5fd7\u7ea7\u522b")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--ignore-aliyun-cli-credentials"),(0,a.kt)("td",{parentName:"tr",align:null},"false"),(0,a.kt)("td",{parentName:"tr",align:null},"\u4e3a true \u65f6\uff0c\u67e5\u627e\u51ed\u8bc1\u4fe1\u606f\u65f6\u5c06\u5ffd\u7565 aliyun cli \u7684\u914d\u7f6e\u6587\u4ef6")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--ignore-env-credentials"),(0,a.kt)("td",{parentName:"tr",align:null},"false"),(0,a.kt)("td",{parentName:"tr",align:null},"\u4e3a true \u65f6\uff0c\u67e5\u627e\u51ed\u8bc1\u4fe1\u606f\u65f6\u5c06\u5ffd\u7565\u73af\u5883\u53d8\u91cf\u4e2d\u7684\u51ed\u8bc1\u4fe1\u606f")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--profile-file"),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"~/.aliyun/config.json")," \u6216 ",(0,a.kt)("inlineCode",{parentName:"td"},"~/.alibabacloud/credentials")),(0,a.kt)("td",{parentName:"tr",align:null},"\u6307\u5b9a\u51ed\u8bc1\u914d\u7f6e\u6587\u4ef6")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--profile-name"),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null},"\u5f53\u4f7f\u7528 aliyun cli \u914d\u7f6e\u6587\u4ef6\u65f6\uff0c\u4f7f\u7528\u6307\u5b9a\u7684\u914d\u7f6e\u96c6\u4e2d\u5b9a\u4e49\u7684\u51ed\u8bc1\u914d\u7f6e")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"--region-id"),(0,a.kt)("td",{parentName:"tr",align:null},"cn-hangzhou"),(0,a.kt)("td",{parentName:"tr",align:null},"\u8bbf\u95ee open api \u65f6\u4f7f\u7528\u7684 region \u4fe1\u606f")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},"-v, --verbose"),(0,a.kt)("td",{parentName:"tr",align:null},"false"),(0,a.kt)("td",{parentName:"tr",align:null},"\u5feb\u901f\u542f\u7528 debug \u6a21\u5f0f")))))}d.isMDXComponent=!0}}]);