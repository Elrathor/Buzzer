let table = document.getElementById("tableBuzzers");
let url = "/data/buzzer"
let xmlhttp = new XMLHttpRequest();
xmlhttp.onreadystatechange = function(){
    if(this.readyState === 4 && this.status === 200) {
        let myArr = JSON.parse(this.responseText);
        updateData(myArr);
        drawTabel();
    }
}

window.setInterval(function () {
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
}, 2500)

function updateData(dataArray) {
    localStorage.setItem('buzzers', JSON.stringify(dataArray));
}

function drawTabel() {
    let data = JSON.parse(localStorage.getItem("buzzers"));
    let numberOfRows = table.rows.length - 1 //DON'T remove the heading

    for (let row = numberOfRows; row > 0; row--) {
        table.deleteRow(row);
    }

    data.forEach(function (value, index) {
        let row = table.insertRow(index + 1);
        let cell0 = row.insertCell(0);
        let cell1 = row.insertCell(1);
        let cell2 = row.insertCell(2);
        cell0.innerHTML = value.BuzzTime;
        cell1.innerHTML = value.PlayerName;
        cell2.innerHTML = value.TeamName;
    })
}