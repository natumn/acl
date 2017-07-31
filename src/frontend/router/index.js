import Vue from 'vue'
import Router from 'vue-router'
import classlist from '@/components/classlist'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'classlist',
      component: classlist
    }
  ]
})
