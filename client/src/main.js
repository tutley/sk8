import '../node_modules/vuetify/src/stylus/app.styl'
import Vue from 'vue'
import {
  Vuetify,
  VApp,
  VNavigationDrawer,
  VList,
  VBtn,
  VIcon,
  VCard,
  VToolbar,
  VFooter,
  VDivider,
  VForm,
  VProgressCircular,
  VProgressLinear,
  VSubheader,
  VTextField,
  VAlert,
  VGrid,
  VDialog,
  transitions
} from 'vuetify'
import App from './App'
import router from './router'
import store from './store/index'

Vue.config.productionTip = false

Vue.use(Vuetify, {
  components: {
    VApp,
    VNavigationDrawer,
    VList,
    VBtn,
    VIcon,
    VCard,
    VToolbar,
    VFooter,
    VDivider,
    VForm,
    VProgressCircular,
    VProgressLinear,
    VSubheader,
    VTextField,
    VAlert,
    VGrid,
    VDialog,
    transitions
  },
  directives: {
    Touch
  },
  theme: {
    primary: '#3ab843',
    secondary: '#435466',
    accent: '#24e132',
    error: '#FF5252',
    info: '#73ea7b',
    success: '#4CAF50',
    warning: '#FFC107'
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: { App },
  created() {
    if (localStorage.getItem('token')) {
      this.$store.dispatch('setIsLoggedIn', true)
    }
  }
})
