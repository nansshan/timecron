let a = 0
let timer = setInterval(() => {
    a++
    console.log(a)
    if (a>= 10) {
            clearInterval(timer)
    }
}, 1000);