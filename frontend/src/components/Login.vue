<template>
    <div class="List col-md-3">
        <div class="form-group">
            <label for="name">User name</label>
            <input v-model="data.login" type="text" class="form-control" id="name" placeholder="Login">
        </div>
        <div class="form-group">
            <label for="email">Email address</label>
            <input v-model="data.email" type="email" class="form-control" id="email" placeholder="Enter email">
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input v-model="data.password" type="password" class="form-control" id="password" placeholder="Password">
        </div>
        <button type="submit" class="btn btn-primary" @click="register">Register</button>
    </div>
</template>

<script>
// import axios from 'axios'
import {HTTP} from '../assets/js/http-common'

export default {
  name: 'Login',
  data () {
    return {
      data: {
        login: '',
        email: '',
        password: ''
      }
    }
  },
  methods: {
    register: function (event) {
      console.log(JSON.stringify(this.data))
      //   var querystring = require('querystring')
      // axios.post('http://something.com/', querystring.stringify({ foo: 'bar' }));
      HTTP({
        method: 'post',
        url: '/api/register',
        data: JSON.stringify(this.data)
        }).then(function (response) {
            console.log(JSON.stringify(response))
        }).catch(e => console.log(e))
    },
    login: function () {},
    logout: function () {}
  }
}
</script>
