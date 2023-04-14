<template>
  <section class="content-main-w3" id="home">
    <SearchPage />
    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">{{ds.c}} Games</div>

        <GameAdvance :GameList="ds.games" />
      </div>
    </div>
    <DisplayAds AdSlot="home-horizontal" />
  </section>
</template>

<script>
import { defineComponent, onMounted, reactive } from "vue"
import SearchPage from "@/components/Search/index.vue"
import GameAdvance from "@/components/GameAdvance/index.vue"
import DisplayAds from "@/components/DisplayAds/index.vue"
import { useRoute, useRouter } from "vue-router"
import { topicGames } from "@/utils/apis"
import { getStorage } from "@/utils/cache"

export default defineComponent({
  components: { GameAdvance, SearchPage, DisplayAds },
  setup() {
    const route = useRoute()
    const router = useRouter()
    const ds = reactive({
      c: "",
      games: [],
    })

    onMounted(() => {
      if (!/^[A-Z]{1}[a-z]{1,20}$/.test(route.query.c)) {
        router.replace({ name: "Home" })
      }
      ds.c = route.query.c
      if (ds.c === "Played") {
        fetchPlayedGames()
      } else {
        fetchGames(ds.c)
      }
    })
    const fetchGames = (topic) => {
      topicGames({ topic })
        .then((res) => {
          if (Array.isArray(res.data)) {
            ds.games = res.data
          }
        })
        .catch((e) => console.log(e))
    }
    const fetchPlayedGames = () => {
      const played = getStorage("played")
      if (Array.isArray(played)) {
        ds.games = played
      }
    }
    return {
      ds,
    }
  },
})
</script>
