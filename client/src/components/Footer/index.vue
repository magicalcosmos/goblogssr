<template lang="html">
  <footer class="footer">
    <p>
      <a href="https://beian.miit.gov.cn/" target="_blank">蜀ICP备14026941号-1</a>
      &copy2016-{{ date }}
    </p>
  </footer>
</template>
<script>
  export default {
    data() {
      return {
        date: `${new Date().getUTCFullYear()}`,
      };
    }
  }
</script>

<style scoped>
.footer {
  border-top: 1px solid #fb923c;
  margin-top: 30px;
  padding: 30px 0;
  color: #fff;
  text-align: center;
}
.footer p {
  margin: 0 0 10px;
}

.footer p a {
  color: #337ab7;
  text-decoration: none;
}

.footer p a:hover {
  text-decoration: underline;
}
</style>
