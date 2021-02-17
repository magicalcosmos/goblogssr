import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);

export function createRouter () {
    return new Router({
        mode: 'history',
        routes: [
            {path: '/', component: () => import('../views/Home')},
            {path: '/:id/details', component: () => import('../views/Details')},
            {path: '/counter', component: () => import('../views/Counter')},
            {path: '/admin/posts', component: () => import('../views/admin/Posts')},
            {path: '/admin/login', component: () => import('../views/admin/Login')}
        ]
    });
}