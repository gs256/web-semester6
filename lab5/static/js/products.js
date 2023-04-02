let cart = {
    products: [],
}

cart = JSON.parse(localStorage.getItem("cart")) || cart

const addToCartButtons = document.querySelectorAll(".product-add-to-cart-button")

console.log(addToCartButtons)

function addToCart(cart, productId) {
    cart.products.push(productId)
    console.log(cart)
}

function saveCart(cart) {
    localStorage.setItem("cart", JSON.stringify(cart))
}

for (let button of addToCartButtons) {
    button.onclick = () => {
        const productId = button.getAttribute("data-id")
        if (!productId) {
            saveCart(cart)
            return
        }
        if (cart.products.includes(productId)) {
            saveCart(cart)
            return
        }
        addToCart(cart, productId)
    }
}
