var heart;

window.onload = function() {
    heart = document.getElementById("heart")
}

function onClick(checkbox) {
    if (checkbox.checked) {
        heart.style.animation = ".6s infinite beatHeart"
    } else {
        heart.style.animation = "1s infinite beatHeart"
    }
}