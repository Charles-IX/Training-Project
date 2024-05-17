<template>
    <div>
        <p></p>
        <h1 class="title"></h1>
        <p class="paragraph">你已注销，将在{{ countdown }}秒内跳转到登录页面。</p>
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
    data() {
        return {
            countdown: 5
        }
    },
    mounted() {
        localStorage.removeItem('token');

        this.startCountdown();
    },
    methods: {
        startCountdown() {
            const countdownInterval = setInterval(() => {
                if (this.countdown > 0) {
                    this.countdown--;
                } else {
                    clearInterval(countdownInterval)
                    this.$router.push('/Login');
                }
            }, 1000);
        }
    }
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

.small,
.inline-content {
    font-size: 14px;
    color: gray;
}
</style>