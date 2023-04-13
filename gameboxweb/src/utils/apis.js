import req from "./request"

export function adConfig() {
  return req({
    url: "/ad-config",
    method: "get",
  })
}

export function gameList() {
  return req({
    url: "/games",
    method: "get",
  })
}

export function getGameInfo(gid) {
  return req({
    url: "/game/" + gid,
    method: "get",
  })
}

export function searchGames(params) {
  return req({
    url: "/search",
    method: "get",
    params,
  })
}

export function topicGames(params) {
  return req({
    url: "/topic",
    method: "get",
    params,
  })
}
