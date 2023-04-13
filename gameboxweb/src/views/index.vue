<template>
  <section class="content-main-w3" id="home">
    <SearchPage />

    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">
          <span class="title">Hot Games</span>
          <span class="action" @click="goMore('Hot')">MORE ></span>
        </div>
        <div class="col-lg-1 col-md-2 col-sm-3 col-4 p-0 snap-img" v-for="item, idx in ds.games.Hot" :key="idx">
          <GameBlock :GameInfo="item" />
        </div>
      </div>
    </div>

    <DisplayAds AdSlot="home-horizontal" />

    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">
          <span class="title">New Games</span>
          <span class="action" @click="goMore('New')">MORE ></span>
        </div>
        <div class="col-lg-1 col-md-2 col-sm-3 col-4 p-0 snap-img" v-for="item, idx in ds.games.New" :key="idx">
          <GameBlock :GameInfo="item" />
        </div>
      </div>
    </div>

    <DisplayAds AdSlot="home-horizontal-1" />

    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">
          <span class="title">All Games</span>
          <span class="action" @click="goMore('All')">MORE ></span>
        </div>
        <div class="col-lg-4 col-md-4 col-sm-6 col-12 p-0 snap-img" v-for="item, idx in ds.games.All" :key="idx">
          <GameAdvance :GameInfo="item" />
        </div>
      </div>
    </div>

    <DisplayAds AdSlot="home-horizontal-2" />
  </section>

</template>


<script>
import { defineComponent, reactive, onMounted } from "vue"
import { gameList } from "@/utils/apis"
import { useRouter } from "vue-router"
import GameAdvance from "@/components/GameAdvance/index.vue"
import GameBlock from "@/components/Game/index.vue"
import SearchPage from "@/components/Search/index.vue"
import DisplayAds from "@/components/DisplayAds/index.vue"
import { getStorage, setStorage } from "@/utils/cache"

export default defineComponent({
  components: { GameBlock, GameAdvance, SearchPage, DisplayAds },
  setup() {
    const ds = reactive({
      games: { Hot: [], New: [], All: [] },
      stars: [1, 2, 3, 4, 5],
    })
    const router = useRouter()
    onMounted(() => {
      // let d = getStorage("games")
      // if (d) {
      //   ds.games = d
      // } else {
      getGames()
      // }
    })
    const getGames = () => {
      gameList()
        .then((res) => {
          ds.games = res.data
          // setStorage("games", res.data)
        })
        .catch((e) => {
          console.log(e)
        })
    }
    const goMore = (t) => {
      router.push({ name: "More", query: { c: t } })
    }
    return {
      ds,
      goMore,
    }
  },
})
</script>
