(window.webpackJsonp=window.webpackJsonp||[]).push([[78],{151:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return r})),a.d(t,"metadata",(function(){return l})),a.d(t,"rightToc",(function(){return s})),a.d(t,"default",(function(){return c}));var n=a(2),i=(a(0),a(243));const r={title:"Patterns"},l={unversionedId:"checklists/backend/backend-kubernetes/patterns",id:"checklists/backend/backend-kubernetes/patterns",isDocsHomePage:!1,title:"Patterns",description:"Predictable Demands",source:"@site/docs/checklists/backend/backend-kubernetes/patterns.md",slug:"/checklists/backend/backend-kubernetes/patterns",permalink:"/docs/checklists/backend/backend-kubernetes/patterns",version:"current",sidebar:"docs",previous:{title:"Namespace Management",permalink:"/docs/checklists/backend/backend-kubernetes/governance"},next:{title:"Production Ready Checklist",permalink:"/docs/checklists/backend/backend-production-ready/production-ready-checklist"}},s=[{value:"Predictable Demands",id:"predictable-demands",children:[]},{value:"Declarative Deployment",id:"declarative-deployment",children:[]},{value:"Health Probe",id:"health-probe",children:[]},{value:"Managed Lifecycle",id:"managed-lifecycle",children:[]},{value:"Automated Placement",id:"automated-placement",children:[]}],o={rightToc:s};function c({components:e,...t}){return Object(i.b)("wrapper",Object(n.a)({},o,t,{components:e,mdxType:"MDXLayout"}),Object(i.b)("h2",{id:"predictable-demands"},"Predictable Demands"),Object(i.b)("p",null,"Containers should know and declare their requirements. This allows the platform (Kubernetes) to put Pods to the right place in the cluster and treat them appropriately."),Object(i.b)("p",null,Object(i.b)("strong",{parentName:"p"},"Types of requirements:")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"Runtime dependencies",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"Persistent storage, host port, configuration (ConfigMaps, Secrets): defined in Pod spec."))),Object(i.b)("li",{parentName:"ul"},"Resource requirements",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"Requests"),Object(i.b)("li",{parentName:"ul"},"Limits"),Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"QoS")," of Pods (defined by requests and limits): defines in what order Pods are killed by the ",Object(i.b)("strong",{parentName:"li"},"kubelet"),Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"Best-effort: no requests and limits, lowest priority, Pod is killed first"),Object(i.b)("li",{parentName:"ul"},"Burstable: limits are higher than requests, middle-priority, killed if no Best-effort Pods exist"),Object(i.b)("li",{parentName:"ul"},"Guaranteed: limits equal requests, highest priority, killed only if no Best-effort and Burstable Pods exist")))))),Object(i.b)("p",null,Object(i.b)("a",Object(n.a)({parentName:"p"},{href:"https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/"}),Object(i.b)("strong",{parentName:"a"},"Pod Priority and Preemption")),":"),Object(i.b)("p",null,Object(i.b)("a",Object(n.a)({parentName:"p"},{href:"https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#priorityclass-v1-scheduling-k8s-io"}),"PriorityClass"),": priority of a Pod relative to other Pods, used by the ",Object(i.b)("strong",{parentName:"p"},"scheduler")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"Defines in what order Pods are scheduled (if multiple Pods are pending)"),Object(i.b)("li",{parentName:"ul"},"If there's no place for a pending Pod, the scheduler may evict running Pods with a lower priority")),Object(i.b)("p",null,"PriorityClass and QoS are unrelated: ",Object(i.b)("strong",{parentName:"p"},"QoS used by kubelet")," for killing, ",Object(i.b)("strong",{parentName:"p"},"PriorityClass used by scheduler")," for scheduling (and evicting)"),Object(i.b)("p",null,Object(i.b)("strong",{parentName:"p"},"Best practices:")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"Make tests to discover resource requirements (CPU and memory) for all containers and set requests and limits in Pod specs"),Object(i.b)("li",{parentName:"ul"},"Use ",Object(i.b)("em",{parentName:"li"},"Guaranteed")," QoS class for critical Pods in production (in dev, you can use ",Object(i.b)("em",{parentName:"li"},"Best-effort")," and ",Object(i.b)("em",{parentName:"li"},"Burstable"),")"),Object(i.b)("li",{parentName:"ul"},"If using ",Object(i.b)("em",{parentName:"li"},"Burstable")," QoS, ensure a low limit/request ratio in the LimitRanges",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"The higher the limit/request ratio, the higher the risk that the node runs out of resources and Pods must be killed (if many containers are close to their limits at the same time)"))),Object(i.b)("li",{parentName:"ul"},"Lock down access to PriorityClass with RBAC, limit PriorityClass in ResourceQuota for each namespace")),Object(i.b)("h2",{id:"declarative-deployment"},"Declarative Deployment"),Object(i.b)("h2",{id:"health-probe"},"Health Probe"),Object(i.b)("p",null,"A way for the platform (Kubernetes) to learn about the internal state of the containerised apps (which otherwise are black boxes for the platform) and take appropriate action."),Object(i.b)("p",null,"The app should implement these interfaces to provide the maximum amount of information about its internal state."),Object(i.b)("p",null,"Kubernetes (kubelet) performs three types of periodic health checks:"),Object(i.b)("ol",null,Object(i.b)("li",{parentName:"ol"},Object(i.b)("strong",{parentName:"li"},"Process:")," is main process of container still running? \u2192 restart container"),Object(i.b)("li",{parentName:"ol"},Object(i.b)("strong",{parentName:"li"},"Liveness:")," implemented by app, is app running? \u2192 restart container"),Object(i.b)("li",{parentName:"ol"},Object(i.b)("strong",{parentName:"li"},"Readiness:")," implemented by app, is app ready to serve requests? \u2192 remove Pod from Service")),Object(i.b)("p",null,Object(i.b)("strong",{parentName:"p"},"Best practices:")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"Define liveness and readiness probes for all containers"),Object(i.b)("li",{parentName:"ul"},"Set initalDelaySeconds for all liveness and readiness probes (make tests to find out how big the value should be)")),Object(i.b)("h2",{id:"managed-lifecycle"},"Managed Lifecycle"),Object(i.b)("p",null,"Signals that the platform emits to the application about the lifecycle of the application (which is managed by the platform)."),Object(i.b)("p",null,"The app should listen to these signals to learn about its own lifecycle (which it otherwise has no access to as it is fully managed by the platform) and react to it accordingly."),Object(i.b)("p",null,"Signals:"),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"SIGTERM:")," emitted to a container when the platform decides to shut down a container (for whatever reason). The container should listen to this and gracefully shut down (i.e. cleaning up and exiting the main container process)."),Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"SIGKILL:")," only emitted after SIGTERM when the container process is still running after ",Object(i.b)("inlineCode",{parentName:"li"},"terminationGracePeriodSeconds"),". Kills the container (hard)."),Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"postStart:")," emitted immediately after creating a container. The implementation must be provided by the app (in the same way as liveness/readiness probes, i.e. ",Object(i.b)("em",{parentName:"li"},"exec"),", ",Object(i.b)("em",{parentName:"li"},"httpGet"),", ",Object(i.b)("em",{parentName:"li"},"tcpSocket"),"). Happens independently from starting the container main process, i.e. calling the postStart hook and starting the container process happens at the same time."),Object(i.b)("li",{parentName:"ul"},Object(i.b)("strong",{parentName:"li"},"preStop:")," emitted when the platform decides to shut down a container. Emitted before SIGTERM and SIGTERM will be only executed after the preStop handler completes. Can be used as an alternative to reacting to SIGTERM.")),Object(i.b)("p",null,Object(i.b)("strong",{parentName:"p"},"Best practices:")),Object(i.b)("ul",null,Object(i.b)("li",{parentName:"ul"},"App should listen to either SIGTERM or implement a preStop hook and gracefully exit",Object(i.b)("ul",{parentName:"li"},Object(i.b)("li",{parentName:"ul"},"Gracefully exiting should be quick, at least within the 30 seconds default ",Object(i.b)("inlineCode",{parentName:"li"},"terminationGracePeriodSeconds")))),Object(i.b)("li",{parentName:"ul"},"If gracefully exiting takes very long, set the ",Object(i.b)("inlineCode",{parentName:"li"},"terminationGracePeriodSeconds")," in the Pod spec"),Object(i.b)("li",{parentName:"ul"},"Implement a postStart hook if there are any start-up tasks to do for a container"),Object(i.b)("li",{parentName:"ul"},"Only use the ",Object(i.b)("em",{parentName:"li"},"exec")," method for the postStart hook, because it is not guaranteed that the container process is already running when the postStart hook is called (so ",Object(i.b)("em",{parentName:"li"},"httpGet")," and ",Object(i.b)("em",{parentName:"li"},"tcpSocket")," could or could not be executed in a race-condition fashion)")),Object(i.b)("h2",{id:"automated-placement"},"Automated Placement"))}c.isMDXComponent=!0}}]);