<template>
    <div>
        <div class="paragraph">每一套试卷包括100道题目，题型为判断题和单项选择题，每道题目1分，满分100分，试题随机来源于题库。</div>
        <div class="paragraph">每一次试卷练习时间规定为45分钟，超时系统会自动交卷结束考试。答题过程中错11分（11道题）即终止本场考试。</div>
        <div class="paragraph">点击交卷后，系统会提供简单统计，比如得分，答对几道题，答错几道题，未答几道题，用时。</div>
        <div> </div>
        <a-button type="primary" @click="showModal" style="margin-top: 20px;"> 开始考试 </a-button>
    </div>

</template>

<script>
import { ref, onMounted, h } from 'vue';
import axios from 'axios';
import { Table, Modal, Button } from 'ant-design-vue';
import { EyeOutlined } from '@ant-design/icons-vue';
import Antd from 'ant-design-vue';

export default {
    components: {
        'a-table': Table,
        'a-modal': Modal,
        'a-button': Button,
        'EyeOutlined': EyeOutlined,
    },
    setup() {
        const screenHeight = ref(window.innerHeight);
        const scrollY = ref(`${screenHeight.value * 0.55}px`);

        const updateTableHeight = () => {
            screenHeight.value = window.innerHeight;
            scrollY.value = `${screenHeight.value * 0.55}px`;
        };
        window.addEventListener('resize', updateTableHeight);

        const questionsData = ref([]);
        const questionTypes = {
            1: '单选题',
            2: '判断题',
            3: '多选题',
        };
        const questionKeys = {
            1: 'A',
            2: 'B',
            3: 'C',
            4: 'D',
        }
        const selectedQuestion = ref({});

        const open = ref(false);
        const showModal = () => {
        open.value = true;
        };

        const fillSendData = () => {
            sendData.value.ID = selectedQuestion.value.id;
            sendData.value.Question = selectedQuestion.value.question;
            sendData.value.Type = selectedQuestion.value.type;
            sendData.value.Chapter = selectedQuestion.value.chapter;
            sendData.value.Key = selectedQuestion.value.key;
            sendData.value.Answer = selectedQuestion.value.answer;
            sendData.value.Times = selectedQuestion.value.times;
        }

        const sendData = ref({
            ID: '',
            Question: '',
            Type: '',
            Chapter: '',
            Key: '',
            Answer: '',
            Times: '',
        });

        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8000/questions');
                questionsData.value = response.data;
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        };

        const submitData = async () => {
            try {
                const response = await axios.post('http://localhost:8000/response', JSON.stringify(sendData.value));
                console.log('Successfully submitted: ', response.data);
            } catch (error) {
                console.error('Failed to submit, ', error);
            }
        };

        // const handleOk = () => {
        //     submitData();
        //     modalVisible.value = false;
        // };

        const handleCancel = () => {
            modalVisible.value = false;
        };

        onMounted(() => {
            fetchData();
            updateTableHeight();
        });

        // onUnmounted(() => {
        //     window.removeEventListener('resize', updateTableHeight);
        // });

        return {
            scrollY,
            questionsData,
            columns,
            selectedQuestion,
            showModal,
            handleOk,
            handleCancel,
            EyeOutlined,
            fillSendData,
            sendData,
            submitData,
        };
    },
};
</script>

<style scoped>
.table {
    white-space: pre-wrap;
    /* 保持空白符和换行 */
}
.question {
    white-space: pre-wrap;
}
.paragraph {
    font-size: 16px;
    font-weight: bold;
    
}
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