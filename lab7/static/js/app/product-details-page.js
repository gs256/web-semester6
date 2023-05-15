import { appendCss } from "./page-utils.js"
import { Navbar } from "./navbar.js"

export class ProductDetailsPage {
    constructor(root) {
        this.root = root
    }

    render() {
        const html = `
            <div class="page-content">
                <div class="item-details">
                    <img class="details-image img-thumbnail" src={{ .ImageUrl }}/>
                    <h1 class="details-name">{{ .Name }}</h1>
                    <h1 class="details-price">{{ .Price }}₽</h1>
                    <h1 class="details-description-header">Описание</h1>
                    <p class="details-description">{{ .Description }}</p>
                </div>
            </div>
        `
        new Navbar(this.root).render()
        this.root.insertAdjacentHTML("beforeend", html)
        appendCss("/static/css/products.css")
    }
}
