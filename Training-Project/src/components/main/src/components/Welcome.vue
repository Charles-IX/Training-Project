<template>
    <div>
      <p></p>
      <h1 class="title">欢迎。</h1>
      <p class="paragraph">请从左侧栏选择你需要使用的功能。</p>
      <p class="small">愿有爱者依旧有爱，无爱者自由。☁️</p>
      <p></p>
      <p class="inline-content">如需技术支持，请访问
        <a href="https://charles-miku.top" target="_blank">Charles的个人主页</a>
        寻求帮助。
      </p>
    </div>
  </template>
  
  <script>
  export default {
    name: 'MyComponent'
  }
  </script>
  
  <style scoped>
  .title {
    font-size: 40px;
    color: navy;
  }
  
  .paragraph {
    font-size: 16px;
    font-weight: bold;
    
  }
  
  .small, .inline-content {
    font-size: 14px;
    color: gray;
  }
  </style>
  