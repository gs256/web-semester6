import { Cart } from "../cart.js"
import { Navbar } from "./navbar.js"
import { appendCss } from "./page-utils.js"

export class ProductPage {
    constructor(root) {
        this.root = root
    }

    render() {
        const html = `
            <div class="page-content product-list-root">
                <div class="top-div card">
                    <h2 class="h2">Ограничение</h2>
                    <input type="number" min="1" class="top-input form-control mt-3"></input>
                    <a href="#" style="text-decoration: none">
                        <button class="item-limit-submit btn btn-primary mt-3">Применить</button>
                        <button class="item-limit-reset btn btn-outline-primary mt-3">Сбросить</button>
                    </a>
                </div>
            </div>
        `
        new Navbar(this.root).render()
        this.root.insertAdjacentHTML("beforeend", html)
        this.productsContentScript()
        this.productPaginationScript()
        appendCss("/static/css/products.css")
    }

    async productsContentScript() {
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

        window.PAGE_NUMBER = pageNumber
        window.PAGE_COUNT = numberOfPages

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
    }

    async productPaginationScript() {
        const urlSearchParams = new URLSearchParams(window.location.search)

        const previousPageButton = document.querySelector(".pagination-button.previous")
        const nextPageButton = document.querySelector(".pagination-button.next")
        const itemLimitInput = document.querySelector(".top-input")
        const itemLimitSubmitButton = document.querySelector(".item-limit-submit")
        const itemLimitResetButton = document.querySelector(".item-limit-reset")

        const pageNumber = getPageNumber()

        function getPageNumber() {
            const params = Object.fromEntries(urlSearchParams.entries())

            let pageNumber = parseInt(params.page)

            if (pageNumber == null || pageNumber == undefined || pageNumber < 0) {
                pageNumber = 1
                updateUrlParam("page", pageNumber)
            }

            return pageNumber
        }

        function previousPage() {
            if (pageNumber <= 1) return
            updateUrlParam("page", pageNumber - 1)
        }

        function nextPage() {
            updateUrlParam("page", pageNumber + 1)
        }

        function updateUrlParam(key, value) {
            urlSearchParams.set(key, value)
            window.location.search = urlSearchParams.toString()
        }

        function deleteUrlParam(name) {
            urlSearchParams.delete(name)
            window.location.search = urlSearchParams.toString()
        }

        function setItemLimit(value) {
            if (!parseInt(value)) return
            updateUrlParam("max", value)
        }

        function submitItemLimit() {
            const limit = parseInt(itemLimitInput.value)
            if (!limit || limit <= 0) {
                alert("Invalid limit")
                return
            }

            setItemLimit(limit)
        }

        function resetItemLimit() {
            deleteUrlParam("max")
            updateUrlParam("page", 1)
        }

        previousPageButton?.addEventListener("click", previousPage)
        nextPageButton?.addEventListener("click", nextPage)
        itemLimitSubmitButton?.addEventListener("click", submitItemLimit)
        itemLimitResetButton?.addEventListener("click", resetItemLimit)
    }
}
