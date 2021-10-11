// Copier l'URL dans le presse-papier (bouton 'Share') :
function CopyURL() {
    document.getElementsByClassName("validation")[0].classList.remove("active");

    console.log("CLICKED!")
    var dummy = document.createElement('input');
    dummy.className = 'dummy';
    myURL = window.location.href;
    dummy.setAttribute('value', myURL)
    document.body.appendChild(dummy);
    console.log(dummy.value);
    dummy.select();
    document.execCommand('copy');
    document.body.removeChild(dummy);

    document.getElementsByClassName("validation")[0].classList.add("active");
}
//Vide le gif si on ne veut finalement pas en mettre dans le commentaire
function empty(me) {
    document.getElementById('giflink').value = 'Choose a gif'
    document.getElementById("inputgif").value = 'Choose a gif'
    me.classList.add('Invisible')
}
//Permet de choisir un gif et d'ouvrir le pannel
function selectGif() {
    grab_data(search_term)
    document.getElementsByClassName("tenorTest")[0].classList.toggle("Invisible")
}
//Ferme le pannel
function closeGif() {
    gif = document.getElementsByClassName("tenorTest")[0]
    if (!gif.classList.contains("Invisible")) {
        gif.classList.add("Invisible")
    }
}
//Rajoute le gif sélectionner au form pour l'envoyer si on valide
function newGifSelected(val) {
    cross = document.getElementById('empty')
    if (cross.classList.contains("Invisible")) {
        cross.classList.remove("Invisible")
    }
    document.getElementById("giflink").value = val
    document.getElementById("inputgif").value = val
}

//TENOR
// url Async requesting function
function httpGetAsync(theUrl, callback) {
    // create the request object
    var xmlHttp = new XMLHttpRequest();

    // set the state change callback to capture when the response comes in
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
            callback(xmlHttp.responseText);
        }
    }

    // open as a GET call, pass in the url and set async = True
    xmlHttp.open("GET", theUrl, true);

    // call send with no params as they were passed in on the url string
    xmlHttp.send(null);

    return;
}

// callback for the top 8 GIFs of search
function tenorCallback_search(responsetext) {
    // parse the json response
    var response_objects = JSON.parse(responsetext);

    top_10_gifs = response_objects["results"];

    // load the GIFs -- for our example we will load the first GIFs preview size (nanogif) and share size (tinygif)
    for (i = 0; i < 8; i++) {
        document.getElementById("share_gif" + (1 + i)).src = top_10_gifs[i]["media"][0]["tinygif"]["url"];
    }
    return;

}

var search_term = ""
    // function to call the trending and category endpoints
function grab_data(st) {
    search_term = st
        // set the apikey and limit
    var apikey = "L8VYXS2JMKAS";
    var lmt = 8;

    // test search term
    // using default locale of en_US
    var search_url = "https://g.tenor.com/v1/search?q=" + search_term + "&key=" +
        apikey + "&limit=" + lmt;

    httpGetAsync(search_url, tenorCallback_search);

    // data will be loaded by each call's callback
    return;
}

//POST REQUEST

//Permet de supprimer un gif de sont commentaire
function DeleteGif(id) {
    var params = new Object()
    params.id = id.toString()
    params.table = "comments"
    params.action = "DELETE"
    params.What = "gif"
    params.is="cell"
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
            document.getElementById("gif" + id).src = ""
        })
        //gestion erreur
        .catch(x => console.log(x))
}
//Permet de valider les modifications
function Delete(Table, id, p, textarea) {
    var params = new Object()
    switch (Table) {
        case "posts":
            var textarea = document.getElementById('EditArea')
            var editBtn = document.getElementById("edit")
            var finishBtn = document.getElementById("validate")
            var delBtn = document.getElementById("delete")
            break
        case "comments":
            var textarea = document.getElementById('EditArea' + id)
            var editBtn = document.getElementById("edit" + id)
            var finishBtn = document.getElementById("validate" + id)
            var p = document.getElementById("content" + id)
            var gif = document.getElementById("deletegif" + id)
            var delBtn = document.getElementById("delete" + id)
            if (gif) {
                gif.classList.toggle("Invisible")
            }
    }
    var table = [editBtn, finishBtn, p, textarea,delBtn]
    table.forEach(function(el) {
        el.classList.toggle("Invisible")
    })
    params.action="UPDATE"
    params.id = id.toString()
    params.table = Table
    params.What = "state"
    params.newVal = "1"
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
            p.innerHTML = x.message
        })
        //gestion erreur
        .catch(x => console.log(x.json()))
}
//Permet d'éditer le post ou le commentaire
function Edit(id) {
    var editBtn = document.getElementById("edit" + id)
    var finishBtn = document.getElementById("validate" + id)
    var delBtn = document.getElementById("delete" + id)
    if (id != "") {
        var p = document.getElementById("content" + id)
        var gif = document.getElementById("deletegif" + id)
        if (gif) {
            gif.classList.toggle("Invisible")
        }
    } else {
        var p = document.getElementById("main-text" + id)
    }
    var textarea = document.getElementById("EditArea" + id)
    var table = [editBtn, finishBtn, p, textarea,delBtn]
    table.forEach(function(el) {
        el.classList.toggle("Invisible")
    })
}
//Permet de valider les modifications
function Validate(Table, id, p, textarea) {
    switch (Table) {
        case "posts":
            var textarea = document.getElementById('EditArea')
            var editBtn = document.getElementById("edit")
            var finishBtn = document.getElementById("validate")
            var delBtn = document.getElementById("delete")
            break
        case "comments":
            var textarea = document.getElementById('EditArea' + id)
            var editBtn = document.getElementById("edit" + id)
            var finishBtn = document.getElementById("validate" + id)
            var p = document.getElementById("content" + id)
            var gif = document.getElementById("deletegif" + id)
            var delBtn = document.getElementById("delete" + id)
            if (gif) {
                gif.classList.toggle("Invisible")
            }
    }
    var table = [editBtn, finishBtn, p, textarea,delBtn]
    table.forEach(function(el) {
        el.classList.toggle("Invisible")
    })
    var params = new Object()
    params.action="UPDATE"
    params.What = "content"
    params.id = id.toString()
    params.newVal = textarea.value
    params.table = Table
    fetch("/fetching", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Accept: "application/json"
            },
            body: JSON.stringify(params)
        }).then(x =>x.json())
        //Réponse qui me dit que la valeur est attribuer dans la base de donnée
        .then(function(x) {
            p.innerHTML = x.newVal
        })
        //gestion erreur
        .catch(x => console.log(x.json()))
}