import { Navbar } from "./navbar.js"

export class HomePage {
    constructor(root) {
        this.root = root
    }

    render() {
        const html = `
            <div class="container" style="margin: 40px">
                <h1>Home Page</h1>
                <a href="/products">/products</a></br>
                <a href="/admin/products">/admin/products</a></br>
                </br>
                <a href="/test/fill-with-test-data">/test/fill-with-test-data</a></br>
            </div>
        `
        new Navbar(this.root).render()
        this.root.insertAdjacentHTML("beforeend", html)
    }
}
