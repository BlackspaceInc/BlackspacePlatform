(window.webpackJsonp=window.webpackJsonp||[]).push([[69],{119:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return o})),n.d(t,"metadata",(function(){return l})),n.d(t,"rightToc",(function(){return s})),n.d(t,"default",(function(){return b}));var a=n(2),r=n(6),i=(n(0),n(146)),o={id:"kubernetes-patterns",sidebar_label:"Kubernetes Patterns",title:"Kubernetes Best Patterns"},l={unversionedId:"technology/checklists/backend/backend-kubernetes/kubernetes-patterns",id:"technology/checklists/backend/backend-kubernetes/kubernetes-patterns",isDocsHomePage:!1,title:"Kubernetes Best Patterns",description:"Kubernetes Patterns",source:"@site/docs/technology/checklists/backend/backend-kubernetes/patterns.md",permalink:"/technology/checklists/backend/backend-kubernetes/kubernetes-patterns",editUrl:"https://github.com/BlackspaceInc/BlackspacePlatform/edit/master/website/docs/technology/checklists/backend/backend-kubernetes/patterns.md",sidebar_label:"Kubernetes Patterns",sidebar:"docs",previous:{title:"Kubernetes Namespace & Governance Management",permalink:"/technology/checklists/backend/backend-kubernetes/kubernetes-governance"},next:{title:"Backend Production Ready Checklist",permalink:"/technology/checklists/backend/backend-production-ready/production-ready-checklist"}},s=[{value:"Chapter 2: Predictable Demands",id:"chapter-2-predictable-demands",children:[]},{value:"Chapter 3: Declarative Deployment",id:"chapter-3-declarative-deployment",children:[]},{value:"Chapter 4: Health Probe",id:"chapter-4-health-probe",children:[]},{value:"Chapter 5: Managed Lifecycle",id:"chapter-5-managed-lifecycle",children:[]},{value:"Chapter 6: Automated Placement",id:"chapter-6-automated-placement",children:[]}],c={rightToc:s};function b(e){var t=e.components,n=Object(r.a)(e,["components"]);return Object(i.b)("wrapper",Object(a.a)({},c,n,{components:t,mdxType:"MDXLayout"}),Object(i.b)("h1",{id:"kubernetes-patterns"},"Kubernetes Patterns"),Object(i.b)("p",null,"O'Reilly, 2019."),Object(i.b)("p",null,'The book is a collection of "patterns" for container-based cloud-native environments and applications, and it shows how Kubernetes implements these patterns.'),Object(i.b)("p",null,"You can learn how to optimally use the tools that Kubernetes provides to you to make the best usage of Kubernetes' implementation of these patterns. "),Object(i.b)("h2",{id:"chapter-2-predictable-demands"},"Chapter 2: Predictable Demands"),Object(i.b)("p",null,"Containers should know and declare their requirements. This allows the platform (Kubernetes) to put Pods to the right place in the cluster and treat them appropriately."),Object(i.b)("p",null,Object(i.b)("strong",{parentName:"p"},"Types of requirements:")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"Runtime dependencies",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"Persistent storage, host port, configuration (ConfigMaps, Secrets): defined in Pod spec."))),Object(i.b)("li",{parentName:"ul"},"Resource requirements",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"Requests"),Object(i.b)("li",{parentName:"ul"},"Limits"),Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"QoS")," of Pods (defined by requests and limits): defines in what order Pods are killed by the ",Object(i.b)("strong",{parentName:"li"},"kubelet"),Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"Best-effort: no requests and limits, lowest priority, Pod is killed first"),Object(i.b)("li",{parentName:"ul"},"Burstable: limits are higher than requests, middle-priority, killed if no Best-effort Pods exist"),Object(i.b)("li",{parentName:"ul"},"Guaranteed: limits equal requests, highest priority, killed only if no Best-effort and Burstable Pods exist")))))),Object(i.b)("p",null,Object(i.b)("a",Object(a.a)({parentName:"p"},{href:"https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/"}),Object(i.b)("strong",{parentName:"a"},"Pod Priority and Preemption")),":"),Object(i.b)("p",null,Object(i.b)("a",Object(a.a)({parentName:"p"},{href:"https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#priorityclass-v1-scheduling-k8s-io"}),"PriorityClass"),": priority of a Pod relative to other Pods, used by the ",Object(i.b)("strong",{parentName:"p"},"scheduler")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"Defines in what order Pods are scheduled (if multiple Pods are pending)"),Object(i.b)("li",{parentName:"ul"},"If there's no place for a pending Pod, the scheduler may evict running Pods with a lower priority")),Object(i.b)("p",null,"PriorityClass and QoS are unrelated: ",Object(i.b)("strong",{parentName:"p"},"QoS used by kubelet")," for killing, ",Object(i.b)("strong",{parentName:"p"},"PriorityClass used by scheduler")," for scheduling (and evicting)"),Object(i.b)("p",null,Object(i.b)("strong",{parentName:"p"},"Best practices:")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"Make tests to discover resource requirements (CPU and memory) for all containers and set requests and limits in Pod specs"),Object(i.b)("li",{parentName:"ul"},"Use ",Object(i.b)("em",{parentName:"li"},"Guaranteed")," QoS class for critical Pods in production (in dev, you can use ",Object(i.b)("em",{parentName:"li"},"Best-effort")," and ",Object(i.b)("em",{parentName:"li"},"Burstable"),")"),Object(i.b)("li",{parentName:"ul"},"If using ",Object(i.b)("em",{parentName:"li"},"Burstable")," QoS, ensure a low limit/request ratio in the LimitRanges",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"The higher the limit/request ratio, the higher the risk that the node runs out of resources and Pods must be killed (if many containers are close to their limits at the same time)"))),Object(i.b)("li",{parentName:"ul"},"Lock down access to PriorityClass with RBAC, limit PriorityClass in ResourceQuota for each namespace")),Object(i.b)("h2",{id:"chapter-3-declarative-deployment"},"Chapter 3: Declarative Deployment"),Object(i.b)("h2",{id:"chapter-4-health-probe"},"Chapter 4: Health Probe"),Object(i.b)("p",null,"A way for the platform (Kubernetes) to learn about the internal state of the containerised apps (which otherwise are black boxes for the platform) and take appropriate action."),Object(i.b)("p",null,"The app should implement these interfaces to provide the maximum amount of information about its internal state."),Object(i.b)("p",null,"Kubernetes (kubelet) performs three types of periodic health checks:"),Object(i.b)("ol",null,Object(i.b)("li",{parentName:"ol"},Object(i.b)("strong",{parentName:"li"},"Process:")," is main process of container still running? \u2192 restart container"),Object(i.b)("li",{parentName:"ol"},Object(i.b)("strong",{parentName:"li"},"Liveness:")," implemented by app, is app running? \u2192 restart container"),Object(i.b)("li",{parentName:"ol"},Object(i.b)("strong",{parentName:"li"},"Readiness:")," implemented by app, is app ready to serve requests? \u2192 remove Pod from Service")),Object(i.b)("p",null,Object(i.b)("strong",{parentName:"p"},"Best practices:")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"Define liveness and readiness probes for all containers"),Object(i.b)("li",{parentName:"ul"},"Set initalDelaySeconds for all liveness and readiness probes (make tests to find out how big the value should be)")),Object(i.b)("h2",{id:"chapter-5-managed-lifecycle"},"Chapter 5: Managed Lifecycle"),Object(i.b)("p",null,"Signals that the platform emits to the application about the lifecycle of the application (which is managed by the platform)."),Object(i.b)("p",null,"The app should listen to these signals to learn about its own lifecycle (which it otherwise has no access to as it is fully managed by the platform) and react to it accordingly."),Object(i.b)("p",null,"Signals:"),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"SIGTERM:")," emitted to a container when the platform decides to shut down a container (for whatever reason). The container should listen to this and gracefully shut down (i.e. cleaning up and exiting the main container process)."),Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"SIGKILL:")," only emitted after SIGTERM when the container process is still running after ",Object(i.b)("inlineCode",{parentName:"li"},"terminationGracePeriodSeconds"),". Kills the container (hard)."),Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"postStart:")," emitted immediately after creating a container. The implementation must be provided by the app (in the same way as liveness/readiness probes, i.e. ",Object(i.b)("em",{parentName:"li"},"exec"),", ",Object(i.b)("em",{parentName:"li"},"httpGet"),", ",Object(i.b)("em",{parentName:"li"},"tcpSocket"),"). Happens independently from starting the container main process, i.e. calling the postStart hook and starting the container process happens at the same time."),Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"preStop:")," emitted when the platform decides to shut down a container. Emitted before SIGTERM and SIGTERM will be only executed after the preStop handler completes. Can be used as an alternative to reacting to SIGTERM.")),Object(i.b)("p",null,Object(i.b)("strong",{parentName:"p"},"Best practices:")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"App should listen to either SIGTERM or implement a preStop hook and gracefully exit",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"Gracefully exiting should be quick, at least within the 30 seconds default ",Object(i.b)("inlineCode",{parentName:"li"},"terminationGracePeriodSeconds")))),Object(i.b)("li",{parentName:"ul"},"If gracefully exiting takes very long, set the ",Object(i.b)("inlineCode",{parentName:"li"},"terminationGracePeriodSeconds")," in the Pod spec"),Object(i.b)("li",{parentName:"ul"},"Implement a postStart hook if there are any start-up tasks to do for a container"),Object(i.b)("li",{parentName:"ul"},"Only use the ",Object(i.b)("em",{parentName:"li"},"exec")," method for the postStart hook, because it is not guaranteed that the container process is already running when the postStart hook is called (so ",Object(i.b)("em",{parentName:"li"},"httpGet")," and ",Object(i.b)("em",{parentName:"li"},"tcpSocket")," could or could not be executed in a race-condition fashion)")),Object(i.b)("h2",{id:"chapter-6-automated-placement"},"Chapter 6: Automated Placement"))}b.isMDXComponent=!0},146:function(e,t,n){"use strict";n.d(t,"a",(function(){return p})),n.d(t,"b",(function(){return m}));var a=n(0),r=n.n(a);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var c=r.a.createContext({}),b=function(e){var t=r.a.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},p=function(e){var t=b(e.components);return r.a.createElement(c.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.a.createElement(r.a.Fragment,{},t)}},d=r.a.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,o=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),p=b(n),d=a,m=p["".concat(o,".").concat(d)]||p[d]||u[d]||i;return n?r.a.createElement(m,l(l({ref:t},c),{},{components:n})):r.a.createElement(m,l({ref:t},c))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=d;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,o[1]=l;for(var c=2;c<i;c++)o[c]=n[c];return r.a.createElement.apply(null,o)}return r.a.createElement.apply(null,n)}d.displayName="MDXCreateElement"}}]);