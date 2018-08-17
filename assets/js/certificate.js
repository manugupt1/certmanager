import {Get, Post, Patch} from './requests.js';
import {updateStatus} from './status.js';
import { download } from './download.js';


$('#certificates').ready(() => {
  $('#create_cert').click(function () {
    create_cert();
  })
  get_active_cert();
});


function get_cust_id() {
  const url = new URL(window.location.href)
  const cust_id = url.searchParams.get('cust_id')
  if (!cust_id) {
    updateStatus("Needs a customer id in the url")
    throw "Needs a customer id in the url"
  }
  return cust_id;
}

function render_action_button(row) {
  if (row.activated != true) {
    return "<td id='"+row.id+"'><a href='#'>Activate</a></td>"
  }
  return "<td id='"+row.id+"'><a href='#'>Deactivate</a></td>"
}

function attach_action_listener(row) {
  let cust_id = get_cust_id();
  let el = document.getElementById(row.id);
  if (row.activated) {
    el.addEventListener('click', () => {
      update_cert(cust_id, row.id, false);
    })
  } else {
    el.addEventListener('click', () => {
      update_cert(cust_id, row.id, true);
    });
  }

}

function get_active_cert() {
  let cust_id = get_cust_id();
  Get("/customer/"+cust_id + "/certificates")
  .then((response) => {
    if (response.status != 200) {
      updateStatus("Some error occured!")
      return;
    }
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
            + render_action_button(row);
            + "</tr>";
          $("#certificate_rows").prepend(el);
          attach_action_listener(row)
          let keyDownloadEl = document.getElementById(row.key_path);
          keyDownloadEl.addEventListener('click', () => {
            download_key(cust_id, row.id, row.key_path);
          });

          let bodyDownloadEl = document.getElementById(row.body_path);
          bodyDownloadEl.addEventListener('click', () => {
            download_cert(cust_id, row.id, row.body_path);
          });
        }
      });
  })
}

function create_cert() {
  let cust_id = get_cust_id();
  if (!cust_id || isNaN(cust_id)) {
    updateStatus("cust_id can't be empty and must be an integer")
    return;
  }
  const createURL = "/customer/" + cust_id + "/certificate";
  Post(createURL)
    .then((response) => {
      const status = response.status;
      if (status == 200) {
        updateStatus("Certificate successfully created!");
        get_active_cert();
      } else {
        response.json()
          .then((data) => {
            updateStatus(data);
          });
      }
    })
}

function download_key(cust_id, cert_id, key) {
  const url = "/customer/" + cust_id + "/certificate/" + cert_id + "/key"
  Get(url).then((response) => {
    const status = response.status;
    if (status == 200) {
      response.json()
        .then((data) => {
          download(key, data);
        })
    } else {
      response.json()
        .then((data) => {
          updateStatus(data);
        })
    }
  })
}

function download_cert(cust_id, cert_id, body) {
  const url = "/customer/" + cust_id + "/certificate/" + cert_id + "/body";
  Get(url).then((response) => {
    const status = response.status;
    if (status == 200) {
      response.json()
        .then((data) => {
          download(body, data);
        })
    } else {
      response.json()
        .then((data) => {
          updateStatus(data);
        })
    }
  })
}


function update_cert(cust_id, cert_id, active) {
  const url = "/customer/" + cust_id + "/certificate/" + cert_id + "?active="+active;
  Patch(url)
    .then((response) => {
      const status = response.status;
      if (status == 200) {
        if (active) {
          updateStatus("Certificate deactivated")
        } else {
          updateStatus("Certificate activated")
        }
        get_active_cert();
      } else {
        response.json()
          .then((data) => {
            updateStatus(data)
          })
      }
    })
}

