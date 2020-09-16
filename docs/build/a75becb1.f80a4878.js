(window.webpackJsonp=window.webpackJsonp||[]).push([[113],{180:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return r})),a.d(t,"metadata",(function(){return l})),a.d(t,"rightToc",(function(){return s})),a.d(t,"default",(function(){return o}));var i=a(2),n=a(6),c=(a(0),a(244)),r={title:"Security Policy"},l={unversionedId:"__guidelines/security",id:"__guidelines/security",isDocsHomePage:!1,title:"Security Policy",description:"Reporting a vulnerability? See the Vulnerability Reporting section",source:"@site/docs/__guidelines/security.md",slug:"/__guidelines/security",permalink:"/docs/__guidelines/security",version:"current",sidebar:"docs",previous:{title:"Reviewing",permalink:"/docs/__guidelines/reviewing"},next:{title:"Clean Go Code",permalink:"/docs/__guidelines/style_guide/golang/clean-code"}},s=[{value:"Project Structure",id:"project-structure",children:[{value:"Transparency",id:"transparency",children:[]},{value:"Version Control",id:"version-control",children:[]}]},{value:"Personnel",id:"personnel",children:[{value:"Education",id:"education",children:[]},{value:"Policies",id:"policies",children:[]},{value:"Two-factor Authentication",id:"two-factor-authentication",children:[]}]},{value:"Development",id:"development",children:[{value:"Design &amp; Architecture",id:"design--architecture",children:[]},{value:"Dependencies",id:"dependencies",children:[]},{value:"Change Control",id:"change-control",children:[]}]},{value:"Building &amp; Releasing",id:"building--releasing",children:[{value:"Network Security",id:"network-security",children:[]},{value:"Runtime Isolation",id:"runtime-isolation",children:[]},{value:"Asset Audit Logging",id:"asset-audit-logging",children:[]},{value:"Asset Signatures &amp; Checksums",id:"asset-signatures--checksums",children:[]}]},{value:"Vulnerability Reporting",id:"vulnerability-reporting",children:[]}],b={rightToc:s};function o(e){var t=e.components,a=Object(n.a)(e,["components"]);return Object(c.b)("wrapper",Object(i.a)({},b,a,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",{align:"center"},Object(c.b)("strong",null,"Reporting a vulnerability? See the ",Object(c.b)("a",{href:"#vulnerability-reporting"},"Vulnerability Reporting section"))),Object(c.b)("hr",null),Object(c.b)("p",null,"We understand that our users place a high level of trust in Blackspace. The security of Blackspace is a top priority.\nThat's why we apply widely accepted best practices when it comes to security.\nThis document will describe these practices and aims to be as transparent as\npossible on our security efforts."),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#project-structure"}),"Project Structure"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#transparency"}),"Transparency"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#open-source"}),"Open Source")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#workflow"}),"Workflow")))),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#version-control"}),"Version Control"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#git"}),"Git")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#signed-commits"}),"Signed Commits")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#protected-branches"}),"Protected Branches")))))),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#personnel"}),"Personnel"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#education"}),"Education")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#policies"}),"Policies")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#two-factor-authentication"}),"Two-factor Authentication")))),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#development"}),"Development"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#design--architecture"}),"Design & Architecture"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#golang"}),"Golang")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#unsafe-code"}),"Unsafe Code")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#user-privileges"}),"User Privileges")))),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#dependencies"}),"Dependencies")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#change-control"}),"Change Control"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#pull-requests"}),"Pull Requests")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#reviews--approvals"}),"Reviews & Approvals")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#merge-policies"}),"Merge Policies")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#automated-checks"}),"Automated Checks"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#vulnerability-scans"}),"Vulnerability Scans")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#fuzz-testing"}),"Fuzz Testing")))))))),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#building--releasing"}),"Building & Releasing"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#network-security"}),"Network Security")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#runtime-isolation"}),"Runtime Isolation")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#asset-audit-logging"}),"Asset Audit Logging")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#asset-signatures--checksums"}),"Asset Signatures & Checksums")))),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(i.a)({parentName:"li"},{href:"#vulnerability-reporting"}),"Vulnerability Reporting"))),Object(c.b)("h2",{id:"project-structure"},"Project Structure"),Object(c.b)("p",null,"Project structure plays an important role in security. It creates guardrails\nthat prevent common security issues. This section will outline our deliberate\nstructural decisions that impact security."),Object(c.b)("h3",{id:"transparency"},"Transparency"),Object(c.b)("p",null,"We believe transparency is a strong deterrent of nefarious behavior that could\notherwise undermine security."),Object(c.b)("h4",{id:"open-source"},"Open Source"),Object(c.b)("p",null,"Blackspace's dependencies are open-source. All code and changes are publicly\navailable at ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"https://github.com/BlackspaceInc/BlackspacePlatform"}),"our Github repo"),"."),Object(c.b)("h4",{id:"workflow"},"Workflow"),Object(c.b)("p",null,"All of Blackspace's workflow is transparent.\n",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"https://github.com/BlackspaceInc/BlackspacePlatform/pulls"}),"Pull requests"),", ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"https://github.com/BlackspaceInc/BlackspacePlatform/issues"}),"issues"),",\n",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"https://gitter.im/BlackspaceInc/community"}),"chats"),", and ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"https://github.com/BlackspaceInc/BlackspacePlatform/milestones"}),"our roadmap"),"\nare all publicly available."),Object(c.b)("h3",{id:"version-control"},"Version Control"),Object(c.b)("p",null,"Version control ensures that all code changes are audited and authentic."),Object(c.b)("h4",{id:"git"},"Git"),Object(c.b)("p",null,"Blackspace leverages the ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"https://git-scm.com/"}),"Git")," version-control system. This ensures all\nchanges are audited and traceable."),Object(c.b)("h4",{id:"signed-commits"},"Signed Commits"),Object(c.b)("p",null,"Because of Blackspace's ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"CONTRIBUTING.md#merge-style"}),"merge style"),", commits to\nrelease branches are signed by Github itself during the squash and merge\nprocess. Commits to development branches are encouraged to be signed but not\nrequired since changes must go through a ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"#reviews--approvals"}),"review process"),"."),Object(c.b)("h4",{id:"protected-branches"},"Protected Branches"),Object(c.b)("p",null,"Blackspace cuts releases from the ",Object(c.b)("inlineCode",{parentName:"p"},"master")," and ",Object(c.b)("inlineCode",{parentName:"p"},"v*")," branches ",Object(c.b)("em",{parentName:"p"},"only"),". These branches\nare protected. The exact requirements are:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"Cannot be deleted."),Object(c.b)("li",{parentName:"ul"},"Force pushes are not allowed."),Object(c.b)("li",{parentName:"ul"},"A linear history is required."),Object(c.b)("li",{parentName:"ul"},"Signed commits are required."),Object(c.b)("li",{parentName:"ul"},"Administrators are included in these checks.")),Object(c.b)("h2",{id:"personnel"},"Personnel"),Object(c.b)("h3",{id:"education"},"Education"),Object(c.b)("p",null,"Blackspace team members are required to review this security document as well as\nthe ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"CONTRIBUTING.md"}),"contributing")," and ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"REVIEWING.md"}),"reviewing")," documents."),Object(c.b)("h3",{id:"policies"},"Policies"),Object(c.b)("p",null,"Blackspace maintains this security policy. Changes are communicated to all Blackspace\nteam members."),Object(c.b)("h3",{id:"two-factor-authentication"},"Two-factor Authentication"),Object(c.b)("p",null,"All Blackspace team members are required to enable two-factor authentication\nfor their Github accounts."),Object(c.b)("h2",{id:"development"},"Development"),Object(c.b)("h3",{id:"design--architecture"},"Design & Architecture"),Object(c.b)("p",null,"The base of Blackspace's security lies in our choice of underlying technology and\ndecisions around design and architecture."),Object(c.b)("h4",{id:"golang"},"Golang"),Object(c.b)("p",null,"The ","[Golang programming language][urls.go]"," is not memory and thread-safe; it will\nnot catch many common sources of concurency vulnerabilities at compile time. Hence, it is the responsibility of the developer\nto properly account for any potential race conditions."),Object(c.b)("h4",{id:"unsafe-code"},"Unsafe Code"),Object(c.b)("p",null,"Blackspace does not allow the use of unsafe code except in circumstances where it\nis required, such as dealing with CFFI."),Object(c.b)("h4",{id:"user-privileges"},"User Privileges"),Object(c.b)("p",null,"Blackspace is always designed to run under non-",Object(c.b)("inlineCode",{parentName:"p"},"root")," privileges, and our\ndocumentation always defaults to non-",Object(c.b)("inlineCode",{parentName:"p"},"root")," use."),Object(c.b)("h3",{id:"dependencies"},"Dependencies"),Object(c.b)("p",null,"Blackspace aims to reduce the number of dependencies it relies on. If a dependency\nis added it goes through a comprehensive review process that is detailed in\nthe ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"REVIEWING.md#dependencies"}),"Reviewing guide"),"."),Object(c.b)("h3",{id:"change-control"},"Change Control"),Object(c.b)("p",null,"As noted above Blackspace uses the Git version control system on Github."),Object(c.b)("h4",{id:"pull-requests"},"Pull Requests"),Object(c.b)("p",null,"All changes to Blackspace must go through a pull request review process."),Object(c.b)("h4",{id:"reviews--approvals"},"Reviews & Approvals"),Object(c.b)("p",null,"All pull requests must be reviewed by at least one Blackspace team member. The\nreview process takes into account many factors, all of which are detailed in\nour ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"REVIEWING.md"}),"Reviewing guide"),". In exceptional circumstances, this\napproval can be retroactive."),Object(c.b)("h4",{id:"merge-policies"},"Merge Policies"),Object(c.b)("p",null,"Blackspace requires pull requests to pass all ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"#automated-checks"}),"automated checks"),".\nOnce passed, the pull request must be squashed and merged. This creates a clean\nlinear history with a Blackspace team member's co-sign."),Object(c.b)("h4",{id:"automated-checks"},"Automated Checks"),Object(c.b)("p",null,"When possible, we'll create automated checks to enforce security policies."),Object(c.b)("h5",{id:"vulnerability-scans"},"Vulnerability Scans"),Object(c.b)("p",null,"Blackspace will implement an automated ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"https://github.com/securego/gosec"}),Object(c.b)("inlineCode",{parentName:"a"},"cargo deny")," check"),". This\nis a security checker for golang source code."),Object(c.b)("h5",{id:"fuzz-testing"},"Fuzz Testing"),Object(c.b)("p",null,"Blackspace will implement automated fuzz testing to probe our code for other sources\nof potential vulnerabilities."),Object(c.b)("h2",{id:"building--releasing"},"Building & Releasing"),Object(c.b)("p",null,"Blackspace takes care to secure the build and release process to prevent unintended\nmodifications."),Object(c.b)("h3",{id:"network-security"},"Network Security"),Object(c.b)("p",null,"All network traffic is secured via TLS and SSH. This includes checking out\nBlackspace's code from the relevant ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"#protected-branches"}),"protected branch"),",\nDocker image retrieval, and publishment of Blackspace's release artifacts."),Object(c.b)("h3",{id:"runtime-isolation"},"Runtime Isolation"),Object(c.b)("p",null,"All builds run in an isolated sandbox that is destroyed after each use."),Object(c.b)("h3",{id:"asset-audit-logging"},"Asset Audit Logging"),Object(c.b)("p",null,"Changes to Blackspace's assets will be logged through Digital Ocean's blob storage audit logging feature."),Object(c.b)("h3",{id:"asset-signatures--checksums"},"Asset Signatures & Checksums"),Object(c.b)("p",null,"All assets are signed with checksums allowing users to verify asset authenticity\nupon download. This verifies that assets have not been modified at rest."),Object(c.b)("h2",{id:"vulnerability-reporting"},"Vulnerability Reporting"),Object(c.b)("p",null,"We deeply appreciate any effort to discover and disclose security\nvulnerabilities responsibly."),Object(c.b)("p",null,"If you would like to report a vulnerability or have any security concerns with\nBlackspace, please e-mail ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"mailto:yoanyombapro@gmail.com"}),"yoanyombapro@gmail.com"),"."),Object(c.b)("p",null,"For non-critical matters, we prefer users ",Object(c.b)("a",Object(i.a)({parentName:"p"},{href:"https://github.com/BlackspaceInc/BlackspacePlatform/issues/new"}),"open an issue"),".\nFor us to best investigate your request, please include any of the\nfollowing when reporting:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"Proof of concept"),Object(c.b)("li",{parentName:"ul"},"Any tools, including versions used"),Object(c.b)("li",{parentName:"ul"},"Any relevant output")),Object(c.b)("p",null,"We take all disclosures very seriously and will do our best to rapidly respond\nand verify the vulnerability before taking the necessary steps to fix it. After\nour initial reply to your disclosure, which should be directly after receiving\nit, we will periodically update you with the status of the fix."))}o.isMDXComponent=!0}}]);