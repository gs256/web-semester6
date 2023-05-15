import { CartPage } from "./cart-page.js"
import { HomePage } from "./home-page.js"
import { ProductPage } from "./product-page.js"

const root = document.getElementById("root")
const route = window.location.pathname

if (route === "/") {
    new HomePage(root).render()
} else if (route == "/products") {
    new ProductPage(root).render()
} else if (route == "/cart") {
    new CartPage(root).render()
}
