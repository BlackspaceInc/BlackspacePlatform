export default {
  "title": "Blackspace Platform",
  "tagline": "Empowering minority owned businesses to better serve their clients <br/>and achieve more.",
  "favicon": "img/favicon/favicon-32x32.png",
  "customFields": {
    "description": "Empowering minority owned businesses to better serve their clients and achieve more."
  },
  "onBrokenLinks": "log",
  "themeConfig": {
    "colorMode": {
      "defaultMode": "light",
      "disableSwitch": false,
      "respectPrefersColorScheme": true,
      "switchConfig": {
        "darkIcon": "🌙",
        "darkIconStyle": {
          "marginLeft": "2px"
        },
        "lightIcon": "☀",
        "lightIconStyle": {
          "marginLeft": "1px"
        }
      }
    },
    "image": "img/share.jpg",
    "announcementBar": {
      "id": "supportus",
      "backgroundColor": "#1064d3",
      "textColor": "white",
      "content": "⭐️ If you like Blackspace, give it a star on <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://github.com/BlackspaceInc/BlackspacePlatform\">GitHub</a>! ⭐️"
    },
    "prism": {
      "additionalLanguages": [
        "nginx"
      ],
      "defaultLanguage": "javascript",
      "theme": {
        "plain": {
          "color": "#393A34",
          "backgroundColor": "#f6f8fa"
        },
        "styles": [
          {
            "types": [
              "comment",
              "prolog",
              "doctype",
              "cdata"
            ],
            "style": {
              "color": "#999988",
              "fontStyle": "italic"
            }
          },
          {
            "types": [
              "namespace"
            ],
            "style": {
              "opacity": 0.7
            }
          },
          {
            "types": [
              "string",
              "attr-value"
            ],
            "style": {
              "color": "#e3116c"
            }
          },
          {
            "types": [
              "punctuation",
              "operator"
            ],
            "style": {
              "color": "#393A34"
            }
          },
          {
            "types": [
              "entity",
              "url",
              "symbol",
              "number",
              "boolean",
              "variable",
              "constant",
              "property",
              "regex",
              "inserted"
            ],
            "style": {
              "color": "#36acaa"
            }
          },
          {
            "types": [
              "atrule",
              "keyword",
              "attr-name",
              "selector"
            ],
            "style": {
              "color": "#00a4db"
            }
          },
          {
            "types": [
              "function",
              "deleted",
              "tag"
            ],
            "style": {
              "color": "#d73a49"
            }
          },
          {
            "types": [
              "function-variable"
            ],
            "style": {
              "color": "#6f42c1"
            }
          },
          {
            "types": [
              "tag",
              "selector",
              "keyword"
            ],
            "style": {
              "color": "#00009f"
            }
          }
        ]
      },
      "darkTheme": {
        "plain": {
          "color": "#F8F8F2",
          "backgroundColor": "#282A36"
        },
        "styles": [
          {
            "types": [
              "prolog",
              "constant",
              "builtin"
            ],
            "style": {
              "color": "rgb(189, 147, 249)"
            }
          },
          {
            "types": [
              "inserted",
              "function"
            ],
            "style": {
              "color": "rgb(80, 250, 123)"
            }
          },
          {
            "types": [
              "deleted"
            ],
            "style": {
              "color": "rgb(255, 85, 85)"
            }
          },
          {
            "types": [
              "changed"
            ],
            "style": {
              "color": "rgb(255, 184, 108)"
            }
          },
          {
            "types": [
              "punctuation",
              "symbol"
            ],
            "style": {
              "color": "rgb(248, 248, 242)"
            }
          },
          {
            "types": [
              "string",
              "char",
              "tag",
              "selector"
            ],
            "style": {
              "color": "rgb(255, 121, 198)"
            }
          },
          {
            "types": [
              "keyword",
              "variable"
            ],
            "style": {
              "color": "rgb(189, 147, 249)",
              "fontStyle": "italic"
            }
          },
          {
            "types": [
              "comment"
            ],
            "style": {
              "color": "rgb(98, 114, 164)"
            }
          },
          {
            "types": [
              "attr-name"
            ],
            "style": {
              "color": "rgb(241, 250, 140)"
            }
          }
        ]
      }
    },
    "footer": {
      "logo": {
        "alt": "Blackspace Logo",
        "src": "img/twemoji_poodle.svg"
      },
      "links": [
        {
          "title": "Docs",
          "items": [
            {
              "label": "Introduction",
              "to": "docs/introduction"
            },
            {
              "label": "Technology",
              "to": "docs/technology"
            }
          ]
        },
        {
          "title": "Products",
          "items": [
            {
              "label": "Product Overview",
              "to": "docs/technology/products/overview"
            },
            {
              "label": "Marketplace",
              "to": "docs/technology/products/marketplace"
            },
            {
              "label": " Business",
              "to": "docs/technology/products/business"
            },
            {
              "label": " Analytics",
              "to": "docs/technology/products/analytics"
            },
            {
              "label": " Makers",
              "to": "docs/technology/products/makers"
            },
            {
              "label": " 3rd Party Integrations",
              "to": "docs/technology/products/integrations"
            },
            {
              "label": " Ads",
              "to": "docs/technology/products/ads"
            },
            {
              "label": " AI & Research",
              "to": "docs/technology/products/research"
            },
            {
              "label": " VC",
              "to": "docs/technology/products/vc"
            }
          ]
        },
        {
          "title": "Social",
          "items": [
            {
              "label": "Blog",
              "href": "https://the-guild.dev/blog"
            },
            {
              "label": "GitHub",
              "href": "https://github.com/kamilkisiela/graphql-inspector"
            },
            {
              "label": "Twitter",
              "href": "https://twitter.com/kamilkisiela"
            },
            {
              "label": "LinkedIn",
              "href": "https://www.linkedin.com/company/the-guild-software"
            }
          ]
        }
      ],
      "copyright": "Made with ❤️ at Blackspace. Apache 2.0 License. Built with Docusaurus.",
      "style": "light"
    },
    "navbar": {
      "title": "Blackspace Platform",
      "logo": {
        "alt": "Blackspace Logo",
        "src": "img/twemoji_poodle.svg"
      },
      "hideOnScroll": true,
      "items": [
        {
          "to": "blog",
          "label": "Blog",
          "position": "left"
        },
        {
          "to": "docs/introduction",
          "label": "Docs",
          "position": "left"
        },
        {
          "href": "https://github.com/BlackspaceInc/BlackspacePlatform",
          "position": "left",
          "label": "Source"
        },
        {
          "href": "https://github.com/BlackspaceInc/BlackspacePlatform",
          "position": "right",
          "className": "header-github-link",
          "aria-label": "GitHub repository"
        }
      ]
    },
    "sidebarCollapsible": true,
    "algolia": {
      "apiKey": "7e47115263beea4eb52978a771750414",
      "indexName": "docs",
      "algoliaOptions": {},
      "appId": "BH4D9OD16A"
    }
  },
  "presets": [
    [
      "@docusaurus/preset-classic",
      {
        "theme": {
          "customCss": [
            "/Users/yoanyomba/go/src/github.com/BlackspaceInc/BlackspacePlatform/docs/src/css/custom.css"
          ]
        },
        "blog": {
          "path": "blog/engineering",
          "routeBasePath": "blog",
          "blogDescription": "Blackspace Engineering"
        },
        "docs": {
          "path": "docs",
          "routeBasePath": "docs",
          "sidebarPath": "/Users/yoanyomba/go/src/github.com/BlackspaceInc/BlackspacePlatform/docs/sidebars.js",
          "editUrl": "https://github.com/BlackspaceInc/BlackspacePlatform/edit/master/website/"
        }
      }
    ]
  ],
  "url": "https://blackspaceinc.github.io",
  "baseUrl": "/",
  "organizationName": "BlackspaceInc",
  "projectName": "BlackspacePlatform",
  "plugins": [],
  "themes": []
};