import { createPinia } from "pinia"
import piniaPluginpersist from "pinia-plugin-persist"

const pinia = createPinia()
pinia.use(piniaPluginpersist)
export default pinia
