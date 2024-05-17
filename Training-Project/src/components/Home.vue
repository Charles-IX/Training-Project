<template>
  <layout id="components-layout-demo-custom-trigger">
    <a-layout-sider
      v-model:collapsed="collapsed"
      collapsible
      trigger="null"
      breakpoint="lg"
      @collapse="handleCollapse"
    >
      <!-- 侧边栏内容 -->
      <div />
      <a-menu theme="dark" mode="inline" default-selected-keys="['1']">
        <a-menu-item key="1">
          <a-icon type="user" />
          <span>题库</span>
        </a-menu-item>
        <a-menu-item key="2">
          <a-icon type="user" />
          <span>另一个题库</span>
        </a-menu-item>
        <!-- 更多菜单项 -->
      </a-menu>
    </a-layout-sider>

    
      <a-layout-header style="background: #fff; padding: 0;">
        <!-- 顶栏内容 -->
        <div>C1小型汽车科目一模拟考试系统</div>
      </a-layout-header>

      <a-layout-content style="margin: 24px 16px; padding: 24px; background: #fff;">
        <!-- 主内容区 -->
        这里是一些需要展示在主页面上的内容。
      </a-layout-content>
    
  </layout>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false,
    };
  },
  methods: {
    handleCollapse(collapsed) {
      this.collapsed = collapsed;
    },
  },
};
</script>

<style scoped>
a-layout-header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 50px;
  background-color: #333;
  color: white;
  text-align: center;
  line-height: 50px;
  z-index: 1;
}

a-layout-sider {
  position: fixed;
  top: 50px;
  left: 0;
  width: 200px;
  height: calc(100vh - 50px);
  background-color: #444;
  color: white;
  z-index: 1;
}

a-layout-content {
  margin-left: 200px;
  margin-top: 50px;
  padding: 20px;
  height: calc(100vh - 50px);
  overflow: auto;
}

/* 响应式调整 */
@media (max-width: 768px) {
  a-layout-header {
    height: 40px;
    line-height: 40px;
  }

  a-layout-sider {
    width: 150px;
    top: 40px;
    height: calc(100vh - 40px);
  }

  a-layout-content {
    margin-left: 150px;
    margin-top: 40px;
  }
}

@media (max-width: 480px) {
  a-layout-sider {
    width: 100%;
    height: auto;
    position: fixed;
    top: 40px;
    padding: 10px;
  }

  a-layout-content {
    margin-left: 0;
    margin-top: 40px;
  }
}
</style>
