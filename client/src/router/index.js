import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import ExampleList from '@/components/examples/List'
import ExampleDetail from '@/components/examples/Detail'
import ExamplePost from '@/components/examples/Post'
import About from '@/components/general/About'
import Protected from '@/components/general/Protected'
import Signin from '@/components/general/Signin'
import NotFound from '@/components/general/404'
import store from '../store'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Hello',
      component: Hello
    },
    {
      path: '/examples',
      name: 'Example List',
      component: ExampleList
    },
    {
      path: '/examples/post',
      name: 'Example Post',
      component: ExamplePost
    },
    {
      path: '/examples/:id',
      name: 'Example Detail',
      component: ExampleDetail
    },
    {
      path: '/about',
      name: 'About',
      component: About
    },
    {
      path: '/protected',
      name: 'Protected',
      meta: { Auth: true },
      component: Protected
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin
    },
    {
      path: '/404',
      component: NotFound
    },
    {
      path: '*',
      redirect: '/404'
    }
  ]
})

// this function looks for an item in the route oject to determine that a route is auth protected
// meta: { Auth: true }
router.beforeEach((to, from, next) => {
  // look for the meta Auth and check it (designated auth-guarded routes)
  if (to.matched.some(record => record.meta.Auth)) {
    if (!store.state.isLoggedIn) {
      next({
        path: '/signin'
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
