import { createApp } from "vue"
import "./style.css"
import "./assets/styles/style.scss"
import * as bootstrap from "bootstrap"
import App from "./App.vue"
import VueLazyLoad from "vue3-lazyload"
import router from "./router"
import pinia from "./store"
import "./interceptor"
const app = createApp(App)

app.use(pinia)
app.use(VueLazyLoad)
app.use(router)

app.mount("#app")
