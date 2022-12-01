<script setup lang="ts">
import { FormInstance } from "element-plus";
import { onMounted, reactive, ref } from "vue";
import type { ElForm } from "element-plus";
import { User as UserIcon } from "@element-plus/icons-vue";
import { useUserStore } from "@/stores/user";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";
import { Failure } from "@/internal/entity/failure";
const imageUrl = new URL("@/assets/images/login.png", import.meta.url).href;

const userStore = useUserStore();
const router = useRouter();
const toast = useToast();

const form = reactive({
  login: "",
  password: "",
});

const ruleFormRef = ref<FormInstance>();
const rules = reactive({
  login: [
    {
      required: true,
      message: "Укажите логин",
      trigger: "blur",
    },
  ],
  password: [
    {
      required: true,
      message: "Укажите пароль",
      trigger: "blur",
    },
  ],
});

const submitLogin = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async (valid) => {
    if (valid) {
      const res = await userStore.auth(form.login, form.password);
      res.match(
        async () => {
          await userStore.profile();
          router.push({ path: "/" });
        },
        async (err: Failure) => {
          toast.error(err.message);
        }
      );
    } else {
      return false;
    }
  });
};

onMounted(() => {});
</script>

<template>
  <div class="main__section container">
    <div class="login__container is__fullheight">
      <div class="login__box">
        <el-form
          ref="ruleFormRef"
          :model="form"
          label-position="top"
          :rules="rules"
        >
          <div class="form__image">
            <el-image :src="imageUrl" fit="contain"></el-image>
          </div>
          <el-form-item prop="login">
            <el-input
              v-model.trim="form.login"
              placeholder="Логин"
              :suffix-icon="UserIcon"
            >
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              type="password"
              v-model.trim="form.password"
              placeholder="Пароль"
              show-password
            >
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button @click="submitLogin(ruleFormRef)" type="primary"
              >Войти</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.form__image {
  display: flex;
  justify-content: center;
  margin-bottom: 0.7rem;
}
.login__container {
  display: flex;
  justify-content: center;
}

.login__box {
  background-color: #fff;
  border-radius: 6px;
  box-shadow: 0 0.5em 1em -0.125em #0a0a0a1a, 0 0 0 1px #0a0a0a05;
  color: #4a4a4a;
  display: block;
  padding: 1.25rem;
  width: 25rem;
}
</style>
