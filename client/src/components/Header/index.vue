<template lang="html">
  <header class="header clearfix">
    <Toast></Toast>
    <div class="logo">
      <span class="web-name">天马星空</span>
      <a href="/" class="navbar-brand">ShareUs</a>
    </div>
    <ul class="nav-list">
      <li v-for="item in navList" :class="`nav-item${ item.isActive ? ' active': ''}`" :name="item.name">
        <a :href="item.href">{{item.label}}</a>
      </li>
    </ul>
    <a href="#/admin/login" class="sigin">{{ $t('button.login')}}</a>
  </header>
</template>
<script>
  import 'primeicons/primeicons.css';
  import Toast from 'primevue/toast';
  export default {
    components: {
      Toast,
    },
    watch: {
     '$route'() {
       this.navSwitch();
     } 
    },
    data() {
      return {
        navList: [{
          name: 'blog',
          label: this.$t('blog.name'),
          isActive: true,
          href: '#/blog'
        }, {
          name: 'gallery',
          label: this.$t('gallery.name'),
          isActive: false,
          href: '#/gallery'
        }, {
          name: 'contact',
          label: this.$t('contact.name'),
          isActive: false,
          href: '#/contact'
        }],
      }
    },
    methods: {
      navSwitch() {
        const path = this.$route.path;
        const pathArr = path.split('/');
        this.navList.forEach((item) => {
          item.isActive = `#${path}` === item.href || (pathArr.length >= 3 ? `#/${pathArr[1]}` === item.href : false);
        });
      }
    },
    mounted() {
      this.navSwitch();
    }
    
  };
</script>

<style scoped>
  .header {
    background: #f8f8f8;
    border: 1px solid #e7e7e7;
    border-radius: 5px;
    box-sizing: border-box;
    margin-top: 30px;
    padding: 0 15px;
    width: 100%;
  }

  .header > * {
    float: left;
  }
  
  .sigin {
    color: #777;
    float: right;
    margin-top: 10px;
    text-decoration: none;
  }

  .sigin:hover {
    color: #555;
  }

  .navbar-brand {
    color: #777;
    font-size: 18px;
    padding: 10px 15px;
    position: relative;
    text-decoration: none;
    top: 7px;
  }

  .web-name {
    color: red;
    float:left;
    font-size: 10px;
    font-weight: bold;
    text-shadow: 0 0 4px white, 0 -6px 4px #ff3, 1px 2px 6px #fd6, -2px -2px 4px #f80, 2px -8px 7px #f20;
  }

  .nav-list {
    overflow: hidden;
  }

  .nav-list .nav-item {
    cursor: pointer;
    float: left;
    padding: 10px 15px;
  }

  .nav-list .nav-item.active {
    background-color: #e7e7e7;
    color: #555;
  }

  .nav-list .nav-item a {
    color: #777;
    text-decoration: none;
  }

</style>
