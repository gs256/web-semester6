import { appendCss, createElementFromHtml } from "./page-utils.js"
import { Navbar } from "./navbar.js"
import { getProducts } from "./api.js"

export class AdminPage {
    constructor(root) {
        this.root = root
    }

    async render() {
        const products = await getProducts()

        const html = `
            <div class="page-content">
                <h1>Products</h1>
                <div class="admin-product-list">
                </div>
                <a href="/admin/products/create">
                    <button>Create</button>
                </a>
            </div>
        `
        new Navbar(this.root).render()
        this.root.insertAdjacentHTML("beforeend", html)
        appendCss("/static/css/admin.css")
        this.populateProducts(products)
    }

    populateProducts(products) {
        const list = document.querySelector(".admin-product-list")

        function createProductItem(product) {
            const html = `
                <div class="admin-product-item">
                    <span class="admin-item-name">
                        <span>${product.name}</span>
                    </span>
                    <div class="admin-item__buttons">
                        <a class="admin-item-edit link-button" href="/admin/products/edit/${product.id}">Edit</a>
                        <a class="admin-item-delete link-button" href="/admin/products/delete/${product.id}">Delete</a>
                    </div>
                </div>
            `
            return createElementFromHtml(html)
        }

        products.forEach(product => {
            list.appendChild(createProductItem(product))
        })
    }
}
