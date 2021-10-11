//Permet d'attribuer un badge a un utilisateur manuellement.
function Badge(textarea, cat) {
    text = textarea.value
    var divVal = document.getElementsByClassName("selected")[0].firstChild;
    //Si le numéro du badge existe (en tout cas que c'est un nombre entier suéprieur a 0)
    if (parseInt(text, 10) > 0 && parseInt(text, 10) % 1 == 0) {
        var params = new Object()
        params.id = divVal.nextElementSibling.id
        params.val = text
        params.cat = cat
        fetch("/badge", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Accept: "application/json"
                },
                body: JSON.stringify(params)
            }).then(x => x.json())
            //Réponse qui me dit que la valeur est attribuer dans la base de donnée
            .then(function() { textarea.style.backgroundColor = "green" })
            //gestion erreur
            .catch(x => console.log(x.json()))
    }
}




//Crée une catégory
function Create() {
    var params = new Object()
    params.newVal = document.getElementById("newValue").value
    fetch("/category/", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Accept: "application/json"
            },
            body: JSON.stringify(params)
        }).then(x => x.json())
        //Réponse qui me dit que la valeur est attribuer dans la base de donnée
        .then(function(x) {
            row = categoryTable.insertRow(categoryTable.rows.length - 1)
            cell1 = row.insertCell(0)
            cell2 = row.insertCell(1)
            cell1.innerHTML = x.id
            cell2.innerHTML = x.new
            cell1.addEventListener('click', function() {
                if (!item.className.includes("selected")) {
                    removeSelected();
                    item.classList.add("selected")
                } else {
                    item.classList.remove("selected");
                }
            }, false);
            cell2.addEventListener('click', function() {
                if (!item.className.includes("selected")) {
                    removeSelected();
                    item.classList.add("selected")
                } else {
                    item.classList.remove("selected");
                }
            }, false);
        })
        //gestion erreur
        .catch(x => console.log(x))
}

//Permet de promouvoir un post en le sélectionnant (un seul peut être promus a la fois)
function Promote() {
    var divVal = document.getElementsByClassName("selected")[0].firstChild;
    var params = new Object();
    params.table = "promote"
    params.id = divVal.nextElementSibling.id
    fetch("/moderation", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Accept: "application/json"
            },
            body: JSON.stringify(params)
        }).then(x => x.json())
        //Réponse qui me dit que la valeur est attribuer dans la base de donnée
        .then(function(x) { document.getElementById("promote").style.color = "green" })
        //gestion erreur
        .catch(x => console.log(x.json()))
}

//Efface la ligne sélectionner de la base de donnée.
function DeleteRow(table) {
    var divVal = document.getElementsByClassName("selected")[0].firstChild;
    var params = new Object();
    params.table = table
    params.action = "DELETE"
    params.id = divVal.nextElementSibling.id
    fetch("/fetching", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Accept: "application/json"
            },
            body: JSON.stringify(params)
        }).then(x => x.json())
        //Réponse qui me dit que la valeur est attribuer dans la base de donnée
        .then(function(x) { divVal.nextElementSibling.innerHTML = "done" })
        //gestion erreur
        .catch(x => console.log(x.json()))
}

//Actualise la valeur ciblé dans la base de donnée.
function PostRequest(role) {
    var divVal = document.getElementsByClassName("selected")[0].firstChild;
    if ((divVal.nextElementSibling.dataset.name == "state" && divVal.nextElementSibling.dataset.role >= role) || (divVal.nextElementSibling.dataset.name == "role" && divVal.nextElementSibling.dataset.value > role)) {
        //Le modérateur ne peut pas changer l'état du compte d'un modérateur ou d'un administrateur, il ne peut pas non plus changer sont rôle
        console.log("PAS LE DROIT")
    } else {
        var params = new Object();
        params.table = actualTable
        params.action = "UPDATE"
        params.id = divVal.nextElementSibling.id
        params.What = divVal.nextElementSibling.dataset.name
        params.val = divVal.nextElementSibling.dataset.value
        params.newVal = document.getElementById("newValue").value
        console.log(params)
        fetch("/fetching", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Accept: "application/json"
                },
                body: JSON.stringify(params)
            }).then(x => x.json())
            //Réponse qui me dit que la valeur est attribuer dans la base de donnée
            .then(function(x) { 
                if(params.cat=="gif"){
                    divVal.nextElementSibling.src=x.message
                }else if(params.cat=="description"){
                    divVal.nextElementSibling.value =params.newVal
                }else{
                    divVal.nextElementSibling.innerHTML = params.newVal
                }
            })
            //gestion erreur
            .catch(x => console.log(x.json()))
    }
}