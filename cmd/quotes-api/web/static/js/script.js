const quoteText = document.querySelector(".quoteText")
const quoteAuthor = document.querySelector(".quoteAuthor")
const quoteFetch = document.querySelector(".quoteFetch")

;(async () => {
    const res = await fetch("/quotes")

    res.json().then((v) => {
        quoteText.innerText = v.quoteText
        quoteAuthor.innerText = v.quoteAuthor
    })
})()


quoteFetch.addEventListener("click", function () {
    fetch("/quotes").then((res) => {
        res.json().then((v) => {
            quoteText.innerText = v.quoteText
            quoteAuthor.innerText = v.quoteAuthor
        })
    })
})