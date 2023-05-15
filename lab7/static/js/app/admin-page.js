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
                    {{range .products}}
                        <div class="admin-product-item">
                            <span class="admin-item-name">
                                <span>{{ .Name }}</span>
                            </span>
                            <div class="admin-item__buttons">
                                <a class="admin-item-edit link-button" href="/admin/products/edit/{{ .Id }}">Edit</a>
                                <a class="admin-item-delete link-button" href="/admin/products/delete/{{ .Id }}">Delete</a>
                            </div>
                        </div>
                    {{end}}
                </div>
                <a href="/admin/products/create">
                    <button>Create</button>
                </a>
            </div>
        `
        new Navbar(this.root).render()
        this.root.insertAdjacentHTML("beforeend", html)
        appendCss("/static/css/admin.css")
    }
}
