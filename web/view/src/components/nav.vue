<template>
    <n-card>
        <div class="profile">
            <n-avatar
            round
            :size="48"
            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
            />
            <p>xianren</p>
        </div>
        <div class="content">
            嘿嘿嘿
        </div>
      </n-card>
</template>

<script setup>

</script>

<style lang="scss" scoped>
.n-card {
    max-width: 300px;
  }
.profile {
    height: 150px;
    padding-top: 30px ;
    padding-left:50px;

  }
.content {
  text-align: center;
}
</style>