import { defineStore } from "pinia";
import { UploadImagesRequest } from "@/api/models/UploadImagesRequest";
import { ImageList } from "@/api/models/ImageList";
import { ImageService } from "@/api/services/ImageService";

export const useImageStore = defineStore("image", {
  actions: {
    async uploadImages(
      productId: string,
      formData: UploadImagesRequest,
    ): Promise<ImageList> {
      return ImageService.postProductImage(productId, formData);
    },
  },
});
