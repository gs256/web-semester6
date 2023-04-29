import { Cart } from "./cart.js"

const cart = new Cart()
const cartProductIds = cart.products()
const itemList = document.querySelector(".cart__item-list")
const orderButton = document.querySelector(".cart__sidebar-form__order-button")
const products = await getProducts()

orderButton.onclick = async () => {
    await createOrder()
    cart.clear()
    location.reload()
}

async function createOrder() {
    const user = getUser()

    if (!isUserValid(user)) {
        alert("Invalid name or phone")
        return
    }

    const userResponse = await fetch("/api/users/create", {
        method: "POST",
        body: JSON.stringify(user),
    })

    if (!userResponse.ok) {
        alert(`Error: '${await userResponse.text()}'`)
        return
    }

    const userId = (await userResponse.json())?.id
    const productIds = cart.products()

    if (!productIds || productIds.length == 0) {
        alert("Nothing to checkout")
        return
    }

    const orderRequestBody = {
        userId: userId,
        productIds: productIds,
    }

    console.log(JSON.stringify(orderRequestBody))

    const orderResponse = await fetch("/api/orders/create", {
        method: "POST",
        body: JSON.stringify(orderRequestBody),
    })

    if (!orderResponse.ok) {
        alert(`Error: '${await orderResponse.text()}'`)
        return
    }

    alert("Successfully created an order")
    console.log(await orderResponse.json())
}

function getUser() {
    const nameInput = document.querySelector(".cart__sidebar-form__name")
    const phoneInput = document.querySelector(".cart__sidebar-form__phone")

    return {
        name: nameInput.value,
        phone: phoneInput.value,
    }
}

function isUserValid(user) {
    return user?.name && user?.phone
}

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
    <img class="cart__item-image img-thumbnail" src="${imageUrl}"/>
    <div class="cart__item-info">
        <h1 class="cart__item-title">${name}</h1>
        <h1 class="cart__item-price">${price} ₽</h1>
    </div>
    <button class="cart__item-remove btn btn-close"></button>
`
    const div = document.createElement("div")
    div.classList.add("card", "cart__item")
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
