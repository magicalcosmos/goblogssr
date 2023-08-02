import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);

export function createRouter () {
    return new Router({
        mode: 'hash',
        routes: [
          {
            path: '',
            component: () => import('@/components/Layout'),
            name: 'Layout',
            hidden: true,
            children: [
              {
                path: '',
                component: () => import('@/views/Home')
              },
              {
                path: '/blog',
                component: () => import('@/views/Blog')
              },
              {
                path: '/blog/:id',
                component: () => import('@/views/BlogDetails')
              },
              {
                path: '/gallery',
                component: () => import('@/views/Gallery')
              },
              {
                path: '/contact',
                component: () => import('@/views/Contact')
              }
            ]
          },
          {
            path: '/admin',
            component: () => import('@/components/Layout/AdminLayout'),
            name: 'AdminLayout',
            hidden: true,
            children: [
              {
                path: '',
                component: () => import('@/views/admin'),
                meta: {
                  hasNav: false
                }
              },
              {
                path: '/admin/posts',
                component: () => import('@/views/admin/Posts'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/posts/:id',
                component: () => import('@/views/admin/PostDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/categories',
                component: () => import('@/views/admin/Categories'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/categories/:id',
                component: () => import('@/views/admin/CategoryDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/category/:id',
                component: () => import('@/views/admin/CategoryDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/galleries',
                component: () => import('@/views/admin/Galleries'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/galleries/:id',
                component: () => import('@/views/admin/GalleryDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/enquiries',
                component: () => import('@/views/admin/Enquiries'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/enquiries/:id',
                component: () => import('@/views/admin/EnquiryDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/users',
                component: () => import('@/views/admin/Users'),
                meta: {
                  hasNav: true
                }
              },
              {
                path: '/admin/users/:id',
                component: () => import('@/views/admin/UserDetail'),
                meta: {
                  hasNav: true
                }
              }
            ]
          },
          {
            path: '/admin/login',
            component: () => import('@/views/admin/Login')
          }
        ]
    });
}
