(window.webpackJsonp=window.webpackJsonp||[]).push([[70],{138:function(e,t,a){"use strict";a.r(t),a.d(t,"frontMatter",(function(){return l})),a.d(t,"metadata",(function(){return i})),a.d(t,"rightToc",(function(){return o})),a.d(t,"default",(function(){return d}));var n=a(2),b=a(6),r=(a(0),a(236)),l={title:"Web Console",sidebar_label:"Web Console",description:"Web Console reference documentation."},i={unversionedId:"__uneeded/reference/web-console",id:"__uneeded/reference/web-console",isDocsHomePage:!1,title:"Web Console",description:"Web Console reference documentation.",source:"@site/docs/__uneeded/reference/web-console.md",slug:"/__uneeded/reference/web-console",permalink:"/docs/__uneeded/reference/web-console",version:"current",sidebar_label:"Web Console"},o=[{value:"SQL Editor",id:"sql-editor",children:[{value:"Shortcuts",id:"shortcuts",children:[]},{value:"Behaviour",id:"behaviour",children:[]}]},{value:"Import",id:"import",children:[{value:"Import details",id:"import-details",children:[]},{value:"Import statuses",id:"import-statuses",children:[]},{value:"Amending the schema",id:"amending-the-schema",children:[]},{value:"Custom import",id:"custom-import",children:[]}]}],c={rightToc:o};function d(e){var t=e.components,l=Object(b.a)(e,["components"]);return Object(r.b)("wrapper",Object(n.a)({},c,l,{components:t,mdxType:"MDXLayout"}),Object(r.b)("p",null,"This is a reference for the Console. If you want to learn how to use it, we\nsuggest you follow the ",Object(r.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/guide/web-console/"}),"Guide"),"."),Object(r.b)("h2",{id:"sql-editor"},"SQL Editor"),Object(r.b)("h3",{id:"shortcuts"},"Shortcuts"),Object(r.b)("table",null,Object(r.b)("thead",{parentName:"table"},Object(r.b)("tr",{parentName:"thead"},Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Command"),Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Action"))),Object(r.b)("tbody",{parentName:"table"},Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Run query"),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"f9")," or ",Object(r.b)("inlineCode",{parentName:"td"},"ctrl/cmd + enter"))),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Locate cursor"),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"f2"),", use this to focus the SQL editor on your cursor in order to locate it")))),Object(r.b)("h3",{id:"behaviour"},"Behaviour"),Object(r.b)("p",null,"As you can write numerous SQL commands separated by semicolon, the Web Console\nuses the following logic to decide what command to execute:"),Object(r.b)("ul",null,Object(r.b)("li",{parentName:"ul"},"Check if you highlighted a query or part of it, if yes then it will be\nexecuted, otherwise:"),Object(r.b)("li",{parentName:"ul"},"Verify if the cursor is within a SQL statement, if yes, the wrapping statement\nwill be executed, otherwise:"),Object(r.b)("li",{parentName:"ul"},"Find out if the cursor is on the same line as a SQL statement and after the\nsemicolon, if yes, this statement will be executed, finally:"),Object(r.b)("li",{parentName:"ul"},"If the cursor is on a line that does not contain a SQL statement, the next\nencountered statement will be executed. If there is no statement following the\ncursor, the previous statement will be used.")),Object(r.b)("h2",{id:"import"},"Import"),Object(r.b)("h3",{id:"import-details"},"Import details"),Object(r.b)("p",null,"Description of the fields in the import details table"),Object(r.b)("table",null,Object(r.b)("thead",{parentName:"table"},Object(r.b)("tr",{parentName:"thead"},Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Column"),Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Description"))),Object(r.b)("tbody",{parentName:"table"},Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"File name")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Name of the file imported. If imported from copy & paste, an automatically-generated file name")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"Size")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Size of the imported file")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"Total rows")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Number of rows successfully imported")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"Failed rows")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Number of rows that failed to import")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"Header row")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Whether the dataset has been recognised to have a header row or not")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"Status")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Status of the import. See below")))),Object(r.b)("h3",{id:"import-statuses"},"Import statuses"),Object(r.b)("p",null,"Description of the import statuses"),Object(r.b)("table",null,Object(r.b)("thead",{parentName:"table"},Object(r.b)("tr",{parentName:"thead"},Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Status"),Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Description"))),Object(r.b)("tbody",{parentName:"table"},Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"importing")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Data is currently being imported")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"failed")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Import failed, nothing was imported")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"imported in [time]")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Import is finished. The completion time will be displayed next to the status")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"exists")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"You are trying to import a file that already exists. To import it regardless, you can either ",Object(r.b)("strong",{parentName:"td"},"append")," or ",Object(r.b)("strong",{parentName:"td"},"override"),". See ",Object(r.b)("a",Object(n.a)({parentName:"td"},{href:"#custom-import"}),"importing again")," for a more exhaustive description")))),Object(r.b)("h3",{id:"amending-the-schema"},"Amending the schema"),Object(r.b)("p",null,"Although the schema is automatically detected, you can amend the type for any\ncolumn using the following steps:"),Object(r.b)("ul",null,Object(r.b)("li",{parentName:"ul"},"Click on the file you want to amend in the Import screen. The schema will be\ndisplayed."),Object(r.b)("li",{parentName:"ul"},"Find and click on the column which type you want to change."),Object(r.b)("li",{parentName:"ul"},"You will then need to ",Object(r.b)("a",Object(n.a)({parentName:"li"},{href:"#custom-import"}),"re-trigger the import"),".")),Object(r.b)("img",{alt:" Change the schema in the Web Console when importing data",className:"screenshot--shadow screenshot--docs",src:"/img/docs/console/amendType.jpg"}),Object(r.b)("h3",{id:"custom-import"},"Custom import"),Object(r.b)("p",null,"You can amend the import behaviour with the following options. This will trigger\nto import the data again."),Object(r.b)("table",null,Object(r.b)("thead",{parentName:"table"},Object(r.b)("tr",{parentName:"thead"},Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Option"),Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Name"),Object(r.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"Description"))),Object(r.b)("tbody",{parentName:"table"},Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"A")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Append"),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Uploaded data will be appended at the end of the table")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"O")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Override"),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Uploaded data will override existing data in the table")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"LEV")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Skip lines with extra values"),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Skips rows that contains dangling values that don't fit the schema")),Object(r.b)("tr",{parentName:"tbody"},Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(r.b)("inlineCode",{parentName:"td"},"H")),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Header row"),Object(r.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Flag whether the first row should be considered header")))),Object(r.b)("p",null,"To start the import, click the following button:"),Object(r.b)("p",null,Object(r.b)("img",{alt:"Upload button from the Web Console",src:a(336).default})))}d.isMDXComponent=!0},336:function(e,t,a){"use strict";a.r(t),t.default="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAYAAACM/rhtAAAMSmlDQ1BJQ0MgUHJvZmlsZQAASImVVwdUU8kanltSSWiBUKSE3kQp0qWE0CIISBVshCSQUGJMCCJ2FlkF1y4ioK7oqoiLrgWQtaKudVHs/aGIysq6WLCh8iYF1tXz3jvvP2fu/fLPP99fMjN3BgCdWp5UmofqApAvKZAlRIayJqals0hdAAEo0AKuAOfx5VJ2fHwMgDL0/qe8uQ6toVxxVXJ92/9fRU8glPMBQOIhzhTI+fkQ7wcAL+VLZQUAEH2h3mZmgVSJJ0NsIIMBQixV4mw1LlXiTDWuUtkkJXAg3gUAmcbjybIB0G6BelYhPxvyaN+E2E0iEEsA0CFDHMQX8QQQR0E8Mj9/uhJDO+CY+QVP9j84M4c5ebzsYazORSXkMLFcmseb9X+W439Lfp5iyIc9bDSRLCpBmTOs283c6dFKTIO4V5IZGwexPsTvxAKVPcQoVaSISlbbo2Z8OQfWDDAhdhPwwqIhNoM4QpIXG6PRZ2aJI7gQwxmCFokLuEmasYuF8vBEDWetbHpC3BDOknHYmrGNPJnKr9L+pCI3ma3hvykScof4XxeLklLVMWPUQnFKLMTaEDPluYnRahvMtljEiR2ykSkSlPHbQuwvlESGqvmxqVmyiASNvSxfPpQvtlgk5sZqcHWBKClKw7OLz1PFbwxxi1DCTh7iEconxgzlIhCGhatzxy4JJcmafLFOaUFogmbsS2levMYepwrzIpV6a4jN5IWJmrF4UAGckGp+PFZaEJ+kjhPPzOGNi1fHgxeBGMABYYAFFLBlgukgB4jbe5t74S91TwTgARnIBkK4KtWaoRGpqh4JfCaCYvAnREIgHx4XquoVgkKo/zSsVT9dQZaqt1A1Ihc8hjgfRIM8+FuhGiUZ9pYCHkGN+BvvfBhrHmzKvm91bKiJ0WgUQ7wsnSFLYjgxjBhFjCA64aZ4EB6Ax8BnCGweuC/uNxTt3/aEx4QOwkPCNUIn4dY0cYnsq3xYYDzohB4iNDlnfpkzbg9ZvfBQPBDyQ26ciZsCV3wM9MTGg6FvL6jlaCJXZv819z9y+KLqGjuKGwWlGFFCKI5fj9R21vYaZlHW9MsKqWPNHK4rZ7jna/+cLyotgO/ory2xxdg+7DR2HDuLHcKaAQs7irVgF7DDSjw8ix6pZtGQtwRVPLmQR/yNP57Gp7KScrcGtx63j+q+AmGRcn8EnOnSWTJxtqiAxYY7v5DFlfBHjWR5uLn7AaD8jqi3qVdM1fcBYZ77W1fyGoBAweDg4KG/dTFwTe9fBAD18d86hyNwOzAC4EwFXyErVOtw5YMAqEAHrigTYAFsgCPMxwN4gwAQAsLBOBAHkkAamAqrLILzWQZmgjlgISgDFWAFWAuqwSawBewAP4O9oBkcAsfBb+A8uASugTtw9nSDZ6APvAEDCIKQEDrCQEwQS8QOcUE8EF8kCAlHYpAEJA3JQLIRCaJA5iDfIRXIKqQa2YzUI78gB5HjyFmkA7mFPEB6kJfIBxRDaagBao7ao6NRX5SNRqNJ6BQ0G52BFqOl6DK0Cq1Dd6FN6HH0PHoN7USfof0YwLQwJmaFuWK+GAeLw9KxLEyGzcPKsUqsDmvEWuH/fAXrxHqx9zgRZ+As3BXO4Cg8GefjM/B5+FK8Gt+BN+En8Sv4A7wP/0ygE8wILgR/ApcwkZBNmEkoI1QSthEOEE7B1dRNeEMkEplEB6IPXI1pxBzibOJS4gbibuIxYgexi9hPIpFMSC6kQFIciUcqIJWR1pN2kY6SLpO6Se/IWmRLsgc5gpxOlpBLyJXkneQj5MvkJ+QBii7FjuJPiaMIKLMoyylbKa2Ui5RuygBVj+pADaQmUXOoC6lV1EbqKepd6istLS1rLT+tCVpirQVaVVp7tM5oPdB6T9OnOdM4tMk0BW0ZbTvtGO0W7RWdTrenh9DT6QX0ZfR6+gn6ffo7bYb2KG2utkB7vnaNdpP2Ze3nOhQdOx22zlSdYp1KnX06F3V6dSm69rocXZ7uPN0a3YO6N3T79Rh67npxevl6S/V26p3Ve6pP0rfXD9cX6Jfqb9E/od/FwBg2DA6Dz/iOsZVxitFtQDRwMOAa5BhUGPxs0G7QZ6hvOMYwxbDIsMbwsGEnE2PaM7nMPOZy5l7mdeYHI3MjtpHQaIlRo9Flo7fGI4xDjIXG5ca7ja8ZfzBhmYSb5JqsNGk2uWeKmzqbTjCdabrR9JRp7wiDEQEj+CPKR+wdcdsMNXM2SzCbbbbF7IJZv7mFeaS51Hy9+QnzXgumRYhFjsUaiyMWPZYMyyBLseUay6OWf7AMWWxWHquKdZLVZ2VmFWWlsNps1W41YO1gnWxdYr3b+p4N1cbXJstmjU2bTZ+tpe142zm2Dba37Sh2vnYiu3V2p+3e2jvYp9p/b99s/9TB2IHrUOzQ4HDXke4Y7DjDsc7xqhPRydcp12mD0yVn1NnLWeRc43zRBXXxdhG7bHDpGEkY6TdSMrJu5A1XmivbtdC1wfXBKOaomFElo5pHPR9tOzp99MrRp0d/dvNyy3Pb6nbHXd99nHuJe6v7Sw9nD75HjcdVT7pnhOd8zxbPF2NcxgjHbBxz04vhNd7re682r0/ePt4y70bvHh9bnwyfWp8bvga+8b5Lfc/4EfxC/eb7HfJ77+/tX+C/1/+vANeA3ICdAU/HOowVjt06tivQOpAXuDmwM4gVlBH0Y1BnsFUwL7gu+GGITYggZFvIE7YTO4e9i/081C1UFnog9C3HnzOXcywMC4sMKw9rD9cPTw6vDr8fYR2RHdEQ0RfpFTk78lgUISo6amXUDa45l8+t5/aN8xk3d9zJaFp0YnR19MMY5xhZTOt4dPy48avH3421i5XENseBOG7c6rh78Q7xM+J/nUCcED+hZsLjBPeEOQmnExmJ0xJ3Jr5JCk1annQn2TFZkdyWopMyOaU+5W1qWOqq1M6JoyfOnXg+zTRNnNaSTkpPSd+W3j8pfNLaSd2TvSaXTb4+xWFK0ZSzU02n5k09PE1nGm/avgxCRmrGzoyPvDheHa8/k5tZm9nH5/DX8Z8JQgRrBD3CQOEq4ZOswKxVWU+zA7NXZ/eIgkWVol4xR1wtfpETlbMp521uXO723MG81Lzd+eT8jPyDEn1JruTkdIvpRdM7pC7SMmnnDP8Za2f0yaJl2+SIfIq8pcAAHtgvKBwVixQPCoMKawrfzUyZua9Ir0hSdGGW86wls54URxT/NBufzZ/dNsdqzsI5D+ay526eh8zLnNc232Z+6fzuBZELdiykLsxd+HuJW8mqktffpX7XWmpeuqC0a1HkooYy7TJZ2Y3vA77ftBhfLF7cvsRzyfoln8sF5ecq3CoqKz4u5S8994P7D1U/DC7LWta+3Hv5xhXEFZIV11cGr9yxSm9V8aqu1eNXN61hrSlf83rttLVnK8dUblpHXadY11kVU9Wy3nb9ivUfq0XV12pCa3bXmtUuqX27QbDh8saQjY2bzDdVbPrwo/jHm5sjNzfV2ddVbiFuKdzyeGvK1tM/+f5Uv810W8W2T9sl2zt3JOw4We9TX7/TbOfyBrRB0dCza/KuSz+H/dzS6Nq4eTdzd8UesEex549fMn65vjd6b9s+332N++321x5gHChvQppmNfU1i5o7W9JaOg6OO9jWGtB64NdRv24/ZHWo5rDh4eVHqEdKjwweLT7af0x6rPd49vGutmltd05MPHH15IST7aeiT535LeK3E6fZp4+eCTxz6Kz/2YPnfM81n/c+33TB68KB371+P9Du3d500ediyyW/S60dYzuOXA6+fPxK2JXfrnKvnr8We63jevL1mzcm3+i8Kbj59FberRe3C28P3Flwl3C3/J7uvcr7Zvfr/uX0r92d3p2HH4Q9uPAw8eGdLn7Xs0fyRx+7Sx/TH1c+sXxS/9Tj6aGeiJ5Lf0z6o/uZ9NlAb9mfen/WPnd8vv+vkL8u9E3s634hezH4cukrk1fbX4953dYf33//Tf6bgbfl70ze7Xjv+/70h9QPTwZmfiR9rPrk9Kn1c/Tnu4P5g4NSnoynOgpgsKFZWQC83A4APQ0AxiV4fpikvuepBFHfTVUI/CesvguqxBuARvhSHtc5xwDYA5v9AsgdAoDyqJ4UAlBPz+GmEXmWp4eaiwZvPIR3g4OvzAEgtQLwSTY4OLBhcPDTVhjsLQCOzVDfL5VChHeDH0OU6JqxYAH4Sv4NWFd/XGjpTwUAAAAJcEhZcwAAFiUAABYlAUlSJPAAAAIEaVRYdFhNTDpjb20uYWRvYmUueG1wAAAAAAA8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIiB4OnhtcHRrPSJYTVAgQ29yZSA1LjQuMCI+CiAgIDxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+CiAgICAgIDxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSIiCiAgICAgICAgICAgIHhtbG5zOmV4aWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20vZXhpZi8xLjAvIgogICAgICAgICAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29tL3RpZmYvMS4wLyI+CiAgICAgICAgIDxleGlmOlBpeGVsWURpbWVuc2lvbj4xODQ8L2V4aWY6UGl4ZWxZRGltZW5zaW9uPgogICAgICAgICA8ZXhpZjpQaXhlbFhEaW1lbnNpb24+MzIyPC9leGlmOlBpeGVsWERpbWVuc2lvbj4KICAgICAgICAgPHRpZmY6T3JpZW50YXRpb24+MTwvdGlmZjpPcmllbnRhdGlvbj4KICAgICAgPC9yZGY6RGVzY3JpcHRpb24+CiAgIDwvcmRmOlJERj4KPC94OnhtcG1ldGE+Co7RGvEAAAIeSURBVFgJ7ZhNLwNBGMf/VHWbJhrSFwdOqKNyUxJCHIgvIF7iStCmVwnfgNSXEOXk7SIk+ASatm6oGxahiaZpFp1NdqKr7Ww7k6aHnWSzz8w+M/Ob/8w+O7MNP/mEOk6NdcymopmAvDNkKmgqyKsAb33hazD1+Ih0Os3LRes3UUuAcXl1jd29KJzOFoSDQXi9Hu5WhSmowRGij49PbEUieHp6rg/Av3AakShIbgWLwYmE5AIsBycKsmpAI3AiIKsCTCST6tuqARi5a2tSURQj7tSnKkBF+YbVaqWNGDUkSUI2mzXqrvo1iNqwrobCyOVyBZ0PBwKYm50pKKs0U5WClXbC428C8qhH6poK1o2CJITok2T/X6b3YeUNhZl4IoFk8hbFDtC9vh74+/rwkErh/v6B9mexWDDQ74fD4cDxySm+Mhn6TDMkmw1DgUG4XC6t6N+dCfj29o71jU2UOt+TgL2yvAQCqk8kLh4eHePs/EL/iOZ9Pd0Ih4I0rzeYG1b59bUkHGmMQGxHdlQV2v9sUMkn7S6vKOvL8fwi65kK8kzAAu8yGVmWQS7RqSZhprurKz+Na+pRgAxgYnwMiwvzhsZSE0DyorjdbjQ321So1rY2eDxuQ4DMKbYLCBU3sRjIpaXo/oFmgtU+E7CzowNjoyOI5/eAon/USZIN01OTFLaYwQwzxSrVsqwma5BnQCYgj3qkrqmgqSCvArz1fwE6e8R9NkWs9gAAAABJRU5ErkJggg=="}}]);