(window.webpackJsonp=window.webpackJsonp||[]).push([[87],{160:function(I,g,e){"use strict";e.r(g),e.d(g,"frontMatter",(function(){return A})),e.d(g,"metadata",(function(){return a})),e.d(g,"rightToc",(function(){return t})),e.d(g,"default",(function(){return l}));var i=e(2),C=(e(0),e(243));const A={title:"WITH keyword",sidebar_label:"WITH",description:"WITH SQL keyword reference documentation."},a={unversionedId:"__uneeded/reference/sql/with",id:"__uneeded/reference/sql/with",isDocsHomePage:!1,title:"WITH keyword",description:"WITH SQL keyword reference documentation.",source:"@site/docs/__uneeded/reference/sql/with.md",slug:"/__uneeded/reference/sql/with",permalink:"/docs/__uneeded/reference/sql/with",version:"current",sidebar_label:"WITH"},t=[{value:"Syntax",id:"syntax",children:[]},{value:"Examples",id:"examples",children:[]}],b={rightToc:t};function l({components:I,...g}){return Object(C.b)("wrapper",Object(i.a)({},b,g,{components:I,mdxType:"MDXLayout"}),Object(C.b)("p",null,"Name one or several sub-queries to be used within the main query."),Object(C.b)("p",null,"This clause makes it easy to simplify large or complex statements which involve\nsub-queries, particularly when such sub-queries are used several times."),Object(C.b)("h2",{id:"syntax"},"Syntax"),Object(C.b)("p",null,Object(C.b)("img",{alt:"Flow chart showing the syntax of the WITH clause",src:e(358).default})),Object(C.b)("p",null,"Where:"),Object(C.b)("ul",null,Object(C.b)("li",{parentName:"ul"},Object(C.b)("inlineCode",{parentName:"li"},"subQueryName")," is the alias for the sub-query"),Object(C.b)("li",{parentName:"ul"},Object(C.b)("inlineCode",{parentName:"li"},"subQuery")," is a SQL query (e.g ",Object(C.b)("inlineCode",{parentName:"li"},"SELECT * FROM table"),")"),Object(C.b)("li",{parentName:"ul"},Object(C.b)("inlineCode",{parentName:"li"},"mainQuery")," is the main SQL query which involves the ",Object(C.b)("inlineCode",{parentName:"li"},"subQuery")," using its\nalias.")),Object(C.b)("h2",{id:"examples"},"Examples"),Object(C.b)("pre",null,Object(C.b)("code",Object(i.a)({parentName:"pre"},{className:"language-questdb-sql",metastring:'title="Single alias"',title:'"Single','alias"':!0}),"WITH first_10_users AS (SELECT * FROM users limit 10)\nSELECT user_name FROM first_10_users;\n")),Object(C.b)("pre",null,Object(C.b)("code",Object(i.a)({parentName:"pre"},{className:"language-questdb-sql",metastring:'title="Using recursively"',title:'"Using','recursively"':!0}),"WITH first_10_users AS (SELECT * FROM users limit 10),\nfirst_5_users AS (SELECT * FROM first_10_users limit 5)\nSELECT user_name FROM first_5_users;\n")),Object(C.b)("pre",null,Object(C.b)("code",Object(i.a)({parentName:"pre"},{className:"language-questdb-sql",metastring:'title="Flag whether individual trips are longer or shorter than average"',title:'"Flag',whether:!0,individual:!0,trips:!0,are:!0,longer:!0,or:!0,shorter:!0,than:!0,'average"':!0}),"WITH avg_distance AS (select avg(trip_distance) average FROM trips)\nSELECT pickup_datetime, trips.trip_distance > avg_distance.average longer_than_average\nFROM trips CROSS JOIN avg_distance;\n")))}l.isMDXComponent=!0},358:function(I,g,e){"use strict";e.r(g),g.default="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI2NTEiIGhlaWdodD0iODEiPgogICAgPGRlZnM+CiAgICAgICAgPHN0eWxlIHR5cGU9InRleHQvY3NzIj4KICAgICAgICAgICAgQG5hbWVzcGFjZSAiaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciOwogICAgICAgICAgICAubGluZSAgICAgICAgICAgICAgICAge2ZpbGw6IG5vbmU7IHN0cm9rZTogIzYzNjI3Mzt9CiAgICAgICAgICAgIC5ib2xkLWxpbmUgICAgICAgICAgICB7c3Ryb2tlOiAjNjM2MjczOyBzaGFwZS1yZW5kZXJpbmc6IGNyaXNwRWRnZXM7IHN0cm9rZS13aWR0aDogMjsgfQogICAgICAgICAgICAudGhpbi1saW5lICAgICAgICAgICAge3N0cm9rZTogIzYzNjI3Mzsgc2hhcGUtcmVuZGVyaW5nOiBjcmlzcEVkZ2VzfQogICAgICAgICAgICAuZmlsbGVkICAgICAgICAgICAgICAge2ZpbGw6ICM2MzYyNzM7IHN0cm9rZTogbm9uZTt9CiAgICAgICAgICAgIHRleHQudGVybWluYWwgICAgICAgICB7Zm9udC1mYW1pbHk6IC1hcHBsZS1zeXN0ZW0sIEJsaW5rTWFjU3lzdGVtRm9udCwgIlNlZ29lIFVJIiwgUm9ib3RvLCBVYnVudHUsIENhbnRhcmVsbCwgSGVsdmV0aWNhLCBzYW5zLXNlcmlmOwogICAgICAgICAgICBmb250LXNpemU6IDEycHg7CiAgICAgICAgICAgIGZpbGw6ICNmZmZmZmY7CiAgICAgICAgICAgIGZvbnQtd2VpZ2h0OiBib2xkOwogICAgICAgICAgICB9CiAgICAgICAgICAgIHRleHQubm9udGVybWluYWwgICAgICB7Zm9udC1mYW1pbHk6IC1hcHBsZS1zeXN0ZW0sIEJsaW5rTWFjU3lzdGVtRm9udCwgIlNlZ29lIFVJIiwgUm9ib3RvLCBVYnVudHUsIENhbnRhcmVsbCwgSGVsdmV0aWNhLCBzYW5zLXNlcmlmOwogICAgICAgICAgICBmb250LXNpemU6IDEycHg7CiAgICAgICAgICAgIGZpbGw6ICNlMjg5YTQ7CiAgICAgICAgICAgIGZvbnQtd2VpZ2h0OiBub3JtYWw7CiAgICAgICAgICAgIH0KICAgICAgICAgICAgdGV4dC5yZWdleHAgICAgICAgICAgIHtmb250LWZhbWlseTogLWFwcGxlLXN5c3RlbSwgQmxpbmtNYWNTeXN0ZW1Gb250LCAiU2Vnb2UgVUkiLCBSb2JvdG8sIFVidW50dSwgQ2FudGFyZWxsLCBIZWx2ZXRpY2EsIHNhbnMtc2VyaWY7CiAgICAgICAgICAgIGZvbnQtc2l6ZTogMTJweDsKICAgICAgICAgICAgZmlsbDogIzAwMTQxRjsKICAgICAgICAgICAgZm9udC13ZWlnaHQ6IG5vcm1hbDsKICAgICAgICAgICAgfQogICAgICAgICAgICByZWN0LCBjaXJjbGUsIHBvbHlnb24ge2ZpbGw6IG5vbmU7IHN0cm9rZTogbm9uZTt9CiAgICAgICAgICAgIHJlY3QudGVybWluYWwgICAgICAgICB7ZmlsbDogbm9uZTsgc3Ryb2tlOiAjYmUyZjViO30KICAgICAgICAgICAgcmVjdC5ub250ZXJtaW5hbCAgICAgIHtmaWxsOiByZ2JhKDI1NSwyNTUsMjU1LDAuMSk7IHN0cm9rZTogbm9uZTt9CiAgICAgICAgICAgIHJlY3QudGV4dCAgICAgICAgICAgICB7ZmlsbDogbm9uZTsgc3Ryb2tlOiBub25lO30KICAgICAgICAgICAgcG9seWdvbi5yZWdleHAgICAgICAgIHtmaWxsOiAjQzdFQ0ZGOyBzdHJva2U6ICMwMzhjYmM7fQogICAgICAgIDwvc3R5bGU+CiAgICA8L2RlZnM+CiAgICA8cG9seWdvbiBwb2ludHM9IjkgNjEgMSA1NyAxIDY1Ii8+CiAgICA8cG9seWdvbiBwb2ludHM9IjE3IDYxIDkgNTcgOSA2NSIvPgogICAgPHJlY3QgeD0iMzEiIHk9IjQ3IiB3aWR0aD0iNTgiIGhlaWdodD0iMzIiIHJ4PSIxMCIvPgogICAgPHJlY3QgeD0iMjkiIHk9IjQ1IiB3aWR0aD0iNTgiIGhlaWdodD0iMzIiIGNsYXNzPSJ0ZXJtaW5hbCIgcng9IjEwIi8+CiAgICA8dGV4dCBjbGFzcz0idGVybWluYWwiIHg9IjM5IiB5PSI2NSI+V0lUSDwvdGV4dD48YSB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeGxpbms6aHJlZj0iI3N1YlF1ZXJ5TmFtZSIgeGxpbms6dGl0bGU9InN1YlF1ZXJ5TmFtZSI+CiAgICA8cmVjdCB4PSIxMjkiIHk9IjQ3IiB3aWR0aD0iMTE2IiBoZWlnaHQ9IjMyIi8+CiAgICA8cmVjdCB4PSIxMjciIHk9IjQ1IiB3aWR0aD0iMTE2IiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iMTM3IiB5PSI2NSI+c3ViUXVlcnlOYW1lPC90ZXh0PjwvYT48cmVjdCB4PSIyNjUiIHk9IjQ3IiB3aWR0aD0iMzgiIGhlaWdodD0iMzIiIHJ4PSIxMCIvPgogICAgPHJlY3QgeD0iMjYzIiB5PSI0NSIgd2lkdGg9IjM4IiBoZWlnaHQ9IjMyIiBjbGFzcz0idGVybWluYWwiIHJ4PSIxMCIvPgogICAgPHRleHQgY2xhc3M9InRlcm1pbmFsIiB4PSIyNzMiIHk9IjY1Ij5BUzwvdGV4dD4KICAgIDxyZWN0IHg9IjMyMyIgeT0iNDciIHdpZHRoPSIyNiIgaGVpZ2h0PSIzMiIgcng9IjEwIi8+CiAgICA8cmVjdCB4PSIzMjEiIHk9IjQ1IiB3aWR0aD0iMjYiIGhlaWdodD0iMzIiIGNsYXNzPSJ0ZXJtaW5hbCIgcng9IjEwIi8+CiAgICA8dGV4dCBjbGFzcz0idGVybWluYWwiIHg9IjMzMSIgeT0iNjUiPig8L3RleHQ+PGEgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHhsaW5rOmhyZWY9IiNzdWJRdWVyeSIgeGxpbms6dGl0bGU9InN1YlF1ZXJ5Ij4KICAgIDxyZWN0IHg9IjM2OSIgeT0iNDciIHdpZHRoPSI4MCIgaGVpZ2h0PSIzMiIvPgogICAgPHJlY3QgeD0iMzY3IiB5PSI0NSIgd2lkdGg9IjgwIiBoZWlnaHQ9IjMyIiBjbGFzcz0ibm9udGVybWluYWwiLz4KICAgIDx0ZXh0IGNsYXNzPSJub250ZXJtaW5hbCIgeD0iMzc3IiB5PSI2NSI+c3ViUXVlcnk8L3RleHQ+PC9hPjxyZWN0IHg9IjQ2OSIgeT0iNDciIHdpZHRoPSIyNiIgaGVpZ2h0PSIzMiIgcng9IjEwIi8+CiAgICA8cmVjdCB4PSI0NjciIHk9IjQ1IiB3aWR0aD0iMjYiIGhlaWdodD0iMzIiIGNsYXNzPSJ0ZXJtaW5hbCIgcng9IjEwIi8+CiAgICA8dGV4dCBjbGFzcz0idGVybWluYWwiIHg9IjQ3NyIgeT0iNjUiPik8L3RleHQ+CiAgICA8cmVjdCB4PSIxMjkiIHk9IjMiIHdpZHRoPSIyNCIgaGVpZ2h0PSIzMiIgcng9IjEwIi8+CiAgICA8cmVjdCB4PSIxMjciIHk9IjEiIHdpZHRoPSIyNCIgaGVpZ2h0PSIzMiIgY2xhc3M9InRlcm1pbmFsIiByeD0iMTAiLz4KICAgIDx0ZXh0IGNsYXNzPSJ0ZXJtaW5hbCIgeD0iMTM3IiB5PSIyMSI+LDwvdGV4dD48YSB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeGxpbms6aHJlZj0iI21haW5RdWVyeSIgeGxpbms6dGl0bGU9Im1haW5RdWVyeSI+CiAgICA8cmVjdCB4PSI1MzUiIHk9IjQ3IiB3aWR0aD0iODgiIGhlaWdodD0iMzIiLz4KICAgIDxyZWN0IHg9IjUzMyIgeT0iNDUiIHdpZHRoPSI4OCIgaGVpZ2h0PSIzMiIgY2xhc3M9Im5vbnRlcm1pbmFsIi8+CiAgICA8dGV4dCBjbGFzcz0ibm9udGVybWluYWwiIHg9IjU0MyIgeT0iNjUiPm1haW5RdWVyeTwvdGV4dD48L2E+PHBhdGggeG1sbnM6c3ZnPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgY2xhc3M9ImxpbmUiIGQ9Im0xNyA2MSBoMiBtMCAwIGgxMCBtNTggMCBoMTAgbTIwIDAgaDEwIG0xMTYgMCBoMTAgbTAgMCBoMTAgbTM4IDAgaDEwIG0wIDAgaDEwIG0yNiAwIGgxMCBtMCAwIGgxMCBtODAgMCBoMTAgbTAgMCBoMTAgbTI2IDAgaDEwIG0tNDA2IDAgbDIwIDAgbS0xIDAgcS05IDAgLTkgLTEwIGwwIC0yNCBxMCAtMTAgMTAgLTEwIG0zODYgNDQgbDIwIDAgbS0yMCAwIHExMCAwIDEwIC0xMCBsMCAtMjQgcTAgLTEwIC0xMCAtMTAgbS0zODYgMCBoMTAgbTI0IDAgaDEwIG0wIDAgaDM0MiBtMjAgNDQgaDEwIG04OCAwIGgxMCBtMyAwIGgtMyIvPgogICAgPHBvbHlnb24gcG9pbnRzPSI2NDEgNjEgNjQ5IDU3IDY0OSA2NSIvPgogICAgPHBvbHlnb24gcG9pbnRzPSI2NDEgNjEgNjMzIDU3IDYzMyA2NSIvPjwvc3ZnPg=="}}]);