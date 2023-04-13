import CryptoJS from "crypto-js"
import Settings from "@/setting"

// cookie保存的天数
const cache_expires = 36000
const SECRET_KEY = CryptoJS.enc.Utf8.parse(Settings.Uid)
const SECRET_IV = CryptoJS.enc.Utf8.parse(Settings.OpenId)

// 类型 window.localStorage,window.sessionStorage,
const config = {
  type: "localStorage", // 本地存储类型 sessionStorage
  prefix: "gigame_box01", // 名称前缀 建议：项目名 + 项目版本
  expire: cache_expires, //过期时间 单位：秒
  isEncrypt: false, // 默认加密 为了调试方便, 开发过程中可以不加密
}

// 设置 setStorage
export const setStorage = (key, value, expire = 0) => {
  if (value === "" || value === null || value === undefined) {
    value = null
  }

  if (isNaN(expire) || expire < 0) throw new Error("Expire must be a number")

  expire = expire ? expire : config.expire
  let data = {
    value: value, // 存储值
    time: Date.now(), //存值时间戳
    expire: expire, // 过期时间
  }

  // const encryptString = encrypt(JSON.stringify(data))
  const encryptString = JSON.stringify(data)
  window.localStorage.setItem(autoAddPrefix(key), encryptString)
}

// 名称前自动添加前缀
const autoAddPrefix = (key) => {
  const prefix = config.prefix ? config.prefix + "_" : ""
  return prefix + key
}

// 获取 getStorage
export const getStorage = (key) => {
  key = autoAddPrefix(key)
  // key 不存在判断
  if (
    !window.localStorage.getItem(key) ||
    JSON.stringify(window.localStorage.getItem(key)) === "null"
  ) {
    return null
  }

  // 优化 持续使用中续期
  // const storage = JSON.parse(decrypt(window.localStorage.getItem(key)))
  const storage = JSON.parse(window.localStorage.getItem(key))
  let nowTime = Date.now()
  // 过期删除
  let setExpire = (storage.expire || config.expire) * 1000,
    expDiff = nowTime - storage.time
  if (setExpire < expDiff) {
    removeStorage(key)
    return null
  } else {
    // 未过期期间被调用 则自动续期 进行保活
    setStorage(autoRemovePrefix(key), storage.value, storage.expire)
    return storage.value
  }
}

// 移除已添加的前缀
const autoRemovePrefix = (key) => {
  const len = config.prefix ? config.prefix.length + 1 : ""
  return key.substr(len)
}

// 是否存在 hasStorage
export const hasStorage = (key) => {
  key = autoAddPrefix(key)
  let arr = getStorageAll().filter((item) => {
    return item.key === key
  })
  return arr.length ? true : false
}

// 获取全部 getAllStorage
export const getStorageAll = () => {
  let len = window.localStorage.length // 获取长度
  let arr = new Array() // 定义数据集
  for (let i = 0; i < len; i++) {
    // 获取key 索引从0开始
    let getKey = window.localStorage.key(i)
    // 获取key对应的值
    let getVal = window.localStorage.getItem(getKey)
    // 放进数组
    arr[i] = { key: getKey, val: getVal }
  }
  return arr
}

// 删除 removeStorage
export const removeStorage = (key) => {
  window.localStorage.removeItem(autoAddPrefix(key))
}

/**
 * 加密方法
 * @param data
 * @returns {string}
 */
const encrypt = (data) => {
  if (typeof data === "object") {
    try {
      data = JSON.stringify(data)
    } catch (error) {
      console.log("encrypt error:", error)
    }
  }
  const dataHex = CryptoJS.enc.Utf8.parse(data)
  const encrypted = CryptoJS.AES.encrypt(dataHex, SECRET_KEY, {
    iv: SECRET_IV,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7,
  })
  return encrypted.ciphertext.toString()
}

/**
 * 解密方法
 * @param data
 * @returns {string}
 */
const decrypt = (data) => {
  const encryptedHexStr = CryptoJS.enc.Hex.parse(data)
  const str = CryptoJS.enc.Base64.stringify(encryptedHexStr)
  const decrypt = CryptoJS.AES.decrypt(str, SECRET_KEY, {
    iv: SECRET_IV,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7,
  })
  const decryptedStr = decrypt.toString(CryptoJS.enc.Utf8)
  return decryptedStr.toString()
}
