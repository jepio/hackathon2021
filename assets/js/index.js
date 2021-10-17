document.getElementById("button-id").addEventListener("click", async function() {
    var content = document.getElementById("yaml").value;
    var output = document.getElementById("json");
    var response = await fetch("https://clctranspiler.azurewebsites.net/api/clctranspilerfunction", {
        method: "POST",
        body: content
    })
    var text = await response.text();
    output.value = text
});

window.addEventListener("load", async function() {
    var preload = await fetch("https://clctranspiler.azurewebsites.net/api/clctranspilerfunction")
});
