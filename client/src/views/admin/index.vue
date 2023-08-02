<template>
  <section>
    <h1>ShareUs</h1>
    <ul v-for="(item, index) in data" :key="`gg${index}`" class="panel-list">
      <li>
        <h2><i :class="`pi ${item.icon}`"></i>{{ item.title }}</h2>
      </li>
      <li v-if="item.children.length" class="clearfix">
        <ul>
          <li v-for="(k, i) in item.children" :key="`yy${i}`" class="tab-item clearfix">
            <div class="content">
              <a :href="k.href">
                <div class="title">{{ k.title }}</div>
                <div class="count">{{ k.count }}</div>
              </a>
            </div>
            <div class="add">
              <a :href="k.addUrl">
                <i class="pi pi-plus"></i>
              </a>
            </div>
          </li>
        </ul>
      </li>
    </ul>
  </section>
</template>
<script>
export default {
    data() {
      return {
        data: [
          {
            title: 'Posts',
            icon: 'pi-book',
            children: [
              {
                title: 'Posts',
                count: '20 items',
                href: '/#/admin/posts',
                addUrl: '/#/admin/posts?create'
              },
              {
                title: 'Post Galleries',
                count: '20 items',
                href: '/#/admin/categories',
                addUrl: '/#/admin/categories?create'
              }
            ]
          },
          {
            title: 'Galleries',
            icon: 'pi-images',
            children: [
              {
                title: 'Galleries',
                count: '20 items',
                href: '/#/admin/galleries',
                addUrl: '/#/admin/galleries?create'
              }
            ]
          },
          {
            title: 'Enquiries',
            icon: 'pi-envelope',
            children: [
              {
                title: 'Enquiries',
                count: '20 items',
                href: '/#/admin/enquiries',
                href: '/#/admin/enquiries',
                addUrl: '/#/admin/enquiries?create'
              }
            ]
          },
          {
            title: 'Users',
            icon: 'pi-users',
            children: [
              {
                title: 'Users',
                count: '20 items',
                href: '/#/admin/users',
                addUrl: '/#/admin/users?create'
              }
            ]
          }
        ]
      }
    }
  
}
</script>
<style scoped>
  h1 {
    border-bottom: 2px solid rgba(0,0,0, 0.1);
    color: #333;
    font-size: 28px;
    font-weight: 300;
    margin: 20px 0;
    padding-bottom: 10px;
  }

  .panel-list {
    margin-top: 20px;
  }

  .panel-list li:nth-child(2) {
    margin-top: 10px;
  }

  h2 {
    color: #999;
    font-size: 24px;
    font-weight: 500;
  }

  h2 i {
    font-size: 14px;
    margin-right: 10px;
    position: relative;
    top: -2px;
  }


  .content {
    float: left;
  }

  .content a {
    color: #1385e5;
    font-weight: 500;
    text-decoration: none;
  }

  .content .count {
    color: #999;
    font-size: 14px;
    margin-top: 3px;
    text-decoration: none;
    padding-bottom: .2rem;
  }

  .add {
    float: right;
  }

  .add i {
    background-color: #eee;
    border-radius: 3px;
    color: #fff;
    font-size: 14px;
    padding: 3px;
  }

  .tab-item {
    background-color: white;
    border-radius: .3em;
    box-shadow: 0 2px 3px rgba(0, 0, 0, 0.075), 0 0 0 1px rgba(0, 0, 0, 0.1);
    cursor: pointer;
    display: inline-block;
    margin-left: 22px;
    transition: box-shadow 80ms linear;
    padding: 10px;
    position: relative;
  }

  .tab-item:hover {
    box-shadow: 0 2px 3px rgba(19, 133, 229, 0.075), 0 0 0 1px rgba(19, 133, 229, 0.5);
  }

  .tab-item:hover .title {
    text-decoration: underline;
  }


  .tab-item:hover .add i {
    background-color: #dfdfdf;
  }

  .tab-item .add i:hover {
    background-color: #34c240;
  }

  @media (min-width: 768px) {
    .tab-item {
      width: 50%;
    }
  }

  @media (min-width: 992px) {
    .tab-item {
      width: 33%;
    }
  }

  @media (min-width: 1200px) {
    .tab-item {
      width: 25%;
    }
  }
</style>
