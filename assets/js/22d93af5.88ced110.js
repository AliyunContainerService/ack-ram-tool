"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[2721],{8270:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>s,contentTitle:()=>a,default:()=>u,frontMatter:()=>c,metadata:()=>l,toc:()=>o});var i=t(4848),r=t(8453);const c={slug:"/zh-CN/credential-plugin/kubeconfig",title:"get-kubeconfig\uff08\u4e2d\u6587\uff09",sidebar_position:1},a="get-kubeconfig",l={id:"credential-plugin/get-kubeconfig.zh-CN",title:"get-kubeconfig\uff08\u4e2d\u6587\uff09",description:"\u83b7\u53d6\u4f7f\u7528 ack-ram-tool \u4f5c\u4e3a credential plugin \u7684 kubeconfig\u3002",source:"@site/versioned_docs/version-v0.13.2/credential-plugin/get-kubeconfig.zh-CN.md",sourceDirName:"credential-plugin",slug:"/zh-CN/credential-plugin/kubeconfig",permalink:"/ack-ram-tool/v0.13.2/zh-CN/credential-plugin/kubeconfig",draft:!1,unlisted:!1,editUrl:"https://github.com/AliyunContainerService/ack-ram-tool/edit/master/website/versioned_docs/version-v0.13.2/credential-plugin/get-kubeconfig.zh-CN.md",tags:[],version:"v0.13.2",sidebarPosition:1,frontMatter:{slug:"/zh-CN/credential-plugin/kubeconfig",title:"get-kubeconfig\uff08\u4e2d\u6587\uff09",sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"get-kubeconfig",permalink:"/ack-ram-tool/v0.13.2/credential-plugin/get-kubeconfig"},next:{title:"get-credential",permalink:"/ack-ram-tool/v0.13.2/credential-plugin/get-credential"}},s={},o=[{value:"\u4f7f\u7528\u793a\u4f8b",id:"\u4f7f\u7528\u793a\u4f8b",level:2},{value:"--mode ram-authenticator-token",id:"--mode-ram-authenticator-token",level:3},{value:"\u547d\u4ee4\u884c\u53c2\u6570",id:"\u547d\u4ee4\u884c\u53c2\u6570",level:2}];function d(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",ul:"ul",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.h1,{id:"get-kubeconfig",children:"get-kubeconfig"}),"\n",(0,i.jsxs)(n.p,{children:["\u83b7\u53d6\u4f7f\u7528 ack-ram-tool \u4f5c\u4e3a ",(0,i.jsx)(n.a,{href:"https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins",children:"credential plugin"})," \u7684 kubeconfig\u3002"]}),"\n",(0,i.jsx)(n.p,{children:"\u5305\u542b\u5982\u4e0b\u7279\u6027\uff1a"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"\u8bc1\u4e66\u8fc7\u671f\u524d\u5c06\u81ea\u52a8\u83b7\u53d6\u65b0\u7684\u8bc1\u4e66\u3002"}),"\n",(0,i.jsx)(n.li,{children:"\u652f\u6301\u4f7f\u7528\u4e34\u65f6\u8bc1\u4e66\u3002"}),"\n",(0,i.jsx)(n.li,{children:"\u96c6\u6210 ack-ram-authenticator\u3002"}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"\u4f7f\u7528\u793a\u4f8b",children:"\u4f7f\u7528\u793a\u4f8b"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-shell",children:'$ ack-ram-tool credential-plugin get-kubeconfig --cluster-id c5e***\n\nkind: Config\napiVersion: v1\nclusters:\n    - name: kubernetes\n      cluster:\n        server: https://106.*.*.*:6443\n        certificate-authority-data: LS0tL***\ncontexts:\n    - name: 272***-c5e***\n      context:\n        cluster: kubernetes\n        user: "272***"\ncurrent-context: 272***-c5e***\nusers:\n    - name: "272***"\n      user:\n        exec:\n            command: ack-ram-tool\n            args:\n                - credential-plugin\n                - get-credential\n                - --cluster-id\n                - c5e***\n                - --api-version\n                - v1beta1\n                - --expiration\n                - 3h\n                - --log-level\n                - error\n            apiVersion: client.authentication.k8s.io/v1beta1\n            provideClusterInfo: false\n            interactiveMode: Never\npreferences: {}\n\n$ ack-ram-tool credential-plugin get-kubeconfig --cluster-id c5e*** > kubeconfig\n$ kubectl --kubeconfig kubeconfig get ns\nNAME                         STATUS   AGE\ndefault                      Active   6d3h\nkube-node-lease              Active   6d3h\nkube-public                  Active   6d3h\nkube-system                  Active   6d3h\n'})}),"\n",(0,i.jsx)(n.h3,{id:"--mode-ram-authenticator-token",children:"--mode ram-authenticator-token"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{children:'$ ack-ram-tool credential-plugin get-kubeconfig --mode ram-authenticator-token --cluster-id c5e***\n\nkind: Config\napiVersion: v1\nclusters:\n    - name: kubernetes\n      cluster:\n        server: https://106.*.*.*:6443\n        certificate-authority-data: LS0t***\ncontexts:\n    - name: 272***-c5e***\n      context:\n        cluster: kubernetes\n        user: "272***"\ncurrent-context: 272***-c5e***\nusers:\n    - name: "272***"\n      user:\n        exec:\n            command: ack-ram-tool\n            args:\n                - credential-plugin\n                - get-token\n                - --cluster-id\n                - c5e***\n                - --api-version\n                - v1beta1\n                - --log-level\n                - error\n            apiVersion: client.authentication.k8s.io/v1beta1\n            provideClusterInfo: false\n            interactiveMode: Never\npreferences: {}\n\n'})}),"\n",(0,i.jsx)(n.h2,{id:"\u547d\u4ee4\u884c\u53c2\u6570",children:"\u547d\u4ee4\u884c\u53c2\u6570"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{children:'Usage:\n  ack-ram-tool credential-plugin get-kubeconfig [flags]\n\nFlags:\n      --api-version string            v1 or v1beta1 (default "v1beta1")\n  -c, --cluster-id string             The cluster id to use\n      --credential-cache-dir string   Directory to cache certificate (default "~/.kube/cache/ack-ram-tool/credential-plugin")\n      --expiration duration           The certificate expiration (default 3h0m0s)\n  -h, --help                          help for get-kubeconfig\n  -m, --mode string                   credential mode: certificate or ram-authenticator-token (default "certificate")\n      --private-address               Use private ip as api-server address\n\nGlobal Flags:\n  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively\n      --ignore-aliyun-cli-credentials   don\'t try to parse credentials from config.json of aliyun cli\n      --ignore-env-credentials          don\'t try to parse credentials from environment variables\n      --log-level string                log level: info, debug, error (default "info")\n      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)\n      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli\n      --region-id string                The region to use (default "cn-hangzhou")\n'})}),"\n",(0,i.jsx)(n.p,{children:"\u53c2\u6570\u8bf4\u660e\uff1a"}),"\n",(0,i.jsxs)(n.table,{children:[(0,i.jsx)(n.thead,{children:(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.th,{children:"\u53c2\u6570\u540d\u79f0"}),(0,i.jsx)(n.th,{children:"\u9ed8\u8ba4\u503c"}),(0,i.jsx)(n.th,{children:"\u5fc5\u9700\u53c2\u6570"}),(0,i.jsx)(n.th,{children:"\u8bf4\u660e"})]})}),(0,i.jsxs)(n.tbody,{children:[(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"-c, --cluster-id"}),(0,i.jsx)(n.td,{children:"\u65e0"}),(0,i.jsx)(n.td,{children:"\u662f"}),(0,i.jsx)(n.td,{children:"\u96c6\u7fa4 ID"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"-m, --mode"}),(0,i.jsx)(n.td,{children:"certificate"}),(0,i.jsx)(n.td,{children:"\u5426"}),(0,i.jsxs)(n.td,{children:["kubeconfig \u4e2d\u7684\u8ba4\u8bc1\u65b9\u6cd5\uff1a ",(0,i.jsx)(n.code,{children:"certificate"})," \u8868\u793a\u8bc1\u4e66\u8ba4\u8bc1\uff0c",(0,i.jsx)(n.code,{children:"ram-authenticator-token"})," \u8868\u793a\u57fa\u4e8e ack-ram-authenticator \u7684 token \u8ba4\u8bc1"]})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"--expiration"}),(0,i.jsx)(n.td,{children:"3h"}),(0,i.jsx)(n.td,{children:"\u5426"}),(0,i.jsxs)(n.td,{children:["--mode \u88ab\u8bbe\u7f6e\u4e3a ",(0,i.jsx)(n.code,{children:"certificate"})," \u65f6\uff0c\u901a\u8fc7\u8fd9\u4e2a\u53c2\u6570\u8bbe\u7f6e\u8bc1\u4e66\u8fc7\u671f\u65f6\u95f4\u3002\u4e3a 0 \u65f6\u8868\u793a\u4e0d\u4f7f\u7528\u4e34\u65f6\u8bc1\u4e66\u800c\u662f\u4f7f\u7528\u6709\u6548\u671f\u66f4\u957f\u7684\u8bc1\u4e66\uff08\u8fc7\u671f\u65f6\u95f4\u7531\u670d\u52a1\u7aef\u81ea\u52a8\u786e\u5b9a\uff09"]})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"--private-address"}),(0,i.jsx)(n.td,{children:"false"}),(0,i.jsx)(n.td,{children:"\u5426"}),(0,i.jsx)(n.td,{children:"\u662f\u5426\u4f7f\u7528\u5185\u7f51 api server \u5730\u5740"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"--api-version"}),(0,i.jsx)(n.td,{children:"v1beta1"}),(0,i.jsx)(n.td,{children:"\u5426"}),(0,i.jsxs)(n.td,{children:["\u6307\u5b9a\u8fd4\u56de\u7684\u6570\u636e\u4e2d\u4f7f\u7528\u54ea\u4e2a\u7248\u672c\u7684 apiVersion\u3002v1beta1 \u8868\u793a ",(0,i.jsx)(n.code,{children:"client.authentication.k8s.io/v1beta1"}),"\uff0cv1 \u8868\u793a ",(0,i.jsx)(n.code,{children:"client.authentication.k8s.io/v1beta1"})]})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:"--credential-cache-dir"}),(0,i.jsx)(n.td,{children:(0,i.jsx)(n.code,{children:"~/.kube/cache/ack-ram-tool/credential-plugin"})}),(0,i.jsx)(n.td,{children:"\u5426"}),(0,i.jsxs)(n.td,{children:["\u7528\u4e8e\u7f13\u5b58\u8bc1\u4e66\u7684\u76ee\u5f55\uff0c\u53ea\u5728 ",(0,i.jsx)(n.code,{children:"--mode"})," \u88ab\u8bbe\u7f6e\u4e3a ",(0,i.jsx)(n.code,{children:"certificate"})," \u65f6\u6709\u6548"]})]})]})]})]})}function u(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},8453:(e,n,t)=>{t.d(n,{R:()=>a,x:()=>l});var i=t(6540);const r={},c=i.createContext(r);function a(e){const n=i.useContext(c);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function l(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:a(e.components),i.createElement(c.Provider,{value:n},e.children)}}}]);