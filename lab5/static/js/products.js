// let cart = {
//     products: [],
// }

import { Cart } from "./cart.js"

// cart = JSON.parse(localStorage.getItem("cart")) || cart

const addToCartButtons = document.querySelectorAll(".product-add-to-cart-button")

// console.log(addToCartButtons)

// function addToCart(cart, productId) {
//     cart.products.push(productId)
//     console.log(cart)
// }

// function saveCart(cart) {
//     localStorage.setItem("cart", JSON.stringify(cart))
// }
const cart = new Cart()
const productsInCart = cart.products()

for (let button of addToCartButtons) {
    const productId = button.getAttribute("data-id")

    if (productsInCart.includes(productId)) {
        button.disabled = true
    } else {
        button.onclick = () => {
            cart.addProduct(productId)
            console.log(cart.products())
            button.disabled = true
        }
    }
}

window.onpageshow = event => {
    if (event.persisted) {
        window.location.reload()
    }
}
