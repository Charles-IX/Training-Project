<template>
    <div>
        <a-table class="table" :columns="columns" :data-source="questionsData" row-key="id"
            :pagination="{ pageSize: 10 }" :scroll="{ y: scrollY }" />
        <!-- <a-modal title="题目详情" v-model:open="modalVisible" @ok="handleOk" @cancel="handleCancel"> -->
        <a-modal class="m-modal" v-model:open="modalVisible" @cancel="handleCancel" @ok="handleOk" ok-text="确认"
            cancel-text="取消" :closable="false" centered :destroyOnClose="true">
            <template #title>
                题目详情
            </template>
            <a-form style="margin-top: 20px;" autocomplete="off">
                <a-form-item label="题目">
                    <a-textarea v-model:value="sendData.Question" :defaultValue="sendData.Question"
                        :autoSize="{ minRows: 2, maxRows: 5 }"></a-textarea>
                </a-form-item>
                <a-form-item label="类型">
                    <a-radio-group v-model:value="sendData.Type" :autofocus="true">
                        <a-radio-button value='1'>单选题</a-radio-button>
                        <a-radio-button value='2'>判断题</a-radio-button>
                        <a-radio-button value='3'>多选题</a-radio-button>
                    </a-radio-group>
                </a-form-item>
                <a-form-item label="所属章节">
                    <a-input v-model:value="sendData.Chapter" :defaultValue="sendData.Chapter"></a-input>
                </a-form-item>
                <a-form-item label="答案">
                    <a-input v-model:value="sendData.Key" :defaultValue="sendData.Key"></a-input>
                </a-form-item>
                <a-form-item label="解析">
                    <a-textarea v-model:value="sendData.Answer" :defaultValue="sendData.Answer"
                        :autosize="{ minRows: 2, maxRows: 5 }"></a-textarea>
                </a-form-item>
                <a-form-item label="做错次数">
                    <a-input v-model:value="sendData.Times" :defaultValue="sendData.Times"></a-input>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script>
import { ref, onMounted, h } from 'vue';
import axios from 'axios';
import { Table, Modal, Button } from 'ant-design-vue';
import { EditOutlined } from '@ant-design/icons-vue';
import Antd from 'ant-design-vue';

export default {
    components: {
        'a-table': Table,
        'a-modal': Modal,
        'a-button': Button,
        'EditOutlined': EditOutlined,
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
        const modalVisible = ref(false);
        const selectedQuestion = ref({});

        const columns = [
            {
                title: '题号',
                dataIndex: 'id',
                key: 'id',
                width: 50,
            },
            {
                title: '题目',
                dataIndex: 'question',
                key: 'question',
            },
            {
                title: '类型',
                dataIndex: 'type',
                key: 'type',
                customRender: ({ text }) => {
                    return questionTypes[text] || '未知题型';
                },
                width: 50,
            },
            {
                title: '所属章节',
                dataIndex: 'chapter',
                key: 'chapter',
                width: 40,
            },
            {
                title: '答案',
                dataIndex: 'key',
                key: 'key',
                customRender: ({ text }) => {
                    return questionKeys[text] || '未知答案';
                },
                width: 30,
            },
            {
                title: '解析',
                dataIndex: 'answer',
                key: 'answer',
            },
            {
                title: '做错次数',
                dataIndex: 'times',
                key: 'times',
                width: 50,
            },
            {
                title: '操作',
                width: 40,
                key: 'action',

                customRender: ({ record }) => {
                    return h('a-button', {
                        onClick: () => {
                            console.log('Edit button clicked', record);
                            selectedQuestion.value = record;
                            fillSendData();
                            modalVisible.value = true;
                        }
                    }, [
                        h(EditOutlined)
                    ]);
                }
            }
        ];

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
                console.error('Error fetching data:', error);
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

        const handleOk = () => {
            submitData();
            modalVisible.value = false;
        };

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
            modalVisible,
            handleOk,
            handleCancel,
            EditOutlined,
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
</style>