<template>
  <SearchPage />

  <section class="content-main-w3" id="home">
    <div class="logo-wthree text-center">
      <a class="navbar-brand" href="#">
        {{ds.game.name}}
      </a>
    </div>

    <div class="container-fluid gallery-lightbox my-2">
      <div class="row mb-3 align-center page-width">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0">
          <div class="game-detail-ad-bg">
            <div class="game-detail-icon mb-2">
              <img :src="ds.game.icon">
            </div>
            <div class="game-desc-title p-2">
              <h6>DESCRIPTION</h6>
            </div>
            <div class="game-desc p-2">{{ds.game.desc ? ds.game.desc : ds.welcome}}</div>
            <div class="game-play" @click="play(ds.game.game_id)"><a>PLAY NOW</a></div>
          </div>

          <DisplayAds AdSlot="game-horizontal" />
          <div class="row m-0 may-like-box">
            <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">
              <span class="title">May Like</span>
              <span class="action"></span>
            </div>
            <div class="col-lg-4 col-md-6 col-sm-6 col-12 p-0 snap-img" v-for="item, idx in ds.games" :key="idx">
              <GameAdvance :GameInfo="item" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { defineComponent, onMounted, reactive } from "vue"
import { useRoute, useRouter } from "vue-router"
import GameAdvance from "@/components/GameAdvance/index.vue"
import SearchPage from "@/components/Search/index.vue"
import DisplayAds from "@/components/DisplayAds/index.vue"
import { getGameInfo, topicGames } from "@/utils/apis"
import Setting from "@/setting"

export default defineComponent({
  components: { GameAdvance, SearchPage, DisplayAds },
  setup() {
    const route = useRoute()
    const router = useRouter()

    const ds = reactive({
      game: {},
      games: [],
      welcome: Setting.CompanyDesc,
    })

    onMounted(() => {
      gameInfo(route.query.gid)
      likeGameInfo()
    })
    const gameInfo = (gid) => {
      if (!/^[a-zA-Z0-9]{10,15}$/.test(gid)) {
        router.push({ name: "Home" })
        return
      }
      getGameInfo(gid)
        .then((res) => {
          ds.game = res.data
        })
        .catch((e) => {
          console.log(e)
        })
    }
    const likeGameInfo = () => {
      topicGames({ topic: "Like" })
        .then((res) => {
          ds.games = res.data
        })
        .catch((e) => {
          console.log(e)
        })
    }
    const play = (gid) => {
      router.push({ name: "Play", query: { gid } })
    }
    const goMore = () => {
      router.push({ name: "More", query: { c: "Like" } })
    }

    return {
      ds,
      play,
      goMore,
    }
  },
})
</script>
