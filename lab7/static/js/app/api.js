export async function getProducts() {
    const response = await fetch("/api/products")
    return response.json()
}
