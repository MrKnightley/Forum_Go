function handleClick(question) {

    switch (question) {
        case "q1":
            document.getElementById('r2').checked = true;
            break;
        case "q2":
            document.getElementById('r3').checked = true;
            break;
        case "q3":
            document.getElementById('r4').checked = true;
            break;
        case "q4":
            document.getElementById('r5').checked = true;
            break;
        case "q5":
            document.getElementById('r6').checked = true;
            break;
        case "q6":
            document.getElementById('r7').checked = true;
            break;
        case "q7":
            document.getElementById('r8').checked = true;
            break;
        case "q8":
            document.getElementById('r9').checked = true;
            break;
        case "q9":
            document.getElementById('r10').checked = true;
            break;
        case "q10":
            document.getElementById('r11').checked = true;
            break;
        case "q11":
            document.getElementById('r12').checked = true;
            break;
        case "q12":
            document.getElementById('r13').checked = true;
            break;
        case "q13":
            document.querySelector("form").submit();
            console.log("Form submitted.")
            break;
    }
}