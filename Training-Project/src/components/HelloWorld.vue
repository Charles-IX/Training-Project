<script setup>
import { ref } from 'vue'

const count = ref(0)
</script>

<template>
  <h1> Hello World. </h1>

  <div class="card">
    <button type="button" @mousedown="handleMouseDown" @mouseup="handleMouseUp">
      {{ buttonText }}
    </button>
    <p>What is Frex's favourite food?</p>
    <p>Hit the button and you'll know.</p>
  </div>

  <p class="copyright">Made with 💩 by Charles</p>
  <p class="copyright">Copyright 2024</p>
</template>

<style scoped>
.copyright {
  color: #888;
  line-height: 5px;
}
</style>

<script>
export default {
  data() {
    return {
      isPressing: false
    };
  },
  computed: {
    buttonText() {
      return this.isPressing ? '💩' : '😋';
    }
  },
  methods: {
    handleMouseDown() {
      this.isPressing = true;
    },
    handleMouseUp() {
      this.isPressing = false;
    }
  }
};
</script>
