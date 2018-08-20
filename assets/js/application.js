require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle");

import { get_customer, add_customer, delete_customer } from './customer.js';

$(() => {

  get_customer();

  $('#add_customer').click(function () {
    add_customer();
  })

  $('#delete_customer').click(function () {
    delete_customer();
  })
});
