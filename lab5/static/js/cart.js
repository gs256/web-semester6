const cartSchema = Object.freeze({
    products: [],
})

export class Cart {
    constructor() {
        this._cart = { ...cartSchema }
    }

    addProduct(productId) {
        if (!productId) {
            this.#save()
            return
        }

        if (this._cart?.products.includes(productId)) {
            this.#save()
            return
        }

        this._cart.products.push(productId)
        this.#save()
    }

    removeProduct(productId) {
        this._cart.products = this._cart.products.filter(id => id != productId)
        this.#save()
    }

    products() {
        this.#restore()
        return this._cart?.products ?? []
    }

    clear() {
        this._cart.products = []
        this.#save()
    }

    #restore() {
        try {
            const cart = JSON.parse(localStorage.getItem("cart"))
            if (cart) this._cart = cart
            else this._cart = { ...cartSchema }
        } catch {
            this._cart = { ...cartSchema }
        }
    }

    #save() {
        localStorage.setItem("cart", JSON.stringify(this._cart))
    }
}
