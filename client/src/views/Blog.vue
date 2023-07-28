<template lang="html">
  <section class="home">
    <div class="content">
      <ul class="content-list">
        <li class="content-item">
          <h2>
            <a href="#/0/details">[Vim] [MacOS] Enable vim shortcuts keys in terminal</a>
          </h2> 
          <div class="author">
            <span>所属目录：Frontend</span>
            <span>创建者：Brody</span>
            <span>创建时间：2012-01-01</span>
          </div>
          <div class="bref">
 1.export与export default均可用于导出常量、函数、文件、模块等
 2.在一个文件或模块中，export、import可以有多个，export default仅有一个
 3.通过export方式导出，在导入时要加{ }，export default则不需要
 4.export能直接导出变量表达式，export default不行。
          </div>
        </li>
        <li class="content-item">
          <h2>
            <a href="#/1/details">[Vim] [MacOS] Enable vim shortcuts keys in terminal</a>
          </h2> 
          <div class="author">
            <span>所属目录：Frontend</span>
            <span>创建者：Brody</span>
            <span>创建时间：2012-01-01</span>
          </div>
          <div class="bref">
1.export与export default均可用于导出常量、函数、文件、模块等
2.在一个文件或模块中，export、import可以有多个，export default仅有一个
3.通过export方式导出，在导入时要加{ }，export default则不需要
4.export能直接导出变量表达式，export default不行。
          </div>
        </li>
      </ul>
    </div>
    <div class="blog-category clearfix">
      <ul class="blog-category-list">
        <li class="blog-category-item all">All Categories</li>
        <li class="blog-category-item">Docker</li>
        <li class="blog-category-item">Frontend</li>
        <li class="blog-category-item">Golang</li>
        <li class="blog-category-item">Ubuntu</li>
        <li class="blog-category-item">other</li>
      </ul>
    </div>
    <Paginator :rows="10" :totalRecords="totalRecords" :rowsPerPageOptions="[10,20,30]"></Paginator>
  </section>
</template>
<script>

  import 'primevue/resources/primevue.min.css';
  import 'primeicons/primeicons.css'
  import 'primevue/resources/themes/lara-light-blue/theme.css';
  import Paginator from 'primevue/paginator';
  export default {
    components: {
      Paginator
    },
    data() {
      return {
        page: 1,
        totalRecords: 120,
      }
    },
    methods: {
      
    },
    mounted() {
      
    }
  }
</script>
<style scoped>
  .home {
    font-size: 14px;
    margin-top: 20px;
    white-space: nowrap;
  }
  
  .content {
    display: inline-block;
    vertical-align: top;
    word-wrap: break-word;
    width: 75%;
  }

  .content .bref {
    white-space: pre-wrap;
  }

  .content h2 a {
    background-color: transparent;
    font-size: 30px;
    color: #337ab7;
    text-decoration: none;
  }

  .content h2 a:hover {
    text-decoration: underline;
  }

  .content .content-list {
    width: 90%;
  }

  .content .content-item {
    border-top: 1px solid #e2e8f0;
    padding: 0.3rem 0;
  }

  .author {
    font-size: 12px;
    margin-top: 5px;
  }

  .author span {
    color: rgba(0, 0, 0, 0.5);
    margin-right: 10px;
  }

  .blog-category {
    display: inline-block;
    vertical-align: top;
    width: 25%;
  }

  .blog-category .blog-category-item {
    background-color: #fff;
    border: 1px solid #ddd;
    border-top: none;
    color: #555;
    text-align: left;
    padding: 10px 15px;
    
  }

  .blog-category .blog-category-item.all {
    background-color: #337ab7;
    border-color: #337ab7;
    border-top-right-radius: 4px;
    border-top-left-radius: 4px;
    color: #fff;
  }

  .blog-category .blog-category-item:last-child {
    border-bottom-right-radius: 4px;
    border-bottom-left-radius: 4px;
  }
</style>
