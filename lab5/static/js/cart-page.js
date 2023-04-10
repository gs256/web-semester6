import { Cart } from "./cart.js"

const cart = new Cart()
const cartProductIds = cart.products()
const itemList = document.querySelector(".cart__item-list")
const products = await getProducts()

for (const productId of cartProductIds) {
    const product = getProductById(products, productId)

    const element = itemList.appendChild(
        createCartItemElement(product.name, product.price, product.imageUrl)
    )

    const removeButton = element.querySelector(".cart__item-remove")
    removeButton.onclick = () => {
        cart.removeProduct(productId)
        element.remove()
    }
}

function createCartItemElement(name, price, imageUrl) {
    const innerHtml = `
    <img class="cart__item-image" src="${imageUrl}"/>
    <div class="cart__item-info">
        <h1 class="cart__item-title">${name}</h1>
        <h1 class="cart__item-price">${price} ₽</h1>
    </div>
    <button class="cart__item-remove">x</button>
`
    const div = document.createElement("div")
    div.classList.add("cart__item")
    div.innerHTML = innerHtml.trim()
    return div
}

async function getProducts() {
    const response = await fetch("/api/products")
    return response.json()
}

function getProductById(products, id) {
    return products.find(product => product.id == id)
}
