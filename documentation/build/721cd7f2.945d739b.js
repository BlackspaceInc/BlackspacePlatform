(window.webpackJsonp=window.webpackJsonp||[]).push([[76],{144:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return r})),a.d(t,"metadata",(function(){return c})),a.d(t,"rightToc",(function(){return O})),a.d(t,"default",(function(){return i}));var n=a(2),b=a(6),l=(a(0),a(236)),r={title:"UNION keyword",sidebar_label:"UNION",description:"UNION SQL keyword reference documentation."},c={unversionedId:"__uneeded/reference/sql/union",id:"__uneeded/reference/sql/union",isDocsHomePage:!1,title:"UNION keyword",description:"UNION SQL keyword reference documentation.",source:"@site/docs/__uneeded/reference/sql/union.md",slug:"/__uneeded/reference/sql/union",permalink:"/docs/__uneeded/reference/sql/union",version:"current",sidebar_label:"UNION"},O=[{value:"Overview",id:"overview",children:[]},{value:"Syntax",id:"syntax",children:[]},{value:"Examples",id:"examples",children:[]}],j={rightToc:O};function i(e){var t=e.components,r=Object(b.a)(e,["components"]);return Object(l.b)("wrapper",Object(n.a)({},j,r,{components:t,mdxType:"MDXLayout"}),Object(l.b)("h2",{id:"overview"},"Overview"),Object(l.b)("p",null,Object(l.b)("inlineCode",{parentName:"p"},"UNION")," is used to combine the results of two or more ",Object(l.b)("inlineCode",{parentName:"p"},"SELECT")," statements. To\nwork properly:"),Object(l.b)("ul",null,Object(l.b)("li",{parentName:"ul"},"Each select statement should return the same number of column"),Object(l.b)("li",{parentName:"ul"},"Each column should have the same type"),Object(l.b)("li",{parentName:"ul"},"Columns should be in the same order")),Object(l.b)("h2",{id:"syntax"},"Syntax"),Object(l.b)("p",null,Object(l.b)("img",{alt:"Flow chart showing the syntax of the UNION keyword",src:a(274).default})),Object(l.b)("ul",null,Object(l.b)("li",{parentName:"ul"},Object(l.b)("inlineCode",{parentName:"li"},"UNION")," will return distinct results."),Object(l.b)("li",{parentName:"ul"},Object(l.b)("inlineCode",{parentName:"li"},"UNION ALL")," will return all results including duplicates.")),Object(l.b)("h2",{id:"examples"},"Examples"),Object(l.b)("p",null,"Let's assume the following two tables listA"),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Description"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"ID"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Red Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Blue Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Green Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"3")))),Object(l.b)("p",null,"listB"),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Description"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"ID"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Pink Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Black Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Green Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"3")))),Object(l.b)("pre",null,Object(l.b)("code",Object(n.a)({parentName:"pre"},{className:"language-questdb-sql"}),"liastA UNION listB\n")),Object(l.b)("p",null,"will return"),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Description"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"ID"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Red Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Blue Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Green Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"3")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Pink Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Black Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"2")))),Object(l.b)("pre",null,Object(l.b)("code",Object(n.a)({parentName:"pre"},{className:"language-questdb-sql"}),"liastA UNION ALL listB\n")),Object(l.b)("p",null,"will return"),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Description"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"ID"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Red Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Blue Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Green Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"3")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Pink Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Black Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Green Pen"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"3")))))}i.isMDXComponent=!0}}]);