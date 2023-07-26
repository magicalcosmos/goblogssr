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
                path: '/',
                component: () => import('@/views/home')
              },
              {
                path: '/:id/details',
                component: () => import('@/views/Details')
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
                component: () => import('@/views/admin')
              },
              {
                path: '/admin/posts',
                component: () => import('@/views/admin/Posts')
              },
              {
                path: '/admin/category',
                component: () => import('@/views/admin/Category')
              },
              {
                path: '/admin/login',
                component: () => import('@/views/admin/Login')
              }
            ]
          }
        ]
    });
}
