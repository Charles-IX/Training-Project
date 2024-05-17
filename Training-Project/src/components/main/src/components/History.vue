<template>
    <div>
        <a-table class="table" :columns="columns" :data-source="questionsData" row-key="id"
            :pagination="{ pageSize: 10 }" />
        <p></p>
        <p class="paragraph"> 前往错题本以查看错题和解析。 </p>
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
        const questionsData = ref([]);

        const modalVisible = ref(false);

        const columns = [
            {
                title: '成绩',
                dataIndex: 'score',
                key: 'score',
                width: 100,
            },
            {
                title: '日期',
                dataIndex: 'date',
                key: 'date',
                width: 300,
            },
            {
                title: '错题',
                dataIndex: 'mistakes',
                key: 'mistakes',
            },
        ];


        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8000/questions');
                questionsData.value = response.data;
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };


        return {
            questionsData,
            columns,
            EyeOutlined,
        };
    },
};
</script>

<style scoped>
.table {
    white-space: pre-wrap;
    /* 保持空白符和换行 */
}

.paragraph {
    font-size: 16px;
    font-weight: bold;
    
}
</style>