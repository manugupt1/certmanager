import {Get, Post, Patch} from './requests.js';
import {updateStatus} from './status.js';
import { download } from './download.js';

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
            + "<td id='"+row.key_path+"'><a href='#'>" + row.key_path + "</a></td>"
            + "<td id='"+row.body_path+"'><a href='#'>" + row.body_path + "</a></td>"
            + "<td id='"+row.id+"'><a href='#'>Deactivate</a></td>"
            + "</tr>";
          $("#certificate_rows").prepend(el);
          let deactEl = document.getElementById(row.id);
          deactEl.addEventListener('click', deactivate_cert.bind({cert_id: row.id}));

          let keyDownloadEl = document.getElementById(row.key_path);
          keyDownloadEl.addEventListener('click', download_key.bind({key: row.key_path}));

          let bodyDownloadEl = document.getElementById(row.body_path);
          bodyDownloadEl.addEventListener('click', download_cert.bind({body: row.body_path}));
        }
      });
  })
}

function create_cert() {

}

function download_key() {
  const url = "/certificate/key/" + this.key;
  Get(url).then((response) => {
    const status = response.status;
    if (status == 200) {
      response.json()
        .then((data) => {
          download(this.key, data);
        })
    } else {
      response.json()
        .then((data) => {
          updateStatus(data);
        })
    }
  })
}

function download_cert() {
  const url = "/certificate/body/" + this.body;
  Get(url).then((response) => {
    const status = response.status;
    if (status == 200) {
      response.json()
        .then((data) => {
          download(this.body, data);
        })
    } else {
      response.json()
        .then((data) => {
          updateStatus(data);
        })
    }
  })
}


function deactivate_cert() {
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