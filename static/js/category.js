var el = document.getElementsByClassName("super")[0]
var posts = document.getElementsByClassName("column")
//Converti la date sql en date js
function convertDate(date) {
    let match = date.split('.')[0].match(/^(\d+)-(\d+)-(\d+) (\d+)\:(\d+)\:(\d+)$/)
    return new Date(match[1], match[2] - 1, match[3], match[4], match[5], match[6]).getTime()/1000
}
//trie les post par date croissant ou décroissant
function sortByDate(){
    alldata = document.querySelectorAll("[data-date]")
    toSort = Array.prototype.slice.call(alldata, 0);
    
    if(toSort[toSort.length-1].dataset.date<toSort[0].dataset.date){
        toSort.sort(function(a, b) {
            return (convertDate(a.dataset.date)) - (convertDate(b.dataset.date));
            });
    }else{
        toSort.sort(function(a, b) {
            return (convertDate(b.dataset.date)) - (convertDate(a.dataset.date));
            });
    }
    for(var i = 0, l = toSort.length; i < l; i++) {
        el.appendChild(toSort[i]);
    }
}
//Filtre pas post le plus commenter ou le moins commenter
function sortByComment(){
    alldata = document.querySelectorAll("[data-comment]")
    toSort = Array.prototype.slice.call(alldata, 0);
    if(toSort[toSort.length-1].dataset.comment<toSort[0].dataset.comment){
        toSort.sort(function(a, b) {
            return a.dataset.comment - b.dataset.comment;
            });
    }else{
        toSort.sort(function(a, b) {
            return  b.dataset.comment - a.dataset.comment;
            });
    }
    for(var i = 0, l = toSort.length; i < l; i++) {
        el.appendChild(toSort[i]);
    }
}
//Filtre par post le plus liké ou le moins liker
function sortByLike(){
    alldata = document.querySelectorAll("[data-like]")
    toSort = Array.prototype.slice.call(alldata, 0);
    if(toSort[toSort.length-1].dataset.like<toSort[0].dataset.like){
        toSort.sort(function(a, b) {
            return a.dataset.like - b.dataset.like;
            });
    }else{
        toSort.sort(function(a, b) {
            return  b.dataset.like - a.dataset.like;
            });
    }
    for(var i = 0, l = toSort.length; i < l; i++) {
        el.appendChild(toSort[i]);
    }
}
//Filtres les posts sur la page pour un mot clé rentré
function filter(val){
    Array.prototype.forEach.call(posts, function(e){
        if(val=="" && e.className.includes("invisible")){
            e.classList.remove("invisible");
        }else{
            if(e.dataset.content!=undefined){
                if(e.dataset.content.toLowerCase().includes(val.toLowerCase()) && e.className.includes("invisible")){
                    e.classList.remove("invisible");
                }else if(!e.dataset.content.toLowerCase().includes(val.toLowerCase()) && !e.className.includes("invisible")){
                    e.classList.add("invisible")
                }
            }else{
                e.classList.add("invisible")
            }
            
        }
    })
}