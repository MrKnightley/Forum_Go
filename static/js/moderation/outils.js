//On place l'event pour la sélection sur chaque cellule
var AllTd = (Array.from(AllTdUser).concat(Array.from(AllTdPost))).concat(Array.from(AllTdComment)).concat(Array.from(AllTdCategory))
Array.prototype.forEach.call(AllTd, function(item) {
    item.addEventListener('click', function() {
        if (!item.className.includes("selected")) {
            removeSelected();
            item.classList.add("selected")
        } else {
            item.classList.remove("selected");
        }
    }, false);
});

//Efface la sélection si on change de tableau.
function removeSelected() {
    var alreadySelected = document.getElementsByClassName("selected")
    for (var i = alreadySelected.length - 1; i >= 0; i--) {
        alreadySelected[i].classList.remove('selected')
    }
}
//Cache les options qui ne correspondent pas au tableau (options du filtre)
function removeVisibility(NameClass) {
    options.forEach(element => {
        if (element[0].className.includes(NameClass)) {
            if (element[0].className.includes("Invisible")) {
                Array.prototype.forEach.call(element, function(el) {
                    el.classList.remove('Invisible');
                    el.classList.add('Visible');
                })
            }
        } else {
            if (!element[0].className.includes("Invisible")) {
                Array.prototype.forEach.call(element, function(el) {
                    el.classList.add('Invisible');
                    el.classList.remove('Visible');
                })
            }
        }
    })
}
//Résultat rendu invisible le temps de la recherche.
function tempInvi() {
    let all = document.getElementsByClassName("Visible")
    while (all.length) {
        Array.prototype.forEach.call(all, function(el) {
            el.classList.remove('Visible');
            el.classList.add('tempInvi')
        })
    }
}
//Réafiche les résultats une fois la barre de recherche vide.
function tempInviReverse() {
    let all = document.getElementsByClassName("tempInvi")
    while (all.length) {
        Array.prototype.forEach.call(all, function(el) {
            el.classList.remove('tempInvi');
            el.classList.add('Visible')
        })
    }
    let b = document.getElementsByClassName("tempVisi")
    while (b.length) {
        Array.prototype.forEach.call(b, function(el) {
            el.classList.remove('tempVisi');
        })
    }
}
//Le filte qui permet d'afficher les résultas qui nous intérresse dans la table
function filter(text) {
    if (text == "") {
        tempInviReverse()
    } else if (select.value != "") {
        tempInvi();
        let searchTable = document.getElementById(actualTable)
        let searchRow = searchTable.querySelectorAll("[data-name='" + select.value + "']")
        let result = Array.from(searchRow).filter(search => search.innerHTML.includes(text))
        result.forEach(elem => {
            elem.parentElement.parentElement.classList.add("tempVisi")
            elem.parentElement.parentElement.parentElement.parentElement.classList.add("tempVisi")
        })
        if (!result.length) {
            let b = document.getElementsByClassName("tempVisi")
            while (b.length) {
                Array.prototype.forEach.call(b, function(el) {
                    el.classList.remove('tempVisi');
                })
            }
        }
    }
}
