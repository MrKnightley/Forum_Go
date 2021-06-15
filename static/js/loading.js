// Lorsque la page est totalement chargée... 
$(window).on("load", function() {
    $(".loader-wrapper").fadeOut("slow"); // Fait disparaître l'écran de chargement
    $("html").css({ overflow: 'auto' }); // Rétablit la barre de scroll
})