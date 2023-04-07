<template>
  <div class="header_top_w3ls">
    <div id="navbarSupportedContent">
      <a href="/"><img src="/images/logo.png" class="gigame-logo" /></a>
    </div>
    <div style="flex: 1;"></div>
    <div class="game-search">
      <div class="search-box">
        <img class="icon-search" src="/images/search.png" />
        <input type="text" placeholder="Search Games" v-model="search.k" @blur="doSearch" />
      </div>
    </div>
  </div>

  <div class="search-container">
    <img class="icon-search" src="/images/search.png" />
    <input type="text" placeholder="Search Games" v-model="search.k" @blur="doSearch" />
  </div>
</template>

<script>
import { defineComponent, onMounted, reactive } from "vue"
import { useRoute, useRouter } from "vue-router"

export default defineComponent({
  setup() {
    const search = reactive({
      k: "",
    })
    const route = useRoute()
    const router = useRouter()

    onMounted(() => {
      if (route.query.k) {
        search.k = route.query.k.trim()
      }
    })

    const doSearch = () => {
      if (search.k.trim() !== "") {
        router.push({ name: "Search", query: { k: search.k } })
      }
    }
    return {
      search,
      doSearch,
    }
  },
})
</script>

