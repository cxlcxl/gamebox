<template>
  <div class="full-page">
    <div class="ad-position">
      <!-- TODO 广告位 -->全屏广告
    </div>

    <!--game-play-->
    <div class="full-ad-page">
      <div class="game-play">
        <a :class="{'gray-btn': ds.wait > 0}" @click="play">PLAY<span class="ad-tip" v-show="ds.wait > 0">Ad close in 00:0{{ ds.wait }}</span></a>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, onMounted, reactive, ref } from "vue"
import { useRouter } from "vue-router"
import { useRoute } from "vue-router"

export default defineComponent({
  setup() {
    const ds = reactive({
      wait: 6,
    })
    let interval = ref({})
    const route = useRoute()
    const router = useRouter()
    onMounted(() => {
      if (!/^[A-Z0-9]{10,15}$/.test(route.query.gid)) {
        router.push({ name: "Home" })
        return
      }
      timeRun()
    })

    const timeRun = () => {
      interval = setInterval(() => {
        ds.wait--
        if (ds.wait <= 0) {
          finish()
        }
      }, 1000)
    }

    const finish = () => {
      clearInterval(interval)
    }

    const play = () => {
      if (ds.wait > 0) {
        return
      }
      router.push({ name: "Play", query: { gid: route.query.gid } })
    }
    return {
      ds,
      play,
    }
  },
})
</script>

<style lang="scss" scoped>
.full-page {
  width: 100%;
  height: 100%;
  background: #666666;
}

.full-ad-page .game-play {
  position: fixed;
  bottom: 30px;
  left: 0;
  width: 100%;
  text-align: center;

  a {
    display: inline-block;
    text-align: center;
    height: 2em;
    line-height: 2em;
    font-size: 1.2em;
    color: #ffffff;
    font-weight: 500;
    background: linear-gradient(to bottom right, #6f42c1, #6610f2, #007bff);
    cursor: pointer;
    border-radius: 10px;
    padding: 0 20px;
  }
}

.ad-tip {
  margin-left: 10px;
  color: #eee;
  font-size: 0.8em !important;
  font-weight: 400 !important;
}

.gray-btn {
  -webkit-filter: grayscale(100%);
  -moz-filter: grayscale(100%);
  -ms-filter: grayscale(100%);
  -o-filter: grayscale(100%);
  filter: grayscale(100%);
}
</style>