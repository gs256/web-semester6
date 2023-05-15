export class Navbar {
    constructor(root) {
        this.root = root
    }

    render() {
        const html = `
            <nav class="navbar">
                <ul class="navbar__list">
                    <li class="navbar__list-item"><a href="/">Home</a></li>
                    <li class="navbar__list-item"><a href="/products">Products</a></li>
                    <li class="navbar__list-item"><a href="/cart">Cart</a></li>
                    <li class="navbar__list-item"><a href="/admin/products">Admin</a></li>
                </ul>
            </nav>
        `
        this.root.insertAdjacentHTML("beforeend", html)
    }
}
