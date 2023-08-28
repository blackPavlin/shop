/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { ImageList } from '../models/ImageList';
import type { UploadImagesRequest } from '../models/UploadImagesRequest';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class ImageService {

    /**
     * Загрузка изображений товара
     * @param productId ID товара
     * @param formData 
     * @returns ImageList OK
     * @throws ApiError
     */
    public static postProductImage(
productId: string,
formData?: UploadImagesRequest,
): CancelablePromise<ImageList> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/product/{productId}/image',
            path: {
                'productId': productId,
            },
            formData: formData,
            mediaType: 'multipart/form-data',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Удаление изображения товара
     * @param productId ID товара
     * @param imageId ID изображения
     * @returns any OK
     * @throws ApiError
     */
    public static deleteProductImage(
productId: string,
imageId: string,
): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/product/{productId}/image/{imageId}',
            path: {
                'productId': productId,
                'imageId': imageId,
            },
            errors: {
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

}
