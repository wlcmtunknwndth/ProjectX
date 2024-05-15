import {createWebHistory, createRouter} from "vue-router";

import MainPage from "../pages/MainPage.vue"
import DetailPage from "../pages/DetailPage.vue"
import HomePage from "../pages/HomePage.vue";

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'home', component: HomePage },
    { path: '/events', name: 'main', component: MainPage },
    { path: '/events/:id', name: 'detail', component: DetailPage },
  ],
  scrollBehavior(to, from, savedPosition) {
    // always scroll to top
    return { top: 0 }
  },
})