(window.webpackJsonp=window.webpackJsonp||[]).push([[55],{126:function(I,g,e){"use strict";e.r(g),e.d(g,"frontMatter",(function(){return d})),e.d(g,"metadata",(function(){return a})),e.d(g,"rightToc",(function(){return t})),e.d(g,"default",(function(){return b}));var A=e(2),i=e(6),C=(e(0),e(243)),d={title:"ALTER TABLE COLUMN ADD INDEX keywords",sidebar_label:"ALTER TABLE COLUMN ADD INDEX",description:"ADD INDEX SQL keword reference documentation."},a={unversionedId:"__uneeded/reference/sql/alter-table-alter-column-add-index",id:"__uneeded/reference/sql/alter-table-alter-column-add-index",isDocsHomePage:!1,title:"ALTER TABLE COLUMN ADD INDEX keywords",description:"ADD INDEX SQL keword reference documentation.",source:"@site/docs/__uneeded/reference/sql/alter-table-alter-column-add-index.md",slug:"/__uneeded/reference/sql/alter-table-alter-column-add-index",permalink:"/docs/__uneeded/reference/sql/alter-table-alter-column-add-index",version:"current",sidebar_label:"ALTER TABLE COLUMN ADD INDEX"},t=[{value:"Syntax",id:"syntax",children:[]},{value:"Description",id:"description",children:[]},{value:"Example",id:"example",children:[]}],n={rightToc:t};function b(I){var g=I.components,d=Object(i.a)(I,["components"]);return Object(C.b)("wrapper",Object(A.a)({},n,d,{components:g,mdxType:"MDXLayout"}),Object(C.b)("p",null,"Adds an index to an existing column."),Object(C.b)("h2",{id:"syntax"},"Syntax"),Object(C.b)("p",null,Object(C.b)("img",{alt:"Flow chart showing the syntax of the ALTER TABLE keyword",src:e(254).default}),"\n",Object(C.b)("img",{alt:"Flow chart showing the syntax of the ALTER TABLE with ADD INDEX keyword",src:e(320).default})),Object(C.b)("h2",{id:"description"},"Description"),Object(C.b)("p",null,"Adds new index to column of type ",Object(C.b)("inlineCode",{parentName:"p"},"symbol"),". Adding index is an atomic,\nnon-blocking and non-waiting operation. Once complete optimiser will start using\nnew index for SQL executions."),Object(C.b)("h2",{id:"example"},"Example"),Object(C.b)("pre",null,Object(C.b)("code",Object(A.a)({parentName:"pre"},{className:"language-questdb-sql",metastring:'title="Adding an index"',title:'"Adding',an:!0,'index"':!0}),"ALTER TABLE trades ALTER COLUMN instrument ADD INDEX;\n")),Object(C.b)("div",{className:"admonition admonition-info alert alert--info"},Object(C.b)("div",Object(A.a)({parentName:"div"},{className:"admonition-heading"}),Object(C.b)("h5",{parentName:"div"},Object(C.b)("span",Object(A.a)({parentName:"h5"},{className:"admonition-icon"}),Object(C.b)("svg",Object(A.a)({parentName:"span"},{xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"}),Object(C.b)("path",Object(A.a)({parentName:"svg"},{fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"})))),"info")),Object(C.b)("div",Object(A.a)({parentName:"div"},{className:"admonition-content"}),Object(C.b)("p",{parentName:"div"},"For more information about indexes please refer to the\n",Object(C.b)("a",Object(A.a)({parentName:"p"},{href:"/docs/concept/indexes/"}),"INDEX section"),"."))))}b.isMDXComponent=!0},320:function(I,g,e){"use strict";e.r(g),g.default="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI0NTkiIGhlaWdodD0iMzciPgogICAgPGRlZnM+CiAgICAgICAgPHN0eWxlIHR5cGU9InRleHQvY3NzIj4KICAgICAgICAgICAgQG5hbWVzcGFjZSAiaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciOwogICAgICAgICAgICAgICAgLmxpbmUgICAgICAgICAgICAgICAgIHtmaWxsOiBub25lOyBzdHJva2U6ICM2MzYyNzM7fQogICAgICAgICAgICAgICAgLmJvbGQtbGluZSAgICAgICAgICAgIHtzdHJva2U6ICM2MzYyNzM7IHNoYXBlLXJlbmRlcmluZzogY3Jpc3BFZGdlczsgc3Ryb2tlLXdpZHRoOiAyOyB9CiAgICAgICAgICAgICAgICAudGhpbi1saW5lICAgICAgICAgICAge3N0cm9rZTogIzYzNjI3Mzsgc2hhcGUtcmVuZGVyaW5nOiBjcmlzcEVkZ2VzfQogICAgICAgICAgICAgICAgLmZpbGxlZCAgICAgICAgICAgICAgIHtmaWxsOiAjNjM2MjczOyBzdHJva2U6IG5vbmU7fQogICAgICAgICAgICAgICAgdGV4dC50ZXJtaW5hbCAgICAgICAgIHtmb250LWZhbWlseTogLWFwcGxlLXN5c3RlbSwgQmxpbmtNYWNTeXN0ZW1Gb250LCAiU2Vnb2UgVUkiLCBSb2JvdG8sIFVidW50dSwgQ2FudGFyZWxsLCBIZWx2ZXRpY2EsIHNhbnMtc2VyaWY7CiAgICAgICAgICAgICAgICBmb250LXNpemU6IDEycHg7CiAgICAgICAgICAgICAgICBmaWxsOiAjZmZmZmZmOwogICAgICAgICAgICAgICAgZm9udC13ZWlnaHQ6IGJvbGQ7CiAgICAgICAgICAgICAgICB9CiAgICAgICAgICAgICAgICB0ZXh0Lm5vbnRlcm1pbmFsICAgICAge2ZvbnQtZmFtaWx5OiAtYXBwbGUtc3lzdGVtLCBCbGlua01hY1N5c3RlbUZvbnQsICJTZWdvZSBVSSIsIFJvYm90bywgVWJ1bnR1LCBDYW50YXJlbGwsIEhlbHZldGljYSwgc2Fucy1zZXJpZjsKICAgICAgICAgICAgICAgIGZvbnQtc2l6ZTogMTJweDsKICAgICAgICAgICAgICAgIGZpbGw6ICNlMjg5YTQ7CiAgICAgICAgICAgICAgICBmb250LXdlaWdodDogbm9ybWFsOwogICAgICAgICAgICAgICAgfQogICAgICAgICAgICAgICAgdGV4dC5yZWdleHAgICAgICAgICAgIHtmb250LWZhbWlseTogLWFwcGxlLXN5c3RlbSwgQmxpbmtNYWNTeXN0ZW1Gb250LCAiU2Vnb2UgVUkiLCBSb2JvdG8sIFVidW50dSwgQ2FudGFyZWxsLCBIZWx2ZXRpY2EsIHNhbnMtc2VyaWY7CiAgICAgICAgICAgICAgICBmb250LXNpemU6IDEycHg7CiAgICAgICAgICAgICAgICBmaWxsOiAjMDAxNDFGOwogICAgICAgICAgICAgICAgZm9udC13ZWlnaHQ6IG5vcm1hbDsKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgICAgIHJlY3QsIGNpcmNsZSwgcG9seWdvbiB7ZmlsbDogbm9uZTsgc3Ryb2tlOiBub25lO30KICAgICAgICAgICAgICAgIHJlY3QudGVybWluYWwgICAgICAgICB7ZmlsbDogbm9uZTsgc3Ryb2tlOiAjYmUyZjViO30KICAgICAgICAgICAgICAgIHJlY3Qubm9udGVybWluYWwgICAgICB7ZmlsbDogcmdiYSgyNTUsMjU1LDI1NSwwLjEpOyBzdHJva2U6IG5vbmU7fQogICAgICAgICAgICAgICAgcmVjdC50ZXh0ICAgICAgICAgICAgIHtmaWxsOiBub25lOyBzdHJva2U6IG5vbmU7fQogICAgICAgICAgICAgICAgcG9seWdvbi5yZWdleHAgICAgICAgIHtmaWxsOiAjQzdFQ0ZGOyBzdHJva2U6ICMwMzhjYmM7fQogICAgICAgIDwvc3R5bGU+CiAgICA8L2RlZnM+CiAgICA8cG9seWdvbiBwb2ludHM9IjkgMTcgMSAxMyAxIDIxIi8+CiAgICA8cG9seWdvbiBwb2ludHM9IjE3IDE3IDkgMTMgOSAyMSIvPjxhIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4bGluazpocmVmPSIjYWx0ZXIiIHhsaW5rOnRpdGxlPSJhbHRlciI+CiAgICA8cmVjdCB4PSIzMSIgeT0iMyIgd2lkdGg9IjQ4IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSIyOSIgeT0iMSIgd2lkdGg9IjQ4IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iMzkiIHk9IjIxIj5hbHRlcjwvdGV4dD48L2E+PGEgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHhsaW5rOmhyZWY9IiNjb2x1bW4iIHhsaW5rOnRpdGxlPSJjb2x1bW4iPgogICAgPHJlY3QgeD0iOTkiIHk9IjMiIHdpZHRoPSI2NCIgaGVpZ2h0PSIzMiIvPgogICAgPHJlY3QgeD0iOTciIHk9IjEiIHdpZHRoPSI2NCIgaGVpZ2h0PSIzMiIgY2xhc3M9Im5vbnRlcm1pbmFsIi8+CiAgICA8dGV4dCBjbGFzcz0ibm9udGVybWluYWwiIHg9IjEwNyIgeT0iMjEiPmNvbHVtbjwvdGV4dD48L2E+PHJlY3QgeD0iMTgzIiB5PSIzIiB3aWR0aD0iMTEwIiBoZWlnaHQ9IjMyIiByeD0iMTAiLz4KICAgIDxyZWN0IHg9IjE4MSIgeT0iMSIgd2lkdGg9IjExMCIgaGVpZ2h0PSIzMiIgY2xhc3M9InRlcm1pbmFsIiByeD0iMTAiLz4KICAgIDx0ZXh0IGNsYXNzPSJ0ZXJtaW5hbCIgeD0iMTkxIiB5PSIyMSI+Y29sdW1uIG5hbWU8L3RleHQ+PGEgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHhsaW5rOmhyZWY9IiNhZGQiIHhsaW5rOnRpdGxlPSJhZGQiPgogICAgPHJlY3QgeD0iMzEzIiB5PSIzIiB3aWR0aD0iNDQiIGhlaWdodD0iMzIiLz4KICAgIDxyZWN0IHg9IjMxMSIgeT0iMSIgd2lkdGg9IjQ0IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iMzIxIiB5PSIyMSI+YWRkPC90ZXh0PjwvYT48YSB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeGxpbms6aHJlZj0iI2luZGV4IiB4bGluazp0aXRsZT0iaW5kZXgiPgogICAgPHJlY3QgeD0iMzc3IiB5PSIzIiB3aWR0aD0iNTQiIGhlaWdodD0iMzIiLz4KICAgIDxyZWN0IHg9IjM3NSIgeT0iMSIgd2lkdGg9IjU0IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iMzg1IiB5PSIyMSI+aW5kZXg8L3RleHQ+PC9hPjxwYXRoIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGNsYXNzPSJsaW5lIiBkPSJtMTcgMTcgaDIgbTAgMCBoMTAgbTQ4IDAgaDEwIG0wIDAgaDEwIG02NCAwIGgxMCBtMCAwIGgxMCBtMTEwIDAgaDEwIG0wIDAgaDEwIG00NCAwIGgxMCBtMCAwIGgxMCBtNTQgMCBoMTAgbTMgMCBoLTMiLz4KICAgIDxwb2x5Z29uIHBvaW50cz0iNDQ5IDE3IDQ1NyAxMyA0NTcgMjEiLz4KICAgIDxwb2x5Z29uIHBvaW50cz0iNDQ5IDE3IDQ0MSAxMyA0NDEgMjEiLz48L3N2Zz4K"}}]);