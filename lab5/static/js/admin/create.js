import { ProductEditForm } from "./product-edit-form.js"

const submitButton = document.querySelector(".product-edit__submit")
const goBackButton = document.querySelector(".product-edit__go-back")

submitButton.addEventListener("click", onSubmitClicked)
goBackButton.addEventListener("click", onGoBackClicked)

async function onSubmitClicked() {
  const form = new ProductEditForm()

  try {
    form.validate()
  } catch (e) {
    alert(e.message)
    return
  }

  const requestBody = form.createBody()

  const response = await fetch("/admin/products/create", {
    method: "POST",
    body: JSON.stringify(requestBody),
  })

  if (!response.ok) {
    alert(`Error: ${await response.text()}`)
  } else {
    alert(`Product added successfully`)
    form.clear()
  }
}

function onGoBackClicked() {
  window.location.replace(document.referrer)
}
