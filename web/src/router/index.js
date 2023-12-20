import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue')
    },
    {
      path: '/products/',
      name: 'products',
      component: () => import('../views/ProductsView.vue')
    },
    {
      path: '/product/:id',
      name: 'product',
      component: () => import('../views/ProductView.vue')
    },
    {
      path: '/orders/new',
      name: 'ordersNew',
      component: () => import('../views/OrderNew.vue')
    },
    {
      path: '/orders/pay',
      name: 'orderPay',
      component: () => import('../views/OrderPay.vue')
    },
    {
      path: '/orders/',
      name: 'orders',
      component: () => import('../views/Orders.vue')
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/seller/productNew',
      name: 'productNew',
      component: () => import('../views/seller/ProductEdit.vue')
    },
    {
      path: '/seller/productEdit',
      name: 'productEdit',
      component: () => import('../views/seller/ProductEdit.vue')
    },
    {
      path: '/seller/productList',
      name: 'productList',
      component: () => import('../views/seller/ProductList.vue')
    },
    {
      path: '/seller/productPushEdit',
      name: 'productPushEdit',
      component: () => import('../views/seller/ProductPushEdit.vue')
    }
  ]
})

export default router
