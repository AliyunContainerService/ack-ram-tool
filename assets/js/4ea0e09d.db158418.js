"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[3544],{5788:(e,t,n)=>{n.d(t,{Iu:()=>p,yg:()=>u});var r=n(1504);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var g=r.createContext({}),d=function(e){var t=r.useContext(g),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},p=function(e){var t=d(e.components);return r.createElement(g.Provider,{value:t},e.children)},y="mdxType",m={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},s=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,g=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),y=d(n),s=a,u=y["".concat(g,".").concat(s)]||y[s]||m[s]||l;return n?r.createElement(u,i(i({ref:t},p),{},{components:n})):r.createElement(u,i({ref:t},p))}));function u(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,i=new Array(l);i[0]=s;var o={};for(var g in t)hasOwnProperty.call(t,g)&&(o[g]=t[g]);o.originalType=e,o[y]="string"==typeof e?e:a,i[1]=o;for(var d=2;d<l;d++)i[d]=n[d];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}s.displayName="MDXCreateElement"},5488:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>g,contentTitle:()=>i,default:()=>m,frontMatter:()=>l,metadata:()=>o,toc:()=>d});var r=n(5072),a=(n(1504),n(5788));const l={slug:"/zh-CN/permissions"},i="\u6743\u9650",o={unversionedId:"permission.zh-CN",id:"version-v0.13.2/permission.zh-CN",title:"\u6743\u9650",description:"\u4e3a\u4e86\u6b63\u5e38\u4f7f\u7528 ack-ram-tool\uff0c\u60a8\u9700\u8981\u4e3a\u4f7f\u7528\u6539\u5de5\u5177\u7684\u963f\u91cc\u4e91 RAM \u7528\u6237\u6216 RAM \u89d2\u8272\u6388\u4e88\u6240\u9700\u7684 RAM \u6743\u9650\u548c RBAC \u6743\u9650\u3002",source:"@site/versioned_docs/version-v0.13.2/permission.zh-CN.md",sourceDirName:".",slug:"/zh-CN/permissions",permalink:"/ack-ram-tool/v0.13.2/zh-CN/permissions",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.13.2/permission.zh-CN.md",tags:[],version:"v0.13.2",frontMatter:{slug:"/zh-CN/permissions"},sidebar:"tutorialSidebar",previous:{title:"Permissions",permalink:"/ack-ram-tool/v0.13.2/permissions"}},g={},d=[],p={toc:d},y="wrapper";function m(e){let{components:t,...n}=e;return(0,a.yg)(y,(0,r.c)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,a.yg)("h1",{id:"\u6743\u9650"},"\u6743\u9650"),(0,a.yg)("p",null,"\u4e3a\u4e86\u6b63\u5e38\u4f7f\u7528 ack-ram-tool\uff0c\u60a8\u9700\u8981\u4e3a\u4f7f\u7528\u6539\u5de5\u5177\u7684\u963f\u91cc\u4e91 RAM \u7528\u6237\u6216 RAM \u89d2\u8272\u6388\u4e88\u6240\u9700\u7684 RAM \u6743\u9650\u548c RBAC \u6743\u9650\u3002\n\u5404\u4e2a\u5b50\u547d\u4ee4\u6240\u9700\u7684\u6700\u5c0f\u6743\u9650\u4fe1\u606f\u5982\u4e0b\u8868\u6240\u793a\uff1a"),(0,a.yg)("table",null,(0,a.yg)("thead",{parentName:"table"},(0,a.yg)("tr",{parentName:"thead"},(0,a.yg)("th",{parentName:"tr",align:null},"\u5b50\u547d\u4ee4"),(0,a.yg)("th",{parentName:"tr",align:null},"RAM \u6743\u9650"),(0,a.yg)("th",{parentName:"tr",align:null},"RBAC \u6743\u9650"))),(0,a.yg)("tbody",{parentName:"table"},(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"rrsa status")),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"rrsa enable")),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"cs:ModifyCluster")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterLogs")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"rrsa associate-role")),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:GetRole")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:CreateRole")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:UpdateRole")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"rrsa install-helper-addon")),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterAddonsVersion")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"cs:InstallClusterAddons")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"rrsa assumerole")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"rrsa disable")),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"cs:ModifyCluster")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterLogs")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"rrsa setup-addon")),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:GetRole")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:CreateRole")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:UpdateRole")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:CreatePolicy")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:ListPoliciesForRole")," ",(0,a.yg)("br",null)," ",(0,a.yg)("inlineCode",{parentName:"td"},"ram:AttachPolicyToRole")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"rrsa demo")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"credential-plugin get-kubeconfig")),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterUserKubeconfig")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"credential-plugin get-credential")),(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"cs:DescribeClusterUserKubeconfig")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"credential-plugin get-token")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")),(0,a.yg)("tr",{parentName:"tbody"},(0,a.yg)("td",{parentName:"tr",align:null},(0,a.yg)("inlineCode",{parentName:"td"},"export-credentials")),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0"),(0,a.yg)("td",{parentName:"tr",align:null},"\u65e0")))))}m.isMDXComponent=!0}}]);