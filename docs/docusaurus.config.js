const publish_blackspace_private = process.env.PUBLISH === 'blackspace_private' 
const publish_blackspace_public  = process.env.PUBLISH === 'blackspace_public' 
const is_for_webapp = !publish_blackspace_private && !publish_blackspace_public

// See https://docusaurus.io/docs/site-config for all the possible site configuration options.
const logo = {
  alt: 'Blackspace Logo',
  // https://commons.wikimedia.org/wiki/File:Twemoji_1f429.svg
  src: 'img/twemoji_poodle.svg',
}

var config = {
  title: 'Blackspace Platform',
  tagline: "Empowering minority owned businesses to better serve their clients <br/>and achieve more.",
  // You may provide arbitrary config keys to be used as needed by your
  // template. For example, if you need your repo's URL...

  // scripts: [
  // Add custom scripts here that would be placed in <script> tags.
  //'https://buttons.github.io/buttons.js'
  // ],

  // https://realfavicongenerator.net/
  favicon: 'img/favicon/favicon-32x32.png',
  customFields: {
    description:
      "Empowering minority owned businesses to better serve their clients and achieve more.",
  },
  onBrokenLinks: 'log',
  themeConfig: {
    image: "img/share.jpg",
    announcementBar: {
      id: 'supportus',
      backgroundColor: '#1064d3',
      textColor: 'white',
      content: '⭐️ If you like Blackspace, give it a star on <a target="_blank" rel="noopener noreferrer" href="https://github.com/BlackspaceInc/BlackspacePlatform">GitHub</a>! ⭐️',
    },
    prism: {
      additionalLanguages: ['nginx'],
      theme: require('prism-react-renderer/themes/github'),
      darkTheme: require('prism-react-renderer/themes/dracula'),
    },
    footer: {
      logo,
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Introduction',
              to: 'docs/introduction',
            },
            {
              label: 'Technology',
              to: 'docs/technology',
            },
          ],
        },
        {
          title: 'Products',
          items: [
            {
              label: 'Product Overview',
              to: 'docs/technology/products/overview',
            },
            {
              label: 'Marketplace',
              to: 'docs/technology/products/marketplace',
            },
            {
              label: ' Business',
              to: 'docs/technology/products/business',
            },
            {
              label: ' Analytics',
              to: 'docs/technology/products/analytics',
            },
            {
              label: ' Makers',
              to: 'docs/technology/products/makers',
            },
            {
              label: ' 3rd Party Integrations',
              to: 'docs/technology/products/integrations',
            },
            {
              label: ' Ads',
              to: 'docs/technology/products/ads',
            },
            {
              label: ' AI & Research',
              to: 'docs/technology/products/research',
            },
            {
              label: ' VC',
              to: 'docs/technology/products/vc',
            },
          ],
        },
        {
          title: 'Social',
          items: [
            {
              label: 'Blog',
              href: 'https://the-guild.dev/blog',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/kamilkisiela/graphql-inspector',
            },
            {
              label: 'Twitter',
              href: 'https://twitter.com/kamilkisiela',
            },
            {
              label: 'LinkedIn',
              href: 'https://www.linkedin.com/company/the-guild-software',
            },
          ],
        },],
      copyright: "Made with ❤️ at Blackspace. Apache 2.0 License. Built with Docusaurus.",
    },
    navbar: {
      title: 'Blackspace Platform',
      logo,
      hideOnScroll: true,
      items: [
        {to: 'blog', label: 'Blog', position: 'left'}, // or position: 'right'
        {to: is_for_webapp ? '/introduction' : 'docs/introduction', label: 'Docs', position: 'left'},
        {
          href: 'https://github.com/BlackspaceInc/BlackspacePlatform',
          position: 'left',
          label: 'Source',
        },
        {
          href: 'https://github.com/BlackspaceInc/BlackspacePlatform',
          position: 'right',
          className: 'header-github-link',
          'aria-label': 'GitHub repository',
        },
      ],
    },
    // removes "active" color on parent sidebar categories :|
    sidebarCollapsible: true,
  },
};

config = {
  ...config,
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        theme: {
          customCss: [require.resolve('./src/css/custom.css'), require.resolve('remark-admonitions/styles/infima.css')],
        },
        blog: {
          path: 'blog/engineering',
          routeBasePath: 'blog',
          blogDescription: 'Blackspace Engineering',
          feedOptions: {
            type: 'all',
            copyright: `Copyright © ${new Date().getFullYear()} Blackspace Platform, Inc.`,
          },
          include: ['*.md', '*.mdx'],
          postsPerPage: 10,
          /**
           * Theme components used by the blog pages.
           */
          blogListComponent: '@theme/BlogListPage',
          blogPostComponent: '@theme/BlogPostPage',
          blogTagsListComponent: '@theme/BlogTagsListPage',
          blogTagsPostsComponent: '@theme/BlogTagsPostsPage',
          /**
           * Remark and Rehype plugins passed to MDX.
           */
          remarkPlugins: [
            /* require('remark-math') */
          ],
          rehypePlugins: [],
          /**
           * Custom Remark and Rehype plugins passed to MDX before
           * the default Docusaurus Remark and Rehype plugins.
           */
          beforeDefaultRemarkPlugins: [],
          beforeDefaultRehypePlugins: [],
          /**
           * Truncate marker, can be a regex or string.
           */
          truncateMarker: /<!--\s*(truncate)\s*-->/,
          /**
           * Show estimated reading time for the blog post.
           */
          showReadingTime: true,
        },
        docs: {
          path: 'docs',
          routeBasePath: 'docs',
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl: 'https://github.com/BlackspaceInc/BlackspacePlatform/edit/master/website/',
          admonitions: {},
          // Show documentation's last contributor's name.
          // enableUpdateBy: true,
          // Show documentation's last update time.
          // enableUpdateTime: true,        
        },
      },
    ],
  ],
}


if (publish_blackspace_public) {
  config.themeConfig.algolia = {
    apiKey: '7e47115263beea4eb52978a771750414',
    indexName: 'docs',
    algoliaOptions: {
      // facetFilters: [`version:${versions[0]}`],
    },
  }
} else {
  config.plugins = [require.resolve('docusaurus-lunr-search')]
}



if (is_for_webapp) {
  // build for the app at /docs
  config = {
    ...config,
    url: 'https://blackspaceinc.github.io',
    baseUrl: '/',
    organizationName: 'BlackspaceInc',
    projectName: 'BlackspacePlatform',
  }
  config.presets[0][1].docs.routeBasePath = '';
} else {
  if (publish_blackspace_private) {
    config = {
      ...config,
      url: 'https://blackspaceinc.github.io',
      baseUrl: '/',
      organizationName: 'BlackspaceInc',
      projectName: 'BlackspacePlatform',
    }  
  }
  if (publish_blackspace_public) {
    config = {
      ...config,
      url: 'https://blackspaceinc.github.io',
      baseUrl: '/',
      organizationName: 'BlackspaceInc',
      projectName: 'BlackspacePlatform',
    }
  }
}

module.exports = config;
