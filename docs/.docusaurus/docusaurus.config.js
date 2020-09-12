export default {
  "title": "Better Ecommerce",
  "tagline": "Blackspace is the fastest social ecommerce solution for minority owned business owners",
  "url": "https://blackspaceInc.io",
  "baseUrl": "/",
  "favicon": "/img/favicon.png",
  "organizationName": "Blackspace",
  "projectName": "blackspace",
  "customFields": {
    "artifactHubUrl": "https://artifacthub.io/packages/helm/BlackspaceInc/BlackspacePlatform",
    "copyright": "Copyright © 2020 Blackspace Platform",
    "description": "Blackspace is an ecommerce platform empowering minoriting owned businesses to achieve more",
    "dockerUrl": "https://hub.docker.com/repository/docker/blackspaceinc/frontend-service",
    "domain": "blackspaceInc.io",
    "githubUrl": "https://github.com/BlackspaceInc/",
    "helmVersion": "0.2.4",
    "linkedInUrl": "https://www.linkedin.com/company/blackspaceInc/",
    "oneLiner": "Ecommerce platform empowering minority owned businesses to achieve more",
    "slackUrl": "https://blackspace.slack.com/join/shared_invite/enQtNzk4Nzg4Mjc2MTE2LTEzZThjMzliMjUzMTBmYzVjYWNmM2UyNWJmNDdkMDYyZmE0ZDliZTQxN2EzNzk5MDE3Zjc1ZmJiZmFiZTIwMGY#/",
    "twitterUrl": "https://twitter.com/blackspace",
    "version": "5.0.3"
  },
  "plugins": [
    "/Users/yoanyomba/go/src/github.com/BlackspaceInc/BlackspacePlatform/docs/plugins/fetch-release.js",
    "/Users/yoanyomba/go/src/github.com/BlackspaceInc/BlackspacePlatform/docs/plugins/lint.js",
    "/Users/yoanyomba/go/src/github.com/BlackspaceInc/BlackspacePlatform/docs/plugins/manifest.js",
    [
      "@docusaurus/plugin-ideal-image",
      {
        "quality": 100,
        "steps": 2
      }
    ],
    [
      "@docusaurus/plugin-pwa",
      {
        "pwaHead": [
          {
            "tagName": "link",
            "rel": "icon",
            "href": "/img/favicon.png"
          },
          {
            "tagName": "link",
            "rel": "manifest",
            "href": "/manifest.webmanifest"
          },
          {
            "tagName": "meta",
            "name": "theme-color",
            "content": "#d14671"
          },
          {
            "tagName": "meta",
            "name": "apple-mobile-web-app-capable",
            "content": "yes"
          },
          {
            "tagName": "meta",
            "name": "apple-mobile-web-app-status-bar-style",
            "content": "#21222c"
          },
          {
            "tagName": "link",
            "rel": "apple-touch-icon",
            "href": "/img/favicon.png"
          },
          {
            "tagName": "link",
            "rel": "mask-icon",
            "href": "/img/favicon.png",
            "content": "#fff"
          },
          {
            "tagName": "meta",
            "name": "msapplication-TileImage",
            "content": "/img/favicon.png"
          },
          {
            "tagName": "meta",
            "name": "msapplication-TileColor",
            "content": "#21222c"
          }
        ]
      }
    ]
  ],
  "themeConfig": {
    "announcementBar": {
      "id": "supportus",
      "backgroundColor": "#1064d3",
      "textColor": "white",
      "content": "⭐️ If you like Blackspace, give it a star on <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://github.com/BlackspaceInc/BlackspacePlatform\">GitHub</a>! ⭐️",
      "isCloseable": true
    },
    "colorMode": {
      "defaultMode": "light",
      "disableSwitch": false,
      "respectPrefersColorScheme": true,
      "switchConfig": {
        "darkIcon": "🌜",
        "darkIconStyle": {},
        "lightIcon": "🌞",
        "lightIconStyle": {}
      }
    },
    "image": "/img/og.png",
    "gtag": {
      "trackingID": "GTM-PVR7M2G",
      "anonymizeIP": true
    },
    "prism": {
      "defaultLanguage": "blackspace",
      "theme": {
        "plain": {
          "color": "#f8f8f2",
          "backgroundColor": "#262833"
        },
        "styles": [
          {
            "types": [
              "prolog",
              "constant",
              "boolean",
              "builtin"
            ],
            "style": {
              "color": "#bd93f9"
            }
          },
          {
            "types": [
              "inserted",
              "function"
            ],
            "style": {
              "color": "#8be9fd"
            }
          },
          {
            "types": [
              "dataType"
            ],
            "style": {
              "color": "#8be9fd",
              "fontStyle": "italic"
            }
          },
          {
            "types": [
              "deleted"
            ],
            "style": {
              "color": "#ff5555"
            }
          },
          {
            "types": [
              "changed"
            ],
            "style": {
              "color": "#ffb86c"
            }
          },
          {
            "types": [
              "punctuation",
              "symbol"
            ],
            "style": {
              "color": "#f8f8f2"
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
              "color": "#f1fa8c"
            }
          },
          {
            "types": [
              "keyword",
              "variable",
              "operator"
            ],
            "style": {
              "color": "#ff79c6"
            }
          },
          {
            "types": [
              "number"
            ],
            "style": {
              "color": "#50fa7b"
            }
          },
          {
            "types": [
              "comment"
            ],
            "style": {
              "color": "#6272a4"
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
    "algolia": {
      "apiKey": "b2a69b4869a2a85284a82fb57519dcda",
      "indexName": "blackspace",
      "appId": "BH4D9OD16A"
    },
    "navbar": {
      "title": " ",
      "logo": {
        "alt": "Blackspace",
        "src": "/img/navbar/logo.ico"
      },
      "items": [
        {
          "label": "Getting Started",
          "position": "left",
          "items": [
            {
              "label": "Docker",
              "to": "/docs/guide/docker/",
              "position": "left"
            },
            {
              "label": "Kubernetes",
              "to": "/docs/guide/kubernetes/",
              "position": "left"
            }
          ]
        },
        {
          "label": "Documentation",
          "position": "left",
          "to": "/docs/introduction/",
          "activeBasePath": "docs"
        },
        {
          "label": "GitHub",
          "className": "navbar__item--github",
          "href": "https://github.com/BlackspaceInc/",
          "position": "right"
        },
        {
          "label": "Join Slack",
          "className": "navbar__item--slack",
          "href": "https://blackspace.slack.com/join/shared_invite/enQtNzk4Nzg4Mjc2MTE2LTEzZThjMzliMjUzMTBmYzVjYWNmM2UyNWJmNDdkMDYyZmE0ZDliZTQxN2EzNzk5MDE3Zjc1ZmJiZmFiZTIwMGY#/",
          "position": "right"
        }
      ],
      "hideOnScroll": false
    },
    "footer": {
      "links": [
        {
          "title": "Blackspace",
          "items": [
            {
              "label": "Documentation",
              "to": "/docs/introduction/"
            },
            {
              "label": "Roadmap",
              "href": "https://github.com/BlackspaceInc//projects/3"
            }
          ]
        },
        {
          "title": "Community",
          "items": [
            {
              "label": "Slack",
              "href": "https://blackspace.slack.com/join/shared_invite/enQtNzk4Nzg4Mjc2MTE2LTEzZThjMzliMjUzMTBmYzVjYWNmM2UyNWJmNDdkMDYyZmE0ZDliZTQxN2EzNzk5MDE3Zjc1ZmJiZmFiZTIwMGY#/"
            },
            {
              "label": "Twitter",
              "href": "https://twitter.com/blackspace"
            }
          ]
        },
        {
          "title": "More",
          "items": [
            {
              "label": "GitHub",
              "href": "https://github.com/BlackspaceInc/"
            }
          ]
        }
      ],
      "style": "light"
    }
  },
  "presets": [
    [
      "@docusaurus/preset-classic",
      {
        "docs": {
          "remarkPlugins": [
            null
          ],
          "sidebarPath": "/Users/yoanyomba/go/src/github.com/BlackspaceInc/BlackspacePlatform/docs/sidebars.js"
        },
        "blog": {
          "remarkPlugins": [
            null
          ],
          "feedOptions": {
            "type": "all",
            "copyright": "Copyright © 2020 Blackspace Platform"
          },
          "showReadingTime": true
        },
        "sitemap": {
          "cacheTime": 600000,
          "changefreq": "daily",
          "priority": 0.7
        },
        "theme": {
          "customCss": "/Users/yoanyomba/go/src/github.com/BlackspaceInc/BlackspacePlatform/docs/src/css/global.css"
        }
      }
    ]
  ],
  "ssrTemplate": "\n<!DOCTYPE html>\n<html lang=\"en\" >\n  <head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=3.0\">\n    <meta httpEquiv=\"x-ua-compatible\" content=\"ie=edge\" />\n    <meta property=\"og:type\" content=\"website\" />\n    <meta name=\"author\" content=\"Blackspace\">\n    <meta name=\"twitter:card\" content=\"summary_large_image\" />\n    <meta name=\"twitter:site\" content=\"@questdb\" />\n    <meta name=\"generator\" content=\"Docusaurus v<%= it.version %>\">\n    <link href=\"https://www.googletagmanager.com\" rel=\"dns-prefetch\" />\n    <link href=\"https://www.google-analytics.com\" rel=\"dns-prefetch\" />\n    <link rel=\"shortcut icon\" href=\"https://blackspaceInc.io/img/favicon.png\" />\n    <link rel=\"apple-touch-icon\" sizes=\"48x48\" href=\"/img/icons/icon-48x48.png\" />\r\n<link rel=\"apple-touch-icon\" sizes=\"72x72\" href=\"/img/icons/icon-72x72.png\" />\r\n<link rel=\"apple-touch-icon\" sizes=\"96x96\" href=\"/img/icons/icon-96x96.png\" />\r\n<link rel=\"apple-touch-icon\" sizes=\"144x144\" href=\"/img/icons/icon-144x144.png\" />\r\n<link rel=\"apple-touch-icon\" sizes=\"192x192\" href=\"/img/icons/icon-192x192.png\" />\r\n<link rel=\"apple-touch-icon\" sizes=\"256x256\" href=\"/img/icons/icon-256x256.png\" />\r\n<link rel=\"apple-touch-icon\" sizes=\"384x384\" href=\"/img/icons/icon-384x384.png\" />\r\n<link rel=\"apple-touch-icon\" sizes=\"512x512\" href=\"/img/icons/icon-512x512.png\" />\n    <link rel=\"sitemap\" type=\"application/xml\" href=\"/sitemap.xml\" />\n    <%~ it.headTags %>\n    <% it.metaAttributes.forEach((metaAttribute) => { %>\n      <%~ metaAttribute %>\n    <% }); %>\n    <% it.stylesheets.forEach((stylesheet) => { %>\n      <link rel=\"stylesheet\" type=\"text/css\" href=\"<%= it.baseUrl %><%= stylesheet %>\" />\n    <% }); %>\n  </head>\n  <body <%~ it.bodyAttributes %> itemscope=\"\" itemtype=\"http://schema.org/Organization\">\n    <meta itemprop=\"name\" content=\"Ecommerce platform empowering minority owned businesses to achieve more\" />\n    <meta itemprop=\"description\" content=\"Blackspace is an ecommerce platform empowering minoriting owned businesses to achieve more\" />\n    <meta itemprop=\"url\" content=\"https://blackspaceInc.io\" />\n    <meta itemprop=\"logo\" content=\"https://blackspaceInc.io/img/favicon.png\" />\n    <meta itemprop=\"sameAs\" content=\"https://twitter.com/blackspace\" />\n    <meta itemprop=\"sameAs\" content=\"https://www.linkedin.com/company/blackspaceInc/\" />\n    <%~ it.preBodyTags %>\n    <div id=\"__docusaurus\">\n      <%~ it.appHtml %>\n    </div>\n    <% it.scripts.forEach((script) => { %>\n      <script type=\"text/javascript\" src=\"<%= it.baseUrl %><%= script %>\" defer></script>\n    <% }); %>\n    <%~ it.postBodyTags %>\n  </body>\n</html>\n",
  "onBrokenLinks": "throw",
  "onDuplicateRoutes": "warn",
  "themes": []
};