import { defineStore } from "pinia"
import { adConfig } from "../utils/apis"

export const adStore = defineStore("ad", {
  state: () => {
    return {
      AdType: "adsense",
      ShowCustomAd: true, // 是否显示定制广告位的广告
      fetched: false,
    }
  },
  getters: {
    showAdType(state) {
      return state.AdType
    },
    showCustomAd(state) {
      return state.ShowCustomAd
    },
  },
  actions: {
    async getAdConfig() {
      const config = await adConfig()
      console.log(config)
      if (config.code === 0) {
        this.$patch({
          AdType: "",
          ShowCustomAd: true,
          fetched: true,
        })
      }
    },
  },
  persist: {
    enabled: true,
  },
})
