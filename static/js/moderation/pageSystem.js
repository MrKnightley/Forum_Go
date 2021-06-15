//----------------------------------INIT----------------------------------\\\

                    //Je chope l'URL ici
                    const queryString = window.location.search;
                    const urlParams = new URLSearchParams(queryString);
                    //Je chope la table si elle est dans l'URL
                    actualTable = urlParams.get('table')
                    if (actualTable == "users") {
                        showUsers()
                    } else if (actualTable == "posts") {
                        showPosts()
                    } else if (actualTable == "comments") {
                        showComments()
                    } else if (actualTable == "categories") {
                        showCategory()
                    }
                    //Je chope la page si elle est dans l'URL
                    const pg = urlParams.get('page')
                    if (pg != null) {
                        setOffset((parseInt(pg, 10) - 1) * 20)
                    }
                    //Permet d'initialiser l'url de la page si on a pas encore l'url custom
                    if (urlParams.get('page') == null && actualTable == null) {
                        window.location.href = "/moderation?page=1&table=null"
                    }

                    //Offset de chaque catégories (pour le systeme de page pour chaque tableau)
                    var UserOffset = 0
                    var PostOffset = 0
                    var CommentOffset = 0
                    var CategoryOffset = 0


//----------------------------------Set Page----------------------------------\\\

            //Me permet de savoir quel offset utiliser pour le tableau actuel
            function getOffSet() {
                switch (actualTable) {
                    case "users":
                        return UserOffset
                    case "comments":
                        return CommentOffset
                    case "posts":
                        return PostOffset
                        break
                    case "categories":
                        return CategoryOffset
                }
            }
           //Me permet de mettre a jour l'offset que j'utilise             
            function setOffset(i) {
                switch (actualTable) {
                    case "users":
                        UserOffset += i
                        break
                    case "comments":
                        CommentOffset += i
                        break
                    case "posts":
                        PostOffset += i
                        break
                    case "categories":
                        CategoryOffset += i
                }
            }
            //Affiche la prochaine page
            function NextPage() {
                let arr
                actualOffSet = getOffSet()
                //Me permet de récuperer le tableau ou je veux changer la page
                switch (actualTable) {
                    case "users":
                        arr = AllTrUser
                        break
                    case "comments":
                        arr = AllTrComment
                        break
                    case "posts":
                        arr = AllTrPost
                        break
                    case "categories":
                        arr = AllTrCategory
                }
                if (actualOffSet < arr.length) {
                    //Rend Invisible la page d'avant si il y en a une
                    if (actualOffSet > 0) {
                        for (var i = actualOffSet - 20; i <= actualOffSet && i < arr.length; i++) {
                            arr[i].classList.add("Invisible")
                            arr[i].classList.remove('Visible');
                        }
                    }
                    //Rend visible l'équivalent de la page d'après
                    for (var i = actualOffSet; i <= actualOffSet + 19 && i < arr.length; i++) {
                        arr[i].classList.remove("Invisible")
                        arr[i].classList.add('Visible');
                    }
                    //Change le liens pour le partager
                    if (actualTable != "") {
                        history.pushState({}, null, "/moderation?page=" + ((actualOffSet / 20) + 1).toString() + "&table=" + actualTable);
                    } else {
                        history.pushState({}, null, "/moderation?page=" + ((actualOffSet / 20) + 1).toString());
                    }
                    //Change l'offset actuel
                    setOffset(20)
                }
            }
            //Affiche la page précédente
            function PreviousPage() {
                let arr
                actualOffSet = getOffSet()
                //Récupère le tableau que je veut changer
                switch (actualTable) {
                    case "users":
                        arr = AllTrUser
                        break
                    case "comments":
                        arr = AllTrComment
                        break
                    case "posts":
                        arr = AllTrPost
                        break
                    case "categories":
                        arr = AllTrCategory
                }
                //Si il y a une page avant alors je masque la page actuelle
                if (actualOffSet > 20) {
                    for (var i = actualOffSet; i >= actualOffSet - 20; i--) {
                        if (i < arr.length) {
                            arr[i].classList.add("Invisible")
                            arr[i].classList.remove('Visible');
                        }
                    }
                    //J'affiche la précedente
                    for (var i = actualOffSet - 20; i >= actualOffSet - 40 && i < arr.length; i--) {
                        arr[i].classList.remove("Invisible")
                        arr[i].classList.add('Visible');
                    }
                    //Je Change le liens
                    if (actualTable != "") {
                        history.pushState({}, null, "/moderation?page=" + ((actualOffSet / 20) - 1).toString() + "&table=" + actualTable);
                    } else {
                        history.pushState({}, null, "/moderation?page=" + ((actualOffSet / 20) - 1).toString());
                    }
                    setOffset(-20)
                }
            }


//Initilalise la première page de chaque table
function initialize() {
    temp = actualTable
    actualTable = "users"
    NextPage()
    actualTable = "posts"
    NextPage()
    actualTable = "comments"
    NextPage()
    actualTable = "categories"
    NextPage()
    actualTable = temp
}
initialize()       