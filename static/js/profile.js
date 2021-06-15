//Si l'utilisateur veut supprimer sont compte
function DeleteAccountRequest(userID) {
    var params = new Object();
    params.id = userID.toString();
    params.table = users
    fetch("/Profile", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Accept: "application/json"
            },
            body: JSON.stringify(params)
        }).then(x => x.json())
        //Réponse qui me dit que la valeur est attribuer dans la base de donnée
        .then(function(x) { divVal.innerHTML = x, divVal.dataset.value = x })
        //gestion erreur
        .catch(x => console.log(x.json()))

}
//Permet de sélectionner les informations de notre pages
var AllTd = document.getElementsByTagName('td')
var AllTr = document.getElementsByTagName('tr')
Array.prototype.forEach.call(AllTd, function(item) {
    item.addEventListener('click', function() {
        if (!item.className.includes("selected")) {
            var alreadySelected = document.getElementsByClassName("selected")
            for (var i = alreadySelected.length - 1; i >= 0; i--) {
                alreadySelected[i].classList.remove('selected')
            }
            item.classList.add("selected")
        } else {
            item.classList.remove("selected");
        }
    }, false);
});
//permet de supprimer un post ou commentaire
function PostRequest() {
    var divVal = document.getElementsByClassName("selected")[0].firstChild;
    let params = new Object();
    params.id = divVal.nextElementSibling.id
    params.table = divVal.nextElementSibling.dataset.name
    if (params.table == "deleted") {
        alert("Already deleted")
    } else {
        fetch("/Profile", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Accept: "application/json"
                },
                body: JSON.stringify(params)
            }).then(x => x.json())
            //Réponse qui me dit que la valeur est attribuer dans la base de donnée
            .then(function(x) { divVal.nextElementSibling.innerHTML = x.message })
            //gestion erreur
            .catch(x => console.log(x.json()))
    }
}

// Modifie le titre de l'onglet :
document.title = document.title.toUpperCase();