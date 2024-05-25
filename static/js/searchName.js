var artistUrl = document.getElementById("artist-url")
var arUrl = artistUrl.getAttribute("jsVarArtistPath");

//------------------------------------------
function displayArtistUrl() {
    let inputValue = document.getElementById("input-value").value.trim();
    if (inputValue !== "") {
        let url = arUrl + encodeURIComponent(inputValue);
       window.location.href = url;
    } else {
        alert("Please enter an artist name.");
    }
}
//------------------------------------------ 
// Function to generate hyperlink dynamically based on user input
function generateArtistLink() {

    let inputValue = document.getElementById("input-value").value.trim();

    if (inputValue !== "") {

        let link = document.createElement('a');
        link.href = arUrl + encodeURIComponent(inputValue);
       // link.textContent = inputValue + "-" + "Link";
        //link.style.color = "red"
       // document.getElementById("display").innerHTML = '';
       // document.getElementById("display").appendChild(link);
    } 
    //else {
       // document.getElementById("display").innerHTML = '';
    //}
}
//------------------------------------------




//------------------------------------------ 
// Event listener to call the function on input change
document.getElementById("input-value").addEventListener("input", generateArtistLink);

// Function to handle Enter key press event
document.getElementById("input-value").addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        displayArtistUrl();
    }
});
//------------------------------------------