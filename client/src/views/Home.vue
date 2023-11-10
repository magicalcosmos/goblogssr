<template lang="html">
  <section class="home">
    <ul id="container" class="home-list container">
      <li v-for="d in data" class="home-item">
        <a>
          <img :src="d.img" />
        </a>
        <a>
          <h5>{{ d.title }}</h5>
        </a>
        <p class="date-time">{{ d.datetime }}</p>
        <p class="brief">{{ d.brief }}</p>
        <div class="tag">
          <a v-for="tag in d.tags" :href="tag.href">{{tag.name}} </a>
        </div>
        <a class="read-more">Read more</a>
      </li>
    </ul>
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

  import { Article } from '@/api';
  import { formatDate } from '@/utils/time';

  let data = [{
    img: 'https://programmingpercy.tech/_app/immutable/assets/img0-b4114dea.webp',
    title: 'Amazon Web Services (AWS)&hairsp;—&hairsp;Serverless Architecture',
    datatime: 'February 19, 2021',
    brief: 'A complete tour for beginners to set up a full-stack application',
    tags: [{
      name: 'aws',
      href: '/'
    }, {
      name: 'cloud',
      href: '/blog'
    }, {
      name: 'frontend',
      href: '/blog'
    }]
  }, {
    img: 'https://programmingpercy.tech/_app/immutable/assets/img0-b4114dea.webp',
    title: 'Amazon Web Services (AWS)&hairsp;—&hairsp;Serverless Architecture',
    datatime: 'February 19, 2021',
    brief: 'A complete tour for beginners to set up a full-stack application',
    tags: [{
      name: 'aws',
      href: ''
    }, {
      name: 'cloud',
      href: ''
    }, {
      name: 'frontend',
      href: ''
    }]
  }, {
    img: 'https://programmingpercy.tech/_app/immutable/assets/img0-b4114dea.webp',
    title: 'Amazon Web Services (AWS)&hairsp;—&hairsp;Serverless Architecture',
    datatime: 'February 19, 2021',
    brief: 'A complete tour for beginners to set up a full-stack application',
    tags: [{
      name: 'aws',
      href: ''
    }, {
      name: 'cloud',
      href: ''
    }, {
      name: 'frontend',
      href: ''
    }]
  }, {
    img: 'https://programmingpercy.tech/_app/immutable/assets/img0-b4114dea.webp',
    title: 'Amazon Web Services (AWS)&hairsp;—&hairsp;Serverless Architecture',
    datatime: 'February 19, 2021',
    brief: 'A complete tour for beginners to set up a full-stack application',
    tags: [{
      name: 'aws',
      href: ''
    }, {
      name: 'cloud',
      href: ''
    }, {
      name: 'frontend',
      href: ''
    }]
  }, {
    img: 'https://programmingpercy.tech/_app/immutable/assets/img0-b4114dea.webp',
    title: 'Amazon Web Services (AWS)&hairsp;—&hairsp;Serverless Architecture, Amazon Web Services (AWS)&hairsp;—&hairsp;Serverless Architecture',
    datatime: 'February 19, 2021',
    brief: 'A complete tour for beginners to set up a full-stack application, A complete tour for beginners to set up a full-stack application, A complete tour for beginners to set up a full-stack application, A complete tour for beginners to set up a full-stack application',
    tags: [{
      name: 'aws',
      href: ''
    }, {
      name: 'cloud',
      href: ''
    }, {
      name: 'frontend',
      href: ''
    }]
  }];


  export default {
    components: {
    },
    data() {
      return {
        gap: 10,
        top: 10,
        column: 3,
        data: data.concat(data),
        formatDate: formatDate,
        page: {
          currentPage: 1,
          pageSize: 10,
        },
        articles: [],
      }
    },
    methods: {
      _GetArticleList() {
        Article.list(this.page).then(res => {
          this.articles = res.articleList.articles;
          this.page = Object.assign(this.page, res.articleList.page);
          this.$nextTick(() => {
            if (this.macy) {
              
            } else {
             // this.initSortItems();
            }
          })
        });
      },
      _InitSortItems() {
        const container = document.getElementById('container');
        const items = container.childNodes;
        const width = items[0].getBoundingClientRect().width;
        const poses = [{
          left: 0,
          top: 0,
          height: 0
        }, {
          left: 0,
          top: 0,
          height: 0
        }, {
          left: 0,
          top: 0,
          height: 0
        }];
        for (let i = 0; items.length > 1 && i < items.length; ++i) {
          const height = items[i].getBoundingClientRect().height;
          const currentIndex = i % 3;
          if (i > 2) {
            poses[currentIndex].top = poses[currentIndex].top + poses[currentIndex].height + 10; 
          } else {
            poses[currentIndex].left = (width + this.gap) * i; 
          }
          poses[currentIndex].height = height; 
          items[i].setAttribute('style', `left:${poses[currentIndex].left}px;top:${poses[currentIndex].top}px`);
        }
        poses.sort((a, b) => {
          return a.top + a.height > b.top + b.height;
        });

        document.getElementById('appMainId').setAttribute('style', `height:${poses[0].top + poses[0].height + 10}px`)
      }
    },

    
    mounted() {
      // this._GetArticleList();
      window.onload = this._InitSortItems;
    }
  }
</script>
<style scoped>
  .home {
  }

  .home-list {
    position: relative;
  }

  .home-item {
    background: rgba(31, 41, 55, 1);
    border-radius: 8px;
    cursor: pointer;
    left: 0;
    max-width: 32%;
    margin-bottom: 1.25rem;
    overflow: hidden;
    padding: 1.25rem;
    position: absolute;
    top: 0;
  }

  .home-item:hover {
    background: rgba(92, 99, 111, 1);
  }
  
  .home-item img {
    object-fit: cover;
    height: 100%;
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
