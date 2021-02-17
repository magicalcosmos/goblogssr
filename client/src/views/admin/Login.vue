<template lang="html">
  <section class="login">
    <div>
      <input ref="refUserName" type="text" class="username"/>
    </div>
    <div>
      <input ref="refPassword" type="password" class="password"/>
    </div>
    <div>
      <input type="submit" :value="$t('button.login')" @click="handleSubmit"/>
    </div>
  </section>
</template>

<script>
  import { Login } from '@/api'
  export default {
    methods: {
      /**
       * user login
       */
      handleSubmit() {
        // this.$router.push('/admin/posts');
        const userName = this.$refs.refUserName.value;
        const password = this.$refs.refPassword.value;
        Login.login({
          userName,
          password
        }).then(() => {
          this.$router.push('/admin/posts');
        });
      }
    }
  };
</script>
