import { appendCss } from "./page-utils.js"
import { Navbar } from "./navbar.js"

export class AdminPage {
    constructor(root) {
        this.root = root
    }

    render() {
        const html = `
            <div class="page-content">
                <h1>Products</h1>
                <div class="admin-product-list">
                </div>
                <a href="/admin/products/create">
                    <button class="btn btn-primary">Create</button>
                </a>
            </div>
        `
        new Navbar(this.root).render()
        this.root.insertAdjacentHTML("beforeend", html)
        appendCss("/static/css/admin.css")
        this.populateItems()
    }

    populateItems() {
        const list = document.querySelector(".admin-product-list")

        for (let i = 0; i < 10; i++) {
            list.appendChild(this.createMockItem())
        }
    }

    createMockItem() {
        const html = `
            <div class="card admin-product-item">
                <span class="admin-item-name">
                    <span>Mock Item</span>
                </span>
                <div class="admin-item__buttons">
                    <a class="admin-item-edit link-button" href="/admin/products/edit/#">Edit</a>
                    <a class="admin-item-delete link-button" href="/admin/products/delete/#">Delete</a>
                </div>
            </div>
        `

        var div = document.createElement("div")
        div.innerHTML = html.trim()
        return div.firstChild
    }
}
