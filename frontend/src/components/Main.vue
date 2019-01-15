<template>
    <div class="main">
      <nav class="navbar navbar-expand-md navbar-dark bg-dark mb-4">
        <a class="navbar-brand" href="#">Top navbar</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarCollapse">
          <ul class="navbar-nav mr-auto">
            <li class="nav-item active">
              <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">Link</a>
            </li>
            <li class="nav-item">
              <a class="nav-link disabled" href="#">Disabled</a>
            </li>
          </ul>
          <Login></Login>
        </div>
      </nav>
      <div>
        <div class="container-fluid" >
          <div class="row">
            <Lists
              :todoLists="data.todoLists"
              :selectedList="selectedListId"
              @addList="addList"
              @select="selectList"
            />
            <Todos
              :todos="currentTodos"
              :listId="selectedListId"
            />
          </div>
        </div>
      </div>
    </div>
</template>

<script>
import Lists from './Lists'
import Todos from './Todos'
import Login from './Login'

export default {
  name: 'Main',
  data () {
    return {
      msg: 'Welcome to Your TODO site App',
      selectedListId: 1,
      data: {
        user: {
          id: 1,
          login: 'sealTV',
          password: 'password',
          email: 'test@mail.com'
        },
        todoLists: [
          {
            list: {
              id: 1,
              name: 'All',
              user_id: 1
            },
            todos: [
              {
                id: 1,
                title: 'Some title here',
                description: 'some description here',
                is_active: true,
                list_id: 1,
                user_id: 1
              },
              {
                id: 2,
                title: 'Some title here 1',
                description: 'some description here 1',
                is_active: true,
                list_id: 1,
                user_id: 1
              }
            ]
          },
          {
            list: {
              id: 2,
              name: 'Work',
              user_id: 1
            },
            todos: [
              {
                id: 3,
                title: 'Some title here 3',
                description: 'some description here',
                is_active: true,
                list_id: 2,
                user_id: 1
              },
              {
                id: 4,
                title: 'Some title here 4',
                description: 'some description here 1',
                is_active: true,
                list_id: 2,
                user_id: 1
              }
            ]
          }
        ]
      }
    }
  },
  components: { Lists, Todos, Login },
  methods: {
    addList (value) {
      var list = {
        list: {
          id: this.maxId() + 1,
          name: value,
          user_id: 1
        },
        todos: []
      }
      this.data.todoLists.push(list)
      this.selectedListId = list.list.id
    },
    maxId () {
      var max = 0
      for (var i = 0; i < this.data.todoLists.length; i++) {
        var list = this.data.todoLists[i].list
        if (max <= list.id) {
          max = list.id
        }
      }
      return max
    },
    selectList (idToSelect) {
      this.selectedListId = idToSelect
    },
    currentTodos: function () {
      for (var i = 0; i < this.data.todoLists.length; i++) {
        var list = this.data.todoLists[i].list
        if (list.id === this.selectedListId) {
          return this.data.todoLists[i].todos
        }
      }
      return []
    }
  }
}
</script>
