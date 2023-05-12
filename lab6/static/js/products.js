import { Cart } from "./cart.js"

window.onpageshow = event => {
    if (event.persisted) {
        window.location.reload()
    }
}

async function getProducts() {
    const response = await fetch("/api/products")
    return response.json()
}

function getSpecificPageUrl(pageNumber) {
    let url = new URL(location.href)
    url.searchParams.set("page", pageNumber)
    return url.toString()
}

function createPaginationElement(currentPage, pageCount) {
    const previousPageUrl = getSpecificPageUrl(currentPage - 1)
    const nextPageUrl = getSpecificPageUrl(currentPage + 1)

    const html = `
        <div class="pagination">
            ${
                currentPage > 1
                    ? `<a href="${previousPageUrl}" class="pagination-button previous">←</a>`
                    : ""
            }
            <span class="pagination-text">${currentPage} / ${pageCount}</span>
            ${
                currentPage < pageCount
                    ? `<a href="${nextPageUrl}" class="pagination-button next">→ </a>`
                    : ""
            }
        </div>
    `
    var div = document.createElement("div")
    div.innerHTML = html.trim()
    return div.firstChild
}

function createItemListElement(productId, name, description, price, imageUrl) {
    const html = `
        <div class="product-item card">
            <div class="product-item__image-container">
                <img class="product-image img-thumbnail" src=${imageUrl}/>
            </div>
            <div class="product-item__name-container">
                <a href="/products/${productId}" class="product-name">${name}</a>
                <p class="product-item__description">${description}</p>
            </div>
            <div class="product-item__side-container">
                <span class="product-price">${price}₽</span>
                <button class="product-add-to-cart-button btn btn-primary" data-id="${productId}">Add to cart</button>
            </div>
        </div>
    `

    var div = document.createElement("div")
    div.innerHTML = html.trim()
    return div.firstChild
}

const productListRoot = document.querySelector(".product-list-root")
const cart = new Cart()
const productsInCart = cart.products()
let products = await getProducts()

let pageNumber = parseInt(new URLSearchParams(location.search).get("page"))
let pageLimit = parseInt(new URLSearchParams(location.search).get("max"))

if (!pageNumber) pageNumber = 1
if (!pageLimit) pageLimit = products.length

const numberOfPages = parseInt(products.length / pageLimit)

const sliceStart = pageLimit * (pageNumber - 1)
const sliceEnd = sliceStart + pageLimit
products = products.slice(sliceStart, sliceEnd)

products.forEach(product => {
    productListRoot.appendChild(
        createItemListElement(
            product.id,
            product.name,
            product.description,
            product.price,
            product.imageUrl
        )
    )
})

const paginationElement = createPaginationElement(pageNumber, numberOfPages)
productListRoot.appendChild(paginationElement)

const addToCartButtons = document.querySelectorAll(".product-add-to-cart-button")
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
