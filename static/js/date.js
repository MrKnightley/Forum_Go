// ========= ⭐ FORMATTAGE DE LA DATE DU JOUR ⭐ =========

// Date d'aujourd'hui au format '2021-06-03T16:10:00.891Z' :
let today = new Date(Date.now());

// Date d'aujourd'hui au format 'Thu Jun 03 2021' :
let date = today.toDateString();

// array = ['Thu', 'Jun', '03', '2021'] :
let array = date.split(' ');

let day = array[0]; // Thu
let month = array[1]; // Jun
let num = array[2]; // 03
let year = array[3]; // 2021

// Formattage du jour de la semaine :
switch (day) {
    case "Mon":
        day = "Monday";
        break;
    case "Tue":
        day = "Tuesday";
        break;
    case "Wed":
        day = "Wednesday";
        break;
    case "Thu":
        day = "Thursday";
        break;
    case "Fri":
        day = "Friday";
        break;
    case "Sat":
        day = "Saturday";
        break;
    case "Sun":
        day = "Sunday";
        break;
}

// Formattage du mois :
function formatMonth(month) {
    switch (month) {
        case "Jan":
        case "01":
            month = "January";
            break;

        case "Feb":
        case "02":
            month = "February";
            break;

        case "Mar":
        case "03":
            month = "March";
            break;

        case "Apr":
        case "04":
            month = "April";
            break;

        case "May":
        case "05":
            month = "May";
            break;

        case "Jun":
        case "06":
            month = "June";
            break;

        case "Jul":
        case "07":
            month = "July";
            break;

        case "Aug":
        case "08":
            month = "August";
            break;

        case "Sep":
        case "09":
            month = "September";
            break;

        case "Oct":
        case "10":
            month = "October";
            break;

        case "Nov":
        case "11":
            month = "November";
            break;

        case "Dec":
        case "12":
            month = "December";
            break;
    }
    return month;
}

// Formattage du jour (nombre) :
function formatNum(num) {

    if (num[0] == '0') { // Si le jour commence par '0', c'est qu'il est compris entre 01 et 09.
        num = num[1]; // Donc on enlève le '0' et on ne garde que le 2nd chiffre.

        // Ajout du suffixe :
        switch (num) {
            case "1":
                num += `<span>st</span>`;
                break;
            case "2":
                num += `<span>nd</span>`;
                break;
            case "3":
                num += `<span>rd</span>`;
                break;
            default:
                num += `<span>th</span>`;
        }
    }
    return num
}


let fullDate = day + ' ' + formatMonth(month) + ' ' + formatNum(num) + ', ' + year;

// Envoi de fullDate dans les div dont la classe est 'today-date' :
let todayDates = document.getElementsByClassName('today-date');

for (let i = 0; i < todayDates.length; i++) {
    todayDates[i].innerHTML = fullDate;
};

// ========= ⭐ FORMATTAGE DE LA DATE DES POSTS ⭐ =========

let dates = document.getElementsByClassName('date');

for (let i = 0; i < dates.length; i++) {
    let raw_date = dates[i].textContent; // Date au format '2021-06-04 00:47:16.6486277 +0200 +0200'

    let arr = raw_date.split(' '); // arr = ['2021-06-04', '00:47:16.6486277', '+0200', '+0200']
    let calendarDate = arr[0] // '2021-06-04'
    let time = arr[1].split('.')[0] // '00:47:16'

    let calendarDateArr = calendarDate.split('-')

    let result = "on " + formatMonth(calendarDateArr[1]) + " " + formatNum(calendarDateArr[2]) + ", " + calendarDateArr[0] + " at " + time;
    dates[i].innerHTML = result;
};


// Multiplication du contenu des posts par 20 (A SUPPRIMER PLUS TARD) :

// let contents = document.getElementsByClassName("content")
// for (let i = 0; i < contents.length; i++) {
//     contents[i].textContent = contents[i].textContent.repeat(20);
// };