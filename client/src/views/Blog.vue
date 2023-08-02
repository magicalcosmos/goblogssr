<template lang="html">
  <section class="home">
    <div class="content">
      <ul class="content-list">
        <li class="content-item">
          <h2>
            <a href="#/blog/0">[Vim] [MacOS] Enable vim shortcuts keys in terminal</a>
          </h2> 
          <div class="author">
            <span>所属目录：Frontend</span>
            <span>创建者：Brody</span>
            <span>创建时间：2012-01-01</span>
          </div>
          <div class="bref">
            <viewer
              initialValue="### test"
            />
          </div>
        </li>
        <li class="content-item">
          <h2>
            <a href="#/blog/1">[Vim] [MacOS] Enable vim shortcuts keys in terminal</a>
          </h2> 
          <div class="author">
            <span>所属目录：Frontend</span>
            <span>创建者：Brody</span>
            <span>创建时间：2012-01-01</span>
          </div>
          <div class="bref">
            <viewer
              initialValue="### test"
            />
          </div>
        </li>
      </ul>
    </div>
    <div class="blog-category clearfix">
      <ul class="blog-category-list">
        <li
          v-for="(item, index) in categories"
          :key="index"
          :class="['blog-category-item', item.className || '']"
        >{{ item.name }}</li>
      </ul>
    </div>
    <Paginator :rows="10" :totalRecords="totalRecords" :rowsPerPageOptions="[10,20,30]"></Paginator>
  </section>
</template>
<script>
  import '@toast-ui/editor/dist/toastui-editor-viewer.css';
  import '@toast-ui/chart/dist/toastui-chart.css';
  import 'prismjs/themes/prism.css';
  import '@toast-ui/editor-plugin-table-merged-cell/dist/toastui-editor-plugin-table-merged-cell.css';
  import 'primevue/resources/primevue.min.css';
  import 'primeicons/primeicons.css'
  import 'primevue/resources/themes/lara-light-blue/theme.css';

  import Paginator from 'primevue/paginator';
  // import { Parser, createRenderHTML } from '@toast-ui/toastmark';
  import { Viewer } from '@toast-ui/vue-editor';
  export default {
    components: {
      Viewer,
      Paginator
    },
    data() {
      return {
        page: 1,
        totalRecords: 120,
        categories: [
          {
            name: 'All Categories',
            className: 'all'
          },
          {
            name: 'Docker'
          },
          {
            name: 'Frontend'
          },
          {
            name: 'Golang'
          },
          {
            name: 'Ubuntu'
          },
          {
            name: 'other'
          }
        ]
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
