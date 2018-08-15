import {Get, Post, Patch} from './requests.js';
import {updateStatus} from './status.js';

$('#certificates').ready(() => {
  get_active_cert();
});

function get_active_cert() {
  Get("/certificate/list"+window.location.search)
  .then((response) => {
    response.json()
      .then((data) => {
        let certRows = document.getElementById("#certificate_rows");
        if (certRows) {
          certRows.innerText = "";
        }
        for (let row of data) {
          const el = "<tr>"
            + "<td>" + row.id + "</td>"
            + "<td>" + row.key_path + "</td>"
            + "<td>" + row.body_path + "</td>"
            + "<td><a>Deactivate</a></td>"
            + "</tr>";
          $("#certificate_rows").prepend(el)
        }
      });
  })
  console.log("get active cert");
}

function create_cert() {

}

function download_key(id) {

}

function download_cert(id) {
  
}

export function deactivate_cert() {

}

