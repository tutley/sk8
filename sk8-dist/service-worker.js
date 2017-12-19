importScripts('workbox-sw.prod.v2.1.2.js');

/**
 * DO NOT EDIT THE FILE MANIFEST ENTRY
 *
 * The method precache() does the following:
 * 1. Cache URLs in the manifest to a local cache.
 * 2. When a network request is made for any of these URLs the response
 *    will ALWAYS comes from the cache, NEVER the network.
 * 3. When the service worker changes ONLY assets with a revision change are
 *    updated, old cache entries are left as is.
 *
 * By changing the file manifest manually, your users may end up not receiving
 * new versions of files because the revision hasn't changed.
 *
 * Please use workbox-build or some other tool / approach to generate the file
 * manifest which accounts for changes to local files and update the revision
 * accordingly.
 */
const fileManifest = [
  {
    "url": "/index.html",
    "revision": "bfddc2dc3eeed87377725eb353b68116"
  },
  {
    "url": "/static/css/app.9f1a7f721c010d86bd770bcfd18c0f20.css",
    "revision": "55f6d715e57a9ae3b61d34e2e4bc702a"
  },
  {
    "url": "/static/img/icons/android-chrome-192x192.png",
    "revision": "fae5b57fe035a1900bfaf73541de3965"
  },
  {
    "url": "/static/img/icons/android-chrome-512x512.png",
    "revision": "6419055aedf4f0fdcf1cf313ca0d530a"
  },
  {
    "url": "/static/img/icons/apple-touch-icon.png",
    "revision": "d5b8030c539af174272a8c12f8dff600"
  },
  {
    "url": "/static/img/icons/favicon-16x16.png",
    "revision": "b6be8968ea62b971281a04ab2f81746a"
  },
  {
    "url": "/static/img/icons/favicon-32x32.png",
    "revision": "e9db6d27ac87370fb5c3102c8b2560ba"
  },
  {
    "url": "/static/img/icons/icon-128x128.png",
    "revision": "507c16595c62217f9c9fa7500c600e3f"
  },
  {
    "url": "/static/img/icons/icon-144x144.png",
    "revision": "c9c5eb02a4c644550e16457a162f4e33"
  },
  {
    "url": "/static/img/icons/icon-152x152.png",
    "revision": "e7389291cd91bfbe574e4f1c36bdb5e0"
  },
  {
    "url": "/static/img/icons/icon-192x192.png",
    "revision": "06753edd96d87bb115c45d2fa3a77257"
  },
  {
    "url": "/static/img/icons/icon-384x384.png",
    "revision": "baa98672cee488fccef4658da3aa4e70"
  },
  {
    "url": "/static/img/icons/icon-512x512.png",
    "revision": "a4be890a52342369ebe1fd1c6f44ac1e"
  },
  {
    "url": "/static/img/icons/icon-72x72.png",
    "revision": "f724f0da0741c76bf20d45e3fcfbbf37"
  },
  {
    "url": "/static/img/icons/icon-96x96.png",
    "revision": "2e3ca42312b84cd9d2598c7ce0f25daa"
  },
  {
    "url": "/static/img/icons/mstile-150x150.png",
    "revision": "56e6ed072e8737c918a560e713dfcf8d"
  },
  {
    "url": "/static/js/app.b7013da65233808880f1.js",
    "revision": "80c3dee64ba9ea27335350654b84c7ee"
  },
  {
    "url": "/static/js/manifest.0780dcce1fd3d16f00d7.js",
    "revision": "4c1d407824f7b00bf4f2ac3966fb1ca4"
  },
  {
    "url": "/static/js/vendor.2a3369f3e3014c90bfa0.js",
    "revision": "733cf152e15879f8d9d6ed5422b76b80"
  }
];

const workboxSW = new self.WorkboxSW();
workboxSW.precache(fileManifest);
workboxSW.router.registerNavigationRoute("/");workboxSW.router.registerRoute(/^https:\/\/i.imgur.com\/..*/, workboxSW.strategies.cacheFirst({
  "cacheName": "imgurfiles"
}), 'GET');
workboxSW.router.registerRoute(/^https:\/\/fonts.gstatic.com\/..*/, workboxSW.strategies.cacheFirst({
  "cacheName": "fonts"
}), 'GET');
