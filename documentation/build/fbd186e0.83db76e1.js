(window.webpackJsonp=window.webpackJsonp||[]).push([[158],{227:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return s})),a.d(t,"metadata",(function(){return o})),a.d(t,"rightToc",(function(){return i})),a.d(t,"default",(function(){return l}));var n=a(2),c=a(6),r=(a(0),a(236)),s={title:"How to use Blackspace with Docker",sidebar_label:"Docker",description:"Tutorial showing how to use Blackspace with Docker."},o={unversionedId:"guide/docker",id:"guide/docker",isDocsHomePage:!1,title:"How to use Blackspace with Docker",description:"Tutorial showing how to use Blackspace with Docker.",source:"@site/docs/guide/docker.md",slug:"/guide/docker",permalink:"/docs/guide/docker",version:"current",sidebar_label:"Docker",sidebar:"docs",previous:{title:"Third Party Api's",permalink:"/docs/technology/third-party-apis/third-party-apis"},next:{title:"How to setup blackspace with Kubernetes",permalink:"/docs/guide/kubernetes"}},i=[{value:"Install Docker",id:"install-docker",children:[]},{value:"Blackspace Backend Service Images",id:"blackspace-backend-service-images",children:[]},{value:"Container status",id:"container-status",children:[]},{value:"Importing data and sending queries",id:"importing-data-and-sending-queries",children:[]},{value:"Data persistence",id:"data-persistence",children:[{value:"Restart an existing container",id:"restart-an-existing-container",children:[]},{value:"Re-run <code>docker run</code>",id:"re-run-docker-run",children:[]}]}],d={rightToc:i};function l(e){var t=e.components,a=Object(c.a)(e,["components"]);return Object(r.b)("wrapper",Object(n.a)({},d,a,{components:t,mdxType:"MDXLayout"}),Object(r.b)("p",null,"Docker is great to get started in minutes with just a few commands. Follow this\nguide to set up and start Blackspace locally. By the end, you will be able to send and\nquery data to backend services using the REST API."),Object(r.b)("h2",{id:"install-docker"},"Install Docker"),Object(r.b)("p",null,"Before we start, you will need to install Docker. You can find guides for your\nplatform ",Object(r.b)("a",Object(n.a)({parentName:"p"},{href:"https://docs.docker.com/get-docker/"}),"on the official documentation"),"."),Object(r.b)("h2",{id:"blackspace-backend-service-images"},"Blackspace Backend Service Images"),Object(r.b)("p",null,"With Docker installed, you will need to pull All Blackspace Backend Service Images and create a\ncontainer. You can do both in one command using ",Object(r.b)("inlineCode",{parentName:"p"},"docker run"),":"),Object(r.b)("pre",null,Object(r.b)("code",Object(n.a)({parentName:"pre"},{className:"language-shell"}),"git clone BlackspaceInc/BlackspacePlatform\nmake clean\nmake up\n")),Object(r.b)("h2",{id:"container-status"},"Container status"),Object(r.b)("p",null,"You can check the status of your container with ",Object(r.b)("strong",{parentName:"p"},"docker ps"),". It also lists the\nports we published:"),Object(r.b)("pre",null,Object(r.b)("code",Object(n.a)({parentName:"pre"},{className:"language-shell"}),"docker ps\n")),Object(r.b)("pre",null,Object(r.b)("code",Object(n.a)({parentName:"pre"},{className:"language-shell",metastring:'title="Result"',title:'"Result"'}),'CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                NAMES\ndd363939f261        blackspaceInc/frontend-service     "/app/bin/go -m io\u2026"   3 seconds ago       Up 2 seconds        8812/tcp, 9000/tcp   frosty_gauss\ndd363939f262        blackspaceInc/company-service      "/app/bin/go -m io\u2026"   3 seconds ago       Up 2 seconds        8813/tcp, 9001/tcp   frosty_gauss\ndd363939f263        blackspaceInc/user-service         "/app/bin/go -m io\u2026"   3 seconds ago       Up 2 seconds        8813/tcp, 9002/tcp   frosty_gauss\ndd363939f264        blackspaceInc/auth-service         "/app/bin/go -m io\u2026"   3 seconds ago       Up 2 seconds        8813/tcp, 9003/tcp   frosty_gauss\n\n')),Object(r.b)("h2",{id:"importing-data-and-sending-queries"},"Importing data and sending queries"),Object(r.b)("p",null,"\ud83c\udf89 Congratulations, you have a running QuestDB server. You can now start to\ninteract with it:"),Object(r.b)("h2",{id:"data-persistence"},"Data persistence"),Object(r.b)("h3",{id:"restart-an-existing-container"},"Restart an existing container"),Object(r.b)("p",null,"When you stop the container, it will not be removed by Docker. This means that\nyou can restart it anytime and your data will be accessible:"),Object(r.b)("pre",null,Object(r.b)("code",Object(n.a)({parentName:"pre"},{className:"language-shell",metastring:"title=\"Start container from the  ID obtained with 'docker ps'\"",title:'"Start',container:!0,from:!0,the:!0,"":!0,ID:!0,obtained:!0,with:!0,"'docker":!0,"ps'\"":!0}),"docker start dd363939f261\n")),Object(r.b)("h3",{id:"re-run-docker-run"},"Re-run ",Object(r.b)("inlineCode",{parentName:"h3"},"docker run")),Object(r.b)("p",null,"If you re-run the command:"),Object(r.b)("pre",null,Object(r.b)("code",Object(n.a)({parentName:"pre"},{className:"language-shell"}),"docker run -p 9000:9000 -p 8812:8812 blackspaceInc/frontend-service\n")),Object(r.b)("p",null,"A new container will be created for the microservice image. This means that the\ncontainer will be fresh, any data you may have created previously won't be\naccessible."))}l.isMDXComponent=!0}}]);