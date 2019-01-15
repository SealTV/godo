<template>
    <div class="Todos col-md-6">
      <div class="list-group" v-if="todos().length">
          <TodoItem
            v-for="t in todos()"
            :key="t.id"
            :todo="t"
            :selected="selectedTodo"
            @remove="removeTodo"
            @select="selectTodo"
          />
      </div>
      <p v-else>
          Nothing left in the list. Add a new todo in the input above.
      </p>
      <CreateTodoItem placeholder="New todo"
        @add="addNewTodo"
      />
    </div>
</template>

<script>
import TodoItem from './TodoItem'
import CreateTodoItem from './CreateTodoItem'

export default {
  name: 'Todos',
  components: { TodoItem, CreateTodoItem },
  props: {
    todos: {
      type: Function,
      required: true
    },
    listId: {
      type: Number,
      required: true
    }
  },
  data () {
    return {
      selectedTodo: 1
    }
  },
  methods: {
    addNewTodo (value) {
      var todo = {
        id: 2,
        title: value,
        description: '',
        is_active: true,
        list_id: this.listId,
        user_id: 1
      }
      this.todos().push(todo)
    },
    removeTodo (idToRemove) {
    },
    selectTodo (idToSelect) {
      this.selectedTodo = idToSelect
    }
  }
}
</script>
