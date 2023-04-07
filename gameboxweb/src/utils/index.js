import Toastify from "toastify-js"
import "toastify-js/src/toastify.css"

export function toastMsg(msg) {
  Toastify({
    text: msg,
    position: "center",
    style: {
      background: "rgba(0, 0, 0, 0.7)",
      padding: "5px 20px",
      fontSize: "0.9rem",
    },
    duration: 5000,
  }).showToast()
}
