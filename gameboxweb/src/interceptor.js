import router from "./router"
import { adStore } from "./store/ad"

router.beforeEach((to) => {
  // 路由已装置，pinia 也将装置
  const store = adStore()
  // store.getAdConfig()
})
