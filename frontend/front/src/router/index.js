import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/Login.vue'
import RegisterView from '../views/RegisterView.vue'
import PostView from '../views/PostView.vue'
import ProfileView from '../views/ProfileView.vue'
import UserPageView from '../views/UserPageView.vue'
import NewPostView from '../views/NewPostView.vue'
import SearchResultsView from '../views/SearchResultsView.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  },
  {
    path: '/new/post',
    name: 'newPost',
    component: NewPostView
  },
  {
    path: '/search/results/:query',
    name: 'searchResults',
    component: SearchResultsView
  },
  {
    path: '/post/:postId',
    name: 'post',
    component: PostView
  },
  {
    path: '/user/:userId',
    name: 'user',
    component: ProfileView
  },
  {
    path: '/myHome',
    name: 'myHome',
    component: UserPageView
  },

  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
