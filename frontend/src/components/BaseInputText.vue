<template>
    <div class="container-fluid">
      <div class="row">
        <div class="input-group mb-3">
          <input
            type="text"
            class="form-control"
            v-model="listName"
            placeholder="New list name"
            aria-label="New list name"
            aria-describedby="basic-addon2">
          <div class="input-group-append">
            <button class="btn btn-success" type="button" @click="click">Add list</button>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
export default {
  name: 'BaseInputText',
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      listName: ''
    }
  },
  methods: {
    click: function (event) {
      if (!this.listName.length) {
        return
      }
      this.$emit('add', this.listName)
      this.listName = ''
    }
  }
}
</script>
