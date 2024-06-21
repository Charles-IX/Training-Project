<template>

    <div class="container">
        <div class="logos">
            <a href="https://vitejs.dev" target="_blank">
                <img src="/vite.svg" class="logo" alt="Vite logo" />
            </a>
            <a href="https://vuejs.org/" target="_blank">
                <img src="../assets/vue.svg" class="logo vue" alt="Vue logo" />
            </a>
        </div>
        <a-card style="width: 320px; top: 0px; left: 7px" hoverable: true>
            <a-space size="15" align="center" direction="vertical" style="font-size: 16px">
                <!-- <a-avatar :size="64">
                <template #icon>
                    <UserOutlined />
                </template>
                </a-avatar> -->
                <h3 style="top: 20px; user-select: none;">C1小型汽车科目一模拟考试系统</h3>
                <div style="width: 100%;">
                    <a-input v-model:value="username" placeholder="用户名">
                        <template #prefix>
                            <UserOutlined />
                        </template>
                    </a-input>
                    <a-input-password v-model:value="password" placeholder="密码">
                        <template #prefix>
                            <KeyOutlined />
                        </template>
                    </a-input-password>
                </div>
                <p></p>
                <a-button @click="login" type="primary" style="font-size: 15px">
                    <template #icon>
                        <RightCircleFilled />
                    </template>
                    登录
                </a-button>
            </a-space>
        </a-card>
    </div>
</template>


<style scoped>
.logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
}

.logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
    filter: drop-shadow(0 0 2em #42b883aa);
}

.container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
}

.logos {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    min-height: 10vh;
}
</style>


<script>
import { defineComponent, ref } from 'vue';
import { UserOutlined, KeyOutlined, RightCircleFilled } from '@ant-design/icons-vue';
export default defineComponent({
    name: 'LoginUI',
    components: {
        UserOutlined,
        KeyOutlined,
        RightCircleFilled,
    },
    setup() {
        const username = ref('');
        const password = ref('');
        return {
            username,
            password,
        };
    },
    methods: {
        login() {
            if (this.username === 'admin' && this.password === 'admin') {
                const token = 'admin_logged_in';
                localStorage.setItem('token', token);
                this.$router.replace('/');
            }
            else if (this.username === 'user' && this.password === 'user') {
                const token = 'user_logged_in';
                localStorage.setItem('token', token);
                this.$router.replace('/');
            }
            else {
                alert('用户名或密码不正确。\n请检查你的用户名或密码。');
            }
        }
    }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>