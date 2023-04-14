<template>
  <section class="content-main-w3" id="home">
    <SearchPage />
    <div class="container-fluid gallery-lightbox my-2">
      <div class="row main-box">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 p-0 u-may-like">Search Result</div>

        <GameAdvance :GameList="ds.games" />

        <div class="empty" v-show="ds.games.length === 0">Query result is empty</div>
      </div>
    </div>

    <DisplayAds AdSlot="search-horizontal" />

  </section>
</template>
<script>
import { defineComponent, onMounted, reactive } from "vue"
import { searchGames } from "@/utils/apis"
import { useRoute } from "vue-router"
import GameAdvance from "@/components/GameAdvance/index.vue"
import SearchPage from "@/components/Search/index.vue"
import DisplayAds from "@/components/DisplayAds/index.vue"

export default defineComponent({
  components: { GameAdvance, SearchPage, DisplayAds },
  setup() {
    const route = useRoute()
    const ds = reactive({
      games: [],
    })

    onMounted(() => {
      if (route.query.k.trim() !== "") {
        fetchGames(route.query.k.trim())
      }
    })
    const fetchGames = (k) => {
      searchGames({ k })
        .then((res) => {
          if (Array.isArray(res.data)) {
            ds.games = res.data
          }
        })
        .catch((e) => console.log(e))
    }
    return {
      ds,
    }
  },
})
</script>

<style lang="scss" scoped>
.main-box {
  .empty {
    text-align: center;
    color: #606266;
    height: 180px;
    line-height: 180px;
    color: var(--danger);
    text-shadow: 0 1px 2px var(--danger);
  }
}
</style>