import Vue from 'vue'
import VueRouter from 'vue-router'
import Books from '../views/Books.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Books
  },
  {
    path: '/singleBook/:id',
    name: 'SingleBook',
    component: () => import( '../views/SingleBook.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import( '../views/Login.vue')
  },
  {
    path: '/signup',
    name: 'Signup',
    component: () => import( '../views/Signup.vue')
  }
]

const router = new VueRouter({
  routes,
  mode: 'history'
})

export default router
