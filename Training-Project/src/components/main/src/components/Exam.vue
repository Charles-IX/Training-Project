<template>
  <div>
    <div class="paragraph">每一套试卷包括 100 道题目，题型为判断题和单项选择题，每道题目 1 分，满分 100 分，试题随机来源于题库。</div>
    <div class="paragraph">每一次试卷练习时间规定为 45 分钟，超时系统会自动交卷结束考试。答题过程中错 11 分（ 11 道题）即终止本场考试。</div>
    <div class="paragraph">点击交卷后，系统会提供简单统计，比如得分，答对几道题，答错几道题，未答几道题，用时。</div>
    <a-button type="primary" @click="showModal" style="margin-top: 20px;">开始考试</a-button>
    <a-modal
      v-model:open="open"
      :closable="false"
      :keyboard="false"
      :footer="null"
      :title="`剩余 ${formattedTime} - 题目 ${currentQuestionIndex + 1}/100`"
      width="100%"
      wrap-class-name="full-modal"
      @ok="handleOk"
    >
      <div v-if="currentQuestion">
        <div style="white-space: pre-wrap;">{{ currentQuestion.question }}</div>
        <a-radio-group v-model:value="selectedAnswer" name="radioGroup">
          <a-radio v-for="(option, index) in options" :key="index" :value="index + 1">{{ option }}</a-radio>
        </a-radio-group>
        <!-- <p v-if="answerStatus !== null">{{ answerStatus ? '回答正确' : '回答错误' }}</p> -->
        <p v-if="answerStatus == true" style="color: green;">回答正确</p>
        <p v-if="answerStatus == false" style="color: red;">回答错误</p>
        <p v-if="notify == true" style="color: orange;">请选择答案</p>
      </div>
      <div style="margin-top: 20px;">
        <a-button type="default" @click="handleExit">退出考试</a-button>
        <a-button type="primary" @click="handleNext" style="margin-left: 10px;">下一题</a-button>
      </div>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from 'vue';
import axios from 'axios';

const open = ref(false);
const showModal = async () => {
  open.value = true;
  await fetchQuestions();
  startCountdown();
};

const handleOk = () => {
  open.value = false;
  countdown.value = 45 * 60;
  clearTimer();
};

const handleExit = () => {
  open.value = false;
  countdown.value = 45 * 60;
  clearTimer();
  window.location.reload();
};

const handleNext = async () => {
  if (selectedAnswer.value == null){
    notify.value = true;
  }
  else if (currentQuestionIndex.value <= 100) {
    notify.value = false;
    await submitAnswer();
    // 等待3秒钟后再切换到下一题
    setTimeout(() => {
      currentQuestionIndex.value++;
      answerStatus.value = null;
      selectedAnswer.value = null; // 重置选择的答案
    }, 1200);
  } else {
    alert('已经是最后一题');
  }
};

const value = ref('1');
const countdown = ref(45 * 60); // 45分钟，转换为秒
let timer = null;

const formattedTime = computed(() => {
  const minutes = Math.floor(countdown.value / 60);
  const seconds = countdown.value % 60;
  return `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
});

const formattedUsedTime = computed(() => {
  const minutesUsed = Math.floor((45*60 - countdown.value) / 60);
  const secondsUsed = (45*60 - countdown.value) % 60;
  return `${minutesUsed.toString().padStart(2, '0')}:${secondsUsed.toString().padStart(2, '0')}`;
});

const startCountdown = () => {
  if (timer) return;
  timer = setInterval(() => {
    if (countdown.value > 0) {
      countdown.value--;
    } else {
      clearTimer();
      alert('倒计时结束');
      handleExit();
    }
  }, 1000);
};

const clearTimer = () => {
  if (timer) {
    clearInterval(timer);
    timer = null;
  }
};

watch(open, (newVal) => {
  if (newVal) {
    startCountdown();
  } else {
    clearTimer();
  }
});

onUnmounted(() => {
  clearTimer();
});

const questions = ref([]);
const currentQuestionIndex = ref(0);
const currentQuestion = computed(() => questions.value[currentQuestionIndex.value] || {});
const selectedAnswer = ref(null);
const answerStatus = ref(null);
const notify = ref(false);

const fetchQuestions = async () => {
  try {
    const response = await axios.get('http://localhost:8000/getRandomQuestions');
    questions.value = response.data;
  } catch (error) {
    console.error('Error fetching questions:', error);
  }
};

const submitAnswer = async () => {
  try {
    const response = await axios.post('http://localhost:8000/answer', {
      ID: currentQuestion.value.id,
      Key: selectedAnswer.value
    });

    answerStatus.value = response.data.correct;
    if (response.data.terminate) {
      alert('已在本次考试中做错11题。\n考试终止，用时 ' + formattedUsedTime.value);
      handleExit();
    }
  } catch (error) {
    console.error('Error submitting answer:', error);
  }
};

const options = ['A', 'B', 'C', 'D']; // Assuming these are the only options for simplicity

</script>

<style lang="less">
.full-modal {
  .ant-modal {
    max-width: 100%;
    top: 0;
    padding-bottom: 0;
    margin: 0;
  }
  .ant-modal-content {
    display: flex;
    flex-direction: column;
    height: calc(100vh);
  }
  .ant-modal-body {
    flex: 1;
  }
}
</style>
