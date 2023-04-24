"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[3655],{3905:(e,t,n)=>{n.d(t,{Zo:()=>m,kt:()=>c});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),d=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},m=function(e){var t=d(e.components);return r.createElement(s.Provider,{value:t},e.children)},p="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},k=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,s=e.parentName,m=o(e,["components","mdxType","originalType","parentName"]),p=d(n),k=a,c=p["".concat(s,".").concat(k)]||p[k]||u[k]||l;return n?r.createElement(c,i(i({ref:t},m),{},{components:n})):r.createElement(c,i({ref:t},m))}));function c(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,i=new Array(l);i[0]=k;var o={};for(var s in t)hasOwnProperty.call(t,s)&&(o[s]=t[s]);o.originalType=e,o[p]="string"==typeof e?e:a,i[1]=o;for(var d=2;d<l;d++)i[d]=n[d];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}k.displayName="MDXCreateElement"},5302:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>i,default:()=>u,frontMatter:()=>l,metadata:()=>o,toc:()=>d});var r=n(7462),a=(n(7294),n(3905));const l={slug:"permissions"},i="Permissions",o={unversionedId:"permission",id:"permission",title:"Permissions",description:"In order to use ack-ram-tool normally, you need to grant the necessary RAM permissions and RBAC permissions for",source:"@site/docs/permission.md",sourceDirName:".",slug:"/permissions",permalink:"/ack-ram-tool/next/permissions",draft:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/docs/permission.md",tags:[],version:"current",frontMatter:{slug:"permissions"},sidebar:"tutorialSidebar",previous:{title:"\u5168\u5c40\u53c2\u6570",permalink:"/ack-ram-tool/next/zh-CN/global-flags"},next:{title:"\u6743\u9650",permalink:"/ack-ram-tool/next/zh-CN/permissions"}},s={},d=[],m={toc:d},p="wrapper";function u(e){let{components:t,...n}=e;return(0,a.kt)(p,(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"permissions"},"Permissions"),(0,a.kt)("p",null,"In order to use ack-ram-tool normally, you need to grant the necessary RAM permissions and RBAC permissions for\nthe Alibaba Cloud RAM user or RAM role that uses this tool.\nThe minimum permission information required for each subcommand is shown in the following table:"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"Command"),(0,a.kt)("th",{parentName:"tr",align:null},"RAM Permissoins"),(0,a.kt)("th",{parentName:"tr",align:null},"RBAC Permissions"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"rrsa status")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"rrsa enable")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"cs:ModifyCluster")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterLogs")),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"rrsa associate-role")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:GetRole")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:CreateRole")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:UpdateRole")),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"rrsa install-helper-addon")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterAddonsVersion")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"cs:InstallClusterAddons")),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"rrsa assumerole")),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"rrsa disable")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"cs:ModifyCluster")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterLogs")),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"rrsa setup-addon")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterDetail")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:GetRole")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:CreateRole")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:UpdateRole")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:CreatePolicy")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:ListPoliciesForRole")," ",(0,a.kt)("br",null)," ",(0,a.kt)("inlineCode",{parentName:"td"},"ram:AttachPolicyToRole")),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"rrsa demo")),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"credential-plugin get-kubeconfig")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterUserKubeconfig")),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"credential-plugin get-credential")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"cs:DescribeClusterUserKubeconfig")),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"credential-plugin get-token")),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null})),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"export-credentials")),(0,a.kt)("td",{parentName:"tr",align:null}),(0,a.kt)("td",{parentName:"tr",align:null})))))}u.isMDXComponent=!0}}]);