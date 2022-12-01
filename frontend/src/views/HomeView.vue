<script setup lang="ts">
import { useProductStore } from "@/stores/product";
import { useUserStore } from "@/stores/user";
import { numberFormatter } from "@/tools/filters";
import { FormInstance } from "element-plus";
import { onMounted, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import NotificationList from "@/components/NotificationList.vue";
const productStore = useProductStore();
const userStore = useUserStore();
const router = useRouter();

const productModal = ref(false);
const notificationModal = ref(false);
const loading = ref(false);

const form = reactive({
  name: "",
  price: "",
});

const ruleFormRef = ref<FormInstance>();
const rules = reactive({
  name: [
    {
      required: true,
      message: "Укажите название",
      trigger: "blur",
    },
  ],
  price: [
    {
      required: true,
      message: "Укажите стоимость",
      trigger: "blur",
    },
  ],
});

const createProduct = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async (valid) => {
    if (valid) {
      const price = +form.price;
      await productStore.create(form.name, price);
      productModal.value = false;
      form.name = "";
      form.price = "";
    } else {
      return false;
    }
  });
};

const deleteProduct = async (productID: number) => {
  await productStore.remove(productID);
  loadProductList();
};

const logout = async () => {
  await userStore.logout();
  router.push({ path: "/login" });
};

const loadProductList = async () => {
  loading.value = true;
  productStore.loadAll();
  loading.value = false;
};

onMounted(() => {
  loadProductList();
});
</script>

<template>
  <el-dialog
    v-model="productModal"
    destroy-on-close
    class="custom__modal"
    :fullscreen="false"
    title="Создание продукта"
  >
    <el-form
      ref="ruleFormRef"
      :model="form"
      label-position="top"
      :rules="rules"
    >
      <el-form-item prop="name">
        <el-input v-model="form.name" placeholder="Название"> </el-input>
      </el-form-item>
      <el-form-item prop="price">
        <el-input
          type="number"
          v-model.trim="form.price"
          placeholder="Стоимость"
        >
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="createProduct(ruleFormRef)" type="primary"
          >Создать продукт</el-button
        >
      </el-form-item>
    </el-form>
  </el-dialog>
  <el-dialog
    v-model="notificationModal"
    destroy-on-close
    class="custom__modal"
    :fullscreen="false"
    title="Уведомления"
  >
    <NotificationList></NotificationList>
  </el-dialog>
  <div class="main__section container">
    <el-card class="box-card">
      <div class="product__admin">
        <div class="control__btn1">
          <el-button @click="productModal = true">Создать продукт</el-button>
          <el-button @click="notificationModal = true">Уведомления</el-button>
        </div>
        <div class="control__btn2">
          <el-button @click="logout">Выход</el-button>
        </div>
      </div>
    </el-card>
    <div class="product__list" v-loading.fullscreen.lock="loading">
      <el-card
        class="box-card"
        v-for="product in productStore.productList"
        :key="product.id"
      >
        <template #header>
          <div class="card-header">
            <div>{{ product.name }}</div>
            <el-popconfirm
              title="Вы уверены, что хотите удалить?"
              @confirm="deleteProduct(product.id)"
              confirm-button-text="Да"
              cancel-button-text="Нет"
              width="230px"
            >
              <template #reference>
                <el-button class="button" type="danger" size="small"
                  >Удалить</el-button
                >
              </template>
            </el-popconfirm>
          </div>
        </template>
        <div>Стоимость: {{ numberFormatter(product.price) }}</div>
      </el-card>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.main__section {
  padding: 30px;
}

.product__admin {
  display: grid;
  margin-top: 0.2rem;
  width: 100%;

  grid-template-columns: repeat(2, 1fr);
  row-gap: 1rem;

  .control__btn2 {
    justify-self: end;
  }

  @media screen and (max-width: 1000px) {
    grid-template-columns: repeat(1, 1fr);

    .control__btn2 {
      justify-self: start;
    }

    .control__btn1 {
      display: grid;
      row-gap: 1rem;

      grid-template-columns: repeat(1, 1fr);
      .el-button {
        margin-left: 0px;
      }
    }
  }
}

.product__list {
  margin-top: 5rem;
  display: grid;
  gap: 2rem;
  grid-template-columns: repeat(4, 1fr);

  @media screen and (max-width: 1000px) {
    grid-template-columns: repeat(1, 1fr);
  }
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
