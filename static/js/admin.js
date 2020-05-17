let tableRef = document.getElementById("tableBuzzers");
let resetButtonRef = document.getElementById("buttonReset");

let url = "/data/buzzer"
let httpRead = new XMLHttpRequest();
httpRead.onreadystatechange = function(){
    if(this.readyState === 4 && this.status === 200) {
        let myArr = JSON.parse(this.responseText);
        updateData(myArr);
        drawTabel();
    }
}

let httpDelete = new XMLHttpRequest();

window.setInterval(function () {
    httpRead.open("GET", url, true);
    httpRead.send();
}, 2500)

resetButtonRef.onclick = function () {
    deleteData();
}

function deleteData() {
    localStorage.setItem('buzzers', null);
    httpDelete.open("DELETE", url);
    httpDelete.send();
}

function updateData(dataArray) {
    localStorage.setItem('buzzers', JSON.stringify(dataArray));
}

function drawTabel() {
    let data = JSON.parse(localStorage.getItem("buzzers"));
    let numberOfRows = tableRef.rows.length - 1 //DON'T remove the heading

    for (let row = numberOfRows; row > 0; row--) {
        tableRef.deleteRow(row);
    }

    data.forEach(function (value, index) {
        let row = tableRef.insertRow(index + 1);
        let cell0 = row.insertCell(0);
        let cell1 = row.insertCell(1);
        let cell2 = row.insertCell(2);
        cell0.innerHTML = value.BuzzTime;
        cell1.innerHTML = value.PlayerName;
        cell2.innerHTML = value.TeamName;
    })
}