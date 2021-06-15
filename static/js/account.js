// var password1 = document.getElementById("pwd-1");
// var password2 = document.getElementById("pwd-2");
// var username = document.getElementById("username")

// let capitals = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// let lowers = "abcdefghijklmnopqrstuvwxyz"
// let nums = "0123456789"

// // Vérification du match des 2 mots de passe :
// function checkPasswords() {
//     if (password1.value != password2.value) {
//         password2.setCustomValidity("Your passwords do not match.")
//         changeToRed(password1)
//         changeToRed(password2)
//     } else {
//         password2.setCustomValidity("")
//         changeToNormal(password1)
//         changeToNormal(password2)
//     }
// }

// password1.onchange = checkPasswords;
// password2.onkeyup = checkPasswords;

// // Vérification du pattern du mot de passe :
// function hasCapitals(str) {
//     for (let i = 0; i < capitals.length; i++) {
//         if (str.includes(capitals[i]))
//             return true
//     }
//     return false
// }

// function hasLowers(str) {
//     for (let i = 0; i < lowers.length; i++) {
//         if (str.includes(lowers[i]))
//             return true
//     }
//     return false
// }

// function hasNums(str) {
//     for (let i = 0; i < nums.length; i++) {
//         if (str.includes(nums[i]))
//             return true
//     }
//     return false
// }

// function checkPattern() {
//     if (!hasCapitals(password1.value) || !hasLowers(password1.value) || !hasNums(password1.value) || password1.value.length < 6 || password1.value.length > 30) {
//         password1.setCustomValidity("Your password must contain at least one lowercase character, one uppercase character, one number and be longer than 6 and shorter than 30.")
//     } else {
//         password1.setCustomValidity("")
//     }
// }

// password1.onchange = checkPattern;
// password1.onkeyup = checkPattern;


// // Vérification que le username soit valide (uniquement lettres/chiffres/espaces/underscore, et pas 2 espaces consécutifs) :
// function checkRegex() {
//     for (let i = 0; i < username.value.length - 1; i++) {
//         if (!capitals.includes(username.value[i]) && !lowers.includes(username.value[i]) && !nums.includes(username.value[i]) && username.value[i] != " " && username.value[i] != "_" && username.value[i] != ".") {
//             return false
//         }
//     }
//     return true
// }

// function checkConsecutiveSpaces() {
//     for (let i = 0; i < username.value.length - 1; i++) {
//         if (username.value[i] == " " && username.value[i + 1] == " ") {
//             return false
//         }
//     }
//     return true
// }

// function isValidUsername() {
//     if (!checkRegex()) {
//         username.setCustomValidity("Your username can only contain letters, numbers, blank spaces, dots and underscores.")
//         changeToRed(username, "pattern")
//     } else if (!checkConsecutiveSpaces()) {
//         username.setCustomValidity("Your username cannot contain several consecutive blank spaces.")
//         changeToRed(username)
//     } else if (username.value.length < 3 || username.value.length > 20) {
//         username.setCustomValidity("Your username must be longer than 3 and shorter than 20 characters.")
//         changeToRed(username, "length")
//     } else {
//         username.setCustomValidity("")
//         changeToNormal(username)
//     }
// }

// function changeToRed(element, str) {
//     element.style.borderTop = "1px solid rgb(233, 0, 31)";
//     element.style.borderBottom = "1px solid rgb(233, 0, 31)";
//     element.style.borderLeft = "1px solid rgb(233, 0, 31)";
//     element.style.borderRight = "1px solid rgb(233, 0, 31)";
//     element.style.boxShadow = "0 0 5px rgb(214, 18, 44)";
//     if (element == username) {
//         // Désactivation du bouton 'Continue'
//         disables(button)
//         if (str == "length") {
//             document.getElementsByClassName("invalid-username")[0].style.display = "block"
//             document.getElementsByClassName("invalid-username")[1].style.display = "none"
//         }
//         if (str == "pattern") {
//             document.getElementsByClassName("invalid-username")[1].style.display = "block"
//             document.getElementsByClassName("invalid-username")[0].style.display = "none"
//         }
//     }
//     if (element == password1) {
//         document.getElementById("invalid-password").style.display = "block"
//         disables(button)
//     }
// }

// function changeToNormal(element) {
//     enables(button)
//     element.style.borderTop = "1px solid rgb(66, 66, 66)";
//     element.style.borderBottom = "1px solid rgb(66, 66, 66)";
//     element.style.borderLeft = "1px solid rgb(66, 66, 66)";
//     element.style.borderRight = "1px solid rgb(66, 66, 66)";
//     element.style.boxShadow = "none";
//     if (element == username) {
//         document.getElementsByClassName("invalid-username")[0].style.display = "none"
//         document.getElementsByClassName("invalid-username")[1].style.display = "none"
//     }
//     if (element == password1) {
//         document.getElementById("invalid-password").style.display = "none"
//     }
// }

// function disables(element) {
//     // Désactivation du bouton "Continue" :
//     element.disabled = true
//     element.style.opacity = 0.4
//     element.style.pointerEvents = "none"
// }

// function enables(element) {
//     // Activation du bouton "Continue" :
//     element.disabled = false
//     element.style.opacity = 1
//     element.style.pointerEvents = "initial"
// }

// username.onchange = isValidUsername;
// username.onkeyup = isValidUsername;

// // Activation des div "requirement" lorsque le mot de passe est valide :
// let letterTyped = "";
// let requirements = document.getElementsByClassName("requirement")
// let ticks = document.querySelectorAll("svg")
// let button = document.querySelector("button")

// document.getElementById("pwd-1").addEventListener("input", (e) => {
//     letterTyped = e.target.value;
//     if (!isLongEnough()) {
//         // Désactivation de la 1ère div :
//         requirements[0].style.opacity = 0.35
//         ticks[0].style.opacity = 0
//     } else {
//         // Activation de la 1ère div :
//         requirements[0].style.opacity = 1
//         ticks[0].style.opacity = 1
//     }
//     if (!hasCorrectPattern()) {
//         // Désactivation de la 2nde div :
//         requirements[1].style.opacity = 0.35
//         ticks[1].style.opacity = 0
//     } else {
//         // Activation de la 2nde div :
//         requirements[1].style.opacity = 1
//         ticks[1].style.opacity = 1
//     }
//     if (!isLongEnough() || !hasCorrectPattern()) {
//         // Désactivation du bouton "Continue" :
//         disables(button)
//     } else {
//         // Activation du bouton "Continue" :
//         enables(button)
//     }
// });

// function hasCorrectPattern() {
//     if (!hasCapitals(letterTyped) || !hasLowers(letterTyped) || !hasNums(letterTyped)) {
//         return false
//     } else {
//         return true
//     }
// }

// function isLongEnough() {
//     if (letterTyped.length < 6 || letterTyped.length > 30) {
//         return false
//     } else {
//         return true
//     }
// }


function Change(cat,val,id){
    var params = new Object()
    params.newVal = val.value
    params.cat = cat
    params.id = id.toString()
    fetch("/edit-account", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Accept: "application/json"
        },
        body: JSON.stringify(params)
    }).then(x=>x.json())
    //Réponse qui me dit que la valeur est attribuer dans la base de donnée
    .then(function(x){
        val.value = x.newVal
    })
    //gestion erreur
    .catch(x=>console.log(x))    
}