<template>
  <div>
    <a v-if="selected == list.id" href="#" class="list-group-item d-flex justify-content-between align-items-center list-group-item-action active" @click="$emit('select', list.id)">
      {{list.name}}
      <span class="badge badge-light badge-pill">{{todosCount}}</span>
    </a>
    <a v-else href="#" class="list-group-item d-flex justify-content-between align-items-center list-group-item-action" @click="$emit('select', list.id)">
      {{list.name}}
      <span class="badge badge-primary badge-pill">{{todosCount}}</span>
    </a>
  </div>
</template>

<script>

export default {
  name: 'ListsItem',
  props: {
    list: {
      type: Object,
      required: true
    },
    selected: {
      type: Number,
      required: true
    },
    todosCount: {
      type: Number,
      required: true
    }
  },
  methods: {
  }
}
</script>
