const dropboxes = document.querySelectorAll(".Project-Box");
const desc = document.querySelectorAll(".hidden-desc")

document.addEventListener("DOMContentLoaded", () => {
    dropboxes.forEach((dropbox, index) => {
        dropbox.dataset.expanded = "false";
        dropbox.addEventListener("click", () => {
            const computedHeight = window.getComputedStyle(dropbox).height;
            if (dropbox.dataset.expanded === "false") {


                desc[index].style.display = "flex"
                dropbox.style.flexDirection = "column"
                dropbox.style.cursor = "auto"
                dropbox.querySelector("h3").style.paddingTop = "2rem"

                let height = dropbox.offsetHeight;
                dropbox.style.height = "auto";
                let autoHeight = dropbox.offsetHeight;
                dropbox.style.height = `${height}px`;


                void dropbox.offsetHeight;


                dropbox.style.transition = "height 0.7s smooth";
                dropbox.style.height = `${autoHeight}px`;
                dropbox.dataset.expanded = "true"

            } else {
                dropbox.style.height = "14vh"
                desc[index].style.display = "none"
                 dropbox.dataset.expanded = "false"
                 dropbox.style.cursor = "cursor"
                dropbox.querySelector("h3").style.paddingTop  = "0"

            }
        });
    });
})
