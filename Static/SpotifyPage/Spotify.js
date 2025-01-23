document.addEventListener("DOMContentLoaded", () => {
    console.log("Script loaded and DOM ready.");
    const form = document.getElementById("MyForm")
    const input = document.getElementById("InputField")
    const spotify = document.getElementById("spotify")

    form.addEventListener("submit", function(event){
        event.preventDefault()


        const inputValue = input.value


        fetch("http://localhost:8080/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({input: inputValue})
        })

        setTimeout(()=> {
            const currentSrc = spotify.src;
            spotify.src = currentSrc;
    }, 2000)


    })
})