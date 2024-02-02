"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[6768],{5788:(e,t,n)=>{n.d(t,{Iu:()=>c,yg:()=>y});var r=n(1504);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var g=r.createContext({}),s=function(e){var t=r.useContext(g),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},c=function(e){var t=s(e.components);return r.createElement(g.Provider,{value:t},e.children)},u="mdxType",p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,g=e.parentName,c=o(e,["components","mdxType","originalType","parentName"]),u=s(n),d=a,y=u["".concat(g,".").concat(d)]||u[d]||p[d]||l;return n?r.createElement(y,i(i({ref:t},c),{},{components:n})):r.createElement(y,i({ref:t},c))}));function y(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,i=new Array(l);i[0]=d;var o={};for(var g in t)hasOwnProperty.call(t,g)&&(o[g]=t[g]);o.originalType=e,o[u]="string"==typeof e?e:a,i[1]=o;for(var s=2;s<l;s++)i[s]=n[s];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},2964:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>g,contentTitle:()=>i,default:()=>p,frontMatter:()=>l,metadata:()=>o,toc:()=>s});var r=n(5072),a=(n(1504),n(5788));const l={slug:"global-flags"},i="Global Flags",o={unversionedId:"global-flags",id:"version-v0.15.0/global-flags",title:"Global Flags",description:"Description",source:"@site/versioned_docs/version-v0.15.0/global-flags.md",sourceDirName:".",slug:"/global-flags",permalink:"/ack-ram-tool/v0.15.0/global-flags",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.15.0/global-flags.md",tags:[],version:"v0.15.0",frontMatter:{slug:"global-flags"},sidebar:"tutorialSidebar",previous:{title:"export-credentials\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.15.0/zh-CN/export-credentials/export-credentials"},next:{title:"\u5168\u5c40\u53c2\u6570",permalink:"/ack-ram-tool/v0.15.0/zh-CN/global-flags"}},g={},s=[{value:"Description",id:"description",level:2}],c={toc:s},u="wrapper";function p(e){let{components:t,...n}=e;return(0,a.yg)(u,(0,r.c)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,a.yg)("h1",{id:"global-flags"},"Global Flags"),(0,a.yg)("pre",null,(0,a.yg)("code",{parentName:"pre"},'Global Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively (env: "ACK_RAM_TOOL_ASSUME_YES")\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli (env: "ACK_RAM_TOOL_IGNORE_ALIYUN_CLI_CREDENTIALS")\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables (env: "ACK_RAM_TOOL_IGNORE_ENV_CREDENTIALS")\n      --log-level string                log level: info, debug, error (default "info") (env: "ACK_RAM_TOOL_LOG_LEVEL")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials) (env: "ACK_RAM_TOOL_PROFILE_FILE")\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli (env: "ACK_RAM_TOOL_PROFIL_ENAME")\n      --region-id string                The region to use (default "cn-hangzhou") (env: "ACK_RAM_TOOL_REGION_ID")\n  -v, --verbose                         Make the operation more talkative\n')),(0,a.yg)("h2",{id:"description"},"Description"),(0,a.yg)("table",null,(0,a.yg)("thead",{parentName:"table"},(0,a.yg)("tr",{parentName:"thead"},(0,a.yg)("th",{parentName:"tr",align:null},"Flag"),(0,a.yg)("th",{parentName:"tr",align:null},"Default"),(0,a.yg)("th",{parentName:"tr",align:null},"Description"))),(0,a.yg)("tbody",{parentName:"table"},(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"-y, --assume-yes"),(0,a.yg)("td",{parentName:"tr",align:null},"false"),(0,a.yg)("td",{parentName:"tr",align:null},"When set to true, the program will automatically execute without asking whether to continue the operation")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"--log-level"),(0,a.yg)("td",{parentName:"tr",align:null},"info"),(0,a.yg)("td",{parentName:"tr",align:null},"Log level")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"--ignore-aliyun-cli-credentials"),(0,a.yg)("td",{parentName:"tr",align:null},"false"),(0,a.yg)("td",{parentName:"tr",align:null},"When set to true, the aliyun cli configuration file will be ignored when searching for credential information")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"--ignore-env-credentials"),(0,a.yg)("td",{parentName:"tr",align:null},"false"),(0,a.yg)("td",{parentName:"tr",align:null},"When set to true, the credential information in the environment variables will be ignored when searching for credential information")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"--profile-file"),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"~/.aliyun/config.json")," or ",(0,a.yg)("inlineCode",{parentName:"td"},"~/.alibabacloud/credentials")),(0,a.yg)("td",{parentName:"tr",align:null},"Specify the credential configuration file")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"--profile-name"),(0,a.yg)("td",{parentName:"tr",align:null},"no default"),(0,a.yg)("td",{parentName:"tr",align:null},"When using the aliyun cli configuration file, use the credential configuration defined in the specified configuration set")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"--region-id"),(0,a.yg)("td",{parentName:"tr",align:null},"cn-hangzhou"),(0,a.yg)("td",{parentName:"tr",align:null},"Region information used when accessing OpenAPI")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"-v, --verbose"),(0,a.yg)("td",{parentName:"tr",align:null},"false"),(0,a.yg)("td",{parentName:"tr",align:null},"Quickly enable debug mode")))))}p.isMDXComponent=!0}}]);