

import {post} from './requests.js';
import {updateStatus} from './status.js';

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

  post("/customer/create", customer)
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