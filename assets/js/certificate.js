import {Get, Post, Patch} from './requests.js';
import {updateStatus} from './status.js';

$('#certificates').ready(() => {
  // debugger;
  get_active_cert();
});


function get_active_cert() {
  Get("/certificate/list"+window.location.search)
  .then((response) => {
    response.json()
      .then((data) => {
        let certRows = document.getElementById("certificate_rows");
        if (certRows) {
          certRows.innerText = "";
        }
        for (let row of data) {
          const el = "<tr>"
            + "<td>" + row.id + "</td>"
            + "<td>" + row.key_path + "</td>"
            + "<td>" + row.body_path + "</td>"
            + "<td id='"+row.id+"'><a href='#'>Deactivate</a></td>"
            + "</tr>";
          $("#certificate_rows").prepend(el);
          var elObj = document.getElementById(row.id);
          elObj.addEventListener('click', deactivate_cert.bind({cert_id: row.id}));
        }
      });
  })
}

function create_cert() {

}

function download_key(id) {

}

function download_cert(id) {

}

export function deactivate_cert() {
  const url = "/certificate/" + this.cert_id + "/update?active=false";
  Patch(url)
    .then((response) => {
      const status = response.status;
      if (status == 200) {
        updateStatus("Certificate deactivated")
        get_active_cert();
      } else {
        response.json()
          .then((data) => {
            updateStatus(data)
          })
      }
    })
}