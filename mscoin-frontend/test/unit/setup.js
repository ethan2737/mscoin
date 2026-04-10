let VueLib = null

try {
  VueLib = require('vue')
} catch (error) {
  VueLib = null
}

const Vue = VueLib && (VueLib.default || VueLib)

if (Vue && Vue.config) {
  Vue.config.productionTip = false
}
