<template>
    <div>
      <b-navbar-nav class="ml-auto">
        <b-navbar right v-if="!$auth.check()">
          <b-nav-form>
            <b-form-input v-model="data.username" class="mr-sm-2" type="text" id="name" placeholder="Login"></b-form-input>
            <b-form-input v-model="data.password" class="mr-sm-2" type="password" id="name" placeholder="Password"></b-form-input>
            <b-button variant="outline-success" class="my-2 my-sm-0" type="submit" @click="login">Login</b-button>
          </b-nav-form>
        </b-navbar>

        <b-nav-item-dropdown right v-if="$auth.check()">
            <template slot="button-content">
              <em>User</em>
            </template>
            <b-dropdown-item href="#">Profile</b-dropdown-item>
            <b-dropdown-item href="#"  @click="logout">Signout</b-dropdown-item>
        </b-nav-item-dropdown>
      </b-navbar-nav>
      <div>
        <b-modal ref="myModalRef" hide-footer title="Login failed">
          <div class="d-block text-center">
            <p>Invalid username or password!</p>
          </div>
          <b-btn class="mt-3" variant="outline-danger" block @click="hideModal">Ok</b-btn>
        </b-modal>
      </div>
    </div>
</template>

<script>

export default {
  name: 'Login',
  data () {
    return {
      data: {
        username: '',
        email: '',
        password: ''
      },
      error: {}
    }
  },
  methods: {
    showModal () {
      this.$refs.myModalRef.show()
    },
    hideModal () {
      this.$refs.myModalRef.hide()
    },
    register: function (event) {
      this.$auth.register({
        data: this.data,
        success: function () {},
        error: function () {},
        autoLogin: true,
        rememberMe: true
      })
    },
    login: function () {
      this.$auth.login({
        data: this.data,
        success: function (response) {
          this.data.username = response.data.username
          this.$auth.user(response)
        },
        error: function () {
          this.showModal()
        },
        rememberMe: true,
        autoLogin: true,
        fetchUser: true
      })
    },
    logout: function () {
      this.$auth.logout({
        makeRequest: true,
        params: {},
        success: function () {},
        error: function () {}
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
