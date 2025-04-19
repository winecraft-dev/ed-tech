let table = document.getElementsByClassName("KN3Kfc")[0];
let rows = table.getElementsByTagName("tr");
let text = "";
for(let row of rows) {
    let a = row.getElementsByTagName("a")[0];
    if(a != null) {
        text += a.innerHTML + "\n";
    }
}
console.log(text);