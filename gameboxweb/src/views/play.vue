<template>
  <div class="play-game-container">
    <div class="back-home" @click="backHome">
      <img src="/images/home.png" />
    </div>
    <div id="gameFrame">
      <div class="play-banner play-banner-left">
        <DisplayAds AdSlot="home-horizontal" />
      </div>

      <iframe v-show="ds.game.url !== ''" ref="game" allowfullscreen="true" webkitallowfullscreen="true" mozallowfullscreen="true" @load="loaded"
        :src="ds.game.url" scrolling="auto" frameborder="0" id="game-iframe">
      </iframe>

      <div class="play-banner play-banner-right">
        <DisplayAds AdSlot="home-horizontal-1" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, onMounted } from "vue"
import { useRoute, useRouter } from "vue-router"
import { getGameInfo } from "@/utils/apis"
import DisplayAds from "@/components/DisplayAds/index.vue"

const ds = reactive({
  more: true,
  game: {},
})
const route = useRoute()
const router = useRouter()

onMounted(() => {
  gameInfo(route.query.gid)
})
const gameInfo = (gid) => {
  if (!/^[a-zA-Z0-9]{10,15}$/.test(gid)) {
    router.push({ name: "Home" })
    return
  }
  getGameInfo(gid)
    .then((res) => {
      ds.game = res.data
      document.getElementById("app").style.height = "100%" // 修正 adsense 导致的高度问题
      document.getElementById("gameFrame").style.height = "100%" // 修正 adsense 导致的高度问题
    })
    .catch((e) => {
      console.log(e)
    })
}
const backHome = () => {
  router.replace({ name: "Home" })
}
const loaded = () => {}
</script>

<style scoped lang="scss">
.play-game-container {
  position: relative;
  width: 100%;
  height: 100%;
  background: grey;
  overflow-y: auto;
  overflow-x: hidden;

  #gameFrame {
    display: flex;
  }
  .back-home {
    width: 35px;
    height: 35px;
    position: fixed;
    background: rgba(255, 255, 255, 0.7);
    background-size: cover;
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px;
    top: 20%;
    left: 0;
    z-index: 9999;
    cursor: pointer;
    border: 1px solid #dedede;

    img {
      position: absolute;
      top: 3px;
      left: 3px;
      width: 25px;
      height: 25px;
    }
  }

  #game-iframe {
    position: relative;
    z-index: 9;
    width: 70%;
    height: 100%;
    vertical-align: middle;
  }

  .play-banner {
    width: 15%;
    overflow: hidden;
    text-align: center;
  }
}

@media (max-width: 1500px) {
  #game-iframe {
    width: 60% !important;
  }
  .play-game-container .play-banner {
    width: 20% !important;
  }
}
@media (max-width: 1100px) {
  #game-iframe {
    width: 70% !important;
  }
  .play-game-container .play-banner-left {
    display: none !important;
  }
  .play-game-container .play-banner-right {
    width: 30% !important;
  }
}
@media (max-width: 900px) {
  #game-iframe {
    width: 100% !important;
    height: 100% !important;
  }
  .play-game-container .play-banner {
    display: none;
  }
}
</style>