function showAlert() {
    alert("Are you sure about that?");
}

document.getElementById("myButton").addEventListener("click", showAlert);
document.getElementById("prj1").addEventListener("click", function() {
    window.location.href = "https://github.com/NFLSTY/LibraryPlus.git";
});

document.getElementById("prj2").addEventListener("click", function() {
    window.location.href = "https://github.com/NFLSTY/BookManagement.git";
});

document.getElementById("prj3").addEventListener("click", function() {
    window.location.href = "https://github.com/NFLSTY/MoneySystem.git";
});