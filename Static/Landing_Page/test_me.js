const dropboxes = document.querySelectorAll(".Project-Box");
const desc = document.querySelectorAll(".hidden-desc")

dropboxes.forEach((dropbox, index) => {
    dropbox.addEventListener("click", () => {
        if (dropbox.style.height === "14vh"){
            dropbox.style.height = "70vh"
            desc[index].style.display = "flex"
            dropbox.style.flexDirection = "column"
        }else{
            dropbox.style.height = "14vh"
            desc[index].style.display = "none"
        }
    });
});
