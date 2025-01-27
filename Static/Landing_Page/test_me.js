const dropboxes = document.querySelectorAll(".Project-Box");
const desc = document.querySelectorAll(".hidden-desc")

dropboxes.forEach((dropbox, index) => {
    dropbox.addEventListener("click", () => {
        if (dropbox.style.height === "14vh"){



            desc[index].style.display = "flex"
            dropbox.style.flexDirection = "column"

            let height = dropbox.offsetHeight;
            dropbox.style.height = "auto";
            let autoHeight = dropbox.offsetHeight;
            dropbox.style.height = `${height}px`;


            void dropbox.offsetHeight;


            dropbox.style.transition = "height 0.5s smooth";
            dropbox.style.height = `${autoHeight}px`;

        }else{
            dropbox.style.height = "14vh"
            desc[index].style.display = "none"
        }
    });
});
