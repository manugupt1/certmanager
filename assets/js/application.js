require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle");
require("./customer.js");

import { add_customer } from './customer.js';

$(() => {

  $('#add_customer').click(function () {
    add_customer();
  })

});
