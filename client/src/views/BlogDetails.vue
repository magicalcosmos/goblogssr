<template>
  <article>
    <p>
      <a href="#/blog">← 返回</a>
    </p>
    <hr>
    <header>
      <h1>{{ data.title }}</h1>
      <div class="author">
        <span>所属目录：Frontend</span>
        <span>创建者：Brody</span>
        <span>创建时间：{{ data.publishAt }}</span>
      </div>
    </header>
    <div class="post" v-html="data.content"></div>
  </article>
</template>
<script>

import '@toast-ui/editor/dist/toastui-editor-viewer.css';
import '@toast-ui/chart/dist/toastui-chart.css';
import 'prismjs/themes/prism.css';
import '@toast-ui/editor-plugin-table-merged-cell/dist/toastui-editor-plugin-table-merged-cell.css';
import { Viewer } from '@toast-ui/vue-editor';
import { formatDate } from '@/utils/time';
import { Article } from '@/api';

export default {
  components: {
    Viewer
  },
  data() {
    return {
      data: {},
      initText: ''
    };
  },
  methods: {
    /**
     * 根据ID获取文章
     */
    getArticleById() {
      Article.getArticleById({
        id: this.$route.params.id
      }).then(res => {
        const data = res.getArticleById;
        data.publishAt = formatDate(new Date(data.publishAt));
        this.data = data;
      });
    },
  },
  mounted() {
    this.getArticleById();
  },

}
</script>
<style scoped>
  h1 {
    font-size: 36px;
  }
  p {
    margin-top: 20px;
  }

  p a {
    text-decoration: none;
  }

  p a:hover {
    text-decoration: underline;
  }

  hr {
    border: 0;
    border-top: 1px solid #eee;
    margin-top: 20px;
    margin-bottom: 20px;
  }

  .author {
    font-size: 12px;
    margin-top: 5px;
  }

  .author span {
    color: rgba(0, 0, 0, .5);
    margin-right: 10px;
  }

  .post {
    white-space: pre-wrap;
  }
</style>
