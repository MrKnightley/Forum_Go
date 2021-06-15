// Lorsque la page est totalement chargée... 
// $(window).on("load", function() {
//     $("#flipping-page").addClass("animated");
// })

function moveLeft() {
    let section = document.getElementById("main");
    section.style.left = "calc(50vw - (var(--baseline) * 60)/1.35)"; // Centre le livre au milieu de la page
}

function moveRight() {
    let section = document.getElementById("main");
    section.style.left = "calc(50vw - (var(--baseline) * 60)/2)"; //
}