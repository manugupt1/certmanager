
export function download(filename, data) {
  let blob = new Blob([JSON.stringify(data)], {type: 'application/json'});
  let a = document.createElement("a");
  a.style = "display: none";
  document.body.appendChild(a);
  let url = window.URL.createObjectURL(blob);
  a.href = url;
  a.download = filename;
  a.click();
  window.URL.revokeObjectURL(url);
}