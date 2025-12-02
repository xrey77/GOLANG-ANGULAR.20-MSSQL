
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
    "route": "/**"
  }
],
  entryPointToBrowserMapping: undefined,
  assets: {
    'index.csr.html': {size: 9776, hash: '971187a933dcf223c9616c6b4aee9c3ca6b0dd5921d57412d70605b11cae7897', text: () => import('./assets-chunks/index_csr_html.mjs').then(m => m.default)},
    'index.server.html': {size: 1023, hash: '7198a568ad19ae302077aa2fed60c60c9947495cbfd10d92a8822adbfa440e14', text: () => import('./assets-chunks/index_server_html.mjs').then(m => m.default)},
    'index.html': {size: 53972, hash: '5c1c8f24c9542b8244a43d668fe882053a3777a71c43b3e6be400deb18f69f2e', text: () => import('./assets-chunks/index_html.mjs').then(m => m.default)},
    'static/index.html': {size: 41584, hash: 'cec982fd68c502ad59ba2c05288f1f8fd4e6a4b478da7d54718867542fdc7898', text: () => import('./assets-chunks/static_index_html.mjs').then(m => m.default)},
    'contact/index.html': {size: 46912, hash: '54559847ed732605722287474e0d66dc9ec1a830079c638e5f406ebe827df02c', text: () => import('./assets-chunks/contact_index_html.mjs').then(m => m.default)},
    'productsearch/index.html': {size: 46986, hash: 'b6bb51690abb55dc6d4ca9844f9d891cf649284b41cd7d2de33b46f56bf4039b', text: () => import('./assets-chunks/productsearch_index_html.mjs').then(m => m.default)},
    'productcatalog/index.html': {size: 47970, hash: '496248c92173c62f431aaec5f530d1a7c9e83d513a9331b1c84886a7f9a23984', text: () => import('./assets-chunks/productcatalog_index_html.mjs').then(m => m.default)},
    'about/index.html': {size: 47467, hash: '647543861d02e7d503c3332f4428228be67c98429409089fbce49e9ecf42add4', text: () => import('./assets-chunks/about_index_html.mjs').then(m => m.default)},
    'productlist/index.html': {size: 52899, hash: '13c7ae4f8583ce5f2f0f49f06c2d3fd97f74a4df0fdd5b44ee4a1e4a6fdcb67d', text: () => import('./assets-chunks/productlist_index_html.mjs').then(m => m.default)},
    'profile/index.html': {size: 61406, hash: '09f692a3c6d718c18b1eb94c895194c5dad12390f6823e024775358cc13b5d63', text: () => import('./assets-chunks/profile_index_html.mjs').then(m => m.default)},
    'styles.css': {size: 396704, hash: 'd+zBPSfIKoU', text: () => import('./assets-chunks/styles_css.mjs').then(m => m.default)}
  },
};
