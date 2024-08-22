import { useStore } from '@/stores/store'
import { message } from 'ant-design-vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path:"/login",
      name:"login",
      component:()=>import("../views/login.vue")
    },
    {
      path:"/admin",
      name:"admin",
      meta:{
        is_login:true,
      },
      component:()=>import("../views/admin/admin.vue"),
      children:[
        {
          path:"",
          name:"admin_index",
          redirect:"/admin/home"
        },
        {
          path:"user_center",
          name:"user_center",
          redirect:"/admin/user_center/user_info",
          children:[
            {
              path:"user_info",
              name:"user_info",
              component:()=>import("../views/admin/user_center/user_info.vue")
            },
            {
              path:"user_article_list",
              name:"user_article_list",
              component:()=>import("../views/admin/user_center/user_create_article_list.vue")
            },
            {
              path:"user_collects",
              name:"user_collects",
              component:()=>import("../views/admin/user_center/user_collects.vue")
            },
            {
              path:"user_messages",
              name:"user_messages",
              component:()=>import("../views/admin/user_center/user_message.vue")
            }
          ]
          
        },
        {
          path:"home",
          name:"home",
          component:()=>import("../views/admin/home/home.vue")
        },
        {
          path:"user_list",
          name:"user_list",
          component:()=>import("../views/admin/user_mgt/user_list.vue")
        },
        {
          path:"image_list",
          name:"image_list",
          component:()=>import("../views/admin/image_mgr/image_list.vue")
        },
        {
          path:"advert_list",
          name:"advert_list",
          component:()=>import("../views/admin/advert_mgr/advert_list.vue")
        },
        {
          path:"menu_list",
          name:"menu_list",
          component:()=>import("../views/admin/menu_mgr/menu_list.vue")
        },
        {
          path:"log_list",
          name:"log_list",
          component:()=>import("../views/admin/log_mgr/log_list.vue")
        },
        {
            path:"chat_list",
            name:"chat_list",
            component:()=>import("../views/admin/chat_list.vue")
        },
        {
          path:"message_list",
          name:"message_list",
          component:()=>import("../views/admin/message_list.vue")
        },
        {
          path:"tag_list",
          name:"tag_list",
          component:()=>import("../views/admin/tag_list.vue")
        },
        {
          path:"system",
          name:"system",
          children:[
            {
              path:"site",
              name:"system_site",
              component:()=>import("../views/admin/system_mgr/system_site.vue"),
            }
          ]
        },
        {
          path:"article_list",
          name:"article_list",
          component:()=>import("../views/admin/article_mgr/article_list.vue")
        },
        {
          path:"comment_list",
          name:"comment_list",
          component:()=>import("../views/admin/article_mgr/comment_list.vue")
        },
        {
          path:"add_article",
          name:"add_article",
          component:()=>import("../views/admin/article_mgr/add_article.vue")
        },
        {
          path:"edit_article/:id",
          name:"edit_article",
          component:()=>import("../views/admin/article_mgr/edit_article.vue")
        },
      ]
    },
  ]
})

export default router

// 路由的前置守卫
router.beforeEach((to, from, next) => {
  const store = useStore()
  
  if(to.meta.is_login && store.userInfo.role===0){
    // 未登录且要登录
    message.warn("需要登录")
    router.push({name:"login"})
    console.log("需要登录")
    return
  }
  // 放行
  next()
})