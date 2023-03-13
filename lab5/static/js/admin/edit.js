import { ProductEditForm } from "./product-edit-form.js"

const applyButton = document.querySelector(".product-edit__apply")
const goBackButton = document.querySelector(".product-edit__go-back")

applyButton.addEventListener("click", onApplyClicked)
goBackButton.addEventListener("click", onGoBackClicked)

async function onApplyClicked() {
  const form = new ProductEditForm()

  try {
    form.validate()
  } catch (e) {
    alert(e.message)
    return
  }

  const requestBody = form.createBody()

  const response = await fetch(window.location, {
    method: "POST",
    body: JSON.stringify(requestBody),
  })

  if (!response.ok) {
    alert(`Error: ${await response.text()}`)
  } else {
    alert(`Product updated successfully`)
  }
}

function onGoBackClicked() {
  window.location.replace(document.referrer)
}
