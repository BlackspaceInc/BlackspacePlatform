(window.webpackJsonp=window.webpackJsonp||[]).push([[100],{168:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return r})),a.d(t,"metadata",(function(){return c})),a.d(t,"rightToc",(function(){return O})),a.d(t,"default",(function(){return i}));var n=a(2),b=a(6),l=(a(0),a(236)),r={title:"FILL keyword",sidebar_label:"FILL",description:"FILL SQL keyword reference documentation."},c={unversionedId:"__uneeded/reference/sql/fill",id:"__uneeded/reference/sql/fill",isDocsHomePage:!1,title:"FILL keyword",description:"FILL SQL keyword reference documentation.",source:"@site/docs/__uneeded/reference/sql/fill.md",slug:"/__uneeded/reference/sql/fill",permalink:"/docs/__uneeded/reference/sql/fill",version:"current",sidebar_label:"FILL"},O=[{value:"Syntax",id:"syntax",children:[{value:"Options",id:"options",children:[]}]},{value:"Examples",id:"examples",children:[]}],j={rightToc:O};function i(e){var t=e.components,r=Object(b.a)(e,["components"]);return Object(l.b)("wrapper",Object(n.a)({},j,r,{components:t,mdxType:"MDXLayout"}),Object(l.b)("p",null,"Specifies fill behavior for missing data for as part of a\n",Object(l.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/sql/sample-by/"}),"SAMPLE BY")," aggregation query."),Object(l.b)("h2",{id:"syntax"},"Syntax"),Object(l.b)("p",null,Object(l.b)("img",{alt:"Flow chart showing the syntax of the FILL keyword",src:a(271).default})),Object(l.b)("h3",{id:"options"},"Options"),Object(l.b)("p",null,"There are as many ",Object(l.b)("inlineCode",{parentName:"p"},"fillOption")," as there are ",Object(l.b)("inlineCode",{parentName:"p"},"aggreate")," columns in your query."),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"fillOption"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Description"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"NONE")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Will not fill. In case there is no data, the time chunk will be skipped in the results. This means your table could potentially be missing intervals.")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"NULL")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Fills with ",Object(l.b)("inlineCode",{parentName:"td"},"null"))),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"PREV")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Fills using the previous value")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"LINEAR")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Fills by linear interpolation of the 2 surrounding points")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"x")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Fills with the constant defined (replace the ",Object(l.b)("inlineCode",{parentName:"td"},"x")," by the value you want. For example ",Object(l.b)("inlineCode",{parentName:"td"},"fill 100.05"))))),Object(l.b)("h2",{id:"examples"},"Examples"),Object(l.b)("p",null,"Consider the following ",Object(l.b)("inlineCode",{parentName:"p"},"prices")," table"),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"timestamp"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"price"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"p1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"p2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts3"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"p3")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"...")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"tsn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"pn")))),Object(l.b)("p",null,"We could run the following to get the minimum, maximum and average price per\nhour using the following query:"),Object(l.b)("pre",null,Object(l.b)("code",Object(n.a)({parentName:"pre"},{className:"language-questdb-sql"}),"SELECT timestamp, min(price) min, max(price) max, avg(price) avg\nFROM PRICES\nSAMPLE BY 1h;\n")),Object(l.b)("p",null,"It would generally return result like this:"),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"timestamp"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"min"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"max"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"average"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"max1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"...")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"tsn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"minn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"maxn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avgn")))),Object(l.b)("p",null,"However, in case there was no ",Object(l.b)("inlineCode",{parentName:"p"},"PRICES")," data for a given hour, your table would\nhave time chunks missing. In the below example, there is no data to generate\naggregates for ",Object(l.b)("inlineCode",{parentName:"p"},"ts3")),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"timestamp"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"min"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"max"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"average"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"max1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"max2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"ts3")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"null")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"null")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"null"))),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts4"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min4"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"max4"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg4")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"...")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"tsn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"minn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"maxn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avgn")))),Object(l.b)("p",null,"Here you can see that the third time chunk is missing. This is because there was\nno price update in the third hour. Let's see what different fill values would\nreturn:"),Object(l.b)("pre",null,Object(l.b)("code",Object(n.a)({parentName:"pre"},{className:"language-questdb-sql"}),"SELECT timestamp, min(price) min, max(price) max, avg(price) avg\nFROM PRICES\nSAMPLE BY 1h\nFILL(null, 0, prev);\n")),Object(l.b)("p",null,"would return the following"),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"timestamp"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"min"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"max"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"average"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"max1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"max2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"ts3")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"NULL")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"0")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"avg2"))),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts4"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min4"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"max4"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg4")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"...")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"tsn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"minn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"maxn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avgn")))),Object(l.b)("p",null,"And the following:"),Object(l.b)("pre",null,Object(l.b)("code",Object(n.a)({parentName:"pre"},{className:"language-questdb-sql"}),"SELECT timestamp, min(price) min, avg(price) avg\nFROM PRICES\nSAMPLE BY 1h\nFILL(25.5, linear);\n")),Object(l.b)("p",null,"Would return:"),Object(l.b)("table",null,Object(l.b)("thead",{parentName:"table"},Object(l.b)("tr",{parentName:"thead"},Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"timestamp"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"min"),Object(l.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"average"))),Object(l.b)("tbody",{parentName:"table"},Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min1"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg1")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min2"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg2")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"ts3")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"25.5")),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(l.b)("inlineCode",{parentName:"td"},"(avg2+avg4)/2"))),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"ts4"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"min4"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avg4")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"..."),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"...")),Object(l.b)("tr",{parentName:"tbody"},Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"tsn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"minn"),Object(l.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"avgn")))))}i.isMDXComponent=!0}}]);