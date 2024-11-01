<template>
  <div class="gvb_tabs">
    <span class="gvb_tab active" v-for="item in tabList" :key="item.name">{{ item.title }}</span>
    <span class="gvb_tab">用户列表
    <IconClose></IconClose>
    </span>
    <span class="gvb_tab">文章列表
    <IconClose></IconClose>
    </span>
    <span class="gvb_tab close_all_tab">全部关闭</span>
  </div>
</template>

<script setup lang="ts">
import {IconClose} from "@arco-design/web-vue/es/icon";
import {ref, type Ref} from "vue";
import {i} from "vite/dist/node/types.d-aGj9QkWt";

interface tabType {
  name: string
  title: string
}

const tabList: Ref<tabType[]> = ref([
  {name: "home", title: "首页"},
  {name: "user_info", title: "个人信息"},
])

</script>

<style scoped>
.gvb_tabs {
  width: 100%;
  height: 30px;
  border-bottom: 1px solid var(--bg);
  padding: 0 20px;
  display: flex;
  align-items: center;

  .gvb_tab {
    border-radius: 5px;
    border: 1px solid var(--bg);
    padding: 2px 8px;
    margin-right: 10px;
    cursor: pointer;
    position: relative;

    &.active {
      background-color: var(--active);
      color: white;
      border: none;
    }

    svg {
      font-size: 12px;
      border-radius: 50%;

      &:hover {
        background-color: var(--bg);
      }
    }
  }

  .close_all_tab {
    position: absolute;
    right: 20px;
    margin-right: 0;
  }
}
</style>