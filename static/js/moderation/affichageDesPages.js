//Affiche seulement le tableau post et cache les autres. Chaque fonction est casi pareil, je cache les autres table SI elle sont visible et j'affiche celle si si elle ne l'est pas.
function showPosts() {
    if (!commentTable.className.includes("Invisible")) {
        commentTable.classList.add("Invisible")
        commentTable.classList.remove('Visible');
    }
    if (postTable.className.includes("Invisible")) {
        removeSelected();
        removeVisibility("postOption");
        postTable.classList.remove("Invisible")
        postTable.classList.add('Visible');
        if (actualTable != "posts") {
            actualTable = "posts"
            if (urlParams.get('page') == null) {
                history.pushState({}, null, "/moderation&table=posts");
            } else {
                history.pushState({}, null, "/moderation?page=" + urlParams.get('page') + "&table=posts");
            }
        }
    }
    if (!userTable.className.includes("Invisible")) {
        userTable.classList.add("Invisible")
        userTable.classList.remove('Visible');
    }
    if (!categoryTable.className.includes("Invisible")) {
        categoryTable.classList.add("Invisible")
        categoryTable.classList.remove('Visible');
    }
}
//Affiche seulement le tableau comments et cache les autres. Chaque fonction est casi pareil, je cache les autres table SI elle sont visible et j'affiche celle si si elle ne l'est pas.
function showComments() {
    if (commentTable.className.includes("Invisible")) {
        removeSelected();
        removeVisibility("commentOption");
        commentTable.classList.remove("Invisible")
        commentTable.classList.add('Visible');
        if (actualTable != "comments") {
            actualTable = "comments"
            if (urlParams.get('page') == null) {
                history.pushState({}, null, "/moderation&table=comments");
            } else {
                history.pushState({}, null, "/moderation?page=" + urlParams.get('page') + "&table=comments");
            }
        }
    }
    if (!postTable.className.includes("Invisible")) {
        postTable.classList.add("Invisible")
        postTable.classList.remove('Visible');
    }
    if (!userTable.className.includes("Invisible")) {
        userTable.classList.add("Invisible")
        userTable.classList.remove('Visible');
    }
    if (!categoryTable.className.includes("Invisible")) {
        categoryTable.classList.add("Invisible")
        categoryTable.classList.remove('Visible');
    }
}
//Affiche seulement le tableau users et cache les autres. Chaque fonction est casi pareil, je cache les autres table SI elle sont visible et j'affiche celle si si elle ne l'est pas.
function showUsers() {
    if (!commentTable.className.includes("Invisible")) {
        commentTable.classList.add("Invisible")
        commentTable.classList.remove('Visible');
    }
    if (!postTable.className.includes("Invisible")) {
        postTable.classList.add("Invisible")
        postTable.classList.remove('Visible');
    }
    if (userTable.className.includes("Invisible")) {
        removeSelected();
        removeVisibility("userOption");
        userTable.classList.remove("Invisible")
        userTable.classList.add('Visible');
        if (actualTable != "users") {
            actualTable = "users"
            if (urlParams.get('page') == null) {
                history.pushState({}, null, "/moderation&table=users");
            } else {
                history.pushState({}, null, "/moderation?page=" + urlParams.get('page') + "&table=users");
            }
        }
    }
    if (!categoryTable.className.includes("Invisible")) {
        categoryTable.classList.add("Invisible")
        categoryTable.classList.remove('Visible');
    }
}
//Affiche seulement le tableau post et cache les autres. Chaque fonction est casi pareil, je cache les autres table SI elle sont visible et j'affiche celle si si elle ne l'est pas.
function showCategory() {
    if (!commentTable.className.includes("Invisible")) {
        commentTable.classList.add("Invisible")
        commentTable.classList.remove('Visible');
    }
    if (!postTable.className.includes("Invisible")) {
        postTable.classList.add("Invisible")
        postTable.classList.remove('Visible');
    }
    if (!userTable.className.includes("Invisible")) {
        userTable.classList.add("Invisible")
        userTable.classList.remove('Visible');
    }
    if (categoryTable.className.includes("Invisible")) {
        removeSelected();
        removeVisibility("postOption");
        categoryTable.classList.remove("Invisible")
        categoryTable.classList.add('Visible');
        if (actualTable != "categories") {
            actualTable = "categories"
            if (urlParams.get('page') == null) {
                history.pushState({}, null, "/moderation&table=categories");
            } else {
                history.pushState({}, null, "/moderation?page=" + urlParams.get('page') + "&table=categories");
            }
        }
    }
}