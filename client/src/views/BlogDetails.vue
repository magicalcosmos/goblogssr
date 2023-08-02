<template>
  <article>
    <p>
      <a href="#/blog">← 返回</a>
    </p>
    <hr>
    <header>
      <h1>[System] Mac系统的环境变量加载顺序</h1>
      <div class="author">
        <span>所属目录：Frontend</span>
        <span>创建者：Brody</span>
        <span>创建时间：2012-01-01</span>
      </div>
    </header>
    <div class="post">
      <viewer
        :initialValue="data"
      />
    </div>
  </article>
</template>
<script>

import '@toast-ui/editor/dist/toastui-editor-viewer.css';
import '@toast-ui/chart/dist/toastui-chart.css';
import 'prismjs/themes/prism.css';
import '@toast-ui/editor-plugin-table-merged-cell/dist/toastui-editor-plugin-table-merged-cell.css';
import { Viewer } from '@toast-ui/vue-editor';

export default {
  components: {
    Viewer
  },
  data() {
    return {
      data: '# test'
    };
  }
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
    color: rgba(0, 0, 0, 0.5);
    margin-right: 10px;
  }

  .post {
    white-space: pre-wrap;
  }
</style>
