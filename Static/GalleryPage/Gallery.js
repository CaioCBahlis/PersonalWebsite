document.addEventListener("DOMContentLoaded", function() {
    const myphotos = document.querySelectorAll('.Photo-Box')
    myphotos.forEach((element) =>{
        element.addEventListener('click', () =>{
            myphotos.forEach((photo) => {
                if (photo !== element) {
                    photo.classList.remove("expanded");
                    photo.querySelector('.PhotoDesc').style.visibility = "hidden"
                }
            });

            element.classList.toggle("expanded")
            if (element.classList.contains("expanded")){
                element.querySelector('.PhotoDesc').style.visibility = "visible"
            }else{
                element.querySelector('.PhotoDesc').style.visibility = "hidden"
            }

        })
    })
})