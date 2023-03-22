export class ProductEditForm {
  constructor() {
    this.nameInput = document.querySelector("#product-edit__name")
    this.priceInput = document.querySelector("#product-edit__price")
    this.descriptionInput = document.querySelector("#product-edit__description")
    this.imageInput = document.querySelector("#product-edit__image-url")
    this.update()
  }

  validate() {
    this.update()
    ensureStringFieldValid(this.name, "name")
    ensureNumberFieldValid(this.price, "price")
    ensureStringFieldValid(this.description, "description")
    ensureFileValid(this.image, "image")
  }

  createBody() {
    this.validate()
    let form = new FormData()
    form.append("name", this.name)
    form.append("price", parseInt(this.price))
    form.append("description", this.description)
    form.append("image", this.image)

    // const requestBody = {
    //   name: this.name,
    //   price: parseInt(this.price),
    //   description: this.description,
    //   image: this.image,
    // }
    // return requestBody
    return form
  }

  update() {
    this.name = this.nameInput.value
    this.price = this.priceInput.value
    this.description = this.descriptionInput.value
    this.image = this.imageInput.files[0]
  }

  clear() {
    this.name = this.nameInput.value = ""
    this.price = this.priceInput.value = ""
    this.description = this.descriptionInput.value = ""
    this.image = undefined
    this.imageInput.files = []
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

function ensureFileValid(file, fieldName) {
  if (!file) throw Error(`Invalid number in ${fieldName}: "${number}"`)
}
