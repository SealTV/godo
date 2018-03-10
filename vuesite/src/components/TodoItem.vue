<template>
  <a
    v-if="selected == todo.id"
    href="#"
    class="list-group-item list-group-item-action active"
    @click="$emit('select', todo.id)"
  >
    {{todo.title}}
  </a>
  <a
    v-else
    href="#"
    class="list-group-item list-group-item-action"
    @click="$emit('select', todo.id)"
  >
    {{todo.title}}
  </a>
</template>

<script>
export default {
  name: 'TodoItem',
  props: {
    todo: {
      type: Object,
      required: true
    },
    selected: {
      type: Number,
      required: true
    }
  }
}
</script>
