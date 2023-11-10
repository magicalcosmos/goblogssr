<template lang="html">
  <section class="home test">
    <div class="test">
      <div class="content">
        <ul class="content-list">
          <li v-for="item in articles" class="content-item">
            <h2>
              <a :href="`#/blog/${item.id}`">{{ item.title }}</a>
            </h2> 
            <div class="author">
              <span>所属目录：Frontend</span>
              <span>创建者：Brody</span>
              <span>创建时间：{{formatDate(new Date(item.publishAt))}}</span>
            </div>
            <div class="brief" v-html="item.brief"></div>
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
    </div>
    <Paginator
      v-if="page.total"
      :rows="page.pageSize"
      :totalRecords="page.total"
      :rowsPerPageOptions="[10,20,30]"
      @page="onPage($event)"
    ></Paginator>
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

  import { Article, Category } from '@/api';
  import { formatDate } from '@/utils/time'

  export default {
    components: {
      Viewer,
      Paginator
    },
    data() {
      return {
        formatDate: formatDate,
        page: {
          currentPage: 1,
          pageSize: 10,
        },
        articles: [],
        categories: [
          {
            name: 'All Categories',
            className: 'all'
          }
        ]
      }
    },
    methods: {
      getArticleList() {
        Article.list(this.page).then(res => {
          this.articles = res.articleList.articles;
          this.page = Object.assign(this.page, res.articleList.page);
        });
      },
      getCategoryList() {
        Category.list({
          currentPage: 1,
          pageSize: 10000,
        }).then(res => {
          res.categoryList.categories.forEach((item) => {
            this.categories.push({
              name: item.name
            })
          });
          this.page = Object.assign(this.page, res.categoryList.page);
        });
      },
      onPage(event) {
        this.page.currentPage = ++event.page;
        this.page.pageSize = event.rows;
        this.getArticleList();
      }
    },
    mounted() {
      this.getArticleList();
      this.getCategoryList();
    }
  }
</script>
<style scoped>
  .home, .test {
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

  .content .brief {
    color: rgba(156, 163, 175, 1);
    padding: 10px 0;
    word-break: break-all;
    white-space: pre-wrap;
  }

  .content h2 a {
    background-color: transparent;
    font-size: 30px;
    color: #fff;
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
    color: rgba(156, 163, 175, 1);
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

  .p-paginator {
    background: none !important;
  }
</style>
