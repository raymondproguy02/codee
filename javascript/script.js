// Counter Program

incrementBtn = document.getElementById("incrementBtn");
resetBtn = document.getElementById("resetBtn");
decrementBtn = document.getElementById("decrementBtn");
countLabel = document.getElementById("countLabel");
count = 0;

incrementBtn.onclick = function() {
    count++;
    countLabel.textContent = count;
}

decrementBtn.onclick = function() {
    count--;
    countLabel.textContent = count;
}

resetBtn.onclick = function() {
    count = 0;
    countLabel.textContent = count;
}
