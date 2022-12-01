<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import { useNotificationStore } from "@/stores/notification";
import { dateFormat } from "@/tools/filters";

const notificationStore = useNotificationStore();

const loading = ref(false);

onMounted(async () => {
  await notificationStore.loadMessages();
  loading.value = false;
});
</script>

<template>
  <div v-loading="loading">
    <div v-if="notificationStore.messageList">
      <el-scrollbar height="400px">
        <div class="message_list">
          <el-card
            class="message"
            shadow="never"
            v-for="message in notificationStore.messageList"
            :key="message.title"
          >
            <template #header>
              <div class="card-header">
                <div class="message__title">
                  <div>{{ message.title }}</div>
                  <div>
                    {{ dateFormat(message.created_date, "DD.MM.YYYY HH:mm") }}
                  </div>
                </div>
              </div>
            </template>

            <div class="message__body">{{ message.message }}</div>
          </el-card>
        </div>
      </el-scrollbar>
    </div>
    <el-empty v-else description="Нет новых уведомлений" />
  </div>
</template>

<style lang="scss" scoped>
.message__title {
  display: flex;
  justify-content: space-between;
}

.message_list {
  display: grid;
  row-gap: 1rem;
}

.message__body {
  white-space: pre-line;
}
</style>
