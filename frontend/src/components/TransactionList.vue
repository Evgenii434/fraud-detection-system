<template>
  <div>
    <h2>Transactions</h2>
    <ul>
      <li v-for="transaction in transactions" :key="transaction.id">
        {{ transaction.details }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      transactions: []
    }
  },
  created() {
    fetch('/transactions')
      .then(response => response.json())
      .then(data => {
        this.transactions = data;
      });
  }
}
</script>
