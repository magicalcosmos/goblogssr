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
                name: 'Home',
                path: '',
                component: () => import('@/views/Home')
              },
              {
                name: 'Blog',
                path: '/blog',
                component: () => import('@/views/Blog')
              },
              {
                name: 'BlogDetails',
                path: '/blog/:id',
                component: () => import('@/views/BlogDetails')
              },
              {
                name: 'Gallery',
                path: '/gallery',
                component: () => import('@/views/Gallery')
              },
              {
                name: 'Contact',
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
                name: 'Admin',
                path: '',
                component: () => import('@/views/admin'),
                meta: {
                  hasNav: false
                }
              },
              {
                name: 'Posts',
                path: '/admin/posts',
                component: () => import('@/views/admin/Posts'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'PostDetail',
                path: '/admin/posts/:id',
                component: () => import('@/views/admin/PostDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'Categories',
                path: '/admin/categories',
                component: () => import('@/views/admin/Categories'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'CategoryDetail',
                path: '/admin/categories/:id',
                component: () => import('@/views/admin/CategoryDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'Galleries',
                path: '/admin/galleries',
                component: () => import('@/views/admin/Galleries'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'GalleryDetail',
                path: '/admin/galleries/:id',
                component: () => import('@/views/admin/GalleryDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'Enquiries',
                path: '/admin/enquiries',
                component: () => import('@/views/admin/Enquiries'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'EnquiryDetail',
                path: '/admin/enquiries/:id',
                component: () => import('@/views/admin/EnquiryDetail'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'Users',
                path: '/admin/users',
                component: () => import('@/views/admin/Users'),
                meta: {
                  hasNav: true
                }
              },
              {
                name: 'UserDetail',
                path: '/admin/users/:id',
                component: () => import('@/views/admin/UserDetail'),
                meta: {
                  hasNav: true
                }
              }
            ]
          },
          {
            name: 'Login',
            path: '/admin/login',
            component: () => import('@/views/admin/Login')
          }
        ]
    });
}
