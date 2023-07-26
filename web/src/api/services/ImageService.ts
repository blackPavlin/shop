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

}
