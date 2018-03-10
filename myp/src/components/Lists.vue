<template>
  <div class="List col-md-3">
    <div class="list-group" id="list" v-if="todoLists.length">
      <ListItem
        v-for="l in todoLists"
        :key="l.list.id"
        :list="l.list"
        :selected="selectedList"
        :todosCount="l.todos.length"
        @remove="removeList"
        @select="selectList"
      />
      <!-- <ul class="list-group flex-column" id="list" v-if="todoLists.length">
      </ul> -->
    </div>
    <p v-else>
      Nothing left in the list. Add a new list in the input above.
    </p>
    <BaseInputText placeholder="New list"
      @add="addNewList"
    />
  </div>
</template>

<script>
import BaseInputText from './BaseInputText'
import ListItem from './ListItem'

export default {
  name: 'Lists',
  components: { BaseInputText, ListItem },
  props: {
    todoLists: {
      type: Array,
      required: true
    },
    selectedList: {
      type: Number,
      required: true
    }
  },
  data () {
    return {
    }
  },
  methods: {
    addNewList (value) {
      this.$emit('addList', value)
    },
    removeList (idToRemove) {
      this.lists = this.lists.filter(list => {
        return list.id !== idToRemove
      })
    },
    selectList (idToSelect) {
      this.$emit('select', idToSelect)
    }
  }
}
</script>
