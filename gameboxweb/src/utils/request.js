import axios from "axios"
import CryptoJS from "crypto-js"
import Config from "@/setting"
import { toastMsg } from "./index"

// 加密
function encrypt(pwd, key, iv) {
  pwd = CryptoJS.enc.Utf8.parse(pwd) // 解析明文
  key = CryptoJS.enc.Utf8.parse(key) // 解析密钥
  iv = CryptoJS.enc.Utf8.parse(iv)
  const encrypted = CryptoJS.AES.encrypt(pwd, key, {
    mode: CryptoJS.mode.CBC, // 加密模式
    padding: CryptoJS.pad.Pkcs7, // 填充方式
    iv, // 向量
  })
  return encrypted.toString() // 加密后的结果是对象，要转为文本
}

function generateSecret() {
  let timestamp = new Date().getTime()
  let second = Math.floor(timestamp / 1000)
  return encrypt(Number(second) + 30, Config.SecretKey, Config.SecretKey)
}

const rs = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 20000,
})

// request拦截器
rs.interceptors.request.use(
  (config) => {
    config.headers["Authorization"] = "Box01 " + generateSecret()
    return config
  },
  (error) => {
    console.log("request", error)
    return Promise.reject(error)
  }
)

// 响应拦截器
rs.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    const err = error.response.data

    const errMsg = err.message ? err.message : "error"
    toastMsg(errMsg)
    return Promise.reject(err)
  }
)

export default rs
