playerNameInputRef = document.getElementById("inputPlayerName");
teamNameInputRef = document.getElementById("inputTeamName");
buzzerButtonRef = document.getElementById("buttonBuzzer");

// Prepare request client
let url = "/data/buzzer"
let xmlhttp = new XMLHttpRequest();

// Prefill data from localstorage if available
if(localStorage.getItem("playerName") !== null){
    playerNameInputRef.value = localStorage.getItem("playerName");
}

if(localStorage.getItem("teamName") !== null){
    teamNameInputRef.value = localStorage.getItem("teamName");
}

// Store values in localstorage
playerNameInputRef.onchange = function (event) {
    localStorage.setItem("playerName", playerNameInputRef.value);
};

teamNameInputRef.onchange = function (event) {
    localStorage.setItem("teamName", teamNameInputRef.value);
};


buzzerButtonRef.onclick = function (event) {
    let data = {
        PlayerName : playerNameInputRef.value,
        TeamName: teamNameInputRef.value,
        BuzzTime: new Date(),
    }

    xmlhttp.open("POST", url, true);
    xmlhttp.setRequestHeader('Content-Type', 'application/json');
    xmlhttp.send(JSON.stringify(data));
}
