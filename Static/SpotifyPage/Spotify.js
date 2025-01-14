document.addEventListener("DOMContentLoaded", () => {
    console.log("Script loaded and DOM ready.");
    const form = document.getElementById("MyForm")
    const input = document.getElementById("input")

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
    })
})