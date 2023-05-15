const urlSearchParams = new URLSearchParams(window.location.search)

const previousPageButton = document.querySelector(".pagination-button.previous")
const nextPageButton = document.querySelector(".pagination-button.next")
const itemLimitInput = document.querySelector(".top-input")
const itemLimitSubmitButton = document.querySelector(".item-limit-submit")
const itemLimitResetButton = document.querySelector(".item-limit-reset")

const pageNumber = getPageNumber()

function getPageNumber() {
    const params = Object.fromEntries(urlSearchParams.entries())

    let pageNumber = window.PAGE_NUMBER

    if (parseInt(params.page) != pageNumber) {
        updateUrlParam("page", pageNumber)
    }

    if (pageNumber == null || pageNumber == undefined || pageNumber < 0) pageNumber = 1

    return pageNumber
}

function previousPage() {
    if (pageNumber <= 1) return
    updateUrlParam("page", pageNumber - 1)
}

function nextPage() {
    updateUrlParam("page", pageNumber + 1)
}

function updateUrlParam(key, value) {
    urlSearchParams.set(key, value)
    window.location.search = urlSearchParams.toString()
}

function deleteUrlParam(name) {
    urlSearchParams.delete(name)
    window.location.search = urlSearchParams.toString()
}

function setItemLimit(value) {
    if (!parseInt(value)) return
    updateUrlParam("max", value)
}

function submitItemLimit() {
    const limit = parseInt(itemLimitInput.value)
    if (!limit || limit <= 0) {
        alert("Invalid limit")
        return
    }

    setItemLimit(limit)
}

function resetItemLimit() {
    deleteUrlParam("max")
}

previousPageButton?.addEventListener("click", previousPage)
nextPageButton?.addEventListener("click", nextPage)
itemLimitSubmitButton?.addEventListener("click", submitItemLimit)
itemLimitResetButton?.addEventListener("click", resetItemLimit)
