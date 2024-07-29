<template lang="html">
  <section>
    <div class="type-writer">
      <div class="role"></div>
      <div class="type-writer-main">
        <a href="" class="typewrite" data-period="2000">
          <span class="wrap"></span>
        </a>
      </div>
    </div>
    <div class="home-main">
      <div id="homeId" class="home">
        <ul id="container" class="home-list container">
          <li v-for="d in data" class="home-item" @click="goBlog(d.id)">
            <div v-html="d.cover" class="cover"></div>
            <a>
              <h5 class="item-title">{{ d.title }}</h5>
            </a>
            <p class="date-time">{{ d.datetime }}</p>
            <p class="brief" v-html="d.brief"></p>
            <div class="tag">
              <a v-for="tag in d.tags" :href="tag.href">{{ tag.name }} </a>
            </div>
            <a class="read-more" :href="`/blog/${d.id}`">{{ $t('button.read_more') }} </a>
          </li>
        </ul>
      </div>
      <div v-if="page.pageSize > 30">
        <Paginator :rows="page.pageSize" :totalRecords="page.total" :rowsPerPageOptions="[30, 60, 90]"
          @page="onPage($event)"></Paginator>
      </div>
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
      typewriterData: ["Hi, I'm Brody.", "I am Creative.", "I Love Design.", "I Love to Develop.", "I Love to Share Something."],
    }
  },
  methods: {
    _GetArticleList() {
      Article.list(this.page).then(res => {
        this.data = res.articleList.articles;
        this.data.map((item) => {
          item.datetime = this.formatDate(new Date(item.publishAt));
        });
        this.page = Object.assign(this.page, res.articleList.page);
      });
    },
    onPage(event) {
      Object.assign(this.page, {
        currentPage: ++event.page,
        pageSize: rows,
      });
      this._GetArticleList();
    },
    handleTypeWriter() {
      const _this = this;
      var TxtType = function (el, toRotate, period) {
        this.toRotate = toRotate;
        this.el = el;
        this.loopNum = 0;
        this.period = parseInt(period, 10) || 2000;
        this.txt = '';
        this.tick();
        this.isDeleting = false;
      };

      TxtType.prototype.tick = function () {
        var i = this.loopNum % this.toRotate.length;
        var fullTxt = this.toRotate[i];

        if (this.isDeleting) {
          this.txt = fullTxt.substring(0, this.txt.length - 1);
        } else {
          this.txt = fullTxt.substring(0, this.txt.length + 1);
        }

        this.el.innerHTML = '<span class="wrap">' + this.txt + '</span>';

        var that = this;
        var delta = 200 - Math.random() * 100;

        if (this.isDeleting) { delta /= 2; }

        if (!this.isDeleting && this.txt === fullTxt) {
          delta = this.period;
          this.isDeleting = true;
        } else if (this.isDeleting && this.txt === '') {
          this.isDeleting = false;
          this.loopNum++;
          delta = 500;
        }

        setTimeout(function () {
          that.tick();
        }, delta);
      };

      window.onload = function () {
        const elements = document.getElementsByClassName('typewrite');
        new TxtType(elements[0], _this.typewriterData, 2000)
        // INJECT CSS
        var css = document.createElement("style");
        css.type = "text/css";
        css.innerHTML = `
          .typewrite > .wrap { 
            border-right: 3px solid #333;  
            animation: glow 800ms ease-out infinite alternate; 
          }
          @keyframes glow {
              0% {
                  border-color: #fff;
              }
              100% {
                  border-color: #333;
              }
          }
        `;
        document.body.appendChild(css);
      };
    },
    goBlog(id) {
      this.$router.push({
        path: '/blog/' + id
      })
    },
  },


  mounted() {
    this._GetArticleList();
    this.handleTypeWriter();
  }
}
</script>
<style scoped>
.type-writer {
  background-repeat: no-repeat;
  background-size: cover;
  background-image: url(@static/img/type-writer-bg.png);
  font-size: 30px;
  height: 18rem;
  overflow: hidden;
  text-align: center;
}

.type-writer a {
  color: #333;
  text-decoration: none;
}

.role {
  background-image: url(@static/img/pangu.jpeg);
  background-repeat: no-repeat;
  background-size: contain;
  background-position: center;
  border-radius: 5px;
  height: 168px;
  margin: 30px auto;
  width: 300px;
}

.home-main {
  max-width: 1170px;
  margin: 15px auto 0 auto;
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
  width: 100%;
}

.home-item>>>img {
  width: 100%;
}

.home-item h5 {
  color: #fff;
  font-weight: 700;
  font-size: 1.5rem;
  margin-top: 5px;
}

.home-item.hover {
  background: #337ab7;
}

.home-item .date-time,
.home-item .brief {
  color: rgba(152, 163, 175, 1);
  font-size: 12px;
  margin-top: 5px;
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
