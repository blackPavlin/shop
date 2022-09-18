/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CreateProductRequest } from '../models/CreateProductRequest';
import type { Product } from '../models/Product';
import type { ProductList } from '../models/ProductList';
import type { UpdateProductRequest } from '../models/UpdateProductRequest';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class ProductService {

    /**
     * Получить список товаров
     * @returns ProductList successful operation
     * @throws ApiError
     */
    public static getProduct(): CancelablePromise<ProductList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/product',
        });
    }

    /**
     * Создать товар
     * @param requestBody 
     * @returns Product successful operation
     * @throws ApiError
     */
    public static postProduct(
requestBody?: CreateProductRequest,
): CancelablePromise<Product> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/product',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Получить детальную карточку товара
     * @param productId ID заказа
     * @returns Product successful operation
     * @throws ApiError
     */
    public static getProduct1(
productId: string,
): CancelablePromise<Product> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/product/{productId}',
            path: {
                'productId': productId,
            },
        });
    }

    /**
     * Изменить товар
     * @param productId ID заказа
     * @param requestBody 
     * @returns Product successful operation
     * @throws ApiError
     */
    public static patchProduct(
productId: string,
requestBody?: UpdateProductRequest,
): CancelablePromise<Product> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/product/{productId}',
            path: {
                'productId': productId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }

}
