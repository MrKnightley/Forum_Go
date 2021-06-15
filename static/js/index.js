// ======= FADE-IN CARD EFFECT =======

let sliders = document.getElementsByClassName("slider");
let radios = document.getElementsByClassName("slider_nav");

let n = 0;

function fade(num) {
    n = num;
    // Pour chaque slider[i]...
    for (let i = 0; i < sliders.length; i++) {
        // J'invisibilise la slider et la met en arrière-plan :
        sliders[i].style.opacity = "0"
        sliders[i].style.zIndex = 10;
    }
    // Je rends visibile et en 1èr plan uniquement la sliders[num] :
    sliders[num].style.opacity = '1'
    sliders[num].style.display = 'inherit'
    sliders[num].style.zIndex = 50; // Classe 'bullets' : z-index = 999 
}

// Quand on clique sur la flèche de droite :
function swipeRight() {
    n++;
    if (n >= radios.length) { // Retour au début de la liste
        n = 0;
    }
    fade(n)
    radios[n].checked = true

}

// Quand on clique sur la flèche de gauche :
function swipeLeft() {
    n--;
    if (n < 0) {
        n = radios.length - 1; // Envoi à la fin de la liste
    }
    fade(n)
    radios[n].checked = true
}
