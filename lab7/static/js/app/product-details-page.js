import { appendCss } from "./page-utils.js"
import { Navbar } from "./navbar.js"
import { getProducts } from "./api.js"

export class ProductDetailsPage {
    constructor(root) {
        this.root = root
    }

    async render() {
        const productId = window.location.pathname.split("/").at(-1)
        const products = await getProducts()
        const product = products.find(p => p.id == productId)

        const html = `
            <div class="page-content">
                <div class="item-details">
                    <img class="details-image img-thumbnail" src=${product.imageUrl}/>
                    <h1 class="details-name">${product.name}</h1>
                    <h1 class="details-price">${product.price}₽</h1>
                    <h1 class="details-description-header">Описание</h1>
                    <p class="details-description">${product.description}</p>
                </div>
            </div>
        `
        new Navbar(this.root).render()
        this.root.insertAdjacentHTML("beforeend", html)
        appendCss("/static/css/products.css")
    }
}
