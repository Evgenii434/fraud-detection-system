<template>
  <div id="app">
    <h1>Fraud Detection System</h1>
    <transaction-list></transaction-list>
  </div>
</template>

<script>
import TransactionList from './components/TransactionList.vue';

export default {
  name: 'App',
  components: {
    TransactionList
  }
}
</script>
