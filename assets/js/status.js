export const updateStatus = (obj) => {
  if (typeof obj == "string") {
    document.getElementById("status").innerText = obj;
    return;
  } else {
    let status = "";
    for (var key in obj) {
      status += "<li>" + obj[key][0] + "</li>";
      document.getElementById("status").innerHTML = status;
    }
  }
}