import { createRouter, createWebHistory } from "vue-router"

const Layout = () => import("@/layout/index.vue")
export const routes = [
  {
    path: "/play-m",
    name: "PlayAd",
    meta: { title: "PlayAd" },
    component: () => import("@v/ad.vue"),
  },
  {
    path: "/play",
    name: "Play",
    meta: { title: "Play" },
    component: () => import("@v/play.vue"),
  },
  {
    path: "/",
    component: Layout,
    children: [
      {
        path: "",
        name: "Home",
        meta: { title: "Home" },
        component: () => import("@v/index.vue"),
      },
      {
        path: "about",
        name: "About",
        meta: { title: "About" },
        component: () => import("@v/about.vue"),
      },
      {
        path: "cookies",
        name: "Cookie",
        meta: { title: "Cookie" },
        component: () => import("@v/cookies.vue"),
      },
      {
        path: "privacy",
        name: "Privacy",
        meta: { title: "Privacy" },
        component: () => import("@v/privacy.vue"),
      },
      {
        path: "copyright",
        name: "Copyright",
        meta: { title: "Copyright" },
        component: () => import("@v/copyright.vue"),
      },
      {
        path: "contact",
        name: "Contact",
        meta: { title: "Contact" },
        component: () => import("@v/contact.vue"),
      },
      {
        path: "game",
        name: "GameInfo",
        meta: { title: "GameInfo" },
        component: () => import("@v/game.vue"),
      },
      {
        path: "search",
        name: "Search",
        meta: { title: "Search" },
        component: () => import("@v/search.vue"),
      },
      {
        path: "more",
        name: "More",
        meta: { title: "More" },
        component: () => import("@v/more.vue"),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})
export default router
