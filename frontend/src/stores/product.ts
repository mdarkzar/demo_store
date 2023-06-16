import { Failure } from "@/internal/entity/failure";
import { Product } from "@/internal/types/product";
import { defineStore } from "pinia";
import { useToast } from "vue-toastification";
import { siteAPI } from "./api";
import { err, ok, Result } from "neverthrow";

const toast = useToast();

interface loadAll {
  productList: Array<Product> | null;
}

interface loadProduct {
  product: Product | null;
}

interface create {
  product_id: number;
}

export const useProductStore = defineStore("product", {
  state: () => ({
    productList: null as Array<Product> | null,
    product: null as Product | null,
  }),
  getters: {},
  actions: {
    async loadAll() {
      const res = await siteAPI.get("/product/load");
      res.match(
        (r: loadAll) => {
          this.productList = r.productList;
        },
        (err: Failure) => {
          toast.error(err.message);
        }
      );
    },
    async create(name: string, price: number) {
      const res = await siteAPI.post("/product/create", {
        name: name,
        price: price,
      });
      return res.match(
        (r: create) => {
          toast.success(`Продукт №${r.product_id} успешно создан`);
          this.loadAll();
        },
        (e: Failure) => {
          toast.error(e.message);
        }
      );
    },
    async remove(productID: number) {
      const res = await siteAPI.post("/product/remove", {
        product_id: productID,
      });
      res.match(
        (_) => {
          toast.success(`Продукт №${productID} успешно удален`);
        },
        (err: Failure) => {
          toast.error(err.message);
        }
      );
    },
    async loadProduct(productID: number) {
      const res = await siteAPI.get(`/product/find/${productID}`);
      res.match(
        (r: loadProduct) => {
          this.product = r.product;
        },
        (err: Failure) => {
          toast.error(err.message);
        }
      );
    },
  },
});
