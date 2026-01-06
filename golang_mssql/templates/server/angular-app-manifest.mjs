
export default {
  bootstrap: () => import('./main.server.mjs').then(m => m.default),
  inlineCriticalCss: true,
  baseHref: '/',
  locale: undefined,
  routes: [
  {
    "renderMode": 2,
    "route": "/"
  },
  {
    "renderMode": 2,
    "route": "/static"
  },
  {
    "renderMode": 2,
    "route": "/about"
  },
  {
    "renderMode": 2,
    "route": "/contact"
  },
  {
    "renderMode": 2,
    "route": "/profile"
  },
  {
    "renderMode": 2,
    "route": "/productlist"
  },
  {
    "renderMode": 2,
    "route": "/productcatalog"
  },
  {
    "renderMode": 2,
    "route": "/productsearch"
  },
  {
    "renderMode": 2,
    "route": "/pdfreport"
  },
  {
    "renderMode": 2,
    "route": "/saleschart"
  },
  {
    "renderMode": 2,
    "route": "/**"
  }
],
  entryPointToBrowserMapping: undefined,
  assets: {
    'index.csr.html': {size: 9778, hash: '8b02141deb72ab61b6f4f6c65affd57a5783c59e3a1a6b53f7e5736705ca6839', text: () => import('./assets-chunks/index_csr_html.mjs').then(m => m.default)},
    'index.server.html': {size: 1025, hash: '03c3af0ab58b4769ff1cf8ba9d2a886344ef5a3aa43a168722d8a9178feee076', text: () => import('./assets-chunks/index_server_html.mjs').then(m => m.default)},
    'index.html': {size: 54747, hash: 'cde75eff93561ae6fa3e60a55bfd24ee073c113019bd691f541943d2948066fc', text: () => import('./assets-chunks/index_html.mjs').then(m => m.default)},
    'static/index.html': {size: 42359, hash: 'e50e0bab741d46221a5ede327c53c0176f2a887be3135e597cbc5c6e2931f1cd', text: () => import('./assets-chunks/static_index_html.mjs').then(m => m.default)},
    'productsearch/index.html': {size: 47761, hash: '0cfaf8845dfda165d3c6d0a26fda58b6572c2a2fc767d35f1fb109d3b180b57b', text: () => import('./assets-chunks/productsearch_index_html.mjs').then(m => m.default)},
    'saleschart/index.html': {size: 41961, hash: 'c6226a5644f1db4e5f426d4c724504e63d7c4f8ce9dddb8c2f5f45c72afcc6d4', text: () => import('./assets-chunks/saleschart_index_html.mjs').then(m => m.default)},
    'contact/index.html': {size: 47687, hash: '5d18c6e92a117924aba44aa6fffb29a939a5e93218c34a95ea2f5a62ef97b556', text: () => import('./assets-chunks/contact_index_html.mjs').then(m => m.default)},
    'productcatalog/index.html': {size: 48745, hash: '4f5b1a7bf257db878396bc81379bf5a66197a07eb002d62a244ede711225e5b5', text: () => import('./assets-chunks/productcatalog_index_html.mjs').then(m => m.default)},
    'about/index.html': {size: 48242, hash: '6af6b66ce18ccb75246f2b4a12a54fb9ec2f9892b70607560fb2e8784172f51d', text: () => import('./assets-chunks/about_index_html.mjs').then(m => m.default)},
    'pdfreport/index.html': {size: 42140, hash: 'fe837bd04be475207eabe89ae94c709c77d7094f2e9c1458a475963a2b59121b', text: () => import('./assets-chunks/pdfreport_index_html.mjs').then(m => m.default)},
    'productlist/index.html': {size: 53674, hash: '046edbb1dbb95c4c8f0f840b03fc3b5beda5203ec04573bfc7b117805c1864d4', text: () => import('./assets-chunks/productlist_index_html.mjs').then(m => m.default)},
    'profile/index.html': {size: 62181, hash: '036d690e73f758c6855f673b7acdc1c45ba695d00a08d01e47d88179799fa4b6', text: () => import('./assets-chunks/profile_index_html.mjs').then(m => m.default)},
    'styles.css': {size: 396704, hash: 'd+zBPSfIKoU', text: () => import('./assets-chunks/styles_css.mjs').then(m => m.default)}
  },
};
