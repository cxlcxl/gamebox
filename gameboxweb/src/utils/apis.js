import req from "./request"

export function gameList() {
  return req({
    url: "/games",
    method: "get",
  })
}

export function getGameInfo(params) {
  return req({
    url: "/game",
    method: "get",
    params,
  })
}

export function getGameDetail(gid) {
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
