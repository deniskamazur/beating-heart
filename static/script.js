var heart;
var message;

window.onload = function() {
    heart = document.getElementById("heart")

    message = document.getElementById("message")
    message.innerText = "When I don't see you"
}

function onClick(checkbox) {
    if (checkbox.checked) {
        heart.style.animation = ".6s infinite beatHeart"
        message.innerText = "When I see you"
    } else {
        heart.style.animation = "1s infinite beatHeart"
        message.innerText = "When I don't see you"
    }
}