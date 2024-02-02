"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[2008],{5788:(e,r,n)=>{n.d(r,{Iu:()=>u,yg:()=>d});var t=n(1504);function a(e,r,n){return r in e?Object.defineProperty(e,r,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[r]=n,e}function l(e,r){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);r&&(t=t.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),n.push.apply(n,t)}return n}function o(e){for(var r=1;r<arguments.length;r++){var n=null!=arguments[r]?arguments[r]:{};r%2?l(Object(n),!0).forEach((function(r){a(e,r,n[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(n,r))}))}return e}function i(e,r){if(null==e)return{};var n,t,a=function(e,r){if(null==e)return{};var n,t,a={},l=Object.keys(e);for(t=0;t<l.length;t++)n=l[t],r.indexOf(n)>=0||(a[n]=e[n]);return a}(e,r);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(t=0;t<l.length;t++)n=l[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=t.createContext({}),c=function(e){var r=t.useContext(s),n=r;return e&&(n="function"==typeof e?e(r):o(o({},r),e)),n},u=function(e){var r=c(e.components);return t.createElement(s.Provider,{value:r},e.children)},p="mdxType",g={inlineCode:"code",wrapper:function(e){var r=e.children;return t.createElement(t.Fragment,{},r)}},f=t.forwardRef((function(e,r){var n=e.components,a=e.mdxType,l=e.originalType,s=e.parentName,u=i(e,["components","mdxType","originalType","parentName"]),p=c(n),f=a,d=p["".concat(s,".").concat(f)]||p[f]||g[f]||l;return n?t.createElement(d,o(o({ref:r},u),{},{components:n})):t.createElement(d,o({ref:r},u))}));function d(e,r){var n=arguments,a=r&&r.mdxType;if("string"==typeof e||a){var l=n.length,o=new Array(l);o[0]=f;var i={};for(var s in r)hasOwnProperty.call(r,s)&&(i[s]=r[s]);i.originalType=e,i[p]="string"==typeof e?e:a,o[1]=i;for(var c=2;c<l;c++)o[c]=n[c];return t.createElement.apply(null,o)}return t.createElement.apply(null,n)}f.displayName="MDXCreateElement"},8357:(e,r,n)=>{n.r(r),n.d(r,{assets:()=>s,contentTitle:()=>o,default:()=>g,frontMatter:()=>l,metadata:()=>i,toc:()=>c});var t=n(5072),a=(n(1504),n(5788));const l={slug:"enable",sidebar_position:1},o="enable",i={unversionedId:"rrsa/enable",id:"version-v0.13.0/rrsa/enable",title:"enable",description:"Enable the RRSA feature for a specific cluster.",source:"@site/versioned_docs/version-v0.13.0/rrsa/enable.md",sourceDirName:"rrsa",slug:"/rrsa/enable",permalink:"/ack-ram-tool/v0.13.0/rrsa/enable",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.13.0/rrsa/enable.md",tags:[],version:"v0.13.0",sidebarPosition:1,frontMatter:{slug:"enable",sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"rrsa",permalink:"/ack-ram-tool/v0.13.0/category/rrsa"},next:{title:"enable\uff08\u4e2d\u6587\uff09",permalink:"/ack-ram-tool/v0.13.0/zh-CN/rrsa/enable"}},s={},c=[{value:"Usage",id:"usage",level:2},{value:"Flags",id:"flags",level:2}],u={toc:c},p="wrapper";function g(e){let{components:r,...n}=e;return(0,a.yg)(p,(0,t.c)({},u,n,{components:r,mdxType:"MDXLayout"}),(0,a.yg)("h1",{id:"enable"},"enable"),(0,a.yg)("p",null,"Enable the RRSA feature for a specific cluster."),(0,a.yg)("h2",{id:"usage"},"Usage"),(0,a.yg)("pre",null,(0,a.yg)("code",{parentName:"pre",className:"language-shell"},"$ ack-ram-tool rrsa enable --cluster-id <clusterId>\n\n? Are you sure you want to enable RRSA feature? Yes\n2023-04-20T14:30:40+08:00 INFO Enable RRSA feature for cluster c86fdd*** successfully\n")),(0,a.yg)("h2",{id:"flags"},"Flags"),(0,a.yg)("pre",null,(0,a.yg)("code",{parentName:"pre"},'Usage:\n  ack-ram-tool rrsa enable [flags]\n\nFlags:\n  -c, --cluster-id string   The cluster id to use\n  -h, --help                help for enable\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n')),(0,a.yg)("p",null,"Descriptions\uff1a"),(0,a.yg)("table",null,(0,a.yg)("thead",{parentName:"table"},(0,a.yg)("tr",{parentName:"thead"},(0,a.yg)("th",{parentName:"tr",align:null},"Flag"),(0,a.yg)("th",{parentName:"tr",align:null},"Default"),(0,a.yg)("th",{parentName:"tr",align:null},"Required"),(0,a.yg)("th",{parentName:"tr",align:null},"Description"))),(0,a.yg)("tbody",{parentName:"table"},(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},"-c, --cluster-id"),(0,a.yg)("td",{parentName:"tr",align:null}),(0,a.yg)("td",{parentName:"tr",align:null},"Yes"),(0,a.yg)("td",{parentName:"tr",align:null},"Cluster ID")))))}g.isMDXComponent=!0}}]);