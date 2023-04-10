const cartSchema = Object.freeze({
    products: [],
})

export class Cart {
    #cart

    constructor() {
        this.#cart = { ...cartSchema }
    }

    addProduct(productId) {
        if (!productId) {
            this.#save()
            return
        }

        if (this.#cart.products.includes(productId)) {
            this.#save()
            return
        }

        this.#cart.products.push(productId)
        this.#save()
    }

    removeProduct(productId) {
        this.#cart.products = this.#cart.products.filter(id => id != productId)
        this.#save()
    }

    products() {
        this.#restore()
        return this.#cart.products ?? []
    }

    #restore() {
        try {
            const cart = JSON.parse(localStorage.getItem("cart"))
            this.#cart = cart
        } catch {
            this.#cart = { ...cartSchema }
        }
    }

    #save() {
        localStorage.setItem("cart", JSON.stringify(this.#cart))
    }
}
