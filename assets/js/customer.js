

import {Get, Post, Delete} from './requests.js';
import {updateStatus} from './status.js';

export function get_customer() {
  Get("/customer/list")
    .then((response) => {
      response.json().then((data) => {
        for (let row of data) {
          const el = "<tr>" 
            + "<td>" + row.name + "</td>"
            + "<td>" + row.email + "</td>"
            + "<td>" + "</td>"
            + "</tr>";
          $("#customer_rows").prepend(el)
        }

      })
    })
}

export function add_customer() {
  const customer = {
    name: document.getElementById("add_cust_name").value,
    email: document.getElementById("add_cust_email").value,
    password: document.getElementById("password1").value
  }

  const password2 = document.getElementById("password2").value

  if (customer.password.length == 0 || password2.length == 0 || customer.password != password2) {
    alert("Make sure that passwords match or are not empty")
  }

  Post("/customer/create", customer)
  .then((response) => {
    const status = response.status;
    if (status == 200) {
      updateStatus("Customer created successfully!")
    } else if (status == 422) {
      response.json()
        .then((data) => {
          updateStatus(data);
        });
    } else {
      console.error("Unknown error occured!")
    }
  }) // parses response to JSON
  .catch(error => console.error(`Fetch Error =\n`, error));


  $('#addCustomer').modal('hide')
}

export function delete_customer() {
  const email = document.getElementById("delete_cust_email").value

  if (!email || email.length == 0) {
    updateStatus("Needs an email to be entered")
  }
  Delete("/customer/delete", {email: email})
  .catch(error => console.error(`Fetch Error =\n`, error));


  $('#deleteCustomer').modal('hide')
}