<template>
  <div id="app">
    <h1>To Do List</h1>
    <table>
      <thead>
        <tr>
          <th>Item</th>
          <th>Due Date</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="todo in todos" :key="todo.id">
          <td>{{ todo.value }}</td>
          <td>{{ todo.due_date }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
const appData = {
  todos: []
}

export default {
  name: 'app',
  data() {
    return appData;
  },
  mounted: function() {
    this.getTodos();
  },
  methods: {
    getTodos: getTodos
  }
}

async function getTodos() {
  try {
    const response = await fetch('api/todos')
    const data = await response.json()
    appData.todos = data.list
  } 
  catch (error) {
    console.error(error)
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  margin-top: 60px;
}
</style>
