<template>
  <div class="gvb_menu">
    <a-menu
        @menu-item-click="clickMenu"
        v-model:selected-keys="selectedKeys"
        v-model:open-keys="openKeys"
    >
      <template v-for="item in menuList" :key="item.name">
        <a-menu-item :key="item.name" v-if="item.child?.length === 0">
          {{ item.title }}
          <template #icon>
            <component :is="item.icon"></component>
          </template>
        </a-menu-item>

        <a-sub-menu v-if="item.child?.length !== 0" :key="item.name">
          <template #icon>
            <component :is="item.icon"></component>
          </template>
          <template #title>{{ item.title }}</template>
          <a-menu-item :key="sub.name" v-for="sub in item.child">{{ sub.title }}</a-menu-item>
        </a-sub-menu>
      </template>
    </a-menu>
  </div>
</template>

<script setup lang="ts">

import {type Component, ref} from "vue";
import {IconMenu, IconUser} from "@arco-design/web-vue/es/icon";
import {type RouteMeta, useRoute, useRouter} from "vue-router";

interface MetaType extends RouteMeta{
  title: string;
}

const route = useRoute();
const router = useRouter();
interface MenuType {
  key: string;
  title: string;
  icon?: Component;
  name?: string; // 路由名字
  child?: MenuType[]
}

const menuList: MenuType[] = [
  {key: "1", title: "首页", icon: IconMenu, name: "home", child: []},
  {
    key: "2", title: "个人中心", icon: IconUser, name: "user_center", child: [
      {key: "2-1", title: "个人信息", icon: IconUser, name: "user_info", child: []},
    ]
  },
  {
    key: "3", title: "文章管理", icon: IconUser, name: "article", child: [
      {key: "3-1", title: "文章列表", icon: IconUser, name: "article_list", child: []},
    ]
  },
  {
    key: "4", title: "用户管理", icon: IconUser, name: "users", child: [
      {key: "4-1", title: "用户列表", icon: IconUser, name: "user_list", child: []},
    ]
  },
  {
    key: "5", title: "群聊管理", icon: IconUser, name: "chat_group", child: [
      {key: "5-1", title: "聊天记录", icon: IconUser, name: "chat_list", child: []},
    ]
  },
  {
    key: "6", title: "系统管理", icon: IconUser, name: "system", child: [
      {key: "6-1", title: "菜单列表", icon: IconUser, name: "menu_list", child: []},
    ]
  },
]

const selectedKeys = ref([route.name])
const openKeys = ref([route.matched[1].name])

function clickMenu(name: string) {
  router.push({
    name: name,
  })
}
</script>