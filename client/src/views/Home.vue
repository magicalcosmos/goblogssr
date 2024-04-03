<template lang="html">
  <section class="home-main">
    <div id="homeId" class="home">
      <ul id="container" class="home-list container">
        <li v-for="d in data" class="home-item">
          <div v-html="d.cover" class="cover"></div>
          <a>
            <h5>{{ d.title }}</h5>
          </a>
          <p class="date-time">{{ d.datetime }}</p>
          <p class="brief" v-html="d.brief"></p>
          <div class="tag">
            <a v-for="tag in d.tags" :href="tag.href">{{tag.name}} </a>
          </div>
          <a class="read-more">Read more</a>
        </li>
      </ul>
    </div>
    <div v-if="page.pageSize > 30">
      <Paginator
        :rows="page.pageSize"
        :totalRecords="page.total"
        :rowsPerPageOptions="[30, 60, 90]"
        @page="onPage($event)"
      ></Paginator>
    </div>
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

  import { Article } from '@/api';
  import { formatDate } from '@/utils/time';

  export default {
    components: {
      Paginator, 
    },
    data() {
      return {
        gap: 10,
        top: 10,
        column: 3,
        data: [],
        formatDate: formatDate,
        page: {
          currentPage: 1,
          pageSize: 30,
          total: 0,
        },
        articles: [],
      }
    },
    methods: {
      _GetArticleList() {
        Article.list(this.page).then(res => {
          this.data = res.articleList.articles;
          this.page = Object.assign(this.page, res.articleList.page);
        });
      },
      onPage(event) {
        Object.assign(this.page, {
          currentPage: ++event.page,
          pageSize: rows,
        })
        this._GetArticleList();
      }
    },

    
    mounted() {
      this._GetArticleList();
      window.onload = this._InitSortItems;
    }
  }
</script>
<style scoped>
  .home {
  }

  .home-list {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    height: 100%;
  }

  .home-list:after {
    content: '';
    width: 32%;
  }

  .home-item {
    background: rgba(31, 41, 55, 1);
    border-radius: 8px;
    cursor: pointer;
    max-width: 32%;
    margin-bottom: 1.25rem;
    overflow: hidden;
    padding: 1.25rem;
  }

  .home-item:hover {
    background: rgba(92, 99, 111, 1);
  }
  .home-item .cover {
    text-align: center;
    width:100%;
  }
  .home-item >>> img {
    width: 100%;
  }

  .home-item h5 {
    color: #fff;
    font-weight: 700;
    font-size: 1.5rem;
  }

  .home-item.hover {
    background: rgb(55 65 81 / var(--tw-border-opacity));
  }

  .home-item .date-time,
  .home-item .brief {
    color: rgba(152, 163, 175, 1);
    font-size: 12px;
  }

  .home-item .brief {
    font-weight: 400;
    font-size: 14px;
    line-height: 24px;
    margin: 5px 0 10px 0;
  }

  .home-item .tag a {
    color: rgba(255, 255, 0, 1);
    font-size: 14px;
    line-height: 20px;
    margin: 4px;
  }

  .home-item .read-more {
    background-color: rgba(251, 146, 60, 1);
    border-radius: 8px;
    color: #000;
    display: inline-block;
    font-size: 14px;
    font-weight: 500;
    margin-top: 10px;
    padding: 8px 16px;
    text-align: center;
  }
</style>
