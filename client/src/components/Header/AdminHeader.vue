<template lang="html">
  <header class="header">
    <nav class="primary-navbar">
      <div class="container clearfix">
        <ul class="left-list">
          <li class="left-item">
            <a href="#/admin/posts">
              <i class="pi pi-home"></i>
            </a>
          </li>
          <li class="left-item active">
            <a href="#/admin/posts">Posts</a>
          </li>
          <li class="left-item">
            <a href="#/admin/galleries">Galleries</a>
          </li>
          <li class="left-item">
            <a href="#/admin/enquires">Enquires</a>
          </li>
          <li class="left-item">
            <a href="#/admin/users">users</a>
          </li>
        </ul>
        <ul class="right-list">
          <li class="left-item">
            <a href="#/admin/posts">
              <i class="pi pi-globe"></i>
            </a>
          </li>
          <li class="left-item">
            <a href="#/admin/posts">
              <i class="pi pi-sign-out"></i>
            </a>
          </li>
        </ul>
      </div>
    </nav>
    <nav class="secondary-navbar">
      <div class="container">
        <ul class="dir-list clearfix">
          <li class="dir-item active">
            <a href="#/admin/posts-list">Posts</a>
          </li>
          <li class="dir-item">
            <a href="#/admin/category-list">Post Categories</a>
          </li>
        </ul>
      </div>
    </nav>
  </header>
</template>
<script>
  import 'primeicons/primeicons.css';
  export default {
    components: {
      
    },
    data: () => {
      return {
        
      }
    },
    methods: {
      
    }
  };
</script>

<style scoped>
  .header {
    font-size: 12px;
    width: 100%;
  }

  .header .pi-home {
    
  }

  .primary-navbar {
    background-color: #1385e5;
    color: #fff;
    padding: 5px 0;
  }

  .primary-navbar .container, .secondary-navbar .container {
    margin: 0 auto;
    max-width: 1170px;
  }

  .secondary-navbar .container {
    padding-left: 10px;
  }

  .primary-navbar .left-list, .secondary-navbar .dir-item {
    float: left;
    vertical-align: top;
  }


  .primary-navbar .left-item {
    cursor: pointer;
    display: inline-block;
    position: relative;
    vertical-align: top;
  }

  .primary-navbar .left-item:not(:first-child):before {
    -ms-transform: translate(0, -50%);
    -o-transform: translate(0, -50%);

    background-color: rgba(255, 255, 255, .2);
    content: " ";
    height: 1em;
    left: 0;
    margin-top: -0.3rem;
    position: absolute;
    top: 50%;
    width: 1px;
    z-index: 1;
  } 

  .primary-navbar .left-item a{
    color: #fff;
    cursor: pointer;
    display: block;
    text-decoration: none;
    padding: 5px 11px;
  }

  .primary-navbar .left-item.active {
    box-shadow: 2px 0 0 #1385e5;
    z-index: 2;
  }

  .primary-navbar .left-item.active a{
    background-color: #1073c6;
    color: white;
    transition: color 180ms;
  }

  .primary-navbar .right-list {
    float: right;
    position: relative;
  }

  .secondary-navbar {
    background-color: #f5f5f5;
    border-bottom: 1px solid #dbdbdb;
    color: #666;
  }

  .secondary-navbar  .dir-item {
    margin-right: 10px;
  }

  .secondary-navbar .dir-item a {
    color: #666;
    display: inline-block;
    text-decoration: none;
    padding: 5px 10px;
    position: relative;
  }

  .secondary-navbar .dir-item a:hover .secondary-navbar .dir-item.active a {
    color: #333;
  }

  .secondary-navbar .dir-item.active a::before,
  .secondary-navbar .dir-item.active a::after {
    border: 8px solid transparent;
    border-bottom-width: 0;
    content: "";
    height: 0;
    left: 50%;
    margin-left: -8px;
    position: absolute;
    width: 0;
  }

  .secondary-navbar .dir-item.active a::after {
    border-top-color: #f5f5f5;
    bottom: -8px;
  }

  .secondary-navbar .dir-item.active a::before {
    border-top-color: #ccc;
    bottom: -9px;
  }

</style>
