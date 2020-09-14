const visit = require("unist-util-visit")
const ssrTemplate = require("./src/internals/ssr.template")

const customFields = {
  artifactHubUrl: "https://artifacthub.io/packages/helm/BlackspaceInc/BlackspacePlatform",
  copyright: `Copyright © ${new Date().getFullYear()} Blackspace Platform`,
  description:
    "Blackspace is an ecommerce platform empowering minoriting owned businesses to achieve more",
  dockerUrl: "https://hub.docker.com/repository/docker/blackspaceinc/frontend-service",
  domain: "blackspaceinc.github.io",
  githubUrl: "https://github.com/BlackspaceInc/",
  helmVersion: "0.2.4",
  linkedInUrl: "https://www.linkedin.com/company/blackspaceInc/",
  oneLiner: "Ecommerce platform empowering minority owned businesses to achieve more",
  slackUrl:
    "https://blackspace.slack.com/join/shared_invite/enQtNzk4Nzg4Mjc2MTE2LTEzZThjMzliMjUzMTBmYzVjYWNmM2UyNWJmNDdkMDYyZmE0ZDliZTQxN2EzNzk5MDE3Zjc1ZmJiZmFiZTIwMGY#/",
  twitterUrl: "https://twitter.com/blackspace",
  version: "5.0.3",
}

function variable() {
  const RE_VAR = /{@([\w-_]+)@}/g
  const getVariable = (full, partial) =>
    partial ? customFields[partial] : full

  function textVisitor(node) {
    node.value = node.value.replace(RE_VAR, getVariable)
  }

  function linkVisitor(node) {
    node.url = node.url.replace(RE_VAR, getVariable)

    if (node.title) {
      node.title = node.title.replace(RE_VAR, getVariable)
    }
  }

  function transformer(ast) {
    visit(ast, "text", textVisitor)
    visit(ast, "code", textVisitor)
    visit(ast, "link", linkVisitor)
  }

  return transformer
}

const config = {
  title: "Internal Documentation",
  tagline: "Blackspace is the fastest social ecommerce solution for minority owned business owners",
  url: `https://${customFields.domain}`,
  baseUrl: "/",
  onBrokenLinks: "warn",
  favicon: "/img/favicon.png",
  organizationName: "BlackspaceInc",
  projectName: "BlackspacePlatform",
  customFields,
  plugins: [
    require.resolve("./plugins/fetch-release"),
    require.resolve("./plugins/lint"),
    require.resolve("./plugins/manifest"),
    [
      "@docusaurus/plugin-ideal-image",
      {
        quality: 100,
        steps: 2, // the max number of images generated between min and max (inclusive)
      },
    ],
    /** 
    [
      "@docusaurus/plugin-pwa",
      {
        pwaHead: [
          {
            tagName: "link",
            rel: "icon",
            href: "/img/favicon.png",
          },
          {
            tagName: "link",
            rel: "manifest",
            href: "/manifest.webmanifest",
          },
          {
            tagName: "meta",
            name: "theme-color",
            content: "#d14671",
          },
          {
            tagName: "meta",
            name: "apple-mobile-web-app-capable",
            content: "yes",
          },
          {
            tagName: "meta",
            name: "apple-mobile-web-app-status-bar-style",
            content: "#21222c",
          },
          {
            tagName: "link",
            rel: "apple-touch-icon",
            href: "/img/favicon.png",
          },
          {
            tagName: "link",
            rel: "mask-icon",
            href: "/img/favicon.png",
            content: "#fff",
          },
          {
            tagName: "meta",
            name: "msapplication-TileImage",
            content: "/img/favicon.png",
          },
          {
            tagName: "meta",
            name: "msapplication-TileColor",
            content: "#21222c",
          },
        ],
      },
    ],*/
  ],
  themeConfig: {
    announcementBar: {
      id: "github-star",
    },
    colorMode: {
      defaultMode: 'light',
      disableSwitch: false,
      respectPrefersColorScheme: true,
    },
    announcementBar: {
      id: 'supportus',
      backgroundColor: '#1064d3',
      textColor: 'white',
      content: '⭐️ If you like Blackspace, give it a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/BlackspaceInc/BlackspacePlatform">GitHub</a>! ⭐️',
    },
    image: "/img/og.png",
    gtag: {
      trackingID: "GTM-PVR7M2G",
      anonymizeIP: true,
    },
    prism: {
      defaultLanguage: "blackspace",
      theme: require("./src/internals/prism-dracula"),
    },
    algolia: {
      apiKey: "b2a69b4869a2a85284a82fb57519dcda",
      indexName: "blackspace",
    },
    navbar: {
      title: " ",
      logo: {
        alt: "Blackspace",
        src: "/img/navbar/logo.ico",
      },
      items: [
        {
          label: "Getting Started",
          position: "left",
          items: [
            {
              label: "Docker",
              to: "/docs/guide/docker/",
            },
            {
              label: "Kubernetes",
              to: "/docs/guide/kubernetes/",
            }
          ],
        },
        {
          label: "Documentation",
          position: "left",
          to: "/docs/introduction/",
          activeBasePath: "docs",
        },
        /** 
        {
          label: "Blog",
          to: "/blog/",
          position: "left",
        },
        */
        {
          label: "GitHub",
          className: "navbar__item--github",
          href: customFields.githubUrl,
          position: "right",
        },
        {
          label: "Join Slack",
          className: "navbar__item--slack",
          href: customFields.slackUrl,
          position: "right",
        },
      ],
    },
    footer: {
      links: [
        {
          title: "Blackspace",
          items: [
            {
              label: "Documentation",
              to: "/docs/introduction/",
            },
            {
              label: "Roadmap",
              href: `${customFields.githubUrl}/projects/3`,
            },
          ],
        },
        {
          title: "Community",
          items: [
            {
              label: "Slack",
              href: customFields.slackUrl,
            },
            {
              label: "Twitter",
              href: customFields.twitterUrl,
            },
          ],
        },
        {
          title: "More",
          items: [
            /*
            {
              label: "Blog",
              to: "/blog/",
            },
            */
            {
              label: "GitHub",
              href: customFields.githubUrl,
            },
          ],
        },
      ],
    },
  },
  presets: [
    [
      "@docusaurus/preset-classic",
      {
        docs: {
          remarkPlugins: [variable],
          sidebarPath: require.resolve("./sidebars.js"),
        },
        blog: {
          remarkPlugins: [variable],
          feedOptions: {
            type: "all",
            copyright: customFields.copyright,
          },
          showReadingTime: true,
        },
        sitemap: {
          cacheTime: 600 * 1000, // 600 sec - cache purge period
          changefreq: "daily",
          priority: 0.7,
        },
        theme: {
          customCss: require.resolve("./src/css/global.css"),
        },
      },
    ],
  ],
}

module.exports = {
  ...config,
  ssrTemplate: ssrTemplate(config),
}
