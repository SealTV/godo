<template>
    <div class="container-fluid">
      <div class="row">
        <div class="input-group mb-4">
          <input
            type="text"
            class="form-control"
            v-model="todoName"
            placeholder="New todo"
            aria-label="New todo"
            aria-describedby="basic-addon2">
          <div class="input-group-append">
            <button class="btn btn-success" type="button" @click="click">Add todo</button>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
export default {
  name: 'CreateTodoItem',
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      todoName: ''
    }
  },
  methods: {
    click: function (event) {
      if (!this.todoName.length) {
        return
      }
      this.$emit('add', this.todoName)
      this.todoName = ''
    }
  }
}
</script>
