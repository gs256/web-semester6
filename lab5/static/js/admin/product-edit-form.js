export class ProductEditForm {
  constructor() {
    this.nameInput = document.querySelector("#product-edit__name")
    this.priceInput = document.querySelector("#product-edit__price")
    this.descriptionInput = document.querySelector("#product-edit__description")
    this.imageUrlInput = document.querySelector("#product-edit__image-url")
    this.update()
  }

  validate() {
    this.update()
    ensureStringFieldValid(this.name, "name")
    ensureNumberFieldValid(this.price, "price")
    ensureStringFieldValid(this.description, "description")
    ensureStringFieldValid(this.imageUrl, "image url")
  }

  createBody() {
    this.validate()

    const requestBody = {
      name: this.name,
      price: parseInt(this.price),
      description: this.description,
      imageUrl: this.imageUrl,
    }
    return requestBody
  }

  update() {
    this.name = this.nameInput.value
    this.price = this.priceInput.value
    this.description = this.descriptionInput.value
    this.imageUrl = this.imageUrlInput.value
  }

  clear() {
    this.name = this.nameInput.value = ""
    this.price = this.priceInput.value = ""
    this.description = this.descriptionInput.value = ""
    this.imageUrl = this.imageUrlInput.value = ""
  }
}

function ensureStringFieldValid(string, fieldName) {
  const isValidString = typeof string == "string" && Boolean(string)
  if (!isValidString) throw Error(`Invalid string in ${fieldName}: "${string}"`)
}

function ensureNumberFieldValid(number, fieldName) {
  try {
    !isNaN(Number(number))
  } catch {
    throw Error(`Invalid number in ${fieldName}: "${number}"`)
  }
}
