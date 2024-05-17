<template>
  <a-layout style="min-height: 100vh">
    <!-- 左则菜单栏 -->
    <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible class="menu-side">
      <a-space align="center" class="logo">
        <!-- <AppstoreTwoTone :style="{ fontSize: '35px' }" /> -->
        <!-- <span v-if="!collapsed" style="font-family: har_bold; font-size:20px; color:grey;">C1小型汽车科目一模拟考试系统</span> -->
      </a-space>
      <a-menu v-model:selectedKeys="selectedKeys" theme="light" mode="inline" style="font-family: pingfang_mem;">
        <a-menu-item key="1">
          <EditOutlined />
          <span>{{ menuName[1] }}</span>
        </a-menu-item>
        <a-menu-item key="2">
          <HistoryOutlined v-if="state === 'user'"/>
          <LogoutOutlined v-if="state === 'admin'"/>
          <span>{{ menuName[2] }}</span>
        </a-menu-item>
        <a-menu-item key="3" v-if="state === 'user'">
          <pie-chart-outlined />
          <span>{{ menuName[3] }}</span>
        </a-menu-item>
        <a-menu-item key="4" v-if="state === 'user'">
          <BookOutlined />
          <span>{{ menuName[4] }}</span>
        </a-menu-item>
        <a-menu-item key="5" v-if="state === 'user'">
          <LogoutOutlined />
          <span>{{ menuName[5] }}</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <!-- 右则导航页面 -->
    <a-layout>
      <!-- 顶部导航名 -->
      <a-layout-header style="background: #fff; padding: 0;">
        <a-space align="center">
          <menu-unfold-outlined v-if="collapsed" class="trigger" @click="() => (collapsed = !collapsed)" />
          <menu-fold-outlined v-else class="trigger" @click="() => (collapsed = !collapsed)" />
          <span class="top-name">{{ menuName[selectedKeys[0].valueOf()] }}</span>
        </a-space>
      </a-layout-header>
      <!-- 模块信息页面 -->
      <a-layout-content style="margin: 0 16px;">
        <Welcome v-if="selectedKeys[0] === '0'" />
        <Edit v-if="selectedKeys[0] === '1' && state === 'admin'" />
        <Exam v-if="selectedKeys[0] === '1' && state === 'user'" />
        <Logout v-if="selectedKeys[0] === '2' && state === 'admin'" />
        <History v-if="selectedKeys[0] === '2' && state === 'user'" />
        <Show v-if="selectedKeys[0] === '3'" />
        <Correction v-if="selectedKeys[0] === '4'" />
        <Logout v-if="selectedKeys[0] === '5'" />
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        <p>您目前以{{ state_ch }}身份登录。</p>
        Charles ©2024-2077
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import { defineComponent, ref } from 'vue';
import Welcome from './components/Welcome.vue';
import Edit from './components/Edit.vue';
import History from './components/History.vue';
import Show from './components/Show.vue';
import Correction from './components/Correction.vue';
import Logout from './components/Logout.vue';
import {EditOutlined, HistoryOutlined, PieChartOutlined, BookOutlined, LogoutOutlined, MenuFoldOutlined, MenuUnfoldOutlined, AppstoreTwoTone } from '@ant-design/icons-vue';
export default defineComponent({
  name: 'App',
  components: {
    Welcome,
    Edit,
    History,
    Show,
    Correction,
    Logout,
    EditOutlined,
    HistoryOutlined,
    PieChartOutlined,
    BookOutlined,
    LogoutOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    AppstoreTwoTone,
  },
  setup() {
    const collapsed = ref(false);
    const selectedKeys = ref(['0']);
    const state = ref('user');
    const state_ch = ref('用户')
    const menuName = ref(['C1小型汽车科目一模拟考试系统', '模拟考试', '历史成绩', '章节和专项练习', '错题本', '退出登录']);
    const token = localStorage.getItem('token');
    

    if (token === 'admin_logged_in') {
      state.value = 'admin';
      state_ch.value = '管理员';
      menuName.value = ['C1小型汽车科目一模拟考试系统', '题库管理', '退出登录'];
    }
    
    return {
      collapsed,
      selectedKeys,
      menuName,
      state,
      state_ch,
    }
  },
})

</script>

<style scoped>
.logo {
  display: flex;
  justify-content: center;
  margin-top: 10px;
  margin-bottom: 5px;
}

.trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

.trigger:hover {
  color: #1890ff;
}

.menu-side {
  background: #fff;
}

.top-name {
  font-family: har_bold;
  font-size: 26px;
}
</style>