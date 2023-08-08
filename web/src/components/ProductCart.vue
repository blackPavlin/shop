<template>
  <el-card style="width: 500px">
    <template #header>
      <div class="card-header">
        <span>{{ product.name }}</span>
        <el-tag class="ml-2" type="success" effect="dark">
          {{ product.price }} ₽
        </el-tag>
      </div>
    </template>

    <el-carousel indicator-position="outside" :autoplay="false">
      <el-carousel-item v-if="!product.images.length">
        <el-image>
          <template #error>
            <div class="image-slot">
              <el-icon><icon-picture /></el-icon>
            </div>
          </template>
        </el-image>
      </el-carousel-item>
      <el-carousel-item v-for="image of product.images" :key="image.id">
        <el-image
          :src="`${objectStorageURL}/${image.name}`"
          fit="fill"
        ></el-image>
      </el-carousel-item>
    </el-carousel>

    <span> {{ product.description }}</span>
  </el-card>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Picture as IconPicture } from "@element-plus/icons-vue";
import type { PropType } from "vue";
import type { Product } from "@/api/models/Product";

export default defineComponent({
  name: "ProductCart",
  components: {
    IconPicture,
  },
  props: {
    product: {
      type: Object as PropType<Product>,
      required: true,
    },
  },
  setup() {
    const objectStorageURL = process.env.VUE_APP_IMAGE_STORAGE_URI;
    return {
      objectStorageURL,
    };
  },
});
</script>

<style lang="scss" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.el-image {
  padding: 0 5px;
  width: 100%;
  height: 100%;
}

.image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: var(--el-fill-color-light);
  color: var(--el-text-color-secondary);
  font-size: 30px;
}

.image-slot .el-icon {
  font-size: 30px;
}
</style>
