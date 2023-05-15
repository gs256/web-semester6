export function appendCss(filename) {
    var fileref = document.createElement("link")
    fileref.setAttribute("rel", "stylesheet")
    fileref.setAttribute("type", "text/css")
    fileref.setAttribute("href", filename)
    document.getElementsByTagName("head")[0].appendChild(fileref)
}

export function createElementFromHtml(html) {
    var div = document.createElement("div")
    div.innerHTML = html.trim()
    return div.firstChild
}
