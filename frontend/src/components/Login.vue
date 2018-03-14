<template>
    <div class="List col-md-3">
        <div>
            <div v-if="this.$auth.ready()">
                <div class="form-group">
                    <label for="name">User name</label>
                    <input v-model="data.username" type="text" class="form-control" id="name" placeholder="Login">
                    </div>
                    <div class="form-group">
                        <label for="email">Email address</label>
                        <input v-model="data.email" type="email" class="form-control" id="email" placeholder="Enter email">
                    </div>
                    <div class="form-group">
                        <label for="password">Password</label>
                        <input v-model="data.password" type="password" class="form-control" id="password" placeholder="Password">
                    </div>
                    <button type="submit" class="btn btn-primary" v-if="!$auth.check()" @click="register">Register</button>
                    <div>
                        {{ this.$auth.user().email }}
                    </div>
                    <button type="submit" class="btn btn-primary" v-if="!$auth.check()" @click="login">Login</button>
                    <button type="submit" class="btn btn-primary" v-if="$auth.check()" @click="ping">Ping</button>
                    <button type="submit" class="btn btn-primary" v-if="$auth.check()" @click="logout">Logout</button>
                    </div>
                    <div>
                        {{ this.$auth.user().email }}
                    </div>
            </div>
            <div v-if="!this.$auth.ready()">
                Loading ...
            </div>
    </div>
</template>

<script>
// import axios from 'axios'
// import {HTTP} from '../assets/js/http-common'

export default {
  name: 'Login',
  data () {
    return {
      data: {
        username: '',
        email: '',
        password: ''
      }
    }
  },
  methods: {
    register: function (event) {
    //   console.log(JSON.stringify(this.data))
      //   var querystring = require('querystring')
      // axios.post('http://something.com/', querystring.stringify({ foo: 'bar' }));
    //   this.$cookie.set('test', 'hello world', 1)
      this.$auth.register({
        data: this.data,
        success: function () {},
        error: function () {},
        autoLogin: true,
        rememberMe: true
        // redirect: {name: 'account'},
        // etc...
      })
    },
    login: function () {
      this.$auth.login({
        data: this.data,
        success: function (response) {
          console.log('success ' + JSON.stringify(response))
          this.$auth.user(response)
        },
        error: function () {
          console.log('error ' + this.context)
        },
        rememberMe: true,
        autoLogin: true,
        fetchUser: true
      })
    },
    logout: function () {
    //   this.$cookie.delete('test')
      this.$auth.logout({
        makeRequest: true,
        params: {}, // data: {} in axios
        success: function () {},
        error: function () {}
        // redirect: '/login',
      })
    },
    ping: function () {
      this.$http.get('/ping?ping=hello').then(
        function (response) {
          console.log(response)
        })
    }
  }
}
</script>
