// Script coupant le contenu des posts pour n'avoir que les ~450 premiers caractères :

let contents = document.getElementsByClassName('content'); // Toutes les div contenant le contenu d'un post

for (let i = 0; i < contents.length; i++) {
    contents[i].textContent = contents[i].textContent.substring(0, 450) + "...";
}